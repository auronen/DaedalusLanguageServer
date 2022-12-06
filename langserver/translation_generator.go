/*
Functionality related to the translation feature
*/
package langserver

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/integralist/go-findroot/find"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "go.lsp.dev/protocol"
	"go.lsp.dev/uri"
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

type simpleCSVentry struct {
	context string
	text    string
}

func (c *simpleCSVentry) GetEscapedSimpleCSV() string {
	return "\"" + c.context + "\",\"" + c.text + "\""
}

// "location","source","target","id","fuzzy","context","translator_comments","developer_comments"
type CSVentry struct {
	location            string // {git repo relative path}:{line number}
	source              string // Source string
	target              string // Translated string - empty
	id                  string // ? internal
	fuzzy               string // ? internal
	context             string // when const - symbol name, when SVM - name, when dia - sound file and so on
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
	func newCSVentryFromConstSymbol(symbol symbol.Constant) CSVentry {
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
*/
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
		entries = append(entries, line /*.getValue()*/)
	}
	return entries
}

// TODO
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

/*
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
*/

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
	if err == nil { // ..............................remove leading "/"
		return strings.ReplaceAll(path, repoRoot, "")[1:], err
	}
	return "", err
}

// old csvWirte func
// it is not possible to force all entries to be escaped
/*
func csvWrite(data [][]string, filename, lang string) error {
	// create path
	repoRoot, err := findRepoRoot()
	if err != nil {
		return err
	}
	_ = os.MkdirAll(filepath.Join(repoRoot, ".translations", lang), os.ModePerm)
	file, err := os.Create(filepath.Join(repoRoot, ".translations", lang, filename+"_"+lang+".csv"))
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
*/

// write the CSVentry slice as an escaped CSV file with the filename: `filename`_`lang`.csv
func csvWrite(entries []CSVentry, filename, lang string) error {
	// create path
	repoRoot, err := findRepoRoot()
	if err != nil {
		return err
	}
	_ = os.MkdirAll(filepath.Join(repoRoot, ".translations", lang), os.ModePerm)
	file, err := os.Create(filepath.Join(repoRoot, ".translations", lang, filename+"_"+lang+".csv"))
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

		for _, r := range records {
			// "location","source","target","id","fuzzy","context","translator_comments","developer_comments"
			translatedStrings = append(translatedStrings, newTranslatedString(r[5], r[2]))
		}
	}
	return
}

// write the CSVentry slice as an escaped CSV file with the filename: `filename`_`lang`.csv
func simplecsvWrite(entries []simpleCSVentry, filename, lang string) error {

	path, _ := os.Getwd()

	_ = os.MkdirAll(filepath.Join(path, ".translations"), os.ModePerm)
	pth := filepath.Join(path, ".translations", filename+"_"+lang+".csv")
	file, err := os.Create(pth)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	for _, entry := range entries {
		writer.WriteString(entry.GetEscapedSimpleCSV() + "\n")
	}
	return nil
}

// Generates one giant escaped csv file with all translatable strings
func (h *LspHandler) generateAllCSV() {
	//h.logger.Infof("src files %v", h.parsedKnownSrcFiles)
	var what string
	h.parsedKnownSrcFiles.m.Range(func(mapKey, any interface{}) bool {
		key := mapKey.(string)
		if strings.Contains(strings.ToLower(key), strings.ToLower("Gothic.src")) {
			what = "Gothic"
		} else if strings.Contains(strings.ToLower(key), strings.ToLower("Menu.src")) {
			what = "Menu"
		}
		return true
	})
	h.logger.Infof("Generating %s", what)

	entries := make([]CSVentry, 0, 200)

	// everything
	for _, res := range h.parsedDocuments.parseResults {
		for _, strLit := range res.StringLiterals {
			entries = append(entries, newCSVentryFromStringLiteral(strLit))
		}
	}
	// const arrays
	consts, err := h.parsedDocuments.GetGlobalSymbols(SymbolConstant)
	if err != nil {
		return
	}
	for _, c := range consts {
		if sym, ok := c.(symbol.ConstantArray); ok {
			if strings.EqualFold(sym.Type, "string") {
				if !isBlacklisted_constArray(sym.NameValue) {
					entries = append(entries, newCSVentryFromConstArraySymbol(sym)...)
				}
			}
		}
	}

	// dialogues
	for _, res := range h.parsedDocuments.parseResults {
		for _, dia := range res.GlobalDialogues {
			dia.text = strings.TrimLeft(res.GlobalComments[dia.line].text, "/")
			entries = append(entries, newCSVentryFromDialogue(dia))
		}
	}

	// SVMs
	for _, res := range h.parsedDocuments.parseResults {
		for _, svm := range res.SVMs {
			svm.text = strings.TrimLeft(res.GlobalComments[svm.line].text, "/")
			entries = append(entries, newCSVentryFromSVM(svm))
		}
	}

	// Sort entries by file and then by line (ensures comfortable translations)
	// and enables "easy" file splitting
	sort.Slice(entries, func(i, j int) bool {
		if !strings.EqualFold(entries[i].path, entries[j].path) {
			return entries[i].path < entries[j].path
		}
		return entries[i].lineNumber < entries[j].lineNumber
	})

	h.logger.Infof("Got %d lines", len(entries))

	//csvWrite(entries, "All", "cs")
	csvWrite(entries, what, "translations")
}

// function black list
var FuncBlackList = []string{
	"PrintDebug",
	"PrintDialog",
	"PrintDebugInst",
	"PrintDebugInstCh",
	"PrintDebugCh",
	"PrintMulti",
	"PlayVideo",
	"PlayVideoEx",
	"Wld_IsMobAvailable",
	"Wld_IsFPAvailable",
	"Wld_IsNextFPAvailable",
	"Npc_GetDisttowp",
	"AI_StartState",
	"AI_OutputSvm",
	"AI_OutputSvm_Overlay",
	"AI_PlayCutscene",
	"AI_PlayAni",
	"AI_PlayAniBS",
	"AI_GoToWP",
	"AI_Teleport",
	"AI_GoToFP",
	"Npc_IsOnFP",
	"AI_GoToNextFP",
	"Npc_StopAni",
	"Npc_PlayAni",
	"Wld_GetMobState",
	"AI_LookAt",
	"AI_PointAt",
	"AI_AskText",
	"AI_Snd_Play",
	"AI_Snd_Play3d",
	"Snd_Play",
	"Snd_Play3d",
	"Mis_AddMissionEntry",
	"Mdl_SetVisual",
	"Mdl_SetVisualBody",
	"Mdl_ApplyOverlayMds",
	"Mdl_ApplyOverlayMdsTimed",
	"Mdl_RemoveOverlayMds",
	"Mdl_ApplyRandomAni",
	"Mdl_ApplyRandomAniFreq",
	"Mdl_StartFaceAni",
	"Mdl_ApplyRandomFaceAni",
	"Wld_InsertNpc",
	"Wld_PlayEffect",
	"Wld_StopEffect",
	"AI_PlayFx",
	"AI_StopFx",
	"Wld_InsertNpcAndRespawn",
	"Wld_InsertItem",
	"Wld_InsertObject",
	"Wld_ExchangeGuildAttitudes",
	"Wld_SetObjectRoutine",
	"Wld_SetMobRoutine",
	"Wld_SendTrigger",
	"Wld_SendUntrigger",
	"AI_TakeMob",
	"AI_UseMob",
	"Mob_CreateItems",
	"Mob_HasItems",
	"Doc_SetPage",
	"Doc_SetFont",
	"Doc_SetLevel",
	"Doc_Open",
	"Doc_Font",
	"Doc_MapCoordinates",
	"TA",
	"TA_Min",
	"TA_Cs",
	"Npc_ExchangeRoutine",
	"RTN_Exchange",
	"Wld_AssignRoomToGuild",
	"Wld_AssignRoomToNpc",
	"Hlp_CutscenePlayed",
	"AI_Output",
	"PrintDebugNpc",
	"B_ExchangeRoutineRun",
	"B_ExchangeRoutine",
	"B_SayOverlay",
	"PrintDebugString",
	"PrintDebugInt",
	"B_GotoFP",
	"B_InterruptMob",
	"ZS_Meditate_Om",
	"B_StartUseMob",
	"B_PracticeCombat",
	"B_StopUseMob",
	"Hlp_StrCmp",
	"VobShowVisual",
	"SetVobVisual",
	"SearchVobsByClass",
	"MEM_SearchVobByName",
}

func IsBlacklisted(e string) bool {
	for _, a := range FuncBlackList {
		if strings.EqualFold(a, e) {
			return true
		}
	}
	return false
}

// const array black list
/*
	This is not very nice, it would probably be better to redo this based on the
	path filters and apply them before appending them when generating the CSV
*/
var constArrBlackList = []string{
	"spellFxAniLetters",    // base game
	"spellFxInstanceNames", // base game
	"GFA_AIM_ANIS",         // Gothic Free Aim
	"PARSER_TOKEN_NAMES",   // Ikarus
}

func isBlacklisted_constArray(e string) bool {
	for _, a := range constArrBlackList {
		if strings.EqualFold(a, e) {
			return true
		}
	}
	return false
}

func (h *LspHandler) generatezPECSV() {
	//TODO
	h.logger.Infof("Generating Autorun\n")

	var entries [8][]simpleCSVentry
	for i := 0; i < 8; i++ {
		entries[i] = make([]simpleCSVentry, 0, 20)
	}

	for _, res := range h.parsedDocuments.parseResults {
		for _, strLit := range res.UnionLocalized {
			for i := 0; i < 8; i++ {
				entries[i] = append(entries[i], newsimpleCSVentryTranslationStringEntry(strLit, i))
			}
		}
	}
	langs := []string{"ru", "en", "de", "pl", "ro", "it", "cs", "es"}

	for i := 0; i < 8; i++ {
		simplecsvWrite(entries[i], "autorun", langs[i])
	}
}

type Comment struct {
	text string
	line int
}

func newCommentFromToken(tok antlr.Token) Comment {
	return Comment{
		text: tok.GetText(),
		line: tok.GetLine(),
	}
}

type Dialogue struct {
	soundName  string
	text       string
	sourceFile string
	line       int
}

func newDialogue(sndName, txt, srcFile string, ln int) Dialogue {
	return Dialogue{
		soundName:  sndName,
		text:       txt,
		sourceFile: srcFile,
		line:       ln,
	}
}

type StringLiteral struct {
	text       string
	sourceFile string
	line       int
	context    string
	devComment string
	trComment  string
}

func newStringLiteral(txt, source string, ln int, ctx, devCmnt, trCmnt string) StringLiteral {
	return StringLiteral{
		text:       txt,
		sourceFile: source,
		line:       ln,
		context:    ctx,
		devComment: devCmnt,
		trComment:  trCmnt,
	}
}

type SVM struct {
	text       string
	sourceFile string
	line       int
	context    string
	devComment string
	trComment  string
}

func newSVM(txt, source string, ln int, ctx string) StringLiteral {
	return StringLiteral{
		text:       txt,
		sourceFile: source,
		line:       ln,
		context:    ctx,
		devComment: "",
		trComment:  "",
	}
}

/*
func (h *LspHandler) parseScripts() {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	GothicSrcPath, err := findPathAnywhereUpToRoot(wd, "Gothic.src")

}
*/

type TranslationStringEntry struct {
	name string
	RUS  string
	ENG  string
	GER  string
	POL  string
	ROU  string
	ITA  string
	CZE  string
	ESP  string
}

func NewTranslationStringEntry(ctx, ru, en, de, pl, ro, it, cs, es string) TranslationStringEntry {
	return TranslationStringEntry{
		name: ctx,
		RUS:  ru,
		ENG:  en,
		GER:  de,
		POL:  pl,
		ROU:  ro,
		ITA:  it,
		CZE:  cs,
		ESP:  es,
	}
}

// test lsp server's ability to edit code
func (h *LspHandler) editCode(ctx context.Context) {

	// create a map of text edits
	edits := make(map[uri.URI][]lsp.TextEdit)

	h.logger.Infof("Creating edits")
	ts, err := csvReadLanguage("cs")
	if err != nil {
		h.logger.Errorf("Error: %s", err)
	}

	for i, entry := range ts {
		edits2 := make([]lsp.TextDocumentEdit, 0, 50)
		//h.logger.Infof("%s - %s", entry.id, entry.content)
		for _, res := range h.parsedDocuments.parseResults {
			if positions, ok := res.StringLocations[entry.id]; ok {
				for _, pos := range positions {
					edits[uri.File(pos.document)] = append(edits[uri.File(pos.document)], lsp.TextEdit{
						Range: lsp.Range{
							Start: lsp.Position{
								Line:      uint32(pos.line - 1),
								Character: uint32(pos.start),
							},
							End: lsp.Position{
								Line:      uint32(pos.line - 1),
								Character: uint32(pos.end),
							},
						},
						NewText: "\"" + entry.content + "\"",
					})

					// second way of doing things
					edits2 = append(edits2, lsp.TextDocumentEdit{
						TextDocument: lsp.OptionalVersionedTextDocumentIdentifier{
							TextDocumentIdentifier: lsp.TextDocumentIdentifier{
								URI: uri.File(pos.document),
							},
							Version: lsp.NewVersion(10),
						},
						Edits: []lsp.TextEdit{
							lsp.TextEdit{
								Range: lsp.Range{
									Start: lsp.Position{
										Line:      uint32(pos.line - 1),
										Character: uint32(pos.start),
									},
									End: lsp.Position{
										Line:      uint32(pos.line - 1),
										Character: uint32(pos.end),
									},
								},
								NewText: "\"" + entry.content + "\"",
							},
						},
					})
				}

			}
		}

		h.logger.Infof("Done with edits %d", i)

		var response lsp.ApplyWorkspaceEditResponse

		// send the request to the editor
		/*id, _ := */
		h.conn.Call(context.Background(), lsp.MethodWorkspaceApplyEdit, lsp.ApplyWorkspaceEditParams{
			Edit: lsp.WorkspaceEdit{
				//Changes: edits,
				DocumentChanges: edits2,
			}}, &response)

	}
	/*
		h.logger.Infof("Done with edits")

		h.logger.Infof("Number of edits: %d", len(edits))
		for key, _ := range edits {
			h.logger.Infof("%v - %d", key, len(edits[key]))
		}
		//h.logger.Infof("%v", edits)

		var response lsp.ApplyWorkspaceEditResponse

		// send the request to the editor
		h.conn.Call(context.Background(), lsp.MethodWorkspaceApplyEdit, lsp.ApplyWorkspaceEditParams{
			Edit: lsp.WorkspaceEdit{
				//Changes: edits,
				DocumentChanges: edits2,
			}}, &response)

		h.logger.Infof("Done with function")
	*/
}

/*
func (h *LspHandler) createEdits(edits map[uri.URI][]lsp.TextEdit) {

}
*/

// TODO: rename this to something reasonable and maybe try to merge it with StringLiterals
type SymbolPosition struct {
	document string
	line     int
	start    int
	end      int
}

func newSymbolPosition(doc string, ln, st, en int) SymbolPosition {
	return SymbolPosition{
		document: doc,
		line:     ln,
		start:    st,
		end:      en,
	}
}

type TranslatedString struct {
	id      string
	content string
}

func newTranslatedString(id, content string) TranslatedString {
	return TranslatedString{
		id:      id,
		content: content,
	}
}
