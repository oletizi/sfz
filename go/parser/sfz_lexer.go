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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 99, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 6, 8, 55, 10, 8, 13, 8, 14, 8, 56, 3, 9,
	3, 9, 3, 9, 5, 9, 62, 10, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11,
	7, 11, 70, 10, 11, 12, 11, 14, 11, 73, 11, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 84, 10, 12, 12, 12, 14, 12,
	87, 11, 12, 3, 12, 3, 12, 3, 13, 3, 13, 7, 13, 93, 10, 13, 12, 13, 14,
	13, 96, 11, 13, 3, 13, 3, 13, 3, 71, 2, 14, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 3, 2, 5, 6, 2,
	11, 12, 15, 15, 34, 34, 62, 64, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11, 34,
	34, 2, 103, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2,
	2, 25, 3, 2, 2, 2, 3, 27, 3, 2, 2, 2, 5, 29, 3, 2, 2, 2, 7, 31, 3, 2, 2,
	2, 9, 38, 3, 2, 2, 2, 11, 44, 3, 2, 2, 2, 13, 51, 3, 2, 2, 2, 15, 54, 3,
	2, 2, 2, 17, 61, 3, 2, 2, 2, 19, 63, 3, 2, 2, 2, 21, 65, 3, 2, 2, 2, 23,
	79, 3, 2, 2, 2, 25, 90, 3, 2, 2, 2, 27, 28, 7, 62, 2, 2, 28, 4, 3, 2, 2,
	2, 29, 30, 7, 64, 2, 2, 30, 6, 3, 2, 2, 2, 31, 32, 7, 105, 2, 2, 32, 33,
	7, 110, 2, 2, 33, 34, 7, 113, 2, 2, 34, 35, 7, 100, 2, 2, 35, 36, 7, 99,
	2, 2, 36, 37, 7, 110, 2, 2, 37, 8, 3, 2, 2, 2, 38, 39, 7, 105, 2, 2, 39,
	40, 7, 116, 2, 2, 40, 41, 7, 113, 2, 2, 41, 42, 7, 119, 2, 2, 42, 43, 7,
	114, 2, 2, 43, 10, 3, 2, 2, 2, 44, 45, 7, 116, 2, 2, 45, 46, 7, 103, 2,
	2, 46, 47, 7, 105, 2, 2, 47, 48, 7, 107, 2, 2, 48, 49, 7, 113, 2, 2, 49,
	50, 7, 112, 2, 2, 50, 12, 3, 2, 2, 2, 51, 52, 7, 63, 2, 2, 52, 14, 3, 2,
	2, 2, 53, 55, 10, 2, 2, 2, 54, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56,
	54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 16, 3, 2, 2, 2, 58, 59, 7, 15,
	2, 2, 59, 62, 7, 12, 2, 2, 60, 62, 9, 3, 2, 2, 61, 58, 3, 2, 2, 2, 61,
	60, 3, 2, 2, 2, 62, 18, 3, 2, 2, 2, 63, 64, 9, 4, 2, 2, 64, 20, 3, 2, 2,
	2, 65, 66, 7, 49, 2, 2, 66, 67, 7, 44, 2, 2, 67, 71, 3, 2, 2, 2, 68, 70,
	11, 2, 2, 2, 69, 68, 3, 2, 2, 2, 70, 73, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2,
	71, 69, 3, 2, 2, 2, 72, 74, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 74, 75, 7,
	44, 2, 2, 75, 76, 7, 49, 2, 2, 76, 77, 3, 2, 2, 2, 77, 78, 8, 11, 2, 2,
	78, 22, 3, 2, 2, 2, 79, 80, 7, 49, 2, 2, 80, 81, 7, 49, 2, 2, 81, 85, 3,
	2, 2, 2, 82, 84, 10, 3, 2, 2, 83, 82, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85,
	83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 88, 3, 2, 2, 2, 87, 85, 3, 2, 2,
	2, 88, 89, 8, 12, 2, 2, 89, 24, 3, 2, 2, 2, 90, 94, 7, 37, 2, 2, 91, 93,
	10, 3, 2, 2, 92, 91, 3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2,
	94, 95, 3, 2, 2, 2, 95, 97, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 97, 98, 8,
	13, 2, 2, 98, 26, 3, 2, 2, 2, 8, 2, 56, 61, 71, 85, 94, 3, 8, 2, 2,
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
	"", "'<'", "'>'", "'global'", "'group'", "'region'", "'='",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "STRING", "NEWLINE", "WHITESPACE", "BLOCK_COMMENT",
	"LINE_COMMENT", "HASH_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "STRING", "NEWLINE", "WHITESPACE",
	"BLOCK_COMMENT", "LINE_COMMENT", "HASH_COMMENT",
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
	SfzLexerT__1          = 2
	SfzLexerT__2          = 3
	SfzLexerT__3          = 4
	SfzLexerT__4          = 5
	SfzLexerT__5          = 6
	SfzLexerSTRING        = 7
	SfzLexerNEWLINE       = 8
	SfzLexerWHITESPACE    = 9
	SfzLexerBLOCK_COMMENT = 10
	SfzLexerLINE_COMMENT  = 11
	SfzLexerHASH_COMMENT  = 12
)
