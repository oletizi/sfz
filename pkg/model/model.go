package model

import (
	"fmt"
	parser2 "github.com/oletizi/sfz-parser/pkg/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Model represents the structure of a (parsed) SFZ file
type Model interface {
	Control() []Section
	Curve() []Section
	Effect() []Section
	Global() Section
	Group() []Section
	Master() []Section
	Midi() []Section
	Region() []Section
	Sample() []Section
}

func New(path string) (Model, error) {
	in, err := antlr.NewFileStream(path)
	if err != nil {
		return nil, err
	}
	lexer := parser2.NewSfzLexer(in)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser2.NewSfzParser(stream)
	listener := &sfzListener{}
	p.AddParseListener(listener)
	p.Sfz()
	config := &model{global: listener.global, sections: make(map[string][]Section)}
	for _, section := range listener.sections {
		sections := append(config.sections[section.Name()], section)
		config.sections[section.Name()] = sections
	}
	// TODO: map the sections in the listener to the model
	return config, nil
}

type model struct {
	sections map[string][]Section
	global   *section
}

func (c *model) Control() []Section {
	return c.sectionsByName("control")
}

func (c *model) Curve() []Section {
	return c.sectionsByName("curve")
}

func (c *model) Effect() []Section {
	return c.sectionsByName("effect")
}

func (c *model) Global() Section {
	return c.global
}

func (c *model) Group() []Section {
	return c.sectionsByName("group")
}

func (c *model) Master() []Section {
	return c.sectionsByName("master")
}

func (c *model) Midi() []Section {
	return c.sectionsByName("midi")
}

func (c *model) Region() []Section {
	return c.sectionsByName("region")
}

func (c *model) Sample() []Section {
	return c.sectionsByName("sample")
}

func (c *model) sectionsByName(name string) []Section {
	rv := c.sections[name]
	if rv == nil {
		rv = make([]Section, 0)
		c.sections[name] = rv
	}
	return rv
}

// Section represents a section of an sfz file in a Model
type Section interface {
	Name() string
	Value(string) string
}

type section struct {
	name   string
	values map[string]string
}

func newSection(name string) *section {
	return &section{name, make(map[string]string)}
}

func (s *section) Name() string {
	return s.name
}

func (s *section) Value(key string) string {
	return s.values[key]
}

// sfzListener is a listener for the antlr parser
type sfzListener struct {
	*parser2.BaseSfzListener
	currentSection *section
	global         *section
	sections       []Section
	currentOpcode  string
}

func (s *sfzListener) ExitHeader(ctx *parser2.HeaderContext) {
	header := ctx.GetText()
	var theSection *section
	if header == "global" {
		if s.global == nil {
			theSection = newSection(header)
			s.global = theSection
		}
	} else {
		theSection = newSection(header)
		s.sections = append(s.sections, theSection)
	}
	s.currentSection = theSection
}

func (s *sfzListener) ExitOpcode(ctx *parser2.OpcodeContext) {
	opcode := ctx.GetText()
	s.currentOpcode = opcode
}

func (s *sfzListener) ExitValue(ctx *parser2.ValueContext) {
	value := ctx.GetText()
	section := s.currentSection
	if section == nil {
		fmt.Printf("section is nil; current opcode: %v", s.currentOpcode)
	} else {
		section.values[s.currentOpcode] = value
	}
}
