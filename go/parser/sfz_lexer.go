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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 28, 263,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19,
	3, 19, 6, 19, 180, 10, 19, 13, 19, 14, 19, 181, 3, 20, 6, 20, 185, 10,
	20, 13, 20, 14, 20, 186, 3, 21, 3, 21, 5, 21, 191, 10, 21, 3, 21, 3, 21,
	5, 21, 195, 10, 21, 3, 22, 7, 22, 198, 10, 22, 12, 22, 14, 22, 201, 11,
	22, 3, 22, 3, 22, 6, 22, 205, 10, 22, 13, 22, 14, 22, 206, 3, 22, 3, 22,
	7, 22, 211, 10, 22, 12, 22, 14, 22, 214, 11, 22, 3, 23, 3, 23, 3, 24, 6,
	24, 219, 10, 24, 13, 24, 14, 24, 220, 3, 25, 3, 25, 3, 25, 5, 25, 226,
	10, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 7, 27, 234, 10, 27, 12,
	27, 14, 27, 237, 11, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28,
	3, 28, 3, 28, 7, 28, 248, 10, 28, 12, 28, 14, 28, 251, 11, 28, 3, 28, 3,
	28, 3, 29, 3, 29, 7, 29, 257, 10, 29, 12, 29, 14, 29, 260, 11, 29, 3, 29,
	3, 29, 3, 235, 2, 30, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35,
	19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 2, 47, 2, 49, 24, 51, 25, 53, 26,
	55, 27, 57, 28, 3, 2, 7, 3, 2, 99, 124, 4, 2, 49, 49, 94, 94, 7, 2, 47,
	48, 50, 59, 67, 92, 97, 97, 99, 124, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11,
	34, 34, 2, 272, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2,
	2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3,
	2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51,
	3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 3,
	59, 3, 2, 2, 2, 5, 61, 3, 2, 2, 2, 7, 63, 3, 2, 2, 2, 9, 70, 3, 2, 2, 2,
	11, 76, 3, 2, 2, 2, 13, 83, 3, 2, 2, 2, 15, 85, 3, 2, 2, 2, 17, 99, 3,
	2, 2, 2, 19, 109, 3, 2, 2, 2, 21, 117, 3, 2, 2, 2, 23, 123, 3, 2, 2, 2,
	25, 127, 3, 2, 2, 2, 27, 133, 3, 2, 2, 2, 29, 140, 3, 2, 2, 2, 31, 151,
	3, 2, 2, 2, 33, 160, 3, 2, 2, 2, 35, 168, 3, 2, 2, 2, 37, 177, 3, 2, 2,
	2, 39, 184, 3, 2, 2, 2, 41, 194, 3, 2, 2, 2, 43, 199, 3, 2, 2, 2, 45, 215,
	3, 2, 2, 2, 47, 218, 3, 2, 2, 2, 49, 225, 3, 2, 2, 2, 51, 227, 3, 2, 2,
	2, 53, 229, 3, 2, 2, 2, 55, 243, 3, 2, 2, 2, 57, 254, 3, 2, 2, 2, 59, 60,
	7, 62, 2, 2, 60, 4, 3, 2, 2, 2, 61, 62, 7, 64, 2, 2, 62, 6, 3, 2, 2, 2,
	63, 64, 7, 105, 2, 2, 64, 65, 7, 110, 2, 2, 65, 66, 7, 113, 2, 2, 66, 67,
	7, 100, 2, 2, 67, 68, 7, 99, 2, 2, 68, 69, 7, 110, 2, 2, 69, 8, 3, 2, 2,
	2, 70, 71, 7, 105, 2, 2, 71, 72, 7, 116, 2, 2, 72, 73, 7, 113, 2, 2, 73,
	74, 7, 119, 2, 2, 74, 75, 7, 114, 2, 2, 75, 10, 3, 2, 2, 2, 76, 77, 7,
	116, 2, 2, 77, 78, 7, 103, 2, 2, 78, 79, 7, 105, 2, 2, 79, 80, 7, 107,
	2, 2, 80, 81, 7, 113, 2, 2, 81, 82, 7, 112, 2, 2, 82, 12, 3, 2, 2, 2, 83,
	84, 7, 63, 2, 2, 84, 14, 3, 2, 2, 2, 85, 86, 7, 99, 2, 2, 86, 87, 7, 111,
	2, 2, 87, 88, 7, 114, 2, 2, 88, 89, 7, 103, 2, 2, 89, 90, 7, 105, 2, 2,
	90, 91, 7, 97, 2, 2, 91, 92, 7, 116, 2, 2, 92, 93, 7, 103, 2, 2, 93, 94,
	7, 110, 2, 2, 94, 95, 7, 103, 2, 2, 95, 96, 7, 99, 2, 2, 96, 97, 7, 117,
	2, 2, 97, 98, 7, 103, 2, 2, 98, 16, 3, 2, 2, 2, 99, 100, 7, 100, 2, 2,
	100, 101, 7, 103, 2, 2, 101, 102, 7, 112, 2, 2, 102, 103, 7, 102, 2, 2,
	103, 104, 7, 97, 2, 2, 104, 105, 7, 102, 2, 2, 105, 106, 7, 113, 2, 2,
	106, 107, 7, 121, 2, 2, 107, 108, 7, 112, 2, 2, 108, 18, 3, 2, 2, 2, 109,
	110, 7, 100, 2, 2, 110, 111, 7, 103, 2, 2, 111, 112, 7, 112, 2, 2, 112,
	113, 7, 102, 2, 2, 113, 114, 7, 97, 2, 2, 114, 115, 7, 119, 2, 2, 115,
	116, 7, 114, 2, 2, 116, 20, 3, 2, 2, 2, 117, 118, 7, 106, 2, 2, 118, 119,
	7, 107, 2, 2, 119, 120, 7, 109, 2, 2, 120, 121, 7, 103, 2, 2, 121, 122,
	7, 123, 2, 2, 122, 22, 3, 2, 2, 2, 123, 124, 7, 109, 2, 2, 124, 125, 7,
	103, 2, 2, 125, 126, 7, 123, 2, 2, 126, 24, 3, 2, 2, 2, 127, 128, 7, 110,
	2, 2, 128, 129, 7, 113, 2, 2, 129, 130, 7, 109, 2, 2, 130, 131, 7, 103,
	2, 2, 131, 132, 7, 123, 2, 2, 132, 26, 3, 2, 2, 2, 133, 134, 7, 117, 2,
	2, 134, 135, 7, 99, 2, 2, 135, 136, 7, 111, 2, 2, 136, 137, 7, 114, 2,
	2, 137, 138, 7, 110, 2, 2, 138, 139, 7, 103, 2, 2, 139, 28, 3, 2, 2, 2,
	140, 141, 7, 117, 2, 2, 141, 142, 7, 121, 2, 2, 142, 143, 7, 97, 2, 2,
	143, 144, 7, 102, 2, 2, 144, 145, 7, 103, 2, 2, 145, 146, 7, 104, 2, 2,
	146, 147, 7, 99, 2, 2, 147, 148, 7, 119, 2, 2, 148, 149, 7, 110, 2, 2,
	149, 150, 7, 118, 2, 2, 150, 30, 3, 2, 2, 2, 151, 152, 7, 117, 2, 2, 152,
	153, 7, 121, 2, 2, 153, 154, 7, 97, 2, 2, 154, 155, 7, 106, 2, 2, 155,
	156, 7, 107, 2, 2, 156, 157, 7, 109, 2, 2, 157, 158, 7, 103, 2, 2, 158,
	159, 7, 123, 2, 2, 159, 32, 3, 2, 2, 2, 160, 161, 7, 117, 2, 2, 161, 162,
	7, 121, 2, 2, 162, 163, 7, 97, 2, 2, 163, 164, 7, 110, 2, 2, 164, 165,
	7, 99, 2, 2, 165, 166, 7, 117, 2, 2, 166, 167, 7, 118, 2, 2, 167, 34, 3,
	2, 2, 2, 168, 169, 7, 117, 2, 2, 169, 170, 7, 121, 2, 2, 170, 171, 7, 97,
	2, 2, 171, 172, 7, 110, 2, 2, 172, 173, 7, 113, 2, 2, 173, 174, 7, 109,
	2, 2, 174, 175, 7, 103, 2, 2, 175, 176, 7, 123, 2, 2, 176, 36, 3, 2, 2,
	2, 177, 179, 7, 44, 2, 2, 178, 180, 9, 2, 2, 2, 179, 178, 3, 2, 2, 2, 180,
	181, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 38, 3,
	2, 2, 2, 183, 185, 4, 50, 59, 2, 184, 183, 3, 2, 2, 2, 185, 186, 3, 2,
	2, 2, 186, 184, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 40, 3, 2, 2, 2,
	188, 195, 5, 39, 20, 2, 189, 191, 5, 39, 20, 2, 190, 189, 3, 2, 2, 2, 190,
	191, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 193, 7, 48, 2, 2, 193, 195,
	5, 39, 20, 2, 194, 188, 3, 2, 2, 2, 194, 190, 3, 2, 2, 2, 195, 42, 3, 2,
	2, 2, 196, 198, 5, 45, 23, 2, 197, 196, 3, 2, 2, 2, 198, 201, 3, 2, 2,
	2, 199, 197, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 202, 3, 2, 2, 2, 201,
	199, 3, 2, 2, 2, 202, 212, 5, 47, 24, 2, 203, 205, 5, 45, 23, 2, 204, 203,
	3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 206, 207, 3, 2,
	2, 2, 207, 208, 3, 2, 2, 2, 208, 209, 5, 47, 24, 2, 209, 211, 3, 2, 2,
	2, 210, 204, 3, 2, 2, 2, 211, 214, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 212,
	213, 3, 2, 2, 2, 213, 44, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2, 215, 216, 9,
	3, 2, 2, 216, 46, 3, 2, 2, 2, 217, 219, 9, 4, 2, 2, 218, 217, 3, 2, 2,
	2, 219, 220, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221,
	48, 3, 2, 2, 2, 222, 223, 7, 15, 2, 2, 223, 226, 7, 12, 2, 2, 224, 226,
	9, 5, 2, 2, 225, 222, 3, 2, 2, 2, 225, 224, 3, 2, 2, 2, 226, 50, 3, 2,
	2, 2, 227, 228, 9, 6, 2, 2, 228, 52, 3, 2, 2, 2, 229, 230, 7, 49, 2, 2,
	230, 231, 7, 44, 2, 2, 231, 235, 3, 2, 2, 2, 232, 234, 11, 2, 2, 2, 233,
	232, 3, 2, 2, 2, 234, 237, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 235, 233,
	3, 2, 2, 2, 236, 238, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 238, 239, 7, 44,
	2, 2, 239, 240, 7, 49, 2, 2, 240, 241, 3, 2, 2, 2, 241, 242, 8, 27, 2,
	2, 242, 54, 3, 2, 2, 2, 243, 244, 7, 49, 2, 2, 244, 245, 7, 49, 2, 2, 245,
	249, 3, 2, 2, 2, 246, 248, 10, 5, 2, 2, 247, 246, 3, 2, 2, 2, 248, 251,
	3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250, 252, 3, 2,
	2, 2, 251, 249, 3, 2, 2, 2, 252, 253, 8, 28, 2, 2, 253, 56, 3, 2, 2, 2,
	254, 258, 7, 37, 2, 2, 255, 257, 10, 5, 2, 2, 256, 255, 3, 2, 2, 2, 257,
	260, 3, 2, 2, 2, 258, 256, 3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 261,
	3, 2, 2, 2, 260, 258, 3, 2, 2, 2, 261, 262, 8, 29, 2, 2, 262, 58, 3, 2,
	2, 2, 15, 2, 181, 186, 190, 194, 199, 206, 212, 220, 225, 235, 249, 258,
	3, 8, 2, 2,
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
	"", "'<'", "'>'", "'global'", "'group'", "'region'", "'='", "'ampeg_release'",
	"'bend_down'", "'bend_up'", "'hikey'", "'key'", "'lokey'", "'sample'",
	"'sw_default'", "'sw_hikey'", "'sw_last'", "'sw_lokey'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"GENERATOR", "INT", "FLOAT", "PATH", "NEWLINE", "WHITESPACE", "BLOCK_COMMENT",
	"LINE_COMMENT", "HASH_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"GENERATOR", "INT", "FLOAT", "PATH", "SEPARATOR", "PATH_SEGMENT", "NEWLINE",
	"WHITESPACE", "BLOCK_COMMENT", "LINE_COMMENT", "HASH_COMMENT",
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
	SfzLexerT__6          = 7
	SfzLexerT__7          = 8
	SfzLexerT__8          = 9
	SfzLexerT__9          = 10
	SfzLexerT__10         = 11
	SfzLexerT__11         = 12
	SfzLexerT__12         = 13
	SfzLexerT__13         = 14
	SfzLexerT__14         = 15
	SfzLexerT__15         = 16
	SfzLexerT__16         = 17
	SfzLexerGENERATOR     = 18
	SfzLexerINT           = 19
	SfzLexerFLOAT         = 20
	SfzLexerPATH          = 21
	SfzLexerNEWLINE       = 22
	SfzLexerWHITESPACE    = 23
	SfzLexerBLOCK_COMMENT = 24
	SfzLexerLINE_COMMENT  = 25
	SfzLexerHASH_COMMENT  = 26
)
