package langserver

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"unicode"
)

// Generates case insensitive pattern used with the filepat Glob() funciton
// source: https://wenzr.wordpress.com/2018/04/09/go-glob-case-insensitive/
func InsensitivePattern(path string) string {
	if runtime.GOOS == "windows" {
		return path
	}
	p := strings.Builder{}
	for _, r := range path {
		if unicode.IsLetter(r) {
			fmt.Fprintf(&p, "[%c%c]", unicode.ToLower(r), unicode.ToUpper(r))
		} else {
			p.WriteRune(r)
		}
	}
	return p.String()
}

func findPaths(file string) ([]string, error) {
	abs, err := filepath.Abs(file)
	if err != nil {
		return nil, err
	}

	pattern := InsensitivePattern(abs)
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	if len(matches) == 0 {
		return nil, os.ErrNotExist
	}
	for i, p := range matches {
		p_s, err := filepath.EvalSymlinks(p)
		if err == nil {
			matches[i] = p_s
		} else {
			return nil, fmt.Errorf("could not eval symlink: %s", err)
		}
	}

	return matches, nil
}

func findPath(file string) (string, error) {
	matches, err := findPaths(file)
	if err != nil {
		return "", err
	}
	if len(matches) == 0 {
		return "", os.ErrNotExist
	}
	return matches[0], nil
}

func findPathAnywhereUpToRoot(dir, segment string) (string, error) {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	p, err := findPath(filepath.Join(dir, segment))
	if err == nil {
		return p, nil
	}

	nextDir := filepath.Dir(absDir)
	if nextDir == absDir {
		return "", fmt.Errorf("could not find %q anywhere", segment)
	}
	return findPathAnywhereUpToRoot(nextDir, segment)
}

func insensitiveRxPattern(path string) string {
	p := strings.Builder{}
	for _, r := range path {
		if unicode.IsLetter(r) {
			fmt.Fprintf(&p, "[%c%c]", unicode.ToLower(r), unicode.ToUpper(r))
		} else if r == '*' {
			p.WriteRune('.')
			p.WriteRune(r)
		} else {
			p.WriteString(regexp.QuoteMeta(string(r)))
		}
	}
	return p.String()
}

func ResolvePathsCaseInsensitive(directory, fileName string) (files []string, err error) {
	containingDir, err := findPath(directory)
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(containingDir)
	if err != nil {
		return nil, err
	}
	pattern := insensitiveRxPattern(fileName)
	rxName, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	for _, file := range entries {
		if rxName.MatchString(file.Name()) {
			files = append(files, filepath.Join(containingDir, file.Name()))
		}
	}

	return files, nil
}
