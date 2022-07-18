package langserver

import (
	"fmt"
	"strings"
)

// Symbol ...
type Symbol interface {
	Name() string
	Source() string
	Documentation() string
	Definition() SymbolDefinition
	String() string
}

var _ Symbol = (*FunctionSymbol)(nil)
var _ Symbol = (*VariableSymbol)(nil)
var _ Symbol = (*ConstantSymbol)(nil)
var _ Symbol = (*ConstantSymbol)(nil)
var _ Symbol = (*ConstantArraySymbol)(nil)
var _ Symbol = (*VariableArraySymbol)(nil)

// ----------------------------------------------------------------------------------------------------
// symbolBase ...
type symbolBase struct {
	NameValue           string
	SymbolSource        string
	SymbolDocumentation string
	SymbolDefinition    SymbolDefinition
}

func newSymbolBase(name, source, doc string, def SymbolDefinition) symbolBase {
	return symbolBase{
		NameValue:           name,
		SymbolSource:        source,
		SymbolDocumentation: doc,
		SymbolDefinition:    def,
	}
}

// Name ...
func (s symbolBase) Name() string {
	return s.NameValue
}

// Source ...
func (s symbolBase) Source() string {
	return s.SymbolSource
}

// Documentation ...
func (s symbolBase) Documentation() string {
	return s.SymbolDocumentation
}

// Definition ...
func (s symbolBase) Definition() SymbolDefinition {
	return s.SymbolDefinition
}

// ----------------------------------------------------------------------------------------------------
// FunctionSymbol ...
type FunctionSymbol struct {
	ReturnType     string
	Parameters     []VariableSymbol
	LocalVariables []Symbol
	symbolBase
	BodyDefinition SymbolDefinition
}

// NewFunctionSymbol ...
func NewFunctionSymbol(name, source, doc string, def SymbolDefinition, retType string, bodyDef SymbolDefinition, params []VariableSymbol, locals []Symbol) FunctionSymbol {
	return FunctionSymbol{
		symbolBase:     newSymbolBase(name, source, doc, def),
		ReturnType:     retType,
		BodyDefinition: bodyDef,
		Parameters:     params,
		LocalVariables: locals,
	}
}

// GetType ...
func (s FunctionSymbol) GetType() string {
	return s.ReturnType
}

// String ...
func (s FunctionSymbol) String() string {
	return "func " + s.ReturnType + " " + s.Name() + "(" + joinVars(s.Parameters) + ")"
}

// ----------------------------------------------------------------------------------------------------
// VariableSymbol ...
type VariableSymbol struct {
	Type string
	symbolBase
}

// NewVariableSymbol ...
func NewVariableSymbol(name, varType, source, documentation string, definiton SymbolDefinition) VariableSymbol {
	return VariableSymbol{
		Type:       varType,
		symbolBase: newSymbolBase(name, source, documentation, definiton),
	}
}

// String ...
func (s VariableSymbol) String() string {
	return "var " + s.Type + " " + s.Name()
}

// GetType ...
func (s VariableSymbol) GetType() string {
	return s.Type
}

// ----------------------------------------------------------------------------------------------------
// VariableArraySymbol ...
type VariableArraySymbol struct {
	Type           string
	ArraySizeText  string
	ArraySizeValue string
	symbolBase
}

// NewArrayVariableSymbol ...
func NewArrayVariableSymbol(name, varType, sizeText, source, documentation string, definiton SymbolDefinition, length string) VariableArraySymbol {
	return VariableArraySymbol{
		Type:           varType,
		ArraySizeText:  sizeText,
		ArraySizeValue: length,
		symbolBase:     newSymbolBase(name, source, documentation, definiton),
	}
}

// String ...
func (s VariableArraySymbol) String() string {
	return "var " + s.Type + " " + s.Name() + "[" /*+ s.ArraySizeText + ": " */ + s.ArraySizeValue + "]"
}

// GetType ...
func (s VariableArraySymbol) GetType() string {
	return s.Type
}

// Name ...
func (s VariableArraySymbol) Name() string {
	return s.NameValue
}

// Source ...
func (s VariableArraySymbol) Source() string {
	return s.SymbolSource
}

// Documentation ...
func (s VariableArraySymbol) Documentation() string {
	return s.SymbolDocumentation
}

// Definition ...
func (s VariableArraySymbol) Definition() SymbolDefinition {
	return s.SymbolDefinition
}

// ----------------------------------------------------------------------------------------------------
// ArrayElement ...
type ArrayElement struct {
	Index      int
	Value      string
	Definition SymbolDefinition
}

func newArrayElement(index int, value string, definiton SymbolDefinition) ArrayElement {
	return ArrayElement{
		Index:      index,
		Value:      value,
		Definition: definiton,
	}
}

func (el ArrayElement) GetValue() string {
	return el.Value
}

// ----------------------------------------------------------------------------------------------------
// ConstantSymbol ...
type ConstantSymbol struct {
	Value string
	VariableSymbol
}

// NewConstantSymbol ...
func NewConstantSymbol(name, varType, source, documentation string, definiton SymbolDefinition, value string) ConstantSymbol {
	return ConstantSymbol{
		VariableSymbol: NewVariableSymbol(name, varType, source, documentation, definiton),
		Value:          value,
	}
}

// String ...
func (s ConstantSymbol) String() string {
	return "const " + s.Type + " " + s.Name() + " = " + s.Value
}

// GetType ...
func (s ConstantSymbol) GetType() string {
	return s.Type
}

// ----------------------------------------------------------------------------------------------------
// ConstArraySymbol ...
type ConstantArraySymbol struct {
	symbolBase           // base symbol
	Type          string // type: string, int, float
	ArraySizeText string
	Values        []ArrayElement // array elements
}

// NewConstantArraySymbol ...
func NewConstantArraySymbol(name, arrType, source, documentation string, definiton SymbolDefinition, values []ArrayElement, sizeText string) ConstantArraySymbol {
	return ConstantArraySymbol{
		symbolBase:    newSymbolBase(name, source, documentation, definiton),
		Type:          arrType,
		Values:        values,
		ArraySizeText: sizeText,
	}
}

// Name ...
func (s ConstantArraySymbol) Name() string {
	return s.NameValue
}

// Source ...
func (s ConstantArraySymbol) Source() string {
	return s.SymbolSource
}

// Documentation ...
func (s ConstantArraySymbol) Documentation() string {
	return s.SymbolDocumentation
}

// Definition ...
func (s ConstantArraySymbol) Definition() SymbolDefinition {
	return s.SymbolDefinition
}

// Determine const array size by the amount of values (FIXME: good method?)
func (s ConstantArraySymbol) GetArrSize() int {
	return len(s.Values)
}

// Generates the list of elements for hover
func (s ConstantArraySymbol) HoverElements() string {
	result := "{ "
	for i, v := range s.Values {
		if i >= 3 {
			break
		}
		if i > 0 {
			result += ", " + v.GetValue()
		} else {
			result += v.GetValue()
		}
	}
	if len(s.Values) <= 3 {
		return result + "}"
	}
	return result + " ... }"
}

// String ...
func (s ConstantArraySymbol) String() string {
	return "const " + s.Type + " " + s.Name() + "[" /* + s.ArraySizeText + ": " */ + fmt.Sprint(s.GetArrSize()) + "]" + " = " + s.HoverElements()
}

// ----------------------------------------------------------------------------------------------------
// NewClassSymbol ...
func NewClassSymbol(name, source, documentation string, definiton SymbolDefinition, bodyDef SymbolDefinition, fields []Symbol) ClassSymbol {
	return ClassSymbol{
		symbolBase:     newSymbolBase(name, source, documentation, definiton),
		BodyDefinition: bodyDef,
		Fields:         fields,
	}
}

// ClassSymbol ...
type ClassSymbol struct {
	Fields []Symbol
	symbolBase
	BodyDefinition SymbolDefinition
}

// String ...
func (s ClassSymbol) String() string {
	return "class " + s.Name()
}

// ----------------------------------------------------------------------------------------------------
// ProtoTypeOrInstanceSymbol ...
type ProtoTypeOrInstanceSymbol struct {
	Parent string
	Fields []VariableSymbol
	symbolBase
	BodyDefinition SymbolDefinition
	IsInstance     bool
}

// NewPrototypeOrInstanceSymbol ...
func NewPrototypeOrInstanceSymbol(name, parent, source, documentation string, definiton SymbolDefinition, bodyDef SymbolDefinition, isInstance bool) ProtoTypeOrInstanceSymbol {
	return ProtoTypeOrInstanceSymbol{
		Parent:         parent,
		symbolBase:     newSymbolBase(name, source, documentation, definiton),
		IsInstance:     isInstance,
		BodyDefinition: bodyDef,
	}
}

// String ...
func (s ProtoTypeOrInstanceSymbol) String() string {
	if s.IsInstance {
		return "instance " + s.Name() + "(" + s.Parent + ")"
	}
	return "prototype " + s.Name() + "(" + s.Parent + ")"
}

// TODO: Refactor for generics, `[]T where T: Stringer`
func joinVars(symbols []VariableSymbol) string {
	syms := make([]string, 0, len(symbols))
	for _, s := range symbols {
		syms = append(syms, s.String())
	}
	return strings.Join(syms, ", ")
}

// SymbolDefinition ...
type SymbolDefinition struct {
	Start DefinitionIndex
	End   DefinitionIndex
}

// NewSymbolDefinition ...
func NewSymbolDefinition(startLine, startCol, endLine, endCol int) SymbolDefinition {
	return SymbolDefinition{
		Start: DefinitionIndex{
			Line:   startLine,
			Column: startCol,
		},
		End: DefinitionIndex{
			Line:   endLine,
			Column: endCol,
		},
	}
}

// InBBox ...
func (sd SymbolDefinition) InBBox(di DefinitionIndex) bool {
	if di.Line < sd.Start.Line {
		return false
	}
	if di.Line > sd.End.Line {
		return false
	}
	if sd.Start.Line <= di.Line && di.Line < sd.End.Line {
		return true
	}
	if di.Line == sd.End.Line && di.Column <= sd.End.Column {
		return true
	}
	if sd.Start.Line <= di.Line &&
		sd.Start.Column <= di.Column &&
		di.Line <= sd.End.Line &&
		di.Column <= sd.End.Column {
		return true
	}

	return false
}

// DefinitionIndex ...
type DefinitionIndex struct {
	Line   int
	Column int
}
