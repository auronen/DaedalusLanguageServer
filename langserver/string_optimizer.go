package langserver

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var newConstants []string

func ExtractString(file string, line int, text, constName string) {
	lines, err := readLines(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
	edit := lines[line-1]
	if strings.Contains(edit, text) {
		edit = strings.Replace(edit, text, constName, -1)
		newConstants = append(newConstants, "const string "+constName+" = "+text+";")
		lines[line-1] = edit
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
