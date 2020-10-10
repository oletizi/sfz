package parser

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/require"
)

func TestBasics(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	ass := require.New(t)
	sfzFile, err := filepath.Abs("../src/test/resources/test.sfz")
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

	// test for some headers
	for _, v := range []string{"global", "group", "region"} {
		ass.True(listener.headers[v], "Header not found: "+v)
	}

	// test for some opcodes
	for _, v := range []string{"sample", "hikey", "lokey", "group", "off_by"} {
		ass.True(listener.opcodes[v], "Opcode not found: "+v)
	}

	// test for some values
	for _, v := range []string{"arco\\arco_c1_pp_down.wav", "-3600", "12"} {
		ass.True(listener.values[v], "value not found: "+v)
	}
}

type sfzListener struct {
	*BaseSfzListener
	headers map[string]bool
	opcodes map[string]bool
	values  map[string]bool
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
	log.Printf("header: %v", header)
	s.headers[header] = true
}

func (s *sfzListener) ExitOpcode(ctx *OpcodeContext) {
	opcode := ctx.GetText()
	log.Printf("opcode: %v", opcode)
	s.opcodes[opcode] = true
}

func (s *sfzListener) ExitValue(ctx *ValueContext) {
	value := ctx.GetText()
	log.Printf("value: %v", value)
	s.values[value] = true
}
