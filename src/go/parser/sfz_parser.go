// Code generated from Sfz.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Sfz

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 35, 56, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 6, 2, 18, 10, 2, 13, 2, 14, 2, 19, 3, 2, 3, 2, 3, 3, 3,
	3, 5, 3, 26, 10, 3, 3, 3, 6, 3, 29, 10, 3, 13, 3, 14, 3, 30, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 41, 10, 6, 3, 6, 3, 6, 5, 6,
	45, 10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 6, 8, 52, 10, 8, 13, 8, 14, 8,
	53, 3, 8, 2, 2, 9, 2, 4, 6, 8, 10, 12, 14, 2, 4, 3, 2, 5, 13, 3, 2, 15,
	29, 2, 54, 2, 17, 3, 2, 2, 2, 4, 25, 3, 2, 2, 2, 6, 32, 3, 2, 2, 2, 8,
	36, 3, 2, 2, 2, 10, 38, 3, 2, 2, 2, 12, 48, 3, 2, 2, 2, 14, 51, 3, 2, 2,
	2, 16, 18, 5, 4, 3, 2, 17, 16, 3, 2, 2, 2, 18, 19, 3, 2, 2, 2, 19, 17,
	3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 22, 7, 2, 2, 3,
	22, 3, 3, 2, 2, 2, 23, 26, 5, 6, 4, 2, 24, 26, 5, 10, 6, 2, 25, 23, 3,
	2, 2, 2, 25, 24, 3, 2, 2, 2, 26, 28, 3, 2, 2, 2, 27, 29, 7, 32, 2, 2, 28,
	27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2,
	2, 31, 5, 3, 2, 2, 2, 32, 33, 7, 3, 2, 2, 33, 34, 5, 8, 5, 2, 34, 35, 7,
	4, 2, 2, 35, 7, 3, 2, 2, 2, 36, 37, 9, 2, 2, 2, 37, 9, 3, 2, 2, 2, 38,
	40, 5, 12, 7, 2, 39, 41, 7, 33, 2, 2, 40, 39, 3, 2, 2, 2, 40, 41, 3, 2,
	2, 2, 41, 42, 3, 2, 2, 2, 42, 44, 7, 14, 2, 2, 43, 45, 7, 33, 2, 2, 44,
	43, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 47, 5, 14,
	8, 2, 47, 11, 3, 2, 2, 2, 48, 49, 9, 3, 2, 2, 49, 13, 3, 2, 2, 2, 50, 52,
	7, 31, 2, 2, 51, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2,
	53, 54, 3, 2, 2, 2, 54, 15, 3, 2, 2, 2, 8, 19, 25, 30, 40, 44, 53,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'<'", "'>'", "'region'", "'group'", "'control'", "'global'", "'curve'",
	"'effect'", "'master'", "'midi'", "'sampler'", "'='", "'ampeg_release'",
	"'bend_down'", "'bend_up'", "'hikey'", "'hivel'", "'lokey'", "'lovel'",
	"'pitch_keycenter'", "'sample'", "'seq_length'", "'seq_position'", "'sw_default'",
	"'sw_hikey'", "'sw_last'", "'sw_lokey'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "Digit", "Text", "Newline", "Whitespace",
	"BlockComment", "LineComment",
}

var ruleNames = []string{
	"sfz", "line", "header_stmt", "header", "opcode_stmt", "opcode", "value",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SfzParser struct {
	*antlr.BaseParser
}

func NewSfzParser(input antlr.TokenStream) *SfzParser {
	this := new(SfzParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Sfz.g4"

	return this
}

// SfzParser tokens.
const (
	SfzParserEOF          = antlr.TokenEOF
	SfzParserT__0         = 1
	SfzParserT__1         = 2
	SfzParserT__2         = 3
	SfzParserT__3         = 4
	SfzParserT__4         = 5
	SfzParserT__5         = 6
	SfzParserT__6         = 7
	SfzParserT__7         = 8
	SfzParserT__8         = 9
	SfzParserT__9         = 10
	SfzParserT__10        = 11
	SfzParserT__11        = 12
	SfzParserT__12        = 13
	SfzParserT__13        = 14
	SfzParserT__14        = 15
	SfzParserT__15        = 16
	SfzParserT__16        = 17
	SfzParserT__17        = 18
	SfzParserT__18        = 19
	SfzParserT__19        = 20
	SfzParserT__20        = 21
	SfzParserT__21        = 22
	SfzParserT__22        = 23
	SfzParserT__23        = 24
	SfzParserT__24        = 25
	SfzParserT__25        = 26
	SfzParserT__26        = 27
	SfzParserDigit        = 28
	SfzParserText         = 29
	SfzParserNewline      = 30
	SfzParserWhitespace   = 31
	SfzParserBlockComment = 32
	SfzParserLineComment  = 33
)

// SfzParser rules.
const (
	SfzParserRULE_sfz         = 0
	SfzParserRULE_line        = 1
	SfzParserRULE_header_stmt = 2
	SfzParserRULE_header      = 3
	SfzParserRULE_opcode_stmt = 4
	SfzParserRULE_opcode      = 5
	SfzParserRULE_value       = 6
)

// ISfzContext is an interface to support dynamic dispatch.
type ISfzContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSfzContext differentiates from other interfaces.
	IsSfzContext()
}

type SfzContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySfzContext() *SfzContext {
	var p = new(SfzContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfzParserRULE_sfz
	return p
}

func (*SfzContext) IsSfzContext() {}

func NewSfzContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SfzContext {
	var p = new(SfzContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfzParserRULE_sfz

	return p
}

func (s *SfzContext) GetParser() antlr.Parser { return s.parser }

func (s *SfzContext) EOF() antlr.TerminalNode {
	return s.GetToken(SfzParserEOF, 0)
}

func (s *SfzContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *SfzContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *SfzContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SfzContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SfzContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.EnterSfz(s)
	}
}

func (s *SfzContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.ExitSfz(s)
	}
}

func (p *SfzParser) Sfz() (localctx ISfzContext) {
	localctx = NewSfzContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SfzParserRULE_sfz)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SfzParserT__0)|(1<<SfzParserT__12)|(1<<SfzParserT__13)|(1<<SfzParserT__14)|(1<<SfzParserT__15)|(1<<SfzParserT__16)|(1<<SfzParserT__17)|(1<<SfzParserT__18)|(1<<SfzParserT__19)|(1<<SfzParserT__20)|(1<<SfzParserT__21)|(1<<SfzParserT__22)|(1<<SfzParserT__23)|(1<<SfzParserT__24)|(1<<SfzParserT__25)|(1<<SfzParserT__26))) != 0) {
		{
			p.SetState(14)
			p.Line()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(19)
		p.Match(SfzParserEOF)
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfzParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfzParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Header_stmt() IHeader_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeader_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHeader_stmtContext)
}

func (s *LineContext) Opcode_stmt() IOpcode_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpcode_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpcode_stmtContext)
}

func (s *LineContext) AllNewline() []antlr.TerminalNode {
	return s.GetTokens(SfzParserNewline)
}

func (s *LineContext) Newline(i int) antlr.TerminalNode {
	return s.GetToken(SfzParserNewline, i)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *SfzParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SfzParserRULE_line)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SfzParserT__0:
		{
			p.SetState(21)
			p.Header_stmt()
		}

	case SfzParserT__12, SfzParserT__13, SfzParserT__14, SfzParserT__15, SfzParserT__16, SfzParserT__17, SfzParserT__18, SfzParserT__19, SfzParserT__20, SfzParserT__21, SfzParserT__22, SfzParserT__23, SfzParserT__24, SfzParserT__25, SfzParserT__26:
		{
			p.SetState(22)
			p.Opcode_stmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SfzParserNewline {
		{
			p.SetState(25)
			p.Match(SfzParserNewline)
		}

		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHeader_stmtContext is an interface to support dynamic dispatch.
type IHeader_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeader_stmtContext differentiates from other interfaces.
	IsHeader_stmtContext()
}

type Header_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeader_stmtContext() *Header_stmtContext {
	var p = new(Header_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfzParserRULE_header_stmt
	return p
}

func (*Header_stmtContext) IsHeader_stmtContext() {}

func NewHeader_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Header_stmtContext {
	var p = new(Header_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfzParserRULE_header_stmt

	return p
}

func (s *Header_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Header_stmtContext) Header() IHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHeaderContext)
}

func (s *Header_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Header_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Header_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.EnterHeader_stmt(s)
	}
}

func (s *Header_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.ExitHeader_stmt(s)
	}
}

func (p *SfzParser) Header_stmt() (localctx IHeader_stmtContext) {
	localctx = NewHeader_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SfzParserRULE_header_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.Match(SfzParserT__0)
	}
	{
		p.SetState(31)
		p.Header()
	}
	{
		p.SetState(32)
		p.Match(SfzParserT__1)
	}

	return localctx
}

// IHeaderContext is an interface to support dynamic dispatch.
type IHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeaderContext differentiates from other interfaces.
	IsHeaderContext()
}

type HeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderContext() *HeaderContext {
	var p = new(HeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfzParserRULE_header
	return p
}

func (*HeaderContext) IsHeaderContext() {}

func NewHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderContext {
	var p = new(HeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfzParserRULE_header

	return p
}

func (s *HeaderContext) GetParser() antlr.Parser { return s.parser }
func (s *HeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.EnterHeader(s)
	}
}

func (s *HeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.ExitHeader(s)
	}
}

func (p *SfzParser) Header() (localctx IHeaderContext) {
	localctx = NewHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SfzParserRULE_header)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SfzParserT__2)|(1<<SfzParserT__3)|(1<<SfzParserT__4)|(1<<SfzParserT__5)|(1<<SfzParserT__6)|(1<<SfzParserT__7)|(1<<SfzParserT__8)|(1<<SfzParserT__9)|(1<<SfzParserT__10))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOpcode_stmtContext is an interface to support dynamic dispatch.
type IOpcode_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpcode_stmtContext differentiates from other interfaces.
	IsOpcode_stmtContext()
}

type Opcode_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpcode_stmtContext() *Opcode_stmtContext {
	var p = new(Opcode_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfzParserRULE_opcode_stmt
	return p
}

func (*Opcode_stmtContext) IsOpcode_stmtContext() {}

func NewOpcode_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Opcode_stmtContext {
	var p = new(Opcode_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfzParserRULE_opcode_stmt

	return p
}

func (s *Opcode_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Opcode_stmtContext) Opcode() IOpcodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpcodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpcodeContext)
}

func (s *Opcode_stmtContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Opcode_stmtContext) AllWhitespace() []antlr.TerminalNode {
	return s.GetTokens(SfzParserWhitespace)
}

func (s *Opcode_stmtContext) Whitespace(i int) antlr.TerminalNode {
	return s.GetToken(SfzParserWhitespace, i)
}

func (s *Opcode_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Opcode_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Opcode_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.EnterOpcode_stmt(s)
	}
}

func (s *Opcode_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.ExitOpcode_stmt(s)
	}
}

func (p *SfzParser) Opcode_stmt() (localctx IOpcode_stmtContext) {
	localctx = NewOpcode_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SfzParserRULE_opcode_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Opcode()
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SfzParserWhitespace {
		{
			p.SetState(37)
			p.Match(SfzParserWhitespace)
		}

	}
	{
		p.SetState(40)
		p.Match(SfzParserT__11)
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SfzParserWhitespace {
		{
			p.SetState(41)
			p.Match(SfzParserWhitespace)
		}

	}
	{
		p.SetState(44)
		p.Value()
	}

	return localctx
}

// IOpcodeContext is an interface to support dynamic dispatch.
type IOpcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpcodeContext differentiates from other interfaces.
	IsOpcodeContext()
}

type OpcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpcodeContext() *OpcodeContext {
	var p = new(OpcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfzParserRULE_opcode
	return p
}

func (*OpcodeContext) IsOpcodeContext() {}

func NewOpcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpcodeContext {
	var p = new(OpcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfzParserRULE_opcode

	return p
}

func (s *OpcodeContext) GetParser() antlr.Parser { return s.parser }
func (s *OpcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.EnterOpcode(s)
	}
}

func (s *OpcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.ExitOpcode(s)
	}
}

func (p *SfzParser) Opcode() (localctx IOpcodeContext) {
	localctx = NewOpcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SfzParserRULE_opcode)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SfzParserT__12)|(1<<SfzParserT__13)|(1<<SfzParserT__14)|(1<<SfzParserT__15)|(1<<SfzParserT__16)|(1<<SfzParserT__17)|(1<<SfzParserT__18)|(1<<SfzParserT__19)|(1<<SfzParserT__20)|(1<<SfzParserT__21)|(1<<SfzParserT__22)|(1<<SfzParserT__23)|(1<<SfzParserT__24)|(1<<SfzParserT__25)|(1<<SfzParserT__26))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfzParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfzParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) AllText() []antlr.TerminalNode {
	return s.GetTokens(SfzParserText)
}

func (s *ValueContext) Text(i int) antlr.TerminalNode {
	return s.GetToken(SfzParserText, i)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *SfzParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SfzParserRULE_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SfzParserText {
		{
			p.SetState(48)
			p.Match(SfzParserText)
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
