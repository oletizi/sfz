package config

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/oletizi/sfz-parser/parser"
)

//
// Config
//

type Config interface {
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

func New(path string) (Config, error) {
	in, err := antlr.NewFileStream(path)
	if err != nil {
		return nil, err
	}
	lexer := parser.NewSfzLexer(in)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSfzParser(stream)
	listener := &sfzListener{}
	p.AddParseListener(listener)
	p.Sfz()
	config := &config{global: listener.global, sections: make(map[string][]Section)}
	for _, section := range listener.sections {
		sections := append(config.sections[section.Name()], section)
		config.sections[section.Name()] = sections
	}
	// TODO: map the sections in the listener to the config
	return config, nil
}

type config struct {
	sections map[string][]Section
	global   *section
}

func (c *config) Control() []Section {
	return c.sectionsByName("control")
}

func (c *config) Curve() []Section {
	return c.sectionsByName("curve")
}

func (c *config) Effect() []Section {
	return c.sectionsByName("effect")
}

func (c *config) Global() Section {
	return c.global
}

func (c *config) Group() []Section {
	return c.sectionsByName("group")
}

func (c *config) Master() []Section {
	return c.sectionsByName("master")
}

func (c *config) Midi() []Section {
	return c.sectionsByName("midi")
}

func (c *config) Region() []Section {
	return c.sectionsByName("region")
}

func (c *config) Sample() []Section {
	return c.sectionsByName("sample")
}

func (c *config) sectionsByName(name string) []Section {
	rv := c.sections[name]
	if rv == nil {
		rv = make([]Section, 0)
		c.sections[name] = rv
	}
	return rv
}

//
// Section
//

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

//
// Parser listener
//
type sfzListener struct {
	*parser.BaseSfzListener
	currentSection *section
	global         *section
	sections       []Section
	currentOpcode  string
}

func (s *sfzListener) ExitHeader(ctx *parser.HeaderContext) {
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

func (s *sfzListener) ExitOpcode(ctx *parser.OpcodeContext) {
	opcode := ctx.GetText()
	s.currentOpcode = opcode
}

func (s *sfzListener) ExitValue(ctx *parser.ValueContext) {
	value := ctx.GetText()
	section := s.currentSection
	if section == nil {
		fmt.Printf("section is nil; current opcode: %v", s.currentOpcode)
	} else {
		section.values[s.currentOpcode] = value
	}

}
