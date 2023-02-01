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
		"", "','", "'<<'", "'>>'", "'<='", "'>='", "'=='", "'!='", "'%'", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "'('", "')'", "'['", "']'", "'{'", "'}'",
		"'&'", "'&&'", "'|'", "'||'", "'+'", "'-'", "'/'", "'*'", "'~'", "'!'",
		"'='", "'<'", "'>'", "'+='", "'-='", "'*='", "'/='", "'&='", "'|='",
		"'.'", "';'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "IntegerLiteral", "FloatLiteral",
		"StringLiteral", "Const", "Var", "If", "Int", "Else", "Func", "StringKeyword",
		"Class", "Void", "Return", "Float", "Prototype", "Instance", "Namespace",
		"Null", "Meta", "While", "Continue", "Break", "Mif", "Melif", "Melse",
		"Mendif", "LeftParen", "RightParen", "LeftBracket", "RightBracket",
		"LeftBrace", "RightBrace", "BitAnd", "And", "BitOr", "Or", "Plus", "Minus",
		"Div", "Star", "Tilde", "Not", "Assign", "Less", "Greater", "PlusAssign",
		"MinusAssign", "StarAssign", "DivAssign", "AndAssign", "OrAssign", "Dot",
		"Semi", "Identifier", "Whitespace", "Newline", "BlockComment", "LineComment",
	}
	staticData.ruleNames = []string{
		"daedalusFile", "blockDef", "inlineDef", "macroBlock", "macroCondition",
		"macroElseBlock", "macroElseIfBlock", "macroIfBlock", "macroDef", "functionDef",
		"constDef", "classDef", "prototypeDef", "instanceDef", "instanceDecl",
		"namespaceDef", "mainBlock", "contentBlock", "varDecl", "metaValue",
		"zParserExtenderMeta", "zParserExtenderMetaBlock", "constArrayDef",
		"constArrayAssignment", "constValueDef", "constValueAssignment", "varArrayDecl",
		"varValueDecl", "parameterList", "parameterDecl", "statementBlock",
		"statement", "funcCall", "assignment", "ifCondition", "elseBlock", "elseIfBlock",
		"ifBlock", "ifBlockStatement", "returnStatement", "whileCondition",
		"whileBlock", "whileBlockStatement", "funcArgExpression", "expressionBlock",
		"expression", "arrayIndex", "arraySize", "value", "referenceAtom", "reference",
		"typeReference", "anyIdentifier", "nameNode", "parentReference", "assignmentOperator",
		"unaryOperator", "addOperator", "bitMoveOperator", "compOperator", "eqOperator",
		"multOperator", "binAndOperator", "binOrOperator", "logAndOperator",
		"logOrOperator",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 66, 577, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 1, 0, 3, 0, 134, 8, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 143, 8, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 3, 2, 150, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 3, 3, 160, 8, 3, 1, 3, 1, 3, 1, 3, 3, 3, 165, 8, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 5, 5, 171, 8, 5, 10, 5, 12, 5, 174, 9, 5, 1, 6, 1, 6, 1, 6, 5, 6,
		179, 8, 6, 10, 6, 12, 6, 182, 9, 6, 1, 7, 1, 7, 1, 7, 5, 7, 187, 8, 7,
		10, 7, 12, 7, 190, 9, 7, 1, 8, 1, 8, 5, 8, 194, 8, 8, 10, 8, 12, 8, 197,
		9, 8, 1, 8, 3, 8, 200, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 214, 8, 10, 1, 10, 1, 10, 1, 10,
		3, 10, 219, 8, 10, 5, 10, 221, 8, 10, 10, 10, 12, 10, 224, 9, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 232, 8, 11, 10, 11, 12, 11, 235,
		9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		5, 14, 257, 8, 14, 10, 14, 12, 14, 260, 9, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 270, 8, 15, 10, 15, 12, 15, 273,
		9, 15, 1, 15, 1, 15, 1, 16, 3, 16, 278, 8, 16, 1, 16, 4, 16, 281, 8, 16,
		11, 16, 12, 16, 282, 1, 17, 1, 17, 1, 17, 3, 17, 288, 8, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 3, 18, 294, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18,
		300, 8, 18, 5, 18, 302, 8, 18, 10, 18, 12, 18, 305, 9, 18, 1, 19, 4, 19,
		308, 8, 19, 11, 19, 12, 19, 309, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 21, 5, 21, 320, 8, 21, 10, 21, 12, 21, 323, 9, 21, 1, 21,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 5, 23, 339, 8, 23, 10, 23, 12, 23, 342, 9, 23, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 363, 8, 28,
		10, 28, 12, 28, 366, 9, 28, 3, 28, 368, 8, 28, 1, 28, 1, 28, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 379, 8, 29, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 387, 8, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 5, 30, 393, 8, 30, 10, 30, 12, 30, 396, 9, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 407, 8, 31, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 5, 32, 414, 8, 32, 10, 32, 12, 32, 417, 9, 32,
		3, 32, 419, 8, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1,
		34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 38, 1, 38, 5, 38, 443, 8, 38, 10, 38, 12, 38, 446, 9,
		38, 1, 38, 3, 38, 449, 8, 38, 1, 39, 1, 39, 3, 39, 453, 8, 39, 1, 40, 1,
		40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44,
		1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 476,
		8, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45,
		1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 5, 45, 514, 8, 45, 10, 45, 12, 45,
		517, 9, 45, 1, 46, 1, 46, 3, 46, 521, 8, 46, 1, 47, 1, 47, 3, 47, 525,
		8, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 3, 48, 533, 8, 48, 1,
		49, 1, 49, 1, 49, 1, 49, 1, 49, 3, 49, 540, 8, 49, 1, 50, 1, 50, 1, 50,
		3, 50, 545, 8, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1,
		54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59,
		1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1,
		65, 1, 65, 1, 65, 11, 195, 233, 258, 271, 309, 321, 340, 364, 394, 415,
		444, 1, 90, 66, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
		66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100,
		102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130,
		0, 9, 6, 0, 15, 15, 17, 18, 20, 20, 22, 22, 24, 24, 62, 62, 5, 0, 13, 13,
		15, 15, 17, 20, 22, 27, 62, 62, 2, 0, 51, 51, 54, 59, 2, 0, 45, 46, 49,
		50, 1, 0, 45, 46, 1, 0, 2, 3, 2, 0, 4, 5, 52, 53, 1, 0, 6, 7, 2, 0, 8,
		8, 47, 48, 582, 0, 133, 1, 0, 0, 0, 2, 142, 1, 0, 0, 0, 4, 149, 1, 0, 0,
		0, 6, 164, 1, 0, 0, 0, 8, 166, 1, 0, 0, 0, 10, 168, 1, 0, 0, 0, 12, 175,
		1, 0, 0, 0, 14, 183, 1, 0, 0, 0, 16, 191, 1, 0, 0, 0, 18, 203, 1, 0, 0,
		0, 20, 209, 1, 0, 0, 0, 22, 225, 1, 0, 0, 0, 24, 238, 1, 0, 0, 0, 26, 245,
		1, 0, 0, 0, 28, 252, 1, 0, 0, 0, 30, 265, 1, 0, 0, 0, 32, 277, 1, 0, 0,
		0, 34, 287, 1, 0, 0, 0, 36, 289, 1, 0, 0, 0, 38, 307, 1, 0, 0, 0, 40, 311,
		1, 0, 0, 0, 42, 316, 1, 0, 0, 0, 44, 327, 1, 0, 0, 0, 46, 333, 1, 0, 0,
		0, 48, 345, 1, 0, 0, 0, 50, 348, 1, 0, 0, 0, 52, 351, 1, 0, 0, 0, 54, 356,
		1, 0, 0, 0, 56, 358, 1, 0, 0, 0, 58, 371, 1, 0, 0, 0, 60, 380, 1, 0, 0,
		0, 62, 406, 1, 0, 0, 0, 64, 408, 1, 0, 0, 0, 66, 422, 1, 0, 0, 0, 68, 426,
		1, 0, 0, 0, 70, 428, 1, 0, 0, 0, 72, 431, 1, 0, 0, 0, 74, 436, 1, 0, 0,
		0, 76, 440, 1, 0, 0, 0, 78, 450, 1, 0, 0, 0, 80, 454, 1, 0, 0, 0, 82, 456,
		1, 0, 0, 0, 84, 460, 1, 0, 0, 0, 86, 462, 1, 0, 0, 0, 88, 464, 1, 0, 0,
		0, 90, 475, 1, 0, 0, 0, 92, 520, 1, 0, 0, 0, 94, 524, 1, 0, 0, 0, 96, 532,
		1, 0, 0, 0, 98, 534, 1, 0, 0, 0, 100, 541, 1, 0, 0, 0, 102, 546, 1, 0,
		0, 0, 104, 548, 1, 0, 0, 0, 106, 550, 1, 0, 0, 0, 108, 552, 1, 0, 0, 0,
		110, 554, 1, 0, 0, 0, 112, 556, 1, 0, 0, 0, 114, 558, 1, 0, 0, 0, 116,
		560, 1, 0, 0, 0, 118, 562, 1, 0, 0, 0, 120, 564, 1, 0, 0, 0, 122, 566,
		1, 0, 0, 0, 124, 568, 1, 0, 0, 0, 126, 570, 1, 0, 0, 0, 128, 572, 1, 0,
		0, 0, 130, 574, 1, 0, 0, 0, 132, 134, 3, 32, 16, 0, 133, 132, 1, 0, 0,
		0, 133, 134, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 136, 5, 0, 0, 1, 136,
		1, 1, 0, 0, 0, 137, 143, 3, 18, 9, 0, 138, 143, 3, 22, 11, 0, 139, 143,
		3, 24, 12, 0, 140, 143, 3, 26, 13, 0, 141, 143, 3, 30, 15, 0, 142, 137,
		1, 0, 0, 0, 142, 138, 1, 0, 0, 0, 142, 139, 1, 0, 0, 0, 142, 140, 1, 0,
		0, 0, 142, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 5, 61, 0, 0,
		145, 3, 1, 0, 0, 0, 146, 150, 3, 20, 10, 0, 147, 150, 3, 36, 18, 0, 148,
		150, 3, 28, 14, 0, 149, 146, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 148,
		1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 152, 5, 61, 0, 0, 152, 5, 1, 0,
		0, 0, 153, 165, 3, 2, 1, 0, 154, 155, 3, 62, 31, 0, 155, 156, 5, 61, 0,
		0, 156, 165, 1, 0, 0, 0, 157, 159, 3, 76, 38, 0, 158, 160, 5, 61, 0, 0,
		159, 158, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 165, 1, 0, 0, 0, 161,
		162, 3, 84, 42, 0, 162, 163, 5, 61, 0, 0, 163, 165, 1, 0, 0, 0, 164, 153,
		1, 0, 0, 0, 164, 154, 1, 0, 0, 0, 164, 157, 1, 0, 0, 0, 164, 161, 1, 0,
		0, 0, 165, 7, 1, 0, 0, 0, 166, 167, 3, 88, 44, 0, 167, 9, 1, 0, 0, 0, 168,
		172, 5, 33, 0, 0, 169, 171, 3, 6, 3, 0, 170, 169, 1, 0, 0, 0, 171, 174,
		1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 11, 1, 0,
		0, 0, 174, 172, 1, 0, 0, 0, 175, 176, 5, 32, 0, 0, 176, 180, 3, 8, 4, 0,
		177, 179, 3, 6, 3, 0, 178, 177, 1, 0, 0, 0, 179, 182, 1, 0, 0, 0, 180,
		178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 13, 1, 0, 0, 0, 182, 180, 1,
		0, 0, 0, 183, 184, 5, 31, 0, 0, 184, 188, 3, 8, 4, 0, 185, 187, 3, 6, 3,
		0, 186, 185, 1, 0, 0, 0, 187, 190, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 188,
		189, 1, 0, 0, 0, 189, 15, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 191, 195, 3,
		14, 7, 0, 192, 194, 3, 12, 6, 0, 193, 192, 1, 0, 0, 0, 194, 197, 1, 0,
		0, 0, 195, 196, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 196, 199, 1, 0, 0, 0,
		197, 195, 1, 0, 0, 0, 198, 200, 3, 10, 5, 0, 199, 198, 1, 0, 0, 0, 199,
		200, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202, 5, 34, 0, 0, 202, 17,
		1, 0, 0, 0, 203, 204, 5, 17, 0, 0, 204, 205, 3, 102, 51, 0, 205, 206, 3,
		106, 53, 0, 206, 207, 3, 56, 28, 0, 207, 208, 3, 60, 30, 0, 208, 19, 1,
		0, 0, 0, 209, 210, 5, 12, 0, 0, 210, 213, 3, 102, 51, 0, 211, 214, 3, 48,
		24, 0, 212, 214, 3, 44, 22, 0, 213, 211, 1, 0, 0, 0, 213, 212, 1, 0, 0,
		0, 214, 222, 1, 0, 0, 0, 215, 218, 5, 1, 0, 0, 216, 219, 3, 48, 24, 0,
		217, 219, 3, 44, 22, 0, 218, 216, 1, 0, 0, 0, 218, 217, 1, 0, 0, 0, 219,
		221, 1, 0, 0, 0, 220, 215, 1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222, 220,
		1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 21, 1, 0, 0, 0, 224, 222, 1, 0,
		0, 0, 225, 226, 5, 19, 0, 0, 226, 227, 3, 106, 53, 0, 227, 233, 5, 39,
		0, 0, 228, 229, 3, 36, 18, 0, 229, 230, 5, 61, 0, 0, 230, 232, 1, 0, 0,
		0, 231, 228, 1, 0, 0, 0, 232, 235, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 233,
		231, 1, 0, 0, 0, 234, 236, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 236, 237,
		5, 40, 0, 0, 237, 23, 1, 0, 0, 0, 238, 239, 5, 23, 0, 0, 239, 240, 3, 106,
		53, 0, 240, 241, 5, 35, 0, 0, 241, 242, 3, 108, 54, 0, 242, 243, 5, 36,
		0, 0, 243, 244, 3, 60, 30, 0, 244, 25, 1, 0, 0, 0, 245, 246, 5, 24, 0,
		0, 246, 247, 3, 106, 53, 0, 247, 248, 5, 35, 0, 0, 248, 249, 3, 108, 54,
		0, 249, 250, 5, 36, 0, 0, 250, 251, 3, 60, 30, 0, 251, 27, 1, 0, 0, 0,
		252, 253, 5, 24, 0, 0, 253, 258, 3, 106, 53, 0, 254, 255, 5, 1, 0, 0, 255,
		257, 3, 106, 53, 0, 256, 254, 1, 0, 0, 0, 257, 260, 1, 0, 0, 0, 258, 259,
		1, 0, 0, 0, 258, 256, 1, 0, 0, 0, 259, 261, 1, 0, 0, 0, 260, 258, 1, 0,
		0, 0, 261, 262, 5, 35, 0, 0, 262, 263, 3, 108, 54, 0, 263, 264, 5, 36,
		0, 0, 264, 29, 1, 0, 0, 0, 265, 266, 5, 25, 0, 0, 266, 267, 3, 106, 53,
		0, 267, 271, 5, 39, 0, 0, 268, 270, 3, 34, 17, 0, 269, 268, 1, 0, 0, 0,
		270, 273, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 272,
		274, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 274, 275, 5, 40, 0, 0, 275, 31,
		1, 0, 0, 0, 276, 278, 3, 42, 21, 0, 277, 276, 1, 0, 0, 0, 277, 278, 1,
		0, 0, 0, 278, 280, 1, 0, 0, 0, 279, 281, 3, 34, 17, 0, 280, 279, 1, 0,
		0, 0, 281, 282, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0,
		283, 33, 1, 0, 0, 0, 284, 288, 3, 2, 1, 0, 285, 288, 3, 4, 2, 0, 286, 288,
		3, 16, 8, 0, 287, 284, 1, 0, 0, 0, 287, 285, 1, 0, 0, 0, 287, 286, 1, 0,
		0, 0, 288, 35, 1, 0, 0, 0, 289, 290, 5, 13, 0, 0, 290, 293, 3, 102, 51,
		0, 291, 294, 3, 54, 27, 0, 292, 294, 3, 52, 26, 0, 293, 291, 1, 0, 0, 0,
		293, 292, 1, 0, 0, 0, 294, 303, 1, 0, 0, 0, 295, 299, 5, 1, 0, 0, 296,
		300, 3, 36, 18, 0, 297, 300, 3, 54, 27, 0, 298, 300, 3, 52, 26, 0, 299,
		296, 1, 0, 0, 0, 299, 297, 1, 0, 0, 0, 299, 298, 1, 0, 0, 0, 300, 302,
		1, 0, 0, 0, 301, 295, 1, 0, 0, 0, 302, 305, 1, 0, 0, 0, 303, 301, 1, 0,
		0, 0, 303, 304, 1, 0, 0, 0, 304, 37, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0,
		306, 308, 9, 0, 0, 0, 307, 306, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309,
		310, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 310, 39, 1, 0, 0, 0, 311, 312, 3,
		106, 53, 0, 312, 313, 5, 51, 0, 0, 313, 314, 3, 38, 19, 0, 314, 315, 5,
		61, 0, 0, 315, 41, 1, 0, 0, 0, 316, 317, 5, 27, 0, 0, 317, 321, 5, 39,
		0, 0, 318, 320, 3, 40, 20, 0, 319, 318, 1, 0, 0, 0, 320, 323, 1, 0, 0,
		0, 321, 322, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0, 322, 324, 1, 0, 0, 0, 323,
		321, 1, 0, 0, 0, 324, 325, 5, 40, 0, 0, 325, 326, 5, 61, 0, 0, 326, 43,
		1, 0, 0, 0, 327, 328, 3, 106, 53, 0, 328, 329, 5, 37, 0, 0, 329, 330, 3,
		94, 47, 0, 330, 331, 5, 38, 0, 0, 331, 332, 3, 46, 23, 0, 332, 45, 1, 0,
		0, 0, 333, 334, 5, 51, 0, 0, 334, 335, 5, 39, 0, 0, 335, 340, 3, 88, 44,
		0, 336, 337, 5, 1, 0, 0, 337, 339, 3, 88, 44, 0, 338, 336, 1, 0, 0, 0,
		339, 342, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 340, 338, 1, 0, 0, 0, 341,
		343, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0, 343, 344, 5, 40, 0, 0, 344, 47,
		1, 0, 0, 0, 345, 346, 3, 106, 53, 0, 346, 347, 3, 50, 25, 0, 347, 49, 1,
		0, 0, 0, 348, 349, 5, 51, 0, 0, 349, 350, 3, 88, 44, 0, 350, 51, 1, 0,
		0, 0, 351, 352, 3, 106, 53, 0, 352, 353, 5, 37, 0, 0, 353, 354, 3, 94,
		47, 0, 354, 355, 5, 38, 0, 0, 355, 53, 1, 0, 0, 0, 356, 357, 3, 106, 53,
		0, 357, 55, 1, 0, 0, 0, 358, 367, 5, 35, 0, 0, 359, 364, 3, 58, 29, 0,
		360, 361, 5, 1, 0, 0, 361, 363, 3, 58, 29, 0, 362, 360, 1, 0, 0, 0, 363,
		366, 1, 0, 0, 0, 364, 365, 1, 0, 0, 0, 364, 362, 1, 0, 0, 0, 365, 368,
		1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 367, 359, 1, 0, 0, 0, 367, 368, 1, 0,
		0, 0, 368, 369, 1, 0, 0, 0, 369, 370, 5, 36, 0, 0, 370, 57, 1, 0, 0, 0,
		371, 372, 5, 13, 0, 0, 372, 373, 3, 102, 51, 0, 373, 378, 3, 106, 53, 0,
		374, 375, 5, 37, 0, 0, 375, 376, 3, 94, 47, 0, 376, 377, 5, 38, 0, 0, 377,
		379, 1, 0, 0, 0, 378, 374, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 59, 1,
		0, 0, 0, 380, 394, 5, 39, 0, 0, 381, 382, 3, 62, 31, 0, 382, 383, 5, 61,
		0, 0, 383, 393, 1, 0, 0, 0, 384, 386, 3, 76, 38, 0, 385, 387, 5, 61, 0,
		0, 386, 385, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 393, 1, 0, 0, 0, 388,
		393, 3, 16, 8, 0, 389, 390, 3, 84, 42, 0, 390, 391, 5, 61, 0, 0, 391, 393,
		1, 0, 0, 0, 392, 381, 1, 0, 0, 0, 392, 384, 1, 0, 0, 0, 392, 388, 1, 0,
		0, 0, 392, 389, 1, 0, 0, 0, 393, 396, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0,
		394, 392, 1, 0, 0, 0, 395, 397, 1, 0, 0, 0, 396, 394, 1, 0, 0, 0, 397,
		398, 5, 40, 0, 0, 398, 61, 1, 0, 0, 0, 399, 407, 3, 66, 33, 0, 400, 407,
		3, 78, 39, 0, 401, 407, 3, 20, 10, 0, 402, 407, 3, 36, 18, 0, 403, 407,
		3, 90, 45, 0, 404, 407, 5, 29, 0, 0, 405, 407, 5, 30, 0, 0, 406, 399, 1,
		0, 0, 0, 406, 400, 1, 0, 0, 0, 406, 401, 1, 0, 0, 0, 406, 402, 1, 0, 0,
		0, 406, 403, 1, 0, 0, 0, 406, 404, 1, 0, 0, 0, 406, 405, 1, 0, 0, 0, 407,
		63, 1, 0, 0, 0, 408, 409, 3, 106, 53, 0, 409, 418, 5, 35, 0, 0, 410, 415,
		3, 86, 43, 0, 411, 412, 5, 1, 0, 0, 412, 414, 3, 86, 43, 0, 413, 411, 1,
		0, 0, 0, 414, 417, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0, 415, 413, 1, 0, 0,
		0, 416, 419, 1, 0, 0, 0, 417, 415, 1, 0, 0, 0, 418, 410, 1, 0, 0, 0, 418,
		419, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420, 421, 5, 36, 0, 0, 421, 65,
		1, 0, 0, 0, 422, 423, 3, 100, 50, 0, 423, 424, 3, 110, 55, 0, 424, 425,
		3, 88, 44, 0, 425, 67, 1, 0, 0, 0, 426, 427, 3, 88, 44, 0, 427, 69, 1,
		0, 0, 0, 428, 429, 5, 16, 0, 0, 429, 430, 3, 60, 30, 0, 430, 71, 1, 0,
		0, 0, 431, 432, 5, 16, 0, 0, 432, 433, 5, 14, 0, 0, 433, 434, 3, 68, 34,
		0, 434, 435, 3, 60, 30, 0, 435, 73, 1, 0, 0, 0, 436, 437, 5, 14, 0, 0,
		437, 438, 3, 68, 34, 0, 438, 439, 3, 60, 30, 0, 439, 75, 1, 0, 0, 0, 440,
		444, 3, 74, 37, 0, 441, 443, 3, 72, 36, 0, 442, 441, 1, 0, 0, 0, 443, 446,
		1, 0, 0, 0, 444, 445, 1, 0, 0, 0, 444, 442, 1, 0, 0, 0, 445, 448, 1, 0,
		0, 0, 446, 444, 1, 0, 0, 0, 447, 449, 3, 70, 35, 0, 448, 447, 1, 0, 0,
		0, 448, 449, 1, 0, 0, 0, 449, 77, 1, 0, 0, 0, 450, 452, 5, 21, 0, 0, 451,
		453, 3, 88, 44, 0, 452, 451, 1, 0, 0, 0, 452, 453, 1, 0, 0, 0, 453, 79,
		1, 0, 0, 0, 454, 455, 3, 88, 44, 0, 455, 81, 1, 0, 0, 0, 456, 457, 5, 28,
		0, 0, 457, 458, 3, 80, 40, 0, 458, 459, 3, 60, 30, 0, 459, 83, 1, 0, 0,
		0, 460, 461, 3, 82, 41, 0, 461, 85, 1, 0, 0, 0, 462, 463, 3, 88, 44, 0,
		463, 87, 1, 0, 0, 0, 464, 465, 3, 90, 45, 0, 465, 89, 1, 0, 0, 0, 466,
		467, 6, 45, -1, 0, 467, 468, 5, 35, 0, 0, 468, 469, 3, 90, 45, 0, 469,
		470, 5, 36, 0, 0, 470, 476, 1, 0, 0, 0, 471, 472, 3, 112, 56, 0, 472, 473,
		3, 90, 45, 11, 473, 476, 1, 0, 0, 0, 474, 476, 3, 96, 48, 0, 475, 466,
		1, 0, 0, 0, 475, 471, 1, 0, 0, 0, 475, 474, 1, 0, 0, 0, 476, 515, 1, 0,
		0, 0, 477, 478, 10, 10, 0, 0, 478, 479, 3, 122, 61, 0, 479, 480, 3, 90,
		45, 11, 480, 514, 1, 0, 0, 0, 481, 482, 10, 9, 0, 0, 482, 483, 3, 114,
		57, 0, 483, 484, 3, 90, 45, 10, 484, 514, 1, 0, 0, 0, 485, 486, 10, 8,
		0, 0, 486, 487, 3, 116, 58, 0, 487, 488, 3, 90, 45, 9, 488, 514, 1, 0,
		0, 0, 489, 490, 10, 7, 0, 0, 490, 491, 3, 118, 59, 0, 491, 492, 3, 90,
		45, 8, 492, 514, 1, 0, 0, 0, 493, 494, 10, 6, 0, 0, 494, 495, 3, 120, 60,
		0, 495, 496, 3, 90, 45, 7, 496, 514, 1, 0, 0, 0, 497, 498, 10, 5, 0, 0,
		498, 499, 3, 124, 62, 0, 499, 500, 3, 90, 45, 6, 500, 514, 1, 0, 0, 0,
		501, 502, 10, 4, 0, 0, 502, 503, 3, 126, 63, 0, 503, 504, 3, 90, 45, 5,
		504, 514, 1, 0, 0, 0, 505, 506, 10, 3, 0, 0, 506, 507, 3, 128, 64, 0, 507,
		508, 3, 90, 45, 4, 508, 514, 1, 0, 0, 0, 509, 510, 10, 2, 0, 0, 510, 511,
		3, 130, 65, 0, 511, 512, 3, 90, 45, 3, 512, 514, 1, 0, 0, 0, 513, 477,
		1, 0, 0, 0, 513, 481, 1, 0, 0, 0, 513, 485, 1, 0, 0, 0, 513, 489, 1, 0,
		0, 0, 513, 493, 1, 0, 0, 0, 513, 497, 1, 0, 0, 0, 513, 501, 1, 0, 0, 0,
		513, 505, 1, 0, 0, 0, 513, 509, 1, 0, 0, 0, 514, 517, 1, 0, 0, 0, 515,
		513, 1, 0, 0, 0, 515, 516, 1, 0, 0, 0, 516, 91, 1, 0, 0, 0, 517, 515, 1,
		0, 0, 0, 518, 521, 5, 9, 0, 0, 519, 521, 3, 98, 49, 0, 520, 518, 1, 0,
		0, 0, 520, 519, 1, 0, 0, 0, 521, 93, 1, 0, 0, 0, 522, 525, 5, 9, 0, 0,
		523, 525, 3, 98, 49, 0, 524, 522, 1, 0, 0, 0, 524, 523, 1, 0, 0, 0, 525,
		95, 1, 0, 0, 0, 526, 533, 5, 9, 0, 0, 527, 533, 5, 10, 0, 0, 528, 533,
		5, 11, 0, 0, 529, 533, 5, 26, 0, 0, 530, 533, 3, 64, 32, 0, 531, 533, 3,
		100, 50, 0, 532, 526, 1, 0, 0, 0, 532, 527, 1, 0, 0, 0, 532, 528, 1, 0,
		0, 0, 532, 529, 1, 0, 0, 0, 532, 530, 1, 0, 0, 0, 532, 531, 1, 0, 0, 0,
		533, 97, 1, 0, 0, 0, 534, 539, 3, 106, 53, 0, 535, 536, 5, 37, 0, 0, 536,
		537, 3, 92, 46, 0, 537, 538, 5, 38, 0, 0, 538, 540, 1, 0, 0, 0, 539, 535,
		1, 0, 0, 0, 539, 540, 1, 0, 0, 0, 540, 99, 1, 0, 0, 0, 541, 544, 3, 98,
		49, 0, 542, 543, 5, 60, 0, 0, 543, 545, 3, 98, 49, 0, 544, 542, 1, 0, 0,
		0, 544, 545, 1, 0, 0, 0, 545, 101, 1, 0, 0, 0, 546, 547, 7, 0, 0, 0, 547,
		103, 1, 0, 0, 0, 548, 549, 7, 1, 0, 0, 549, 105, 1, 0, 0, 0, 550, 551,
		3, 104, 52, 0, 551, 107, 1, 0, 0, 0, 552, 553, 5, 62, 0, 0, 553, 109, 1,
		0, 0, 0, 554, 555, 7, 2, 0, 0, 555, 111, 1, 0, 0, 0, 556, 557, 7, 3, 0,
		0, 557, 113, 1, 0, 0, 0, 558, 559, 7, 4, 0, 0, 559, 115, 1, 0, 0, 0, 560,
		561, 7, 5, 0, 0, 561, 117, 1, 0, 0, 0, 562, 563, 7, 6, 0, 0, 563, 119,
		1, 0, 0, 0, 564, 565, 7, 7, 0, 0, 565, 121, 1, 0, 0, 0, 566, 567, 7, 8,
		0, 0, 567, 123, 1, 0, 0, 0, 568, 569, 5, 41, 0, 0, 569, 125, 1, 0, 0, 0,
		570, 571, 5, 43, 0, 0, 571, 127, 1, 0, 0, 0, 572, 573, 5, 42, 0, 0, 573,
		129, 1, 0, 0, 0, 574, 575, 5, 44, 0, 0, 575, 131, 1, 0, 0, 0, 45, 133,
		142, 149, 159, 164, 172, 180, 188, 195, 199, 213, 218, 222, 233, 258, 271,
		277, 282, 287, 293, 299, 303, 309, 321, 340, 364, 367, 378, 386, 392, 394,
		406, 415, 418, 444, 448, 452, 475, 513, 515, 520, 524, 532, 539, 544,
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
	DaedalusParserIntegerLiteral = 9
	DaedalusParserFloatLiteral   = 10
	DaedalusParserStringLiteral  = 11
	DaedalusParserConst          = 12
	DaedalusParserVar            = 13
	DaedalusParserIf             = 14
	DaedalusParserInt            = 15
	DaedalusParserElse           = 16
	DaedalusParserFunc           = 17
	DaedalusParserStringKeyword  = 18
	DaedalusParserClass          = 19
	DaedalusParserVoid           = 20
	DaedalusParserReturn         = 21
	DaedalusParserFloat          = 22
	DaedalusParserPrototype      = 23
	DaedalusParserInstance       = 24
	DaedalusParserNamespace      = 25
	DaedalusParserNull           = 26
	DaedalusParserMeta           = 27
	DaedalusParserWhile          = 28
	DaedalusParserContinue       = 29
	DaedalusParserBreak          = 30
	DaedalusParserMif            = 31
	DaedalusParserMelif          = 32
	DaedalusParserMelse          = 33
	DaedalusParserMendif         = 34
	DaedalusParserLeftParen      = 35
	DaedalusParserRightParen     = 36
	DaedalusParserLeftBracket    = 37
	DaedalusParserRightBracket   = 38
	DaedalusParserLeftBrace      = 39
	DaedalusParserRightBrace     = 40
	DaedalusParserBitAnd         = 41
	DaedalusParserAnd            = 42
	DaedalusParserBitOr          = 43
	DaedalusParserOr             = 44
	DaedalusParserPlus           = 45
	DaedalusParserMinus          = 46
	DaedalusParserDiv            = 47
	DaedalusParserStar           = 48
	DaedalusParserTilde          = 49
	DaedalusParserNot            = 50
	DaedalusParserAssign         = 51
	DaedalusParserLess           = 52
	DaedalusParserGreater        = 53
	DaedalusParserPlusAssign     = 54
	DaedalusParserMinusAssign    = 55
	DaedalusParserStarAssign     = 56
	DaedalusParserDivAssign      = 57
	DaedalusParserAndAssign      = 58
	DaedalusParserOrAssign       = 59
	DaedalusParserDot            = 60
	DaedalusParserSemi           = 61
	DaedalusParserIdentifier     = 62
	DaedalusParserWhitespace     = 63
	DaedalusParserNewline        = 64
	DaedalusParserBlockComment   = 65
	DaedalusParserLineComment    = 66
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
	DaedalusParserRULE_functionDef              = 9
	DaedalusParserRULE_constDef                 = 10
	DaedalusParserRULE_classDef                 = 11
	DaedalusParserRULE_prototypeDef             = 12
	DaedalusParserRULE_instanceDef              = 13
	DaedalusParserRULE_instanceDecl             = 14
	DaedalusParserRULE_namespaceDef             = 15
	DaedalusParserRULE_mainBlock                = 16
	DaedalusParserRULE_contentBlock             = 17
	DaedalusParserRULE_varDecl                  = 18
	DaedalusParserRULE_metaValue                = 19
	DaedalusParserRULE_zParserExtenderMeta      = 20
	DaedalusParserRULE_zParserExtenderMetaBlock = 21
	DaedalusParserRULE_constArrayDef            = 22
	DaedalusParserRULE_constArrayAssignment     = 23
	DaedalusParserRULE_constValueDef            = 24
	DaedalusParserRULE_constValueAssignment     = 25
	DaedalusParserRULE_varArrayDecl             = 26
	DaedalusParserRULE_varValueDecl             = 27
	DaedalusParserRULE_parameterList            = 28
	DaedalusParserRULE_parameterDecl            = 29
	DaedalusParserRULE_statementBlock           = 30
	DaedalusParserRULE_statement                = 31
	DaedalusParserRULE_funcCall                 = 32
	DaedalusParserRULE_assignment               = 33
	DaedalusParserRULE_ifCondition              = 34
	DaedalusParserRULE_elseBlock                = 35
	DaedalusParserRULE_elseIfBlock              = 36
	DaedalusParserRULE_ifBlock                  = 37
	DaedalusParserRULE_ifBlockStatement         = 38
	DaedalusParserRULE_returnStatement          = 39
	DaedalusParserRULE_whileCondition           = 40
	DaedalusParserRULE_whileBlock               = 41
	DaedalusParserRULE_whileBlockStatement      = 42
	DaedalusParserRULE_funcArgExpression        = 43
	DaedalusParserRULE_expressionBlock          = 44
	DaedalusParserRULE_expression               = 45
	DaedalusParserRULE_arrayIndex               = 46
	DaedalusParserRULE_arraySize                = 47
	DaedalusParserRULE_value                    = 48
	DaedalusParserRULE_referenceAtom            = 49
	DaedalusParserRULE_reference                = 50
	DaedalusParserRULE_typeReference            = 51
	DaedalusParserRULE_anyIdentifier            = 52
	DaedalusParserRULE_nameNode                 = 53
	DaedalusParserRULE_parentReference          = 54
	DaedalusParserRULE_assignmentOperator       = 55
	DaedalusParserRULE_unaryOperator            = 56
	DaedalusParserRULE_addOperator              = 57
	DaedalusParserRULE_bitMoveOperator          = 58
	DaedalusParserRULE_compOperator             = 59
	DaedalusParserRULE_eqOperator               = 60
	DaedalusParserRULE_multOperator             = 61
	DaedalusParserRULE_binAndOperator           = 62
	DaedalusParserRULE_binOrOperator            = 63
	DaedalusParserRULE_logAndOperator           = 64
	DaedalusParserRULE_logOrOperator            = 65
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
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2341089280) != 0 {
		{
			p.SetState(132)
			p.MainBlock()
		}

	}
	{
		p.SetState(135)
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
	p.SetState(142)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserFunc:
		{
			p.SetState(137)
			p.FunctionDef()
		}

	case DaedalusParserClass:
		{
			p.SetState(138)
			p.ClassDef()
		}

	case DaedalusParserPrototype:
		{
			p.SetState(139)
			p.PrototypeDef()
		}

	case DaedalusParserInstance:
		{
			p.SetState(140)
			p.InstanceDef()
		}

	case DaedalusParserNamespace:
		{
			p.SetState(141)
			p.NamespaceDef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(144)
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
	p.SetState(149)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserConst:
		{
			p.SetState(146)
			p.ConstDef()
		}

	case DaedalusParserVar:
		{
			p.SetState(147)
			p.VarDecl()
		}

	case DaedalusParserInstance:
		{
			p.SetState(148)
			p.InstanceDecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(151)
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
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(153)
			p.BlockDef()
		}

	case 2:
		{
			p.SetState(154)
			p.Statement()
		}
		{
			p.SetState(155)
			p.Match(DaedalusParserSemi)
		}

	case 3:
		{
			p.SetState(157)
			p.IfBlockStatement()
		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DaedalusParserSemi {
			{
				p.SetState(158)
				p.Match(DaedalusParserSemi)
			}

		}

	case 4:
		{
			p.SetState(161)
			p.WhileBlockStatement()
		}
		{
			p.SetState(162)
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
		p.SetState(166)
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
		p.SetState(168)
		p.Match(DaedalusParserMelse)
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4613480457911074304) != 0 {
		{
			p.SetState(169)
			p.MacroBlock()
		}

		p.SetState(174)
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
		p.SetState(175)
		p.Match(DaedalusParserMelif)
	}
	{
		p.SetState(176)
		p.MacroCondition()
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4613480457911074304) != 0 {
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
		p.SetState(183)
		p.Match(DaedalusParserMif)
	}
	{
		p.SetState(184)
		p.MacroCondition()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4613480457911074304) != 0 {
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
		p.SetState(191)
		p.MacroIfBlock()
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(192)
				p.MacroElseIfBlock()
			}

		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserMelse {
		{
			p.SetState(198)
			p.MacroElseBlock()
		}

	}
	{
		p.SetState(201)
		p.Match(DaedalusParserMendif)
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
	p.EnterRule(localctx, 18, DaedalusParserRULE_functionDef)

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
		p.SetState(203)
		p.Match(DaedalusParserFunc)
	}
	{
		p.SetState(204)
		p.TypeReference()
	}
	{
		p.SetState(205)
		p.NameNode()
	}
	{
		p.SetState(206)
		p.ParameterList()
	}
	{
		p.SetState(207)
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
	p.EnterRule(localctx, 20, DaedalusParserRULE_constDef)
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
		p.SetState(209)
		p.Match(DaedalusParserConst)
	}
	{
		p.SetState(210)
		p.TypeReference()
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(211)
			p.ConstValueDef()
		}

	case 2:
		{
			p.SetState(212)
			p.ConstArrayDef()
		}

	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DaedalusParserT__0 {
		{
			p.SetState(215)
			p.Match(DaedalusParserT__0)
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(216)
				p.ConstValueDef()
			}

		case 2:
			{
				p.SetState(217)
				p.ConstArrayDef()
			}

		}

		p.SetState(224)
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
	p.EnterRule(localctx, 22, DaedalusParserRULE_classDef)

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
		p.SetState(225)
		p.Match(DaedalusParserClass)
	}
	{
		p.SetState(226)
		p.NameNode()
	}
	{
		p.SetState(227)
		p.Match(DaedalusParserLeftBrace)
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(228)
				p.VarDecl()
			}
			{
				p.SetState(229)
				p.Match(DaedalusParserSemi)
			}

		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}
	{
		p.SetState(236)
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
	p.EnterRule(localctx, 24, DaedalusParserRULE_prototypeDef)

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
		p.SetState(238)
		p.Match(DaedalusParserPrototype)
	}
	{
		p.SetState(239)
		p.NameNode()
	}
	{
		p.SetState(240)
		p.Match(DaedalusParserLeftParen)
	}
	{
		p.SetState(241)
		p.ParentReference()
	}
	{
		p.SetState(242)
		p.Match(DaedalusParserRightParen)
	}
	{
		p.SetState(243)
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
	p.EnterRule(localctx, 26, DaedalusParserRULE_instanceDef)

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
		p.SetState(245)
		p.Match(DaedalusParserInstance)
	}
	{
		p.SetState(246)
		p.NameNode()
	}
	{
		p.SetState(247)
		p.Match(DaedalusParserLeftParen)
	}
	{
		p.SetState(248)
		p.ParentReference()
	}
	{
		p.SetState(249)
		p.Match(DaedalusParserRightParen)
	}
	{
		p.SetState(250)
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
	p.EnterRule(localctx, 28, DaedalusParserRULE_instanceDecl)

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
		p.Match(DaedalusParserInstance)
	}
	{
		p.SetState(253)
		p.NameNode()
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(254)
				p.Match(DaedalusParserT__0)
			}
			{
				p.SetState(255)
				p.NameNode()
			}

		}
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}
	{
		p.SetState(261)
		p.Match(DaedalusParserLeftParen)
	}
	{
		p.SetState(262)
		p.ParentReference()
	}
	{
		p.SetState(263)
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
	p.EnterRule(localctx, 30, DaedalusParserRULE_namespaceDef)

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
		p.SetState(265)
		p.Match(DaedalusParserNamespace)
	}
	{
		p.SetState(266)
		p.NameNode()
	}
	{
		p.SetState(267)
		p.Match(DaedalusParserLeftBrace)
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(268)
				p.ContentBlock()
			}

		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}
	{
		p.SetState(274)
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
	p.EnterRule(localctx, 32, DaedalusParserRULE_mainBlock)
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
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserMeta {
		{
			p.SetState(276)
			p.ZParserExtenderMetaBlock()
		}

	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2206871552) != 0 {
		{
			p.SetState(279)
			p.ContentBlock()
		}

		p.SetState(282)
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
	p.EnterRule(localctx, 34, DaedalusParserRULE_contentBlock)

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
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(284)
			p.BlockDef()
		}

	case 2:
		{
			p.SetState(285)
			p.InlineDef()
		}

	case 3:
		{
			p.SetState(286)
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
	p.EnterRule(localctx, 36, DaedalusParserRULE_varDecl)

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
		p.SetState(289)
		p.Match(DaedalusParserVar)
	}
	{
		p.SetState(290)
		p.TypeReference()
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(291)
			p.VarValueDecl()
		}

	case 2:
		{
			p.SetState(292)
			p.VarArrayDecl()
		}

	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(295)
				p.Match(DaedalusParserT__0)
			}
			p.SetState(299)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(296)
					p.VarDecl()
				}

			case 2:
				{
					p.SetState(297)
					p.VarValueDecl()
				}

			case 3:
				{
					p.SetState(298)
					p.VarArrayDecl()
				}

			}

		}
		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 38, DaedalusParserRULE_metaValue)

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
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	_alt = 1 + 1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1 + 1:
			p.SetState(306)
			p.MatchWildcard()

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 40, DaedalusParserRULE_zParserExtenderMeta)

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
		p.SetState(311)
		p.NameNode()
	}
	{
		p.SetState(312)
		p.Match(DaedalusParserAssign)
	}
	{
		p.SetState(313)
		p.MetaValue()
	}
	{
		p.SetState(314)
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
	p.EnterRule(localctx, 42, DaedalusParserRULE_zParserExtenderMetaBlock)

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
		p.Match(DaedalusParserMeta)
	}
	{
		p.SetState(317)
		p.Match(DaedalusParserLeftBrace)
	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(318)
				p.ZParserExtenderMeta()
			}

		}
		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}
	{
		p.SetState(324)
		p.Match(DaedalusParserRightBrace)
	}
	{
		p.SetState(325)
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
	p.EnterRule(localctx, 44, DaedalusParserRULE_constArrayDef)

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
		p.SetState(327)
		p.NameNode()
	}
	{
		p.SetState(328)
		p.Match(DaedalusParserLeftBracket)
	}
	{
		p.SetState(329)
		p.ArraySize()
	}
	{
		p.SetState(330)
		p.Match(DaedalusParserRightBracket)
	}
	{
		p.SetState(331)
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
	p.EnterRule(localctx, 46, DaedalusParserRULE_constArrayAssignment)

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
		p.SetState(333)
		p.Match(DaedalusParserAssign)
	}
	{
		p.SetState(334)
		p.Match(DaedalusParserLeftBrace)
	}

	{
		p.SetState(335)
		p.ExpressionBlock()
	}
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(336)
				p.Match(DaedalusParserT__0)
			}
			{
				p.SetState(337)
				p.ExpressionBlock()
			}

		}
		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}

	{
		p.SetState(343)
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
	p.EnterRule(localctx, 48, DaedalusParserRULE_constValueDef)

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
		p.SetState(345)
		p.NameNode()
	}
	{
		p.SetState(346)
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
	p.EnterRule(localctx, 50, DaedalusParserRULE_constValueAssignment)

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
		p.SetState(348)
		p.Match(DaedalusParserAssign)
	}
	{
		p.SetState(349)
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
	p.EnterRule(localctx, 52, DaedalusParserRULE_varArrayDecl)

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
		p.SetState(351)
		p.NameNode()
	}
	{
		p.SetState(352)
		p.Match(DaedalusParserLeftBracket)
	}
	{
		p.SetState(353)
		p.ArraySize()
	}
	{
		p.SetState(354)
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
	p.EnterRule(localctx, 54, DaedalusParserRULE_varValueDecl)

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
		p.SetState(356)
		p.NameNode()
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
	p.EnterRule(localctx, 56, DaedalusParserRULE_parameterList)
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
		p.SetState(358)
		p.Match(DaedalusParserLeftParen)
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserVar {
		{
			p.SetState(359)
			p.ParameterDecl()
		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

		for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1+1 {
				{
					p.SetState(360)
					p.Match(DaedalusParserT__0)
				}
				{
					p.SetState(361)
					p.ParameterDecl()
				}

			}
			p.SetState(366)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
		}

	}
	{
		p.SetState(369)
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
	p.EnterRule(localctx, 58, DaedalusParserRULE_parameterDecl)
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
		p.SetState(371)
		p.Match(DaedalusParserVar)
	}
	{
		p.SetState(372)
		p.TypeReference()
	}
	{
		p.SetState(373)
		p.NameNode()
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserLeftBracket {
		{
			p.SetState(374)
			p.Match(DaedalusParserLeftBracket)
		}
		{
			p.SetState(375)
			p.ArraySize()
		}
		{
			p.SetState(376)
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
	p.EnterRule(localctx, 60, DaedalusParserRULE_statementBlock)
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
		p.SetState(380)
		p.Match(DaedalusParserLeftBrace)
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(392)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case DaedalusParserIntegerLiteral, DaedalusParserFloatLiteral, DaedalusParserStringLiteral, DaedalusParserConst, DaedalusParserVar, DaedalusParserInt, DaedalusParserFunc, DaedalusParserStringKeyword, DaedalusParserClass, DaedalusParserVoid, DaedalusParserReturn, DaedalusParserFloat, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNamespace, DaedalusParserNull, DaedalusParserMeta, DaedalusParserContinue, DaedalusParserBreak, DaedalusParserLeftParen, DaedalusParserPlus, DaedalusParserMinus, DaedalusParserTilde, DaedalusParserNot, DaedalusParserIdentifier:
				{
					p.SetState(381)
					p.Statement()
				}
				{
					p.SetState(382)
					p.Match(DaedalusParserSemi)
				}

			case DaedalusParserIf:
				{
					p.SetState(384)
					p.IfBlockStatement()
				}
				p.SetState(386)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == DaedalusParserSemi {
					{
						p.SetState(385)
						p.Match(DaedalusParserSemi)
					}

				}

			case DaedalusParserMif:
				{
					p.SetState(388)
					p.MacroDef()
				}

			case DaedalusParserWhile:
				{
					p.SetState(389)
					p.WhileBlockStatement()
				}
				{
					p.SetState(390)
					p.Match(DaedalusParserSemi)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(396)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
	}
	{
		p.SetState(397)
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
	p.EnterRule(localctx, 62, DaedalusParserRULE_statement)

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

	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(399)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(400)
			p.ReturnStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(401)
			p.ConstDef()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(402)
			p.VarDecl()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(403)
			p.expression(0)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(404)
			p.Match(DaedalusParserContinue)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(405)
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
	p.EnterRule(localctx, 64, DaedalusParserRULE_funcCall)
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
		p.SetState(408)
		p.NameNode()
	}
	{
		p.SetState(409)
		p.Match(DaedalusParserLeftParen)
	}
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4613480456029908480) != 0 {
		{
			p.SetState(410)
			p.FuncArgExpression()
		}
		p.SetState(415)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

		for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1+1 {
				{
					p.SetState(411)
					p.Match(DaedalusParserT__0)
				}
				{
					p.SetState(412)
					p.FuncArgExpression()
				}

			}
			p.SetState(417)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
		}

	}
	{
		p.SetState(420)
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
	p.EnterRule(localctx, 66, DaedalusParserRULE_assignment)

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
		p.SetState(422)
		p.Reference()
	}
	{
		p.SetState(423)
		p.AssignmentOperator()
	}
	{
		p.SetState(424)
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
	p.EnterRule(localctx, 68, DaedalusParserRULE_ifCondition)

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
		p.SetState(426)
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
	p.EnterRule(localctx, 70, DaedalusParserRULE_elseBlock)

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
		p.SetState(428)
		p.Match(DaedalusParserElse)
	}
	{
		p.SetState(429)
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
	p.EnterRule(localctx, 72, DaedalusParserRULE_elseIfBlock)

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
		p.SetState(431)
		p.Match(DaedalusParserElse)
	}
	{
		p.SetState(432)
		p.Match(DaedalusParserIf)
	}
	{
		p.SetState(433)
		p.IfCondition()
	}
	{
		p.SetState(434)
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
	p.EnterRule(localctx, 74, DaedalusParserRULE_ifBlock)

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
		p.SetState(436)
		p.Match(DaedalusParserIf)
	}
	{
		p.SetState(437)
		p.IfCondition()
	}
	{
		p.SetState(438)
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
	p.EnterRule(localctx, 76, DaedalusParserRULE_ifBlockStatement)
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
		p.SetState(440)
		p.IfBlock()
	}
	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(441)
				p.ElseIfBlock()
			}

		}
		p.SetState(446)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}
	p.SetState(448)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserElse {
		{
			p.SetState(447)
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
	p.EnterRule(localctx, 78, DaedalusParserRULE_returnStatement)
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
		p.SetState(450)
		p.Match(DaedalusParserReturn)
	}
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4613480456029908480) != 0 {
		{
			p.SetState(451)
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
	p.EnterRule(localctx, 80, DaedalusParserRULE_whileCondition)

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
		p.SetState(454)
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
	p.EnterRule(localctx, 82, DaedalusParserRULE_whileBlock)

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
		p.SetState(456)
		p.Match(DaedalusParserWhile)
	}
	{
		p.SetState(457)
		p.WhileCondition()
	}
	{
		p.SetState(458)
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
	p.EnterRule(localctx, 84, DaedalusParserRULE_whileBlockStatement)

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
		p.SetState(460)
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
	p.EnterRule(localctx, 86, DaedalusParserRULE_funcArgExpression)

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
		p.SetState(462)
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
	p.EnterRule(localctx, 88, DaedalusParserRULE_expressionBlock)

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
	_startState := 90
	p.EnterRecursionRule(localctx, 90, DaedalusParserRULE_expression, _p)

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
	p.SetState(475)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserLeftParen:
		localctx = NewBracketExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(467)
			p.Match(DaedalusParserLeftParen)
		}
		{
			p.SetState(468)
			p.expression(0)
		}
		{
			p.SetState(469)
			p.Match(DaedalusParserRightParen)
		}

	case DaedalusParserPlus, DaedalusParserMinus, DaedalusParserTilde, DaedalusParserNot:
		localctx = NewUnaryOperationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(471)
			p.UnaryOperator()
		}
		{
			p.SetState(472)
			p.expression(11)
		}

	case DaedalusParserIntegerLiteral, DaedalusParserFloatLiteral, DaedalusParserStringLiteral, DaedalusParserVar, DaedalusParserInt, DaedalusParserFunc, DaedalusParserStringKeyword, DaedalusParserClass, DaedalusParserVoid, DaedalusParserFloat, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNamespace, DaedalusParserNull, DaedalusParserMeta, DaedalusParserIdentifier:
		localctx = NewValExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(474)
			p.Value()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(513)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(477)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(478)
					p.MultOperator()
				}
				{
					p.SetState(479)
					p.expression(11)
				}

			case 2:
				localctx = NewAddExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(481)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(482)
					p.AddOperator()
				}
				{
					p.SetState(483)
					p.expression(10)
				}

			case 3:
				localctx = NewBitMoveExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(485)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(486)
					p.BitMoveOperator()
				}
				{
					p.SetState(487)
					p.expression(9)
				}

			case 4:
				localctx = NewCompExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(489)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(490)
					p.CompOperator()
				}
				{
					p.SetState(491)
					p.expression(8)
				}

			case 5:
				localctx = NewEqExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(493)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(494)
					p.EqOperator()
				}
				{
					p.SetState(495)
					p.expression(7)
				}

			case 6:
				localctx = NewBinAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(497)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(498)
					p.BinAndOperator()
				}
				{
					p.SetState(499)
					p.expression(6)
				}

			case 7:
				localctx = NewBinOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(501)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(502)
					p.BinOrOperator()
				}
				{
					p.SetState(503)
					p.expression(5)
				}

			case 8:
				localctx = NewLogAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(505)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(506)
					p.LogAndOperator()
				}
				{
					p.SetState(507)
					p.expression(4)
				}

			case 9:
				localctx = NewLogOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(509)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(510)
					p.LogOrOperator()
				}
				{
					p.SetState(511)
					p.expression(3)
				}

			}

		}
		p.SetState(517)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 92, DaedalusParserRULE_arrayIndex)

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

	p.SetState(520)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(518)
			p.Match(DaedalusParserIntegerLiteral)
		}

	case DaedalusParserVar, DaedalusParserInt, DaedalusParserFunc, DaedalusParserStringKeyword, DaedalusParserClass, DaedalusParserVoid, DaedalusParserFloat, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNamespace, DaedalusParserNull, DaedalusParserMeta, DaedalusParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(519)
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
	p.EnterRule(localctx, 94, DaedalusParserRULE_arraySize)

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

	p.SetState(524)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(522)
			p.Match(DaedalusParserIntegerLiteral)
		}

	case DaedalusParserVar, DaedalusParserInt, DaedalusParserFunc, DaedalusParserStringKeyword, DaedalusParserClass, DaedalusParserVoid, DaedalusParserFloat, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNamespace, DaedalusParserNull, DaedalusParserMeta, DaedalusParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(523)
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
	p.EnterRule(localctx, 96, DaedalusParserRULE_value)

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

	p.SetState(532)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIntegerLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(526)
			p.Match(DaedalusParserIntegerLiteral)
		}

	case 2:
		localctx = NewFloatLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(527)
			p.Match(DaedalusParserFloatLiteral)
		}

	case 3:
		localctx = NewStringLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(528)
			p.Match(DaedalusParserStringLiteral)
		}

	case 4:
		localctx = NewNullLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(529)
			p.Match(DaedalusParserNull)
		}

	case 5:
		localctx = NewFuncCallValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(530)
			p.FuncCall()
		}

	case 6:
		localctx = NewReferenceValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(531)
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
	p.EnterRule(localctx, 98, DaedalusParserRULE_referenceAtom)

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
		p.SetState(534)
		p.NameNode()
	}
	p.SetState(539)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(535)
			p.Match(DaedalusParserLeftBracket)
		}
		{
			p.SetState(536)
			p.ArrayIndex()
		}
		{
			p.SetState(537)
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
	p.EnterRule(localctx, 100, DaedalusParserRULE_reference)

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
		p.SetState(541)
		p.ReferenceAtom()
	}
	p.SetState(544)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(542)
			p.Match(DaedalusParserDot)
		}
		{
			p.SetState(543)
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
	p.EnterRule(localctx, 102, DaedalusParserRULE_typeReference)
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
		p.SetState(546)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4611686018449833984) != 0) {
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
	p.EnterRule(localctx, 104, DaedalusParserRULE_anyIdentifier)
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
		p.SetState(548)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4611686018693636096) != 0) {
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
	p.EnterRule(localctx, 106, DaedalusParserRULE_nameNode)

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
		p.SetState(550)
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
	p.EnterRule(localctx, 108, DaedalusParserRULE_parentReference)

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
		p.SetState(552)
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
	p.EnterRule(localctx, 110, DaedalusParserRULE_assignmentOperator)
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
		p.SetState(554)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1137158905911050240) != 0) {
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
	p.EnterRule(localctx, 112, DaedalusParserRULE_unaryOperator)
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
		p.SetState(556)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1794402976530432) != 0) {
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
	p.EnterRule(localctx, 114, DaedalusParserRULE_addOperator)
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
		p.SetState(558)
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
	p.EnterRule(localctx, 116, DaedalusParserRULE_bitMoveOperator)
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
		p.SetState(560)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DaedalusParserT__1 || _la == DaedalusParserT__2) {
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
	p.EnterRule(localctx, 118, DaedalusParserRULE_compOperator)
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
		p.SetState(562)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13510798882111536) != 0) {
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
	p.EnterRule(localctx, 120, DaedalusParserRULE_eqOperator)
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
		p.SetState(564)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DaedalusParserT__5 || _la == DaedalusParserT__6) {
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
	p.EnterRule(localctx, 122, DaedalusParserRULE_multOperator)
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
		p.SetState(566)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&422212465066240) != 0) {
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
	p.EnterRule(localctx, 124, DaedalusParserRULE_binAndOperator)

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
		p.SetState(568)
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
	p.EnterRule(localctx, 126, DaedalusParserRULE_binOrOperator)

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
		p.SetState(570)
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
	p.EnterRule(localctx, 128, DaedalusParserRULE_logAndOperator)

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
		p.SetState(572)
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
	p.EnterRule(localctx, 130, DaedalusParserRULE_logOrOperator)

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
		p.Match(DaedalusParserOr)
	}

	return localctx
}

func (p *DaedalusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 45:
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
