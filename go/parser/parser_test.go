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
	sfzFile, err := filepath.Abs("../../src/test/resources/foo.sfz")
	ass.Nil(err)

	log.Printf("sfzFile: %v", sfzFile)

	in, err := antlr.NewFileStream(sfzFile)
	ass.Nil(err)
	lexer := NewSfzLexer(in)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSfzParser(stream)
	var listener SfzListener = &sfzListener{}
	p.AddParseListener(listener)
	p.Sfz()
	//antlr.ParseTreeWalkerDefault.Walk(listener, p.Sfz())
}

type sfzListener struct {
	*BaseSfzListener
}

func (s *sfzListener) ExitHeaderName(ctx *HeaderNameContext) {
	log.Printf("Header: %v", ctx.GetText())
}

func (s *sfzListener) ExitOpcode(ctx *OpcodeContext) {
	log.Printf("opcode: %v", ctx.GetText())
}

func (s *sfzListener) ExitValue(ctx *ValueContext) {
	log.Printf("value: %v", ctx.GetText())
}

//func (s sfzListener) ExitHeader(c *HeaderContext) {
//	log.Printf("Exit header: %v", c.GetText())
//}
/*
func (s sfzListener) ExitInt_opcode(c *Int_opcodeContext) {
	log.Printf("Int opcode: %v", c.GetText())
}

func (s sfzListener) ExitFloat_opcode(c *Float_opcodeContext) {
	log.Printf("Float opcode: %v", c.GetText())
}

func (s sfzListener) ExitText_opcode(c *Text_opcodeContext) {
	log.Printf("Text opcode: %v", c.GetText())
}

func (s sfzListener) ExitInt_value(c *Int_valueContext) {
	log.Printf("int value: %v", c.GetText())
}

func (s sfzListener) ExitFloat_value(c *Float_valueContext) {
	log.Printf("float value: %v", c.GetText())
}

func (s sfzListener) ExitText_value(c *Text_valueContext) {
	log.Printf("text value: %v", c.GetText())
}
*/
