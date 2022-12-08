package langserver

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
)

// DaedalusTranslatingListener ...
type DaedalusTranslatingListener struct {
	parser.BaseDaedalusListener
	StringLocations map[string][]SymbolPosition
	Handled         map[string]bool
	source          string
}

// DaedalusTranslatingListener ...
func NewDaedalusTranslatingListener(source string, knownSymbols SymbolProvider) *DaedalusTranslatingListener {
	return &DaedalusTranslatingListener{
		StringLocations: map[string][]SymbolPosition{},
		source:          source,
		Handled:         make(map[string]bool),
	}
}

var NewConstantNames = make(map[string]int)

//NewConstantNames = make(map[string]int)

// EnterFunctionDef ...
func (l *DaedalusTranslatingListener) EnterFunctionDef(ctx *parser.FunctionDefContext) {
	statements := ctx.StatementBlock().(*parser.StatementBlockContext)

	// Parse dialogues
	l.parseAI_OutputFuncCallStatements(statements)

}

// / Searches the statement block for AI_Output calls
func (l *DaedalusTranslatingListener) parseAI_OutputFuncCallStatements(root antlr.Tree) {
	for _, s := range root.GetChildren() {
		if funcCall, ok := s.(*parser.FuncCallContext); ok {
			if strings.EqualFold(funcCall.NameNode().GetText(), "AI_Output") {
				// We found the AI_Output call

				symbolID := strings.Trim(funcCall.AllFuncArgExpression()[2].GetText(), "\"")
				document := l.source
				line := funcCall.GetStart().GetLine()
				var start, end int
				//var content string

				parserGetter, ok := root.(interface{ GetParser() antlr.Parser })
				if !ok {
					return
				}
				tokenStream := parserGetter.GetParser().GetTokenStream()
				common, ok := tokenStream.(interface {
					GetHiddenTokensToRight(tokenIndex int, channel int) []antlr.Token
				})
				if !ok {
					return
				}
				hidden := common.GetHiddenTokensToRight(funcCall.GetStop().GetTokenIndex()+1 /* + 1 -> skipping the semicolon :) */, parser.COMMENTS)
				for _, h := range hidden {
					txt := h.GetText()
					if strings.HasPrefix(txt, "//") {
						start = h.GetColumn()
						end = start + utf8.RuneCountInString(strings.TrimSpace(txt))
						//	content = txt
						//fmt.Fprintf(os.Stderr, "\n%s\nlength: %d", txt, end-start)
						break
					}
				}

				//	fmt.Fprintf(os.Stderr, "\n[%s] %s { %d: %d-%d}\n", symbolID, content, line, start, end)

				l.StringLocations[symbolID] = append(l.StringLocations[symbolID], newSymbolPosition(document, line, start, end, false))
				//_ = newSymbolPosition(document+symbolID, line, start, end, false)
			}
		} else if s.GetChildCount() > 0 {
			l.parseAI_OutputFuncCallStatements(s)
		}
	}
}

func (l *DaedalusTranslatingListener) EnterStringLiteralValue(ctx *parser.StringLiteralValueContext) {
	//fmt.Fprintf(os.Stderr, "EnterStringLiterlValue\n")
	/*
		// Filter Ikarus, Lego, GFA and AFSP and various files TODO: Introduce config file for this
		if strings.Contains(strings.ToLower(l.source), strings.ToLower("LeGo")) ||
			strings.Contains(strings.ToLower(l.source), strings.ToLower("ikarus")) ||
			strings.Contains(strings.ToLower(l.source), strings.ToLower("AFSP")) ||
			strings.Contains(strings.ToLower(l.source), strings.ToLower("BODYSTATES")) ||
			strings.Contains(strings.ToLower(l.source), strings.ToLower("TestModelle")) ||
			strings.Contains(strings.ToLower(l.source), strings.ToLower("TestZustaende/")) ||
			strings.Contains(strings.ToLower(l.source), strings.ToLower("B_Plunder.d")) ||
			strings.Contains(strings.ToLower(l.source), strings.ToLower("B_MoveMob.d")) ||
			strings.Contains(strings.ToLower(l.source), strings.ToLower("PrintDebug.d")) ||
			strings.Contains(strings.ToLower(l.source), strings.ToLower("B_TestReaction.d")) ||
			strings.Contains(strings.ToLower(l.source), strings.ToLower("B_Functions.d")) || // asdf
			strings.Contains(strings.ToLower(l.source), strings.ToLower("/G_Constants.d")) || // asdf
			strings.Contains(strings.ToLower(l.source), strings.ToLower("/insertAnything.d")) || // asdf
			strings.Contains(strings.ToLower(l.source), strings.ToLower("/CC_Call.d")) || // asdf
			strings.Contains(strings.ToLower(l.source), strings.ToLower("GFA")) {
			return
		}

		// known instances
		if ass, ok := ctx.GetParent().GetParent().GetParent().(*parser.AssignmentContext); ok {
			if l.DidFindMemberVarStringLiteral(&l.StringLiterals, ass, "text[0]") { // C_ITEM and menus
				return
			}
			if l.DidFindMemberVarStringLiteral(&l.StringLiterals, ass, "text[1]") { // C_ITEM
				return
			}
			if l.DidFindMemberVarStringLiteral(&l.StringLiterals, ass, "text[2]") { // C_ITEM
				return
			}
			if l.DidFindMemberVarStringLiteral(&l.StringLiterals, ass, "text[3]") { // C_ITEM
				return
			}
			if l.DidFindMemberVarStringLiteral(&l.StringLiterals, ass, "text[4]") { // C_ITEM
				return
			}
			if l.DidFindMemberVarStringLiteral(&l.StringLiterals, ass, "text[5]") { // C_ITEM
				return
			}
		}
	*/

	content := ctx.ValueContext.GetText()
	if len(content) <= 2 { // skip empty strings ""
		//TODO:
		//return
	}
	// filter out string not containing letters, TODO: might be dangerous for special characters?
	if !HasLetter(content) {
		return
	}
	// Filter SVMs
	if strings.HasPrefix(content, "\"$") {
		return
		// Filter files by their extensions
	} else if strings.HasSuffix(
		strings.ToLower(content),
		strings.ToLower(".TGA\"")) {
		return
	} else if strings.HasSuffix(
		strings.ToLower(content),
		strings.ToLower(".3ds\"")) {
		return
	} else if strings.HasSuffix(
		strings.ToLower(content),
		strings.ToLower(".mms\"")) {
		return
	} else if strings.HasSuffix(
		strings.ToLower(content),
		strings.ToLower(".asc\"")) {
		return
	} else if strings.HasSuffix(
		strings.ToLower(content),
		strings.ToLower(".zen\"")) {
		return
	} else if strings.HasSuffix(
		strings.ToLower(content),
		strings.ToLower(".wav\"")) {
		return
	}

	if fncCtx, ok := ctx.GetParent().GetParent().GetParent().GetParent().(*parser.FuncCallContext); ok {
		fncName := fncCtx.NameNode().GetText()

		// Filter out TA_ and ZS_ functions, I think we can safely assume these will be used in the majority of script bases
		if strings.HasPrefix(strings.ToLower(fncName), strings.ToLower("TA")) {
			return
		} else if strings.HasPrefix(strings.ToLower(fncName), strings.ToLower("ZS")) {
			return
		}
		if strings.EqualFold(fncName, "Info_AddChoice") {
			symbolID := fncCtx.AllFuncArgExpression()[0].GetText() + "." + fncCtx.AllFuncArgExpression()[2].GetText()
			line := ctx.ValueContext.GetStart().GetLine()
			start := ctx.ValueContext.GetStart().GetColumn()
			end := start + utf8.RuneCountInString(content)
			document := l.source

			l.StringLocations[symbolID] = append(l.StringLocations[symbolID], newSymbolPosition(document, line, start, end))
			/*
				l.StringLiterals = append(l.StringLiterals,
					newStringLiteral(
						content,
						l.source,
						ctx.ValueContext.GetStart().GetLine(),
						fncCtx.AllFuncArgExpression()[0].GetText()+"."+fncCtx.AllFuncArgExpression()[2].GetText(),
						"",
						"Info_AddChoice", // translation comment
					),
				)
			*/
			return
		}
		// avoid black listed functions
		// TODO: add user defined function blacklist
		if IsBlacklisted(fncName) {
			return
		}
		/*
			else {
				const_name := ""

				if strings.EqualFold(fncName, "B_LogEntry") {
					const_name = fncCtx.AllFuncArgExpression()[0].GetText()
				} else if strings.EqualFold(fncName, "Log_AddEntry") {
					const_name = fncCtx.AllFuncArgExpression()[0].GetText()
				} else if strings.EqualFold(fncName, "EndeQuest") {
					const_name = fncCtx.AllFuncArgExpression()[0].GetText()
				} else if strings.EqualFold(fncName, "QuestTagebucheintrag") {
					const_name = fncCtx.AllFuncArgExpression()[0].GetText()
				} else if strings.EqualFold(fncName, "NeueQuest") {
					const_name = fncCtx.AllFuncArgExpression()[0].GetText()
				} else if strings.EqualFold(fncName, "Doc_PrintLine") {
					if fnc, ok := fncCtx.GetParent().GetParent().GetParent().(*parser.FunctionDefContext); ok {
						const_name = fnc.NameNode().GetText()
					}
				} else if strings.EqualFold(fncName, "Doc_PrintLines") {
					if fnc, ok := fncCtx.GetParent().GetParent().GetParent().(*parser.FunctionDefContext); ok {
						const_name = fnc.NameNode().GetText()
					}
				} else if strings.EqualFold(fncName, "PrintScreen") {
					const_name = fncName
				}

				NewConstantNames[const_name] += 1

				l.StringLiterals = append(l.StringLiterals,
					newStringLiteral(
						content,
						l.source,
						ctx.ValueContext.GetStart().GetLine(),
						"",
						fncName+" call.", // add a developer comment to tell translators/translation preparators in which function is the string used as a parameter
						const_name+"_"+fmt.Sprintf("%d", NewConstantNames[const_name]),
					),
				)
				return
			}
		*/
	}

	// known instances
	if ass, ok := ctx.GetParent().GetParent().GetParent().(*parser.AssignmentContext); ok {
		//	if ass, ok := ctx.GetParent().GetParent().GetParent().GetParent().GetParent().GetParent().(*parser.InstanceDefContext); ok {
		if l.DidFindMemberVarStringLiteral( /*&l.StringLiterals,*/ ass, "description") { // C_INFO and C_ITEM
			return
		}
		if l.DidFindMemberVarStringLiteral( /*&l.StringLiterals,*/ ass, "name") { // C_ITEM and C_NPC
			return
		}
		if l.DidFindMemberVarStringLiteral( /*&l.StringLiterals,*/ ass, "name[0]") { // name is actually a string array, some crazy people specify the first element :herosmile:
			return
		}
		if l.DidFindMemberVarStringLiteral( /*&l.StringLiterals,*/ ass, "text[0]") { // C_ITEM and menus
			return
		}
		if l.DidFindMemberVarStringLiteral( /*&l.StringLiterals,*/ ass, "text[1]") { // C_ITEM
			return
		}
		if l.DidFindMemberVarStringLiteral( /*&l.StringLiterals,*/ ass, "text[2]") { // C_ITEM
			return
		}
		if l.DidFindMemberVarStringLiteral( /*&l.StringLiterals,*/ ass, "text[3]") { // C_ITEM
			return
		}
		if l.DidFindMemberVarStringLiteral( /*&l.StringLiterals,*/ ass, "text[4]") { // C_ITEM
			return
		}
		if l.DidFindMemberVarStringLiteral( /*&l.StringLiterals,*/ ass, "text[5]") { // C_ITEM
			return
		}
	}

	// string constants
	if constCtx, ok := ctx.GetParent().GetParent().GetParent().GetParent().(*parser.ConstValueDefContext); ok {

		symbolID := constCtx.NameNode().GetText()
		line := ctx.ValueContext.GetStart().GetLine()
		start := ctx.ValueContext.GetStart().GetColumn()
		end := start + utf8.RuneCountInString(content)
		document := l.source

		l.StringLocations[symbolID] = append(l.StringLocations[symbolID], newSymbolPosition(document, line, start, end))
		/*
			l.StringLiterals = append(l.StringLiterals,
				newStringLiteral(
					content,
					document,
					line,
					symbolID,
					"",
					"",
				),
			)
		*/
		return
	}

	// SVMs
	if inst, ok := ctx.GetParent().GetParent().GetParent().GetParent().GetParent().GetParent().(*parser.InstanceDefContext); ok {
		if strings.EqualFold(inst.ParentReference().GetText(), "C_SVM") {

			symbolID := strings.Trim(content, "\"")
			document := l.source
			line := ctx.ValueContext.GetStart().GetLine()
			var start, end int

			parserGetter, ok := antlr.Tree(ctx.GetParent().GetParent().GetParent().GetParent().(*parser.StatementContext)).(interface{ GetParser() antlr.Parser })
			if !ok {
				return
			}
			tokenStream := parserGetter.GetParser().GetTokenStream()
			common, ok := tokenStream.(interface {
				GetHiddenTokensToRight(tokenIndex int, channel int) []antlr.Token
			})
			if !ok {
				return
			}
			hidden := common.GetHiddenTokensToRight(ctx.GetParent().GetParent().GetParent().GetParent().(*parser.StatementContext).GetStop().GetTokenIndex()+1 /* + 1 -> skipping the semicolon :) */, parser.COMMENTS)
			//var cntnt string
			for _, h := range hidden {
				txt := h.GetText()
				if strings.HasPrefix(txt, "//") {
					start = h.GetColumn()
					end = start + utf8.RuneCountInString(strings.TrimSpace(txt))
					//		cntnt = txt
					//fmt.Fprintf(os.Stderr, "\n%s\nlength: %d", txt, end-start)
					break
				}
			}

			//fmt.Fprintf(os.Stderr, "\n[%s] %s { %d: %d-%d}\n", symbolID, cntnt, line, start, end)

			l.StringLocations[symbolID] = append(l.StringLocations[symbolID], newSymbolPosition(document, line, start, end, false))
			/*
				l.SVMs = append(l.SVMs, SVM(newSVM(
					"",
					l.source,
					ctx.ValueContext.GetStart().GetLine(),
					content,
				)))
			*/
		}
		return
	}

	// constant arrays
	if /* constArrCtx */ c, ok := ctx.GetParent().GetParent().GetParent().(*parser.ConstArrayAssignmentContext); ok {

		for i, v := range c.GetChildren() {
			if s, ok := v.(*parser.ExpressionBlockContext); ok {
				if s == ctx.GetParent().GetParent() {

					symbolID := c.GetParent().(*parser.ConstArrayDefContext).NameNode().GetText() + "[" + fmt.Sprint(i/2-1) + "]" // TODO: Not sure, if this

					document := l.source
					line := ctx.GetStart().GetLine()
					start := ctx.GetStart().GetColumn()
					end := start + utf8.RuneCountInString(content)

					//fmt.Fprintf(os.Stderr, "\n[%s] %s { %d: %d-%d}\n", symbolID, content, line, start, end)

					l.StringLocations[symbolID] = append(l.StringLocations[symbolID], newSymbolPosition(document, line, start, end))
				}
			}
		}
		return
	}

	// if anything slips by, label it NOT HANDLED, useful for catching unusual cases
	/*
		l.StringLiterals = append(l.StringLiterals,
			newStringLiteral(
				content,
				l.source,
				ctx.ValueContext.GetStart().GetLine(),
				"",
				"NOT HANDLED",
				"",
			),
		)
	*/
}

// Checks, wheter thi string literal is assigned to the `field` specified
func (l *DaedalusTranslatingListener) DidFindMemberVarStringLiteral( /*list *[]StringLiteral,*/ ass *parser.AssignmentContext, field string) bool {
	if strings.EqualFold(ass.Reference().GetText(), field) {
		if strings.Contains(ass.ExpressionBlock().GetText(), "\"") {
			if efs, ok := ass.GetParent().GetParent().GetParent().(*parser.InstanceDefContext); ok {

				content := ass.ExpressionBlock().GetText()
				symbolID := efs.NameNode().GetText() + "." + field
				line := ass.ExpressionBlock().GetStart().GetLine()
				start := ass.ExpressionBlock().GetStart().GetColumn()
				end := start + utf8.RuneCountInString(content)
				document := l.source

				l.StringLocations[symbolID] = append(l.StringLocations[symbolID], newSymbolPosition(document, line, start, end))
				/*
					*list = append(*list,
						newStringLiteral(
							content,
							l.source,
							ass.ExpressionBlock().GetStart().GetLine(),
							efs.NameNode().GetText()+"."+field,
							"",
							"",
						),
					)
				*/
				return true
				// also check in prototypes
			} else if efs, ok := ass.GetParent().GetParent().GetParent().(*parser.PrototypeDefContext); ok {

				//content := ass.ExpressionBlock().GetText()
				symbolID := efs.NameNode().GetText() + "." + field
				line := ass.ExpressionBlock().GetStart().GetLine()
				start := ass.ExpressionBlock().GetStart().GetColumn()
				end := start + utf8.RuneCountInString(ass.ExpressionBlock().GetText())
				document := l.source

				l.StringLocations[symbolID] = append(l.StringLocations[symbolID], newSymbolPosition(document, line, start, end))
				/*
					*list = append(*list,
						newStringLiteral(
							content,
							l.source,
							ass.ExpressionBlock().GetStart().GetLine(),
							efs.NameNode().GetText()+"."+field,
							"",
							"",
						),
					)
				*/
				return true
			}

		}
	}
	return false
}

func HasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// TODO: move this to json config
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
