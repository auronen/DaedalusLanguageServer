// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type DaedalusLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var daedaluslexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func daedaluslexerLexerInit() {
	staticData := &daedaluslexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "','", "'<<'", "'>>'", "'<='", "'>='", "'=='", "'!='", "'%'", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "'('",
		"')'", "'['", "']'", "'{'", "'}'", "'&'", "'&&'", "'|'", "'||'", "'+'",
		"'-'", "'/'", "'*'", "'~'", "'!'", "'='", "'<'", "'>'", "'+='", "'-='",
		"'*='", "'/='", "'&='", "'|='", "'.'", "';'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "IntegerLiteral", "FloatLiteral",
		"StringLiteral", "Const", "Var", "If", "Int", "Else", "Func", "StringKeyword",
		"Class", "Void", "Return", "Float", "Prototype", "Instance", "Null",
		"LeftParen", "RightParen", "LeftBracket", "RightBracket", "LeftBrace",
		"RightBrace", "BitAnd", "And", "BitOr", "Or", "Plus", "Minus", "Div",
		"Star", "Tilde", "Not", "Assign", "Less", "Greater", "PlusAssign", "MinusAssign",
		"StarAssign", "DivAssign", "AndAssign", "OrAssign", "Dot", "Semi", "Identifier",
		"Whitespace", "Newline", "BlockComment", "LineComment",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "IntegerLiteral",
		"FloatLiteral", "StringLiteral", "Const", "Var", "If", "Int", "Else",
		"Func", "StringKeyword", "Class", "Void", "Return", "Float", "Prototype",
		"Instance", "Null", "LeftParen", "RightParen", "LeftBracket", "RightBracket",
		"LeftBrace", "RightBrace", "BitAnd", "And", "BitOr", "Or", "Plus", "Minus",
		"Div", "Star", "Tilde", "Not", "Assign", "Less", "Greater", "PlusAssign",
		"MinusAssign", "StarAssign", "DivAssign", "AndAssign", "OrAssign", "Dot",
		"Semi", "Identifier", "Whitespace", "Newline", "BlockComment", "LineComment",
		"NonDigit", "IdContinue", "IdSpecial", "GermanCharacter", "Digit", "PointFloat",
		"ExponentFloat", "Exponent",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 57, 421, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 8, 4, 8, 155, 8, 8, 11, 8, 12, 8, 156, 1, 9, 1,
		9, 3, 9, 161, 8, 9, 1, 10, 1, 10, 5, 10, 165, 8, 10, 10, 10, 12, 10, 168,
		9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1,
		29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1,
		39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44,
		1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1,
		47, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51,
		1, 52, 1, 52, 3, 52, 318, 8, 52, 1, 52, 5, 52, 321, 8, 52, 10, 52, 12,
		52, 324, 9, 52, 1, 53, 4, 53, 327, 8, 53, 11, 53, 12, 53, 328, 1, 53, 1,
		53, 1, 54, 1, 54, 3, 54, 335, 8, 54, 1, 54, 3, 54, 338, 8, 54, 1, 54, 1,
		54, 1, 55, 1, 55, 1, 55, 1, 55, 5, 55, 346, 8, 55, 10, 55, 12, 55, 349,
		9, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 56, 5,
		56, 360, 8, 56, 10, 56, 12, 56, 363, 9, 56, 1, 56, 1, 56, 1, 57, 1, 57,
		3, 57, 369, 8, 57, 1, 58, 1, 58, 1, 58, 3, 58, 374, 8, 58, 1, 59, 1, 59,
		1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 5, 62, 383, 8, 62, 10, 62, 12, 62, 386,
		9, 62, 1, 62, 1, 62, 4, 62, 390, 8, 62, 11, 62, 12, 62, 391, 1, 62, 4,
		62, 395, 8, 62, 11, 62, 12, 62, 396, 1, 62, 1, 62, 3, 62, 401, 8, 62, 1,
		63, 4, 63, 404, 8, 63, 11, 63, 12, 63, 405, 1, 63, 3, 63, 409, 8, 63, 1,
		63, 1, 63, 1, 64, 1, 64, 3, 64, 415, 8, 64, 1, 64, 4, 64, 418, 8, 64, 11,
		64, 12, 64, 419, 1, 347, 0, 65, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25,
		51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34,
		69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43,
		87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52,
		105, 53, 107, 54, 109, 55, 111, 56, 113, 57, 115, 0, 117, 0, 119, 0, 121,
		0, 123, 0, 125, 0, 127, 0, 129, 0, 1, 0, 25, 3, 0, 10, 10, 13, 13, 34,
		34, 2, 0, 67, 67, 99, 99, 2, 0, 79, 79, 111, 111, 2, 0, 78, 78, 110, 110,
		2, 0, 83, 83, 115, 115, 2, 0, 84, 84, 116, 116, 2, 0, 86, 86, 118, 118,
		2, 0, 65, 65, 97, 97, 2, 0, 82, 82, 114, 114, 2, 0, 73, 73, 105, 105, 2,
		0, 70, 70, 102, 102, 2, 0, 69, 69, 101, 101, 2, 0, 76, 76, 108, 108, 2,
		0, 85, 85, 117, 117, 2, 0, 71, 71, 103, 103, 2, 0, 68, 68, 100, 100, 2,
		0, 80, 80, 112, 112, 2, 0, 89, 89, 121, 121, 2, 0, 9, 9, 32, 32, 2, 0,
		10, 10, 13, 13, 3, 0, 65, 90, 95, 95, 97, 122, 2, 0, 64, 64, 94, 94, 7,
		0, 196, 196, 214, 214, 220, 220, 223, 223, 228, 228, 246, 246, 252, 252,
		1, 0, 48, 57, 2, 0, 43, 43, 45, 45, 433, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0,
		0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1,
		0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73,
		1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0,
		81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0,
		0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0,
		0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1,
		0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0,
		111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 1, 131, 1, 0, 0, 0, 3, 133, 1, 0,
		0, 0, 5, 136, 1, 0, 0, 0, 7, 139, 1, 0, 0, 0, 9, 142, 1, 0, 0, 0, 11, 145,
		1, 0, 0, 0, 13, 148, 1, 0, 0, 0, 15, 151, 1, 0, 0, 0, 17, 154, 1, 0, 0,
		0, 19, 160, 1, 0, 0, 0, 21, 162, 1, 0, 0, 0, 23, 171, 1, 0, 0, 0, 25, 177,
		1, 0, 0, 0, 27, 181, 1, 0, 0, 0, 29, 184, 1, 0, 0, 0, 31, 188, 1, 0, 0,
		0, 33, 193, 1, 0, 0, 0, 35, 198, 1, 0, 0, 0, 37, 205, 1, 0, 0, 0, 39, 211,
		1, 0, 0, 0, 41, 216, 1, 0, 0, 0, 43, 223, 1, 0, 0, 0, 45, 229, 1, 0, 0,
		0, 47, 239, 1, 0, 0, 0, 49, 248, 1, 0, 0, 0, 51, 253, 1, 0, 0, 0, 53, 255,
		1, 0, 0, 0, 55, 257, 1, 0, 0, 0, 57, 259, 1, 0, 0, 0, 59, 261, 1, 0, 0,
		0, 61, 263, 1, 0, 0, 0, 63, 265, 1, 0, 0, 0, 65, 267, 1, 0, 0, 0, 67, 270,
		1, 0, 0, 0, 69, 272, 1, 0, 0, 0, 71, 275, 1, 0, 0, 0, 73, 277, 1, 0, 0,
		0, 75, 279, 1, 0, 0, 0, 77, 281, 1, 0, 0, 0, 79, 283, 1, 0, 0, 0, 81, 285,
		1, 0, 0, 0, 83, 287, 1, 0, 0, 0, 85, 289, 1, 0, 0, 0, 87, 291, 1, 0, 0,
		0, 89, 293, 1, 0, 0, 0, 91, 296, 1, 0, 0, 0, 93, 299, 1, 0, 0, 0, 95, 302,
		1, 0, 0, 0, 97, 305, 1, 0, 0, 0, 99, 308, 1, 0, 0, 0, 101, 311, 1, 0, 0,
		0, 103, 313, 1, 0, 0, 0, 105, 317, 1, 0, 0, 0, 107, 326, 1, 0, 0, 0, 109,
		337, 1, 0, 0, 0, 111, 341, 1, 0, 0, 0, 113, 355, 1, 0, 0, 0, 115, 368,
		1, 0, 0, 0, 117, 373, 1, 0, 0, 0, 119, 375, 1, 0, 0, 0, 121, 377, 1, 0,
		0, 0, 123, 379, 1, 0, 0, 0, 125, 400, 1, 0, 0, 0, 127, 408, 1, 0, 0, 0,
		129, 412, 1, 0, 0, 0, 131, 132, 5, 44, 0, 0, 132, 2, 1, 0, 0, 0, 133, 134,
		5, 60, 0, 0, 134, 135, 5, 60, 0, 0, 135, 4, 1, 0, 0, 0, 136, 137, 5, 62,
		0, 0, 137, 138, 5, 62, 0, 0, 138, 6, 1, 0, 0, 0, 139, 140, 5, 60, 0, 0,
		140, 141, 5, 61, 0, 0, 141, 8, 1, 0, 0, 0, 142, 143, 5, 62, 0, 0, 143,
		144, 5, 61, 0, 0, 144, 10, 1, 0, 0, 0, 145, 146, 5, 61, 0, 0, 146, 147,
		5, 61, 0, 0, 147, 12, 1, 0, 0, 0, 148, 149, 5, 33, 0, 0, 149, 150, 5, 61,
		0, 0, 150, 14, 1, 0, 0, 0, 151, 152, 5, 37, 0, 0, 152, 16, 1, 0, 0, 0,
		153, 155, 3, 123, 61, 0, 154, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156,
		154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 18, 1, 0, 0, 0, 158, 161, 3,
		125, 62, 0, 159, 161, 3, 127, 63, 0, 160, 158, 1, 0, 0, 0, 160, 159, 1,
		0, 0, 0, 161, 20, 1, 0, 0, 0, 162, 166, 5, 34, 0, 0, 163, 165, 8, 0, 0,
		0, 164, 163, 1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166,
		167, 1, 0, 0, 0, 167, 169, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169, 170,
		5, 34, 0, 0, 170, 22, 1, 0, 0, 0, 171, 172, 7, 1, 0, 0, 172, 173, 7, 2,
		0, 0, 173, 174, 7, 3, 0, 0, 174, 175, 7, 4, 0, 0, 175, 176, 7, 5, 0, 0,
		176, 24, 1, 0, 0, 0, 177, 178, 7, 6, 0, 0, 178, 179, 7, 7, 0, 0, 179, 180,
		7, 8, 0, 0, 180, 26, 1, 0, 0, 0, 181, 182, 7, 9, 0, 0, 182, 183, 7, 10,
		0, 0, 183, 28, 1, 0, 0, 0, 184, 185, 7, 9, 0, 0, 185, 186, 7, 3, 0, 0,
		186, 187, 7, 5, 0, 0, 187, 30, 1, 0, 0, 0, 188, 189, 7, 11, 0, 0, 189,
		190, 7, 12, 0, 0, 190, 191, 7, 4, 0, 0, 191, 192, 7, 11, 0, 0, 192, 32,
		1, 0, 0, 0, 193, 194, 7, 10, 0, 0, 194, 195, 7, 13, 0, 0, 195, 196, 7,
		3, 0, 0, 196, 197, 7, 1, 0, 0, 197, 34, 1, 0, 0, 0, 198, 199, 7, 4, 0,
		0, 199, 200, 7, 5, 0, 0, 200, 201, 7, 8, 0, 0, 201, 202, 7, 9, 0, 0, 202,
		203, 7, 3, 0, 0, 203, 204, 7, 14, 0, 0, 204, 36, 1, 0, 0, 0, 205, 206,
		7, 1, 0, 0, 206, 207, 7, 12, 0, 0, 207, 208, 7, 7, 0, 0, 208, 209, 7, 4,
		0, 0, 209, 210, 7, 4, 0, 0, 210, 38, 1, 0, 0, 0, 211, 212, 7, 6, 0, 0,
		212, 213, 7, 2, 0, 0, 213, 214, 7, 9, 0, 0, 214, 215, 7, 15, 0, 0, 215,
		40, 1, 0, 0, 0, 216, 217, 7, 8, 0, 0, 217, 218, 7, 11, 0, 0, 218, 219,
		7, 5, 0, 0, 219, 220, 7, 13, 0, 0, 220, 221, 7, 8, 0, 0, 221, 222, 7, 3,
		0, 0, 222, 42, 1, 0, 0, 0, 223, 224, 7, 10, 0, 0, 224, 225, 7, 12, 0, 0,
		225, 226, 7, 2, 0, 0, 226, 227, 7, 7, 0, 0, 227, 228, 7, 5, 0, 0, 228,
		44, 1, 0, 0, 0, 229, 230, 7, 16, 0, 0, 230, 231, 7, 8, 0, 0, 231, 232,
		7, 2, 0, 0, 232, 233, 7, 5, 0, 0, 233, 234, 7, 2, 0, 0, 234, 235, 7, 5,
		0, 0, 235, 236, 7, 17, 0, 0, 236, 237, 7, 16, 0, 0, 237, 238, 7, 11, 0,
		0, 238, 46, 1, 0, 0, 0, 239, 240, 7, 9, 0, 0, 240, 241, 7, 3, 0, 0, 241,
		242, 7, 4, 0, 0, 242, 243, 7, 5, 0, 0, 243, 244, 7, 7, 0, 0, 244, 245,
		7, 3, 0, 0, 245, 246, 7, 1, 0, 0, 246, 247, 7, 11, 0, 0, 247, 48, 1, 0,
		0, 0, 248, 249, 7, 3, 0, 0, 249, 250, 7, 13, 0, 0, 250, 251, 7, 12, 0,
		0, 251, 252, 7, 12, 0, 0, 252, 50, 1, 0, 0, 0, 253, 254, 5, 40, 0, 0, 254,
		52, 1, 0, 0, 0, 255, 256, 5, 41, 0, 0, 256, 54, 1, 0, 0, 0, 257, 258, 5,
		91, 0, 0, 258, 56, 1, 0, 0, 0, 259, 260, 5, 93, 0, 0, 260, 58, 1, 0, 0,
		0, 261, 262, 5, 123, 0, 0, 262, 60, 1, 0, 0, 0, 263, 264, 5, 125, 0, 0,
		264, 62, 1, 0, 0, 0, 265, 266, 5, 38, 0, 0, 266, 64, 1, 0, 0, 0, 267, 268,
		5, 38, 0, 0, 268, 269, 5, 38, 0, 0, 269, 66, 1, 0, 0, 0, 270, 271, 5, 124,
		0, 0, 271, 68, 1, 0, 0, 0, 272, 273, 5, 124, 0, 0, 273, 274, 5, 124, 0,
		0, 274, 70, 1, 0, 0, 0, 275, 276, 5, 43, 0, 0, 276, 72, 1, 0, 0, 0, 277,
		278, 5, 45, 0, 0, 278, 74, 1, 0, 0, 0, 279, 280, 5, 47, 0, 0, 280, 76,
		1, 0, 0, 0, 281, 282, 5, 42, 0, 0, 282, 78, 1, 0, 0, 0, 283, 284, 5, 126,
		0, 0, 284, 80, 1, 0, 0, 0, 285, 286, 5, 33, 0, 0, 286, 82, 1, 0, 0, 0,
		287, 288, 5, 61, 0, 0, 288, 84, 1, 0, 0, 0, 289, 290, 5, 60, 0, 0, 290,
		86, 1, 0, 0, 0, 291, 292, 5, 62, 0, 0, 292, 88, 1, 0, 0, 0, 293, 294, 5,
		43, 0, 0, 294, 295, 5, 61, 0, 0, 295, 90, 1, 0, 0, 0, 296, 297, 5, 45,
		0, 0, 297, 298, 5, 61, 0, 0, 298, 92, 1, 0, 0, 0, 299, 300, 5, 42, 0, 0,
		300, 301, 5, 61, 0, 0, 301, 94, 1, 0, 0, 0, 302, 303, 5, 47, 0, 0, 303,
		304, 5, 61, 0, 0, 304, 96, 1, 0, 0, 0, 305, 306, 5, 38, 0, 0, 306, 307,
		5, 61, 0, 0, 307, 98, 1, 0, 0, 0, 308, 309, 5, 124, 0, 0, 309, 310, 5,
		61, 0, 0, 310, 100, 1, 0, 0, 0, 311, 312, 5, 46, 0, 0, 312, 102, 1, 0,
		0, 0, 313, 314, 5, 59, 0, 0, 314, 104, 1, 0, 0, 0, 315, 318, 3, 115, 57,
		0, 316, 318, 3, 123, 61, 0, 317, 315, 1, 0, 0, 0, 317, 316, 1, 0, 0, 0,
		318, 322, 1, 0, 0, 0, 319, 321, 3, 117, 58, 0, 320, 319, 1, 0, 0, 0, 321,
		324, 1, 0, 0, 0, 322, 320, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323, 106,
		1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 325, 327, 7, 18, 0, 0, 326, 325, 1, 0,
		0, 0, 327, 328, 1, 0, 0, 0, 328, 326, 1, 0, 0, 0, 328, 329, 1, 0, 0, 0,
		329, 330, 1, 0, 0, 0, 330, 331, 6, 53, 0, 0, 331, 108, 1, 0, 0, 0, 332,
		334, 5, 13, 0, 0, 333, 335, 5, 10, 0, 0, 334, 333, 1, 0, 0, 0, 334, 335,
		1, 0, 0, 0, 335, 338, 1, 0, 0, 0, 336, 338, 5, 10, 0, 0, 337, 332, 1, 0,
		0, 0, 337, 336, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 340, 6, 54, 0, 0,
		340, 110, 1, 0, 0, 0, 341, 342, 5, 47, 0, 0, 342, 343, 5, 42, 0, 0, 343,
		347, 1, 0, 0, 0, 344, 346, 9, 0, 0, 0, 345, 344, 1, 0, 0, 0, 346, 349,
		1, 0, 0, 0, 347, 348, 1, 0, 0, 0, 347, 345, 1, 0, 0, 0, 348, 350, 1, 0,
		0, 0, 349, 347, 1, 0, 0, 0, 350, 351, 5, 42, 0, 0, 351, 352, 5, 47, 0,
		0, 352, 353, 1, 0, 0, 0, 353, 354, 6, 55, 0, 0, 354, 112, 1, 0, 0, 0, 355,
		356, 5, 47, 0, 0, 356, 357, 5, 47, 0, 0, 357, 361, 1, 0, 0, 0, 358, 360,
		8, 19, 0, 0, 359, 358, 1, 0, 0, 0, 360, 363, 1, 0, 0, 0, 361, 359, 1, 0,
		0, 0, 361, 362, 1, 0, 0, 0, 362, 364, 1, 0, 0, 0, 363, 361, 1, 0, 0, 0,
		364, 365, 6, 56, 1, 0, 365, 114, 1, 0, 0, 0, 366, 369, 3, 121, 60, 0, 367,
		369, 7, 20, 0, 0, 368, 366, 1, 0, 0, 0, 368, 367, 1, 0, 0, 0, 369, 116,
		1, 0, 0, 0, 370, 374, 3, 115, 57, 0, 371, 374, 3, 119, 59, 0, 372, 374,
		3, 123, 61, 0, 373, 370, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 373, 372, 1,
		0, 0, 0, 374, 118, 1, 0, 0, 0, 375, 376, 7, 21, 0, 0, 376, 120, 1, 0, 0,
		0, 377, 378, 7, 22, 0, 0, 378, 122, 1, 0, 0, 0, 379, 380, 7, 23, 0, 0,
		380, 124, 1, 0, 0, 0, 381, 383, 3, 123, 61, 0, 382, 381, 1, 0, 0, 0, 383,
		386, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 387,
		1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 387, 389, 5, 46, 0, 0, 388, 390, 3, 123,
		61, 0, 389, 388, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0, 391, 389, 1, 0, 0, 0,
		391, 392, 1, 0, 0, 0, 392, 401, 1, 0, 0, 0, 393, 395, 3, 123, 61, 0, 394,
		393, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 394, 1, 0, 0, 0, 396, 397,
		1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 399, 5, 46, 0, 0, 399, 401, 1, 0,
		0, 0, 400, 384, 1, 0, 0, 0, 400, 394, 1, 0, 0, 0, 401, 126, 1, 0, 0, 0,
		402, 404, 3, 123, 61, 0, 403, 402, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405,
		403, 1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406, 409, 1, 0, 0, 0, 407, 409,
		3, 125, 62, 0, 408, 403, 1, 0, 0, 0, 408, 407, 1, 0, 0, 0, 409, 410, 1,
		0, 0, 0, 410, 411, 3, 129, 64, 0, 411, 128, 1, 0, 0, 0, 412, 414, 7, 11,
		0, 0, 413, 415, 7, 24, 0, 0, 414, 413, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0,
		415, 417, 1, 0, 0, 0, 416, 418, 3, 123, 61, 0, 417, 416, 1, 0, 0, 0, 418,
		419, 1, 0, 0, 0, 419, 417, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420, 130,
		1, 0, 0, 0, 21, 0, 156, 160, 166, 317, 322, 328, 334, 337, 347, 361, 368,
		373, 384, 391, 396, 400, 405, 408, 414, 419, 2, 6, 0, 0, 0, 2, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// DaedalusLexerInit initializes any static state used to implement DaedalusLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewDaedalusLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func DaedalusLexerInit() {
	staticData := &daedaluslexerLexerStaticData
	staticData.once.Do(daedaluslexerLexerInit)
}

// NewDaedalusLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewDaedalusLexer(input antlr.CharStream) *DaedalusLexer {
	DaedalusLexerInit()
	l := new(DaedalusLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &daedaluslexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Daedalus.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DaedalusLexer tokens.
const (
	DaedalusLexerT__0           = 1
	DaedalusLexerT__1           = 2
	DaedalusLexerT__2           = 3
	DaedalusLexerT__3           = 4
	DaedalusLexerT__4           = 5
	DaedalusLexerT__5           = 6
	DaedalusLexerT__6           = 7
	DaedalusLexerT__7           = 8
	DaedalusLexerIntegerLiteral = 9
	DaedalusLexerFloatLiteral   = 10
	DaedalusLexerStringLiteral  = 11
	DaedalusLexerConst          = 12
	DaedalusLexerVar            = 13
	DaedalusLexerIf             = 14
	DaedalusLexerInt            = 15
	DaedalusLexerElse           = 16
	DaedalusLexerFunc           = 17
	DaedalusLexerStringKeyword  = 18
	DaedalusLexerClass          = 19
	DaedalusLexerVoid           = 20
	DaedalusLexerReturn         = 21
	DaedalusLexerFloat          = 22
	DaedalusLexerPrototype      = 23
	DaedalusLexerInstance       = 24
	DaedalusLexerNull           = 25
	DaedalusLexerLeftParen      = 26
	DaedalusLexerRightParen     = 27
	DaedalusLexerLeftBracket    = 28
	DaedalusLexerRightBracket   = 29
	DaedalusLexerLeftBrace      = 30
	DaedalusLexerRightBrace     = 31
	DaedalusLexerBitAnd         = 32
	DaedalusLexerAnd            = 33
	DaedalusLexerBitOr          = 34
	DaedalusLexerOr             = 35
	DaedalusLexerPlus           = 36
	DaedalusLexerMinus          = 37
	DaedalusLexerDiv            = 38
	DaedalusLexerStar           = 39
	DaedalusLexerTilde          = 40
	DaedalusLexerNot            = 41
	DaedalusLexerAssign         = 42
	DaedalusLexerLess           = 43
	DaedalusLexerGreater        = 44
	DaedalusLexerPlusAssign     = 45
	DaedalusLexerMinusAssign    = 46
	DaedalusLexerStarAssign     = 47
	DaedalusLexerDivAssign      = 48
	DaedalusLexerAndAssign      = 49
	DaedalusLexerOrAssign       = 50
	DaedalusLexerDot            = 51
	DaedalusLexerSemi           = 52
	DaedalusLexerIdentifier     = 53
	DaedalusLexerWhitespace     = 54
	DaedalusLexerNewline        = 55
	DaedalusLexerBlockComment   = 56
	DaedalusLexerLineComment    = 57
)

const (
	// Channel that contains comments
	COMMENTS int = 2
)
