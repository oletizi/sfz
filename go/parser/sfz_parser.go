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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 70, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 7, 2, 18, 10, 2, 12, 2, 14, 2, 21, 11, 2, 3, 2, 3, 2, 3,
	3, 7, 3, 26, 10, 3, 12, 3, 14, 3, 29, 11, 3, 3, 3, 3, 3, 7, 3, 33, 10,
	3, 12, 3, 14, 3, 36, 11, 3, 3, 3, 3, 3, 6, 3, 40, 10, 3, 13, 3, 14, 3,
	41, 3, 3, 7, 3, 45, 10, 3, 12, 3, 14, 3, 48, 11, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 58, 10, 6, 3, 6, 3, 6, 5, 6, 62, 10, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 2, 2, 9, 2, 4, 6, 8, 10, 12,
	14, 2, 6, 3, 2, 29, 30, 3, 2, 5, 7, 3, 2, 9, 24, 3, 2, 25, 28, 2, 70, 2,
	19, 3, 2, 2, 2, 4, 27, 3, 2, 2, 2, 6, 49, 3, 2, 2, 2, 8, 53, 3, 2, 2, 2,
	10, 55, 3, 2, 2, 2, 12, 65, 3, 2, 2, 2, 14, 67, 3, 2, 2, 2, 16, 18, 5,
	4, 3, 2, 17, 16, 3, 2, 2, 2, 18, 21, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19,
	20, 3, 2, 2, 2, 20, 22, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 22, 23, 7, 2, 2,
	3, 23, 3, 3, 2, 2, 2, 24, 26, 9, 2, 2, 2, 25, 24, 3, 2, 2, 2, 26, 29, 3,
	2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 30, 3, 2, 2, 2, 29,
	27, 3, 2, 2, 2, 30, 46, 5, 6, 4, 2, 31, 33, 9, 2, 2, 2, 32, 31, 3, 2, 2,
	2, 33, 36, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 37,
	3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 37, 45, 5, 6, 4, 2, 38, 40, 9, 2, 2, 2,
	39, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3,
	2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 45, 5, 10, 6, 2, 44, 34, 3, 2, 2, 2, 44,
	39, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2,
	2, 47, 5, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 50, 7, 3, 2, 2, 50, 51, 5,
	8, 5, 2, 51, 52, 7, 4, 2, 2, 52, 7, 3, 2, 2, 2, 53, 54, 9, 3, 2, 2, 54,
	9, 3, 2, 2, 2, 55, 57, 5, 12, 7, 2, 56, 58, 7, 30, 2, 2, 57, 56, 3, 2,
	2, 2, 57, 58, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 61, 7, 8, 2, 2, 60, 62,
	7, 30, 2, 2, 61, 60, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2,
	63, 64, 5, 14, 8, 2, 64, 11, 3, 2, 2, 2, 65, 66, 9, 4, 2, 2, 66, 13, 3,
	2, 2, 2, 67, 68, 9, 5, 2, 2, 68, 15, 3, 2, 2, 2, 10, 19, 27, 34, 41, 44,
	46, 57, 61,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'<'", "'>'", "'global'", "'group'", "'region'", "'='", "'ampeg_release'",
	"'bend_down'", "'bend_up'", "'hikey'", "'hivel'", "'key'", "'lokey'", "'lovel'",
	"'pitch_keycenter'", "'sample'", "'seq_length'", "'seq_position'", "'sw_default'",
	"'sw_hikey'", "'sw_last'", "'sw_lokey'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "GENERATOR", "INT", "FLOAT", "PATH", "NEWLINE", "WHITESPACE",
	"BLOCK_COMMENT", "LINE_COMMENT", "HASH_COMMENT",
}

var ruleNames = []string{
	"sfz", "sfzObject", "headerObject", "header", "opcodeStatement", "opcode",
	"value",
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
	SfzParserEOF           = antlr.TokenEOF
	SfzParserT__0          = 1
	SfzParserT__1          = 2
	SfzParserT__2          = 3
	SfzParserT__3          = 4
	SfzParserT__4          = 5
	SfzParserT__5          = 6
	SfzParserT__6          = 7
	SfzParserT__7          = 8
	SfzParserT__8          = 9
	SfzParserT__9          = 10
	SfzParserT__10         = 11
	SfzParserT__11         = 12
	SfzParserT__12         = 13
	SfzParserT__13         = 14
	SfzParserT__14         = 15
	SfzParserT__15         = 16
	SfzParserT__16         = 17
	SfzParserT__17         = 18
	SfzParserT__18         = 19
	SfzParserT__19         = 20
	SfzParserT__20         = 21
	SfzParserT__21         = 22
	SfzParserGENERATOR     = 23
	SfzParserINT           = 24
	SfzParserFLOAT         = 25
	SfzParserPATH          = 26
	SfzParserNEWLINE       = 27
	SfzParserWHITESPACE    = 28
	SfzParserBLOCK_COMMENT = 29
	SfzParserLINE_COMMENT  = 30
	SfzParserHASH_COMMENT  = 31
)

// SfzParser rules.
const (
	SfzParserRULE_sfz             = 0
	SfzParserRULE_sfzObject       = 1
	SfzParserRULE_headerObject    = 2
	SfzParserRULE_header          = 3
	SfzParserRULE_opcodeStatement = 4
	SfzParserRULE_opcode          = 5
	SfzParserRULE_value           = 6
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

func (s *SfzContext) AllSfzObject() []ISfzObjectContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISfzObjectContext)(nil)).Elem())
	var tst = make([]ISfzObjectContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISfzObjectContext)
		}
	}

	return tst
}

func (s *SfzContext) SfzObject(i int) ISfzObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISfzObjectContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISfzObjectContext)
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SfzParserT__0)|(1<<SfzParserNEWLINE)|(1<<SfzParserWHITESPACE))) != 0 {
		{
			p.SetState(14)
			p.SfzObject()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(20)
		p.Match(SfzParserEOF)
	}

	return localctx
}

// ISfzObjectContext is an interface to support dynamic dispatch.
type ISfzObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSfzObjectContext differentiates from other interfaces.
	IsSfzObjectContext()
}

type SfzObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySfzObjectContext() *SfzObjectContext {
	var p = new(SfzObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfzParserRULE_sfzObject
	return p
}

func (*SfzObjectContext) IsSfzObjectContext() {}

func NewSfzObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SfzObjectContext {
	var p = new(SfzObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfzParserRULE_sfzObject

	return p
}

func (s *SfzObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *SfzObjectContext) AllHeaderObject() []IHeaderObjectContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHeaderObjectContext)(nil)).Elem())
	var tst = make([]IHeaderObjectContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHeaderObjectContext)
		}
	}

	return tst
}

func (s *SfzObjectContext) HeaderObject(i int) IHeaderObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeaderObjectContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHeaderObjectContext)
}

func (s *SfzObjectContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(SfzParserWHITESPACE)
}

func (s *SfzObjectContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(SfzParserWHITESPACE, i)
}

func (s *SfzObjectContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(SfzParserNEWLINE)
}

func (s *SfzObjectContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(SfzParserNEWLINE, i)
}

func (s *SfzObjectContext) AllOpcodeStatement() []IOpcodeStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOpcodeStatementContext)(nil)).Elem())
	var tst = make([]IOpcodeStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOpcodeStatementContext)
		}
	}

	return tst
}

func (s *SfzObjectContext) OpcodeStatement(i int) IOpcodeStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpcodeStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOpcodeStatementContext)
}

func (s *SfzObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SfzObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SfzObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.EnterSfzObject(s)
	}
}

func (s *SfzObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.ExitSfzObject(s)
	}
}

func (p *SfzParser) SfzObject() (localctx ISfzObjectContext) {
	localctx = NewSfzObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SfzParserRULE_sfzObject)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SfzParserNEWLINE || _la == SfzParserWHITESPACE {
		{
			p.SetState(22)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SfzParserNEWLINE || _la == SfzParserWHITESPACE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	{
		p.SetState(28)
		p.HeaderObject()
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(42)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				p.SetState(32)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == SfzParserNEWLINE || _la == SfzParserWHITESPACE {
					{
						p.SetState(29)
						_la = p.GetTokenStream().LA(1)

						if !(_la == SfzParserNEWLINE || _la == SfzParserWHITESPACE) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

					p.SetState(34)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(35)
					p.HeaderObject()
				}

			case 2:
				p.SetState(37)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for ok := true; ok; ok = _la == SfzParserNEWLINE || _la == SfzParserWHITESPACE {
					{
						p.SetState(36)
						_la = p.GetTokenStream().LA(1)

						if !(_la == SfzParserNEWLINE || _la == SfzParserWHITESPACE) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

					p.SetState(39)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(41)
					p.OpcodeStatement()
				}

			}

		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IHeaderObjectContext is an interface to support dynamic dispatch.
type IHeaderObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeaderObjectContext differentiates from other interfaces.
	IsHeaderObjectContext()
}

type HeaderObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderObjectContext() *HeaderObjectContext {
	var p = new(HeaderObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfzParserRULE_headerObject
	return p
}

func (*HeaderObjectContext) IsHeaderObjectContext() {}

func NewHeaderObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderObjectContext {
	var p = new(HeaderObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfzParserRULE_headerObject

	return p
}

func (s *HeaderObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderObjectContext) Header() IHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHeaderContext)
}

func (s *HeaderObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.EnterHeaderObject(s)
	}
}

func (s *HeaderObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.ExitHeaderObject(s)
	}
}

func (p *SfzParser) HeaderObject() (localctx IHeaderObjectContext) {
	localctx = NewHeaderObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SfzParserRULE_headerObject)

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
		p.SetState(47)
		p.Match(SfzParserT__0)
	}
	{
		p.SetState(48)
		p.Header()
	}
	{
		p.SetState(49)
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
		p.SetState(51)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SfzParserT__2)|(1<<SfzParserT__3)|(1<<SfzParserT__4))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOpcodeStatementContext is an interface to support dynamic dispatch.
type IOpcodeStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpcodeStatementContext differentiates from other interfaces.
	IsOpcodeStatementContext()
}

type OpcodeStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpcodeStatementContext() *OpcodeStatementContext {
	var p = new(OpcodeStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfzParserRULE_opcodeStatement
	return p
}

func (*OpcodeStatementContext) IsOpcodeStatementContext() {}

func NewOpcodeStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpcodeStatementContext {
	var p = new(OpcodeStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfzParserRULE_opcodeStatement

	return p
}

func (s *OpcodeStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *OpcodeStatementContext) Opcode() IOpcodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpcodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpcodeContext)
}

func (s *OpcodeStatementContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *OpcodeStatementContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(SfzParserWHITESPACE)
}

func (s *OpcodeStatementContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(SfzParserWHITESPACE, i)
}

func (s *OpcodeStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpcodeStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpcodeStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.EnterOpcodeStatement(s)
	}
}

func (s *OpcodeStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfzListener); ok {
		listenerT.ExitOpcodeStatement(s)
	}
}

func (p *SfzParser) OpcodeStatement() (localctx IOpcodeStatementContext) {
	localctx = NewOpcodeStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SfzParserRULE_opcodeStatement)
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
		p.SetState(53)
		p.Opcode()
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SfzParserWHITESPACE {
		{
			p.SetState(54)
			p.Match(SfzParserWHITESPACE)
		}

	}
	{
		p.SetState(57)
		p.Match(SfzParserT__5)
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SfzParserWHITESPACE {
		{
			p.SetState(58)
			p.Match(SfzParserWHITESPACE)
		}

	}
	{
		p.SetState(61)
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
		p.SetState(63)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SfzParserT__6)|(1<<SfzParserT__7)|(1<<SfzParserT__8)|(1<<SfzParserT__9)|(1<<SfzParserT__10)|(1<<SfzParserT__11)|(1<<SfzParserT__12)|(1<<SfzParserT__13)|(1<<SfzParserT__14)|(1<<SfzParserT__15)|(1<<SfzParserT__16)|(1<<SfzParserT__17)|(1<<SfzParserT__18)|(1<<SfzParserT__19)|(1<<SfzParserT__20)|(1<<SfzParserT__21))) != 0) {
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

func (s *ValueContext) INT() antlr.TerminalNode {
	return s.GetToken(SfzParserINT, 0)
}

func (s *ValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SfzParserFLOAT, 0)
}

func (s *ValueContext) PATH() antlr.TerminalNode {
	return s.GetToken(SfzParserPATH, 0)
}

func (s *ValueContext) GENERATOR() antlr.TerminalNode {
	return s.GetToken(SfzParserGENERATOR, 0)
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
	{
		p.SetState(65)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SfzParserGENERATOR)|(1<<SfzParserINT)|(1<<SfzParserFLOAT)|(1<<SfzParserPATH))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
