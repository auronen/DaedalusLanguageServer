/*
	Functionality related to the translation feature
*/
package langserver

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/integralist/go-findroot/find"
)

/* What has to be done
- SVMs
- string constants
- string literals (don't include irrelevant strings - blacklist for functions maybe?)
- AI_Output line comments - dialogues

Order of operations:
	I think it is only needed to parse and put all the relevant data by parsing
	in the order of the .src file
*/

// "location","source","target","id","fuzzy","context","translator_comments","developer_comments"
type CSVentry struct {
	location            string // {git repo relative path}:{line number}
	source              string // Source string
	target              string // Translated string - empty
	id                  string // ?
	fuzzy               string // ?
	context             string // when const - symbol name, when SVM - name, when dia - sound file
	translator_comments string // empty
	developer_comments  string // Maybe we could parse something here?
	// for sorting
	path       string
	lineNumber int
}

func newCSVentry(loc, src, trgt, id, fuzz, ctx, tr_c, dev_c, pth string, lN int) CSVentry {
	return CSVentry{
		location:            loc,
		source:              src,
		target:              trgt,
		id:                  id,
		fuzzy:               fuzz,
		context:             ctx,
		translator_comments: tr_c,
		developer_comments:  dev_c,
		path:                pth,
		lineNumber:          lN,
	}
}

func trimQuotes(input string) string {
	return strings.Trim(input, "\"")
}

func newCSVentryFromConstSymbol(symbol ConstantSymbol) CSVentry {
	line := symbol.Definition().Start.Line
	// TODO: Do not ignore the error!
	file, _ := getRepoRelativePath(symbol.Source())
	return CSVentry{
		path:       file,
		lineNumber: line,

		location:            file + ":" + fmt.Sprint(line),
		source:              trimQuotes(symbol.Value),
		target:              "",
		id:                  "",
		fuzzy:               "",
		context:             symbol.NameValue,
		translator_comments: "",
		developer_comments:  "",
	}
}

func (e CSVentry) getValue() []string {
	return []string{e.location, e.source, e.target, e.id, e.fuzzy, e.context, e.translator_comments, e.developer_comments}
} // header for the CSV files
var CSVHeader CSVentry = newCSVentry("location", "source", "target", "id", "fuzzy", "context", "translator_comments", "developer_comments", "", 0)

// Returns absolute path of the git repo
func findRepoRoot() (string, error) {
	root, err := find.Repo()
	if err != nil {
		return "", err
	}
	return root.Path, nil
}

// Returns relative path of files inside the git repo
//
// takes absolute path of a file and returns path relative to git repo
func getRepoRelativePath(path string) (string, error) {
	repoRoot, err := findRepoRoot()
	if err == nil {
		return strings.ReplaceAll(path, repoRoot, "")[1:], err
	}
	return "", err
}

func csvWrite(data [][]string, filename, lang string) error {
	// create path
	repoRoot, err := findRepoRoot()
	if err != nil {
		return err
	}
	file, err := os.Create(filepath.Join(repoRoot, "translations", lang, filename+".csv"))
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		if err := writer.Write(value); err != nil {
			return err
		}
	}
	return nil
}

func (h *LspHandler) generateConstStringCSV() {
	entries := make( /*[][]string*/ []CSVentry, 0, 200)
	entries = append(entries, CSVHeader /*.getValue()*/) // add header
	consts, err := h.parsedDocuments.GetGlobalSymbols(SymbolConstant)
	if err != nil {
		return
	}
	for _, c := range consts {
		sym, ok := c.(ConstantSymbol)
		if ok {
			if strings.EqualFold(sym.Type, "string") {
				line := newCSVentryFromConstSymbol(sym)
				if line.source == "" {
					continue
				}
				entries = append(entries, line /*.getValue()*/)
			}
		}
	}
	constsArr, err := h.parsedDocuments.GetGlobalSymbols(SymbolConstantArray)
	if err != nil {
		return
	}
	for _, c := range constsArr {
		sym, ok := c.(ConstantArraySymbol)
		if ok {
			if strings.EqualFold(sym.Type, "string") {
				// TODO: Do not ignore the error!
				file, _ := getRepoRelativePath(sym.Source())
				parentName := sym.Name()
				for i, el := range sym.Values {
					lineNumber := el.Definition.Start.Line
					line := newCSVentry(file+":"+fmt.Sprint(lineNumber),
						trimQuotes(el.Value),
						"",
						"",
						"",
						parentName+"["+fmt.Sprint(i)+"]",
						"",
						"",
						file,
						lineNumber,
					)
					if line.source == "" {
						continue
					}
					entries = append(entries, line /*.getValue()*/)

				}
			}
		}
	}

	//sort.Sort(SortByLineNumber(entries))
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].path != entries[j].path {
			return entries[i].path < entries[j].path
		}
		return entries[i].lineNumber < entries[j].lineNumber
	})

	data := make([][]string, 0, len(entries))
	for _, l := range entries {
		data = append(data, l.getValue())
	}

	h.logger.Infof("Got %d string constants", len(entries))
	csvWrite(data, "Const", "de")
}
