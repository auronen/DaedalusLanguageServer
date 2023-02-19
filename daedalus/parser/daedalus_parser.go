// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Daedalus

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type DaedalusParser struct {
	*antlr.BaseParser
}

var daedalusParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func daedalusParserInit() {
	staticData := &daedalusParserStaticData
	staticData.literalNames = []string{
		"", "','", "'...'", "'<<'", "'>>'", "'<='", "'>='", "'=='", "'!='",
		"'%'", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "'('", "')'", "'['",
		"']'", "'{'", "'}'", "'&'", "'&&'", "'|'", "'||'", "'+'", "'-'", "'/'",
		"'*'", "'~'", "'!'", "'='", "'<'", "'>'", "'+='", "'-='", "'*='", "'/='",
		"'&='", "'|='", "'.'", "';'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "IntegerLiteral", "FloatLiteral",
		"StringLiteral", "Var", "Const", "Class", "Prototype", "Instance", "If",
		"Else", "Null", "Return", "Namespace", "Func", "Int", "Float", "StringKeyword",
		"Void", "Event", "Meta", "While", "Continue", "Break", "Test", "Mif",
		"Melif", "Melse", "Mendif", "LeftParen", "RightParen", "LeftBracket",
		"RightBracket", "LeftBrace", "RightBrace", "BitAnd", "And", "BitOr",
		"Or", "Plus", "Minus", "Div", "Star", "Tilde", "Not", "Assign", "Less",
		"Greater", "PlusAssign", "MinusAssign", "StarAssign", "DivAssign", "AndAssign",
		"OrAssign", "Dot", "Semi", "Identifier", "Whitespace", "Newline", "BlockComment",
		"LineComment",
	}
	staticData.ruleNames = []string{
		"daedalusFile", "blockDef", "inlineDef", "macroBlock", "macroCondition",
		"macroElseBlock", "macroElseIfBlock", "macroIfBlock", "macroDef", "testCondition",
		"testBlock", "testBlockStatement", "functionDef", "constDef", "classDef",
		"prototypeDef", "instanceDef", "instanceDecl", "namespaceDef", "mainBlock",
		"contentBlock", "varDecl", "metaValue", "zParserExtenderMeta", "zParserExtenderMetaBlock",
		"constArrayDef", "constArrayAssignment", "constValueDef", "constValueAssignment",
		"varArrayDecl", "varValueDecl", "variadic", "parameterList", "parameterDecl",
		"statementBlock", "statement", "funcCall", "assignment", "ifCondition",
		"elseBlock", "elseIfBlock", "ifBlock", "ifBlockStatement", "returnStatement",
		"whileCondition", "whileBlock", "whileBlockStatement", "funcArgExpression",
		"expressionBlock", "expression", "arrayIndex", "arraySize", "value",
		"referenceAtom", "reference", "typeReference", "anyIdentifier", "nameNode",
		"parentReference", "assignmentOperator", "unaryOperator", "addOperator",
		"bitMoveOperator", "compOperator", "eqOperator", "multOperator", "binAndOperator",
		"binOrOperator", "logAndOperator", "logOrOperator",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 69, 610, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67, 2, 68,
		7, 68, 2, 69, 7, 69, 1, 0, 3, 0, 142, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 151, 8, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 3, 2, 158,
		8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 168, 8, 3,
		1, 3, 1, 3, 1, 3, 3, 3, 173, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 179, 8,
		5, 10, 5, 12, 5, 182, 9, 5, 1, 6, 1, 6, 1, 6, 5, 6, 187, 8, 6, 10, 6, 12,
		6, 190, 9, 6, 1, 7, 1, 7, 1, 7, 5, 7, 195, 8, 7, 10, 7, 12, 7, 198, 9,
		7, 1, 8, 1, 8, 5, 8, 202, 8, 8, 10, 8, 12, 8, 205, 9, 8, 1, 8, 3, 8, 208,
		8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		5, 11, 220, 8, 11, 10, 11, 12, 11, 223, 9, 11, 1, 11, 3, 11, 226, 8, 11,
		1, 12, 1, 12, 1, 12, 3, 12, 231, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 3, 13, 241, 8, 13, 1, 13, 1, 13, 1, 13, 3, 13,
		246, 8, 13, 5, 13, 248, 8, 13, 10, 13, 12, 13, 251, 9, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 259, 8, 14, 10, 14, 12, 14, 262, 9,
		14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5,
		17, 284, 8, 17, 10, 17, 12, 17, 287, 9, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 297, 8, 18, 10, 18, 12, 18, 300, 9,
		18, 1, 18, 1, 18, 1, 19, 3, 19, 305, 8, 19, 1, 19, 4, 19, 308, 8, 19, 11,
		19, 12, 19, 309, 1, 20, 1, 20, 1, 20, 3, 20, 315, 8, 20, 1, 21, 1, 21,
		1, 21, 1, 21, 3, 21, 321, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 327,
		8, 21, 5, 21, 329, 8, 21, 10, 21, 12, 21, 332, 9, 21, 1, 22, 4, 22, 335,
		8, 22, 11, 22, 12, 22, 336, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 24, 5, 24, 347, 8, 24, 10, 24, 12, 24, 350, 9, 24, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 5, 26, 366, 8, 26, 10, 26, 12, 26, 369, 9, 26, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 392,
		8, 32, 10, 32, 12, 32, 395, 9, 32, 3, 32, 397, 8, 32, 1, 32, 1, 32, 3,
		32, 401, 8, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 3, 33, 412, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3,
		34, 420, 8, 34, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 426, 8, 34, 10, 34,
		12, 34, 429, 9, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 3, 35, 440, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 5, 36,
		447, 8, 36, 10, 36, 12, 36, 450, 9, 36, 3, 36, 452, 8, 36, 1, 36, 1, 36,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 5, 42,
		476, 8, 42, 10, 42, 12, 42, 479, 9, 42, 1, 42, 3, 42, 482, 8, 42, 1, 43,
		1, 43, 3, 43, 486, 8, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1,
		46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 49, 1, 49, 3, 49, 509, 8, 49, 1, 49, 1, 49, 1, 49, 1,
		49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1,
		49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49,
		1, 49, 5, 49, 547, 8, 49, 10, 49, 12, 49, 550, 9, 49, 1, 50, 1, 50, 3,
		50, 554, 8, 50, 1, 51, 1, 51, 3, 51, 558, 8, 51, 1, 52, 1, 52, 1, 52, 1,
		52, 1, 52, 1, 52, 3, 52, 566, 8, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53,
		3, 53, 573, 8, 53, 1, 54, 1, 54, 1, 54, 3, 54, 578, 8, 54, 1, 55, 1, 55,
		1, 56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1,
		61, 1, 61, 1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65, 1, 65, 1, 66,
		1, 66, 1, 67, 1, 67, 1, 68, 1, 68, 1, 69, 1, 69, 1, 69, 12, 203, 221, 260,
		285, 298, 336, 348, 367, 393, 427, 448, 477, 1, 98, 70, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80,
		82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112,
		114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 0, 9,
		3, 0, 17, 17, 23, 27, 65, 65, 6, 0, 13, 13, 15, 17, 20, 20, 22, 27, 29,
		29, 65, 65, 2, 0, 54, 54, 57, 62, 2, 0, 48, 49, 52, 53, 1, 0, 48, 49, 1,
		0, 3, 4, 2, 0, 5, 6, 55, 56, 1, 0, 7, 8, 2, 0, 9, 9, 50, 51, 615, 0, 141,
		1, 0, 0, 0, 2, 150, 1, 0, 0, 0, 4, 157, 1, 0, 0, 0, 6, 172, 1, 0, 0, 0,
		8, 174, 1, 0, 0, 0, 10, 176, 1, 0, 0, 0, 12, 183, 1, 0, 0, 0, 14, 191,
		1, 0, 0, 0, 16, 199, 1, 0, 0, 0, 18, 211, 1, 0, 0, 0, 20, 213, 1, 0, 0,
		0, 22, 217, 1, 0, 0, 0, 24, 227, 1, 0, 0, 0, 26, 236, 1, 0, 0, 0, 28, 252,
		1, 0, 0, 0, 30, 265, 1, 0, 0, 0, 32, 272, 1, 0, 0, 0, 34, 279, 1, 0, 0,
		0, 36, 292, 1, 0, 0, 0, 38, 304, 1, 0, 0, 0, 40, 314, 1, 0, 0, 0, 42, 316,
		1, 0, 0, 0, 44, 334, 1, 0, 0, 0, 46, 338, 1, 0, 0, 0, 48, 343, 1, 0, 0,
		0, 50, 354, 1, 0, 0, 0, 52, 360, 1, 0, 0, 0, 54, 372, 1, 0, 0, 0, 56, 375,
		1, 0, 0, 0, 58, 378, 1, 0, 0, 0, 60, 383, 1, 0, 0, 0, 62, 385, 1, 0, 0,
		0, 64, 387, 1, 0, 0, 0, 66, 404, 1, 0, 0, 0, 68, 413, 1, 0, 0, 0, 70, 439,
		1, 0, 0, 0, 72, 441, 1, 0, 0, 0, 74, 455, 1, 0, 0, 0, 76, 459, 1, 0, 0,
		0, 78, 461, 1, 0, 0, 0, 80, 464, 1, 0, 0, 0, 82, 469, 1, 0, 0, 0, 84, 473,
		1, 0, 0, 0, 86, 483, 1, 0, 0, 0, 88, 487, 1, 0, 0, 0, 90, 489, 1, 0, 0,
		0, 92, 493, 1, 0, 0, 0, 94, 495, 1, 0, 0, 0, 96, 497, 1, 0, 0, 0, 98, 508,
		1, 0, 0, 0, 100, 553, 1, 0, 0, 0, 102, 557, 1, 0, 0, 0, 104, 565, 1, 0,
		0, 0, 106, 567, 1, 0, 0, 0, 108, 574, 1, 0, 0, 0, 110, 579, 1, 0, 0, 0,
		112, 581, 1, 0, 0, 0, 114, 583, 1, 0, 0, 0, 116, 585, 1, 0, 0, 0, 118,
		587, 1, 0, 0, 0, 120, 589, 1, 0, 0, 0, 122, 591, 1, 0, 0, 0, 124, 593,
		1, 0, 0, 0, 126, 595, 1, 0, 0, 0, 128, 597, 1, 0, 0, 0, 130, 599, 1, 0,
		0, 0, 132, 601, 1, 0, 0, 0, 134, 603, 1, 0, 0, 0, 136, 605, 1, 0, 0, 0,
		138, 607, 1, 0, 0, 0, 140, 142, 3, 38, 19, 0, 141, 140, 1, 0, 0, 0, 141,
		142, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 5, 0, 0, 1, 144, 1, 1,
		0, 0, 0, 145, 151, 3, 24, 12, 0, 146, 151, 3, 28, 14, 0, 147, 151, 3, 30,
		15, 0, 148, 151, 3, 32, 16, 0, 149, 151, 3, 36, 18, 0, 150, 145, 1, 0,
		0, 0, 150, 146, 1, 0, 0, 0, 150, 147, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0,
		150, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 5, 64, 0, 0, 153,
		3, 1, 0, 0, 0, 154, 158, 3, 26, 13, 0, 155, 158, 3, 42, 21, 0, 156, 158,
		3, 34, 17, 0, 157, 154, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 156, 1,
		0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 160, 5, 64, 0, 0, 160, 5, 1, 0, 0,
		0, 161, 173, 3, 2, 1, 0, 162, 163, 3, 70, 35, 0, 163, 164, 5, 64, 0, 0,
		164, 173, 1, 0, 0, 0, 165, 167, 3, 84, 42, 0, 166, 168, 5, 64, 0, 0, 167,
		166, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 173, 1, 0, 0, 0, 169, 170,
		3, 92, 46, 0, 170, 171, 5, 64, 0, 0, 171, 173, 1, 0, 0, 0, 172, 161, 1,
		0, 0, 0, 172, 162, 1, 0, 0, 0, 172, 165, 1, 0, 0, 0, 172, 169, 1, 0, 0,
		0, 173, 7, 1, 0, 0, 0, 174, 175, 3, 96, 48, 0, 175, 9, 1, 0, 0, 0, 176,
		180, 5, 36, 0, 0, 177, 179, 3, 6, 3, 0, 178, 177, 1, 0, 0, 0, 179, 182,
		1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 11, 1, 0,
		0, 0, 182, 180, 1, 0, 0, 0, 183, 184, 5, 35, 0, 0, 184, 188, 3, 8, 4, 0,
		185, 187, 3, 6, 3, 0, 186, 185, 1, 0, 0, 0, 187, 190, 1, 0, 0, 0, 188,
		186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 13, 1, 0, 0, 0, 190, 188, 1,
		0, 0, 0, 191, 192, 5, 34, 0, 0, 192, 196, 3, 8, 4, 0, 193, 195, 3, 6, 3,
		0, 194, 193, 1, 0, 0, 0, 195, 198, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196,
		197, 1, 0, 0, 0, 197, 15, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 203, 3,
		14, 7, 0, 200, 202, 3, 12, 6, 0, 201, 200, 1, 0, 0, 0, 202, 205, 1, 0,
		0, 0, 203, 204, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 204, 207, 1, 0, 0, 0,
		205, 203, 1, 0, 0, 0, 206, 208, 3, 10, 5, 0, 207, 206, 1, 0, 0, 0, 207,
		208, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 5, 37, 0, 0, 210, 17,
		1, 0, 0, 0, 211, 212, 3, 96, 48, 0, 212, 19, 1, 0, 0, 0, 213, 214, 5, 33,
		0, 0, 214, 215, 3, 18, 9, 0, 215, 216, 3, 68, 34, 0, 216, 21, 1, 0, 0,
		0, 217, 221, 3, 20, 10, 0, 218, 220, 3, 80, 40, 0, 219, 218, 1, 0, 0, 0,
		220, 223, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 222,
		225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 224, 226, 3, 78, 39, 0, 225, 224,
		1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 23, 1, 0, 0, 0, 227, 230, 5, 23,
		0, 0, 228, 231, 3, 110, 55, 0, 229, 231, 5, 28, 0, 0, 230, 228, 1, 0, 0,
		0, 230, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 233, 3, 114, 57, 0,
		233, 234, 3, 64, 32, 0, 234, 235, 3, 68, 34, 0, 235, 25, 1, 0, 0, 0, 236,
		237, 5, 14, 0, 0, 237, 240, 3, 110, 55, 0, 238, 241, 3, 54, 27, 0, 239,
		241, 3, 50, 25, 0, 240, 238, 1, 0, 0, 0, 240, 239, 1, 0, 0, 0, 241, 249,
		1, 0, 0, 0, 242, 245, 5, 1, 0, 0, 243, 246, 3, 54, 27, 0, 244, 246, 3,
		50, 25, 0, 245, 243, 1, 0, 0, 0, 245, 244, 1, 0, 0, 0, 246, 248, 1, 0,
		0, 0, 247, 242, 1, 0, 0, 0, 248, 251, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0,
		249, 250, 1, 0, 0, 0, 250, 27, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 252, 253,
		5, 15, 0, 0, 253, 254, 3, 114, 57, 0, 254, 260, 5, 42, 0, 0, 255, 256,
		3, 42, 21, 0, 256, 257, 5, 64, 0, 0, 257, 259, 1, 0, 0, 0, 258, 255, 1,
		0, 0, 0, 259, 262, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 260, 258, 1, 0, 0,
		0, 261, 263, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 263, 264, 5, 43, 0, 0, 264,
		29, 1, 0, 0, 0, 265, 266, 5, 16, 0, 0, 266, 267, 3, 114, 57, 0, 267, 268,
		5, 38, 0, 0, 268, 269, 3, 116, 58, 0, 269, 270, 5, 39, 0, 0, 270, 271,
		3, 68, 34, 0, 271, 31, 1, 0, 0, 0, 272, 273, 5, 17, 0, 0, 273, 274, 3,
		114, 57, 0, 274, 275, 5, 38, 0, 0, 275, 276, 3, 116, 58, 0, 276, 277, 5,
		39, 0, 0, 277, 278, 3, 68, 34, 0, 278, 33, 1, 0, 0, 0, 279, 280, 5, 17,
		0, 0, 280, 285, 3, 114, 57, 0, 281, 282, 5, 1, 0, 0, 282, 284, 3, 114,
		57, 0, 283, 281, 1, 0, 0, 0, 284, 287, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0,
		285, 283, 1, 0, 0, 0, 286, 288, 1, 0, 0, 0, 287, 285, 1, 0, 0, 0, 288,
		289, 5, 38, 0, 0, 289, 290, 3, 116, 58, 0, 290, 291, 5, 39, 0, 0, 291,
		35, 1, 0, 0, 0, 292, 293, 5, 22, 0, 0, 293, 294, 3, 114, 57, 0, 294, 298,
		5, 42, 0, 0, 295, 297, 3, 40, 20, 0, 296, 295, 1, 0, 0, 0, 297, 300, 1,
		0, 0, 0, 298, 299, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 299, 301, 1, 0, 0,
		0, 300, 298, 1, 0, 0, 0, 301, 302, 5, 43, 0, 0, 302, 37, 1, 0, 0, 0, 303,
		305, 3, 48, 24, 0, 304, 303, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 307,
		1, 0, 0, 0, 306, 308, 3, 40, 20, 0, 307, 306, 1, 0, 0, 0, 308, 309, 1,
		0, 0, 0, 309, 307, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 39, 1, 0, 0,
		0, 311, 315, 3, 2, 1, 0, 312, 315, 3, 4, 2, 0, 313, 315, 3, 16, 8, 0, 314,
		311, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314, 313, 1, 0, 0, 0, 315, 41, 1,
		0, 0, 0, 316, 317, 5, 13, 0, 0, 317, 320, 3, 110, 55, 0, 318, 321, 3, 60,
		30, 0, 319, 321, 3, 58, 29, 0, 320, 318, 1, 0, 0, 0, 320, 319, 1, 0, 0,
		0, 321, 330, 1, 0, 0, 0, 322, 326, 5, 1, 0, 0, 323, 327, 3, 42, 21, 0,
		324, 327, 3, 60, 30, 0, 325, 327, 3, 58, 29, 0, 326, 323, 1, 0, 0, 0, 326,
		324, 1, 0, 0, 0, 326, 325, 1, 0, 0, 0, 327, 329, 1, 0, 0, 0, 328, 322,
		1, 0, 0, 0, 329, 332, 1, 0, 0, 0, 330, 328, 1, 0, 0, 0, 330, 331, 1, 0,
		0, 0, 331, 43, 1, 0, 0, 0, 332, 330, 1, 0, 0, 0, 333, 335, 9, 0, 0, 0,
		334, 333, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 336,
		334, 1, 0, 0, 0, 337, 45, 1, 0, 0, 0, 338, 339, 3, 114, 57, 0, 339, 340,
		5, 54, 0, 0, 340, 341, 3, 44, 22, 0, 341, 342, 5, 64, 0, 0, 342, 47, 1,
		0, 0, 0, 343, 344, 5, 29, 0, 0, 344, 348, 5, 42, 0, 0, 345, 347, 3, 46,
		23, 0, 346, 345, 1, 0, 0, 0, 347, 350, 1, 0, 0, 0, 348, 349, 1, 0, 0, 0,
		348, 346, 1, 0, 0, 0, 349, 351, 1, 0, 0, 0, 350, 348, 1, 0, 0, 0, 351,
		352, 5, 43, 0, 0, 352, 353, 5, 64, 0, 0, 353, 49, 1, 0, 0, 0, 354, 355,
		3, 114, 57, 0, 355, 356, 5, 40, 0, 0, 356, 357, 3, 102, 51, 0, 357, 358,
		5, 41, 0, 0, 358, 359, 3, 52, 26, 0, 359, 51, 1, 0, 0, 0, 360, 361, 5,
		54, 0, 0, 361, 362, 5, 42, 0, 0, 362, 367, 3, 96, 48, 0, 363, 364, 5, 1,
		0, 0, 364, 366, 3, 96, 48, 0, 365, 363, 1, 0, 0, 0, 366, 369, 1, 0, 0,
		0, 367, 368, 1, 0, 0, 0, 367, 365, 1, 0, 0, 0, 368, 370, 1, 0, 0, 0, 369,
		367, 1, 0, 0, 0, 370, 371, 5, 43, 0, 0, 371, 53, 1, 0, 0, 0, 372, 373,
		3, 114, 57, 0, 373, 374, 3, 56, 28, 0, 374, 55, 1, 0, 0, 0, 375, 376, 5,
		54, 0, 0, 376, 377, 3, 96, 48, 0, 377, 57, 1, 0, 0, 0, 378, 379, 3, 114,
		57, 0, 379, 380, 5, 40, 0, 0, 380, 381, 3, 102, 51, 0, 381, 382, 5, 41,
		0, 0, 382, 59, 1, 0, 0, 0, 383, 384, 3, 114, 57, 0, 384, 61, 1, 0, 0, 0,
		385, 386, 5, 2, 0, 0, 386, 63, 1, 0, 0, 0, 387, 396, 5, 38, 0, 0, 388,
		393, 3, 66, 33, 0, 389, 390, 5, 1, 0, 0, 390, 392, 3, 66, 33, 0, 391, 389,
		1, 0, 0, 0, 392, 395, 1, 0, 0, 0, 393, 394, 1, 0, 0, 0, 393, 391, 1, 0,
		0, 0, 394, 397, 1, 0, 0, 0, 395, 393, 1, 0, 0, 0, 396, 388, 1, 0, 0, 0,
		396, 397, 1, 0, 0, 0, 397, 400, 1, 0, 0, 0, 398, 399, 5, 1, 0, 0, 399,
		401, 3, 62, 31, 0, 400, 398, 1, 0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 402,
		1, 0, 0, 0, 402, 403, 5, 39, 0, 0, 403, 65, 1, 0, 0, 0, 404, 405, 5, 13,
		0, 0, 405, 406, 3, 110, 55, 0, 406, 411, 3, 114, 57, 0, 407, 408, 5, 40,
		0, 0, 408, 409, 3, 102, 51, 0, 409, 410, 5, 41, 0, 0, 410, 412, 1, 0, 0,
		0, 411, 407, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412, 67, 1, 0, 0, 0, 413,
		427, 5, 42, 0, 0, 414, 415, 3, 70, 35, 0, 415, 416, 5, 64, 0, 0, 416, 426,
		1, 0, 0, 0, 417, 419, 3, 84, 42, 0, 418, 420, 5, 64, 0, 0, 419, 418, 1,
		0, 0, 0, 419, 420, 1, 0, 0, 0, 420, 426, 1, 0, 0, 0, 421, 426, 3, 16, 8,
		0, 422, 423, 3, 92, 46, 0, 423, 424, 5, 64, 0, 0, 424, 426, 1, 0, 0, 0,
		425, 414, 1, 0, 0, 0, 425, 417, 1, 0, 0, 0, 425, 421, 1, 0, 0, 0, 425,
		422, 1, 0, 0, 0, 426, 429, 1, 0, 0, 0, 427, 428, 1, 0, 0, 0, 427, 425,
		1, 0, 0, 0, 428, 430, 1, 0, 0, 0, 429, 427, 1, 0, 0, 0, 430, 431, 5, 43,
		0, 0, 431, 69, 1, 0, 0, 0, 432, 440, 3, 74, 37, 0, 433, 440, 3, 86, 43,
		0, 434, 440, 3, 26, 13, 0, 435, 440, 3, 42, 21, 0, 436, 440, 3, 98, 49,
		0, 437, 440, 5, 31, 0, 0, 438, 440, 5, 32, 0, 0, 439, 432, 1, 0, 0, 0,
		439, 433, 1, 0, 0, 0, 439, 434, 1, 0, 0, 0, 439, 435, 1, 0, 0, 0, 439,
		436, 1, 0, 0, 0, 439, 437, 1, 0, 0, 0, 439, 438, 1, 0, 0, 0, 440, 71, 1,
		0, 0, 0, 441, 442, 3, 114, 57, 0, 442, 451, 5, 38, 0, 0, 443, 448, 3, 94,
		47, 0, 444, 445, 5, 1, 0, 0, 445, 447, 3, 94, 47, 0, 446, 444, 1, 0, 0,
		0, 447, 450, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 448, 446, 1, 0, 0, 0, 449,
		452, 1, 0, 0, 0, 450, 448, 1, 0, 0, 0, 451, 443, 1, 0, 0, 0, 451, 452,
		1, 0, 0, 0, 452, 453, 1, 0, 0, 0, 453, 454, 5, 39, 0, 0, 454, 73, 1, 0,
		0, 0, 455, 456, 3, 108, 54, 0, 456, 457, 3, 118, 59, 0, 457, 458, 3, 96,
		48, 0, 458, 75, 1, 0, 0, 0, 459, 460, 3, 96, 48, 0, 460, 77, 1, 0, 0, 0,
		461, 462, 5, 19, 0, 0, 462, 463, 3, 68, 34, 0, 463, 79, 1, 0, 0, 0, 464,
		465, 5, 19, 0, 0, 465, 466, 5, 18, 0, 0, 466, 467, 3, 76, 38, 0, 467, 468,
		3, 68, 34, 0, 468, 81, 1, 0, 0, 0, 469, 470, 5, 18, 0, 0, 470, 471, 3,
		76, 38, 0, 471, 472, 3, 68, 34, 0, 472, 83, 1, 0, 0, 0, 473, 477, 3, 82,
		41, 0, 474, 476, 3, 80, 40, 0, 475, 474, 1, 0, 0, 0, 476, 479, 1, 0, 0,
		0, 477, 478, 1, 0, 0, 0, 477, 475, 1, 0, 0, 0, 478, 481, 1, 0, 0, 0, 479,
		477, 1, 0, 0, 0, 480, 482, 3, 78, 39, 0, 481, 480, 1, 0, 0, 0, 481, 482,
		1, 0, 0, 0, 482, 85, 1, 0, 0, 0, 483, 485, 5, 21, 0, 0, 484, 486, 3, 96,
		48, 0, 485, 484, 1, 0, 0, 0, 485, 486, 1, 0, 0, 0, 486, 87, 1, 0, 0, 0,
		487, 488, 3, 96, 48, 0, 488, 89, 1, 0, 0, 0, 489, 490, 5, 30, 0, 0, 490,
		491, 3, 88, 44, 0, 491, 492, 3, 68, 34, 0, 492, 91, 1, 0, 0, 0, 493, 494,
		3, 90, 45, 0, 494, 93, 1, 0, 0, 0, 495, 496, 3, 96, 48, 0, 496, 95, 1,
		0, 0, 0, 497, 498, 3, 98, 49, 0, 498, 97, 1, 0, 0, 0, 499, 500, 6, 49,
		-1, 0, 500, 501, 5, 38, 0, 0, 501, 502, 3, 98, 49, 0, 502, 503, 5, 39,
		0, 0, 503, 509, 1, 0, 0, 0, 504, 505, 3, 120, 60, 0, 505, 506, 3, 98, 49,
		11, 506, 509, 1, 0, 0, 0, 507, 509, 3, 104, 52, 0, 508, 499, 1, 0, 0, 0,
		508, 504, 1, 0, 0, 0, 508, 507, 1, 0, 0, 0, 509, 548, 1, 0, 0, 0, 510,
		511, 10, 10, 0, 0, 511, 512, 3, 130, 65, 0, 512, 513, 3, 98, 49, 11, 513,
		547, 1, 0, 0, 0, 514, 515, 10, 9, 0, 0, 515, 516, 3, 122, 61, 0, 516, 517,
		3, 98, 49, 10, 517, 547, 1, 0, 0, 0, 518, 519, 10, 8, 0, 0, 519, 520, 3,
		124, 62, 0, 520, 521, 3, 98, 49, 9, 521, 547, 1, 0, 0, 0, 522, 523, 10,
		7, 0, 0, 523, 524, 3, 126, 63, 0, 524, 525, 3, 98, 49, 8, 525, 547, 1,
		0, 0, 0, 526, 527, 10, 6, 0, 0, 527, 528, 3, 128, 64, 0, 528, 529, 3, 98,
		49, 7, 529, 547, 1, 0, 0, 0, 530, 531, 10, 5, 0, 0, 531, 532, 3, 132, 66,
		0, 532, 533, 3, 98, 49, 6, 533, 547, 1, 0, 0, 0, 534, 535, 10, 4, 0, 0,
		535, 536, 3, 134, 67, 0, 536, 537, 3, 98, 49, 5, 537, 547, 1, 0, 0, 0,
		538, 539, 10, 3, 0, 0, 539, 540, 3, 136, 68, 0, 540, 541, 3, 98, 49, 4,
		541, 547, 1, 0, 0, 0, 542, 543, 10, 2, 0, 0, 543, 544, 3, 138, 69, 0, 544,
		545, 3, 98, 49, 3, 545, 547, 1, 0, 0, 0, 546, 510, 1, 0, 0, 0, 546, 514,
		1, 0, 0, 0, 546, 518, 1, 0, 0, 0, 546, 522, 1, 0, 0, 0, 546, 526, 1, 0,
		0, 0, 546, 530, 1, 0, 0, 0, 546, 534, 1, 0, 0, 0, 546, 538, 1, 0, 0, 0,
		546, 542, 1, 0, 0, 0, 547, 550, 1, 0, 0, 0, 548, 546, 1, 0, 0, 0, 548,
		549, 1, 0, 0, 0, 549, 99, 1, 0, 0, 0, 550, 548, 1, 0, 0, 0, 551, 554, 5,
		10, 0, 0, 552, 554, 3, 106, 53, 0, 553, 551, 1, 0, 0, 0, 553, 552, 1, 0,
		0, 0, 554, 101, 1, 0, 0, 0, 555, 558, 5, 10, 0, 0, 556, 558, 3, 106, 53,
		0, 557, 555, 1, 0, 0, 0, 557, 556, 1, 0, 0, 0, 558, 103, 1, 0, 0, 0, 559,
		566, 5, 10, 0, 0, 560, 566, 5, 11, 0, 0, 561, 566, 5, 12, 0, 0, 562, 566,
		5, 20, 0, 0, 563, 566, 3, 72, 36, 0, 564, 566, 3, 108, 54, 0, 565, 559,
		1, 0, 0, 0, 565, 560, 1, 0, 0, 0, 565, 561, 1, 0, 0, 0, 565, 562, 1, 0,
		0, 0, 565, 563, 1, 0, 0, 0, 565, 564, 1, 0, 0, 0, 566, 105, 1, 0, 0, 0,
		567, 572, 3, 114, 57, 0, 568, 569, 5, 40, 0, 0, 569, 570, 3, 100, 50, 0,
		570, 571, 5, 41, 0, 0, 571, 573, 1, 0, 0, 0, 572, 568, 1, 0, 0, 0, 572,
		573, 1, 0, 0, 0, 573, 107, 1, 0, 0, 0, 574, 577, 3, 106, 53, 0, 575, 576,
		5, 63, 0, 0, 576, 578, 3, 106, 53, 0, 577, 575, 1, 0, 0, 0, 577, 578, 1,
		0, 0, 0, 578, 109, 1, 0, 0, 0, 579, 580, 7, 0, 0, 0, 580, 111, 1, 0, 0,
		0, 581, 582, 7, 1, 0, 0, 582, 113, 1, 0, 0, 0, 583, 584, 3, 112, 56, 0,
		584, 115, 1, 0, 0, 0, 585, 586, 5, 65, 0, 0, 586, 117, 1, 0, 0, 0, 587,
		588, 7, 2, 0, 0, 588, 119, 1, 0, 0, 0, 589, 590, 7, 3, 0, 0, 590, 121,
		1, 0, 0, 0, 591, 592, 7, 4, 0, 0, 592, 123, 1, 0, 0, 0, 593, 594, 7, 5,
		0, 0, 594, 125, 1, 0, 0, 0, 595, 596, 7, 6, 0, 0, 596, 127, 1, 0, 0, 0,
		597, 598, 7, 7, 0, 0, 598, 129, 1, 0, 0, 0, 599, 600, 7, 8, 0, 0, 600,
		131, 1, 0, 0, 0, 601, 602, 5, 44, 0, 0, 602, 133, 1, 0, 0, 0, 603, 604,
		5, 46, 0, 0, 604, 135, 1, 0, 0, 0, 605, 606, 5, 45, 0, 0, 606, 137, 1,
		0, 0, 0, 607, 608, 5, 47, 0, 0, 608, 139, 1, 0, 0, 0, 49, 141, 150, 157,
		167, 172, 180, 188, 196, 203, 207, 221, 225, 230, 240, 245, 249, 260, 285,
		298, 304, 309, 314, 320, 326, 330, 336, 348, 367, 393, 396, 400, 411, 419,
		425, 427, 439, 448, 451, 477, 481, 485, 508, 546, 548, 553, 557, 565, 572,
		577,
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

// DaedalusParserInit initializes any static state used to implement DaedalusParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDaedalusParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DaedalusParserInit() {
	staticData := &daedalusParserStaticData
	staticData.once.Do(daedalusParserInit)
}

// NewDaedalusParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDaedalusParser(input antlr.TokenStream) *DaedalusParser {
	DaedalusParserInit()
	this := new(DaedalusParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &daedalusParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// DaedalusParser tokens.
const (
	DaedalusParserEOF            = antlr.TokenEOF
	DaedalusParserT__0           = 1
	DaedalusParserT__1           = 2
	DaedalusParserT__2           = 3
	DaedalusParserT__3           = 4
	DaedalusParserT__4           = 5
	DaedalusParserT__5           = 6
	DaedalusParserT__6           = 7
	DaedalusParserT__7           = 8
	DaedalusParserT__8           = 9
	DaedalusParserIntegerLiteral = 10
	DaedalusParserFloatLiteral   = 11
	DaedalusParserStringLiteral  = 12
	DaedalusParserVar            = 13
	DaedalusParserConst          = 14
	DaedalusParserClass          = 15
	DaedalusParserPrototype      = 16
	DaedalusParserInstance       = 17
	DaedalusParserIf             = 18
	DaedalusParserElse           = 19
	DaedalusParserNull           = 20
	DaedalusParserReturn         = 21
	DaedalusParserNamespace      = 22
	DaedalusParserFunc           = 23
	DaedalusParserInt            = 24
	DaedalusParserFloat          = 25
	DaedalusParserStringKeyword  = 26
	DaedalusParserVoid           = 27
	DaedalusParserEvent          = 28
	DaedalusParserMeta           = 29
	DaedalusParserWhile          = 30
	DaedalusParserContinue       = 31
	DaedalusParserBreak          = 32
	DaedalusParserTest           = 33
	DaedalusParserMif            = 34
	DaedalusParserMelif          = 35
	DaedalusParserMelse          = 36
	DaedalusParserMendif         = 37
	DaedalusParserLeftParen      = 38
	DaedalusParserRightParen     = 39
	DaedalusParserLeftBracket    = 40
	DaedalusParserRightBracket   = 41
	DaedalusParserLeftBrace      = 42
	DaedalusParserRightBrace     = 43
	DaedalusParserBitAnd         = 44
	DaedalusParserAnd            = 45
	DaedalusParserBitOr          = 46
	DaedalusParserOr             = 47
	DaedalusParserPlus           = 48
	DaedalusParserMinus          = 49
	DaedalusParserDiv            = 50
	DaedalusParserStar           = 51
	DaedalusParserTilde          = 52
	DaedalusParserNot            = 53
	DaedalusParserAssign         = 54
	DaedalusParserLess           = 55
	DaedalusParserGreater        = 56
	DaedalusParserPlusAssign     = 57
	DaedalusParserMinusAssign    = 58
	DaedalusParserStarAssign     = 59
	DaedalusParserDivAssign      = 60
	DaedalusParserAndAssign      = 61
	DaedalusParserOrAssign       = 62
	DaedalusParserDot            = 63
	DaedalusParserSemi           = 64
	DaedalusParserIdentifier     = 65
	DaedalusParserWhitespace     = 66
	DaedalusParserNewline        = 67
	DaedalusParserBlockComment   = 68
	DaedalusParserLineComment    = 69
)

// DaedalusParser rules.
const (
	DaedalusParserRULE_daedalusFile             = 0
	DaedalusParserRULE_blockDef                 = 1
	DaedalusParserRULE_inlineDef                = 2
	DaedalusParserRULE_macroBlock               = 3
	DaedalusParserRULE_macroCondition           = 4
	DaedalusParserRULE_macroElseBlock           = 5
	DaedalusParserRULE_macroElseIfBlock         = 6
	DaedalusParserRULE_macroIfBlock             = 7
	DaedalusParserRULE_macroDef                 = 8
	DaedalusParserRULE_testCondition            = 9
	DaedalusParserRULE_testBlock                = 10
	DaedalusParserRULE_testBlockStatement       = 11
	DaedalusParserRULE_functionDef              = 12
	DaedalusParserRULE_constDef                 = 13
	DaedalusParserRULE_classDef                 = 14
	DaedalusParserRULE_prototypeDef             = 15
	DaedalusParserRULE_instanceDef              = 16
	DaedalusParserRULE_instanceDecl             = 17
	DaedalusParserRULE_namespaceDef             = 18
	DaedalusParserRULE_mainBlock                = 19
	DaedalusParserRULE_contentBlock             = 20
	DaedalusParserRULE_varDecl                  = 21
	DaedalusParserRULE_metaValue                = 22
	DaedalusParserRULE_zParserExtenderMeta      = 23
	DaedalusParserRULE_zParserExtenderMetaBlock = 24
	DaedalusParserRULE_constArrayDef            = 25
	DaedalusParserRULE_constArrayAssignment     = 26
	DaedalusParserRULE_constValueDef            = 27
	DaedalusParserRULE_constValueAssignment     = 28
	DaedalusParserRULE_varArrayDecl             = 29
	DaedalusParserRULE_varValueDecl             = 30
	DaedalusParserRULE_variadic                 = 31
	DaedalusParserRULE_parameterList            = 32
	DaedalusParserRULE_parameterDecl            = 33
	DaedalusParserRULE_statementBlock           = 34
	DaedalusParserRULE_statement                = 35
	DaedalusParserRULE_funcCall                 = 36
	DaedalusParserRULE_assignment               = 37
	DaedalusParserRULE_ifCondition              = 38
	DaedalusParserRULE_elseBlock                = 39
	DaedalusParserRULE_elseIfBlock              = 40
	DaedalusParserRULE_ifBlock                  = 41
	DaedalusParserRULE_ifBlockStatement         = 42
	DaedalusParserRULE_returnStatement          = 43
	DaedalusParserRULE_whileCondition           = 44
	DaedalusParserRULE_whileBlock               = 45
	DaedalusParserRULE_whileBlockStatement      = 46
	DaedalusParserRULE_funcArgExpression        = 47
	DaedalusParserRULE_expressionBlock          = 48
	DaedalusParserRULE_expression               = 49
	DaedalusParserRULE_arrayIndex               = 50
	DaedalusParserRULE_arraySize                = 51
	DaedalusParserRULE_value                    = 52
	DaedalusParserRULE_referenceAtom            = 53
	DaedalusParserRULE_reference                = 54
	DaedalusParserRULE_typeReference            = 55
	DaedalusParserRULE_anyIdentifier            = 56
	DaedalusParserRULE_nameNode                 = 57
	DaedalusParserRULE_parentReference          = 58
	DaedalusParserRULE_assignmentOperator       = 59
	DaedalusParserRULE_unaryOperator            = 60
	DaedalusParserRULE_addOperator              = 61
	DaedalusParserRULE_bitMoveOperator          = 62
	DaedalusParserRULE_compOperator             = 63
	DaedalusParserRULE_eqOperator               = 64
	DaedalusParserRULE_multOperator             = 65
	DaedalusParserRULE_binAndOperator           = 66
	DaedalusParserRULE_binOrOperator            = 67
	DaedalusParserRULE_logAndOperator           = 68
	DaedalusParserRULE_logOrOperator            = 69
)

// IDaedalusFileContext is an interface to support dynamic dispatch.
type IDaedalusFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDaedalusFileContext differentiates from other interfaces.
	IsDaedalusFileContext()
}

type DaedalusFileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDaedalusFileContext() *DaedalusFileContext {
	var p = new(DaedalusFileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_daedalusFile
	return p
}

func (*DaedalusFileContext) IsDaedalusFileContext() {}

func NewDaedalusFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DaedalusFileContext {
	var p = new(DaedalusFileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_daedalusFile

	return p
}

func (s *DaedalusFileContext) GetParser() antlr.Parser { return s.parser }

func (s *DaedalusFileContext) EOF() antlr.TerminalNode {
	return s.GetToken(DaedalusParserEOF, 0)
}

func (s *DaedalusFileContext) MainBlock() IMainBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainBlockContext)
}

func (s *DaedalusFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DaedalusFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DaedalusFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterDaedalusFile(s)
	}
}

func (s *DaedalusFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitDaedalusFile(s)
	}
}

func (p *DaedalusParser) DaedalusFile() (localctx IDaedalusFileContext) {
	this := p
	_ = this

	localctx = NewDaedalusFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DaedalusParserRULE_daedalusFile)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17729576960) != 0 {
		{
			p.SetState(140)
			p.MainBlock()
		}

	}
	{
		p.SetState(143)
		p.Match(DaedalusParserEOF)
	}

	return localctx
}

// IBlockDefContext is an interface to support dynamic dispatch.
type IBlockDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockDefContext differentiates from other interfaces.
	IsBlockDefContext()
}

type BlockDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockDefContext() *BlockDefContext {
	var p = new(BlockDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_blockDef
	return p
}

func (*BlockDefContext) IsBlockDefContext() {}

func NewBlockDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockDefContext {
	var p = new(BlockDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_blockDef

	return p
}

func (s *BlockDefContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockDefContext) Semi() antlr.TerminalNode {
	return s.GetToken(DaedalusParserSemi, 0)
}

func (s *BlockDefContext) FunctionDef() IFunctionDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefContext)
}

func (s *BlockDefContext) ClassDef() IClassDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassDefContext)
}

func (s *BlockDefContext) PrototypeDef() IPrototypeDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrototypeDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrototypeDefContext)
}

func (s *BlockDefContext) InstanceDef() IInstanceDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstanceDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstanceDefContext)
}

func (s *BlockDefContext) NamespaceDef() INamespaceDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespaceDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespaceDefContext)
}

func (s *BlockDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterBlockDef(s)
	}
}

func (s *BlockDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitBlockDef(s)
	}
}

func (p *DaedalusParser) BlockDef() (localctx IBlockDefContext) {
	this := p
	_ = this

	localctx = NewBlockDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DaedalusParserRULE_blockDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserFunc:
		{
			p.SetState(145)
			p.FunctionDef()
		}

	case DaedalusParserClass:
		{
			p.SetState(146)
			p.ClassDef()
		}

	case DaedalusParserPrototype:
		{
			p.SetState(147)
			p.PrototypeDef()
		}

	case DaedalusParserInstance:
		{
			p.SetState(148)
			p.InstanceDef()
		}

	case DaedalusParserNamespace:
		{
			p.SetState(149)
			p.NamespaceDef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(152)
		p.Match(DaedalusParserSemi)
	}

	return localctx
}

// IInlineDefContext is an interface to support dynamic dispatch.
type IInlineDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineDefContext differentiates from other interfaces.
	IsInlineDefContext()
}

type InlineDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineDefContext() *InlineDefContext {
	var p = new(InlineDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_inlineDef
	return p
}

func (*InlineDefContext) IsInlineDefContext() {}

func NewInlineDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineDefContext {
	var p = new(InlineDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_inlineDef

	return p
}

func (s *InlineDefContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineDefContext) Semi() antlr.TerminalNode {
	return s.GetToken(DaedalusParserSemi, 0)
}

func (s *InlineDefContext) ConstDef() IConstDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstDefContext)
}

func (s *InlineDefContext) VarDecl() IVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *InlineDefContext) InstanceDecl() IInstanceDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstanceDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstanceDeclContext)
}

func (s *InlineDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterInlineDef(s)
	}
}

func (s *InlineDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitInlineDef(s)
	}
}

func (p *DaedalusParser) InlineDef() (localctx IInlineDefContext) {
	this := p
	_ = this

	localctx = NewInlineDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DaedalusParserRULE_inlineDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(157)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserConst:
		{
			p.SetState(154)
			p.ConstDef()
		}

	case DaedalusParserVar:
		{
			p.SetState(155)
			p.VarDecl()
		}

	case DaedalusParserInstance:
		{
			p.SetState(156)
			p.InstanceDecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(159)
		p.Match(DaedalusParserSemi)
	}

	return localctx
}

// IMacroBlockContext is an interface to support dynamic dispatch.
type IMacroBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMacroBlockContext differentiates from other interfaces.
	IsMacroBlockContext()
}

type MacroBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacroBlockContext() *MacroBlockContext {
	var p = new(MacroBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_macroBlock
	return p
}

func (*MacroBlockContext) IsMacroBlockContext() {}

func NewMacroBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MacroBlockContext {
	var p = new(MacroBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_macroBlock

	return p
}

func (s *MacroBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *MacroBlockContext) BlockDef() IBlockDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockDefContext)
}

func (s *MacroBlockContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *MacroBlockContext) Semi() antlr.TerminalNode {
	return s.GetToken(DaedalusParserSemi, 0)
}

func (s *MacroBlockContext) IfBlockStatement() IIfBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockStatementContext)
}

func (s *MacroBlockContext) WhileBlockStatement() IWhileBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileBlockStatementContext)
}

func (s *MacroBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MacroBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterMacroBlock(s)
	}
}

func (s *MacroBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitMacroBlock(s)
	}
}

func (p *DaedalusParser) MacroBlock() (localctx IMacroBlockContext) {
	this := p
	_ = this

	localctx = NewMacroBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DaedalusParserRULE_macroBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(161)
			p.BlockDef()
		}

	case 2:
		{
			p.SetState(162)
			p.Statement()
		}
		{
			p.SetState(163)
			p.Match(DaedalusParserSemi)
		}

	case 3:
		{
			p.SetState(165)
			p.IfBlockStatement()
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DaedalusParserSemi {
			{
				p.SetState(166)
				p.Match(DaedalusParserSemi)
			}

		}

	case 4:
		{
			p.SetState(169)
			p.WhileBlockStatement()
		}
		{
			p.SetState(170)
			p.Match(DaedalusParserSemi)
		}

	}

	return localctx
}

// IMacroConditionContext is an interface to support dynamic dispatch.
type IMacroConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMacroConditionContext differentiates from other interfaces.
	IsMacroConditionContext()
}

type MacroConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacroConditionContext() *MacroConditionContext {
	var p = new(MacroConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_macroCondition
	return p
}

func (*MacroConditionContext) IsMacroConditionContext() {}

func NewMacroConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MacroConditionContext {
	var p = new(MacroConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_macroCondition

	return p
}

func (s *MacroConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *MacroConditionContext) ExpressionBlock() IExpressionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionBlockContext)
}

func (s *MacroConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MacroConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterMacroCondition(s)
	}
}

func (s *MacroConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitMacroCondition(s)
	}
}

func (p *DaedalusParser) MacroCondition() (localctx IMacroConditionContext) {
	this := p
	_ = this

	localctx = NewMacroConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DaedalusParserRULE_macroCondition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.ExpressionBlock()
	}

	return localctx
}

// IMacroElseBlockContext is an interface to support dynamic dispatch.
type IMacroElseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMacroElseBlockContext differentiates from other interfaces.
	IsMacroElseBlockContext()
}

type MacroElseBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacroElseBlockContext() *MacroElseBlockContext {
	var p = new(MacroElseBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_macroElseBlock
	return p
}

func (*MacroElseBlockContext) IsMacroElseBlockContext() {}

func NewMacroElseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MacroElseBlockContext {
	var p = new(MacroElseBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_macroElseBlock

	return p
}

func (s *MacroElseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *MacroElseBlockContext) Melse() antlr.TerminalNode {
	return s.GetToken(DaedalusParserMelse, 0)
}

func (s *MacroElseBlockContext) AllMacroBlock() []IMacroBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMacroBlockContext); ok {
			len++
		}
	}

	tst := make([]IMacroBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMacroBlockContext); ok {
			tst[i] = t.(IMacroBlockContext)
			i++
		}
	}

	return tst
}

func (s *MacroElseBlockContext) MacroBlock(i int) IMacroBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacroBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacroBlockContext)
}

func (s *MacroElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroElseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MacroElseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterMacroElseBlock(s)
	}
}

func (s *MacroElseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitMacroElseBlock(s)
	}
}

func (p *DaedalusParser) MacroElseBlock() (localctx IMacroElseBlockContext) {
	this := p
	_ = this

	localctx = NewMacroElseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DaedalusParserRULE_macroElseBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(DaedalusParserMelse)
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-10)) & ^0x3f) == 0 && ((int64(1)<<(_la-10))&36042816068779519) != 0 {
		{
			p.SetState(177)
			p.MacroBlock()
		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMacroElseIfBlockContext is an interface to support dynamic dispatch.
type IMacroElseIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMacroElseIfBlockContext differentiates from other interfaces.
	IsMacroElseIfBlockContext()
}

type MacroElseIfBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacroElseIfBlockContext() *MacroElseIfBlockContext {
	var p = new(MacroElseIfBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_macroElseIfBlock
	return p
}

func (*MacroElseIfBlockContext) IsMacroElseIfBlockContext() {}

func NewMacroElseIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MacroElseIfBlockContext {
	var p = new(MacroElseIfBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_macroElseIfBlock

	return p
}

func (s *MacroElseIfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *MacroElseIfBlockContext) Melif() antlr.TerminalNode {
	return s.GetToken(DaedalusParserMelif, 0)
}

func (s *MacroElseIfBlockContext) MacroCondition() IMacroConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacroConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacroConditionContext)
}

func (s *MacroElseIfBlockContext) AllMacroBlock() []IMacroBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMacroBlockContext); ok {
			len++
		}
	}

	tst := make([]IMacroBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMacroBlockContext); ok {
			tst[i] = t.(IMacroBlockContext)
			i++
		}
	}

	return tst
}

func (s *MacroElseIfBlockContext) MacroBlock(i int) IMacroBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacroBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacroBlockContext)
}

func (s *MacroElseIfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroElseIfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MacroElseIfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterMacroElseIfBlock(s)
	}
}

func (s *MacroElseIfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitMacroElseIfBlock(s)
	}
}

func (p *DaedalusParser) MacroElseIfBlock() (localctx IMacroElseIfBlockContext) {
	this := p
	_ = this

	localctx = NewMacroElseIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DaedalusParserRULE_macroElseIfBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(DaedalusParserMelif)
	}
	{
		p.SetState(184)
		p.MacroCondition()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-10)) & ^0x3f) == 0 && ((int64(1)<<(_la-10))&36042816068779519) != 0 {
		{
			p.SetState(185)
			p.MacroBlock()
		}

		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMacroIfBlockContext is an interface to support dynamic dispatch.
type IMacroIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMacroIfBlockContext differentiates from other interfaces.
	IsMacroIfBlockContext()
}

type MacroIfBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacroIfBlockContext() *MacroIfBlockContext {
	var p = new(MacroIfBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_macroIfBlock
	return p
}

func (*MacroIfBlockContext) IsMacroIfBlockContext() {}

func NewMacroIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MacroIfBlockContext {
	var p = new(MacroIfBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_macroIfBlock

	return p
}

func (s *MacroIfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *MacroIfBlockContext) Mif() antlr.TerminalNode {
	return s.GetToken(DaedalusParserMif, 0)
}

func (s *MacroIfBlockContext) MacroCondition() IMacroConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacroConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacroConditionContext)
}

func (s *MacroIfBlockContext) AllMacroBlock() []IMacroBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMacroBlockContext); ok {
			len++
		}
	}

	tst := make([]IMacroBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMacroBlockContext); ok {
			tst[i] = t.(IMacroBlockContext)
			i++
		}
	}

	return tst
}

func (s *MacroIfBlockContext) MacroBlock(i int) IMacroBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacroBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacroBlockContext)
}

func (s *MacroIfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroIfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MacroIfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterMacroIfBlock(s)
	}
}

func (s *MacroIfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitMacroIfBlock(s)
	}
}

func (p *DaedalusParser) MacroIfBlock() (localctx IMacroIfBlockContext) {
	this := p
	_ = this

	localctx = NewMacroIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DaedalusParserRULE_macroIfBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Match(DaedalusParserMif)
	}
	{
		p.SetState(192)
		p.MacroCondition()
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-10)) & ^0x3f) == 0 && ((int64(1)<<(_la-10))&36042816068779519) != 0 {
		{
			p.SetState(193)
			p.MacroBlock()
		}

		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMacroDefContext is an interface to support dynamic dispatch.
type IMacroDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMacroDefContext differentiates from other interfaces.
	IsMacroDefContext()
}

type MacroDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacroDefContext() *MacroDefContext {
	var p = new(MacroDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_macroDef
	return p
}

func (*MacroDefContext) IsMacroDefContext() {}

func NewMacroDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MacroDefContext {
	var p = new(MacroDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_macroDef

	return p
}

func (s *MacroDefContext) GetParser() antlr.Parser { return s.parser }

func (s *MacroDefContext) MacroIfBlock() IMacroIfBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacroIfBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacroIfBlockContext)
}

func (s *MacroDefContext) Mendif() antlr.TerminalNode {
	return s.GetToken(DaedalusParserMendif, 0)
}

func (s *MacroDefContext) AllMacroElseIfBlock() []IMacroElseIfBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMacroElseIfBlockContext); ok {
			len++
		}
	}

	tst := make([]IMacroElseIfBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMacroElseIfBlockContext); ok {
			tst[i] = t.(IMacroElseIfBlockContext)
			i++
		}
	}

	return tst
}

func (s *MacroDefContext) MacroElseIfBlock(i int) IMacroElseIfBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacroElseIfBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacroElseIfBlockContext)
}

func (s *MacroDefContext) MacroElseBlock() IMacroElseBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacroElseBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacroElseBlockContext)
}

func (s *MacroDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MacroDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterMacroDef(s)
	}
}

func (s *MacroDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitMacroDef(s)
	}
}

func (p *DaedalusParser) MacroDef() (localctx IMacroDefContext) {
	this := p
	_ = this

	localctx = NewMacroDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DaedalusParserRULE_macroDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.MacroIfBlock()
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(200)
				p.MacroElseIfBlock()
			}

		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserMelse {
		{
			p.SetState(206)
			p.MacroElseBlock()
		}

	}
	{
		p.SetState(209)
		p.Match(DaedalusParserMendif)
	}

	return localctx
}

// ITestConditionContext is an interface to support dynamic dispatch.
type ITestConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTestConditionContext differentiates from other interfaces.
	IsTestConditionContext()
}

type TestConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestConditionContext() *TestConditionContext {
	var p = new(TestConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_testCondition
	return p
}

func (*TestConditionContext) IsTestConditionContext() {}

func NewTestConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestConditionContext {
	var p = new(TestConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_testCondition

	return p
}

func (s *TestConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *TestConditionContext) ExpressionBlock() IExpressionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionBlockContext)
}

func (s *TestConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TestConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterTestCondition(s)
	}
}

func (s *TestConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitTestCondition(s)
	}
}

func (p *DaedalusParser) TestCondition() (localctx ITestConditionContext) {
	this := p
	_ = this

	localctx = NewTestConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DaedalusParserRULE_testCondition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.ExpressionBlock()
	}

	return localctx
}

// ITestBlockContext is an interface to support dynamic dispatch.
type ITestBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTestBlockContext differentiates from other interfaces.
	IsTestBlockContext()
}

type TestBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestBlockContext() *TestBlockContext {
	var p = new(TestBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_testBlock
	return p
}

func (*TestBlockContext) IsTestBlockContext() {}

func NewTestBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestBlockContext {
	var p = new(TestBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_testBlock

	return p
}

func (s *TestBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *TestBlockContext) Test() antlr.TerminalNode {
	return s.GetToken(DaedalusParserTest, 0)
}

func (s *TestBlockContext) TestCondition() ITestConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITestConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITestConditionContext)
}

func (s *TestBlockContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *TestBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TestBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterTestBlock(s)
	}
}

func (s *TestBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitTestBlock(s)
	}
}

func (p *DaedalusParser) TestBlock() (localctx ITestBlockContext) {
	this := p
	_ = this

	localctx = NewTestBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DaedalusParserRULE_testBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(DaedalusParserTest)
	}
	{
		p.SetState(214)
		p.TestCondition()
	}
	{
		p.SetState(215)
		p.StatementBlock()
	}

	return localctx
}

// ITestBlockStatementContext is an interface to support dynamic dispatch.
type ITestBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTestBlockStatementContext differentiates from other interfaces.
	IsTestBlockStatementContext()
}

type TestBlockStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestBlockStatementContext() *TestBlockStatementContext {
	var p = new(TestBlockStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_testBlockStatement
	return p
}

func (*TestBlockStatementContext) IsTestBlockStatementContext() {}

func NewTestBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestBlockStatementContext {
	var p = new(TestBlockStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_testBlockStatement

	return p
}

func (s *TestBlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *TestBlockStatementContext) TestBlock() ITestBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITestBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITestBlockContext)
}

func (s *TestBlockStatementContext) AllElseIfBlock() []IElseIfBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseIfBlockContext); ok {
			len++
		}
	}

	tst := make([]IElseIfBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseIfBlockContext); ok {
			tst[i] = t.(IElseIfBlockContext)
			i++
		}
	}

	return tst
}

func (s *TestBlockStatementContext) ElseIfBlock(i int) IElseIfBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseIfBlockContext)
}

func (s *TestBlockStatementContext) ElseBlock() IElseBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseBlockContext)
}

func (s *TestBlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestBlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TestBlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterTestBlockStatement(s)
	}
}

func (s *TestBlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitTestBlockStatement(s)
	}
}

func (p *DaedalusParser) TestBlockStatement() (localctx ITestBlockStatementContext) {
	this := p
	_ = this

	localctx = NewTestBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DaedalusParserRULE_testBlockStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.TestBlock()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(218)
				p.ElseIfBlock()
			}

		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserElse {
		{
			p.SetState(224)
			p.ElseBlock()
		}

	}

	return localctx
}

// IFunctionDefContext is an interface to support dynamic dispatch.
type IFunctionDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDefContext differentiates from other interfaces.
	IsFunctionDefContext()
}

type FunctionDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefContext() *FunctionDefContext {
	var p = new(FunctionDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_functionDef
	return p
}

func (*FunctionDefContext) IsFunctionDefContext() {}

func NewFunctionDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefContext {
	var p = new(FunctionDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_functionDef

	return p
}

func (s *FunctionDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefContext) Func() antlr.TerminalNode {
	return s.GetToken(DaedalusParserFunc, 0)
}

func (s *FunctionDefContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *FunctionDefContext) ParameterList() IParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionDefContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *FunctionDefContext) TypeReference() ITypeReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *FunctionDefContext) Event() antlr.TerminalNode {
	return s.GetToken(DaedalusParserEvent, 0)
}

func (s *FunctionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterFunctionDef(s)
	}
}

func (s *FunctionDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitFunctionDef(s)
	}
}

func (p *DaedalusParser) FunctionDef() (localctx IFunctionDefContext) {
	this := p
	_ = this

	localctx = NewFunctionDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DaedalusParserRULE_functionDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(DaedalusParserFunc)
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserInstance, DaedalusParserFunc, DaedalusParserInt, DaedalusParserFloat, DaedalusParserStringKeyword, DaedalusParserVoid, DaedalusParserIdentifier:
		{
			p.SetState(228)
			p.TypeReference()
		}

	case DaedalusParserEvent:
		{
			p.SetState(229)
			p.Match(DaedalusParserEvent)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(232)
		p.NameNode()
	}
	{
		p.SetState(233)
		p.ParameterList()
	}
	{
		p.SetState(234)
		p.StatementBlock()
	}

	return localctx
}

// IConstDefContext is an interface to support dynamic dispatch.
type IConstDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstDefContext differentiates from other interfaces.
	IsConstDefContext()
}

type ConstDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstDefContext() *ConstDefContext {
	var p = new(ConstDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_constDef
	return p
}

func (*ConstDefContext) IsConstDefContext() {}

func NewConstDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstDefContext {
	var p = new(ConstDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_constDef

	return p
}

func (s *ConstDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstDefContext) Const() antlr.TerminalNode {
	return s.GetToken(DaedalusParserConst, 0)
}

func (s *ConstDefContext) TypeReference() ITypeReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *ConstDefContext) AllConstValueDef() []IConstValueDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstValueDefContext); ok {
			len++
		}
	}

	tst := make([]IConstValueDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstValueDefContext); ok {
			tst[i] = t.(IConstValueDefContext)
			i++
		}
	}

	return tst
}

func (s *ConstDefContext) ConstValueDef(i int) IConstValueDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstValueDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstValueDefContext)
}

func (s *ConstDefContext) AllConstArrayDef() []IConstArrayDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstArrayDefContext); ok {
			len++
		}
	}

	tst := make([]IConstArrayDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstArrayDefContext); ok {
			tst[i] = t.(IConstArrayDefContext)
			i++
		}
	}

	return tst
}

func (s *ConstDefContext) ConstArrayDef(i int) IConstArrayDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstArrayDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstArrayDefContext)
}

func (s *ConstDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterConstDef(s)
	}
}

func (s *ConstDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitConstDef(s)
	}
}

func (p *DaedalusParser) ConstDef() (localctx IConstDefContext) {
	this := p
	_ = this

	localctx = NewConstDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DaedalusParserRULE_constDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(DaedalusParserConst)
	}
	{
		p.SetState(237)
		p.TypeReference()
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(238)
			p.ConstValueDef()
		}

	case 2:
		{
			p.SetState(239)
			p.ConstArrayDef()
		}

	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DaedalusParserT__0 {
		{
			p.SetState(242)
			p.Match(DaedalusParserT__0)
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(243)
				p.ConstValueDef()
			}

		case 2:
			{
				p.SetState(244)
				p.ConstArrayDef()
			}

		}

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IClassDefContext is an interface to support dynamic dispatch.
type IClassDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassDefContext differentiates from other interfaces.
	IsClassDefContext()
}

type ClassDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDefContext() *ClassDefContext {
	var p = new(ClassDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_classDef
	return p
}

func (*ClassDefContext) IsClassDefContext() {}

func NewClassDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDefContext {
	var p = new(ClassDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_classDef

	return p
}

func (s *ClassDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDefContext) Class() antlr.TerminalNode {
	return s.GetToken(DaedalusParserClass, 0)
}

func (s *ClassDefContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ClassDefContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBrace, 0)
}

func (s *ClassDefContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBrace, 0)
}

func (s *ClassDefContext) AllVarDecl() []IVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclContext); ok {
			tst[i] = t.(IVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *ClassDefContext) VarDecl(i int) IVarDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *ClassDefContext) AllSemi() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserSemi)
}

func (s *ClassDefContext) Semi(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserSemi, i)
}

func (s *ClassDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterClassDef(s)
	}
}

func (s *ClassDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitClassDef(s)
	}
}

func (p *DaedalusParser) ClassDef() (localctx IClassDefContext) {
	this := p
	_ = this

	localctx = NewClassDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DaedalusParserRULE_classDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Match(DaedalusParserClass)
	}
	{
		p.SetState(253)
		p.NameNode()
	}
	{
		p.SetState(254)
		p.Match(DaedalusParserLeftBrace)
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(255)
				p.VarDecl()
			}
			{
				p.SetState(256)
				p.Match(DaedalusParserSemi)
			}

		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}
	{
		p.SetState(263)
		p.Match(DaedalusParserRightBrace)
	}

	return localctx
}

// IPrototypeDefContext is an interface to support dynamic dispatch.
type IPrototypeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrototypeDefContext differentiates from other interfaces.
	IsPrototypeDefContext()
}

type PrototypeDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrototypeDefContext() *PrototypeDefContext {
	var p = new(PrototypeDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_prototypeDef
	return p
}

func (*PrototypeDefContext) IsPrototypeDefContext() {}

func NewPrototypeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrototypeDefContext {
	var p = new(PrototypeDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_prototypeDef

	return p
}

func (s *PrototypeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *PrototypeDefContext) Prototype() antlr.TerminalNode {
	return s.GetToken(DaedalusParserPrototype, 0)
}

func (s *PrototypeDefContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *PrototypeDefContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftParen, 0)
}

func (s *PrototypeDefContext) ParentReference() IParentReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParentReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParentReferenceContext)
}

func (s *PrototypeDefContext) RightParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightParen, 0)
}

func (s *PrototypeDefContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *PrototypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrototypeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrototypeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterPrototypeDef(s)
	}
}

func (s *PrototypeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitPrototypeDef(s)
	}
}

func (p *DaedalusParser) PrototypeDef() (localctx IPrototypeDefContext) {
	this := p
	_ = this

	localctx = NewPrototypeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DaedalusParserRULE_prototypeDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.Match(DaedalusParserPrototype)
	}
	{
		p.SetState(266)
		p.NameNode()
	}
	{
		p.SetState(267)
		p.Match(DaedalusParserLeftParen)
	}
	{
		p.SetState(268)
		p.ParentReference()
	}
	{
		p.SetState(269)
		p.Match(DaedalusParserRightParen)
	}
	{
		p.SetState(270)
		p.StatementBlock()
	}

	return localctx
}

// IInstanceDefContext is an interface to support dynamic dispatch.
type IInstanceDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstanceDefContext differentiates from other interfaces.
	IsInstanceDefContext()
}

type InstanceDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstanceDefContext() *InstanceDefContext {
	var p = new(InstanceDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_instanceDef
	return p
}

func (*InstanceDefContext) IsInstanceDefContext() {}

func NewInstanceDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstanceDefContext {
	var p = new(InstanceDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_instanceDef

	return p
}

func (s *InstanceDefContext) GetParser() antlr.Parser { return s.parser }

func (s *InstanceDefContext) Instance() antlr.TerminalNode {
	return s.GetToken(DaedalusParserInstance, 0)
}

func (s *InstanceDefContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *InstanceDefContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftParen, 0)
}

func (s *InstanceDefContext) ParentReference() IParentReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParentReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParentReferenceContext)
}

func (s *InstanceDefContext) RightParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightParen, 0)
}

func (s *InstanceDefContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *InstanceDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstanceDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterInstanceDef(s)
	}
}

func (s *InstanceDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitInstanceDef(s)
	}
}

func (p *DaedalusParser) InstanceDef() (localctx IInstanceDefContext) {
	this := p
	_ = this

	localctx = NewInstanceDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DaedalusParserRULE_instanceDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(DaedalusParserInstance)
	}
	{
		p.SetState(273)
		p.NameNode()
	}
	{
		p.SetState(274)
		p.Match(DaedalusParserLeftParen)
	}
	{
		p.SetState(275)
		p.ParentReference()
	}
	{
		p.SetState(276)
		p.Match(DaedalusParserRightParen)
	}
	{
		p.SetState(277)
		p.StatementBlock()
	}

	return localctx
}

// IInstanceDeclContext is an interface to support dynamic dispatch.
type IInstanceDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstanceDeclContext differentiates from other interfaces.
	IsInstanceDeclContext()
}

type InstanceDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstanceDeclContext() *InstanceDeclContext {
	var p = new(InstanceDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_instanceDecl
	return p
}

func (*InstanceDeclContext) IsInstanceDeclContext() {}

func NewInstanceDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstanceDeclContext {
	var p = new(InstanceDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_instanceDecl

	return p
}

func (s *InstanceDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *InstanceDeclContext) Instance() antlr.TerminalNode {
	return s.GetToken(DaedalusParserInstance, 0)
}

func (s *InstanceDeclContext) AllNameNode() []INameNodeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameNodeContext); ok {
			len++
		}
	}

	tst := make([]INameNodeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameNodeContext); ok {
			tst[i] = t.(INameNodeContext)
			i++
		}
	}

	return tst
}

func (s *InstanceDeclContext) NameNode(i int) INameNodeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *InstanceDeclContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftParen, 0)
}

func (s *InstanceDeclContext) ParentReference() IParentReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParentReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParentReferenceContext)
}

func (s *InstanceDeclContext) RightParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightParen, 0)
}

func (s *InstanceDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstanceDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterInstanceDecl(s)
	}
}

func (s *InstanceDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitInstanceDecl(s)
	}
}

func (p *DaedalusParser) InstanceDecl() (localctx IInstanceDeclContext) {
	this := p
	_ = this

	localctx = NewInstanceDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DaedalusParserRULE_instanceDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(DaedalusParserInstance)
	}
	{
		p.SetState(280)
		p.NameNode()
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(281)
				p.Match(DaedalusParserT__0)
			}
			{
				p.SetState(282)
				p.NameNode()
			}

		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}
	{
		p.SetState(288)
		p.Match(DaedalusParserLeftParen)
	}
	{
		p.SetState(289)
		p.ParentReference()
	}
	{
		p.SetState(290)
		p.Match(DaedalusParserRightParen)
	}

	return localctx
}

// INamespaceDefContext is an interface to support dynamic dispatch.
type INamespaceDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespaceDefContext differentiates from other interfaces.
	IsNamespaceDefContext()
}

type NamespaceDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceDefContext() *NamespaceDefContext {
	var p = new(NamespaceDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_namespaceDef
	return p
}

func (*NamespaceDefContext) IsNamespaceDefContext() {}

func NewNamespaceDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceDefContext {
	var p = new(NamespaceDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_namespaceDef

	return p
}

func (s *NamespaceDefContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceDefContext) Namespace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserNamespace, 0)
}

func (s *NamespaceDefContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *NamespaceDefContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBrace, 0)
}

func (s *NamespaceDefContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBrace, 0)
}

func (s *NamespaceDefContext) AllContentBlock() []IContentBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IContentBlockContext); ok {
			len++
		}
	}

	tst := make([]IContentBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IContentBlockContext); ok {
			tst[i] = t.(IContentBlockContext)
			i++
		}
	}

	return tst
}

func (s *NamespaceDefContext) ContentBlock(i int) IContentBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContentBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContentBlockContext)
}

func (s *NamespaceDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterNamespaceDef(s)
	}
}

func (s *NamespaceDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitNamespaceDef(s)
	}
}

func (p *DaedalusParser) NamespaceDef() (localctx INamespaceDefContext) {
	this := p
	_ = this

	localctx = NewNamespaceDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, DaedalusParserRULE_namespaceDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)
		p.Match(DaedalusParserNamespace)
	}
	{
		p.SetState(293)
		p.NameNode()
	}
	{
		p.SetState(294)
		p.Match(DaedalusParserLeftBrace)
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(295)
				p.ContentBlock()
			}

		}
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}
	{
		p.SetState(301)
		p.Match(DaedalusParserRightBrace)
	}

	return localctx
}

// IMainBlockContext is an interface to support dynamic dispatch.
type IMainBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMainBlockContext differentiates from other interfaces.
	IsMainBlockContext()
}

type MainBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainBlockContext() *MainBlockContext {
	var p = new(MainBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_mainBlock
	return p
}

func (*MainBlockContext) IsMainBlockContext() {}

func NewMainBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainBlockContext {
	var p = new(MainBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_mainBlock

	return p
}

func (s *MainBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *MainBlockContext) ZParserExtenderMetaBlock() IZParserExtenderMetaBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IZParserExtenderMetaBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IZParserExtenderMetaBlockContext)
}

func (s *MainBlockContext) AllContentBlock() []IContentBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IContentBlockContext); ok {
			len++
		}
	}

	tst := make([]IContentBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IContentBlockContext); ok {
			tst[i] = t.(IContentBlockContext)
			i++
		}
	}

	return tst
}

func (s *MainBlockContext) ContentBlock(i int) IContentBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContentBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContentBlockContext)
}

func (s *MainBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterMainBlock(s)
	}
}

func (s *MainBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitMainBlock(s)
	}
}

func (p *DaedalusParser) MainBlock() (localctx IMainBlockContext) {
	this := p
	_ = this

	localctx = NewMainBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DaedalusParserRULE_mainBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserMeta {
		{
			p.SetState(303)
			p.ZParserExtenderMetaBlock()
		}

	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17192706048) != 0 {
		{
			p.SetState(306)
			p.ContentBlock()
		}

		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IContentBlockContext is an interface to support dynamic dispatch.
type IContentBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContentBlockContext differentiates from other interfaces.
	IsContentBlockContext()
}

type ContentBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentBlockContext() *ContentBlockContext {
	var p = new(ContentBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_contentBlock
	return p
}

func (*ContentBlockContext) IsContentBlockContext() {}

func NewContentBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentBlockContext {
	var p = new(ContentBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_contentBlock

	return p
}

func (s *ContentBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentBlockContext) BlockDef() IBlockDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockDefContext)
}

func (s *ContentBlockContext) InlineDef() IInlineDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineDefContext)
}

func (s *ContentBlockContext) MacroDef() IMacroDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacroDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacroDefContext)
}

func (s *ContentBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterContentBlock(s)
	}
}

func (s *ContentBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitContentBlock(s)
	}
}

func (p *DaedalusParser) ContentBlock() (localctx IContentBlockContext) {
	this := p
	_ = this

	localctx = NewContentBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, DaedalusParserRULE_contentBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(311)
			p.BlockDef()
		}

	case 2:
		{
			p.SetState(312)
			p.InlineDef()
		}

	case 3:
		{
			p.SetState(313)
			p.MacroDef()
		}

	}

	return localctx
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_varDecl
	return p
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) Var() antlr.TerminalNode {
	return s.GetToken(DaedalusParserVar, 0)
}

func (s *VarDeclContext) TypeReference() ITypeReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *VarDeclContext) AllVarValueDecl() []IVarValueDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarValueDeclContext); ok {
			len++
		}
	}

	tst := make([]IVarValueDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarValueDeclContext); ok {
			tst[i] = t.(IVarValueDeclContext)
			i++
		}
	}

	return tst
}

func (s *VarDeclContext) VarValueDecl(i int) IVarValueDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarValueDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarValueDeclContext)
}

func (s *VarDeclContext) AllVarArrayDecl() []IVarArrayDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarArrayDeclContext); ok {
			len++
		}
	}

	tst := make([]IVarArrayDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarArrayDeclContext); ok {
			tst[i] = t.(IVarArrayDeclContext)
			i++
		}
	}

	return tst
}

func (s *VarDeclContext) VarArrayDecl(i int) IVarArrayDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarArrayDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarArrayDeclContext)
}

func (s *VarDeclContext) AllVarDecl() []IVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclContext); ok {
			tst[i] = t.(IVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *VarDeclContext) VarDecl(i int) IVarDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (p *DaedalusParser) VarDecl() (localctx IVarDeclContext) {
	this := p
	_ = this

	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, DaedalusParserRULE_varDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		p.Match(DaedalusParserVar)
	}
	{
		p.SetState(317)
		p.TypeReference()
	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(318)
			p.VarValueDecl()
		}

	case 2:
		{
			p.SetState(319)
			p.VarArrayDecl()
		}

	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(322)
				p.Match(DaedalusParserT__0)
			}
			p.SetState(326)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(323)
					p.VarDecl()
				}

			case 2:
				{
					p.SetState(324)
					p.VarValueDecl()
				}

			case 3:
				{
					p.SetState(325)
					p.VarArrayDecl()
				}

			}

		}
		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}

	return localctx
}

// IMetaValueContext is an interface to support dynamic dispatch.
type IMetaValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMetaValueContext differentiates from other interfaces.
	IsMetaValueContext()
}

type MetaValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetaValueContext() *MetaValueContext {
	var p = new(MetaValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_metaValue
	return p
}

func (*MetaValueContext) IsMetaValueContext() {}

func NewMetaValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetaValueContext {
	var p = new(MetaValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_metaValue

	return p
}

func (s *MetaValueContext) GetParser() antlr.Parser { return s.parser }
func (s *MetaValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetaValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetaValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterMetaValue(s)
	}
}

func (s *MetaValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitMetaValue(s)
	}
}

func (p *DaedalusParser) MetaValue() (localctx IMetaValueContext) {
	this := p
	_ = this

	localctx = NewMetaValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, DaedalusParserRULE_metaValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	_alt = 1 + 1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1 + 1:
			p.SetState(333)
			p.MatchWildcard()

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// IZParserExtenderMetaContext is an interface to support dynamic dispatch.
type IZParserExtenderMetaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsZParserExtenderMetaContext differentiates from other interfaces.
	IsZParserExtenderMetaContext()
}

type ZParserExtenderMetaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyZParserExtenderMetaContext() *ZParserExtenderMetaContext {
	var p = new(ZParserExtenderMetaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_zParserExtenderMeta
	return p
}

func (*ZParserExtenderMetaContext) IsZParserExtenderMetaContext() {}

func NewZParserExtenderMetaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ZParserExtenderMetaContext {
	var p = new(ZParserExtenderMetaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_zParserExtenderMeta

	return p
}

func (s *ZParserExtenderMetaContext) GetParser() antlr.Parser { return s.parser }

func (s *ZParserExtenderMetaContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ZParserExtenderMetaContext) Assign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserAssign, 0)
}

func (s *ZParserExtenderMetaContext) MetaValue() IMetaValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaValueContext)
}

func (s *ZParserExtenderMetaContext) Semi() antlr.TerminalNode {
	return s.GetToken(DaedalusParserSemi, 0)
}

func (s *ZParserExtenderMetaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ZParserExtenderMetaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ZParserExtenderMetaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterZParserExtenderMeta(s)
	}
}

func (s *ZParserExtenderMetaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitZParserExtenderMeta(s)
	}
}

func (p *DaedalusParser) ZParserExtenderMeta() (localctx IZParserExtenderMetaContext) {
	this := p
	_ = this

	localctx = NewZParserExtenderMetaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, DaedalusParserRULE_zParserExtenderMeta)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(338)
		p.NameNode()
	}
	{
		p.SetState(339)
		p.Match(DaedalusParserAssign)
	}
	{
		p.SetState(340)
		p.MetaValue()
	}
	{
		p.SetState(341)
		p.Match(DaedalusParserSemi)
	}

	return localctx
}

// IZParserExtenderMetaBlockContext is an interface to support dynamic dispatch.
type IZParserExtenderMetaBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsZParserExtenderMetaBlockContext differentiates from other interfaces.
	IsZParserExtenderMetaBlockContext()
}

type ZParserExtenderMetaBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyZParserExtenderMetaBlockContext() *ZParserExtenderMetaBlockContext {
	var p = new(ZParserExtenderMetaBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_zParserExtenderMetaBlock
	return p
}

func (*ZParserExtenderMetaBlockContext) IsZParserExtenderMetaBlockContext() {}

func NewZParserExtenderMetaBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ZParserExtenderMetaBlockContext {
	var p = new(ZParserExtenderMetaBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_zParserExtenderMetaBlock

	return p
}

func (s *ZParserExtenderMetaBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ZParserExtenderMetaBlockContext) Meta() antlr.TerminalNode {
	return s.GetToken(DaedalusParserMeta, 0)
}

func (s *ZParserExtenderMetaBlockContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBrace, 0)
}

func (s *ZParserExtenderMetaBlockContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBrace, 0)
}

func (s *ZParserExtenderMetaBlockContext) Semi() antlr.TerminalNode {
	return s.GetToken(DaedalusParserSemi, 0)
}

func (s *ZParserExtenderMetaBlockContext) AllZParserExtenderMeta() []IZParserExtenderMetaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IZParserExtenderMetaContext); ok {
			len++
		}
	}

	tst := make([]IZParserExtenderMetaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IZParserExtenderMetaContext); ok {
			tst[i] = t.(IZParserExtenderMetaContext)
			i++
		}
	}

	return tst
}

func (s *ZParserExtenderMetaBlockContext) ZParserExtenderMeta(i int) IZParserExtenderMetaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IZParserExtenderMetaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IZParserExtenderMetaContext)
}

func (s *ZParserExtenderMetaBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ZParserExtenderMetaBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ZParserExtenderMetaBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterZParserExtenderMetaBlock(s)
	}
}

func (s *ZParserExtenderMetaBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitZParserExtenderMetaBlock(s)
	}
}

func (p *DaedalusParser) ZParserExtenderMetaBlock() (localctx IZParserExtenderMetaBlockContext) {
	this := p
	_ = this

	localctx = NewZParserExtenderMetaBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, DaedalusParserRULE_zParserExtenderMetaBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(343)
		p.Match(DaedalusParserMeta)
	}
	{
		p.SetState(344)
		p.Match(DaedalusParserLeftBrace)
	}
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(345)
				p.ZParserExtenderMeta()
			}

		}
		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}
	{
		p.SetState(351)
		p.Match(DaedalusParserRightBrace)
	}
	{
		p.SetState(352)
		p.Match(DaedalusParserSemi)
	}

	return localctx
}

// IConstArrayDefContext is an interface to support dynamic dispatch.
type IConstArrayDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstArrayDefContext differentiates from other interfaces.
	IsConstArrayDefContext()
}

type ConstArrayDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstArrayDefContext() *ConstArrayDefContext {
	var p = new(ConstArrayDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_constArrayDef
	return p
}

func (*ConstArrayDefContext) IsConstArrayDefContext() {}

func NewConstArrayDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstArrayDefContext {
	var p = new(ConstArrayDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_constArrayDef

	return p
}

func (s *ConstArrayDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstArrayDefContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ConstArrayDefContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBracket, 0)
}

func (s *ConstArrayDefContext) ArraySize() IArraySizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArraySizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArraySizeContext)
}

func (s *ConstArrayDefContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBracket, 0)
}

func (s *ConstArrayDefContext) ConstArrayAssignment() IConstArrayAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstArrayAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstArrayAssignmentContext)
}

func (s *ConstArrayDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstArrayDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstArrayDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterConstArrayDef(s)
	}
}

func (s *ConstArrayDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitConstArrayDef(s)
	}
}

func (p *DaedalusParser) ConstArrayDef() (localctx IConstArrayDefContext) {
	this := p
	_ = this

	localctx = NewConstArrayDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, DaedalusParserRULE_constArrayDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(354)
		p.NameNode()
	}
	{
		p.SetState(355)
		p.Match(DaedalusParserLeftBracket)
	}
	{
		p.SetState(356)
		p.ArraySize()
	}
	{
		p.SetState(357)
		p.Match(DaedalusParserRightBracket)
	}
	{
		p.SetState(358)
		p.ConstArrayAssignment()
	}

	return localctx
}

// IConstArrayAssignmentContext is an interface to support dynamic dispatch.
type IConstArrayAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstArrayAssignmentContext differentiates from other interfaces.
	IsConstArrayAssignmentContext()
}

type ConstArrayAssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstArrayAssignmentContext() *ConstArrayAssignmentContext {
	var p = new(ConstArrayAssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_constArrayAssignment
	return p
}

func (*ConstArrayAssignmentContext) IsConstArrayAssignmentContext() {}

func NewConstArrayAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstArrayAssignmentContext {
	var p = new(ConstArrayAssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_constArrayAssignment

	return p
}

func (s *ConstArrayAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstArrayAssignmentContext) Assign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserAssign, 0)
}

func (s *ConstArrayAssignmentContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBrace, 0)
}

func (s *ConstArrayAssignmentContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBrace, 0)
}

func (s *ConstArrayAssignmentContext) AllExpressionBlock() []IExpressionBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionBlockContext); ok {
			len++
		}
	}

	tst := make([]IExpressionBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionBlockContext); ok {
			tst[i] = t.(IExpressionBlockContext)
			i++
		}
	}

	return tst
}

func (s *ConstArrayAssignmentContext) ExpressionBlock(i int) IExpressionBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionBlockContext)
}

func (s *ConstArrayAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstArrayAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstArrayAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterConstArrayAssignment(s)
	}
}

func (s *ConstArrayAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitConstArrayAssignment(s)
	}
}

func (p *DaedalusParser) ConstArrayAssignment() (localctx IConstArrayAssignmentContext) {
	this := p
	_ = this

	localctx = NewConstArrayAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, DaedalusParserRULE_constArrayAssignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		p.Match(DaedalusParserAssign)
	}
	{
		p.SetState(361)
		p.Match(DaedalusParserLeftBrace)
	}

	{
		p.SetState(362)
		p.ExpressionBlock()
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(363)
				p.Match(DaedalusParserT__0)
			}
			{
				p.SetState(364)
				p.ExpressionBlock()
			}

		}
		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}

	{
		p.SetState(370)
		p.Match(DaedalusParserRightBrace)
	}

	return localctx
}

// IConstValueDefContext is an interface to support dynamic dispatch.
type IConstValueDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstValueDefContext differentiates from other interfaces.
	IsConstValueDefContext()
}

type ConstValueDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstValueDefContext() *ConstValueDefContext {
	var p = new(ConstValueDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_constValueDef
	return p
}

func (*ConstValueDefContext) IsConstValueDefContext() {}

func NewConstValueDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstValueDefContext {
	var p = new(ConstValueDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_constValueDef

	return p
}

func (s *ConstValueDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstValueDefContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ConstValueDefContext) ConstValueAssignment() IConstValueAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstValueAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstValueAssignmentContext)
}

func (s *ConstValueDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstValueDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstValueDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterConstValueDef(s)
	}
}

func (s *ConstValueDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitConstValueDef(s)
	}
}

func (p *DaedalusParser) ConstValueDef() (localctx IConstValueDefContext) {
	this := p
	_ = this

	localctx = NewConstValueDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, DaedalusParserRULE_constValueDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.NameNode()
	}
	{
		p.SetState(373)
		p.ConstValueAssignment()
	}

	return localctx
}

// IConstValueAssignmentContext is an interface to support dynamic dispatch.
type IConstValueAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstValueAssignmentContext differentiates from other interfaces.
	IsConstValueAssignmentContext()
}

type ConstValueAssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstValueAssignmentContext() *ConstValueAssignmentContext {
	var p = new(ConstValueAssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_constValueAssignment
	return p
}

func (*ConstValueAssignmentContext) IsConstValueAssignmentContext() {}

func NewConstValueAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstValueAssignmentContext {
	var p = new(ConstValueAssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_constValueAssignment

	return p
}

func (s *ConstValueAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstValueAssignmentContext) Assign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserAssign, 0)
}

func (s *ConstValueAssignmentContext) ExpressionBlock() IExpressionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionBlockContext)
}

func (s *ConstValueAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstValueAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstValueAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterConstValueAssignment(s)
	}
}

func (s *ConstValueAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitConstValueAssignment(s)
	}
}

func (p *DaedalusParser) ConstValueAssignment() (localctx IConstValueAssignmentContext) {
	this := p
	_ = this

	localctx = NewConstValueAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, DaedalusParserRULE_constValueAssignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(375)
		p.Match(DaedalusParserAssign)
	}
	{
		p.SetState(376)
		p.ExpressionBlock()
	}

	return localctx
}

// IVarArrayDeclContext is an interface to support dynamic dispatch.
type IVarArrayDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarArrayDeclContext differentiates from other interfaces.
	IsVarArrayDeclContext()
}

type VarArrayDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarArrayDeclContext() *VarArrayDeclContext {
	var p = new(VarArrayDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_varArrayDecl
	return p
}

func (*VarArrayDeclContext) IsVarArrayDeclContext() {}

func NewVarArrayDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarArrayDeclContext {
	var p = new(VarArrayDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_varArrayDecl

	return p
}

func (s *VarArrayDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarArrayDeclContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *VarArrayDeclContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBracket, 0)
}

func (s *VarArrayDeclContext) ArraySize() IArraySizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArraySizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArraySizeContext)
}

func (s *VarArrayDeclContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBracket, 0)
}

func (s *VarArrayDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarArrayDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarArrayDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterVarArrayDecl(s)
	}
}

func (s *VarArrayDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitVarArrayDecl(s)
	}
}

func (p *DaedalusParser) VarArrayDecl() (localctx IVarArrayDeclContext) {
	this := p
	_ = this

	localctx = NewVarArrayDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, DaedalusParserRULE_varArrayDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.NameNode()
	}
	{
		p.SetState(379)
		p.Match(DaedalusParserLeftBracket)
	}
	{
		p.SetState(380)
		p.ArraySize()
	}
	{
		p.SetState(381)
		p.Match(DaedalusParserRightBracket)
	}

	return localctx
}

// IVarValueDeclContext is an interface to support dynamic dispatch.
type IVarValueDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarValueDeclContext differentiates from other interfaces.
	IsVarValueDeclContext()
}

type VarValueDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarValueDeclContext() *VarValueDeclContext {
	var p = new(VarValueDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_varValueDecl
	return p
}

func (*VarValueDeclContext) IsVarValueDeclContext() {}

func NewVarValueDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarValueDeclContext {
	var p = new(VarValueDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_varValueDecl

	return p
}

func (s *VarValueDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarValueDeclContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *VarValueDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarValueDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarValueDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterVarValueDecl(s)
	}
}

func (s *VarValueDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitVarValueDecl(s)
	}
}

func (p *DaedalusParser) VarValueDecl() (localctx IVarValueDeclContext) {
	this := p
	_ = this

	localctx = NewVarValueDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, DaedalusParserRULE_varValueDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.NameNode()
	}

	return localctx
}

// IVariadicContext is an interface to support dynamic dispatch.
type IVariadicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariadicContext differentiates from other interfaces.
	IsVariadicContext()
}

type VariadicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariadicContext() *VariadicContext {
	var p = new(VariadicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_variadic
	return p
}

func (*VariadicContext) IsVariadicContext() {}

func NewVariadicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariadicContext {
	var p = new(VariadicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_variadic

	return p
}

func (s *VariadicContext) GetParser() antlr.Parser { return s.parser }
func (s *VariadicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariadicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariadicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterVariadic(s)
	}
}

func (s *VariadicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitVariadic(s)
	}
}

func (p *DaedalusParser) Variadic() (localctx IVariadicContext) {
	this := p
	_ = this

	localctx = NewVariadicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, DaedalusParserRULE_variadic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Match(DaedalusParserT__1)
	}

	return localctx
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_parameterList
	return p
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftParen, 0)
}

func (s *ParameterListContext) RightParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightParen, 0)
}

func (s *ParameterListContext) AllParameterDecl() []IParameterDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterDeclContext); ok {
			len++
		}
	}

	tst := make([]IParameterDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterDeclContext); ok {
			tst[i] = t.(IParameterDeclContext)
			i++
		}
	}

	return tst
}

func (s *ParameterListContext) ParameterDecl(i int) IParameterDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterDeclContext)
}

func (s *ParameterListContext) Variadic() IVariadicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariadicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariadicContext)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (p *DaedalusParser) ParameterList() (localctx IParameterListContext) {
	this := p
	_ = this

	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, DaedalusParserRULE_parameterList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
		p.Match(DaedalusParserLeftParen)
	}
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserVar {
		{
			p.SetState(388)
			p.ParameterDecl()
		}
		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

		for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1+1 {
				{
					p.SetState(389)
					p.Match(DaedalusParserT__0)
				}
				{
					p.SetState(390)
					p.ParameterDecl()
				}

			}
			p.SetState(395)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
		}

	}
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserT__0 {
		{
			p.SetState(398)
			p.Match(DaedalusParserT__0)
		}
		{
			p.SetState(399)
			p.Variadic()
		}

	}
	{
		p.SetState(402)
		p.Match(DaedalusParserRightParen)
	}

	return localctx
}

// IParameterDeclContext is an interface to support dynamic dispatch.
type IParameterDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterDeclContext differentiates from other interfaces.
	IsParameterDeclContext()
}

type ParameterDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterDeclContext() *ParameterDeclContext {
	var p = new(ParameterDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_parameterDecl
	return p
}

func (*ParameterDeclContext) IsParameterDeclContext() {}

func NewParameterDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterDeclContext {
	var p = new(ParameterDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_parameterDecl

	return p
}

func (s *ParameterDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterDeclContext) Var() antlr.TerminalNode {
	return s.GetToken(DaedalusParserVar, 0)
}

func (s *ParameterDeclContext) TypeReference() ITypeReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *ParameterDeclContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ParameterDeclContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBracket, 0)
}

func (s *ParameterDeclContext) ArraySize() IArraySizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArraySizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArraySizeContext)
}

func (s *ParameterDeclContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBracket, 0)
}

func (s *ParameterDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterParameterDecl(s)
	}
}

func (s *ParameterDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitParameterDecl(s)
	}
}

func (p *DaedalusParser) ParameterDecl() (localctx IParameterDeclContext) {
	this := p
	_ = this

	localctx = NewParameterDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, DaedalusParserRULE_parameterDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(404)
		p.Match(DaedalusParserVar)
	}
	{
		p.SetState(405)
		p.TypeReference()
	}
	{
		p.SetState(406)
		p.NameNode()
	}
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserLeftBracket {
		{
			p.SetState(407)
			p.Match(DaedalusParserLeftBracket)
		}
		{
			p.SetState(408)
			p.ArraySize()
		}
		{
			p.SetState(409)
			p.Match(DaedalusParserRightBracket)
		}

	}

	return localctx
}

// IStatementBlockContext is an interface to support dynamic dispatch.
type IStatementBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementBlockContext differentiates from other interfaces.
	IsStatementBlockContext()
}

type StatementBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementBlockContext() *StatementBlockContext {
	var p = new(StatementBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_statementBlock
	return p
}

func (*StatementBlockContext) IsStatementBlockContext() {}

func NewStatementBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementBlockContext {
	var p = new(StatementBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_statementBlock

	return p
}

func (s *StatementBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementBlockContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBrace, 0)
}

func (s *StatementBlockContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBrace, 0)
}

func (s *StatementBlockContext) AllMacroDef() []IMacroDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMacroDefContext); ok {
			len++
		}
	}

	tst := make([]IMacroDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMacroDefContext); ok {
			tst[i] = t.(IMacroDefContext)
			i++
		}
	}

	return tst
}

func (s *StatementBlockContext) MacroDef(i int) IMacroDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacroDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacroDefContext)
}

func (s *StatementBlockContext) AllWhileBlockStatement() []IWhileBlockStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWhileBlockStatementContext); ok {
			len++
		}
	}

	tst := make([]IWhileBlockStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWhileBlockStatementContext); ok {
			tst[i] = t.(IWhileBlockStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementBlockContext) WhileBlockStatement(i int) IWhileBlockStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileBlockStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileBlockStatementContext)
}

func (s *StatementBlockContext) AllSemi() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserSemi)
}

func (s *StatementBlockContext) Semi(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserSemi, i)
}

func (s *StatementBlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementBlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementBlockContext) AllIfBlockStatement() []IIfBlockStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfBlockStatementContext); ok {
			len++
		}
	}

	tst := make([]IIfBlockStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfBlockStatementContext); ok {
			tst[i] = t.(IIfBlockStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementBlockContext) IfBlockStatement(i int) IIfBlockStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockStatementContext)
}

func (s *StatementBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterStatementBlock(s)
	}
}

func (s *StatementBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitStatementBlock(s)
	}
}

func (p *DaedalusParser) StatementBlock() (localctx IStatementBlockContext) {
	this := p
	_ = this

	localctx = NewStatementBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, DaedalusParserRULE_statementBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(413)
		p.Match(DaedalusParserLeftBrace)
	}
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(425)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case DaedalusParserIntegerLiteral, DaedalusParserFloatLiteral, DaedalusParserStringLiteral, DaedalusParserVar, DaedalusParserConst, DaedalusParserClass, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNull, DaedalusParserReturn, DaedalusParserNamespace, DaedalusParserFunc, DaedalusParserInt, DaedalusParserFloat, DaedalusParserStringKeyword, DaedalusParserVoid, DaedalusParserMeta, DaedalusParserContinue, DaedalusParserBreak, DaedalusParserLeftParen, DaedalusParserPlus, DaedalusParserMinus, DaedalusParserTilde, DaedalusParserNot, DaedalusParserIdentifier:
				{
					p.SetState(414)
					p.Statement()
				}
				{
					p.SetState(415)
					p.Match(DaedalusParserSemi)
				}

			case DaedalusParserIf:
				{
					p.SetState(417)
					p.IfBlockStatement()
				}
				p.SetState(419)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == DaedalusParserSemi {
					{
						p.SetState(418)
						p.Match(DaedalusParserSemi)
					}

				}

			case DaedalusParserMif:
				{
					p.SetState(421)
					p.MacroDef()
				}

			case DaedalusParserWhile:
				{
					p.SetState(422)
					p.WhileBlockStatement()
				}
				{
					p.SetState(423)
					p.Match(DaedalusParserSemi)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(429)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}
	{
		p.SetState(430)
		p.Match(DaedalusParserRightBrace)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) ConstDef() IConstDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstDefContext)
}

func (s *StatementContext) VarDecl() IVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) Continue() antlr.TerminalNode {
	return s.GetToken(DaedalusParserContinue, 0)
}

func (s *StatementContext) Break() antlr.TerminalNode {
	return s.GetToken(DaedalusParserBreak, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *DaedalusParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, DaedalusParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(432)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(433)
			p.ReturnStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(434)
			p.ConstDef()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(435)
			p.VarDecl()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(436)
			p.expression(0)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(437)
			p.Match(DaedalusParserContinue)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(438)
			p.Match(DaedalusParserBreak)
		}

	}

	return localctx
}

// IFuncCallContext is an interface to support dynamic dispatch.
type IFuncCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncCallContext differentiates from other interfaces.
	IsFuncCallContext()
}

type FuncCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallContext() *FuncCallContext {
	var p = new(FuncCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_funcCall
	return p
}

func (*FuncCallContext) IsFuncCallContext() {}

func NewFuncCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallContext {
	var p = new(FuncCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_funcCall

	return p
}

func (s *FuncCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *FuncCallContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftParen, 0)
}

func (s *FuncCallContext) RightParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightParen, 0)
}

func (s *FuncCallContext) AllFuncArgExpression() []IFuncArgExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncArgExpressionContext); ok {
			len++
		}
	}

	tst := make([]IFuncArgExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncArgExpressionContext); ok {
			tst[i] = t.(IFuncArgExpressionContext)
			i++
		}
	}

	return tst
}

func (s *FuncCallContext) FuncArgExpression(i int) IFuncArgExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncArgExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncArgExpressionContext)
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

func (p *DaedalusParser) FuncCall() (localctx IFuncCallContext) {
	this := p
	_ = this

	localctx = NewFuncCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, DaedalusParserRULE_funcCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(441)
		p.NameNode()
	}
	{
		p.SetState(442)
		p.Match(DaedalusParserLeftParen)
	}
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-10)) & ^0x3f) == 0 && ((int64(1)<<(_la-10))&36042816061437167) != 0 {
		{
			p.SetState(443)
			p.FuncArgExpression()
		}
		p.SetState(448)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

		for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1+1 {
				{
					p.SetState(444)
					p.Match(DaedalusParserT__0)
				}
				{
					p.SetState(445)
					p.FuncArgExpression()
				}

			}
			p.SetState(450)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
		}

	}
	{
		p.SetState(453)
		p.Match(DaedalusParserRightParen)
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *AssignmentContext) AssignmentOperator() IAssignmentOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentOperatorContext)
}

func (s *AssignmentContext) ExpressionBlock() IExpressionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionBlockContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *DaedalusParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, DaedalusParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(455)
		p.Reference()
	}
	{
		p.SetState(456)
		p.AssignmentOperator()
	}
	{
		p.SetState(457)
		p.ExpressionBlock()
	}

	return localctx
}

// IIfConditionContext is an interface to support dynamic dispatch.
type IIfConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfConditionContext differentiates from other interfaces.
	IsIfConditionContext()
}

type IfConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfConditionContext() *IfConditionContext {
	var p = new(IfConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_ifCondition
	return p
}

func (*IfConditionContext) IsIfConditionContext() {}

func NewIfConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfConditionContext {
	var p = new(IfConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_ifCondition

	return p
}

func (s *IfConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *IfConditionContext) ExpressionBlock() IExpressionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionBlockContext)
}

func (s *IfConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterIfCondition(s)
	}
}

func (s *IfConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitIfCondition(s)
	}
}

func (p *DaedalusParser) IfCondition() (localctx IIfConditionContext) {
	this := p
	_ = this

	localctx = NewIfConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, DaedalusParserRULE_ifCondition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(459)
		p.ExpressionBlock()
	}

	return localctx
}

// IElseBlockContext is an interface to support dynamic dispatch.
type IElseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseBlockContext differentiates from other interfaces.
	IsElseBlockContext()
}

type ElseBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBlockContext() *ElseBlockContext {
	var p = new(ElseBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_elseBlock
	return p
}

func (*ElseBlockContext) IsElseBlockContext() {}

func NewElseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBlockContext {
	var p = new(ElseBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_elseBlock

	return p
}

func (s *ElseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBlockContext) Else() antlr.TerminalNode {
	return s.GetToken(DaedalusParserElse, 0)
}

func (s *ElseBlockContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *ElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterElseBlock(s)
	}
}

func (s *ElseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitElseBlock(s)
	}
}

func (p *DaedalusParser) ElseBlock() (localctx IElseBlockContext) {
	this := p
	_ = this

	localctx = NewElseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, DaedalusParserRULE_elseBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(461)
		p.Match(DaedalusParserElse)
	}
	{
		p.SetState(462)
		p.StatementBlock()
	}

	return localctx
}

// IElseIfBlockContext is an interface to support dynamic dispatch.
type IElseIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseIfBlockContext differentiates from other interfaces.
	IsElseIfBlockContext()
}

type ElseIfBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfBlockContext() *ElseIfBlockContext {
	var p = new(ElseIfBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_elseIfBlock
	return p
}

func (*ElseIfBlockContext) IsElseIfBlockContext() {}

func NewElseIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfBlockContext {
	var p = new(ElseIfBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_elseIfBlock

	return p
}

func (s *ElseIfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfBlockContext) Else() antlr.TerminalNode {
	return s.GetToken(DaedalusParserElse, 0)
}

func (s *ElseIfBlockContext) If() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIf, 0)
}

func (s *ElseIfBlockContext) IfCondition() IIfConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfConditionContext)
}

func (s *ElseIfBlockContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *ElseIfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterElseIfBlock(s)
	}
}

func (s *ElseIfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitElseIfBlock(s)
	}
}

func (p *DaedalusParser) ElseIfBlock() (localctx IElseIfBlockContext) {
	this := p
	_ = this

	localctx = NewElseIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, DaedalusParserRULE_elseIfBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(464)
		p.Match(DaedalusParserElse)
	}
	{
		p.SetState(465)
		p.Match(DaedalusParserIf)
	}
	{
		p.SetState(466)
		p.IfCondition()
	}
	{
		p.SetState(467)
		p.StatementBlock()
	}

	return localctx
}

// IIfBlockContext is an interface to support dynamic dispatch.
type IIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfBlockContext differentiates from other interfaces.
	IsIfBlockContext()
}

type IfBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBlockContext() *IfBlockContext {
	var p = new(IfBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_ifBlock
	return p
}

func (*IfBlockContext) IsIfBlockContext() {}

func NewIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBlockContext {
	var p = new(IfBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_ifBlock

	return p
}

func (s *IfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBlockContext) If() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIf, 0)
}

func (s *IfBlockContext) IfCondition() IIfConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfConditionContext)
}

func (s *IfBlockContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *IfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterIfBlock(s)
	}
}

func (s *IfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitIfBlock(s)
	}
}

func (p *DaedalusParser) IfBlock() (localctx IIfBlockContext) {
	this := p
	_ = this

	localctx = NewIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, DaedalusParserRULE_ifBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(469)
		p.Match(DaedalusParserIf)
	}
	{
		p.SetState(470)
		p.IfCondition()
	}
	{
		p.SetState(471)
		p.StatementBlock()
	}

	return localctx
}

// IIfBlockStatementContext is an interface to support dynamic dispatch.
type IIfBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfBlockStatementContext differentiates from other interfaces.
	IsIfBlockStatementContext()
}

type IfBlockStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBlockStatementContext() *IfBlockStatementContext {
	var p = new(IfBlockStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_ifBlockStatement
	return p
}

func (*IfBlockStatementContext) IsIfBlockStatementContext() {}

func NewIfBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBlockStatementContext {
	var p = new(IfBlockStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_ifBlockStatement

	return p
}

func (s *IfBlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBlockStatementContext) IfBlock() IIfBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *IfBlockStatementContext) AllElseIfBlock() []IElseIfBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseIfBlockContext); ok {
			len++
		}
	}

	tst := make([]IElseIfBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseIfBlockContext); ok {
			tst[i] = t.(IElseIfBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfBlockStatementContext) ElseIfBlock(i int) IElseIfBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseIfBlockContext)
}

func (s *IfBlockStatementContext) ElseBlock() IElseBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseBlockContext)
}

func (s *IfBlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfBlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterIfBlockStatement(s)
	}
}

func (s *IfBlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitIfBlockStatement(s)
	}
}

func (p *DaedalusParser) IfBlockStatement() (localctx IIfBlockStatementContext) {
	this := p
	_ = this

	localctx = NewIfBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, DaedalusParserRULE_ifBlockStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(473)
		p.IfBlock()
	}
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(474)
				p.ElseIfBlock()
			}

		}
		p.SetState(479)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
	}
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserElse {
		{
			p.SetState(480)
			p.ElseBlock()
		}

	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Return() antlr.TerminalNode {
	return s.GetToken(DaedalusParserReturn, 0)
}

func (s *ReturnStatementContext) ExpressionBlock() IExpressionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionBlockContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (p *DaedalusParser) ReturnStatement() (localctx IReturnStatementContext) {
	this := p
	_ = this

	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, DaedalusParserRULE_returnStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)
		p.Match(DaedalusParserReturn)
	}
	p.SetState(485)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-10)) & ^0x3f) == 0 && ((int64(1)<<(_la-10))&36042816061437167) != 0 {
		{
			p.SetState(484)
			p.ExpressionBlock()
		}

	}

	return localctx
}

// IWhileConditionContext is an interface to support dynamic dispatch.
type IWhileConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileConditionContext differentiates from other interfaces.
	IsWhileConditionContext()
}

type WhileConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileConditionContext() *WhileConditionContext {
	var p = new(WhileConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_whileCondition
	return p
}

func (*WhileConditionContext) IsWhileConditionContext() {}

func NewWhileConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileConditionContext {
	var p = new(WhileConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_whileCondition

	return p
}

func (s *WhileConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileConditionContext) ExpressionBlock() IExpressionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionBlockContext)
}

func (s *WhileConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterWhileCondition(s)
	}
}

func (s *WhileConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitWhileCondition(s)
	}
}

func (p *DaedalusParser) WhileCondition() (localctx IWhileConditionContext) {
	this := p
	_ = this

	localctx = NewWhileConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, DaedalusParserRULE_whileCondition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(487)
		p.ExpressionBlock()
	}

	return localctx
}

// IWhileBlockContext is an interface to support dynamic dispatch.
type IWhileBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileBlockContext differentiates from other interfaces.
	IsWhileBlockContext()
}

type WhileBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileBlockContext() *WhileBlockContext {
	var p = new(WhileBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_whileBlock
	return p
}

func (*WhileBlockContext) IsWhileBlockContext() {}

func NewWhileBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileBlockContext {
	var p = new(WhileBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_whileBlock

	return p
}

func (s *WhileBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileBlockContext) While() antlr.TerminalNode {
	return s.GetToken(DaedalusParserWhile, 0)
}

func (s *WhileBlockContext) WhileCondition() IWhileConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileConditionContext)
}

func (s *WhileBlockContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *WhileBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterWhileBlock(s)
	}
}

func (s *WhileBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitWhileBlock(s)
	}
}

func (p *DaedalusParser) WhileBlock() (localctx IWhileBlockContext) {
	this := p
	_ = this

	localctx = NewWhileBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, DaedalusParserRULE_whileBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(489)
		p.Match(DaedalusParserWhile)
	}
	{
		p.SetState(490)
		p.WhileCondition()
	}
	{
		p.SetState(491)
		p.StatementBlock()
	}

	return localctx
}

// IWhileBlockStatementContext is an interface to support dynamic dispatch.
type IWhileBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileBlockStatementContext differentiates from other interfaces.
	IsWhileBlockStatementContext()
}

type WhileBlockStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileBlockStatementContext() *WhileBlockStatementContext {
	var p = new(WhileBlockStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_whileBlockStatement
	return p
}

func (*WhileBlockStatementContext) IsWhileBlockStatementContext() {}

func NewWhileBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileBlockStatementContext {
	var p = new(WhileBlockStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_whileBlockStatement

	return p
}

func (s *WhileBlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileBlockStatementContext) WhileBlock() IWhileBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileBlockContext)
}

func (s *WhileBlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileBlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileBlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterWhileBlockStatement(s)
	}
}

func (s *WhileBlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitWhileBlockStatement(s)
	}
}

func (p *DaedalusParser) WhileBlockStatement() (localctx IWhileBlockStatementContext) {
	this := p
	_ = this

	localctx = NewWhileBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, DaedalusParserRULE_whileBlockStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(493)
		p.WhileBlock()
	}

	return localctx
}

// IFuncArgExpressionContext is an interface to support dynamic dispatch.
type IFuncArgExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncArgExpressionContext differentiates from other interfaces.
	IsFuncArgExpressionContext()
}

type FuncArgExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgExpressionContext() *FuncArgExpressionContext {
	var p = new(FuncArgExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_funcArgExpression
	return p
}

func (*FuncArgExpressionContext) IsFuncArgExpressionContext() {}

func NewFuncArgExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgExpressionContext {
	var p = new(FuncArgExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_funcArgExpression

	return p
}

func (s *FuncArgExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgExpressionContext) ExpressionBlock() IExpressionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionBlockContext)
}

func (s *FuncArgExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncArgExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterFuncArgExpression(s)
	}
}

func (s *FuncArgExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitFuncArgExpression(s)
	}
}

func (p *DaedalusParser) FuncArgExpression() (localctx IFuncArgExpressionContext) {
	this := p
	_ = this

	localctx = NewFuncArgExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, DaedalusParserRULE_funcArgExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(495)
		p.ExpressionBlock()
	}

	return localctx
}

// IExpressionBlockContext is an interface to support dynamic dispatch.
type IExpressionBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionBlockContext differentiates from other interfaces.
	IsExpressionBlockContext()
}

type ExpressionBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionBlockContext() *ExpressionBlockContext {
	var p = new(ExpressionBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_expressionBlock
	return p
}

func (*ExpressionBlockContext) IsExpressionBlockContext() {}

func NewExpressionBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionBlockContext {
	var p = new(ExpressionBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_expressionBlock

	return p
}

func (s *ExpressionBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionBlockContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterExpressionBlock(s)
	}
}

func (s *ExpressionBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitExpressionBlock(s)
	}
}

func (p *DaedalusParser) ExpressionBlock() (localctx IExpressionBlockContext) {
	this := p
	_ = this

	localctx = NewExpressionBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, DaedalusParserRULE_expressionBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(497)
		p.expression(0)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BitMoveExpressionContext struct {
	*ExpressionContext
}

func NewBitMoveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitMoveExpressionContext {
	var p = new(BitMoveExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BitMoveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitMoveExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitMoveExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitMoveExpressionContext) BitMoveOperator() IBitMoveOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBitMoveOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBitMoveOperatorContext)
}

func (s *BitMoveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterBitMoveExpression(s)
	}
}

func (s *BitMoveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitBitMoveExpression(s)
	}
}

type EqExpressionContext struct {
	*ExpressionContext
}

func NewEqExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqExpressionContext {
	var p = new(EqExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EqExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqExpressionContext) EqOperator() IEqOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqOperatorContext)
}

func (s *EqExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterEqExpression(s)
	}
}

func (s *EqExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitEqExpression(s)
	}
}

type ValExpressionContext struct {
	*ExpressionContext
}

func NewValExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValExpressionContext {
	var p = new(ValExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ValExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValExpressionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterValExpression(s)
	}
}

func (s *ValExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitValExpression(s)
	}
}

type AddExpressionContext struct {
	*ExpressionContext
}

func NewAddExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExpressionContext {
	var p = new(AddExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddExpressionContext) AddOperator() IAddOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddOperatorContext)
}

func (s *AddExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterAddExpression(s)
	}
}

func (s *AddExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitAddExpression(s)
	}
}

type CompExpressionContext struct {
	*ExpressionContext
}

func NewCompExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompExpressionContext {
	var p = new(CompExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CompExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CompExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompExpressionContext) CompOperator() ICompOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompOperatorContext)
}

func (s *CompExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterCompExpression(s)
	}
}

func (s *CompExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitCompExpression(s)
	}
}

type LogOrExpressionContext struct {
	*ExpressionContext
}

func NewLogOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogOrExpressionContext {
	var p = new(LogOrExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LogOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogOrExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogOrExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogOrExpressionContext) LogOrOperator() ILogOrOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogOrOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogOrOperatorContext)
}

func (s *LogOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterLogOrExpression(s)
	}
}

func (s *LogOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitLogOrExpression(s)
	}
}

type BinAndExpressionContext struct {
	*ExpressionContext
}

func NewBinAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinAndExpressionContext {
	var p = new(BinAndExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinAndExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinAndExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinAndExpressionContext) BinAndOperator() IBinAndOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinAndOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinAndOperatorContext)
}

func (s *BinAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterBinAndExpression(s)
	}
}

func (s *BinAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitBinAndExpression(s)
	}
}

type BinOrExpressionContext struct {
	*ExpressionContext
}

func NewBinOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinOrExpressionContext {
	var p = new(BinOrExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinOrExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinOrExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinOrExpressionContext) BinOrOperator() IBinOrOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinOrOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinOrOperatorContext)
}

func (s *BinOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterBinOrExpression(s)
	}
}

func (s *BinOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitBinOrExpression(s)
	}
}

type MultExpressionContext struct {
	*ExpressionContext
}

func NewMultExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultExpressionContext {
	var p = new(MultExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MultExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MultExpressionContext) MultOperator() IMultOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultOperatorContext)
}

func (s *MultExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterMultExpression(s)
	}
}

func (s *MultExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitMultExpression(s)
	}
}

type BracketExpressionContext struct {
	*ExpressionContext
}

func NewBracketExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketExpressionContext {
	var p = new(BracketExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BracketExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketExpressionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftParen, 0)
}

func (s *BracketExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BracketExpressionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightParen, 0)
}

func (s *BracketExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterBracketExpression(s)
	}
}

func (s *BracketExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitBracketExpression(s)
	}
}

type UnaryOperationContext struct {
	*ExpressionContext
}

func NewUnaryOperationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryOperationContext {
	var p = new(UnaryOperationContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperationContext) UnaryOperator() IUnaryOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryOperatorContext)
}

func (s *UnaryOperationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterUnaryOperation(s)
	}
}

func (s *UnaryOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitUnaryOperation(s)
	}
}

type LogAndExpressionContext struct {
	*ExpressionContext
}

func NewLogAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogAndExpressionContext {
	var p = new(LogAndExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LogAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogAndExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogAndExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogAndExpressionContext) LogAndOperator() ILogAndOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogAndOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogAndOperatorContext)
}

func (s *LogAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterLogAndExpression(s)
	}
}

func (s *LogAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitLogAndExpression(s)
	}
}

func (p *DaedalusParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *DaedalusParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 98
	p.EnterRecursionRule(localctx, 98, DaedalusParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(508)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserLeftParen:
		localctx = NewBracketExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(500)
			p.Match(DaedalusParserLeftParen)
		}
		{
			p.SetState(501)
			p.expression(0)
		}
		{
			p.SetState(502)
			p.Match(DaedalusParserRightParen)
		}

	case DaedalusParserPlus, DaedalusParserMinus, DaedalusParserTilde, DaedalusParserNot:
		localctx = NewUnaryOperationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(504)
			p.UnaryOperator()
		}
		{
			p.SetState(505)
			p.expression(11)
		}

	case DaedalusParserIntegerLiteral, DaedalusParserFloatLiteral, DaedalusParserStringLiteral, DaedalusParserVar, DaedalusParserClass, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNull, DaedalusParserNamespace, DaedalusParserFunc, DaedalusParserInt, DaedalusParserFloat, DaedalusParserStringKeyword, DaedalusParserVoid, DaedalusParserMeta, DaedalusParserIdentifier:
		localctx = NewValExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(507)
			p.Value()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(548)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(546)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(510)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(511)
					p.MultOperator()
				}
				{
					p.SetState(512)
					p.expression(11)
				}

			case 2:
				localctx = NewAddExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(514)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(515)
					p.AddOperator()
				}
				{
					p.SetState(516)
					p.expression(10)
				}

			case 3:
				localctx = NewBitMoveExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(518)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(519)
					p.BitMoveOperator()
				}
				{
					p.SetState(520)
					p.expression(9)
				}

			case 4:
				localctx = NewCompExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(522)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(523)
					p.CompOperator()
				}
				{
					p.SetState(524)
					p.expression(8)
				}

			case 5:
				localctx = NewEqExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(526)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(527)
					p.EqOperator()
				}
				{
					p.SetState(528)
					p.expression(7)
				}

			case 6:
				localctx = NewBinAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(530)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(531)
					p.BinAndOperator()
				}
				{
					p.SetState(532)
					p.expression(6)
				}

			case 7:
				localctx = NewBinOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(534)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(535)
					p.BinOrOperator()
				}
				{
					p.SetState(536)
					p.expression(5)
				}

			case 8:
				localctx = NewLogAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(538)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(539)
					p.LogAndOperator()
				}
				{
					p.SetState(540)
					p.expression(4)
				}

			case 9:
				localctx = NewLogOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(542)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(543)
					p.LogOrOperator()
				}
				{
					p.SetState(544)
					p.expression(3)
				}

			}

		}
		p.SetState(550)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
	}

	return localctx
}

// IArrayIndexContext is an interface to support dynamic dispatch.
type IArrayIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayIndexContext differentiates from other interfaces.
	IsArrayIndexContext()
}

type ArrayIndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayIndexContext() *ArrayIndexContext {
	var p = new(ArrayIndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_arrayIndex
	return p
}

func (*ArrayIndexContext) IsArrayIndexContext() {}

func NewArrayIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayIndexContext {
	var p = new(ArrayIndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_arrayIndex

	return p
}

func (s *ArrayIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayIndexContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIntegerLiteral, 0)
}

func (s *ArrayIndexContext) ReferenceAtom() IReferenceAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceAtomContext)
}

func (s *ArrayIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterArrayIndex(s)
	}
}

func (s *ArrayIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitArrayIndex(s)
	}
}

func (p *DaedalusParser) ArrayIndex() (localctx IArrayIndexContext) {
	this := p
	_ = this

	localctx = NewArrayIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, DaedalusParserRULE_arrayIndex)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(553)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(551)
			p.Match(DaedalusParserIntegerLiteral)
		}

	case DaedalusParserVar, DaedalusParserClass, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNull, DaedalusParserNamespace, DaedalusParserFunc, DaedalusParserInt, DaedalusParserFloat, DaedalusParserStringKeyword, DaedalusParserVoid, DaedalusParserMeta, DaedalusParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(552)
			p.ReferenceAtom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArraySizeContext is an interface to support dynamic dispatch.
type IArraySizeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArraySizeContext differentiates from other interfaces.
	IsArraySizeContext()
}

type ArraySizeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArraySizeContext() *ArraySizeContext {
	var p = new(ArraySizeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_arraySize
	return p
}

func (*ArraySizeContext) IsArraySizeContext() {}

func NewArraySizeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArraySizeContext {
	var p = new(ArraySizeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_arraySize

	return p
}

func (s *ArraySizeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArraySizeContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIntegerLiteral, 0)
}

func (s *ArraySizeContext) ReferenceAtom() IReferenceAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceAtomContext)
}

func (s *ArraySizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArraySizeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArraySizeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterArraySize(s)
	}
}

func (s *ArraySizeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitArraySize(s)
	}
}

func (p *DaedalusParser) ArraySize() (localctx IArraySizeContext) {
	this := p
	_ = this

	localctx = NewArraySizeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, DaedalusParserRULE_arraySize)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(557)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(555)
			p.Match(DaedalusParserIntegerLiteral)
		}

	case DaedalusParserVar, DaedalusParserClass, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNull, DaedalusParserNamespace, DaedalusParserFunc, DaedalusParserInt, DaedalusParserFloat, DaedalusParserStringKeyword, DaedalusParserVoid, DaedalusParserMeta, DaedalusParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(556)
			p.ReferenceAtom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IntegerLiteralValueContext struct {
	*ValueContext
}

func NewIntegerLiteralValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerLiteralValueContext {
	var p = new(IntegerLiteralValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *IntegerLiteralValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralValueContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIntegerLiteral, 0)
}

func (s *IntegerLiteralValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterIntegerLiteralValue(s)
	}
}

func (s *IntegerLiteralValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitIntegerLiteralValue(s)
	}
}

type FloatLiteralValueContext struct {
	*ValueContext
}

func NewFloatLiteralValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLiteralValueContext {
	var p = new(FloatLiteralValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *FloatLiteralValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralValueContext) FloatLiteral() antlr.TerminalNode {
	return s.GetToken(DaedalusParserFloatLiteral, 0)
}

func (s *FloatLiteralValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterFloatLiteralValue(s)
	}
}

func (s *FloatLiteralValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitFloatLiteralValue(s)
	}
}

type StringLiteralValueContext struct {
	*ValueContext
}

func NewStringLiteralValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralValueContext {
	var p = new(StringLiteralValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *StringLiteralValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralValueContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(DaedalusParserStringLiteral, 0)
}

func (s *StringLiteralValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterStringLiteralValue(s)
	}
}

func (s *StringLiteralValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitStringLiteralValue(s)
	}
}

type NullLiteralValueContext struct {
	*ValueContext
}

func NewNullLiteralValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullLiteralValueContext {
	var p = new(NullLiteralValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NullLiteralValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullLiteralValueContext) Null() antlr.TerminalNode {
	return s.GetToken(DaedalusParserNull, 0)
}

func (s *NullLiteralValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterNullLiteralValue(s)
	}
}

func (s *NullLiteralValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitNullLiteralValue(s)
	}
}

type FuncCallValueContext struct {
	*ValueContext
}

func NewFuncCallValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallValueContext {
	var p = new(FuncCallValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *FuncCallValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallValueContext) FuncCall() IFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *FuncCallValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterFuncCallValue(s)
	}
}

func (s *FuncCallValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitFuncCallValue(s)
	}
}

type ReferenceValueContext struct {
	*ValueContext
}

func NewReferenceValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReferenceValueContext {
	var p = new(ReferenceValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ReferenceValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceValueContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *ReferenceValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterReferenceValue(s)
	}
}

func (s *ReferenceValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitReferenceValue(s)
	}
}

func (p *DaedalusParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, DaedalusParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(565)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIntegerLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(559)
			p.Match(DaedalusParserIntegerLiteral)
		}

	case 2:
		localctx = NewFloatLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(560)
			p.Match(DaedalusParserFloatLiteral)
		}

	case 3:
		localctx = NewStringLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(561)
			p.Match(DaedalusParserStringLiteral)
		}

	case 4:
		localctx = NewNullLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(562)
			p.Match(DaedalusParserNull)
		}

	case 5:
		localctx = NewFuncCallValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(563)
			p.FuncCall()
		}

	case 6:
		localctx = NewReferenceValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(564)
			p.Reference()
		}

	}

	return localctx
}

// IReferenceAtomContext is an interface to support dynamic dispatch.
type IReferenceAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReferenceAtomContext differentiates from other interfaces.
	IsReferenceAtomContext()
}

type ReferenceAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceAtomContext() *ReferenceAtomContext {
	var p = new(ReferenceAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_referenceAtom
	return p
}

func (*ReferenceAtomContext) IsReferenceAtomContext() {}

func NewReferenceAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceAtomContext {
	var p = new(ReferenceAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_referenceAtom

	return p
}

func (s *ReferenceAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceAtomContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ReferenceAtomContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBracket, 0)
}

func (s *ReferenceAtomContext) ArrayIndex() IArrayIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayIndexContext)
}

func (s *ReferenceAtomContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBracket, 0)
}

func (s *ReferenceAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterReferenceAtom(s)
	}
}

func (s *ReferenceAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitReferenceAtom(s)
	}
}

func (p *DaedalusParser) ReferenceAtom() (localctx IReferenceAtomContext) {
	this := p
	_ = this

	localctx = NewReferenceAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, DaedalusParserRULE_referenceAtom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(567)
		p.NameNode()
	}
	p.SetState(572)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(568)
			p.Match(DaedalusParserLeftBracket)
		}
		{
			p.SetState(569)
			p.ArrayIndex()
		}
		{
			p.SetState(570)
			p.Match(DaedalusParserRightBracket)
		}

	}

	return localctx
}

// IReferenceContext is an interface to support dynamic dispatch.
type IReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReferenceContext differentiates from other interfaces.
	IsReferenceContext()
}

type ReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceContext() *ReferenceContext {
	var p = new(ReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_reference
	return p
}

func (*ReferenceContext) IsReferenceContext() {}

func NewReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceContext {
	var p = new(ReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_reference

	return p
}

func (s *ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceContext) AllReferenceAtom() []IReferenceAtomContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IReferenceAtomContext); ok {
			len++
		}
	}

	tst := make([]IReferenceAtomContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IReferenceAtomContext); ok {
			tst[i] = t.(IReferenceAtomContext)
			i++
		}
	}

	return tst
}

func (s *ReferenceContext) ReferenceAtom(i int) IReferenceAtomContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceAtomContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceAtomContext)
}

func (s *ReferenceContext) Dot() antlr.TerminalNode {
	return s.GetToken(DaedalusParserDot, 0)
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterReference(s)
	}
}

func (s *ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitReference(s)
	}
}

func (p *DaedalusParser) Reference() (localctx IReferenceContext) {
	this := p
	_ = this

	localctx = NewReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, DaedalusParserRULE_reference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(574)
		p.ReferenceAtom()
	}
	p.SetState(577)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(575)
			p.Match(DaedalusParserDot)
		}
		{
			p.SetState(576)
			p.ReferenceAtom()
		}

	}

	return localctx
}

// ITypeReferenceContext is an interface to support dynamic dispatch.
type ITypeReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeReferenceContext differentiates from other interfaces.
	IsTypeReferenceContext()
}

type TypeReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeReferenceContext() *TypeReferenceContext {
	var p = new(TypeReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_typeReference
	return p
}

func (*TypeReferenceContext) IsTypeReferenceContext() {}

func NewTypeReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeReferenceContext {
	var p = new(TypeReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_typeReference

	return p
}

func (s *TypeReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeReferenceContext) Identifier() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIdentifier, 0)
}

func (s *TypeReferenceContext) Void() antlr.TerminalNode {
	return s.GetToken(DaedalusParserVoid, 0)
}

func (s *TypeReferenceContext) Int() antlr.TerminalNode {
	return s.GetToken(DaedalusParserInt, 0)
}

func (s *TypeReferenceContext) Float() antlr.TerminalNode {
	return s.GetToken(DaedalusParserFloat, 0)
}

func (s *TypeReferenceContext) StringKeyword() antlr.TerminalNode {
	return s.GetToken(DaedalusParserStringKeyword, 0)
}

func (s *TypeReferenceContext) Func() antlr.TerminalNode {
	return s.GetToken(DaedalusParserFunc, 0)
}

func (s *TypeReferenceContext) Instance() antlr.TerminalNode {
	return s.GetToken(DaedalusParserInstance, 0)
}

func (s *TypeReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterTypeReference(s)
	}
}

func (s *TypeReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitTypeReference(s)
	}
}

func (p *DaedalusParser) TypeReference() (localctx ITypeReferenceContext) {
	this := p
	_ = this

	localctx = NewTypeReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, DaedalusParserRULE_typeReference)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(579)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-17)) & ^0x3f) == 0 && ((int64(1)<<(_la-17))&281474976712641) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAnyIdentifierContext is an interface to support dynamic dispatch.
type IAnyIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnyIdentifierContext differentiates from other interfaces.
	IsAnyIdentifierContext()
}

type AnyIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnyIdentifierContext() *AnyIdentifierContext {
	var p = new(AnyIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_anyIdentifier
	return p
}

func (*AnyIdentifierContext) IsAnyIdentifierContext() {}

func NewAnyIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnyIdentifierContext {
	var p = new(AnyIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_anyIdentifier

	return p
}

func (s *AnyIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *AnyIdentifierContext) Void() antlr.TerminalNode {
	return s.GetToken(DaedalusParserVoid, 0)
}

func (s *AnyIdentifierContext) Var() antlr.TerminalNode {
	return s.GetToken(DaedalusParserVar, 0)
}

func (s *AnyIdentifierContext) Int() antlr.TerminalNode {
	return s.GetToken(DaedalusParserInt, 0)
}

func (s *AnyIdentifierContext) Float() antlr.TerminalNode {
	return s.GetToken(DaedalusParserFloat, 0)
}

func (s *AnyIdentifierContext) StringKeyword() antlr.TerminalNode {
	return s.GetToken(DaedalusParserStringKeyword, 0)
}

func (s *AnyIdentifierContext) Func() antlr.TerminalNode {
	return s.GetToken(DaedalusParserFunc, 0)
}

func (s *AnyIdentifierContext) Instance() antlr.TerminalNode {
	return s.GetToken(DaedalusParserInstance, 0)
}

func (s *AnyIdentifierContext) Class() antlr.TerminalNode {
	return s.GetToken(DaedalusParserClass, 0)
}

func (s *AnyIdentifierContext) Prototype() antlr.TerminalNode {
	return s.GetToken(DaedalusParserPrototype, 0)
}

func (s *AnyIdentifierContext) Null() antlr.TerminalNode {
	return s.GetToken(DaedalusParserNull, 0)
}

func (s *AnyIdentifierContext) Meta() antlr.TerminalNode {
	return s.GetToken(DaedalusParserMeta, 0)
}

func (s *AnyIdentifierContext) Namespace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserNamespace, 0)
}

func (s *AnyIdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIdentifier, 0)
}

func (s *AnyIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnyIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterAnyIdentifier(s)
	}
}

func (s *AnyIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitAnyIdentifier(s)
	}
}

func (p *DaedalusParser) AnyIdentifier() (localctx IAnyIdentifierContext) {
	this := p
	_ = this

	localctx = NewAnyIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, DaedalusParserRULE_anyIdentifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(581)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-13)) & ^0x3f) == 0 && ((int64(1)<<(_la-13))&4503599627468445) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INameNodeContext is an interface to support dynamic dispatch.
type INameNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameNodeContext differentiates from other interfaces.
	IsNameNodeContext()
}

type NameNodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameNodeContext() *NameNodeContext {
	var p = new(NameNodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_nameNode
	return p
}

func (*NameNodeContext) IsNameNodeContext() {}

func NewNameNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameNodeContext {
	var p = new(NameNodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_nameNode

	return p
}

func (s *NameNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *NameNodeContext) AnyIdentifier() IAnyIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyIdentifierContext)
}

func (s *NameNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterNameNode(s)
	}
}

func (s *NameNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitNameNode(s)
	}
}

func (p *DaedalusParser) NameNode() (localctx INameNodeContext) {
	this := p
	_ = this

	localctx = NewNameNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, DaedalusParserRULE_nameNode)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(583)
		p.AnyIdentifier()
	}

	return localctx
}

// IParentReferenceContext is an interface to support dynamic dispatch.
type IParentReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParentReferenceContext differentiates from other interfaces.
	IsParentReferenceContext()
}

type ParentReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParentReferenceContext() *ParentReferenceContext {
	var p = new(ParentReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_parentReference
	return p
}

func (*ParentReferenceContext) IsParentReferenceContext() {}

func NewParentReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParentReferenceContext {
	var p = new(ParentReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_parentReference

	return p
}

func (s *ParentReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ParentReferenceContext) Identifier() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIdentifier, 0)
}

func (s *ParentReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParentReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterParentReference(s)
	}
}

func (s *ParentReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitParentReference(s)
	}
}

func (p *DaedalusParser) ParentReference() (localctx IParentReferenceContext) {
	this := p
	_ = this

	localctx = NewParentReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, DaedalusParserRULE_parentReference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(585)
		p.Match(DaedalusParserIdentifier)
	}

	return localctx
}

// IAssignmentOperatorContext is an interface to support dynamic dispatch.
type IAssignmentOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentOperatorContext differentiates from other interfaces.
	IsAssignmentOperatorContext()
}

type AssignmentOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentOperatorContext() *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_assignmentOperator
	return p
}

func (*AssignmentOperatorContext) IsAssignmentOperatorContext() {}

func NewAssignmentOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_assignmentOperator

	return p
}

func (s *AssignmentOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentOperatorContext) Assign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserAssign, 0)
}

func (s *AssignmentOperatorContext) StarAssign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserStarAssign, 0)
}

func (s *AssignmentOperatorContext) DivAssign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserDivAssign, 0)
}

func (s *AssignmentOperatorContext) PlusAssign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserPlusAssign, 0)
}

func (s *AssignmentOperatorContext) MinusAssign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserMinusAssign, 0)
}

func (s *AssignmentOperatorContext) AndAssign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserAndAssign, 0)
}

func (s *AssignmentOperatorContext) OrAssign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserOrAssign, 0)
}

func (s *AssignmentOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitAssignmentOperator(s)
	}
}

func (p *DaedalusParser) AssignmentOperator() (localctx IAssignmentOperatorContext) {
	this := p
	_ = this

	localctx = NewAssignmentOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, DaedalusParserRULE_assignmentOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(587)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9097271247288401920) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnaryOperatorContext is an interface to support dynamic dispatch.
type IUnaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOperatorContext differentiates from other interfaces.
	IsUnaryOperatorContext()
}

type UnaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperatorContext() *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_unaryOperator
	return p
}

func (*UnaryOperatorContext) IsUnaryOperatorContext() {}

func NewUnaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_unaryOperator

	return p
}

func (s *UnaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOperatorContext) Plus() antlr.TerminalNode {
	return s.GetToken(DaedalusParserPlus, 0)
}

func (s *UnaryOperatorContext) Tilde() antlr.TerminalNode {
	return s.GetToken(DaedalusParserTilde, 0)
}

func (s *UnaryOperatorContext) Minus() antlr.TerminalNode {
	return s.GetToken(DaedalusParserMinus, 0)
}

func (s *UnaryOperatorContext) Not() antlr.TerminalNode {
	return s.GetToken(DaedalusParserNot, 0)
}

func (s *UnaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitUnaryOperator(s)
	}
}

func (p *DaedalusParser) UnaryOperator() (localctx IUnaryOperatorContext) {
	this := p
	_ = this

	localctx = NewUnaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, DaedalusParserRULE_unaryOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(589)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14355223812243456) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAddOperatorContext is an interface to support dynamic dispatch.
type IAddOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddOperatorContext differentiates from other interfaces.
	IsAddOperatorContext()
}

type AddOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddOperatorContext() *AddOperatorContext {
	var p = new(AddOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_addOperator
	return p
}

func (*AddOperatorContext) IsAddOperatorContext() {}

func NewAddOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddOperatorContext {
	var p = new(AddOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_addOperator

	return p
}

func (s *AddOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AddOperatorContext) Plus() antlr.TerminalNode {
	return s.GetToken(DaedalusParserPlus, 0)
}

func (s *AddOperatorContext) Minus() antlr.TerminalNode {
	return s.GetToken(DaedalusParserMinus, 0)
}

func (s *AddOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterAddOperator(s)
	}
}

func (s *AddOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitAddOperator(s)
	}
}

func (p *DaedalusParser) AddOperator() (localctx IAddOperatorContext) {
	this := p
	_ = this

	localctx = NewAddOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, DaedalusParserRULE_addOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(591)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DaedalusParserPlus || _la == DaedalusParserMinus) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBitMoveOperatorContext is an interface to support dynamic dispatch.
type IBitMoveOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBitMoveOperatorContext differentiates from other interfaces.
	IsBitMoveOperatorContext()
}

type BitMoveOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitMoveOperatorContext() *BitMoveOperatorContext {
	var p = new(BitMoveOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_bitMoveOperator
	return p
}

func (*BitMoveOperatorContext) IsBitMoveOperatorContext() {}

func NewBitMoveOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitMoveOperatorContext {
	var p = new(BitMoveOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_bitMoveOperator

	return p
}

func (s *BitMoveOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *BitMoveOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitMoveOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitMoveOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterBitMoveOperator(s)
	}
}

func (s *BitMoveOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitBitMoveOperator(s)
	}
}

func (p *DaedalusParser) BitMoveOperator() (localctx IBitMoveOperatorContext) {
	this := p
	_ = this

	localctx = NewBitMoveOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, DaedalusParserRULE_bitMoveOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(593)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DaedalusParserT__2 || _la == DaedalusParserT__3) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICompOperatorContext is an interface to support dynamic dispatch.
type ICompOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompOperatorContext differentiates from other interfaces.
	IsCompOperatorContext()
}

type CompOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompOperatorContext() *CompOperatorContext {
	var p = new(CompOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_compOperator
	return p
}

func (*CompOperatorContext) IsCompOperatorContext() {}

func NewCompOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompOperatorContext {
	var p = new(CompOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_compOperator

	return p
}

func (s *CompOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *CompOperatorContext) Less() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLess, 0)
}

func (s *CompOperatorContext) Greater() antlr.TerminalNode {
	return s.GetToken(DaedalusParserGreater, 0)
}

func (s *CompOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterCompOperator(s)
	}
}

func (s *CompOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitCompOperator(s)
	}
}

func (p *DaedalusParser) CompOperator() (localctx ICompOperatorContext) {
	this := p
	_ = this

	localctx = NewCompOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, DaedalusParserRULE_compOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(595)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&108086391056892000) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEqOperatorContext is an interface to support dynamic dispatch.
type IEqOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqOperatorContext differentiates from other interfaces.
	IsEqOperatorContext()
}

type EqOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqOperatorContext() *EqOperatorContext {
	var p = new(EqOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_eqOperator
	return p
}

func (*EqOperatorContext) IsEqOperatorContext() {}

func NewEqOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqOperatorContext {
	var p = new(EqOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_eqOperator

	return p
}

func (s *EqOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *EqOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterEqOperator(s)
	}
}

func (s *EqOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitEqOperator(s)
	}
}

func (p *DaedalusParser) EqOperator() (localctx IEqOperatorContext) {
	this := p
	_ = this

	localctx = NewEqOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, DaedalusParserRULE_eqOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(597)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DaedalusParserT__6 || _la == DaedalusParserT__7) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMultOperatorContext is an interface to support dynamic dispatch.
type IMultOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultOperatorContext differentiates from other interfaces.
	IsMultOperatorContext()
}

type MultOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultOperatorContext() *MultOperatorContext {
	var p = new(MultOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_multOperator
	return p
}

func (*MultOperatorContext) IsMultOperatorContext() {}

func NewMultOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultOperatorContext {
	var p = new(MultOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_multOperator

	return p
}

func (s *MultOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MultOperatorContext) Star() antlr.TerminalNode {
	return s.GetToken(DaedalusParserStar, 0)
}

func (s *MultOperatorContext) Div() antlr.TerminalNode {
	return s.GetToken(DaedalusParserDiv, 0)
}

func (s *MultOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterMultOperator(s)
	}
}

func (s *MultOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitMultOperator(s)
	}
}

func (p *DaedalusParser) MultOperator() (localctx IMultOperatorContext) {
	this := p
	_ = this

	localctx = NewMultOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, DaedalusParserRULE_multOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(599)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3377699720528384) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBinAndOperatorContext is an interface to support dynamic dispatch.
type IBinAndOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinAndOperatorContext differentiates from other interfaces.
	IsBinAndOperatorContext()
}

type BinAndOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinAndOperatorContext() *BinAndOperatorContext {
	var p = new(BinAndOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_binAndOperator
	return p
}

func (*BinAndOperatorContext) IsBinAndOperatorContext() {}

func NewBinAndOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinAndOperatorContext {
	var p = new(BinAndOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_binAndOperator

	return p
}

func (s *BinAndOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *BinAndOperatorContext) BitAnd() antlr.TerminalNode {
	return s.GetToken(DaedalusParserBitAnd, 0)
}

func (s *BinAndOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinAndOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinAndOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterBinAndOperator(s)
	}
}

func (s *BinAndOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitBinAndOperator(s)
	}
}

func (p *DaedalusParser) BinAndOperator() (localctx IBinAndOperatorContext) {
	this := p
	_ = this

	localctx = NewBinAndOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, DaedalusParserRULE_binAndOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(601)
		p.Match(DaedalusParserBitAnd)
	}

	return localctx
}

// IBinOrOperatorContext is an interface to support dynamic dispatch.
type IBinOrOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinOrOperatorContext differentiates from other interfaces.
	IsBinOrOperatorContext()
}

type BinOrOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinOrOperatorContext() *BinOrOperatorContext {
	var p = new(BinOrOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_binOrOperator
	return p
}

func (*BinOrOperatorContext) IsBinOrOperatorContext() {}

func NewBinOrOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinOrOperatorContext {
	var p = new(BinOrOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_binOrOperator

	return p
}

func (s *BinOrOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *BinOrOperatorContext) BitOr() antlr.TerminalNode {
	return s.GetToken(DaedalusParserBitOr, 0)
}

func (s *BinOrOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinOrOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinOrOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterBinOrOperator(s)
	}
}

func (s *BinOrOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitBinOrOperator(s)
	}
}

func (p *DaedalusParser) BinOrOperator() (localctx IBinOrOperatorContext) {
	this := p
	_ = this

	localctx = NewBinOrOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, DaedalusParserRULE_binOrOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(603)
		p.Match(DaedalusParserBitOr)
	}

	return localctx
}

// ILogAndOperatorContext is an interface to support dynamic dispatch.
type ILogAndOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogAndOperatorContext differentiates from other interfaces.
	IsLogAndOperatorContext()
}

type LogAndOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogAndOperatorContext() *LogAndOperatorContext {
	var p = new(LogAndOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_logAndOperator
	return p
}

func (*LogAndOperatorContext) IsLogAndOperatorContext() {}

func NewLogAndOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogAndOperatorContext {
	var p = new(LogAndOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_logAndOperator

	return p
}

func (s *LogAndOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LogAndOperatorContext) And() antlr.TerminalNode {
	return s.GetToken(DaedalusParserAnd, 0)
}

func (s *LogAndOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogAndOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogAndOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterLogAndOperator(s)
	}
}

func (s *LogAndOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitLogAndOperator(s)
	}
}

func (p *DaedalusParser) LogAndOperator() (localctx ILogAndOperatorContext) {
	this := p
	_ = this

	localctx = NewLogAndOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, DaedalusParserRULE_logAndOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(605)
		p.Match(DaedalusParserAnd)
	}

	return localctx
}

// ILogOrOperatorContext is an interface to support dynamic dispatch.
type ILogOrOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogOrOperatorContext differentiates from other interfaces.
	IsLogOrOperatorContext()
}

type LogOrOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogOrOperatorContext() *LogOrOperatorContext {
	var p = new(LogOrOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_logOrOperator
	return p
}

func (*LogOrOperatorContext) IsLogOrOperatorContext() {}

func NewLogOrOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogOrOperatorContext {
	var p = new(LogOrOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_logOrOperator

	return p
}

func (s *LogOrOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LogOrOperatorContext) Or() antlr.TerminalNode {
	return s.GetToken(DaedalusParserOr, 0)
}

func (s *LogOrOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogOrOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogOrOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterLogOrOperator(s)
	}
}

func (s *LogOrOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitLogOrOperator(s)
	}
}

func (p *DaedalusParser) LogOrOperator() (localctx ILogOrOperatorContext) {
	this := p
	_ = this

	localctx = NewLogOrOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, DaedalusParserRULE_logOrOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(607)
		p.Match(DaedalusParserOr)
	}

	return localctx
}

func (p *DaedalusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 49:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *DaedalusParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
