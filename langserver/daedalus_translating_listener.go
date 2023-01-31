package langserver

import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	// lsp "github.com/kirides/DaedalusLanguageServer/protocol"
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
	// INFO: I check these before the empty string check because of how "dynamically" these fields are handled in translations
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
		return
	}
	// filter out string not containing letters, FIXME: might be dangerous for special characters?
	if !HasLetter(content) {
		return
	}

	if isFileterString(content) {
		return
	}

	// string literals as a function call parameter
	if fncCtx, ok := ctx.GetParent().GetParent().GetParent().GetParent().(*parser.FuncCallContext); ok {
		fncName := fncCtx.NameNode().GetText()

		// // Filter out TA_ and ZS_ functions, I think we can safely assume these will be used in the majority of script bases
		// // TODO: Add this to a filter also
		// if strings.HasPrefix(strings.ToLower(fncName), strings.ToLower("TA")) {
		// 	return
		// } else if strings.HasPrefix(strings.ToLower(fncName), strings.ToLower("ZS")) {
		// 	return
		// }
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

		// TODO: Somehow build a call stack for all cunftion calls and check, if the string literal ends up "terminating" in an non blacklisted external
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

		l.StringLocations["NOT HANDLED"] = append(
			l.StringLocations["NOT HANDLED"],
			newSymbolPosition(
				l.source,
				content,
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				ctx.GetStart().GetLine() + utf8.RuneCountInString(content),
				"NOT HADNDLED"))
}

// Checks, wheter the string literal is assigned to the `field` specified
func (l *DaedalusTranslatingListener) DidFindMemberVarStringLiteral(ass *parser.AssignmentContext, field string) bool {
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

// Is the function blacklisted?
func (l *DaedalusTranslatingListener) isBlacklistedFunc(functionName string) bool {
	for _, a := range l.config.FunctionBlackList {
		if strings.EqualFold(a, functionName) {
			return true
		}
	}
	return false
}

// Is the path blacklisted?
func (l *DaedalusTranslatingListener) isBlacklistedPath(pathMask string) bool {
	for _, mask := range l.config.FileMasks {
		if strings.Contains(strings.ToLower(mask), strings.ToLower(pathMask)) {
			return true
		}
	}
	return false
}


// Return true if the string contains $ in the prefix (SVMs) or
// if it cointans any of the file extensions
func isFileterString(s string) bool {
	if strings.HasPrefix(s, "\"$") {
		return true
		// Filter files by their extensions
	} else if strings.HasSuffix(
		strings.ToLower(s),
		strings.ToLower(".Tga\"")) {
		return true
	} else if strings.HasSuffix(
		strings.ToLower(s),
		strings.ToLower(".3ds\"")) {
		return true
	} else if strings.HasSuffix(
		strings.ToLower(s),
		strings.ToLower(".mms\"")) {
		return true
	} else if strings.HasSuffix(
		strings.ToLower(s),
		strings.ToLower(".asc\"")) {
		return true
	} else if strings.HasSuffix(
		strings.ToLower(s),
		strings.ToLower(".zen\"")) {
		return true
	} else if strings.HasSuffix(
		strings.ToLower(s),
		strings.ToLower(".wav\"")) {
		return true
	}
	return false
}

func (l *DaedalusTranslatingListener) EnterMacroDef(ctx *parser.MacroDefContext) {
	fmt.Fprintf(os.Stderr, "Enter Macro Definition: %v\n", ctx.GetChildren())
	start_ln := ctx.GetStart().GetLine()
	start_col := ctx.GetStart().GetColumn()
	end_ln := ctx. GetStop().GetLine()
	end_col := ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText())+1

	fmt.Fprintf(os.Stderr, "(%d, %d) - (%d, %d)\n", start_ln, start_col, end_ln, end_col)
	for i, ch := range ctx.GetChildren() {
		numOfBlocks := len(ctx.GetChildren())
		if i == numOfBlocks - 1 { // last one is #endif
			break
		}
		if i == 0 { // first element is a if block
			// fmt.Fprintf(os.Stderr, "%v\n", ch.(*parser.MacroIfBlockContext).GetChild(2).(*parser.MacroBlockContext).GetText())
			l.ParseMacroIfBlock(ch.(*parser.MacroIfBlockContext))
		} else if i == numOfBlocks - 2 { // last element is a else block
			l.ParseMacroElseBlock(ch.(*parser.MacroElseBlockContext))
			// fmt.Fprintf(os.Stderr, "%v\n", ch.(*parser.MacroElseBlockContext).GetChild(1).(*parser.MacroBlockContext).GetText())
		} else { // middle elements are elif blocks
			// fmt.Fprintf(os.Stderr, "%v\n", ch.(*parser.MacroElseIfBlockContext).GetChild(2).(*parser.MacroBlockContext).GetText())
			l.ParseMacroElseIfBlock(ch.(*parser.MacroElseIfBlockContext))
		}
	}
}

type MacroBlock struct {
    condition string
	content   string
//	position  lsp.Position
}


func (l *DaedalusTranslatingListener) ParseMacroIfBlock(ctx *parser.MacroIfBlockContext) {
	condition := ctx.GetChild(1).(*parser.MacroConditionContext).GetText()
	fmt.Fprintf(os.Stderr, "condition: %s\n", condition)
	content := ctx.GetChild(2).(*parser.MacroBlockContext).GetText()
	fmt.Fprintf(os.Stderr, "macro block content: %s\n", content)
}

func (l *DaedalusTranslatingListener) ParseMacroElseIfBlock(ctx *parser.MacroElseIfBlockContext) {
	condition := ctx.GetChild(1).(*parser.MacroConditionContext).GetText()
	fmt.Fprintf(os.Stderr, "condition: %s\n", condition)
	content := ctx.GetChild(2).(*parser.MacroBlockContext).GetText()
	fmt.Fprintf(os.Stderr, "macro block content: %s\n", content)
}

func (l *DaedalusTranslatingListener) ParseMacroElseBlock(ctx *parser.MacroElseBlockContext) {
	content := ctx.GetChild(1).(*parser.MacroBlockContext).GetText()
	fmt.Fprintf(os.Stderr, "macro block content: %s\n", content)
}


/*
func (l *DaedalusTranslatingListener) EnterMacroIfBlock(ctx *parser.MacroIfBlockContext) {
	fmt.Fprintf(os.Stderr, "\n\nEnterMacroIfBlock\n\n")
}

func (l *DaedalusTranslatingListener) EnterMacroElseBlock (ctx *parser.MacroElseBlockContext) {
	fmt.Fprintf(os.Stderr, "\n\nMacro ElseBlock\n\n")
}
*/
