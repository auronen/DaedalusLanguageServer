// Code generated from Daedalus.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Daedalus

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 60, 496,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 3, 2,
	6, 2, 110, 10, 2, 13, 2, 14, 2, 111, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	7, 4, 120, 10, 4, 12, 4, 14, 4, 123, 11, 4, 3, 4, 3, 4, 3, 5, 7, 5, 128,
	10, 5, 12, 5, 14, 5, 131, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 137, 10,
	5, 3, 5, 5, 5, 140, 10, 5, 3, 6, 7, 6, 143, 10, 6, 12, 6, 14, 6, 146, 11,
	6, 3, 6, 3, 6, 3, 6, 5, 6, 151, 10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 165, 10, 8, 3, 8, 3, 8, 3,
	8, 5, 8, 170, 10, 8, 7, 8, 172, 10, 8, 12, 8, 14, 8, 175, 11, 8, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 184, 10, 9, 12, 9, 14, 9, 187,
	11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12,
	7, 12, 209, 10, 12, 12, 12, 14, 12, 212, 11, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 222, 10, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 5, 13, 228, 10, 13, 7, 13, 230, 10, 13, 12, 13, 14, 13, 233, 11,
	13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 7, 15, 246, 10, 15, 12, 15, 14, 15, 249, 11, 15, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 270, 10, 20, 12, 20, 14,
	20, 273, 11, 20, 5, 20, 275, 10, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 286, 10, 21, 3, 22, 7, 22, 289, 10,
	22, 12, 22, 14, 22, 292, 11, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 298,
	10, 22, 3, 22, 7, 22, 301, 10, 22, 12, 22, 14, 22, 304, 11, 22, 3, 22,
	3, 22, 7, 22, 308, 10, 22, 12, 22, 14, 22, 311, 11, 22, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 5, 23, 329, 10, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 7, 24, 336, 10, 24, 12, 24, 14, 24, 339, 11, 24, 5, 24, 341, 10, 24,
	3, 24, 3, 24, 3, 24, 5, 24, 346, 10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 7, 30, 368, 10, 30, 12, 30, 14, 30,
	371, 11, 30, 3, 30, 5, 30, 374, 10, 30, 3, 31, 3, 31, 5, 31, 378, 10, 31,
	3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 5, 34, 393, 10, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 7,
	34, 431, 10, 34, 12, 34, 14, 34, 434, 11, 34, 3, 35, 3, 35, 5, 35, 438,
	10, 35, 3, 36, 3, 36, 5, 36, 442, 10, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 5, 37, 452, 10, 37, 3, 38, 3, 38, 3, 38, 3, 38,
	3, 38, 5, 38, 459, 10, 38, 3, 39, 3, 39, 3, 39, 5, 39, 464, 10, 39, 3,
	40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45,
	3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3,
	50, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 10,
	121, 185, 210, 247, 271, 302, 337, 369, 3, 66, 55, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
	50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84,
	86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 2, 11, 8, 2, 39, 39, 41,
	42, 44, 44, 46, 46, 48, 48, 51, 51, 7, 2, 37, 37, 39, 39, 41, 44, 46, 49,
	51, 51, 4, 2, 11, 11, 13, 16, 3, 2, 17, 18, 3, 2, 19, 20, 3, 2, 21, 24,
	3, 2, 25, 26, 4, 2, 17, 18, 27, 28, 3, 2, 29, 31, 2, 508, 2, 109, 3, 2,
	2, 2, 4, 113, 3, 2, 2, 2, 6, 121, 3, 2, 2, 2, 8, 129, 3, 2, 2, 2, 10, 144,
	3, 2, 2, 2, 12, 154, 3, 2, 2, 2, 14, 160, 3, 2, 2, 2, 16, 176, 3, 2, 2,
	2, 18, 190, 3, 2, 2, 2, 20, 197, 3, 2, 2, 2, 22, 204, 3, 2, 2, 2, 24, 217,
	3, 2, 2, 2, 26, 234, 3, 2, 2, 2, 28, 240, 3, 2, 2, 2, 30, 252, 3, 2, 2,
	2, 32, 255, 3, 2, 2, 2, 34, 258, 3, 2, 2, 2, 36, 263, 3, 2, 2, 2, 38, 265,
	3, 2, 2, 2, 40, 278, 3, 2, 2, 2, 42, 290, 3, 2, 2, 2, 44, 328, 3, 2, 2,
	2, 46, 330, 3, 2, 2, 2, 48, 347, 3, 2, 2, 2, 50, 351, 3, 2, 2, 2, 52, 353,
	3, 2, 2, 2, 54, 356, 3, 2, 2, 2, 56, 361, 3, 2, 2, 2, 58, 365, 3, 2, 2,
	2, 60, 375, 3, 2, 2, 2, 62, 379, 3, 2, 2, 2, 64, 381, 3, 2, 2, 2, 66, 392,
	3, 2, 2, 2, 68, 437, 3, 2, 2, 2, 70, 441, 3, 2, 2, 2, 72, 451, 3, 2, 2,
	2, 74, 453, 3, 2, 2, 2, 76, 460, 3, 2, 2, 2, 78, 465, 3, 2, 2, 2, 80, 467,
	3, 2, 2, 2, 82, 469, 3, 2, 2, 2, 84, 471, 3, 2, 2, 2, 86, 473, 3, 2, 2,
	2, 88, 475, 3, 2, 2, 2, 90, 477, 3, 2, 2, 2, 92, 479, 3, 2, 2, 2, 94, 481,
	3, 2, 2, 2, 96, 483, 3, 2, 2, 2, 98, 485, 3, 2, 2, 2, 100, 487, 3, 2, 2,
	2, 102, 489, 3, 2, 2, 2, 104, 491, 3, 2, 2, 2, 106, 493, 3, 2, 2, 2, 108,
	110, 7, 57, 2, 2, 109, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 109,
	3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 3, 3, 2, 2, 2, 113, 114, 7, 60,
	2, 2, 114, 5, 3, 2, 2, 2, 115, 120, 5, 8, 5, 2, 116, 120, 5, 10, 6, 2,
	117, 120, 5, 2, 2, 2, 118, 120, 5, 4, 3, 2, 119, 115, 3, 2, 2, 2, 119,
	116, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 118, 3, 2, 2, 2, 120, 123,
	3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 124, 3, 2,
	2, 2, 123, 121, 3, 2, 2, 2, 124, 125, 7, 2, 2, 3, 125, 7, 3, 2, 2, 2, 126,
	128, 5, 2, 2, 2, 127, 126, 3, 2, 2, 2, 128, 131, 3, 2, 2, 2, 129, 127,
	3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 136, 3, 2, 2, 2, 131, 129, 3, 2,
	2, 2, 132, 137, 5, 12, 7, 2, 133, 137, 5, 16, 9, 2, 134, 137, 5, 18, 10,
	2, 135, 137, 5, 20, 11, 2, 136, 132, 3, 2, 2, 2, 136, 133, 3, 2, 2, 2,
	136, 134, 3, 2, 2, 2, 136, 135, 3, 2, 2, 2, 137, 139, 3, 2, 2, 2, 138,
	140, 7, 3, 2, 2, 139, 138, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 9, 3,
	2, 2, 2, 141, 143, 5, 2, 2, 2, 142, 141, 3, 2, 2, 2, 143, 146, 3, 2, 2,
	2, 144, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 150, 3, 2, 2, 2, 146,
	144, 3, 2, 2, 2, 147, 151, 5, 14, 8, 2, 148, 151, 5, 24, 13, 2, 149, 151,
	5, 22, 12, 2, 150, 147, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 150, 149, 3,
	2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 153, 7, 3, 2, 2, 153, 11, 3, 2, 2,
	2, 154, 155, 7, 41, 2, 2, 155, 156, 5, 78, 40, 2, 156, 157, 5, 82, 42,
	2, 157, 158, 5, 38, 20, 2, 158, 159, 5, 42, 22, 2, 159, 13, 3, 2, 2, 2,
	160, 161, 7, 36, 2, 2, 161, 164, 5, 78, 40, 2, 162, 165, 5, 30, 16, 2,
	163, 165, 5, 26, 14, 2, 164, 162, 3, 2, 2, 2, 164, 163, 3, 2, 2, 2, 165,
	173, 3, 2, 2, 2, 166, 169, 7, 4, 2, 2, 167, 170, 5, 30, 16, 2, 168, 170,
	5, 26, 14, 2, 169, 167, 3, 2, 2, 2, 169, 168, 3, 2, 2, 2, 170, 172, 3,
	2, 2, 2, 171, 166, 3, 2, 2, 2, 172, 175, 3, 2, 2, 2, 173, 171, 3, 2, 2,
	2, 173, 174, 3, 2, 2, 2, 174, 15, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 176,
	177, 7, 43, 2, 2, 177, 178, 5, 82, 42, 2, 178, 185, 7, 5, 2, 2, 179, 180,
	5, 24, 13, 2, 180, 181, 7, 3, 2, 2, 181, 184, 3, 2, 2, 2, 182, 184, 5,
	2, 2, 2, 183, 179, 3, 2, 2, 2, 183, 182, 3, 2, 2, 2, 184, 187, 3, 2, 2,
	2, 185, 186, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 186, 188, 3, 2, 2, 2, 187,
	185, 3, 2, 2, 2, 188, 189, 7, 6, 2, 2, 189, 17, 3, 2, 2, 2, 190, 191, 7,
	47, 2, 2, 191, 192, 5, 82, 42, 2, 192, 193, 7, 7, 2, 2, 193, 194, 5, 84,
	43, 2, 194, 195, 7, 8, 2, 2, 195, 196, 5, 42, 22, 2, 196, 19, 3, 2, 2,
	2, 197, 198, 7, 48, 2, 2, 198, 199, 5, 82, 42, 2, 199, 200, 7, 7, 2, 2,
	200, 201, 5, 84, 43, 2, 201, 202, 7, 8, 2, 2, 202, 203, 5, 42, 22, 2, 203,
	21, 3, 2, 2, 2, 204, 205, 7, 48, 2, 2, 205, 210, 5, 82, 42, 2, 206, 207,
	7, 4, 2, 2, 207, 209, 5, 82, 42, 2, 208, 206, 3, 2, 2, 2, 209, 212, 3,
	2, 2, 2, 210, 211, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 211, 213, 3, 2, 2,
	2, 212, 210, 3, 2, 2, 2, 213, 214, 7, 7, 2, 2, 214, 215, 5, 84, 43, 2,
	215, 216, 7, 8, 2, 2, 216, 23, 3, 2, 2, 2, 217, 218, 7, 37, 2, 2, 218,
	221, 5, 78, 40, 2, 219, 222, 5, 36, 19, 2, 220, 222, 5, 34, 18, 2, 221,
	219, 3, 2, 2, 2, 221, 220, 3, 2, 2, 2, 222, 231, 3, 2, 2, 2, 223, 227,
	7, 4, 2, 2, 224, 228, 5, 24, 13, 2, 225, 228, 5, 36, 19, 2, 226, 228, 5,
	34, 18, 2, 227, 224, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 227, 226, 3, 2,
	2, 2, 228, 230, 3, 2, 2, 2, 229, 223, 3, 2, 2, 2, 230, 233, 3, 2, 2, 2,
	231, 229, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232, 25, 3, 2, 2, 2, 233, 231,
	3, 2, 2, 2, 234, 235, 5, 82, 42, 2, 235, 236, 7, 9, 2, 2, 236, 237, 5,
	70, 36, 2, 237, 238, 7, 10, 2, 2, 238, 239, 5, 28, 15, 2, 239, 27, 3, 2,
	2, 2, 240, 241, 7, 11, 2, 2, 241, 242, 7, 5, 2, 2, 242, 247, 5, 64, 33,
	2, 243, 244, 7, 4, 2, 2, 244, 246, 5, 64, 33, 2, 245, 243, 3, 2, 2, 2,
	246, 249, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 248,
	250, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 250, 251, 7, 6, 2, 2, 251, 29, 3,
	2, 2, 2, 252, 253, 5, 82, 42, 2, 253, 254, 5, 32, 17, 2, 254, 31, 3, 2,
	2, 2, 255, 256, 7, 11, 2, 2, 256, 257, 5, 64, 33, 2, 257, 33, 3, 2, 2,
	2, 258, 259, 5, 82, 42, 2, 259, 260, 7, 9, 2, 2, 260, 261, 5, 70, 36, 2,
	261, 262, 7, 10, 2, 2, 262, 35, 3, 2, 2, 2, 263, 264, 5, 82, 42, 2, 264,
	37, 3, 2, 2, 2, 265, 274, 7, 7, 2, 2, 266, 271, 5, 40, 21, 2, 267, 268,
	7, 4, 2, 2, 268, 270, 5, 40, 21, 2, 269, 267, 3, 2, 2, 2, 270, 273, 3,
	2, 2, 2, 271, 272, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 272, 275, 3, 2, 2,
	2, 273, 271, 3, 2, 2, 2, 274, 266, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275,
	276, 3, 2, 2, 2, 276, 277, 7, 8, 2, 2, 277, 39, 3, 2, 2, 2, 278, 279, 7,
	37, 2, 2, 279, 280, 5, 78, 40, 2, 280, 285, 5, 82, 42, 2, 281, 282, 7,
	9, 2, 2, 282, 283, 5, 70, 36, 2, 283, 284, 7, 10, 2, 2, 284, 286, 3, 2,
	2, 2, 285, 281, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 41, 3, 2, 2, 2,
	287, 289, 5, 2, 2, 2, 288, 287, 3, 2, 2, 2, 289, 292, 3, 2, 2, 2, 290,
	288, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 293, 3, 2, 2, 2, 292, 290,
	3, 2, 2, 2, 293, 302, 7, 5, 2, 2, 294, 301, 5, 44, 23, 2, 295, 297, 5,
	58, 30, 2, 296, 298, 7, 3, 2, 2, 297, 296, 3, 2, 2, 2, 297, 298, 3, 2,
	2, 2, 298, 301, 3, 2, 2, 2, 299, 301, 5, 2, 2, 2, 300, 294, 3, 2, 2, 2,
	300, 295, 3, 2, 2, 2, 300, 299, 3, 2, 2, 2, 301, 304, 3, 2, 2, 2, 302,
	303, 3, 2, 2, 2, 302, 300, 3, 2, 2, 2, 303, 305, 3, 2, 2, 2, 304, 302,
	3, 2, 2, 2, 305, 309, 7, 6, 2, 2, 306, 308, 5, 2, 2, 2, 307, 306, 3, 2,
	2, 2, 308, 311, 3, 2, 2, 2, 309, 307, 3, 2, 2, 2, 309, 310, 3, 2, 2, 2,
	310, 43, 3, 2, 2, 2, 311, 309, 3, 2, 2, 2, 312, 313, 5, 48, 25, 2, 313,
	314, 7, 3, 2, 2, 314, 329, 3, 2, 2, 2, 315, 316, 5, 60, 31, 2, 316, 317,
	7, 3, 2, 2, 317, 329, 3, 2, 2, 2, 318, 319, 5, 14, 8, 2, 319, 320, 7, 3,
	2, 2, 320, 329, 3, 2, 2, 2, 321, 322, 5, 24, 13, 2, 322, 323, 7, 3, 2,
	2, 323, 329, 3, 2, 2, 2, 324, 329, 5, 46, 24, 2, 325, 326, 5, 66, 34, 2,
	326, 327, 7, 3, 2, 2, 327, 329, 3, 2, 2, 2, 328, 312, 3, 2, 2, 2, 328,
	315, 3, 2, 2, 2, 328, 318, 3, 2, 2, 2, 328, 321, 3, 2, 2, 2, 328, 324,
	3, 2, 2, 2, 328, 325, 3, 2, 2, 2, 329, 45, 3, 2, 2, 2, 330, 331, 5, 82,
	42, 2, 331, 340, 7, 7, 2, 2, 332, 337, 5, 62, 32, 2, 333, 334, 7, 4, 2,
	2, 334, 336, 5, 62, 32, 2, 335, 333, 3, 2, 2, 2, 336, 339, 3, 2, 2, 2,
	337, 338, 3, 2, 2, 2, 337, 335, 3, 2, 2, 2, 338, 341, 3, 2, 2, 2, 339,
	337, 3, 2, 2, 2, 340, 332, 3, 2, 2, 2, 340, 341, 3, 2, 2, 2, 341, 342,
	3, 2, 2, 2, 342, 343, 7, 8, 2, 2, 343, 345, 7, 3, 2, 2, 344, 346, 5, 4,
	3, 2, 345, 344, 3, 2, 2, 2, 345, 346, 3, 2, 2, 2, 346, 47, 3, 2, 2, 2,
	347, 348, 5, 76, 39, 2, 348, 349, 5, 86, 44, 2, 349, 350, 5, 64, 33, 2,
	350, 49, 3, 2, 2, 2, 351, 352, 5, 64, 33, 2, 352, 51, 3, 2, 2, 2, 353,
	354, 7, 40, 2, 2, 354, 355, 5, 42, 22, 2, 355, 53, 3, 2, 2, 2, 356, 357,
	7, 40, 2, 2, 357, 358, 7, 38, 2, 2, 358, 359, 5, 50, 26, 2, 359, 360, 5,
	42, 22, 2, 360, 55, 3, 2, 2, 2, 361, 362, 7, 38, 2, 2, 362, 363, 5, 50,
	26, 2, 363, 364, 5, 42, 22, 2, 364, 57, 3, 2, 2, 2, 365, 369, 5, 56, 29,
	2, 366, 368, 5, 54, 28, 2, 367, 366, 3, 2, 2, 2, 368, 371, 3, 2, 2, 2,
	369, 370, 3, 2, 2, 2, 369, 367, 3, 2, 2, 2, 370, 373, 3, 2, 2, 2, 371,
	369, 3, 2, 2, 2, 372, 374, 5, 52, 27, 2, 373, 372, 3, 2, 2, 2, 373, 374,
	3, 2, 2, 2, 374, 59, 3, 2, 2, 2, 375, 377, 7, 45, 2, 2, 376, 378, 5, 64,
	33, 2, 377, 376, 3, 2, 2, 2, 377, 378, 3, 2, 2, 2, 378, 61, 3, 2, 2, 2,
	379, 380, 5, 64, 33, 2, 380, 63, 3, 2, 2, 2, 381, 382, 5, 66, 34, 2, 382,
	65, 3, 2, 2, 2, 383, 384, 8, 34, 1, 2, 384, 385, 7, 7, 2, 2, 385, 386,
	5, 66, 34, 2, 386, 387, 7, 8, 2, 2, 387, 393, 3, 2, 2, 2, 388, 389, 5,
	96, 49, 2, 389, 390, 5, 66, 34, 13, 390, 393, 3, 2, 2, 2, 391, 393, 5,
	72, 37, 2, 392, 383, 3, 2, 2, 2, 392, 388, 3, 2, 2, 2, 392, 391, 3, 2,
	2, 2, 393, 432, 3, 2, 2, 2, 394, 395, 12, 12, 2, 2, 395, 396, 5, 98, 50,
	2, 396, 397, 5, 66, 34, 13, 397, 431, 3, 2, 2, 2, 398, 399, 12, 11, 2,
	2, 399, 400, 5, 88, 45, 2, 400, 401, 5, 66, 34, 12, 401, 431, 3, 2, 2,
	2, 402, 403, 12, 10, 2, 2, 403, 404, 5, 90, 46, 2, 404, 405, 5, 66, 34,
	11, 405, 431, 3, 2, 2, 2, 406, 407, 12, 9, 2, 2, 407, 408, 5, 92, 47, 2,
	408, 409, 5, 66, 34, 10, 409, 431, 3, 2, 2, 2, 410, 411, 12, 8, 2, 2, 411,
	412, 5, 94, 48, 2, 412, 413, 5, 66, 34, 9, 413, 431, 3, 2, 2, 2, 414, 415,
	12, 7, 2, 2, 415, 416, 5, 100, 51, 2, 416, 417, 5, 66, 34, 8, 417, 431,
	3, 2, 2, 2, 418, 419, 12, 6, 2, 2, 419, 420, 5, 102, 52, 2, 420, 421, 5,
	66, 34, 7, 421, 431, 3, 2, 2, 2, 422, 423, 12, 5, 2, 2, 423, 424, 5, 104,
	53, 2, 424, 425, 5, 66, 34, 6, 425, 431, 3, 2, 2, 2, 426, 427, 12, 4, 2,
	2, 427, 428, 5, 106, 54, 2, 428, 429, 5, 66, 34, 5, 429, 431, 3, 2, 2,
	2, 430, 394, 3, 2, 2, 2, 430, 398, 3, 2, 2, 2, 430, 402, 3, 2, 2, 2, 430,
	406, 3, 2, 2, 2, 430, 410, 3, 2, 2, 2, 430, 414, 3, 2, 2, 2, 430, 418,
	3, 2, 2, 2, 430, 422, 3, 2, 2, 2, 430, 426, 3, 2, 2, 2, 431, 434, 3, 2,
	2, 2, 432, 430, 3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 67, 3, 2, 2, 2,
	434, 432, 3, 2, 2, 2, 435, 438, 7, 52, 2, 2, 436, 438, 5, 74, 38, 2, 437,
	435, 3, 2, 2, 2, 437, 436, 3, 2, 2, 2, 438, 69, 3, 2, 2, 2, 439, 442, 7,
	52, 2, 2, 440, 442, 5, 74, 38, 2, 441, 439, 3, 2, 2, 2, 441, 440, 3, 2,
	2, 2, 442, 71, 3, 2, 2, 2, 443, 452, 7, 52, 2, 2, 444, 452, 7, 53, 2, 2,
	445, 452, 7, 54, 2, 2, 446, 452, 7, 49, 2, 2, 447, 452, 7, 50, 2, 2, 448,
	452, 5, 46, 24, 2, 449, 452, 5, 76, 39, 2, 450, 452, 5, 82, 42, 2, 451,
	443, 3, 2, 2, 2, 451, 444, 3, 2, 2, 2, 451, 445, 3, 2, 2, 2, 451, 446,
	3, 2, 2, 2, 451, 447, 3, 2, 2, 2, 451, 448, 3, 2, 2, 2, 451, 449, 3, 2,
	2, 2, 451, 450, 3, 2, 2, 2, 452, 73, 3, 2, 2, 2, 453, 458, 5, 82, 42, 2,
	454, 455, 7, 9, 2, 2, 455, 456, 5, 68, 35, 2, 456, 457, 7, 10, 2, 2, 457,
	459, 3, 2, 2, 2, 458, 454, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 75, 3,
	2, 2, 2, 460, 463, 5, 74, 38, 2, 461, 462, 7, 12, 2, 2, 462, 464, 5, 74,
	38, 2, 463, 461, 3, 2, 2, 2, 463, 464, 3, 2, 2, 2, 464, 77, 3, 2, 2, 2,
	465, 466, 9, 2, 2, 2, 466, 79, 3, 2, 2, 2, 467, 468, 9, 3, 2, 2, 468, 81,
	3, 2, 2, 2, 469, 470, 5, 80, 41, 2, 470, 83, 3, 2, 2, 2, 471, 472, 7, 51,
	2, 2, 472, 85, 3, 2, 2, 2, 473, 474, 9, 4, 2, 2, 474, 87, 3, 2, 2, 2, 475,
	476, 9, 5, 2, 2, 476, 89, 3, 2, 2, 2, 477, 478, 9, 6, 2, 2, 478, 91, 3,
	2, 2, 2, 479, 480, 9, 7, 2, 2, 480, 93, 3, 2, 2, 2, 481, 482, 9, 8, 2,
	2, 482, 95, 3, 2, 2, 2, 483, 484, 9, 9, 2, 2, 484, 97, 3, 2, 2, 2, 485,
	486, 9, 10, 2, 2, 486, 99, 3, 2, 2, 2, 487, 488, 7, 32, 2, 2, 488, 101,
	3, 2, 2, 2, 489, 490, 7, 33, 2, 2, 490, 103, 3, 2, 2, 2, 491, 492, 7, 34,
	2, 2, 492, 105, 3, 2, 2, 2, 493, 494, 7, 35, 2, 2, 494, 107, 3, 2, 2, 2,
	43, 111, 119, 121, 129, 136, 139, 144, 150, 164, 169, 173, 183, 185, 210,
	221, 227, 231, 247, 271, 274, 285, 290, 297, 300, 302, 309, 328, 337, 340,
	345, 369, 373, 377, 392, 430, 432, 437, 441, 451, 458, 463,
}
var literalNames = []string{
	"", "';'", "','", "'{'", "'}'", "'('", "')'", "'['", "']'", "'='", "'.'",
	"'+='", "'-='", "'*='", "'/='", "'+'", "'-'", "'<<'", "'>>'", "'<'", "'>'",
	"'<='", "'>='", "'=='", "'!='", "'!'", "'~'", "'*'", "'/'", "'%'", "'&'",
	"'|'", "'&&'", "'||'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Const",
	"Var", "If", "Int", "Else", "Func", "StringKeyword", "Class", "Void", "Return",
	"Float", "Prototype", "Instance", "Null", "NoFunc", "Identifier", "IntegerLiteral",
	"FloatLiteral", "StringLiteral", "Whitespace", "TooManyCommentlines", "SummaryComment",
	"Newline", "BlockComment", "LineComment",
}

var ruleNames = []string{
	"symbolSummary", "lineComment", "daedalusFile", "blockDef", "inlineDef",
	"functionDef", "constDef", "classDef", "prototypeDef", "instanceDef", "instanceDecl",
	"varDecl", "constArrayDef", "constArrayAssignment", "constValueDef", "constValueAssignment",
	"varArrayDecl", "varValueDecl", "parameterList", "parameterDecl", "statementBlock",
	"statement", "funcCall", "assignment", "ifCondition", "elseBlock", "elseIfBlock",
	"ifBlock", "ifBlockStatement", "returnStatement", "funcArgExpression",
	"expressionBlock", "expression", "arrayIndex", "arraySize", "value", "referenceAtom",
	"reference", "typeReference", "anyIdentifier", "nameNode", "parentReference",
	"assignmentOperator", "addOperator", "bitMoveOperator", "compOperator",
	"eqOperator", "oneArgOperator", "multOperator", "binAndOperator", "binOrOperator",
	"logAndOperator", "logOrOperator",
}

type DaedalusParser struct {
	*antlr.BaseParser
}

// NewDaedalusParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *DaedalusParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDaedalusParser(input antlr.TokenStream) *DaedalusParser {
	this := new(DaedalusParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Daedalus.g4"

	return this
}

// DaedalusParser tokens.
const (
	DaedalusParserEOF                 = antlr.TokenEOF
	DaedalusParserT__0                = 1
	DaedalusParserT__1                = 2
	DaedalusParserT__2                = 3
	DaedalusParserT__3                = 4
	DaedalusParserT__4                = 5
	DaedalusParserT__5                = 6
	DaedalusParserT__6                = 7
	DaedalusParserT__7                = 8
	DaedalusParserT__8                = 9
	DaedalusParserT__9                = 10
	DaedalusParserT__10               = 11
	DaedalusParserT__11               = 12
	DaedalusParserT__12               = 13
	DaedalusParserT__13               = 14
	DaedalusParserT__14               = 15
	DaedalusParserT__15               = 16
	DaedalusParserT__16               = 17
	DaedalusParserT__17               = 18
	DaedalusParserT__18               = 19
	DaedalusParserT__19               = 20
	DaedalusParserT__20               = 21
	DaedalusParserT__21               = 22
	DaedalusParserT__22               = 23
	DaedalusParserT__23               = 24
	DaedalusParserT__24               = 25
	DaedalusParserT__25               = 26
	DaedalusParserT__26               = 27
	DaedalusParserT__27               = 28
	DaedalusParserT__28               = 29
	DaedalusParserT__29               = 30
	DaedalusParserT__30               = 31
	DaedalusParserT__31               = 32
	DaedalusParserT__32               = 33
	DaedalusParserConst               = 34
	DaedalusParserVar                 = 35
	DaedalusParserIf                  = 36
	DaedalusParserInt                 = 37
	DaedalusParserElse                = 38
	DaedalusParserFunc                = 39
	DaedalusParserStringKeyword       = 40
	DaedalusParserClass               = 41
	DaedalusParserVoid                = 42
	DaedalusParserReturn              = 43
	DaedalusParserFloat               = 44
	DaedalusParserPrototype           = 45
	DaedalusParserInstance            = 46
	DaedalusParserNull                = 47
	DaedalusParserNoFunc              = 48
	DaedalusParserIdentifier          = 49
	DaedalusParserIntegerLiteral      = 50
	DaedalusParserFloatLiteral        = 51
	DaedalusParserStringLiteral       = 52
	DaedalusParserWhitespace          = 53
	DaedalusParserTooManyCommentlines = 54
	DaedalusParserSummaryComment      = 55
	DaedalusParserNewline             = 56
	DaedalusParserBlockComment        = 57
	DaedalusParserLineComment         = 58
)

// DaedalusParser rules.
const (
	DaedalusParserRULE_symbolSummary        = 0
	DaedalusParserRULE_lineComment          = 1
	DaedalusParserRULE_daedalusFile         = 2
	DaedalusParserRULE_blockDef             = 3
	DaedalusParserRULE_inlineDef            = 4
	DaedalusParserRULE_functionDef          = 5
	DaedalusParserRULE_constDef             = 6
	DaedalusParserRULE_classDef             = 7
	DaedalusParserRULE_prototypeDef         = 8
	DaedalusParserRULE_instanceDef          = 9
	DaedalusParserRULE_instanceDecl         = 10
	DaedalusParserRULE_varDecl              = 11
	DaedalusParserRULE_constArrayDef        = 12
	DaedalusParserRULE_constArrayAssignment = 13
	DaedalusParserRULE_constValueDef        = 14
	DaedalusParserRULE_constValueAssignment = 15
	DaedalusParserRULE_varArrayDecl         = 16
	DaedalusParserRULE_varValueDecl         = 17
	DaedalusParserRULE_parameterList        = 18
	DaedalusParserRULE_parameterDecl        = 19
	DaedalusParserRULE_statementBlock       = 20
	DaedalusParserRULE_statement            = 21
	DaedalusParserRULE_funcCall             = 22
	DaedalusParserRULE_assignment           = 23
	DaedalusParserRULE_ifCondition          = 24
	DaedalusParserRULE_elseBlock            = 25
	DaedalusParserRULE_elseIfBlock          = 26
	DaedalusParserRULE_ifBlock              = 27
	DaedalusParserRULE_ifBlockStatement     = 28
	DaedalusParserRULE_returnStatement      = 29
	DaedalusParserRULE_funcArgExpression    = 30
	DaedalusParserRULE_expressionBlock      = 31
	DaedalusParserRULE_expression           = 32
	DaedalusParserRULE_arrayIndex           = 33
	DaedalusParserRULE_arraySize            = 34
	DaedalusParserRULE_value                = 35
	DaedalusParserRULE_referenceAtom        = 36
	DaedalusParserRULE_reference            = 37
	DaedalusParserRULE_typeReference        = 38
	DaedalusParserRULE_anyIdentifier        = 39
	DaedalusParserRULE_nameNode             = 40
	DaedalusParserRULE_parentReference      = 41
	DaedalusParserRULE_assignmentOperator   = 42
	DaedalusParserRULE_addOperator          = 43
	DaedalusParserRULE_bitMoveOperator      = 44
	DaedalusParserRULE_compOperator         = 45
	DaedalusParserRULE_eqOperator           = 46
	DaedalusParserRULE_oneArgOperator       = 47
	DaedalusParserRULE_multOperator         = 48
	DaedalusParserRULE_binAndOperator       = 49
	DaedalusParserRULE_binOrOperator        = 50
	DaedalusParserRULE_logAndOperator       = 51
	DaedalusParserRULE_logOrOperator        = 52
)

// ISymbolSummaryContext is an interface to support dynamic dispatch.
type ISymbolSummaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolSummaryContext differentiates from other interfaces.
	IsSymbolSummaryContext()
}

type SymbolSummaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolSummaryContext() *SymbolSummaryContext {
	var p = new(SymbolSummaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_symbolSummary
	return p
}

func (*SymbolSummaryContext) IsSymbolSummaryContext() {}

func NewSymbolSummaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolSummaryContext {
	var p = new(SymbolSummaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_symbolSummary

	return p
}

func (s *SymbolSummaryContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolSummaryContext) AllSummaryComment() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserSummaryComment)
}

func (s *SymbolSummaryContext) SummaryComment(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserSummaryComment, i)
}

func (s *SymbolSummaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolSummaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolSummaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterSymbolSummary(s)
	}
}

func (s *SymbolSummaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitSymbolSummary(s)
	}
}

func (p *DaedalusParser) SymbolSummary() (localctx ISymbolSummaryContext) {
	this := p
	_ = this

	localctx = NewSymbolSummaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DaedalusParserRULE_symbolSummary)

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
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(106)
				p.Match(DaedalusParserSummaryComment)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// ILineCommentContext is an interface to support dynamic dispatch.
type ILineCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineCommentContext differentiates from other interfaces.
	IsLineCommentContext()
}

type LineCommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineCommentContext() *LineCommentContext {
	var p = new(LineCommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_lineComment
	return p
}

func (*LineCommentContext) IsLineCommentContext() {}

func NewLineCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineCommentContext {
	var p = new(LineCommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_lineComment

	return p
}

func (s *LineCommentContext) GetParser() antlr.Parser { return s.parser }

func (s *LineCommentContext) LineComment() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLineComment, 0)
}

func (s *LineCommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineCommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineCommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterLineComment(s)
	}
}

func (s *LineCommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitLineComment(s)
	}
}

func (p *DaedalusParser) LineComment() (localctx ILineCommentContext) {
	this := p
	_ = this

	localctx = NewLineCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DaedalusParserRULE_lineComment)

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
		p.SetState(111)
		p.Match(DaedalusParserLineComment)
	}

	return localctx
}

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

func (s *DaedalusFileContext) AllBlockDef() []IBlockDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockDefContext)(nil)).Elem())
	var tst = make([]IBlockDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockDefContext)
		}
	}

	return tst
}

func (s *DaedalusFileContext) BlockDef(i int) IBlockDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockDefContext)
}

func (s *DaedalusFileContext) AllInlineDef() []IInlineDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInlineDefContext)(nil)).Elem())
	var tst = make([]IInlineDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInlineDefContext)
		}
	}

	return tst
}

func (s *DaedalusFileContext) InlineDef(i int) IInlineDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInlineDefContext)
}

func (s *DaedalusFileContext) AllSymbolSummary() []ISymbolSummaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISymbolSummaryContext)(nil)).Elem())
	var tst = make([]ISymbolSummaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISymbolSummaryContext)
		}
	}

	return tst
}

func (s *DaedalusFileContext) SymbolSummary(i int) ISymbolSummaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolSummaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISymbolSummaryContext)
}

func (s *DaedalusFileContext) AllLineComment() []ILineCommentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineCommentContext)(nil)).Elem())
	var tst = make([]ILineCommentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineCommentContext)
		}
	}

	return tst
}

func (s *DaedalusFileContext) LineComment(i int) ILineCommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineCommentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineCommentContext)
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
	p.EnterRule(localctx, 4, DaedalusParserRULE_daedalusFile)

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
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(117)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(113)
					p.BlockDef()
				}

			case 2:
				{
					p.SetState(114)
					p.InlineDef()
				}

			case 3:
				{
					p.SetState(115)
					p.SymbolSummary()
				}

			case 4:
				{
					p.SetState(116)
					p.LineComment()
				}

			}

		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	{
		p.SetState(122)
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

func (s *BlockDefContext) FunctionDef() IFunctionDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDefContext)
}

func (s *BlockDefContext) ClassDef() IClassDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassDefContext)
}

func (s *BlockDefContext) PrototypeDef() IPrototypeDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrototypeDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrototypeDefContext)
}

func (s *BlockDefContext) InstanceDef() IInstanceDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstanceDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstanceDefContext)
}

func (s *BlockDefContext) AllSymbolSummary() []ISymbolSummaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISymbolSummaryContext)(nil)).Elem())
	var tst = make([]ISymbolSummaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISymbolSummaryContext)
		}
	}

	return tst
}

func (s *BlockDefContext) SymbolSummary(i int) ISymbolSummaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolSummaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISymbolSummaryContext)
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
	p.EnterRule(localctx, 6, DaedalusParserRULE_blockDef)
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
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DaedalusParserSummaryComment {
		{
			p.SetState(124)
			p.SymbolSummary()
		}

		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserFunc:
		{
			p.SetState(130)
			p.FunctionDef()
		}

	case DaedalusParserClass:
		{
			p.SetState(131)
			p.ClassDef()
		}

	case DaedalusParserPrototype:
		{
			p.SetState(132)
			p.PrototypeDef()
		}

	case DaedalusParserInstance:
		{
			p.SetState(133)
			p.InstanceDef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserT__0 {
		{
			p.SetState(136)
			p.Match(DaedalusParserT__0)
		}

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

func (s *InlineDefContext) ConstDef() IConstDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstDefContext)
}

func (s *InlineDefContext) VarDecl() IVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *InlineDefContext) InstanceDecl() IInstanceDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstanceDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstanceDeclContext)
}

func (s *InlineDefContext) AllSymbolSummary() []ISymbolSummaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISymbolSummaryContext)(nil)).Elem())
	var tst = make([]ISymbolSummaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISymbolSummaryContext)
		}
	}

	return tst
}

func (s *InlineDefContext) SymbolSummary(i int) ISymbolSummaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolSummaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISymbolSummaryContext)
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
	p.EnterRule(localctx, 8, DaedalusParserRULE_inlineDef)
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
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DaedalusParserSummaryComment {
		{
			p.SetState(139)
			p.SymbolSummary()
		}

		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserConst:
		{
			p.SetState(145)
			p.ConstDef()
		}

	case DaedalusParserVar:
		{
			p.SetState(146)
			p.VarDecl()
		}

	case DaedalusParserInstance:
		{
			p.SetState(147)
			p.InstanceDecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(150)
		p.Match(DaedalusParserT__0)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *FunctionDefContext) NameNode() INameNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *FunctionDefContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionDefContext) StatementBlock() IStatementBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementBlockContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 10, DaedalusParserRULE_functionDef)

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
		p.SetState(152)
		p.Match(DaedalusParserFunc)
	}
	{
		p.SetState(153)
		p.TypeReference()
	}
	{
		p.SetState(154)
		p.NameNode()
	}
	{
		p.SetState(155)
		p.ParameterList()
	}
	{
		p.SetState(156)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *ConstDefContext) AllConstValueDef() []IConstValueDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstValueDefContext)(nil)).Elem())
	var tst = make([]IConstValueDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstValueDefContext)
		}
	}

	return tst
}

func (s *ConstDefContext) ConstValueDef(i int) IConstValueDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstValueDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstValueDefContext)
}

func (s *ConstDefContext) AllConstArrayDef() []IConstArrayDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstArrayDefContext)(nil)).Elem())
	var tst = make([]IConstArrayDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstArrayDefContext)
		}
	}

	return tst
}

func (s *ConstDefContext) ConstArrayDef(i int) IConstArrayDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstArrayDefContext)(nil)).Elem(), i)

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
	p.EnterRule(localctx, 12, DaedalusParserRULE_constDef)
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
		p.SetState(158)
		p.Match(DaedalusParserConst)
	}
	{
		p.SetState(159)
		p.TypeReference()
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(160)
			p.ConstValueDef()
		}

	case 2:
		{
			p.SetState(161)
			p.ConstArrayDef()
		}

	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DaedalusParserT__1 {
		{
			p.SetState(164)
			p.Match(DaedalusParserT__1)
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(165)
				p.ConstValueDef()
			}

		case 2:
			{
				p.SetState(166)
				p.ConstArrayDef()
			}

		}

		p.SetState(173)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ClassDefContext) AllVarDecl() []IVarDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarDeclContext)(nil)).Elem())
	var tst = make([]IVarDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarDeclContext)
		}
	}

	return tst
}

func (s *ClassDefContext) VarDecl(i int) IVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *ClassDefContext) AllSymbolSummary() []ISymbolSummaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISymbolSummaryContext)(nil)).Elem())
	var tst = make([]ISymbolSummaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISymbolSummaryContext)
		}
	}

	return tst
}

func (s *ClassDefContext) SymbolSummary(i int) ISymbolSummaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolSummaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISymbolSummaryContext)
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
	p.EnterRule(localctx, 14, DaedalusParserRULE_classDef)

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
		p.SetState(174)
		p.Match(DaedalusParserClass)
	}
	{
		p.SetState(175)
		p.NameNode()
	}
	{
		p.SetState(176)
		p.Match(DaedalusParserT__2)
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(181)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case DaedalusParserVar:
				{
					p.SetState(177)
					p.VarDecl()
				}
				{
					p.SetState(178)
					p.Match(DaedalusParserT__0)
				}

			case DaedalusParserSummaryComment:
				{
					p.SetState(180)
					p.SymbolSummary()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}
	{
		p.SetState(186)
		p.Match(DaedalusParserT__3)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *PrototypeDefContext) ParentReference() IParentReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParentReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParentReferenceContext)
}

func (s *PrototypeDefContext) StatementBlock() IStatementBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementBlockContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 16, DaedalusParserRULE_prototypeDef)

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
		p.SetState(188)
		p.Match(DaedalusParserPrototype)
	}
	{
		p.SetState(189)
		p.NameNode()
	}
	{
		p.SetState(190)
		p.Match(DaedalusParserT__4)
	}
	{
		p.SetState(191)
		p.ParentReference()
	}
	{
		p.SetState(192)
		p.Match(DaedalusParserT__5)
	}
	{
		p.SetState(193)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *InstanceDefContext) ParentReference() IParentReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParentReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParentReferenceContext)
}

func (s *InstanceDefContext) StatementBlock() IStatementBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementBlockContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 18, DaedalusParserRULE_instanceDef)

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
		p.SetState(195)
		p.Match(DaedalusParserInstance)
	}
	{
		p.SetState(196)
		p.NameNode()
	}
	{
		p.SetState(197)
		p.Match(DaedalusParserT__4)
	}
	{
		p.SetState(198)
		p.ParentReference()
	}
	{
		p.SetState(199)
		p.Match(DaedalusParserT__5)
	}
	{
		p.SetState(200)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameNodeContext)(nil)).Elem())
	var tst = make([]INameNodeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameNodeContext)
		}
	}

	return tst
}

func (s *InstanceDeclContext) NameNode(i int) INameNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameNodeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *InstanceDeclContext) ParentReference() IParentReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParentReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParentReferenceContext)
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
	p.EnterRule(localctx, 20, DaedalusParserRULE_instanceDecl)

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
		p.SetState(202)
		p.Match(DaedalusParserInstance)
	}
	{
		p.SetState(203)
		p.NameNode()
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(204)
				p.Match(DaedalusParserT__1)
			}
			{
				p.SetState(205)
				p.NameNode()
			}

		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}
	{
		p.SetState(211)
		p.Match(DaedalusParserT__4)
	}
	{
		p.SetState(212)
		p.ParentReference()
	}
	{
		p.SetState(213)
		p.Match(DaedalusParserT__5)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *VarDeclContext) AllVarValueDecl() []IVarValueDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarValueDeclContext)(nil)).Elem())
	var tst = make([]IVarValueDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarValueDeclContext)
		}
	}

	return tst
}

func (s *VarDeclContext) VarValueDecl(i int) IVarValueDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarValueDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarValueDeclContext)
}

func (s *VarDeclContext) AllVarArrayDecl() []IVarArrayDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarArrayDeclContext)(nil)).Elem())
	var tst = make([]IVarArrayDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarArrayDeclContext)
		}
	}

	return tst
}

func (s *VarDeclContext) VarArrayDecl(i int) IVarArrayDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarArrayDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarArrayDeclContext)
}

func (s *VarDeclContext) AllVarDecl() []IVarDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarDeclContext)(nil)).Elem())
	var tst = make([]IVarDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarDeclContext)
		}
	}

	return tst
}

func (s *VarDeclContext) VarDecl(i int) IVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclContext)(nil)).Elem(), i)

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
	p.EnterRule(localctx, 22, DaedalusParserRULE_varDecl)

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
		p.SetState(215)
		p.Match(DaedalusParserVar)
	}
	{
		p.SetState(216)
		p.TypeReference()
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(217)
			p.VarValueDecl()
		}

	case 2:
		{
			p.SetState(218)
			p.VarArrayDecl()
		}

	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(221)
				p.Match(DaedalusParserT__1)
			}
			p.SetState(225)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(222)
					p.VarDecl()
				}

			case 2:
				{
					p.SetState(223)
					p.VarValueDecl()
				}

			case 3:
				{
					p.SetState(224)
					p.VarArrayDecl()
				}

			}

		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ConstArrayDefContext) ArraySize() IArraySizeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArraySizeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArraySizeContext)
}

func (s *ConstArrayDefContext) ConstArrayAssignment() IConstArrayAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstArrayAssignmentContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 24, DaedalusParserRULE_constArrayDef)

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
		p.SetState(232)
		p.NameNode()
	}
	{
		p.SetState(233)
		p.Match(DaedalusParserT__6)
	}
	{
		p.SetState(234)
		p.ArraySize()
	}
	{
		p.SetState(235)
		p.Match(DaedalusParserT__7)
	}
	{
		p.SetState(236)
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

func (s *ConstArrayAssignmentContext) AllExpressionBlock() []IExpressionBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionBlockContext)(nil)).Elem())
	var tst = make([]IExpressionBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionBlockContext)
		}
	}

	return tst
}

func (s *ConstArrayAssignmentContext) ExpressionBlock(i int) IExpressionBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionBlockContext)(nil)).Elem(), i)

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
	p.EnterRule(localctx, 26, DaedalusParserRULE_constArrayAssignment)

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
		p.SetState(238)
		p.Match(DaedalusParserT__8)
	}
	{
		p.SetState(239)
		p.Match(DaedalusParserT__2)
	}

	{
		p.SetState(240)
		p.ExpressionBlock()
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(241)
				p.Match(DaedalusParserT__1)
			}
			{
				p.SetState(242)
				p.ExpressionBlock()
			}

		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	{
		p.SetState(248)
		p.Match(DaedalusParserT__3)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ConstValueDefContext) ConstValueAssignment() IConstValueAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstValueAssignmentContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 28, DaedalusParserRULE_constValueDef)

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
		p.SetState(250)
		p.NameNode()
	}
	{
		p.SetState(251)
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

func (s *ConstValueAssignmentContext) ExpressionBlock() IExpressionBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionBlockContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 30, DaedalusParserRULE_constValueAssignment)

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
		p.SetState(253)
		p.Match(DaedalusParserT__8)
	}
	{
		p.SetState(254)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *VarArrayDeclContext) ArraySize() IArraySizeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArraySizeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArraySizeContext)
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
	p.EnterRule(localctx, 32, DaedalusParserRULE_varArrayDecl)

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
		p.SetState(256)
		p.NameNode()
	}
	{
		p.SetState(257)
		p.Match(DaedalusParserT__6)
	}
	{
		p.SetState(258)
		p.ArraySize()
	}
	{
		p.SetState(259)
		p.Match(DaedalusParserT__7)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameNodeContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 34, DaedalusParserRULE_varValueDecl)

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
		p.SetState(261)
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

func (s *ParameterListContext) AllParameterDecl() []IParameterDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterDeclContext)(nil)).Elem())
	var tst = make([]IParameterDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterDeclContext)
		}
	}

	return tst
}

func (s *ParameterListContext) ParameterDecl(i int) IParameterDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterDeclContext)(nil)).Elem(), i)

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
	p.EnterRule(localctx, 36, DaedalusParserRULE_parameterList)
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
		p.SetState(263)
		p.Match(DaedalusParserT__4)
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserVar {
		{
			p.SetState(264)
			p.ParameterDecl()
		}
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

		for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1+1 {
				{
					p.SetState(265)
					p.Match(DaedalusParserT__1)
				}
				{
					p.SetState(266)
					p.ParameterDecl()
				}

			}
			p.SetState(271)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
		}

	}
	{
		p.SetState(274)
		p.Match(DaedalusParserT__5)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *ParameterDeclContext) NameNode() INameNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ParameterDeclContext) ArraySize() IArraySizeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArraySizeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArraySizeContext)
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
	p.EnterRule(localctx, 38, DaedalusParserRULE_parameterDecl)
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
		p.SetState(276)
		p.Match(DaedalusParserVar)
	}
	{
		p.SetState(277)
		p.TypeReference()
	}
	{
		p.SetState(278)
		p.NameNode()
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserT__6 {
		{
			p.SetState(279)
			p.Match(DaedalusParserT__6)
		}
		{
			p.SetState(280)
			p.ArraySize()
		}
		{
			p.SetState(281)
			p.Match(DaedalusParserT__7)
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

func (s *StatementBlockContext) AllSymbolSummary() []ISymbolSummaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISymbolSummaryContext)(nil)).Elem())
	var tst = make([]ISymbolSummaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISymbolSummaryContext)
		}
	}

	return tst
}

func (s *StatementBlockContext) SymbolSummary(i int) ISymbolSummaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolSummaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISymbolSummaryContext)
}

func (s *StatementBlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementBlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementBlockContext) AllIfBlockStatement() []IIfBlockStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIfBlockStatementContext)(nil)).Elem())
	var tst = make([]IIfBlockStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIfBlockStatementContext)
		}
	}

	return tst
}

func (s *StatementBlockContext) IfBlockStatement(i int) IIfBlockStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfBlockStatementContext)(nil)).Elem(), i)

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
	p.EnterRule(localctx, 40, DaedalusParserRULE_statementBlock)
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
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DaedalusParserSummaryComment {
		{
			p.SetState(285)
			p.SymbolSummary()
		}

		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(291)
		p.Match(DaedalusParserT__2)
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(298)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case DaedalusParserT__4, DaedalusParserT__14, DaedalusParserT__15, DaedalusParserT__24, DaedalusParserT__25, DaedalusParserConst, DaedalusParserVar, DaedalusParserInt, DaedalusParserFunc, DaedalusParserStringKeyword, DaedalusParserClass, DaedalusParserVoid, DaedalusParserReturn, DaedalusParserFloat, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNull, DaedalusParserNoFunc, DaedalusParserIdentifier, DaedalusParserIntegerLiteral, DaedalusParserFloatLiteral, DaedalusParserStringLiteral:
				{
					p.SetState(292)
					p.Statement()
				}

			case DaedalusParserIf:
				{
					p.SetState(293)
					p.IfBlockStatement()
				}
				p.SetState(295)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == DaedalusParserT__0 {
					{
						p.SetState(294)
						p.Match(DaedalusParserT__0)
					}

				}

			case DaedalusParserSummaryComment:
				{
					p.SetState(297)
					p.SymbolSummary()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}
	{
		p.SetState(303)
		p.Match(DaedalusParserT__3)
	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(304)
				p.SymbolSummary()
			}

		}
		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) ConstDef() IConstDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstDefContext)
}

func (s *StatementContext) VarDecl() IVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *StatementContext) FuncCall() IFuncCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
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
	p.EnterRule(localctx, 42, DaedalusParserRULE_statement)

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

	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(310)
			p.Assignment()
		}
		{
			p.SetState(311)
			p.Match(DaedalusParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(313)
			p.ReturnStatement()
		}
		{
			p.SetState(314)
			p.Match(DaedalusParserT__0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(316)
			p.ConstDef()
		}
		{
			p.SetState(317)
			p.Match(DaedalusParserT__0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(319)
			p.VarDecl()
		}
		{
			p.SetState(320)
			p.Match(DaedalusParserT__0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(322)
			p.FuncCall()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(323)
			p.expression(0)
		}
		{
			p.SetState(324)
			p.Match(DaedalusParserT__0)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *FuncCallContext) AllFuncArgExpression() []IFuncArgExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncArgExpressionContext)(nil)).Elem())
	var tst = make([]IFuncArgExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncArgExpressionContext)
		}
	}

	return tst
}

func (s *FuncCallContext) FuncArgExpression(i int) IFuncArgExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncArgExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncArgExpressionContext)
}

func (s *FuncCallContext) LineComment() ILineCommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineCommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineCommentContext)
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
	p.EnterRule(localctx, 44, DaedalusParserRULE_funcCall)
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
		p.SetState(328)
		p.NameNode()
	}
	{
		p.SetState(329)
		p.Match(DaedalusParserT__4)
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DaedalusParserT__4)|(1<<DaedalusParserT__14)|(1<<DaedalusParserT__15)|(1<<DaedalusParserT__24)|(1<<DaedalusParserT__25))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(DaedalusParserVar-35))|(1<<(DaedalusParserInt-35))|(1<<(DaedalusParserFunc-35))|(1<<(DaedalusParserStringKeyword-35))|(1<<(DaedalusParserClass-35))|(1<<(DaedalusParserVoid-35))|(1<<(DaedalusParserFloat-35))|(1<<(DaedalusParserPrototype-35))|(1<<(DaedalusParserInstance-35))|(1<<(DaedalusParserNull-35))|(1<<(DaedalusParserNoFunc-35))|(1<<(DaedalusParserIdentifier-35))|(1<<(DaedalusParserIntegerLiteral-35))|(1<<(DaedalusParserFloatLiteral-35))|(1<<(DaedalusParserStringLiteral-35)))) != 0) {
		{
			p.SetState(330)
			p.FuncArgExpression()
		}
		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

		for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1+1 {
				{
					p.SetState(331)
					p.Match(DaedalusParserT__1)
				}
				{
					p.SetState(332)
					p.FuncArgExpression()
				}

			}
			p.SetState(337)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
		}

	}
	{
		p.SetState(340)
		p.Match(DaedalusParserT__5)
	}
	{
		p.SetState(341)
		p.Match(DaedalusParserT__0)
	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(342)
			p.LineComment()
		}

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *AssignmentContext) AssignmentOperator() IAssignmentOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentOperatorContext)
}

func (s *AssignmentContext) ExpressionBlock() IExpressionBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionBlockContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 46, DaedalusParserRULE_assignment)

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
		p.Reference()
	}
	{
		p.SetState(346)
		p.AssignmentOperator()
	}
	{
		p.SetState(347)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionBlockContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 48, DaedalusParserRULE_ifCondition)

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
		p.SetState(349)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementBlockContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 50, DaedalusParserRULE_elseBlock)

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
		p.Match(DaedalusParserElse)
	}
	{
		p.SetState(352)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfConditionContext)
}

func (s *ElseIfBlockContext) StatementBlock() IStatementBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementBlockContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 52, DaedalusParserRULE_elseIfBlock)

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
		p.Match(DaedalusParserElse)
	}
	{
		p.SetState(355)
		p.Match(DaedalusParserIf)
	}
	{
		p.SetState(356)
		p.IfCondition()
	}
	{
		p.SetState(357)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfConditionContext)
}

func (s *IfBlockContext) StatementBlock() IStatementBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementBlockContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 54, DaedalusParserRULE_ifBlock)

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
		p.SetState(359)
		p.Match(DaedalusParserIf)
	}
	{
		p.SetState(360)
		p.IfCondition()
	}
	{
		p.SetState(361)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *IfBlockStatementContext) AllElseIfBlock() []IElseIfBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElseIfBlockContext)(nil)).Elem())
	var tst = make([]IElseIfBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElseIfBlockContext)
		}
	}

	return tst
}

func (s *IfBlockStatementContext) ElseIfBlock(i int) IElseIfBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseIfBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElseIfBlockContext)
}

func (s *IfBlockStatementContext) ElseBlock() IElseBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseBlockContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 56, DaedalusParserRULE_ifBlockStatement)
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
		p.SetState(363)
		p.IfBlock()
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(364)
				p.ElseIfBlock()
			}

		}
		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
	}
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserElse {
		{
			p.SetState(370)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionBlockContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 58, DaedalusParserRULE_returnStatement)
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
		p.SetState(373)
		p.Match(DaedalusParserReturn)
	}
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DaedalusParserT__4)|(1<<DaedalusParserT__14)|(1<<DaedalusParserT__15)|(1<<DaedalusParserT__24)|(1<<DaedalusParserT__25))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(DaedalusParserVar-35))|(1<<(DaedalusParserInt-35))|(1<<(DaedalusParserFunc-35))|(1<<(DaedalusParserStringKeyword-35))|(1<<(DaedalusParserClass-35))|(1<<(DaedalusParserVoid-35))|(1<<(DaedalusParserFloat-35))|(1<<(DaedalusParserPrototype-35))|(1<<(DaedalusParserInstance-35))|(1<<(DaedalusParserNull-35))|(1<<(DaedalusParserNoFunc-35))|(1<<(DaedalusParserIdentifier-35))|(1<<(DaedalusParserIntegerLiteral-35))|(1<<(DaedalusParserFloatLiteral-35))|(1<<(DaedalusParserStringLiteral-35)))) != 0) {
		{
			p.SetState(374)
			p.ExpressionBlock()
		}

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionBlockContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 60, DaedalusParserRULE_funcArgExpression)

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
		p.SetState(377)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 62, DaedalusParserRULE_expressionBlock)

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
		p.SetState(379)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BitMoveExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitMoveExpressionContext) BitMoveOperator() IBitMoveOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitMoveOperatorContext)(nil)).Elem(), 0)

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

type OneArgExpressionContext struct {
	*ExpressionContext
}

func NewOneArgExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OneArgExpressionContext {
	var p = new(OneArgExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OneArgExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneArgExpressionContext) OneArgOperator() IOneArgOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneArgOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOneArgOperatorContext)
}

func (s *OneArgExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OneArgExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterOneArgExpression(s)
	}
}

func (s *OneArgExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitOneArgExpression(s)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EqExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqExpressionContext) EqOperator() IEqOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqOperatorContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddExpressionContext) AddOperator() IAddOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddOperatorContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CompExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompExpressionContext) CompOperator() ICompOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompOperatorContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LogOrExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogOrExpressionContext) LogOrOperator() ILogOrOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogOrOperatorContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BinAndExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinAndExpressionContext) BinAndOperator() IBinAndOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinAndOperatorContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BinOrExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinOrExpressionContext) BinOrOperator() IBinOrOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinOrOperatorContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MultExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MultExpressionContext) MultOperator() IMultOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultOperatorContext)(nil)).Elem(), 0)

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

func (s *BracketExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LogAndExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogAndExpressionContext) LogAndOperator() ILogAndOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogAndOperatorContext)(nil)).Elem(), 0)

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
	_startState := 64
	p.EnterRecursionRule(localctx, 64, DaedalusParserRULE_expression, _p)

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
	p.SetState(390)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserT__4:
		localctx = NewBracketExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(382)
			p.Match(DaedalusParserT__4)
		}
		{
			p.SetState(383)
			p.expression(0)
		}
		{
			p.SetState(384)
			p.Match(DaedalusParserT__5)
		}

	case DaedalusParserT__14, DaedalusParserT__15, DaedalusParserT__24, DaedalusParserT__25:
		localctx = NewOneArgExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(386)
			p.OneArgOperator()
		}
		{
			p.SetState(387)
			p.expression(11)
		}

	case DaedalusParserVar, DaedalusParserInt, DaedalusParserFunc, DaedalusParserStringKeyword, DaedalusParserClass, DaedalusParserVoid, DaedalusParserFloat, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNull, DaedalusParserNoFunc, DaedalusParserIdentifier, DaedalusParserIntegerLiteral, DaedalusParserFloatLiteral, DaedalusParserStringLiteral:
		localctx = NewValExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(389)
			p.Value()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(428)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(392)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(393)
					p.MultOperator()
				}
				{
					p.SetState(394)
					p.expression(11)
				}

			case 2:
				localctx = NewAddExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(396)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(397)
					p.AddOperator()
				}
				{
					p.SetState(398)
					p.expression(10)
				}

			case 3:
				localctx = NewBitMoveExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(400)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(401)
					p.BitMoveOperator()
				}
				{
					p.SetState(402)
					p.expression(9)
				}

			case 4:
				localctx = NewCompExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(404)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(405)
					p.CompOperator()
				}
				{
					p.SetState(406)
					p.expression(8)
				}

			case 5:
				localctx = NewEqExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(408)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(409)
					p.EqOperator()
				}
				{
					p.SetState(410)
					p.expression(7)
				}

			case 6:
				localctx = NewBinAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(412)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(413)
					p.BinAndOperator()
				}
				{
					p.SetState(414)
					p.expression(6)
				}

			case 7:
				localctx = NewBinOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(416)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(417)
					p.BinOrOperator()
				}
				{
					p.SetState(418)
					p.expression(5)
				}

			case 8:
				localctx = NewLogAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(420)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(421)
					p.LogAndOperator()
				}
				{
					p.SetState(422)
					p.expression(4)
				}

			case 9:
				localctx = NewLogOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(424)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(425)
					p.LogOrOperator()
				}
				{
					p.SetState(426)
					p.expression(3)
				}

			}

		}
		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferenceAtomContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 66, DaedalusParserRULE_arrayIndex)

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

	p.SetState(435)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(433)
			p.Match(DaedalusParserIntegerLiteral)
		}

	case DaedalusParserVar, DaedalusParserInt, DaedalusParserFunc, DaedalusParserStringKeyword, DaedalusParserClass, DaedalusParserVoid, DaedalusParserFloat, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNull, DaedalusParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(434)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferenceAtomContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 68, DaedalusParserRULE_arraySize)

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

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(437)
			p.Match(DaedalusParserIntegerLiteral)
		}

	case DaedalusParserVar, DaedalusParserInt, DaedalusParserFunc, DaedalusParserStringKeyword, DaedalusParserClass, DaedalusParserVoid, DaedalusParserFloat, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNull, DaedalusParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(438)
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

type AnyIdentifierValueContext struct {
	*ValueContext
}

func NewAnyIdentifierValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AnyIdentifierValueContext {
	var p = new(AnyIdentifierValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *AnyIdentifierValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyIdentifierValueContext) NameNode() INameNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *AnyIdentifierValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterAnyIdentifierValue(s)
	}
}

func (s *AnyIdentifierValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitAnyIdentifierValue(s)
	}
}

type NoFuncLiteralValueContext struct {
	*ValueContext
}

func NewNoFuncLiteralValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NoFuncLiteralValueContext {
	var p = new(NoFuncLiteralValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NoFuncLiteralValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoFuncLiteralValueContext) NoFunc() antlr.TerminalNode {
	return s.GetToken(DaedalusParserNoFunc, 0)
}

func (s *NoFuncLiteralValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterNoFuncLiteralValue(s)
	}
}

func (s *NoFuncLiteralValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitNoFuncLiteralValue(s)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferenceContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 70, DaedalusParserRULE_value)

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

	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIntegerLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(441)
			p.Match(DaedalusParserIntegerLiteral)
		}

	case 2:
		localctx = NewFloatLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(442)
			p.Match(DaedalusParserFloatLiteral)
		}

	case 3:
		localctx = NewStringLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(443)
			p.Match(DaedalusParserStringLiteral)
		}

	case 4:
		localctx = NewNullLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(444)
			p.Match(DaedalusParserNull)
		}

	case 5:
		localctx = NewNoFuncLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(445)
			p.Match(DaedalusParserNoFunc)
		}

	case 6:
		localctx = NewFuncCallValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(446)
			p.FuncCall()
		}

	case 7:
		localctx = NewReferenceValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(447)
			p.Reference()
		}

	case 8:
		localctx = NewAnyIdentifierValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(448)
			p.NameNode()
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ReferenceAtomContext) ArrayIndex() IArrayIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayIndexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayIndexContext)
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
	p.EnterRule(localctx, 72, DaedalusParserRULE_referenceAtom)

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
		p.SetState(451)
		p.NameNode()
	}
	p.SetState(456)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(452)
			p.Match(DaedalusParserT__6)
		}
		{
			p.SetState(453)
			p.ArrayIndex()
		}
		{
			p.SetState(454)
			p.Match(DaedalusParserT__7)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReferenceAtomContext)(nil)).Elem())
	var tst = make([]IReferenceAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReferenceAtomContext)
		}
	}

	return tst
}

func (s *ReferenceContext) ReferenceAtom(i int) IReferenceAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferenceAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReferenceAtomContext)
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
	p.EnterRule(localctx, 74, DaedalusParserRULE_reference)

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
		p.SetState(458)
		p.ReferenceAtom()
	}
	p.SetState(461)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(459)
			p.Match(DaedalusParserT__9)
		}
		{
			p.SetState(460)
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
	p.EnterRule(localctx, 76, DaedalusParserRULE_typeReference)
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
		p.SetState(463)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(DaedalusParserInt-37))|(1<<(DaedalusParserFunc-37))|(1<<(DaedalusParserStringKeyword-37))|(1<<(DaedalusParserVoid-37))|(1<<(DaedalusParserFloat-37))|(1<<(DaedalusParserInstance-37))|(1<<(DaedalusParserIdentifier-37)))) != 0) {
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
	p.EnterRule(localctx, 78, DaedalusParserRULE_anyIdentifier)
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
		p.SetState(465)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(DaedalusParserVar-35))|(1<<(DaedalusParserInt-35))|(1<<(DaedalusParserFunc-35))|(1<<(DaedalusParserStringKeyword-35))|(1<<(DaedalusParserClass-35))|(1<<(DaedalusParserVoid-35))|(1<<(DaedalusParserFloat-35))|(1<<(DaedalusParserPrototype-35))|(1<<(DaedalusParserInstance-35))|(1<<(DaedalusParserNull-35))|(1<<(DaedalusParserIdentifier-35)))) != 0) {
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnyIdentifierContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 80, DaedalusParserRULE_nameNode)

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
		p.SetState(467)
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
	p.EnterRule(localctx, 82, DaedalusParserRULE_parentReference)

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
	p.EnterRule(localctx, 84, DaedalusParserRULE_assignmentOperator)
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
		p.SetState(471)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DaedalusParserT__8)|(1<<DaedalusParserT__10)|(1<<DaedalusParserT__11)|(1<<DaedalusParserT__12)|(1<<DaedalusParserT__13))) != 0) {
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
	p.EnterRule(localctx, 86, DaedalusParserRULE_addOperator)
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
		p.SetState(473)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DaedalusParserT__14 || _la == DaedalusParserT__15) {
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
	p.EnterRule(localctx, 88, DaedalusParserRULE_bitMoveOperator)
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
		p.SetState(475)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DaedalusParserT__16 || _la == DaedalusParserT__17) {
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
	p.EnterRule(localctx, 90, DaedalusParserRULE_compOperator)
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
		p.SetState(477)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DaedalusParserT__18)|(1<<DaedalusParserT__19)|(1<<DaedalusParserT__20)|(1<<DaedalusParserT__21))) != 0) {
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
	p.EnterRule(localctx, 92, DaedalusParserRULE_eqOperator)
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
		p.SetState(479)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DaedalusParserT__22 || _la == DaedalusParserT__23) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOneArgOperatorContext is an interface to support dynamic dispatch.
type IOneArgOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneArgOperatorContext differentiates from other interfaces.
	IsOneArgOperatorContext()
}

type OneArgOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneArgOperatorContext() *OneArgOperatorContext {
	var p = new(OneArgOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DaedalusParserRULE_oneArgOperator
	return p
}

func (*OneArgOperatorContext) IsOneArgOperatorContext() {}

func NewOneArgOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneArgOperatorContext {
	var p = new(OneArgOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_oneArgOperator

	return p
}

func (s *OneArgOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *OneArgOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneArgOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneArgOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterOneArgOperator(s)
	}
}

func (s *OneArgOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitOneArgOperator(s)
	}
}

func (p *DaedalusParser) OneArgOperator() (localctx IOneArgOperatorContext) {
	this := p
	_ = this

	localctx = NewOneArgOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, DaedalusParserRULE_oneArgOperator)
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
		p.SetState(481)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DaedalusParserT__14)|(1<<DaedalusParserT__15)|(1<<DaedalusParserT__24)|(1<<DaedalusParserT__25))) != 0) {
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
	p.EnterRule(localctx, 96, DaedalusParserRULE_multOperator)
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
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DaedalusParserT__26)|(1<<DaedalusParserT__27)|(1<<DaedalusParserT__28))) != 0) {
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
	p.EnterRule(localctx, 98, DaedalusParserRULE_binAndOperator)

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
		p.SetState(485)
		p.Match(DaedalusParserT__29)
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
	p.EnterRule(localctx, 100, DaedalusParserRULE_binOrOperator)

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
		p.Match(DaedalusParserT__30)
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
	p.EnterRule(localctx, 102, DaedalusParserRULE_logAndOperator)

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
		p.Match(DaedalusParserT__31)
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
	p.EnterRule(localctx, 104, DaedalusParserRULE_logOrOperator)

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
		p.SetState(491)
		p.Match(DaedalusParserT__32)
	}

	return localctx
}

func (p *DaedalusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 32:
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
