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

	SVMs could be sorted by the order from the file
	AI_Output - added and sorted with the relevant string literals from C_INFO instances (descriptions)
	Other String literals - C_ITEM instances - to put the in neat files
	String literals - the rest just dump into one big file and order them by file, then by line. Probably not needed to do anything else, majority strings are Dialogues, items, random log entires and the rest

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

func newCSVentryFromDialogue(dia Dialogue) CSVentry {
	// TODO: Do not ignore the error!
	file, _ := getRepoRelativePath(dia.sourceFile)
	return CSVentry{
		path:       file,
		lineNumber: dia.line,

		location:            file + ":" + fmt.Sprint(dia.line),
		source:              trimQuotes(dia.text),
		target:              "",
		id:                  "",
		fuzzy:               "",
		context:             trimQuotes(dia.soundName),
		translator_comments: "",
		developer_comments:  "",
	}
}

func newCSVentryFromConstArraySymbol(sym ConstantArraySymbol) []CSVentry {
	entries := make([]CSVentry, 0, 100)
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
	return entries
}

func (e CSVentry) getValue() []string {
	return []string{e.location,
		e.source,
		e.target,
		e.id,
		e.fuzzy,
		e.context,
		e.translator_comments,
		e.developer_comments}
}

// header for the CSV files
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

// Generates CSV containing all dialogues
// TODO: add C_INFO.description entries
func (h *LspHandler) generateDialogueCSV() {
	entries := make( /*[][]string*/ []CSVentry, 0, 200)

	for _, res := range h.parsedDocuments.parseResults {
		for _, dia := range res.GlobalDialogues {
			entries = append(entries, newCSVentryFromDialogue(dia))
		}
	}

	// Sort entries by file and then by line (ensures comfortable translations)
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

	h.logger.Infof("Got %d dialogues", len(entries))
	csvWrite(data, "DIA", "de")
}

// Generates CSV containing all constStrings
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
				entries = append(entries, newCSVentryFromConstArraySymbol(sym)...)
			}
		}
	}

	// Sort entries by file and then by line (ensures comfortable translations)
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

/*
func decodeString(s string) string {
	// the string we want to transform
	s := "今日は"
	fmt.Println(s)

	// --- Encoding: convert s from UTF-8 to ShiftJIS
	// declare a bytes.Buffer b and an encoder which will write into this buffer
	var b bytes.Buffer
	wInUTF8 := transform.NewWriter(&b, charmap.Windows1252.NewEncoder())
	// encode our string
	wInUTF8.Write([]byte(s))
	wInUTF8.Close()
	// print the encoded bytes
	fmt.Printf("%#v\n", b)
	encS := b.String()
	fmt.Println(encS)

	// --- Decoding: convert encS from ShiftJIS to UTF8
	// declare a decoder which reads from the string we have just encoded
	rInUTF8 := transform.NewReader(strings.NewReader(encS), japanese.ShiftJIS.NewDecoder())
	// decode our string
	decBytes, _ := ioutil.ReadAll(rInUTF8)
	decS := string(decBytes)
	fmt.Println(decS)
}
*/
