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
	listener := New(t)
	p.AddParseListener(listener)
	p.Sfz()

	// test for some headers
	for _, v := range []string{"global", "group", "region"} {
		ass.True(listener.headers[v], "Header not found: "+v)
	}

	// test for some opcodes
	for _, v := range []string{"sample", "hikey", "lokey", "group", "off_by", "on_locc64", "on_hicc64"} {
		ass.True(listener.opcodes[v], "Opcode not found: "+v)
	}

	// test for some values
	for _, v := range []string{"arco\\arco_c1_pp_down.wav", "-3600", "12"} {
		ass.True(listener.values[v], "value not found: "+v)
	}

	// test one-line section
	found := 0
	for _, section := range listener.sections {
		if section.header == "group" {
			// <group> group=32 lokey=-1 on_locc64=126 on_hicc64=127 off_by=33 volume=-9
			if section.values["group"] == "33" {
				found += 1
				ass.Equal("-1", section.values["lokey"])
				ass.Equal("1", section.values["on_hicc64"])
				ass.Equal("0", section.values["on_locc64"])
				ass.Equal("-9", section.values["volume"])
			}
		}
	}
	ass.True(found > 0, "Still haven't found what I'm looking for")
}

type section struct {
	header        string
	currentOpcode string
	values        map[string]string
}

type sfzListener struct {
	*BaseSfzListener
	t              *testing.T
	currentSection *section
	sections       []*section
	headers        map[string]bool
	opcodes        map[string]bool
	values         map[string]bool
}

func New(t *testing.T) *sfzListener {
	return &sfzListener{
		t:       t,
		headers: make(map[string]bool),
		opcodes: make(map[string]bool),
		values:  make(map[string]bool),
	}
}

func (s *sfzListener) ExitHeader(ctx *HeaderContext) {
	header := ctx.GetText()
	log.Printf("header: %v", header)
	section := &section{header: header, values: make(map[string]string)}
	s.headers[header] = true
	s.sections = append(s.sections, section)
	s.currentSection = section
}

func (s *sfzListener) ExitOpcode(ctx *OpcodeContext) {
	opcode := ctx.GetText()
	log.Printf("opcode: %v", opcode)
	section := s.currentSection
	section.currentOpcode = opcode
	s.opcodes[opcode] = true
}

func (s *sfzListener) ExitValue(ctx *ValueContext) {
	value := ctx.GetText()
	log.Printf("value: %v", value)
	section := s.currentSection
	s.values[value] = true
	section.values[section.currentOpcode] = value
}
