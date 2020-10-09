package parser

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/require"
)

func TestBasics(t *testing.T) {
	ass := require.New(t)
	sfzFile, err := filepath.Abs("../../src/test/resources/test-big.sfz")
	ass.Nil(err)

	log.Printf("sfzFile: %v", sfzFile)

	in, err := antlr.NewFileStream(sfzFile)
	ass.Nil(err)
	lexer := NewSfzLexer(in)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSfzParser(stream)
	listener := New()
	p.AddParseListener(listener)
	p.Sfz()

	ass.NotNil(listener.sections)
	ass.NotEqual(0, len(listener.sections))

	for _, v := range SupportedOpcodes() {
		ass.True(listener.opcodes[v], "opcode not found: %v", v)
	}
}

type section struct {
	header      string
	currentPair *nvpair
	pairs       []nvpair
}

type nvpair struct {
	name  string
	value string
}

type sfzListener struct {
	*BaseSfzListener
	currentSection *section
	sections       []section
	headers        map[string]bool
	opcodes        map[string]bool
	values         map[string]bool
}

func New() *sfzListener {
	return &sfzListener{
		headers: make(map[string]bool),
		opcodes: make(map[string]bool),
		values:  make(map[string]bool),
	}
}

func (s *sfzListener) ExitHeader(ctx *HeaderContext) {
	header := ctx.GetText()
	s.headers[header] = true

	current := &section{header: header}
	s.currentSection = current
	s.sections = append(s.sections, *current)
}
func (s *sfzListener) ExitOpcodeStatement(ctx *OpcodeStatementContext) {
	log.Printf("opcode statement; %v", ctx.GetText())
}

func (s *sfzListener) ExitOpcode(ctx *OpcodeContext) {
	log.Printf("opcode: %v", ctx.GetText())
	opcode := ctx.GetText()
	s.opcodes[opcode] = true

	currentPair := &nvpair{name: opcode}
	currentSection := s.currentSection
	currentSection.currentPair = currentPair
	currentSection.pairs = append(currentSection.pairs, *currentPair)
}

//func (s *sfzListener) ExitValue(ctx *ValueContext) {
//	value := ctx.GetText()
//	s.values[value] = true
//
//	currentSection := s.currentSection
//	currentSection.currentPair.value = value
//}
