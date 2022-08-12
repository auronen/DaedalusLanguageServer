package langserver

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
)

// DaedalusTranslationStatefulListener ...
type DaedalusTranslationStatefulListener struct {
	parser.BaseDaedalusListener
	source string

	// Dialogues
	GlobalDialogues []Dialogue
	StringLiterals  []StringLiteral
	SVMs            []SVM
}

// NewDaedalusTranslationStatefulListener ...
func NewDaedalusTranslationStatefulListener(source string) *DaedalusTranslationStatefulListener {
	return &DaedalusTranslationStatefulListener{
		source:          source,
		GlobalDialogues: []Dialogue{},
		StringLiterals:  []StringLiteral{},
		SVMs:            []SVM{},
	}
}

// EnterFunctionDef ...
func (l *DaedalusTranslationStatefulListener) EnterFunctionDef(ctx *parser.FunctionDefContext) {
	statements := ctx.StatementBlock().(*parser.StatementBlockContext)
	// Parse dialogues
	l.GlobalDialogues = append(l.GlobalDialogues, l.findAI_OutputFuncCallsStatements(statements)...)
}

// Searches the statement block for AI_Output calls
func (l *DaedalusTranslationStatefulListener) findAI_OutputFuncCallsStatements(root antlr.Tree) []Dialogue {
	var foundCalls []Dialogue
	for _, s := range root.GetChildren() {
		if funcCall, ok := s.(*parser.FuncCallContext); ok {
			if strings.EqualFold(funcCall.NameNode().GetText(), "AI_Output") {
				foundCalls = append(foundCalls, newDialogue(funcCall.AllFuncArgExpression()[2].GetText(),
					"", // filled in in the csv generation phase
					l.source,
					funcCall.GetStart().GetLine()))
			}
		} else if s.GetChildCount() > 0 {
			foundCalls = append(foundCalls, l.findAI_OutputFuncCallsStatements(s)...)
		}
	}
	return foundCalls
}

func (l *DaedalusTranslationStatefulListener) EnterStringLiteralValue(ctx *parser.StringLiteralValueContext) {

	// Filter Ikarus, Lego, GFA and AFSP
	if strings.Contains(strings.ToLower(l.source), strings.ToLower("LeGo")) ||
		strings.Contains(strings.ToLower(l.source), strings.ToLower("ikarus")) ||
		strings.Contains(strings.ToLower(l.source), strings.ToLower("AFSP")) ||
		strings.Contains(strings.ToLower(l.source), strings.ToLower("BODYSTATES")) ||
		strings.Contains(strings.ToLower(l.source), strings.ToLower("GFA")) {
		return
	}

	content := ctx.ValueContext.GetText()
	if len(content) <= 2 { // skip empty strings ""
		return
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
			l.StringLiterals = append(l.StringLiterals,
				newStringLiteral(
					content,
					l.source,
					ctx.ValueContext.GetStart().GetLine(),
					fncCtx.AllFuncArgExpression()[2].GetText(),
					"",
				),
			)
			return
		}
		// avoid black listed functions
		// TODO: add user defined function blacklist
		if IsBlacklisted(fncName) {
			return
		} else {
			l.StringLiterals = append(l.StringLiterals,
				newStringLiteral(
					content,
					l.source,
					ctx.ValueContext.GetStart().GetLine(),
					"",
					fncName+" call.", // add a developer comment to tell translators/translation preparators in which function is the string used as a parameter
				),
			)
			return
		}
	}

	// known instances
	if ass, ok := ctx.GetParent().GetParent().GetParent().(*parser.AssignmentContext); ok {
		//	if ass, ok := ctx.GetParent().GetParent().GetParent().GetParent().GetParent().GetParent().(*parser.InstanceDefContext); ok {
		if l.DidFindMemberVarStringLiteral(&l.StringLiterals, ass, "description") { // C_INFO and C_ITEM
			return
		}
		if l.DidFindMemberVarStringLiteral(&l.StringLiterals, ass, "name") { // C_ITEM and C_NPC
			return
		}
		if l.DidFindMemberVarStringLiteral(&l.StringLiterals, ass, "name[0]") { // name is actually a string array, some crazy people specify the first element :herosmile:
			return
		}
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

	// string constants
	if constCtx, ok := ctx.GetParent().GetParent().GetParent().GetParent().(*parser.ConstValueDefContext); ok {
		l.StringLiterals = append(l.StringLiterals,
			newStringLiteral(
				content,
				l.source,
				ctx.ValueContext.GetStart().GetLine(),
				constCtx.NameNode().GetText(),
				"",
			),
		)
		return
	}

	// SVMs
	if inst, ok := ctx.GetParent().GetParent().GetParent().GetParent().GetParent().GetParent().(*parser.InstanceDefContext); ok {
		if strings.EqualFold(inst.ParentReference().GetText(), "C_SVM") {
			l.SVMs = append(l.SVMs, SVM(newSVM(
				"",
				l.source,
				ctx.ValueContext.GetStart().GetLine(),
				content,
			)))
		}
		return
	}

	// constat arrays
	if /* constArrCtx */ _, ok := ctx.GetParent().GetParent().GetParent().(*parser.ConstArrayAssignmentContext); ok {
		// skip them, the are already parsed
		return
	}

	// if anything slips by, label it NOT HANDLED, useful for catching unusual cases
	l.StringLiterals = append(l.StringLiterals,
		newStringLiteral(
			content,
			l.source,
			ctx.ValueContext.GetStart().GetLine(),
			"",
			"NOT HANDLED",
		),
	)

}

// Checks, wheter thi string literal is assigned to the `field` specified
func (l *DaedalusTranslationStatefulListener) DidFindMemberVarStringLiteral(list *[]StringLiteral, ass *parser.AssignmentContext, field string) bool {
	if strings.EqualFold(ass.Reference().GetText(), field) {
		if strings.Contains(ass.ExpressionBlock().GetText(), "\"") {
			if efs, ok := ass.GetParent().GetParent().GetParent().(*parser.InstanceDefContext); ok {
				*list = append(*list,
					newStringLiteral(
						ass.ExpressionBlock().GetText(),
						l.source,
						ass.ExpressionBlock().GetStart().GetLine(),
						efs.NameNode().GetText()+"."+field,
						"",
					),
				)
				return true
				// also check in prototypes
			} else if efs, ok := ass.GetParent().GetParent().GetParent().(*parser.PrototypeDefContext); ok {
				*list = append(*list,
					newStringLiteral(
						ass.ExpressionBlock().GetText(),
						l.source,
						ass.ExpressionBlock().GetStart().GetLine(),
						efs.NameNode().GetText()+"."+field,
						"",
					),
				)
				return true
			}

		}
	}
	return false
}
