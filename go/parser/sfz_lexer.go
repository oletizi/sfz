// Code generated from Sfz.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 20, 156,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 6, 13,
	104, 10, 13, 13, 13, 14, 13, 105, 3, 14, 3, 14, 5, 14, 110, 10, 14, 3,
	14, 3, 14, 5, 14, 114, 10, 14, 3, 15, 3, 15, 3, 15, 5, 15, 119, 10, 15,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 127, 10, 17, 12, 17, 14,
	17, 130, 11, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18,
	3, 18, 7, 18, 141, 10, 18, 12, 18, 14, 18, 144, 11, 18, 3, 18, 3, 18, 3,
	19, 3, 19, 7, 19, 150, 10, 19, 12, 19, 14, 19, 153, 11, 19, 3, 19, 3, 19,
	3, 128, 2, 20, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 3, 2, 4, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11, 34, 34, 2, 162, 2, 3,
	3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11,
	3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2,
	19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2,
	2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2,
	2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 3, 39, 3, 2, 2, 2, 5, 41, 3, 2,
	2, 2, 7, 43, 3, 2, 2, 2, 9, 45, 3, 2, 2, 2, 11, 52, 3, 2, 2, 2, 13, 58,
	3, 2, 2, 2, 15, 65, 3, 2, 2, 2, 17, 79, 3, 2, 2, 2, 19, 85, 3, 2, 2, 2,
	21, 89, 3, 2, 2, 2, 23, 95, 3, 2, 2, 2, 25, 103, 3, 2, 2, 2, 27, 113, 3,
	2, 2, 2, 29, 118, 3, 2, 2, 2, 31, 120, 3, 2, 2, 2, 33, 122, 3, 2, 2, 2,
	35, 136, 3, 2, 2, 2, 37, 147, 3, 2, 2, 2, 39, 40, 7, 63, 2, 2, 40, 4, 3,
	2, 2, 2, 41, 42, 7, 62, 2, 2, 42, 6, 3, 2, 2, 2, 43, 44, 7, 64, 2, 2, 44,
	8, 3, 2, 2, 2, 45, 46, 7, 105, 2, 2, 46, 47, 7, 110, 2, 2, 47, 48, 7, 113,
	2, 2, 48, 49, 7, 100, 2, 2, 49, 50, 7, 99, 2, 2, 50, 51, 7, 110, 2, 2,
	51, 10, 3, 2, 2, 2, 52, 53, 7, 105, 2, 2, 53, 54, 7, 116, 2, 2, 54, 55,
	7, 113, 2, 2, 55, 56, 7, 119, 2, 2, 56, 57, 7, 114, 2, 2, 57, 12, 3, 2,
	2, 2, 58, 59, 7, 116, 2, 2, 59, 60, 7, 103, 2, 2, 60, 61, 7, 105, 2, 2,
	61, 62, 7, 107, 2, 2, 62, 63, 7, 113, 2, 2, 63, 64, 7, 112, 2, 2, 64, 14,
	3, 2, 2, 2, 65, 66, 7, 99, 2, 2, 66, 67, 7, 111, 2, 2, 67, 68, 7, 114,
	2, 2, 68, 69, 7, 103, 2, 2, 69, 70, 7, 105, 2, 2, 70, 71, 7, 97, 2, 2,
	71, 72, 7, 116, 2, 2, 72, 73, 7, 103, 2, 2, 73, 74, 7, 110, 2, 2, 74, 75,
	7, 103, 2, 2, 75, 76, 7, 99, 2, 2, 76, 77, 7, 117, 2, 2, 77, 78, 7, 103,
	2, 2, 78, 16, 3, 2, 2, 2, 79, 80, 7, 106, 2, 2, 80, 81, 7, 107, 2, 2, 81,
	82, 7, 109, 2, 2, 82, 83, 7, 103, 2, 2, 83, 84, 7, 123, 2, 2, 84, 18, 3,
	2, 2, 2, 85, 86, 7, 109, 2, 2, 86, 87, 7, 103, 2, 2, 87, 88, 7, 123, 2,
	2, 88, 20, 3, 2, 2, 2, 89, 90, 7, 110, 2, 2, 90, 91, 7, 113, 2, 2, 91,
	92, 7, 109, 2, 2, 92, 93, 7, 103, 2, 2, 93, 94, 7, 123, 2, 2, 94, 22, 3,
	2, 2, 2, 95, 96, 7, 117, 2, 2, 96, 97, 7, 99, 2, 2, 97, 98, 7, 111, 2,
	2, 98, 99, 7, 114, 2, 2, 99, 100, 7, 110, 2, 2, 100, 101, 7, 103, 2, 2,
	101, 24, 3, 2, 2, 2, 102, 104, 4, 50, 59, 2, 103, 102, 3, 2, 2, 2, 104,
	105, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 26, 3,
	2, 2, 2, 107, 114, 5, 25, 13, 2, 108, 110, 5, 25, 13, 2, 109, 108, 3, 2,
	2, 2, 109, 110, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 112, 7, 48, 2, 2,
	112, 114, 5, 25, 13, 2, 113, 107, 3, 2, 2, 2, 113, 109, 3, 2, 2, 2, 114,
	28, 3, 2, 2, 2, 115, 116, 7, 15, 2, 2, 116, 119, 7, 12, 2, 2, 117, 119,
	9, 2, 2, 2, 118, 115, 3, 2, 2, 2, 118, 117, 3, 2, 2, 2, 119, 30, 3, 2,
	2, 2, 120, 121, 9, 3, 2, 2, 121, 32, 3, 2, 2, 2, 122, 123, 7, 49, 2, 2,
	123, 124, 7, 44, 2, 2, 124, 128, 3, 2, 2, 2, 125, 127, 11, 2, 2, 2, 126,
	125, 3, 2, 2, 2, 127, 130, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 128, 126,
	3, 2, 2, 2, 129, 131, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 131, 132, 7, 44,
	2, 2, 132, 133, 7, 49, 2, 2, 133, 134, 3, 2, 2, 2, 134, 135, 8, 17, 2,
	2, 135, 34, 3, 2, 2, 2, 136, 137, 7, 49, 2, 2, 137, 138, 7, 49, 2, 2, 138,
	142, 3, 2, 2, 2, 139, 141, 10, 2, 2, 2, 140, 139, 3, 2, 2, 2, 141, 144,
	3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 145, 3, 2,
	2, 2, 144, 142, 3, 2, 2, 2, 145, 146, 8, 18, 2, 2, 146, 36, 3, 2, 2, 2,
	147, 151, 7, 37, 2, 2, 148, 150, 10, 2, 2, 2, 149, 148, 3, 2, 2, 2, 150,
	153, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 154,
	3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 154, 155, 8, 19, 2, 2, 155, 38, 3, 2,
	2, 2, 10, 2, 105, 109, 113, 118, 128, 142, 151, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'<'", "'>'", "'global'", "'group'", "'region'", "'ampeg_release'",
	"'hikey'", "'key'", "'lokey'", "'sample'",
}

var lexerSymbolicNames = []string{
	"", "", "LT", "GT", "GLOBAL", "GROUP", "REGION", "AMPEG_RELEASE", "HIKEY",
	"KEY", "LOKEY", "SAMPLE", "INT", "FLOAT", "NEWLINE", "WHITESPACE", "BLOCK_COMMENT",
	"LINE_COMMENT", "HASH_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "LT", "GT", "GLOBAL", "GROUP", "REGION", "AMPEG_RELEASE", "HIKEY",
	"KEY", "LOKEY", "SAMPLE", "INT", "FLOAT", "NEWLINE", "WHITESPACE", "BLOCK_COMMENT",
	"LINE_COMMENT", "HASH_COMMENT",
}

type SfzLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewSfzLexer(input antlr.CharStream) *SfzLexer {

	l := new(SfzLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Sfz.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SfzLexer tokens.
const (
	SfzLexerT__0          = 1
	SfzLexerLT            = 2
	SfzLexerGT            = 3
	SfzLexerGLOBAL        = 4
	SfzLexerGROUP         = 5
	SfzLexerREGION        = 6
	SfzLexerAMPEG_RELEASE = 7
	SfzLexerHIKEY         = 8
	SfzLexerKEY           = 9
	SfzLexerLOKEY         = 10
	SfzLexerSAMPLE        = 11
	SfzLexerINT           = 12
	SfzLexerFLOAT         = 13
	SfzLexerNEWLINE       = 14
	SfzLexerWHITESPACE    = 15
	SfzLexerBLOCK_COMMENT = 16
	SfzLexerLINE_COMMENT  = 17
	SfzLexerHASH_COMMENT  = 18
)
