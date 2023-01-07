// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package terraform // Terraform
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

type TerraformParser struct {
	*antlr.BaseParser
}

var terraformParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func terraformParserInit() {
	staticData := &terraformParserStaticData
	staticData.literalNames = []string{
		"", "'terraform'", "'resource'", "'data'", "'output'", "'locals'", "'module'",
		"'='", "'local'", "'var'", "'?'", "':'", "'for'", "'jsonencode'", "','",
		"'['", "']'", "'file'", "'+'", "'-'", "'/'", "'%'", "'>'", "'>='", "'<'",
		"'<='", "'=='", "'!='", "'&&'", "'||'", "'variable'", "'provider'",
		"'in'", "'*'", "'.'", "'{'", "'}'", "'('", "')'", "", "'nul'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "VARIABLE", "PROVIDER",
		"IN", "STAR", "DOT", "LCURL", "RCURL", "LPAREN", "RPAREN", "EOF_", "NULL_",
		"NATURAL_NUMBER", "BOOL", "DESCRIPTION", "MULTILINESTRING", "STRING",
		"IDENTIFIER", "COMMENT", "BLOCKCOMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"file_", "terraform", "resource", "data", "provider", "output", "local",
		"module", "variable", "block", "blocktype", "resourcetype", "name",
		"label", "blockbody", "argument", "identifier", "identifierchain", "inline_index",
		"expression", "forloop", "section", "val", "functioncall", "functionname",
		"functionarguments", "index", "filedecl", "list_", "map_", "string",
		"signed_number", "operator_", "number",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 49, 310, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 4, 0, 77, 8, 0, 11, 0, 12, 0, 78, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 5, 9, 117, 8, 9, 10, 9, 12, 9,
		120, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 5, 14, 135, 8, 14, 10, 14, 12, 14, 138, 9,
		14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 3, 16, 148,
		8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 3, 17, 154, 8, 17, 1, 17, 1, 17, 5,
		17, 158, 8, 17, 10, 17, 12, 17, 161, 9, 17, 1, 17, 1, 17, 1, 17, 5, 17,
		166, 8, 17, 10, 17, 12, 17, 169, 9, 17, 1, 17, 1, 17, 1, 17, 5, 17, 174,
		8, 17, 10, 17, 12, 17, 177, 9, 17, 3, 17, 179, 8, 17, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 190, 8, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 202,
		8, 19, 10, 19, 12, 19, 205, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 21, 3, 21, 217, 8, 21, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 228, 8, 22, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 238, 8, 23, 10, 23,
		12, 23, 241, 9, 23, 1, 23, 3, 23, 244, 8, 23, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 25, 1, 25, 5, 25, 252, 8, 25, 10, 25, 12, 25, 255, 9, 25, 3, 25, 257,
		8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 1, 28, 5, 28, 272, 8, 28, 10, 28, 12, 28, 275, 9, 28,
		1, 28, 3, 28, 278, 8, 28, 3, 28, 280, 8, 28, 1, 28, 1, 28, 1, 29, 1, 29,
		1, 29, 3, 29, 287, 8, 29, 5, 29, 289, 8, 29, 10, 29, 12, 29, 292, 9, 29,
		1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 3, 31, 299, 8, 31, 1, 31, 1, 31, 1,
		32, 1, 32, 1, 33, 1, 33, 1, 33, 3, 33, 308, 8, 33, 1, 33, 1, 239, 1, 38,
		34, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 0, 5, 3,
		0, 3, 3, 6, 6, 8, 9, 2, 0, 30, 32, 46, 46, 1, 0, 44, 45, 1, 0, 18, 19,
		2, 0, 18, 29, 33, 33, 318, 0, 76, 1, 0, 0, 0, 2, 82, 1, 0, 0, 0, 4, 85,
		1, 0, 0, 0, 6, 90, 1, 0, 0, 0, 8, 95, 1, 0, 0, 0, 10, 99, 1, 0, 0, 0, 12,
		103, 1, 0, 0, 0, 14, 106, 1, 0, 0, 0, 16, 110, 1, 0, 0, 0, 18, 114, 1,
		0, 0, 0, 20, 123, 1, 0, 0, 0, 22, 125, 1, 0, 0, 0, 24, 127, 1, 0, 0, 0,
		26, 129, 1, 0, 0, 0, 28, 131, 1, 0, 0, 0, 30, 141, 1, 0, 0, 0, 32, 147,
		1, 0, 0, 0, 34, 178, 1, 0, 0, 0, 36, 180, 1, 0, 0, 0, 38, 189, 1, 0, 0,
		0, 40, 206, 1, 0, 0, 0, 42, 216, 1, 0, 0, 0, 44, 227, 1, 0, 0, 0, 46, 243,
		1, 0, 0, 0, 48, 245, 1, 0, 0, 0, 50, 256, 1, 0, 0, 0, 52, 258, 1, 0, 0,
		0, 54, 262, 1, 0, 0, 0, 56, 267, 1, 0, 0, 0, 58, 283, 1, 0, 0, 0, 60, 295,
		1, 0, 0, 0, 62, 298, 1, 0, 0, 0, 64, 302, 1, 0, 0, 0, 66, 304, 1, 0, 0,
		0, 68, 77, 3, 12, 6, 0, 69, 77, 3, 14, 7, 0, 70, 77, 3, 10, 5, 0, 71, 77,
		3, 8, 4, 0, 72, 77, 3, 16, 8, 0, 73, 77, 3, 6, 3, 0, 74, 77, 3, 4, 2, 0,
		75, 77, 3, 2, 1, 0, 76, 68, 1, 0, 0, 0, 76, 69, 1, 0, 0, 0, 76, 70, 1,
		0, 0, 0, 76, 71, 1, 0, 0, 0, 76, 72, 1, 0, 0, 0, 76, 73, 1, 0, 0, 0, 76,
		74, 1, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 76, 1, 0, 0,
		0, 78, 79, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 81, 5, 0, 0, 1, 81, 1, 1,
		0, 0, 0, 82, 83, 5, 1, 0, 0, 83, 84, 3, 28, 14, 0, 84, 3, 1, 0, 0, 0, 85,
		86, 5, 2, 0, 0, 86, 87, 3, 22, 11, 0, 87, 88, 3, 24, 12, 0, 88, 89, 3,
		28, 14, 0, 89, 5, 1, 0, 0, 0, 90, 91, 5, 3, 0, 0, 91, 92, 3, 22, 11, 0,
		92, 93, 3, 24, 12, 0, 93, 94, 3, 28, 14, 0, 94, 7, 1, 0, 0, 0, 95, 96,
		5, 31, 0, 0, 96, 97, 3, 22, 11, 0, 97, 98, 3, 28, 14, 0, 98, 9, 1, 0, 0,
		0, 99, 100, 5, 4, 0, 0, 100, 101, 3, 24, 12, 0, 101, 102, 3, 28, 14, 0,
		102, 11, 1, 0, 0, 0, 103, 104, 5, 5, 0, 0, 104, 105, 3, 28, 14, 0, 105,
		13, 1, 0, 0, 0, 106, 107, 5, 6, 0, 0, 107, 108, 3, 24, 12, 0, 108, 109,
		3, 28, 14, 0, 109, 15, 1, 0, 0, 0, 110, 111, 5, 30, 0, 0, 111, 112, 3,
		24, 12, 0, 112, 113, 3, 28, 14, 0, 113, 17, 1, 0, 0, 0, 114, 118, 3, 20,
		10, 0, 115, 117, 3, 26, 13, 0, 116, 115, 1, 0, 0, 0, 117, 120, 1, 0, 0,
		0, 118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 121, 1, 0, 0, 0, 120,
		118, 1, 0, 0, 0, 121, 122, 3, 28, 14, 0, 122, 19, 1, 0, 0, 0, 123, 124,
		5, 46, 0, 0, 124, 21, 1, 0, 0, 0, 125, 126, 5, 45, 0, 0, 126, 23, 1, 0,
		0, 0, 127, 128, 5, 45, 0, 0, 128, 25, 1, 0, 0, 0, 129, 130, 5, 45, 0, 0,
		130, 27, 1, 0, 0, 0, 131, 136, 5, 35, 0, 0, 132, 135, 3, 30, 15, 0, 133,
		135, 3, 18, 9, 0, 134, 132, 1, 0, 0, 0, 134, 133, 1, 0, 0, 0, 135, 138,
		1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 139, 1, 0,
		0, 0, 138, 136, 1, 0, 0, 0, 139, 140, 5, 36, 0, 0, 140, 29, 1, 0, 0, 0,
		141, 142, 3, 32, 16, 0, 142, 143, 5, 7, 0, 0, 143, 144, 3, 38, 19, 0, 144,
		31, 1, 0, 0, 0, 145, 146, 7, 0, 0, 0, 146, 148, 5, 34, 0, 0, 147, 145,
		1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150, 3, 34,
		17, 0, 150, 33, 1, 0, 0, 0, 151, 153, 7, 1, 0, 0, 152, 154, 3, 52, 26,
		0, 153, 152, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 159, 1, 0, 0, 0, 155,
		156, 5, 34, 0, 0, 156, 158, 3, 34, 17, 0, 157, 155, 1, 0, 0, 0, 158, 161,
		1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 179, 1, 0,
		0, 0, 161, 159, 1, 0, 0, 0, 162, 167, 5, 33, 0, 0, 163, 164, 5, 34, 0,
		0, 164, 166, 3, 34, 17, 0, 165, 163, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0,
		167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 179, 1, 0, 0, 0, 169,
		167, 1, 0, 0, 0, 170, 175, 3, 36, 18, 0, 171, 172, 5, 34, 0, 0, 172, 174,
		3, 34, 17, 0, 173, 171, 1, 0, 0, 0, 174, 177, 1, 0, 0, 0, 175, 173, 1,
		0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0,
		0, 178, 151, 1, 0, 0, 0, 178, 162, 1, 0, 0, 0, 178, 170, 1, 0, 0, 0, 179,
		35, 1, 0, 0, 0, 180, 181, 5, 41, 0, 0, 181, 37, 1, 0, 0, 0, 182, 183, 6,
		19, -1, 0, 183, 190, 3, 42, 21, 0, 184, 185, 5, 37, 0, 0, 185, 186, 3,
		38, 19, 0, 186, 187, 5, 38, 0, 0, 187, 190, 1, 0, 0, 0, 188, 190, 3, 40,
		20, 0, 189, 182, 1, 0, 0, 0, 189, 184, 1, 0, 0, 0, 189, 188, 1, 0, 0, 0,
		190, 203, 1, 0, 0, 0, 191, 192, 10, 4, 0, 0, 192, 193, 3, 64, 32, 0, 193,
		194, 3, 38, 19, 5, 194, 202, 1, 0, 0, 0, 195, 196, 10, 2, 0, 0, 196, 197,
		5, 10, 0, 0, 197, 198, 3, 38, 19, 0, 198, 199, 5, 11, 0, 0, 199, 200, 3,
		38, 19, 3, 200, 202, 1, 0, 0, 0, 201, 191, 1, 0, 0, 0, 201, 195, 1, 0,
		0, 0, 202, 205, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0,
		204, 39, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 206, 207, 5, 12, 0, 0, 207,
		208, 3, 32, 16, 0, 208, 209, 5, 32, 0, 0, 209, 210, 3, 38, 19, 0, 210,
		211, 5, 11, 0, 0, 211, 212, 3, 38, 19, 0, 212, 41, 1, 0, 0, 0, 213, 217,
		3, 56, 28, 0, 214, 217, 3, 58, 29, 0, 215, 217, 3, 44, 22, 0, 216, 213,
		1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216, 215, 1, 0, 0, 0, 217, 43, 1, 0,
		0, 0, 218, 228, 5, 40, 0, 0, 219, 228, 3, 62, 31, 0, 220, 228, 3, 60, 30,
		0, 221, 228, 3, 32, 16, 0, 222, 228, 5, 42, 0, 0, 223, 228, 5, 43, 0, 0,
		224, 228, 3, 54, 27, 0, 225, 228, 3, 46, 23, 0, 226, 228, 5, 39, 0, 0,
		227, 218, 1, 0, 0, 0, 227, 219, 1, 0, 0, 0, 227, 220, 1, 0, 0, 0, 227,
		221, 1, 0, 0, 0, 227, 222, 1, 0, 0, 0, 227, 223, 1, 0, 0, 0, 227, 224,
		1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 227, 226, 1, 0, 0, 0, 228, 45, 1, 0,
		0, 0, 229, 230, 3, 48, 24, 0, 230, 231, 5, 37, 0, 0, 231, 232, 3, 50, 25,
		0, 232, 233, 5, 38, 0, 0, 233, 244, 1, 0, 0, 0, 234, 235, 5, 13, 0, 0,
		235, 239, 5, 37, 0, 0, 236, 238, 9, 0, 0, 0, 237, 236, 1, 0, 0, 0, 238,
		241, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 240, 242,
		1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 242, 244, 5, 38, 0, 0, 243, 229, 1, 0,
		0, 0, 243, 234, 1, 0, 0, 0, 244, 47, 1, 0, 0, 0, 245, 246, 5, 46, 0, 0,
		246, 49, 1, 0, 0, 0, 247, 257, 1, 0, 0, 0, 248, 253, 3, 38, 19, 0, 249,
		250, 5, 14, 0, 0, 250, 252, 3, 38, 19, 0, 251, 249, 1, 0, 0, 0, 252, 255,
		1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 257, 1, 0,
		0, 0, 255, 253, 1, 0, 0, 0, 256, 247, 1, 0, 0, 0, 256, 248, 1, 0, 0, 0,
		257, 51, 1, 0, 0, 0, 258, 259, 5, 15, 0, 0, 259, 260, 3, 38, 19, 0, 260,
		261, 5, 16, 0, 0, 261, 53, 1, 0, 0, 0, 262, 263, 5, 17, 0, 0, 263, 264,
		5, 37, 0, 0, 264, 265, 3, 38, 19, 0, 265, 266, 5, 38, 0, 0, 266, 55, 1,
		0, 0, 0, 267, 279, 5, 15, 0, 0, 268, 273, 3, 38, 19, 0, 269, 270, 5, 14,
		0, 0, 270, 272, 3, 38, 19, 0, 271, 269, 1, 0, 0, 0, 272, 275, 1, 0, 0,
		0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 277, 1, 0, 0, 0, 275,
		273, 1, 0, 0, 0, 276, 278, 5, 14, 0, 0, 277, 276, 1, 0, 0, 0, 277, 278,
		1, 0, 0, 0, 278, 280, 1, 0, 0, 0, 279, 268, 1, 0, 0, 0, 279, 280, 1, 0,
		0, 0, 280, 281, 1, 0, 0, 0, 281, 282, 5, 16, 0, 0, 282, 57, 1, 0, 0, 0,
		283, 290, 5, 35, 0, 0, 284, 286, 3, 30, 15, 0, 285, 287, 5, 14, 0, 0, 286,
		285, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 289, 1, 0, 0, 0, 288, 284,
		1, 0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 291, 1, 0,
		0, 0, 291, 293, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 293, 294, 5, 36, 0, 0,
		294, 59, 1, 0, 0, 0, 295, 296, 7, 2, 0, 0, 296, 61, 1, 0, 0, 0, 297, 299,
		7, 3, 0, 0, 298, 297, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 300, 1, 0,
		0, 0, 300, 301, 3, 66, 33, 0, 301, 63, 1, 0, 0, 0, 302, 303, 7, 4, 0, 0,
		303, 65, 1, 0, 0, 0, 304, 307, 5, 41, 0, 0, 305, 306, 5, 34, 0, 0, 306,
		308, 5, 41, 0, 0, 307, 305, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 67,
		1, 0, 0, 0, 27, 76, 78, 118, 134, 136, 147, 153, 159, 167, 175, 178, 189,
		201, 203, 216, 227, 239, 243, 253, 256, 273, 277, 279, 286, 290, 298, 307,
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

// TerraformParserInit initializes any static state used to implement TerraformParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTerraformParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TerraformParserInit() {
	staticData := &terraformParserStaticData
	staticData.once.Do(terraformParserInit)
}

// NewTerraformParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTerraformParser(input antlr.TokenStream) *TerraformParser {
	TerraformParserInit()
	this := new(TerraformParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &terraformParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// TerraformParser tokens.
const (
	TerraformParserEOF             = antlr.TokenEOF
	TerraformParserT__0            = 1
	TerraformParserT__1            = 2
	TerraformParserT__2            = 3
	TerraformParserT__3            = 4
	TerraformParserT__4            = 5
	TerraformParserT__5            = 6
	TerraformParserT__6            = 7
	TerraformParserT__7            = 8
	TerraformParserT__8            = 9
	TerraformParserT__9            = 10
	TerraformParserT__10           = 11
	TerraformParserT__11           = 12
	TerraformParserT__12           = 13
	TerraformParserT__13           = 14
	TerraformParserT__14           = 15
	TerraformParserT__15           = 16
	TerraformParserT__16           = 17
	TerraformParserT__17           = 18
	TerraformParserT__18           = 19
	TerraformParserT__19           = 20
	TerraformParserT__20           = 21
	TerraformParserT__21           = 22
	TerraformParserT__22           = 23
	TerraformParserT__23           = 24
	TerraformParserT__24           = 25
	TerraformParserT__25           = 26
	TerraformParserT__26           = 27
	TerraformParserT__27           = 28
	TerraformParserT__28           = 29
	TerraformParserVARIABLE        = 30
	TerraformParserPROVIDER        = 31
	TerraformParserIN              = 32
	TerraformParserSTAR            = 33
	TerraformParserDOT             = 34
	TerraformParserLCURL           = 35
	TerraformParserRCURL           = 36
	TerraformParserLPAREN          = 37
	TerraformParserRPAREN          = 38
	TerraformParserEOF_            = 39
	TerraformParserNULL_           = 40
	TerraformParserNATURAL_NUMBER  = 41
	TerraformParserBOOL            = 42
	TerraformParserDESCRIPTION     = 43
	TerraformParserMULTILINESTRING = 44
	TerraformParserSTRING          = 45
	TerraformParserIDENTIFIER      = 46
	TerraformParserCOMMENT         = 47
	TerraformParserBLOCKCOMMENT    = 48
	TerraformParserWS              = 49
)

// TerraformParser rules.
const (
	TerraformParserRULE_file_             = 0
	TerraformParserRULE_terraform         = 1
	TerraformParserRULE_resource          = 2
	TerraformParserRULE_data              = 3
	TerraformParserRULE_provider          = 4
	TerraformParserRULE_output            = 5
	TerraformParserRULE_local             = 6
	TerraformParserRULE_module            = 7
	TerraformParserRULE_variable          = 8
	TerraformParserRULE_block             = 9
	TerraformParserRULE_blocktype         = 10
	TerraformParserRULE_resourcetype      = 11
	TerraformParserRULE_name              = 12
	TerraformParserRULE_label             = 13
	TerraformParserRULE_blockbody         = 14
	TerraformParserRULE_argument          = 15
	TerraformParserRULE_identifier        = 16
	TerraformParserRULE_identifierchain   = 17
	TerraformParserRULE_inline_index      = 18
	TerraformParserRULE_expression        = 19
	TerraformParserRULE_forloop           = 20
	TerraformParserRULE_section           = 21
	TerraformParserRULE_val               = 22
	TerraformParserRULE_functioncall      = 23
	TerraformParserRULE_functionname      = 24
	TerraformParserRULE_functionarguments = 25
	TerraformParserRULE_index             = 26
	TerraformParserRULE_filedecl          = 27
	TerraformParserRULE_list_             = 28
	TerraformParserRULE_map_              = 29
	TerraformParserRULE_string            = 30
	TerraformParserRULE_signed_number     = 31
	TerraformParserRULE_operator_         = 32
	TerraformParserRULE_number            = 33
)

// IFile_Context is an interface to support dynamic dispatch.
type IFile_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFile_Context differentiates from other interfaces.
	IsFile_Context()
}

type File_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFile_Context() *File_Context {
	var p = new(File_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_file_
	return p
}

func (*File_Context) IsFile_Context() {}

func NewFile_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *File_Context {
	var p = new(File_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_file_

	return p
}

func (s *File_Context) GetParser() antlr.Parser { return s.parser }

func (s *File_Context) EOF() antlr.TerminalNode {
	return s.GetToken(TerraformParserEOF, 0)
}

func (s *File_Context) AllLocal() []ILocalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILocalContext); ok {
			len++
		}
	}

	tst := make([]ILocalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILocalContext); ok {
			tst[i] = t.(ILocalContext)
			i++
		}
	}

	return tst
}

func (s *File_Context) Local(i int) ILocalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalContext); ok {
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

	return t.(ILocalContext)
}

func (s *File_Context) AllModule() []IModuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IModuleContext); ok {
			len++
		}
	}

	tst := make([]IModuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IModuleContext); ok {
			tst[i] = t.(IModuleContext)
			i++
		}
	}

	return tst
}

func (s *File_Context) Module(i int) IModuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModuleContext); ok {
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

	return t.(IModuleContext)
}

func (s *File_Context) AllOutput() []IOutputContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOutputContext); ok {
			len++
		}
	}

	tst := make([]IOutputContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOutputContext); ok {
			tst[i] = t.(IOutputContext)
			i++
		}
	}

	return tst
}

func (s *File_Context) Output(i int) IOutputContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutputContext); ok {
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

	return t.(IOutputContext)
}

func (s *File_Context) AllProvider() []IProviderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProviderContext); ok {
			len++
		}
	}

	tst := make([]IProviderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProviderContext); ok {
			tst[i] = t.(IProviderContext)
			i++
		}
	}

	return tst
}

func (s *File_Context) Provider(i int) IProviderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProviderContext); ok {
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

	return t.(IProviderContext)
}

func (s *File_Context) AllVariable() []IVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableContext); ok {
			len++
		}
	}

	tst := make([]IVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableContext); ok {
			tst[i] = t.(IVariableContext)
			i++
		}
	}

	return tst
}

func (s *File_Context) Variable(i int) IVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
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

	return t.(IVariableContext)
}

func (s *File_Context) AllData() []IDataContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDataContext); ok {
			len++
		}
	}

	tst := make([]IDataContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDataContext); ok {
			tst[i] = t.(IDataContext)
			i++
		}
	}

	return tst
}

func (s *File_Context) Data(i int) IDataContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataContext); ok {
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

	return t.(IDataContext)
}

func (s *File_Context) AllResource() []IResourceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IResourceContext); ok {
			len++
		}
	}

	tst := make([]IResourceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IResourceContext); ok {
			tst[i] = t.(IResourceContext)
			i++
		}
	}

	return tst
}

func (s *File_Context) Resource(i int) IResourceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourceContext); ok {
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

	return t.(IResourceContext)
}

func (s *File_Context) AllTerraform() []ITerraformContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerraformContext); ok {
			len++
		}
	}

	tst := make([]ITerraformContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerraformContext); ok {
			tst[i] = t.(ITerraformContext)
			i++
		}
	}

	return tst
}

func (s *File_Context) Terraform(i int) ITerraformContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerraformContext); ok {
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

	return t.(ITerraformContext)
}

func (s *File_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *File_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *File_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterFile_(s)
	}
}

func (s *File_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitFile_(s)
	}
}

func (p *TerraformParser) File_() (localctx IFile_Context) {
	this := p
	_ = this

	localctx = NewFile_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TerraformParserRULE_file_)
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
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3221225598) != 0 {
		p.SetState(76)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case TerraformParserT__4:
			{
				p.SetState(68)
				p.Local()
			}

		case TerraformParserT__5:
			{
				p.SetState(69)
				p.Module()
			}

		case TerraformParserT__3:
			{
				p.SetState(70)
				p.Output()
			}

		case TerraformParserPROVIDER:
			{
				p.SetState(71)
				p.Provider()
			}

		case TerraformParserVARIABLE:
			{
				p.SetState(72)
				p.Variable()
			}

		case TerraformParserT__2:
			{
				p.SetState(73)
				p.Data()
			}

		case TerraformParserT__1:
			{
				p.SetState(74)
				p.Resource()
			}

		case TerraformParserT__0:
			{
				p.SetState(75)
				p.Terraform()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(80)
		p.Match(TerraformParserEOF)
	}

	return localctx
}

// ITerraformContext is an interface to support dynamic dispatch.
type ITerraformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTerraformContext differentiates from other interfaces.
	IsTerraformContext()
}

type TerraformContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerraformContext() *TerraformContext {
	var p = new(TerraformContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_terraform
	return p
}

func (*TerraformContext) IsTerraformContext() {}

func NewTerraformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TerraformContext {
	var p = new(TerraformContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_terraform

	return p
}

func (s *TerraformContext) GetParser() antlr.Parser { return s.parser }

func (s *TerraformContext) Blockbody() IBlockbodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockbodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockbodyContext)
}

func (s *TerraformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TerraformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TerraformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterTerraform(s)
	}
}

func (s *TerraformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitTerraform(s)
	}
}

func (p *TerraformParser) Terraform() (localctx ITerraformContext) {
	this := p
	_ = this

	localctx = NewTerraformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TerraformParserRULE_terraform)

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
		p.SetState(82)
		p.Match(TerraformParserT__0)
	}
	{
		p.SetState(83)
		p.Blockbody()
	}

	return localctx
}

// IResourceContext is an interface to support dynamic dispatch.
type IResourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResourceContext differentiates from other interfaces.
	IsResourceContext()
}

type ResourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResourceContext() *ResourceContext {
	var p = new(ResourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_resource
	return p
}

func (*ResourceContext) IsResourceContext() {}

func NewResourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourceContext {
	var p = new(ResourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_resource

	return p
}

func (s *ResourceContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourceContext) Resourcetype() IResourcetypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourcetypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResourcetypeContext)
}

func (s *ResourceContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ResourceContext) Blockbody() IBlockbodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockbodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockbodyContext)
}

func (s *ResourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterResource(s)
	}
}

func (s *ResourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitResource(s)
	}
}

func (p *TerraformParser) Resource() (localctx IResourceContext) {
	this := p
	_ = this

	localctx = NewResourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TerraformParserRULE_resource)

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
		p.SetState(85)
		p.Match(TerraformParserT__1)
	}
	{
		p.SetState(86)
		p.Resourcetype()
	}
	{
		p.SetState(87)
		p.Name()
	}
	{
		p.SetState(88)
		p.Blockbody()
	}

	return localctx
}

// IDataContext is an interface to support dynamic dispatch.
type IDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataContext differentiates from other interfaces.
	IsDataContext()
}

type DataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataContext() *DataContext {
	var p = new(DataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_data
	return p
}

func (*DataContext) IsDataContext() {}

func NewDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataContext {
	var p = new(DataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_data

	return p
}

func (s *DataContext) GetParser() antlr.Parser { return s.parser }

func (s *DataContext) Resourcetype() IResourcetypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourcetypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResourcetypeContext)
}

func (s *DataContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DataContext) Blockbody() IBlockbodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockbodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockbodyContext)
}

func (s *DataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterData(s)
	}
}

func (s *DataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitData(s)
	}
}

func (p *TerraformParser) Data() (localctx IDataContext) {
	this := p
	_ = this

	localctx = NewDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TerraformParserRULE_data)

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
		p.SetState(90)
		p.Match(TerraformParserT__2)
	}
	{
		p.SetState(91)
		p.Resourcetype()
	}
	{
		p.SetState(92)
		p.Name()
	}
	{
		p.SetState(93)
		p.Blockbody()
	}

	return localctx
}

// IProviderContext is an interface to support dynamic dispatch.
type IProviderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProviderContext differentiates from other interfaces.
	IsProviderContext()
}

type ProviderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProviderContext() *ProviderContext {
	var p = new(ProviderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_provider
	return p
}

func (*ProviderContext) IsProviderContext() {}

func NewProviderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProviderContext {
	var p = new(ProviderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_provider

	return p
}

func (s *ProviderContext) GetParser() antlr.Parser { return s.parser }

func (s *ProviderContext) PROVIDER() antlr.TerminalNode {
	return s.GetToken(TerraformParserPROVIDER, 0)
}

func (s *ProviderContext) Resourcetype() IResourcetypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourcetypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResourcetypeContext)
}

func (s *ProviderContext) Blockbody() IBlockbodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockbodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockbodyContext)
}

func (s *ProviderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProviderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProviderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterProvider(s)
	}
}

func (s *ProviderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitProvider(s)
	}
}

func (p *TerraformParser) Provider() (localctx IProviderContext) {
	this := p
	_ = this

	localctx = NewProviderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TerraformParserRULE_provider)

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
		p.SetState(95)
		p.Match(TerraformParserPROVIDER)
	}
	{
		p.SetState(96)
		p.Resourcetype()
	}
	{
		p.SetState(97)
		p.Blockbody()
	}

	return localctx
}

// IOutputContext is an interface to support dynamic dispatch.
type IOutputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOutputContext differentiates from other interfaces.
	IsOutputContext()
}

type OutputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutputContext() *OutputContext {
	var p = new(OutputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_output
	return p
}

func (*OutputContext) IsOutputContext() {}

func NewOutputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputContext {
	var p = new(OutputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_output

	return p
}

func (s *OutputContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *OutputContext) Blockbody() IBlockbodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockbodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockbodyContext)
}

func (s *OutputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterOutput(s)
	}
}

func (s *OutputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitOutput(s)
	}
}

func (p *TerraformParser) Output() (localctx IOutputContext) {
	this := p
	_ = this

	localctx = NewOutputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TerraformParserRULE_output)

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
		p.SetState(99)
		p.Match(TerraformParserT__3)
	}
	{
		p.SetState(100)
		p.Name()
	}
	{
		p.SetState(101)
		p.Blockbody()
	}

	return localctx
}

// ILocalContext is an interface to support dynamic dispatch.
type ILocalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalContext differentiates from other interfaces.
	IsLocalContext()
}

type LocalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalContext() *LocalContext {
	var p = new(LocalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_local
	return p
}

func (*LocalContext) IsLocalContext() {}

func NewLocalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalContext {
	var p = new(LocalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_local

	return p
}

func (s *LocalContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalContext) Blockbody() IBlockbodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockbodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockbodyContext)
}

func (s *LocalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterLocal(s)
	}
}

func (s *LocalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitLocal(s)
	}
}

func (p *TerraformParser) Local() (localctx ILocalContext) {
	this := p
	_ = this

	localctx = NewLocalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TerraformParserRULE_local)

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
		p.SetState(103)
		p.Match(TerraformParserT__4)
	}
	{
		p.SetState(104)
		p.Blockbody()
	}

	return localctx
}

// IModuleContext is an interface to support dynamic dispatch.
type IModuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleContext differentiates from other interfaces.
	IsModuleContext()
}

type ModuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_module
	return p
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ModuleContext) Blockbody() IBlockbodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockbodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockbodyContext)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterModule(s)
	}
}

func (s *ModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitModule(s)
	}
}

func (p *TerraformParser) Module() (localctx IModuleContext) {
	this := p
	_ = this

	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TerraformParserRULE_module)

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
		p.SetState(106)
		p.Match(TerraformParserT__5)
	}
	{
		p.SetState(107)
		p.Name()
	}
	{
		p.SetState(108)
		p.Blockbody()
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(TerraformParserVARIABLE, 0)
}

func (s *VariableContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *VariableContext) Blockbody() IBlockbodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockbodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockbodyContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *TerraformParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TerraformParserRULE_variable)

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
		p.SetState(110)
		p.Match(TerraformParserVARIABLE)
	}
	{
		p.SetState(111)
		p.Name()
	}
	{
		p.SetState(112)
		p.Blockbody()
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Blocktype() IBlocktypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlocktypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlocktypeContext)
}

func (s *BlockContext) Blockbody() IBlockbodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockbodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockbodyContext)
}

func (s *BlockContext) AllLabel() []ILabelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILabelContext); ok {
			len++
		}
	}

	tst := make([]ILabelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILabelContext); ok {
			tst[i] = t.(ILabelContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Label(i int) ILabelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelContext); ok {
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

	return t.(ILabelContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *TerraformParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TerraformParserRULE_block)
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
		p.SetState(114)
		p.Blocktype()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TerraformParserSTRING {
		{
			p.SetState(115)
			p.Label()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(121)
		p.Blockbody()
	}

	return localctx
}

// IBlocktypeContext is an interface to support dynamic dispatch.
type IBlocktypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlocktypeContext differentiates from other interfaces.
	IsBlocktypeContext()
}

type BlocktypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlocktypeContext() *BlocktypeContext {
	var p = new(BlocktypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_blocktype
	return p
}

func (*BlocktypeContext) IsBlocktypeContext() {}

func NewBlocktypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlocktypeContext {
	var p = new(BlocktypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_blocktype

	return p
}

func (s *BlocktypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BlocktypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TerraformParserIDENTIFIER, 0)
}

func (s *BlocktypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlocktypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlocktypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterBlocktype(s)
	}
}

func (s *BlocktypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitBlocktype(s)
	}
}

func (p *TerraformParser) Blocktype() (localctx IBlocktypeContext) {
	this := p
	_ = this

	localctx = NewBlocktypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TerraformParserRULE_blocktype)

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
		p.SetState(123)
		p.Match(TerraformParserIDENTIFIER)
	}

	return localctx
}

// IResourcetypeContext is an interface to support dynamic dispatch.
type IResourcetypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResourcetypeContext differentiates from other interfaces.
	IsResourcetypeContext()
}

type ResourcetypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResourcetypeContext() *ResourcetypeContext {
	var p = new(ResourcetypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_resourcetype
	return p
}

func (*ResourcetypeContext) IsResourcetypeContext() {}

func NewResourcetypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourcetypeContext {
	var p = new(ResourcetypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_resourcetype

	return p
}

func (s *ResourcetypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourcetypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(TerraformParserSTRING, 0)
}

func (s *ResourcetypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourcetypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourcetypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterResourcetype(s)
	}
}

func (s *ResourcetypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitResourcetype(s)
	}
}

func (p *TerraformParser) Resourcetype() (localctx IResourcetypeContext) {
	this := p
	_ = this

	localctx = NewResourcetypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TerraformParserRULE_resourcetype)

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
		p.SetState(125)
		p.Match(TerraformParserSTRING)
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) STRING() antlr.TerminalNode {
	return s.GetToken(TerraformParserSTRING, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *TerraformParser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TerraformParserRULE_name)

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
		p.SetState(127)
		p.Match(TerraformParserSTRING)
	}

	return localctx
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) STRING() antlr.TerminalNode {
	return s.GetToken(TerraformParserSTRING, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (p *TerraformParser) Label() (localctx ILabelContext) {
	this := p
	_ = this

	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TerraformParserRULE_label)

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
		p.SetState(129)
		p.Match(TerraformParserSTRING)
	}

	return localctx
}

// IBlockbodyContext is an interface to support dynamic dispatch.
type IBlockbodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockbodyContext differentiates from other interfaces.
	IsBlockbodyContext()
}

type BlockbodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockbodyContext() *BlockbodyContext {
	var p = new(BlockbodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_blockbody
	return p
}

func (*BlockbodyContext) IsBlockbodyContext() {}

func NewBlockbodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockbodyContext {
	var p = new(BlockbodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_blockbody

	return p
}

func (s *BlockbodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockbodyContext) LCURL() antlr.TerminalNode {
	return s.GetToken(TerraformParserLCURL, 0)
}

func (s *BlockbodyContext) RCURL() antlr.TerminalNode {
	return s.GetToken(TerraformParserRCURL, 0)
}

func (s *BlockbodyContext) AllArgument() []IArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentContext); ok {
			len++
		}
	}

	tst := make([]IArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentContext); ok {
			tst[i] = t.(IArgumentContext)
			i++
		}
	}

	return tst
}

func (s *BlockbodyContext) Argument(i int) IArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
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

	return t.(IArgumentContext)
}

func (s *BlockbodyContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *BlockbodyContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *BlockbodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockbodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockbodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterBlockbody(s)
	}
}

func (s *BlockbodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitBlockbody(s)
	}
}

func (p *TerraformParser) Blockbody() (localctx IBlockbodyContext) {
	this := p
	_ = this

	localctx = NewBlockbodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TerraformParserRULE_blockbody)
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
		p.SetState(131)
		p.Match(TerraformParserLCURL)
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72583873561416) != 0 {
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(132)
				p.Argument()
			}

		case 2:
			{
				p.SetState(133)
				p.Block()
			}

		}

		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(139)
		p.Match(TerraformParserRCURL)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ArgumentContext) Expression() IExpressionContext {
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

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *TerraformParser) Argument() (localctx IArgumentContext) {
	this := p
	_ = this

	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TerraformParserRULE_argument)

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
		p.SetState(141)
		p.Identifier()
	}
	{
		p.SetState(142)
		p.Match(TerraformParserT__6)
	}
	{
		p.SetState(143)
		p.expression(0)
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) Identifierchain() IIdentifierchainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierchainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierchainContext)
}

func (s *IdentifierContext) DOT() antlr.TerminalNode {
	return s.GetToken(TerraformParserDOT, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *TerraformParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TerraformParserRULE_identifier)
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
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&840) != 0 {
		{
			p.SetState(145)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&840) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(146)
			p.Match(TerraformParserDOT)
		}

	}
	{
		p.SetState(149)
		p.Identifierchain()
	}

	return localctx
}

// IIdentifierchainContext is an interface to support dynamic dispatch.
type IIdentifierchainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierchainContext differentiates from other interfaces.
	IsIdentifierchainContext()
}

type IdentifierchainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierchainContext() *IdentifierchainContext {
	var p = new(IdentifierchainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_identifierchain
	return p
}

func (*IdentifierchainContext) IsIdentifierchainContext() {}

func NewIdentifierchainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierchainContext {
	var p = new(IdentifierchainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_identifierchain

	return p
}

func (s *IdentifierchainContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierchainContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TerraformParserIDENTIFIER, 0)
}

func (s *IdentifierchainContext) IN() antlr.TerminalNode {
	return s.GetToken(TerraformParserIN, 0)
}

func (s *IdentifierchainContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(TerraformParserVARIABLE, 0)
}

func (s *IdentifierchainContext) PROVIDER() antlr.TerminalNode {
	return s.GetToken(TerraformParserPROVIDER, 0)
}

func (s *IdentifierchainContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *IdentifierchainContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TerraformParserDOT)
}

func (s *IdentifierchainContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TerraformParserDOT, i)
}

func (s *IdentifierchainContext) AllIdentifierchain() []IIdentifierchainContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierchainContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierchainContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierchainContext); ok {
			tst[i] = t.(IIdentifierchainContext)
			i++
		}
	}

	return tst
}

func (s *IdentifierchainContext) Identifierchain(i int) IIdentifierchainContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierchainContext); ok {
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

	return t.(IIdentifierchainContext)
}

func (s *IdentifierchainContext) STAR() antlr.TerminalNode {
	return s.GetToken(TerraformParserSTAR, 0)
}

func (s *IdentifierchainContext) Inline_index() IInline_indexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInline_indexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInline_indexContext)
}

func (s *IdentifierchainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierchainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierchainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterIdentifierchain(s)
	}
}

func (s *IdentifierchainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitIdentifierchain(s)
	}
}

func (p *TerraformParser) Identifierchain() (localctx IIdentifierchainContext) {
	this := p
	_ = this

	localctx = NewIdentifierchainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TerraformParserRULE_identifierchain)
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

	p.SetState(178)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TerraformParserVARIABLE, TerraformParserPROVIDER, TerraformParserIN, TerraformParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70376260370432) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(152)
				p.Index()
			}

		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(155)
					p.Match(TerraformParserDOT)
				}
				{
					p.SetState(156)
					p.Identifierchain()
				}

			}
			p.SetState(161)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
		}

	case TerraformParserSTAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Match(TerraformParserSTAR)
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(163)
					p.Match(TerraformParserDOT)
				}
				{
					p.SetState(164)
					p.Identifierchain()
				}

			}
			p.SetState(169)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
		}

	case TerraformParserNATURAL_NUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(170)
			p.Inline_index()
		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(171)
					p.Match(TerraformParserDOT)
				}
				{
					p.SetState(172)
					p.Identifierchain()
				}

			}
			p.SetState(177)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInline_indexContext is an interface to support dynamic dispatch.
type IInline_indexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInline_indexContext differentiates from other interfaces.
	IsInline_indexContext()
}

type Inline_indexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInline_indexContext() *Inline_indexContext {
	var p = new(Inline_indexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_inline_index
	return p
}

func (*Inline_indexContext) IsInline_indexContext() {}

func NewInline_indexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inline_indexContext {
	var p = new(Inline_indexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_inline_index

	return p
}

func (s *Inline_indexContext) GetParser() antlr.Parser { return s.parser }

func (s *Inline_indexContext) NATURAL_NUMBER() antlr.TerminalNode {
	return s.GetToken(TerraformParserNATURAL_NUMBER, 0)
}

func (s *Inline_indexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inline_indexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inline_indexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterInline_index(s)
	}
}

func (s *Inline_indexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitInline_index(s)
	}
}

func (p *TerraformParser) Inline_index() (localctx IInline_indexContext) {
	this := p
	_ = this

	localctx = NewInline_indexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TerraformParserRULE_inline_index)

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
		p.SetState(180)
		p.Match(TerraformParserNATURAL_NUMBER)
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
	p.RuleIndex = TerraformParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Section() ISectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISectionContext)
}

func (s *ExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TerraformParserLPAREN, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
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

func (s *ExpressionContext) Expression(i int) IExpressionContext {
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

func (s *ExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TerraformParserRPAREN, 0)
}

func (s *ExpressionContext) Forloop() IForloopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForloopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForloopContext)
}

func (s *ExpressionContext) Operator_() IOperator_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperator_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperator_Context)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TerraformParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *TerraformParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, TerraformParserRULE_expression, _p)

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
	p.SetState(189)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TerraformParserT__2, TerraformParserT__5, TerraformParserT__7, TerraformParserT__8, TerraformParserT__12, TerraformParserT__14, TerraformParserT__16, TerraformParserT__17, TerraformParserT__18, TerraformParserVARIABLE, TerraformParserPROVIDER, TerraformParserIN, TerraformParserSTAR, TerraformParserLCURL, TerraformParserEOF_, TerraformParserNULL_, TerraformParserNATURAL_NUMBER, TerraformParserBOOL, TerraformParserDESCRIPTION, TerraformParserMULTILINESTRING, TerraformParserSTRING, TerraformParserIDENTIFIER:
		{
			p.SetState(183)
			p.Section()
		}

	case TerraformParserLPAREN:
		{
			p.SetState(184)
			p.Match(TerraformParserLPAREN)
		}
		{
			p.SetState(185)
			p.expression(0)
		}
		{
			p.SetState(186)
			p.Match(TerraformParserRPAREN)
		}

	case TerraformParserT__11:
		{
			p.SetState(188)
			p.Forloop()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TerraformParserRULE_expression)
				p.SetState(191)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(192)
					p.Operator_()
				}
				{
					p.SetState(193)
					p.expression(5)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TerraformParserRULE_expression)
				p.SetState(195)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(196)
					p.Match(TerraformParserT__9)
				}
				{
					p.SetState(197)
					p.expression(0)
				}
				{
					p.SetState(198)
					p.Match(TerraformParserT__10)
				}
				{
					p.SetState(199)
					p.expression(3)
				}

			}

		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IForloopContext is an interface to support dynamic dispatch.
type IForloopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForloopContext differentiates from other interfaces.
	IsForloopContext()
}

type ForloopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForloopContext() *ForloopContext {
	var p = new(ForloopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_forloop
	return p
}

func (*ForloopContext) IsForloopContext() {}

func NewForloopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForloopContext {
	var p = new(ForloopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_forloop

	return p
}

func (s *ForloopContext) GetParser() antlr.Parser { return s.parser }

func (s *ForloopContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ForloopContext) IN() antlr.TerminalNode {
	return s.GetToken(TerraformParserIN, 0)
}

func (s *ForloopContext) AllExpression() []IExpressionContext {
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

func (s *ForloopContext) Expression(i int) IExpressionContext {
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

func (s *ForloopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForloopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForloopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterForloop(s)
	}
}

func (s *ForloopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitForloop(s)
	}
}

func (p *TerraformParser) Forloop() (localctx IForloopContext) {
	this := p
	_ = this

	localctx = NewForloopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, TerraformParserRULE_forloop)

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
		p.SetState(206)
		p.Match(TerraformParserT__11)
	}
	{
		p.SetState(207)
		p.Identifier()
	}
	{
		p.SetState(208)
		p.Match(TerraformParserIN)
	}
	{
		p.SetState(209)
		p.expression(0)
	}
	{
		p.SetState(210)
		p.Match(TerraformParserT__10)
	}
	{
		p.SetState(211)
		p.expression(0)
	}

	return localctx
}

// ISectionContext is an interface to support dynamic dispatch.
type ISectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionContext differentiates from other interfaces.
	IsSectionContext()
}

type SectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionContext() *SectionContext {
	var p = new(SectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_section
	return p
}

func (*SectionContext) IsSectionContext() {}

func NewSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionContext {
	var p = new(SectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_section

	return p
}

func (s *SectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionContext) List_() IList_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_Context)
}

func (s *SectionContext) Map_() IMap_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMap_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMap_Context)
}

func (s *SectionContext) Val() IValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *SectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterSection(s)
	}
}

func (s *SectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitSection(s)
	}
}

func (p *TerraformParser) Section() (localctx ISectionContext) {
	this := p
	_ = this

	localctx = NewSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, TerraformParserRULE_section)

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

	p.SetState(216)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TerraformParserT__14:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(213)
			p.List_()
		}

	case TerraformParserLCURL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
			p.Map_()
		}

	case TerraformParserT__2, TerraformParserT__5, TerraformParserT__7, TerraformParserT__8, TerraformParserT__12, TerraformParserT__16, TerraformParserT__17, TerraformParserT__18, TerraformParserVARIABLE, TerraformParserPROVIDER, TerraformParserIN, TerraformParserSTAR, TerraformParserEOF_, TerraformParserNULL_, TerraformParserNATURAL_NUMBER, TerraformParserBOOL, TerraformParserDESCRIPTION, TerraformParserMULTILINESTRING, TerraformParserSTRING, TerraformParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(215)
			p.Val()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValContext is an interface to support dynamic dispatch.
type IValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValContext differentiates from other interfaces.
	IsValContext()
}

type ValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValContext() *ValContext {
	var p = new(ValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_val
	return p
}

func (*ValContext) IsValContext() {}

func NewValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValContext {
	var p = new(ValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_val

	return p
}

func (s *ValContext) GetParser() antlr.Parser { return s.parser }

func (s *ValContext) NULL_() antlr.TerminalNode {
	return s.GetToken(TerraformParserNULL_, 0)
}

func (s *ValContext) Signed_number() ISigned_numberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISigned_numberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISigned_numberContext)
}

func (s *ValContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *ValContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ValContext) BOOL() antlr.TerminalNode {
	return s.GetToken(TerraformParserBOOL, 0)
}

func (s *ValContext) DESCRIPTION() antlr.TerminalNode {
	return s.GetToken(TerraformParserDESCRIPTION, 0)
}

func (s *ValContext) Filedecl() IFiledeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFiledeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFiledeclContext)
}

func (s *ValContext) Functioncall() IFunctioncallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctioncallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctioncallContext)
}

func (s *ValContext) EOF_() antlr.TerminalNode {
	return s.GetToken(TerraformParserEOF_, 0)
}

func (s *ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterVal(s)
	}
}

func (s *ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitVal(s)
	}
}

func (p *TerraformParser) Val() (localctx IValContext) {
	this := p
	_ = this

	localctx = NewValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TerraformParserRULE_val)

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

	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(218)
			p.Match(TerraformParserNULL_)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(219)
			p.Signed_number()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(220)
			p.String_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(221)
			p.Identifier()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(222)
			p.Match(TerraformParserBOOL)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(223)
			p.Match(TerraformParserDESCRIPTION)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(224)
			p.Filedecl()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(225)
			p.Functioncall()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(226)
			p.Match(TerraformParserEOF_)
		}

	}

	return localctx
}

// IFunctioncallContext is an interface to support dynamic dispatch.
type IFunctioncallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctioncallContext differentiates from other interfaces.
	IsFunctioncallContext()
}

type FunctioncallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctioncallContext() *FunctioncallContext {
	var p = new(FunctioncallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_functioncall
	return p
}

func (*FunctioncallContext) IsFunctioncallContext() {}

func NewFunctioncallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctioncallContext {
	var p = new(FunctioncallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_functioncall

	return p
}

func (s *FunctioncallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctioncallContext) Functionname() IFunctionnameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionnameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionnameContext)
}

func (s *FunctioncallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TerraformParserLPAREN, 0)
}

func (s *FunctioncallContext) Functionarguments() IFunctionargumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionargumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionargumentsContext)
}

func (s *FunctioncallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TerraformParserRPAREN, 0)
}

func (s *FunctioncallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctioncallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctioncallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterFunctioncall(s)
	}
}

func (s *FunctioncallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitFunctioncall(s)
	}
}

func (p *TerraformParser) Functioncall() (localctx IFunctioncallContext) {
	this := p
	_ = this

	localctx = NewFunctioncallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TerraformParserRULE_functioncall)

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

	p.SetState(243)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TerraformParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.Functionname()
		}
		{
			p.SetState(230)
			p.Match(TerraformParserLPAREN)
		}
		{
			p.SetState(231)
			p.Functionarguments()
		}
		{
			p.SetState(232)
			p.Match(TerraformParserRPAREN)
		}

	case TerraformParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(234)
			p.Match(TerraformParserT__12)
		}
		{
			p.SetState(235)
			p.Match(TerraformParserLPAREN)
		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

		for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1+1 {
				p.SetState(236)
				p.MatchWildcard()

			}
			p.SetState(241)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
		}
		{
			p.SetState(242)
			p.Match(TerraformParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionnameContext is an interface to support dynamic dispatch.
type IFunctionnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionnameContext differentiates from other interfaces.
	IsFunctionnameContext()
}

type FunctionnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionnameContext() *FunctionnameContext {
	var p = new(FunctionnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_functionname
	return p
}

func (*FunctionnameContext) IsFunctionnameContext() {}

func NewFunctionnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionnameContext {
	var p = new(FunctionnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_functionname

	return p
}

func (s *FunctionnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionnameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TerraformParserIDENTIFIER, 0)
}

func (s *FunctionnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterFunctionname(s)
	}
}

func (s *FunctionnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitFunctionname(s)
	}
}

func (p *TerraformParser) Functionname() (localctx IFunctionnameContext) {
	this := p
	_ = this

	localctx = NewFunctionnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, TerraformParserRULE_functionname)

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
		p.Match(TerraformParserIDENTIFIER)
	}

	return localctx
}

// IFunctionargumentsContext is an interface to support dynamic dispatch.
type IFunctionargumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionargumentsContext differentiates from other interfaces.
	IsFunctionargumentsContext()
}

type FunctionargumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionargumentsContext() *FunctionargumentsContext {
	var p = new(FunctionargumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_functionarguments
	return p
}

func (*FunctionargumentsContext) IsFunctionargumentsContext() {}

func NewFunctionargumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionargumentsContext {
	var p = new(FunctionargumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_functionarguments

	return p
}

func (s *FunctionargumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionargumentsContext) AllExpression() []IExpressionContext {
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

func (s *FunctionargumentsContext) Expression(i int) IExpressionContext {
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

func (s *FunctionargumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionargumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionargumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterFunctionarguments(s)
	}
}

func (s *FunctionargumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitFunctionarguments(s)
	}
}

func (p *TerraformParser) Functionarguments() (localctx IFunctionargumentsContext) {
	this := p
	_ = this

	localctx = NewFunctionargumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, TerraformParserRULE_functionarguments)
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

	p.SetState(256)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TerraformParserRPAREN:
		p.EnterOuterAlt(localctx, 1)

	case TerraformParserT__2, TerraformParserT__5, TerraformParserT__7, TerraformParserT__8, TerraformParserT__11, TerraformParserT__12, TerraformParserT__14, TerraformParserT__16, TerraformParserT__17, TerraformParserT__18, TerraformParserVARIABLE, TerraformParserPROVIDER, TerraformParserIN, TerraformParserSTAR, TerraformParserLCURL, TerraformParserLPAREN, TerraformParserEOF_, TerraformParserNULL_, TerraformParserNATURAL_NUMBER, TerraformParserBOOL, TerraformParserDESCRIPTION, TerraformParserMULTILINESTRING, TerraformParserSTRING, TerraformParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(248)
			p.expression(0)
		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TerraformParserT__13 {
			{
				p.SetState(249)
				p.Match(TerraformParserT__13)
			}
			{
				p.SetState(250)
				p.expression(0)
			}

			p.SetState(255)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_index
	return p
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) Expression() IExpressionContext {
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

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (p *TerraformParser) Index() (localctx IIndexContext) {
	this := p
	_ = this

	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, TerraformParserRULE_index)

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
		p.SetState(258)
		p.Match(TerraformParserT__14)
	}
	{
		p.SetState(259)
		p.expression(0)
	}
	{
		p.SetState(260)
		p.Match(TerraformParserT__15)
	}

	return localctx
}

// IFiledeclContext is an interface to support dynamic dispatch.
type IFiledeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFiledeclContext differentiates from other interfaces.
	IsFiledeclContext()
}

type FiledeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFiledeclContext() *FiledeclContext {
	var p = new(FiledeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_filedecl
	return p
}

func (*FiledeclContext) IsFiledeclContext() {}

func NewFiledeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FiledeclContext {
	var p = new(FiledeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_filedecl

	return p
}

func (s *FiledeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FiledeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TerraformParserLPAREN, 0)
}

func (s *FiledeclContext) Expression() IExpressionContext {
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

func (s *FiledeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TerraformParserRPAREN, 0)
}

func (s *FiledeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FiledeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FiledeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterFiledecl(s)
	}
}

func (s *FiledeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitFiledecl(s)
	}
}

func (p *TerraformParser) Filedecl() (localctx IFiledeclContext) {
	this := p
	_ = this

	localctx = NewFiledeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, TerraformParserRULE_filedecl)

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
		p.SetState(262)
		p.Match(TerraformParserT__16)
	}
	{
		p.SetState(263)
		p.Match(TerraformParserLPAREN)
	}
	{
		p.SetState(264)
		p.expression(0)
	}
	{
		p.SetState(265)
		p.Match(TerraformParserRPAREN)
	}

	return localctx
}

// IList_Context is an interface to support dynamic dispatch.
type IList_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsList_Context differentiates from other interfaces.
	IsList_Context()
}

type List_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_Context() *List_Context {
	var p = new(List_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_list_
	return p
}

func (*List_Context) IsList_Context() {}

func NewList_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_Context {
	var p = new(List_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_list_

	return p
}

func (s *List_Context) GetParser() antlr.Parser { return s.parser }

func (s *List_Context) AllExpression() []IExpressionContext {
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

func (s *List_Context) Expression(i int) IExpressionContext {
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

func (s *List_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterList_(s)
	}
}

func (s *List_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitList_(s)
	}
}

func (p *TerraformParser) List_() (localctx IList_Context) {
	this := p
	_ = this

	localctx = NewList_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, TerraformParserRULE_list_)
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
		p.SetState(267)
		p.Match(TerraformParserT__14)
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140375638324040) != 0 {
		{
			p.SetState(268)
			p.expression(0)
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(269)
					p.Match(TerraformParserT__13)
				}
				{
					p.SetState(270)
					p.expression(0)
				}

			}
			p.SetState(275)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
		}
		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TerraformParserT__13 {
			{
				p.SetState(276)
				p.Match(TerraformParserT__13)
			}

		}

	}
	{
		p.SetState(281)
		p.Match(TerraformParserT__15)
	}

	return localctx
}

// IMap_Context is an interface to support dynamic dispatch.
type IMap_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMap_Context differentiates from other interfaces.
	IsMap_Context()
}

type Map_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMap_Context() *Map_Context {
	var p = new(Map_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_map_
	return p
}

func (*Map_Context) IsMap_Context() {}

func NewMap_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Map_Context {
	var p = new(Map_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_map_

	return p
}

func (s *Map_Context) GetParser() antlr.Parser { return s.parser }

func (s *Map_Context) LCURL() antlr.TerminalNode {
	return s.GetToken(TerraformParserLCURL, 0)
}

func (s *Map_Context) RCURL() antlr.TerminalNode {
	return s.GetToken(TerraformParserRCURL, 0)
}

func (s *Map_Context) AllArgument() []IArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentContext); ok {
			len++
		}
	}

	tst := make([]IArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentContext); ok {
			tst[i] = t.(IArgumentContext)
			i++
		}
	}

	return tst
}

func (s *Map_Context) Argument(i int) IArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
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

	return t.(IArgumentContext)
}

func (s *Map_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Map_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Map_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterMap_(s)
	}
}

func (s *Map_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitMap_(s)
	}
}

func (p *TerraformParser) Map_() (localctx IMap_Context) {
	this := p
	_ = this

	localctx = NewMap_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, TerraformParserRULE_map_)
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
		p.SetState(283)
		p.Match(TerraformParserLCURL)
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72583873561416) != 0 {
		{
			p.SetState(284)
			p.Argument()
		}
		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TerraformParserT__13 {
			{
				p.SetState(285)
				p.Match(TerraformParserT__13)
			}

		}

		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(293)
		p.Match(TerraformParserRCURL)
	}

	return localctx
}

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_string
	return p
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(TerraformParserSTRING, 0)
}

func (s *StringContext) MULTILINESTRING() antlr.TerminalNode {
	return s.GetToken(TerraformParserMULTILINESTRING, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitString(s)
	}
}

func (p *TerraformParser) String_() (localctx IStringContext) {
	this := p
	_ = this

	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, TerraformParserRULE_string)
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
		p.SetState(295)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TerraformParserMULTILINESTRING || _la == TerraformParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISigned_numberContext is an interface to support dynamic dispatch.
type ISigned_numberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSigned_numberContext differentiates from other interfaces.
	IsSigned_numberContext()
}

type Signed_numberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySigned_numberContext() *Signed_numberContext {
	var p = new(Signed_numberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_signed_number
	return p
}

func (*Signed_numberContext) IsSigned_numberContext() {}

func NewSigned_numberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Signed_numberContext {
	var p = new(Signed_numberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_signed_number

	return p
}

func (s *Signed_numberContext) GetParser() antlr.Parser { return s.parser }

func (s *Signed_numberContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *Signed_numberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Signed_numberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Signed_numberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterSigned_number(s)
	}
}

func (s *Signed_numberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitSigned_number(s)
	}
}

func (p *TerraformParser) Signed_number() (localctx ISigned_numberContext) {
	this := p
	_ = this

	localctx = NewSigned_numberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, TerraformParserRULE_signed_number)
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
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TerraformParserT__17 || _la == TerraformParserT__18 {
		{
			p.SetState(297)
			_la = p.GetTokenStream().LA(1)

			if !(_la == TerraformParserT__17 || _la == TerraformParserT__18) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(300)
		p.Number()
	}

	return localctx
}

// IOperator_Context is an interface to support dynamic dispatch.
type IOperator_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperator_Context differentiates from other interfaces.
	IsOperator_Context()
}

type Operator_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperator_Context() *Operator_Context {
	var p = new(Operator_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_operator_
	return p
}

func (*Operator_Context) IsOperator_Context() {}

func NewOperator_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Operator_Context {
	var p = new(Operator_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_operator_

	return p
}

func (s *Operator_Context) GetParser() antlr.Parser { return s.parser }

func (s *Operator_Context) STAR() antlr.TerminalNode {
	return s.GetToken(TerraformParserSTAR, 0)
}

func (s *Operator_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Operator_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Operator_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterOperator_(s)
	}
}

func (s *Operator_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitOperator_(s)
	}
}

func (p *TerraformParser) Operator_() (localctx IOperator_Context) {
	this := p
	_ = this

	localctx = NewOperator_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, TerraformParserRULE_operator_)
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
		p.SetState(302)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9663414272) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraformParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraformParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) AllNATURAL_NUMBER() []antlr.TerminalNode {
	return s.GetTokens(TerraformParserNATURAL_NUMBER)
}

func (s *NumberContext) NATURAL_NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(TerraformParserNATURAL_NUMBER, i)
}

func (s *NumberContext) DOT() antlr.TerminalNode {
	return s.GetToken(TerraformParserDOT, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraformListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *TerraformParser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, TerraformParserRULE_number)

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
		p.SetState(304)
		p.Match(TerraformParserNATURAL_NUMBER)
	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(305)
			p.Match(TerraformParserDOT)
		}
		{
			p.SetState(306)
			p.Match(TerraformParserNATURAL_NUMBER)
		}

	}

	return localctx
}

func (p *TerraformParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 19:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TerraformParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
