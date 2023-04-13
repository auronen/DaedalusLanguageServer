package langserver

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sort"
)

// simpleCSVentry struct holds the id and the text
//
// this is used for simple CSV weblate format
type simpleCSVentry struct {
	context string
	text    string

	// for sorting
	pth     string
	ln      int
}

func (c *simpleCSVentry) GetEscapedCSV() string {
	return "\"" + c.context + "\",\"" + c.text + "\""
}

func newSimpleCSVentryFromStringLocations(id string, sl SymbolPosition) simpleCSVentry {
	if sl.quotes {
		return simpleCSVentry{
			context: id,
			text:    sl.content,
			pth:     sl.document,
			ln:      sl.line,
		}
	} else {
		return simpleCSVentry{
			context: id,
			text:    strings.TrimPrefix(sl.content, "//"),
			pth:     sl.document,
			ln:      sl.line,
		}
	}
}

// CSVentry represents one entry in the complex tranlsation csv file
type CSVentry struct {
	location            string // {git repo relative path}:{line number}
	source              string // source string
	target              string // translated string - empty
	id                  string // ? internal
	fuzzy               string // ? internal
	context             string // string ID
	translator_comments string // empty
	developer_comments  string //

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


func newCSVentryFromStringLocation(id string, sl SymbolPosition) CSVentry {
	if sl.quotes {
		return CSVentry{
			location:            sl.document + ":" + fmt.Sprintf("%d",sl.line),
			source:              trimQuotes(sl.content),
			target:              trimQuotes(sl.content),
			id:                  "",
			fuzzy:               "",
			context:             id,
			translator_comments: sl.translation_comment,
			developer_comments:  "",
			path:                sl.document,
			lineNumber:          sl.line,
		}
	} else {
		return CSVentry{
			location:            sl.document + ":" + fmt.Sprintf("%d",sl.line),
			source:              strings.TrimPrefix(sl.content, "//"),
			target:              strings.TrimPrefix(sl.content, "//"),
			id:                  "",
			fuzzy:               "",
			context:             id,
			translator_comments: sl.translation_comment,
			developer_comments:  "",
			path:                sl.document,
			lineNumber:          sl.line,
		}
	}
}

func (c *CSVentry) GetEscapedCSV() string {
	src := c.source
	if strings.Contains(c.source, "\"") {
		src = strings.ReplaceAll(src, "\"", "'")
	}
	trgt := c.target
	if strings.Contains(c.target, "\"") {
		trgt = strings.ReplaceAll(trgt, "\"", "'")
	}
	return "\"" + c.location + "\",\"" +
		src + "\",\"" +
		trgt + "\",\"" +
		c.id + "\",\"" +
		c.fuzzy + "\",\"" +
		c.context + "\",\"" +
		c.translator_comments + "\",\"" +
		c.developer_comments + "\""
}

func trimQuotes(input string) string {
	return strings.Trim(input, "\"")
}

/*
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

func newCSVentryFromStringLiteral(s StringLiteral) CSVentry {
	file, _ := getRepoRelativePath(s.sourceFile)
	return CSVentry{
		path:       file,
		lineNumber: s.line,

		location:            file + ":" + fmt.Sprint(s.line),
		source:              trimQuotes(s.text),
		target:              "",
		id:                  "",
		fuzzy:               "",
		context:             trimQuotes(s.context),
		translator_comments: s.trComment,
		developer_comments:  s.devComment,
	}
}

func newCSVentryFromSVM(s SVM) CSVentry {
	file, _ := getRepoRelativePath(s.sourceFile)
	return CSVentry{
		path:       file,
		lineNumber: s.line,

		location:            file + ":" + fmt.Sprint(s.line),
		source:              trimQuotes(s.text),
		target:              "",
		id:                  "",
		fuzzy:               "",
		context:             trimQuotes(s.context),
		translator_comments: "",
		developer_comments:  "",
	}
}

func newCSVentryFromConstArraySymbol(sym symbol.ConstantArray) []CSVentry {
	entries := make([]CSVentry, 0, 100)
	file, _ := getRepoRelativePath(sym.Source())
	parentName := sym.Name()
	for i, el := range sym.Elements {
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
		entries = append(entries, line )
	}
	return entries
}
*/
// TODO - Union plugin translations

/*
func parseText(in string) string {
	if strings.Contains(in, "\"") {
		return strings.ReplaceAll(in, "\"", "")
	}
	return "" // if it is not a string literal, return empty string
}

func newsimpleCSVentryTranslationStringEntry(s TranslationStringEntry, id int) simpleCSVentry {
	var txt string
	switch id {
	case 0:
		txt = parseText(s.RUS)
	case 1:
		txt = parseText(s.ENG)
	case 2:
		txt = parseText(s.GER)
	case 3:
		txt = parseText(s.POL)
	case 4:
		txt = parseText(s.ROU)
	case 5:
		txt = parseText(s.ITA)
	case 6:
		txt = parseText(s.CZE)
	case 7:
		txt = parseText(s.ESP)
	}
	return simpleCSVentry{
		context: s.name,
		text:    txt,
	}
}
*/

// write the simple CSVentry slice as an escaped CSV file with the filename: `filename`_`lang`.csv
func simpleCsvWrite(entries []simpleCSVentry, filename, lang string) error {
	// create path
	repoRoot, err := findRepoRoot()
	if err != nil {
		return err
	}
	// _ = os.MkdirAll(filepath.Join(repoRoot, ".translations", lang), os.ModePerm)
	// file, err := os.Create(filepath.Join(repoRoot, ".translations", lang, filename+"_"+lang+".csv"))
	_ = os.MkdirAll(filepath.Join(repoRoot, ".translations"), os.ModePerm)
	file, err := os.Create(filepath.Join(repoRoot, ".translations", filename + ".csv"))
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	for _, entry := range entries {
		writer.WriteString(entry.GetEscapedCSV() + "\n")
	}
	return nil
}

// header for the CSV files
//                                   "location", "source", "target", "id", "fuzzy", "context", "translator_comments", "developer_comments"
var CSVHeader CSVentry = newCSVentry("location", "source", "target", "id", "fuzzy", "context", "translator_comments", "developer_comments", "", 0)

// write the CSVentry slice as an escaped CSV file with the filename: `filename`_`lang`.csv
func csvWrite(entries []CSVentry, filename, lang string) error {
	// create path
	repoRoot, err := findRepoRoot()
	if err != nil {
		return err
	}
	_ = os.MkdirAll(filepath.Join(repoRoot, ".translations"), os.ModePerm)
	file, err := os.Create(filepath.Join(repoRoot, ".translations", filename + ".csv"))
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	writer.WriteString("\"location\",\"source\",\"target\",\"id\",\"fuzzy\",\"context\",\"translator_comments\",\"developer_comments\"\n")

	for _, entry := range entries {
		writer.WriteString(entry.GetEscapedCSV() + "\n")
	}
	return nil
}

// Read all csv files from a specified language and return all of thes as TranslatedString array
func csvReadLanguage(lang string) (translatedStrings []TranslatedString, err error) {
	// create path
	repoRoot, err := findRepoRoot()
	if err != nil {
		return
	}
	path := filepath.Join(repoRoot, ".translations", lang)

	//fmt.Fprintf(os.Stderr, "%s\n", path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured while reading files: %s\n", err)
	}

	translatedStrings = make([]TranslatedString, 0, 200)

	for _, f := range files {
		// skip subdirs
		if f.IsDir() {
			continue
		}
		file, err := os.Open(filepath.Join(path, f.Name()))

		if err != nil {
			break
		}
		defer file.Close()

		csvReader := csv.NewReader(file)

		records, er := csvReader.ReadAll()
		if er != nil {
			break
		}

		if len(records) > 2 { // getting rid of the header - like this because some types of csvs do not have the header
			if strings.EqualFold(records[0][5], "context") && strings.EqualFold(records[0][2], "target") {
				_, records = records[0], records[1:]
			}
		}

		for _, r := range records {
			// "location","source","target","id","fuzzy","context","translator_comments","developer_comments"
			translatedStrings = append(translatedStrings, newTranslatedString(r[5], r[2]))
		}

	}
	return
}

func csvGetLanguages() (languages []string, err error) {
	// create path
	repoRoot, err := findRepoRoot()
	if err != nil {
		return
	}
	path := filepath.Join(repoRoot, ".translations")

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	langs := make([]string, 0)

	fmt.Fprintf(os.Stderr, "Printing files: ")

	for _, f := range files {
		if f.IsDir() {
			langs = append(langs, f.Name())
		}
	}
	return langs, nil
}

// Generates one giant escaped csv file with all translatable strings
func (h *LspHandler) generateAllsimpleCSV() {
	var what string
	var ws_key string
	for key, val := range h.workspaces {
		if val.wsID == GOTHIC {
			what = val.wsID
			ws_key = key
		} else if val.wsID == MENU {
			what = val.wsID
			ws_key = key
		}
	}

	h.logger.Infof("Generating %s", what)

	entries := make([]simpleCSVentry, 0, 200)

	for _, res := range h.workspaces[ws_key].parsedDocuments.parseResults {
		for content, pos_array := range res.StringLocations {
			if len(pos_array) > 0 {
				entries = append(entries, newSimpleCSVentryFromStringLocations(content, pos_array[0]))
			}
		}
	}

	h.logger.Infof("Got %d lines", len(entries))

	simpleCsvWrite(entries, what, "translations")
}







// Generates one giant escaped csv file with all translatable strings
func (h *LspHandler) generateAllCSV() {
	for _, ws := range h.workspaces {
		h.logger.Infof("Generating: %s", ws.wsID)
		entries := make([]CSVentry, 0, 2000)

		for _, res := range ws.parsedDocuments.parseResults {
			for content, pos_array := range res.StringLocations {
				if len(pos_array) > 0 {
					entries = append(entries, newCSVentryFromStringLocation(content, pos_array[0]))
				}
			}
		}

		h.logger.Infof("Got %d lines", len(entries))

		// Sort entries by file and then by line (ensures comfortable translations)
		sort.Slice(entries, func(i, j int) bool {
			if entries[i].path != entries[j].path {
				return entries[i].path < entries[j].path
			}
			return entries[i].lineNumber < entries[j].lineNumber
		})

		csvWrite(entries, ws.wsID, "translations")
	}
}
