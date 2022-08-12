/*
	Functionality related to the translation feature
*/
package langserver

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/integralist/go-findroot/find"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
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
		translator_comments: "",
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

// Generates one giant escaped csv file with all translatable strings
func (h *LspHandler) generateAllCSV() {
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
	csvWrite(entries, "All", "cs")
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
}

func newStringLiteral(txt, source string, ln int, ctx, devCmnt string) StringLiteral {
	return StringLiteral{
		text:       txt,
		sourceFile: source,
		line:       ln,
		context:    ctx,
		devComment: devCmnt,
	}
}

type SVM struct {
	text       string
	sourceFile string
	line       int
	context    string
	devComment string
}

func newSVM(txt, source string, ln int, ctx string) StringLiteral {
	return StringLiteral{
		text:       txt,
		sourceFile: source,
		line:       ln,
		context:    ctx,
		devComment: "",
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
