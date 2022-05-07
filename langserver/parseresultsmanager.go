package langserver

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

type parseResultsManager struct {
	parseResults map[string]*ParseResult
	mtx          sync.RWMutex
	logger       Logger
}

func newParseResultsManager(logger Logger) *parseResultsManager {
	return &parseResultsManager{
		parseResults: make(map[string]*ParseResult),
		logger:       logger,
	}
}

func (m *parseResultsManager) WalkGlobalSymbols(walkFn func(Symbol) error, types SymbolType) error {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, p := range m.parseResults {
		err := p.WalkGlobalSymbols(func(s Symbol) error {
			return walkFn(s)
		}, types)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *parseResultsManager) LookupGlobalSymbol(name string, types SymbolType) (Symbol, bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, p := range m.parseResults {
		if s, ok := p.LookupGlobalSymbol(name, types); ok {
			return s, true
		}
	}

	return nil, false
}

func (m *parseResultsManager) GetGlobalSymbols(types SymbolType) ([]Symbol, error) {
	result := make([]Symbol, 0, 200)

	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, p := range m.parseResults {
		p.WalkGlobalSymbols(func(s Symbol) error {
			result = append(result, s)
			return nil
		}, types)
	}

	return result, nil
}

func (m *parseResultsManager) Get(documentURI string) (*ParseResult, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	if r, ok := m.parseResults[documentURI]; ok {
		return r, nil
	}
	return nil, fmt.Errorf("document %q not found", documentURI)
}

func (m *parseResultsManager) Update(documentURI, content string) (*ParseResult, error) {
	r := m.ParseAndValidateScript(documentURI, content)

	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.parseResults[documentURI] = r
	return r, nil
}

func (m *parseResultsManager) resolveSrcPaths(srcFile, prefixDir string) []string {
	fileBytes, err := os.ReadFile(srcFile)
	if err != nil {
		return []string{}
	}
	decodedContent, err := charmap.Windows1252.NewDecoder().Bytes(fileBytes)
	if err != nil {
		return []string{}
	}
	srcContent := string(decodedContent)

	// Auronen: On Linux and macOS replace backslashes with forwardslashes
	if runtime.GOOS != "windows" {
		srcContent = strings.Replace(srcContent, "\\", string(os.PathSeparator), -1)
	}

	scanner := bufio.NewScanner(strings.NewReader(srcContent))

	resolvedPaths := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		dir := filepath.Dir(line)
		fname := filepath.Base(line)
		ext := strings.ToLower(filepath.Ext(fname))

		if ext == ".d" {
			if strings.Contains(fname, "*") {
				rxFilePath := regexp.MustCompile(strings.ReplaceAll(regexp.QuoteMeta(strings.ToLower(fname)), "\\*", ".*"))

				entries, err := os.ReadDir(filepath.Join(prefixDir, dir))
				if err != nil {
					continue
				}
				for _, e := range entries {
					if e.IsDir() {
						continue
					}
					if rxFilePath.MatchString(strings.ToLower(e.Name())) {
						absPath, err := filepath.Abs(filepath.Join(prefixDir, dir, e.Name()))
						if err != nil {
							continue
						}
						resolvedPaths = append(resolvedPaths, absPath)
					}
				}
			} else {
				absPath, err := filepath.Abs(filepath.Join(prefixDir, dir, fname))
				if err != nil {
					continue
				}
				resolvedPaths = append(resolvedPaths, absPath)
			}
		} else if ext == ".src" {
			m.logger.Infof("Collecting scripts from %q", filepath.Join(prefixDir, dir, fname))
			resolvedPaths = append(resolvedPaths, m.resolveSrcPaths(filepath.Join(prefixDir, line), filepath.Join(prefixDir, dir))...)
		}
	}

	return resolvedPaths
}

var (
	bufferPool  = sync.Pool{New: func() interface{} { b := new(bytes.Buffer); b.Grow(2048); return b }}
	decoderPool = sync.Pool{New: func() interface{} { return charmap.Windows1252.NewDecoder() }}
)

func (m *parseResultsManager) validateFiles(resolvedPaths []string) map[string][]SyntaxError {
	results := make(map[string][]SyntaxError)

	chanPaths := make(chan string, len(resolvedPaths))
	for _, r := range resolvedPaths {
		chanPaths <- r
	}
	close(chanPaths)

	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()
	if numWorkers > 2 {
		numWorkers = numWorkers / 2
	}
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			buf := bufferPool.Get().(*bytes.Buffer)
			defer bufferPool.Put(buf)

			decoder := decoderPool.Get().(*encoding.Decoder)
			defer decoderPool.Put(decoder)

			for r := range chanPaths {
				f, err := os.Open(r)
				if err != nil {
					continue
				}
				decoder.Reset()
				translated := decoder.Reader(f)
				buf.Reset()
				_, err = buf.ReadFrom(translated)
				f.Close()
				if err != nil {
					continue
				}

				parsed := m.ValidateScript(r, buf.String())
				if len(parsed) > 0 {
					m.mtx.Lock()
					results[r] = parsed
					m.mtx.Unlock()
				}
			}
		}(&wg)
	}

	wg.Wait()
	return results
}

func (m *parseResultsManager) ParseSource(srcFile string) ([]*ParseResult, error) {
	resolvedPaths := m.resolveSrcPaths(srcFile, filepath.Dir(srcFile))
	m.logger.Infof("Parsing %q. This might take a while.", srcFile)

	results := make([]*ParseResult, 0, len(resolvedPaths))

	chanPaths := make(chan string, len(resolvedPaths))
	for _, r := range resolvedPaths {
		chanPaths <- r
	}
	close(chanPaths)

	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()
	if numWorkers > 2 {
		numWorkers = numWorkers / 2
	}
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			buf := bufferPool.Get().(*bytes.Buffer)
			defer bufferPool.Put(buf)

			decoder := decoderPool.Get().(*encoding.Decoder)
			defer decoderPool.Put(decoder)

			for r := range chanPaths {
				f, err := os.Open(r)
				if err != nil {
					continue
				}
				decoder.Reset()
				translated := decoder.Reader(f)
				buf.Reset()
				_, err = buf.ReadFrom(translated)
				f.Close()
				if err != nil {
					continue
				}

				parsed := m.ParseScript(r, buf.String())

				m.mtx.Lock()
				m.parseResults[parsed.Source] = parsed
				results = append(results, parsed)
				m.mtx.Unlock()
			}
		}(&wg)
	}

	wg.Wait()

	validations := m.validateFiles(resolvedPaths)

	for _, v := range results {
		if e, ok := validations[v.Source]; ok {
			v.SyntaxErrors = append(v.SyntaxErrors, e...)
		}
	}

	m.logger.Infof("Done parsing %q: %d scripts.", srcFile, len(results))
	return results, nil
}
