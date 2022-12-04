package langserver

import (
	"strings"

	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
)

// ParseResult ...
type ParseResult struct {
	Instances       map[string]symbol.ProtoTypeOrInstance
	GlobalVariables map[string]symbol.Symbol
	GlobalConstants map[string]symbol.Symbol
	Functions       map[string]symbol.Function
	Classes         map[string]symbol.Class
	Prototypes      map[string]symbol.ProtoTypeOrInstance
	Source          string
	SyntaxErrors    []SyntaxError

	GlobalComments map[int]Comment
	// Dialogues
	GlobalDialogues []Dialogue
	StringLiterals  []StringLiteral
	SVMs            []SVM
	UnionLocalized  []TranslationStringEntry
	StringLocations map[string][]SymbolPosition
}

func (r *ParseResult) CountSymbols() int64 {
	numSymbols := int64(len(r.Classes)) +
		int64(len(r.Functions)) +
		int64(len(r.GlobalConstants)) +
		int64(len(r.GlobalVariables)) +
		int64(len(r.Instances)) +
		int64(len(r.Prototypes))

	for _, v := range r.Classes {
		numSymbols += int64(len(v.Fields))
	}

	for _, v := range r.Instances {
		numSymbols += int64(len(v.Fields))
	}
	for _, v := range r.Prototypes {
		numSymbols += int64(len(v.Fields))
	}

	// for _, v := range r.Functions {
	// 	numSymbols += int64(len(v.Parameters))
	// 	numSymbols += int64(len(v.LocalVariables))
	// }

	return numSymbols
}

func (parsedDoc *ParseResult) WalkScopedVariables(di symbol.DefinitionIndex, walkFn func(symbol.Symbol) bool) {
	for _, fn := range parsedDoc.Functions {
		if fn.BodyDefinition.InBBox(di) {
			for _, p := range fn.Parameters {
				if !walkFn(p) {
					return
				}
			}
			for _, p := range fn.LocalVariables {
				if !walkFn(p) {
					return
				}
			}
			break
		}
	}
}

func (parsedDoc *ParseResult) ScopedVariables(di symbol.DefinitionIndex) map[string]symbol.Symbol {
	locals := map[string]symbol.Symbol{}

	parsedDoc.WalkScopedVariables(di, func(s symbol.Symbol) bool {
		locals[strings.ToUpper(s.Name())] = s
		return true
	})

	return locals
}
