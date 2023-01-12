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
	source          string

	config translationConfiguration
}

// DaedalusTranslatingListener ...
func NewDaedalusTranslatingListener(source string, conf translationConfiguration) *DaedalusTranslatingListener {
	return &DaedalusTranslatingListener{
		StringLocations: map[string][]SymbolPosition{},
		source:          source,
		config:          conf,
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
				var content string

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
						end = start + utf8.RuneCountInString(txt)
						content = txt
						break
					}
				}

				l.StringLocations[symbolID] = append(l.StringLocations[symbolID], newSymbolPosition(document, content, line, start, end, false))
				//_ = newSymbolPosition(document+symbolID, line, start, end, false)
			}
		} else if s.GetChildCount() > 0 {
			l.parseAI_OutputFuncCallStatements(s)
		}
	}
}

func (l *DaedalusTranslatingListener) EnterStringLiteralValue(ctx *parser.StringLiteralValueContext) {
	// skip blacklisted paths
	if l.isBlacklistedPath(l.source) {
		return
	}

	// known instance members
	// INFO: I check these before the empty string check because of how "dynamic" these fields are handled in translations
	if ass, ok := ctx.GetParent().GetParent().GetParent().(*parser.AssignmentContext); ok {
		for _, field := range l.config.MemberVariables {
			if l.DidFindMemberVarStringLiteral(ass, field) {
				return
			}
		}
	}

	// get the string content
	content := ctx.ValueContext.GetText()

	// skip empty strings ""
	if len(content) <= 2 {
		//TODO:
		return
	}
	// filter out string not containing letters, TODO: might be dangerous for special characters?
	if !HasLetter(content) {
		return
	}
	// Filter SVM calls
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

	// string literals as a function call parameter
	if fncCtx, ok := ctx.GetParent().GetParent().GetParent().GetParent().(*parser.FuncCallContext); ok {
		fncName := fncCtx.NameNode().GetText()

		// Filter out TA_ and ZS_ functions, I think we can safely assume these will be used in the majority of script bases
		// TODO: Add this to a filter also
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

			l.StringLocations[symbolID] = append(l.StringLocations[symbolID], newSymbolPosition(document, trimQuotes(content), line, start, end, "Info_AddChoice"))

			return
		}
		// avoid black listed functions
		if l.isBlacklistedFunc(fncName) {
			return
		}

		// Somehow build a call stack and check, if the string literal ends up "terminating" in an non blacklisted external

		/*
			// These are examples of functions I had to extract const strings from. TODO: Make it configurable from a .json file
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

	// string constants
	if constCtx, ok := ctx.GetParent().GetParent().GetParent().GetParent().(*parser.ConstValueDefContext); ok {

		symbolID := constCtx.NameNode().GetText()
		line := ctx.ValueContext.GetStart().GetLine()
		start := ctx.ValueContext.GetStart().GetColumn()
		end := start + utf8.RuneCountInString(content)
		document := l.source

		l.StringLocations[symbolID] = append(l.StringLocations[symbolID], newSymbolPosition(document, trimQuotes(content), line, start, end))

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
			var cntnt string
			for _, h := range hidden {
				txt := h.GetText()
				if strings.HasPrefix(txt, "//") {
					start = h.GetColumn()
					end = start + utf8.RuneCountInString(strings.TrimSpace(txt))
					cntnt = txt
					break
				}
			}

			l.StringLocations[symbolID] = append(l.StringLocations[symbolID], newSymbolPosition(document, cntnt, line, start, end, false))
		}
		return
	}

	// constant arrays
	if c, ok := ctx.GetParent().GetParent().GetParent().(*parser.ConstArrayAssignmentContext); ok {

		for i, v := range c.GetChildren() {
			if s, ok := v.(*parser.ExpressionBlockContext); ok {
				if s == ctx.GetParent().GetParent() {

					symbolID := c.GetParent().(*parser.ConstArrayDefContext).NameNode().GetText() + "[" + fmt.Sprint(i/2-1) + "]"

					document := l.source
					line := ctx.GetStart().GetLine()
					start := ctx.GetStart().GetColumn()
					end := start + utf8.RuneCountInString(content)

					l.StringLocations[symbolID] = append(l.StringLocations[symbolID], newSymbolPosition(document, trimQuotes(content), line, start, end))
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

				l.StringLocations[symbolID] = append(l.StringLocations[symbolID], newSymbolPosition(document, trimQuotes(content), line, start, end))

				return true
				// also check in prototypes
			} else if efs, ok := ass.GetParent().GetParent().GetParent().(*parser.PrototypeDefContext); ok {

				content := ass.ExpressionBlock().GetText()
				symbolID := efs.NameNode().GetText() + "." + field
				line := ass.ExpressionBlock().GetStart().GetLine()
				start := ass.ExpressionBlock().GetStart().GetColumn()
				end := start + utf8.RuneCountInString(content)
				document := l.source

				l.StringLocations[symbolID] = append(l.StringLocations[symbolID], newSymbolPosition(document, trimQuotes(content), line, start, end))

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

func (l *DaedalusTranslatingListener) isBlacklistedFunc(functionName string) bool {
	for _, a := range l.config.FunctionBlackList {
		if strings.EqualFold(a, functionName) {
			return true
		}
	}
	return false
}

func (l *DaedalusTranslatingListener) isBlacklistedPath(pathMask string) bool {
	for _, mask := range l.config.FileMasks {
		if strings.Contains(strings.ToLower(mask), strings.ToLower(pathMask)) {
			return true
		}
	}
	return false
}
