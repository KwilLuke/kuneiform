// Code generated from KuneiformLexer.g4 by ANTLR 4.12.0. DO NOT EDIT.

package kfgrammar

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

type KuneiformLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var kuneiformlexerLexerStaticData struct {
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

func kuneiformlexerLexerInit() {
	staticData := &kuneiformlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "':'", "';'", "'('", "'{'", "')'", "'}'", "','", "'$'", "'#'", "'@'",
		"'.'", "'='", "'database'", "'use'", "'as'", "'table'", "'action'",
		"'init'", "'public'", "'private'", "'view'", "'mustsign'", "'owner'",
		"'int'", "'text'", "'blob'", "'min'", "'max'", "'minlen'", "'maxlen'",
		"'notnull'", "'primary'", "'default'", "'unique'", "'index'", "'foreign_key'",
		"'fk'", "'references'", "'ref'", "'on_update'", "'on_delete'", "'do'",
		"'no_action'", "'cascade'", "'set_null'", "'set_default'", "'restrict'",
	}
	staticData.symbolicNames = []string{
		"", "COL", "SCOL", "L_PAREN", "L_BRACE", "R_PAREN", "R_BRACE", "COMMA",
		"DOLLAR", "HASH", "AT", "PERIOD", "EQ", "DATABASE_", "USE_", "AS_",
		"TABLE_", "ACTION_", "INIT_", "PUBLIC_", "PRIVATE_", "VIEW_", "MUSTSIGN_",
		"OWNER_", "INT_", "TEXT_", "BLOB_", "MIN_", "MAX_", "MIN_LEN_", "MAX_LEN_",
		"NOT_NULL_", "PRIMARY_", "DEFAULT_", "UNIQUE_", "INDEX_", "FOREIGN_KEY_",
		"FOREIGN_KEY_ABBR_", "REFERENCES_", "REFERENCES_ABBR_", "ACTION_ON_UPDATE_",
		"ACTION_ON_DELETE_", "ACTION_DO_", "ACTION_DO_NO_ACTION_", "ACTION_DO_CASCADE_",
		"ACTION_DO_SET_NULL_", "ACTION_DO_SET_DEFAULT_", "ACTION_DO_RESTRICT_",
		"SELECT_", "INSERT_", "UPDATE_", "DELETE_", "WITH_", "IDENTIFIER", "INDEX_NAME",
		"PARAM_OR_VAR", "BLOCK_VAR", "UNSIGNED_NUMBER_LITERAL", "SIGNED_NUMBER_LITERAL",
		"STRING_LITERAL", "SQL_KEYWORDS", "SQL_STMT", "WS", "TERMINATOR", "BLOCK_COMMENT",
		"LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"COL", "SCOL", "L_PAREN", "L_BRACE", "R_PAREN", "R_BRACE", "COMMA",
		"DOLLAR", "HASH", "AT", "PERIOD", "EQ", "DATABASE_", "USE_", "AS_",
		"TABLE_", "ACTION_", "INIT_", "PUBLIC_", "PRIVATE_", "VIEW_", "MUSTSIGN_",
		"OWNER_", "INT_", "TEXT_", "BLOB_", "MIN_", "MAX_", "MIN_LEN_", "MAX_LEN_",
		"NOT_NULL_", "PRIMARY_", "DEFAULT_", "UNIQUE_", "INDEX_", "FOREIGN_KEY_",
		"FOREIGN_KEY_ABBR_", "REFERENCES_", "REFERENCES_ABBR_", "ACTION_ON_UPDATE_",
		"ACTION_ON_DELETE_", "ACTION_DO_", "ACTION_DO_NO_ACTION_", "ACTION_DO_CASCADE_",
		"ACTION_DO_SET_NULL_", "ACTION_DO_SET_DEFAULT_", "ACTION_DO_RESTRICT_",
		"SELECT_", "INSERT_", "UPDATE_", "DELETE_", "WITH_", "IDENTIFIER", "INDEX_NAME",
		"PARAM_OR_VAR", "BLOCK_VAR", "UNSIGNED_NUMBER_LITERAL", "SIGNED_NUMBER_LITERAL",
		"STRING_LITERAL", "SQL_KEYWORDS", "SQL_STMT", "WS", "TERMINATOR", "BLOCK_COMMENT",
		"LINE_COMMENT", "WSNL", "DIGIT", "DOUBLE_QUOTE_STRING_CHAR", "SINGLE_QUOTE_STRING_CHAR",
		"DOUBLE_QUOTE_STRING", "SINGLE_QUOTE_STRING",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 65, 559, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41,
		1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1,
		42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44,
		1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46,
		1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48,
		1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1,
		50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52,
		1, 52, 5, 52, 446, 8, 52, 10, 52, 12, 52, 449, 9, 52, 1, 53, 1, 53, 1,
		53, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 56, 4, 56, 461, 8, 56,
		11, 56, 12, 56, 462, 1, 57, 3, 57, 466, 8, 57, 1, 57, 4, 57, 469, 8, 57,
		11, 57, 12, 57, 470, 1, 58, 1, 58, 3, 58, 475, 8, 58, 1, 59, 1, 59, 1,
		59, 1, 59, 1, 59, 3, 59, 482, 8, 59, 1, 60, 1, 60, 4, 60, 486, 8, 60, 11,
		60, 12, 60, 487, 1, 60, 4, 60, 491, 8, 60, 11, 60, 12, 60, 492, 1, 61,
		1, 61, 1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 62, 1, 63, 1, 63, 1, 63, 1,
		63, 5, 63, 507, 8, 63, 10, 63, 12, 63, 510, 9, 63, 1, 63, 1, 63, 1, 63,
		1, 63, 1, 63, 1, 64, 1, 64, 1, 64, 1, 64, 5, 64, 521, 8, 64, 10, 64, 12,
		64, 524, 9, 64, 1, 64, 1, 64, 1, 65, 1, 65, 1, 66, 1, 66, 1, 67, 1, 67,
		1, 67, 3, 67, 535, 8, 67, 1, 68, 1, 68, 1, 68, 3, 68, 540, 8, 68, 1, 69,
		1, 69, 5, 69, 544, 8, 69, 10, 69, 12, 69, 547, 9, 69, 1, 69, 1, 69, 1,
		70, 1, 70, 5, 70, 553, 8, 70, 10, 70, 12, 70, 556, 9, 70, 1, 70, 1, 70,
		1, 508, 0, 71, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18,
		37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27,
		55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36,
		73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45,
		91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107,
		54, 109, 55, 111, 56, 113, 57, 115, 58, 117, 59, 119, 60, 121, 61, 123,
		62, 125, 63, 127, 64, 129, 65, 131, 0, 133, 0, 135, 0, 137, 0, 139, 0,
		141, 0, 1, 0, 24, 2, 0, 83, 83, 115, 115, 2, 0, 69, 69, 101, 101, 2, 0,
		76, 76, 108, 108, 2, 0, 67, 67, 99, 99, 2, 0, 84, 84, 116, 116, 2, 0, 73,
		73, 105, 105, 2, 0, 78, 78, 110, 110, 2, 0, 82, 82, 114, 114, 2, 0, 85,
		85, 117, 117, 2, 0, 80, 80, 112, 112, 2, 0, 68, 68, 100, 100, 2, 0, 65,
		65, 97, 97, 2, 0, 87, 87, 119, 119, 2, 0, 72, 72, 104, 104, 2, 0, 65, 90,
		97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 43, 43, 45, 45, 2,
		0, 59, 59, 125, 125, 2, 0, 9, 9, 32, 32, 2, 0, 10, 10, 13, 13, 3, 0, 9,
		10, 13, 13, 32, 32, 1, 0, 48, 57, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92,
		4, 0, 10, 10, 13, 13, 39, 39, 92, 92, 569, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
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
		111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0,
		0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125,
		1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0, 1, 143, 1, 0, 0, 0,
		3, 145, 1, 0, 0, 0, 5, 147, 1, 0, 0, 0, 7, 149, 1, 0, 0, 0, 9, 151, 1,
		0, 0, 0, 11, 153, 1, 0, 0, 0, 13, 155, 1, 0, 0, 0, 15, 157, 1, 0, 0, 0,
		17, 159, 1, 0, 0, 0, 19, 161, 1, 0, 0, 0, 21, 163, 1, 0, 0, 0, 23, 165,
		1, 0, 0, 0, 25, 167, 1, 0, 0, 0, 27, 176, 1, 0, 0, 0, 29, 180, 1, 0, 0,
		0, 31, 183, 1, 0, 0, 0, 33, 189, 1, 0, 0, 0, 35, 196, 1, 0, 0, 0, 37, 201,
		1, 0, 0, 0, 39, 208, 1, 0, 0, 0, 41, 216, 1, 0, 0, 0, 43, 221, 1, 0, 0,
		0, 45, 230, 1, 0, 0, 0, 47, 236, 1, 0, 0, 0, 49, 240, 1, 0, 0, 0, 51, 245,
		1, 0, 0, 0, 53, 250, 1, 0, 0, 0, 55, 254, 1, 0, 0, 0, 57, 258, 1, 0, 0,
		0, 59, 265, 1, 0, 0, 0, 61, 272, 1, 0, 0, 0, 63, 280, 1, 0, 0, 0, 65, 288,
		1, 0, 0, 0, 67, 296, 1, 0, 0, 0, 69, 303, 1, 0, 0, 0, 71, 309, 1, 0, 0,
		0, 73, 321, 1, 0, 0, 0, 75, 324, 1, 0, 0, 0, 77, 335, 1, 0, 0, 0, 79, 339,
		1, 0, 0, 0, 81, 349, 1, 0, 0, 0, 83, 359, 1, 0, 0, 0, 85, 362, 1, 0, 0,
		0, 87, 372, 1, 0, 0, 0, 89, 380, 1, 0, 0, 0, 91, 389, 1, 0, 0, 0, 93, 401,
		1, 0, 0, 0, 95, 410, 1, 0, 0, 0, 97, 417, 1, 0, 0, 0, 99, 424, 1, 0, 0,
		0, 101, 431, 1, 0, 0, 0, 103, 438, 1, 0, 0, 0, 105, 443, 1, 0, 0, 0, 107,
		450, 1, 0, 0, 0, 109, 453, 1, 0, 0, 0, 111, 456, 1, 0, 0, 0, 113, 460,
		1, 0, 0, 0, 115, 465, 1, 0, 0, 0, 117, 474, 1, 0, 0, 0, 119, 481, 1, 0,
		0, 0, 121, 483, 1, 0, 0, 0, 123, 494, 1, 0, 0, 0, 125, 498, 1, 0, 0, 0,
		127, 502, 1, 0, 0, 0, 129, 516, 1, 0, 0, 0, 131, 527, 1, 0, 0, 0, 133,
		529, 1, 0, 0, 0, 135, 534, 1, 0, 0, 0, 137, 539, 1, 0, 0, 0, 139, 541,
		1, 0, 0, 0, 141, 550, 1, 0, 0, 0, 143, 144, 5, 58, 0, 0, 144, 2, 1, 0,
		0, 0, 145, 146, 5, 59, 0, 0, 146, 4, 1, 0, 0, 0, 147, 148, 5, 40, 0, 0,
		148, 6, 1, 0, 0, 0, 149, 150, 5, 123, 0, 0, 150, 8, 1, 0, 0, 0, 151, 152,
		5, 41, 0, 0, 152, 10, 1, 0, 0, 0, 153, 154, 5, 125, 0, 0, 154, 12, 1, 0,
		0, 0, 155, 156, 5, 44, 0, 0, 156, 14, 1, 0, 0, 0, 157, 158, 5, 36, 0, 0,
		158, 16, 1, 0, 0, 0, 159, 160, 5, 35, 0, 0, 160, 18, 1, 0, 0, 0, 161, 162,
		5, 64, 0, 0, 162, 20, 1, 0, 0, 0, 163, 164, 5, 46, 0, 0, 164, 22, 1, 0,
		0, 0, 165, 166, 5, 61, 0, 0, 166, 24, 1, 0, 0, 0, 167, 168, 5, 100, 0,
		0, 168, 169, 5, 97, 0, 0, 169, 170, 5, 116, 0, 0, 170, 171, 5, 97, 0, 0,
		171, 172, 5, 98, 0, 0, 172, 173, 5, 97, 0, 0, 173, 174, 5, 115, 0, 0, 174,
		175, 5, 101, 0, 0, 175, 26, 1, 0, 0, 0, 176, 177, 5, 117, 0, 0, 177, 178,
		5, 115, 0, 0, 178, 179, 5, 101, 0, 0, 179, 28, 1, 0, 0, 0, 180, 181, 5,
		97, 0, 0, 181, 182, 5, 115, 0, 0, 182, 30, 1, 0, 0, 0, 183, 184, 5, 116,
		0, 0, 184, 185, 5, 97, 0, 0, 185, 186, 5, 98, 0, 0, 186, 187, 5, 108, 0,
		0, 187, 188, 5, 101, 0, 0, 188, 32, 1, 0, 0, 0, 189, 190, 5, 97, 0, 0,
		190, 191, 5, 99, 0, 0, 191, 192, 5, 116, 0, 0, 192, 193, 5, 105, 0, 0,
		193, 194, 5, 111, 0, 0, 194, 195, 5, 110, 0, 0, 195, 34, 1, 0, 0, 0, 196,
		197, 5, 105, 0, 0, 197, 198, 5, 110, 0, 0, 198, 199, 5, 105, 0, 0, 199,
		200, 5, 116, 0, 0, 200, 36, 1, 0, 0, 0, 201, 202, 5, 112, 0, 0, 202, 203,
		5, 117, 0, 0, 203, 204, 5, 98, 0, 0, 204, 205, 5, 108, 0, 0, 205, 206,
		5, 105, 0, 0, 206, 207, 5, 99, 0, 0, 207, 38, 1, 0, 0, 0, 208, 209, 5,
		112, 0, 0, 209, 210, 5, 114, 0, 0, 210, 211, 5, 105, 0, 0, 211, 212, 5,
		118, 0, 0, 212, 213, 5, 97, 0, 0, 213, 214, 5, 116, 0, 0, 214, 215, 5,
		101, 0, 0, 215, 40, 1, 0, 0, 0, 216, 217, 5, 118, 0, 0, 217, 218, 5, 105,
		0, 0, 218, 219, 5, 101, 0, 0, 219, 220, 5, 119, 0, 0, 220, 42, 1, 0, 0,
		0, 221, 222, 5, 109, 0, 0, 222, 223, 5, 117, 0, 0, 223, 224, 5, 115, 0,
		0, 224, 225, 5, 116, 0, 0, 225, 226, 5, 115, 0, 0, 226, 227, 5, 105, 0,
		0, 227, 228, 5, 103, 0, 0, 228, 229, 5, 110, 0, 0, 229, 44, 1, 0, 0, 0,
		230, 231, 5, 111, 0, 0, 231, 232, 5, 119, 0, 0, 232, 233, 5, 110, 0, 0,
		233, 234, 5, 101, 0, 0, 234, 235, 5, 114, 0, 0, 235, 46, 1, 0, 0, 0, 236,
		237, 5, 105, 0, 0, 237, 238, 5, 110, 0, 0, 238, 239, 5, 116, 0, 0, 239,
		48, 1, 0, 0, 0, 240, 241, 5, 116, 0, 0, 241, 242, 5, 101, 0, 0, 242, 243,
		5, 120, 0, 0, 243, 244, 5, 116, 0, 0, 244, 50, 1, 0, 0, 0, 245, 246, 5,
		98, 0, 0, 246, 247, 5, 108, 0, 0, 247, 248, 5, 111, 0, 0, 248, 249, 5,
		98, 0, 0, 249, 52, 1, 0, 0, 0, 250, 251, 5, 109, 0, 0, 251, 252, 5, 105,
		0, 0, 252, 253, 5, 110, 0, 0, 253, 54, 1, 0, 0, 0, 254, 255, 5, 109, 0,
		0, 255, 256, 5, 97, 0, 0, 256, 257, 5, 120, 0, 0, 257, 56, 1, 0, 0, 0,
		258, 259, 5, 109, 0, 0, 259, 260, 5, 105, 0, 0, 260, 261, 5, 110, 0, 0,
		261, 262, 5, 108, 0, 0, 262, 263, 5, 101, 0, 0, 263, 264, 5, 110, 0, 0,
		264, 58, 1, 0, 0, 0, 265, 266, 5, 109, 0, 0, 266, 267, 5, 97, 0, 0, 267,
		268, 5, 120, 0, 0, 268, 269, 5, 108, 0, 0, 269, 270, 5, 101, 0, 0, 270,
		271, 5, 110, 0, 0, 271, 60, 1, 0, 0, 0, 272, 273, 5, 110, 0, 0, 273, 274,
		5, 111, 0, 0, 274, 275, 5, 116, 0, 0, 275, 276, 5, 110, 0, 0, 276, 277,
		5, 117, 0, 0, 277, 278, 5, 108, 0, 0, 278, 279, 5, 108, 0, 0, 279, 62,
		1, 0, 0, 0, 280, 281, 5, 112, 0, 0, 281, 282, 5, 114, 0, 0, 282, 283, 5,
		105, 0, 0, 283, 284, 5, 109, 0, 0, 284, 285, 5, 97, 0, 0, 285, 286, 5,
		114, 0, 0, 286, 287, 5, 121, 0, 0, 287, 64, 1, 0, 0, 0, 288, 289, 5, 100,
		0, 0, 289, 290, 5, 101, 0, 0, 290, 291, 5, 102, 0, 0, 291, 292, 5, 97,
		0, 0, 292, 293, 5, 117, 0, 0, 293, 294, 5, 108, 0, 0, 294, 295, 5, 116,
		0, 0, 295, 66, 1, 0, 0, 0, 296, 297, 5, 117, 0, 0, 297, 298, 5, 110, 0,
		0, 298, 299, 5, 105, 0, 0, 299, 300, 5, 113, 0, 0, 300, 301, 5, 117, 0,
		0, 301, 302, 5, 101, 0, 0, 302, 68, 1, 0, 0, 0, 303, 304, 5, 105, 0, 0,
		304, 305, 5, 110, 0, 0, 305, 306, 5, 100, 0, 0, 306, 307, 5, 101, 0, 0,
		307, 308, 5, 120, 0, 0, 308, 70, 1, 0, 0, 0, 309, 310, 5, 102, 0, 0, 310,
		311, 5, 111, 0, 0, 311, 312, 5, 114, 0, 0, 312, 313, 5, 101, 0, 0, 313,
		314, 5, 105, 0, 0, 314, 315, 5, 103, 0, 0, 315, 316, 5, 110, 0, 0, 316,
		317, 5, 95, 0, 0, 317, 318, 5, 107, 0, 0, 318, 319, 5, 101, 0, 0, 319,
		320, 5, 121, 0, 0, 320, 72, 1, 0, 0, 0, 321, 322, 5, 102, 0, 0, 322, 323,
		5, 107, 0, 0, 323, 74, 1, 0, 0, 0, 324, 325, 5, 114, 0, 0, 325, 326, 5,
		101, 0, 0, 326, 327, 5, 102, 0, 0, 327, 328, 5, 101, 0, 0, 328, 329, 5,
		114, 0, 0, 329, 330, 5, 101, 0, 0, 330, 331, 5, 110, 0, 0, 331, 332, 5,
		99, 0, 0, 332, 333, 5, 101, 0, 0, 333, 334, 5, 115, 0, 0, 334, 76, 1, 0,
		0, 0, 335, 336, 5, 114, 0, 0, 336, 337, 5, 101, 0, 0, 337, 338, 5, 102,
		0, 0, 338, 78, 1, 0, 0, 0, 339, 340, 5, 111, 0, 0, 340, 341, 5, 110, 0,
		0, 341, 342, 5, 95, 0, 0, 342, 343, 5, 117, 0, 0, 343, 344, 5, 112, 0,
		0, 344, 345, 5, 100, 0, 0, 345, 346, 5, 97, 0, 0, 346, 347, 5, 116, 0,
		0, 347, 348, 5, 101, 0, 0, 348, 80, 1, 0, 0, 0, 349, 350, 5, 111, 0, 0,
		350, 351, 5, 110, 0, 0, 351, 352, 5, 95, 0, 0, 352, 353, 5, 100, 0, 0,
		353, 354, 5, 101, 0, 0, 354, 355, 5, 108, 0, 0, 355, 356, 5, 101, 0, 0,
		356, 357, 5, 116, 0, 0, 357, 358, 5, 101, 0, 0, 358, 82, 1, 0, 0, 0, 359,
		360, 5, 100, 0, 0, 360, 361, 5, 111, 0, 0, 361, 84, 1, 0, 0, 0, 362, 363,
		5, 110, 0, 0, 363, 364, 5, 111, 0, 0, 364, 365, 5, 95, 0, 0, 365, 366,
		5, 97, 0, 0, 366, 367, 5, 99, 0, 0, 367, 368, 5, 116, 0, 0, 368, 369, 5,
		105, 0, 0, 369, 370, 5, 111, 0, 0, 370, 371, 5, 110, 0, 0, 371, 86, 1,
		0, 0, 0, 372, 373, 5, 99, 0, 0, 373, 374, 5, 97, 0, 0, 374, 375, 5, 115,
		0, 0, 375, 376, 5, 99, 0, 0, 376, 377, 5, 97, 0, 0, 377, 378, 5, 100, 0,
		0, 378, 379, 5, 101, 0, 0, 379, 88, 1, 0, 0, 0, 380, 381, 5, 115, 0, 0,
		381, 382, 5, 101, 0, 0, 382, 383, 5, 116, 0, 0, 383, 384, 5, 95, 0, 0,
		384, 385, 5, 110, 0, 0, 385, 386, 5, 117, 0, 0, 386, 387, 5, 108, 0, 0,
		387, 388, 5, 108, 0, 0, 388, 90, 1, 0, 0, 0, 389, 390, 5, 115, 0, 0, 390,
		391, 5, 101, 0, 0, 391, 392, 5, 116, 0, 0, 392, 393, 5, 95, 0, 0, 393,
		394, 5, 100, 0, 0, 394, 395, 5, 101, 0, 0, 395, 396, 5, 102, 0, 0, 396,
		397, 5, 97, 0, 0, 397, 398, 5, 117, 0, 0, 398, 399, 5, 108, 0, 0, 399,
		400, 5, 116, 0, 0, 400, 92, 1, 0, 0, 0, 401, 402, 5, 114, 0, 0, 402, 403,
		5, 101, 0, 0, 403, 404, 5, 115, 0, 0, 404, 405, 5, 116, 0, 0, 405, 406,
		5, 114, 0, 0, 406, 407, 5, 105, 0, 0, 407, 408, 5, 99, 0, 0, 408, 409,
		5, 116, 0, 0, 409, 94, 1, 0, 0, 0, 410, 411, 7, 0, 0, 0, 411, 412, 7, 1,
		0, 0, 412, 413, 7, 2, 0, 0, 413, 414, 7, 1, 0, 0, 414, 415, 7, 3, 0, 0,
		415, 416, 7, 4, 0, 0, 416, 96, 1, 0, 0, 0, 417, 418, 7, 5, 0, 0, 418, 419,
		7, 6, 0, 0, 419, 420, 7, 0, 0, 0, 420, 421, 7, 1, 0, 0, 421, 422, 7, 7,
		0, 0, 422, 423, 7, 4, 0, 0, 423, 98, 1, 0, 0, 0, 424, 425, 7, 8, 0, 0,
		425, 426, 7, 9, 0, 0, 426, 427, 7, 10, 0, 0, 427, 428, 7, 11, 0, 0, 428,
		429, 7, 4, 0, 0, 429, 430, 7, 1, 0, 0, 430, 100, 1, 0, 0, 0, 431, 432,
		7, 10, 0, 0, 432, 433, 7, 1, 0, 0, 433, 434, 7, 2, 0, 0, 434, 435, 7, 1,
		0, 0, 435, 436, 7, 4, 0, 0, 436, 437, 7, 1, 0, 0, 437, 102, 1, 0, 0, 0,
		438, 439, 7, 12, 0, 0, 439, 440, 7, 5, 0, 0, 440, 441, 7, 4, 0, 0, 441,
		442, 7, 13, 0, 0, 442, 104, 1, 0, 0, 0, 443, 447, 7, 14, 0, 0, 444, 446,
		7, 15, 0, 0, 445, 444, 1, 0, 0, 0, 446, 449, 1, 0, 0, 0, 447, 445, 1, 0,
		0, 0, 447, 448, 1, 0, 0, 0, 448, 106, 1, 0, 0, 0, 449, 447, 1, 0, 0, 0,
		450, 451, 3, 17, 8, 0, 451, 452, 3, 105, 52, 0, 452, 108, 1, 0, 0, 0, 453,
		454, 3, 15, 7, 0, 454, 455, 3, 105, 52, 0, 455, 110, 1, 0, 0, 0, 456, 457,
		3, 19, 9, 0, 457, 458, 3, 105, 52, 0, 458, 112, 1, 0, 0, 0, 459, 461, 3,
		133, 66, 0, 460, 459, 1, 0, 0, 0, 461, 462, 1, 0, 0, 0, 462, 460, 1, 0,
		0, 0, 462, 463, 1, 0, 0, 0, 463, 114, 1, 0, 0, 0, 464, 466, 7, 16, 0, 0,
		465, 464, 1, 0, 0, 0, 465, 466, 1, 0, 0, 0, 466, 468, 1, 0, 0, 0, 467,
		469, 3, 133, 66, 0, 468, 467, 1, 0, 0, 0, 469, 470, 1, 0, 0, 0, 470, 468,
		1, 0, 0, 0, 470, 471, 1, 0, 0, 0, 471, 116, 1, 0, 0, 0, 472, 475, 3, 139,
		69, 0, 473, 475, 3, 141, 70, 0, 474, 472, 1, 0, 0, 0, 474, 473, 1, 0, 0,
		0, 475, 118, 1, 0, 0, 0, 476, 482, 3, 95, 47, 0, 477, 482, 3, 97, 48, 0,
		478, 482, 3, 99, 49, 0, 479, 482, 3, 101, 50, 0, 480, 482, 3, 103, 51,
		0, 481, 476, 1, 0, 0, 0, 481, 477, 1, 0, 0, 0, 481, 478, 1, 0, 0, 0, 481,
		479, 1, 0, 0, 0, 481, 480, 1, 0, 0, 0, 482, 120, 1, 0, 0, 0, 483, 485,
		3, 119, 59, 0, 484, 486, 3, 131, 65, 0, 485, 484, 1, 0, 0, 0, 486, 487,
		1, 0, 0, 0, 487, 485, 1, 0, 0, 0, 487, 488, 1, 0, 0, 0, 488, 490, 1, 0,
		0, 0, 489, 491, 8, 17, 0, 0, 490, 489, 1, 0, 0, 0, 491, 492, 1, 0, 0, 0,
		492, 490, 1, 0, 0, 0, 492, 493, 1, 0, 0, 0, 493, 122, 1, 0, 0, 0, 494,
		495, 7, 18, 0, 0, 495, 496, 1, 0, 0, 0, 496, 497, 6, 61, 0, 0, 497, 124,
		1, 0, 0, 0, 498, 499, 7, 19, 0, 0, 499, 500, 1, 0, 0, 0, 500, 501, 6, 62,
		0, 0, 501, 126, 1, 0, 0, 0, 502, 503, 5, 47, 0, 0, 503, 504, 5, 42, 0,
		0, 504, 508, 1, 0, 0, 0, 505, 507, 9, 0, 0, 0, 506, 505, 1, 0, 0, 0, 507,
		510, 1, 0, 0, 0, 508, 509, 1, 0, 0, 0, 508, 506, 1, 0, 0, 0, 509, 511,
		1, 0, 0, 0, 510, 508, 1, 0, 0, 0, 511, 512, 5, 42, 0, 0, 512, 513, 5, 47,
		0, 0, 513, 514, 1, 0, 0, 0, 514, 515, 6, 63, 0, 0, 515, 128, 1, 0, 0, 0,
		516, 517, 5, 47, 0, 0, 517, 518, 5, 47, 0, 0, 518, 522, 1, 0, 0, 0, 519,
		521, 8, 19, 0, 0, 520, 519, 1, 0, 0, 0, 521, 524, 1, 0, 0, 0, 522, 520,
		1, 0, 0, 0, 522, 523, 1, 0, 0, 0, 523, 525, 1, 0, 0, 0, 524, 522, 1, 0,
		0, 0, 525, 526, 6, 64, 0, 0, 526, 130, 1, 0, 0, 0, 527, 528, 7, 20, 0,
		0, 528, 132, 1, 0, 0, 0, 529, 530, 7, 21, 0, 0, 530, 134, 1, 0, 0, 0, 531,
		535, 8, 22, 0, 0, 532, 533, 5, 92, 0, 0, 533, 535, 9, 0, 0, 0, 534, 531,
		1, 0, 0, 0, 534, 532, 1, 0, 0, 0, 535, 136, 1, 0, 0, 0, 536, 540, 8, 23,
		0, 0, 537, 538, 5, 92, 0, 0, 538, 540, 9, 0, 0, 0, 539, 536, 1, 0, 0, 0,
		539, 537, 1, 0, 0, 0, 540, 138, 1, 0, 0, 0, 541, 545, 5, 34, 0, 0, 542,
		544, 3, 135, 67, 0, 543, 542, 1, 0, 0, 0, 544, 547, 1, 0, 0, 0, 545, 543,
		1, 0, 0, 0, 545, 546, 1, 0, 0, 0, 546, 548, 1, 0, 0, 0, 547, 545, 1, 0,
		0, 0, 548, 549, 5, 34, 0, 0, 549, 140, 1, 0, 0, 0, 550, 554, 5, 39, 0,
		0, 551, 553, 3, 137, 68, 0, 552, 551, 1, 0, 0, 0, 553, 556, 1, 0, 0, 0,
		554, 552, 1, 0, 0, 0, 554, 555, 1, 0, 0, 0, 555, 557, 1, 0, 0, 0, 556,
		554, 1, 0, 0, 0, 557, 558, 5, 39, 0, 0, 558, 142, 1, 0, 0, 0, 15, 0, 447,
		462, 465, 470, 474, 481, 487, 492, 508, 522, 534, 539, 545, 554, 1, 0,
		1, 0,
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

// KuneiformLexerInit initializes any static state used to implement KuneiformLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewKuneiformLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func KuneiformLexerInit() {
	staticData := &kuneiformlexerLexerStaticData
	staticData.once.Do(kuneiformlexerLexerInit)
}

// NewKuneiformLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewKuneiformLexer(input antlr.CharStream) *KuneiformLexer {
	KuneiformLexerInit()
	l := new(KuneiformLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &kuneiformlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "KuneiformLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// KuneiformLexer tokens.
const (
	KuneiformLexerCOL                     = 1
	KuneiformLexerSCOL                    = 2
	KuneiformLexerL_PAREN                 = 3
	KuneiformLexerL_BRACE                 = 4
	KuneiformLexerR_PAREN                 = 5
	KuneiformLexerR_BRACE                 = 6
	KuneiformLexerCOMMA                   = 7
	KuneiformLexerDOLLAR                  = 8
	KuneiformLexerHASH                    = 9
	KuneiformLexerAT                      = 10
	KuneiformLexerPERIOD                  = 11
	KuneiformLexerEQ                      = 12
	KuneiformLexerDATABASE_               = 13
	KuneiformLexerUSE_                    = 14
	KuneiformLexerAS_                     = 15
	KuneiformLexerTABLE_                  = 16
	KuneiformLexerACTION_                 = 17
	KuneiformLexerINIT_                   = 18
	KuneiformLexerPUBLIC_                 = 19
	KuneiformLexerPRIVATE_                = 20
	KuneiformLexerVIEW_                   = 21
	KuneiformLexerMUSTSIGN_               = 22
	KuneiformLexerOWNER_                  = 23
	KuneiformLexerINT_                    = 24
	KuneiformLexerTEXT_                   = 25
	KuneiformLexerBLOB_                   = 26
	KuneiformLexerMIN_                    = 27
	KuneiformLexerMAX_                    = 28
	KuneiformLexerMIN_LEN_                = 29
	KuneiformLexerMAX_LEN_                = 30
	KuneiformLexerNOT_NULL_               = 31
	KuneiformLexerPRIMARY_                = 32
	KuneiformLexerDEFAULT_                = 33
	KuneiformLexerUNIQUE_                 = 34
	KuneiformLexerINDEX_                  = 35
	KuneiformLexerFOREIGN_KEY_            = 36
	KuneiformLexerFOREIGN_KEY_ABBR_       = 37
	KuneiformLexerREFERENCES_             = 38
	KuneiformLexerREFERENCES_ABBR_        = 39
	KuneiformLexerACTION_ON_UPDATE_       = 40
	KuneiformLexerACTION_ON_DELETE_       = 41
	KuneiformLexerACTION_DO_              = 42
	KuneiformLexerACTION_DO_NO_ACTION_    = 43
	KuneiformLexerACTION_DO_CASCADE_      = 44
	KuneiformLexerACTION_DO_SET_NULL_     = 45
	KuneiformLexerACTION_DO_SET_DEFAULT_  = 46
	KuneiformLexerACTION_DO_RESTRICT_     = 47
	KuneiformLexerSELECT_                 = 48
	KuneiformLexerINSERT_                 = 49
	KuneiformLexerUPDATE_                 = 50
	KuneiformLexerDELETE_                 = 51
	KuneiformLexerWITH_                   = 52
	KuneiformLexerIDENTIFIER              = 53
	KuneiformLexerINDEX_NAME              = 54
	KuneiformLexerPARAM_OR_VAR            = 55
	KuneiformLexerBLOCK_VAR               = 56
	KuneiformLexerUNSIGNED_NUMBER_LITERAL = 57
	KuneiformLexerSIGNED_NUMBER_LITERAL   = 58
	KuneiformLexerSTRING_LITERAL          = 59
	KuneiformLexerSQL_KEYWORDS            = 60
	KuneiformLexerSQL_STMT                = 61
	KuneiformLexerWS                      = 62
	KuneiformLexerTERMINATOR              = 63
	KuneiformLexerBLOCK_COMMENT           = 64
	KuneiformLexerLINE_COMMENT            = 65
)
