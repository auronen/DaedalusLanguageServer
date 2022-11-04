package langserver

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"

	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// DaedalusStatefulListener ...
type DaedalusStatefulListener struct {
	parser.BaseDaedalusListener
	knownSymbols    SymbolProvider
	Instances       map[string]symbol.ProtoTypeOrInstance
	GlobalVariables map[string]symbol.Symbol
	GlobalConstants map[string]symbol.Symbol
	Functions       map[string]symbol.Function
	Classes         map[string]symbol.Class
	Prototypes      map[string]symbol.ProtoTypeOrInstance
	summaryBuilder  *bytes.Buffer
	source          string

	Tokens antlr.CommonTokenStream
	// Translations
	GlobalDialogues []Dialogue
	StringLiterals  []StringLiteral
	SVMs            []SVM
	UnionLocalized  []TranslationStringEntry
}

type SymbolProvider interface {
	WalkGlobalSymbols(walkFn func(symbol.Symbol) error, types SymbolType) error
	LookupGlobalSymbol(name string, types SymbolType) (symbol.Symbol, bool)
}

// NewDaedalusStatefulListener ...
func NewDaedalusStatefulListener(source string, knownSymbols SymbolProvider) *DaedalusStatefulListener {
	return &DaedalusStatefulListener{
		source:          source,
		GlobalVariables: map[string]symbol.Symbol{},
		GlobalConstants: map[string]symbol.Symbol{},
		Functions:       map[string]symbol.Function{},
		Classes:         map[string]symbol.Class{},
		Prototypes:      map[string]symbol.ProtoTypeOrInstance{},
		Instances:       map[string]symbol.ProtoTypeOrInstance{},
		summaryBuilder:  &bytes.Buffer{},
		knownSymbols:    knownSymbols,

		// Translations
		GlobalDialogues: []Dialogue{},
		StringLiterals:  []StringLiteral{},
		SVMs:            []SVM{},
		//	newConstantNames: make(map[string]int),
		UnionLocalized: []TranslationStringEntry{},
	}
}

func symbolDefinitionForRuleContext(ctx antlr.ParserRuleContext) symbol.Definition {
	return symbol.NewDefinition(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
}

func (l *DaedalusStatefulListener) symbolSummaryForContext(ctx antlr.ParserRuleContext) string {
	l.summaryBuilder.Reset()

	walkSymbols(ctx, func(sum *parser.SymbolSummaryContext) error {
		if l.summaryBuilder.Len() > 0 {
			l.summaryBuilder.WriteRune('\n')
		}
		sumReplaced := strings.TrimLeft(sum.GetText(), "/ \t")
		sumReplaced = strings.ReplaceAll(sumReplaced, "/// ", "  \n")
		l.summaryBuilder.WriteString(sumReplaced)

		return nil
	})

	return l.summaryBuilder.String()
}

func (l *DaedalusStatefulListener) variablesFromContext(v *parser.VarDeclContext) []symbol.Symbol {
	result := []symbol.Symbol{}
	summary := ""
	if p, ok := v.GetParent().(parser.IInlineDefContext); ok {
		summary = l.symbolSummaryForContext(p)
	}

	for _, ch := range v.GetChildren() {
		if val, ok := ch.(*parser.VarValueDeclContext); ok {
			result = append(result,
				symbol.NewVariable(val.NameNode().GetText(),
					v.TypeReference().GetText(),
					l.source,
					summary, // documentation
					symbolDefinitionForRuleContext(val.NameNode()),
				))
		} else if innerVal, ok := ch.(*parser.VarDeclContext); ok {
			for _, ival := range innerVal.AllVarValueDecl() {
				val := ival.(*parser.VarValueDeclContext)
				result = append(result,
					symbol.NewVariable(val.NameNode().GetText(),
						v.TypeReference().GetText(),
						l.source,
						summary, // documentation
						symbolDefinitionForRuleContext(val.NameNode()),
					))
			}
		} else if innerVal, ok := ch.(*parser.VarArrayDeclContext); ok {
			result = append(result,
				symbol.NewArrayVariable(innerVal.NameNode().GetText(),
					v.TypeReference().GetText(),
					innerVal.ArraySize().GetText(),
					l.source,
					summary, // documentation
					symbolDefinitionForRuleContext(innerVal.NameNode()),
				))
		}
	}

	return result
}

func (l *DaedalusStatefulListener) maxNOfConstValues(n int, c *parser.ConstArrayAssignmentContext) string {
	result := &strings.Builder{}
	result.WriteString("{ ")
	counter := 0

	walkSymbols(c, func(v parser.IExpressionBlockContext) error {
		if counter > n {
			return fmt.Errorf("ok")
		}
		if counter > 0 {
			result.WriteString(", ")
		}
		result.WriteString(v.(*parser.ExpressionBlockContext).GetText())
		counter++
		return nil
	})

	if counter == 0 {
		result.WriteString("}")
	} else {
		result.WriteString(" ... }")
	}
	return result.String()
}

func arrayElementsFromContext(c *parser.ConstArrayAssignmentContext) []symbol.ArrayElement {
	result := make([]symbol.ArrayElement, 0, 66)
	for i, v := range c.AllExpressionBlock() {
		result = append(result, symbol.NewArrayElement(
			i,
			v.GetText(),
			symbol.NewDefinition(v.GetStart().GetLine(),
				v.GetStart().GetColumn(),
				v.GetStop().GetLine(),
				v.GetStop().GetColumn(),
			)))
	}
	return result
}

func (l *DaedalusStatefulListener) constsFromContext(c *parser.ConstDefContext) []symbol.Symbol {
	result := []symbol.Symbol{}
	summary := ""
	if p, ok := c.GetParent().(parser.IInlineDefContext); ok {
		summary = l.symbolSummaryForContext(p)
	}

	walkSymbols(c, func(cv *parser.ConstValueDefContext) error {
		result = append(result,
			symbol.NewConstant(cv.NameNode().GetText(),
				c.TypeReference().GetText(),
				l.source,
				summary, // documentation
				symbolDefinitionForRuleContext(cv.NameNode()),
				cv.ConstValueAssignment().(*parser.ConstValueAssignmentContext).ExpressionBlock().GetText(),
			))
		return nil
	})

	walkSymbols(c, func(innerVal *parser.ConstArrayDefContext) error {
		result = append(result,
			symbol.NewConstantArray(innerVal.NameNode().GetText(),
				c.TypeReference().GetText(),
				innerVal.ArraySize().GetText(),
				l.source,
				summary, // documentation
				symbolDefinitionForRuleContext(innerVal.NameNode()),
				l.maxNOfConstValues(3, innerVal.ConstArrayAssignment().(*parser.ConstArrayAssignmentContext)),
				arrayElementsFromContext(innerVal.ConstArrayAssignment().(*parser.ConstArrayAssignmentContext)),
			))
		return nil
	})

	return result
}

// EnterInlineDef ...
func (l *DaedalusStatefulListener) EnterInlineDef(ctx *parser.InlineDefContext) {
	for _, ch := range ctx.GetChildren() {
		if c, ok := ch.(*parser.ConstDefContext); ok {
			for _, s := range l.constsFromContext(c) {
				l.GlobalConstants[strings.ToUpper(s.Name())] = s
			}
		} else if v, ok := ch.(*parser.VarDeclContext); ok {
			for _, s := range l.variablesFromContext(v) {
				l.GlobalVariables[strings.ToUpper(s.Name())] = s
			}
		} else if c, ok := ch.(*parser.InstanceDeclContext); ok {
			walkSymbols(c, func(name parser.INameNodeContext) error {
				psym := symbol.NewPrototypeOrInstance(
					name.GetText(),
					c.ParentReference().GetText(),
					l.source,
					"",
					symbolDefinitionForRuleContext(name),
					symbolDefinitionForRuleContext(c.ParentReference()),
					true)
				l.Instances[strings.ToUpper(psym.Name())] = psym
				return nil
			})
		}
	}
}

func walkSymbols[T antlr.ParserRuleContext](rule antlr.RuleNode, walkFn func(symbol T) error) {
	for _, v := range rule.GetChildren() {
		if s, ok := v.(T); ok {
			if err := walkFn(s); err != nil {
				return
			}
		}
	}
}

// EnterBlockDef ...
func (l *DaedalusStatefulListener) EnterBlockDef(ctx *parser.BlockDefContext) {
	// Classes
	walkSymbols(ctx, func(c *parser.ClassDefContext) error {
		cFields := []symbol.Symbol{}
		walkSymbols(c, func(v *parser.VarDeclContext) error {
			for _, ch := range v.GetChildren() {
				if vv, ok := ch.(*parser.VarValueDeclContext); ok {
					cFields = append(cFields,
						symbol.NewVariable(vv.NameNode().GetText(),
							v.TypeReference().GetText(),
							l.source,
							"",
							symbolDefinitionForRuleContext(vv.NameNode())),
					)
				} else if vv, ok := ch.(*parser.VarArrayDeclContext); ok {
					cFields = append(cFields,
						symbol.NewArrayVariable(vv.NameNode().GetText(),
							v.TypeReference().GetText(),
							vv.ArraySize().GetText(),
							l.source,
							"",
							symbolDefinitionForRuleContext(vv.NameNode())),
					)
				}
			}
			return nil
		})

		csym := symbol.NewClass(c.NameNode().GetText(),
			l.source,
			l.symbolSummaryForContext(c),
			symbolDefinitionForRuleContext(c.NameNode()),
			symbol.NewDefinition(
				c.GetChild(2).(antlr.TerminalNode).GetSymbol().GetLine(),
				c.GetChild(2).(antlr.TerminalNode).GetSymbol().GetColumn(),
				c.GetChild(c.GetChildCount()-1).(antlr.TerminalNode).GetSymbol().GetLine(),
				c.GetChild(c.GetChildCount()-1).(antlr.TerminalNode).GetSymbol().GetColumn(),
			),
			cFields)
		l.Classes[strings.ToUpper(csym.Name())] = csym
		return nil
	})

	for _, ch := range ctx.GetChildren() {
		if c, ok := ch.(*parser.PrototypeDefContext); ok {
			psym := symbol.NewPrototypeOrInstance(
				c.NameNode().GetText(),
				c.ParentReference().GetText(),
				l.source,
				"",
				symbolDefinitionForRuleContext(c.NameNode()),
				symbolDefinitionForRuleContext(c.StatementBlock()),
				false)
			l.Prototypes[strings.ToUpper(psym.Name())] = psym
		} else if c, ok := ch.(*parser.InstanceDefContext); ok {
			psym := symbol.NewPrototypeOrInstance(
				c.NameNode().GetText(),
				c.ParentReference().GetText(),
				l.source,
				"",
				symbolDefinitionForRuleContext(c.NameNode()),
				symbolDefinitionForRuleContext(c.StatementBlock()),
				true)
			l.Instances[strings.ToUpper(psym.Name())] = psym
		}
	}
}

func (l *DaedalusStatefulListener) findVarsConstsInStatements(root antlr.Tree) []symbol.Symbol {
	result := []symbol.Symbol{}

	for _, s := range root.GetChildren() {
		if varDecl, ok := s.(*parser.VarDeclContext); ok {
			result = append(result, l.variablesFromContext(varDecl)...)
		} else if constDef, ok := s.(*parser.ConstDefContext); ok {
			result = append(result, l.constsFromContext(constDef)...)
		} else if s.GetChildCount() > 0 {
			result = append(result, l.findVarsConstsInStatements(s)...)
		}
	}

	return result
}

// EnterFunctionDef ...
func (l *DaedalusStatefulListener) EnterFunctionDef(ctx *parser.FunctionDefContext) {
	statements := ctx.StatementBlock().(*parser.StatementBlockContext)

	// Parse dialogues
	l.GlobalDialogues = append(l.GlobalDialogues, l.findAI_OutputFuncCallsStatements(statements)...)

	params := []symbol.Variable{}
	locals := l.findVarsConstsInStatements(statements)

	walkSymbols(ctx.ParameterList(), func(pdef *parser.ParameterDeclContext) error {
		params = append(params,
			symbol.NewVariable(pdef.NameNode().GetText(),
				pdef.TypeReference().GetText(),
				l.source,
				"",
				symbolDefinitionForRuleContext(pdef.NameNode())))
		return nil
	})

	fnc := symbol.NewFunction(ctx.NameNode().GetText(),
		l.source,
		l.symbolSummaryForContext(ctx.GetParent().(*parser.BlockDefContext)),
		symbolDefinitionForRuleContext(ctx.NameNode()),
		ctx.TypeReference().GetText(),
		symbolDefinitionForRuleContext(ctx.StatementBlock()),
		params,
		locals)

	l.Functions[strings.ToUpper(fnc.Name())] = fnc
}

// Searches the statement block for AI_Output calls
func (l *DaedalusStatefulListener) findAI_OutputFuncCallsStatements(root antlr.Tree) []Dialogue {
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

func (l *DaedalusStatefulListener) EnterFuncCall(ctx *parser.FuncCallContext) {

	funcName := ctx.NameNode().GetText()

	if strings.EqualFold(funcName, "Str_GetLocalizedString") {
		l.UnionLocalized = append(l.UnionLocalized, NewTranslationStringEntry(
			l.getContext(ctx),
			ctx.AllFuncArgExpression()[0].GetText(),
			ctx.AllFuncArgExpression()[1].GetText(),
			ctx.AllFuncArgExpression()[2].GetText(),
			ctx.AllFuncArgExpression()[3].GetText(),
			"",
			"",
			"",
			"",
		))
	} else if strings.EqualFold(funcName, "Str_GetLocalizedStringEx") {
		l.UnionLocalized = append(l.UnionLocalized, NewTranslationStringEntry(
			l.getContext(ctx),
			ctx.AllFuncArgExpression()[0].GetText(),
			ctx.AllFuncArgExpression()[1].GetText(),
			ctx.AllFuncArgExpression()[2].GetText(),
			ctx.AllFuncArgExpression()[3].GetText(),
			ctx.AllFuncArgExpression()[4].GetText(),
			ctx.AllFuncArgExpression()[5].GetText(),
			ctx.AllFuncArgExpression()[6].GetText(),
			ctx.AllFuncArgExpression()[7].GetText(),
		))
	}
}

func (l *DaedalusStatefulListener) getContext(root antlr.Tree) string {
	if root == nil {
		return ""
	}
	var inst, memvar string
	if ass, ok := root.GetParent().GetParent().GetParent().GetParent().GetParent().GetParent().GetParent().(*parser.InstanceDefContext); ok {
		inst = ass.NameNode().GetText()
	}

	if ass, ok := root.GetParent().GetParent().GetParent().GetParent().(*parser.AssignmentContext); ok {
		memvar = ass.Reference().GetText()
	}

	return inst + "." + memvar
}

/*
func (l *DaedalusStatefulListener) findStr_GetLocalizedStringEx(root antlr.Tree) []Dialogue {
	var foundCalls []Dialogue
	for _, s := range root.GetChildren() {
		if funcCall, ok := s.(*parser.FuncCallContext); ok {
			if strings.EqualFold(funcCall.NameNode().GetText(), "Str_GetLocalizedString") {
				// TODO: Parse languages and the left side of the assignment

				l.UnionLocalized = append(l.UnionLocalized, NewTranslationStringEntry(
					"TODO",
					funcCall.AllFuncArgExpression()[0].GetText(),
					funcCall.AllFuncArgExpression()[1].GetText(),
					funcCall.AllFuncArgExpression()[2].GetText(),
					funcCall.AllFuncArgExpression()[3].GetText(),
					"",
					"",
					"",
					"",
				))
			} else if strings.EqualFold(funcCall.NameNode().GetText(), "Str_GetLocalizedStringEx") {
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
*/
var NewConstantNames = make(map[string]int)

//NewConstantNames = make(map[string]int)

func (l *DaedalusStatefulListener) EnterStringLiteralValue(ctx *parser.StringLiteralValueContext) {
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
					fncCtx.AllFuncArgExpression()[0].GetText()+"."+fncCtx.AllFuncArgExpression()[2].GetText(),
					"",
					"Info_AddChoice", // translation comment
				),
			)
			return
		}
		// avoid black listed functions
		// TODO: add user defined function blacklist
		if IsBlacklisted(fncName) {
			return
		} else {
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
		// skip them, they are already parsed
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
			"",
		),
	)

}

// Checks, wheter thi string literal is assigned to the `field` specified
func (l *DaedalusStatefulListener) DidFindMemberVarStringLiteral(list *[]StringLiteral, ass *parser.AssignmentContext, field string) bool {
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
						"",
					),
				)
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
