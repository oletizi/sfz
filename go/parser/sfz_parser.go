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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 61, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 7, 2, 16, 10, 2, 12, 2, 14, 2, 19, 11, 2, 3, 2, 3, 2, 3, 3, 7, 3, 24,
	10, 3, 12, 3, 14, 3, 27, 11, 3, 3, 3, 3, 3, 7, 3, 31, 10, 3, 12, 3, 14,
	3, 34, 11, 3, 3, 3, 3, 3, 6, 3, 38, 10, 3, 13, 3, 14, 3, 39, 3, 3, 7, 3,
	43, 10, 3, 12, 3, 14, 3, 46, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 5, 7, 59, 10, 7, 3, 7, 2, 2, 8, 2, 4, 6,
	8, 10, 12, 2, 4, 3, 2, 10, 11, 3, 2, 5, 7, 2, 61, 2, 17, 3, 2, 2, 2, 4,
	25, 3, 2, 2, 2, 6, 47, 3, 2, 2, 2, 8, 51, 3, 2, 2, 2, 10, 53, 3, 2, 2,
	2, 12, 58, 3, 2, 2, 2, 14, 16, 5, 4, 3, 2, 15, 14, 3, 2, 2, 2, 16, 19,
	3, 2, 2, 2, 17, 15, 3, 2, 2, 2, 17, 18, 3, 2, 2, 2, 18, 20, 3, 2, 2, 2,
	19, 17, 3, 2, 2, 2, 20, 21, 7, 2, 2, 3, 21, 3, 3, 2, 2, 2, 22, 24, 9, 2,
	2, 2, 23, 22, 3, 2, 2, 2, 24, 27, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 26,
	3, 2, 2, 2, 26, 28, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 28, 44, 5, 6, 4, 2,
	29, 31, 9, 2, 2, 2, 30, 29, 3, 2, 2, 2, 31, 34, 3, 2, 2, 2, 32, 30, 3,
	2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 35, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 35,
	43, 5, 6, 4, 2, 36, 38, 9, 2, 2, 2, 37, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2,
	2, 39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 43,
	5, 10, 6, 2, 42, 32, 3, 2, 2, 2, 42, 37, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2,
	44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 5, 3, 2, 2, 2, 46, 44, 3, 2,
	2, 2, 47, 48, 7, 3, 2, 2, 48, 49, 5, 8, 5, 2, 49, 50, 7, 4, 2, 2, 50, 7,
	3, 2, 2, 2, 51, 52, 9, 3, 2, 2, 52, 9, 3, 2, 2, 2, 53, 54, 5, 12, 7, 2,
	54, 55, 7, 8, 2, 2, 55, 56, 7, 9, 2, 2, 56, 11, 3, 2, 2, 2, 57, 59, 7,
	9, 2, 2, 58, 57, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 13, 3, 2, 2, 2, 9,
	17, 25, 32, 39, 42, 44, 58,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'<'", "'>'", "'global'", "'group'", "'region'", "'='",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "STRING", "NEWLINE", "WHITESPACE", "BLOCK_COMMENT",
	"LINE_COMMENT", "HASH_COMMENT",
}

var ruleNames = []string{
	"sfz", "sfzObject", "headerObject", "header", "opcodeStatement", "opcode",
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
	SfzParserSTRING        = 7
	SfzParserNEWLINE       = 8
	SfzParserWHITESPACE    = 9
	SfzParserBLOCK_COMMENT = 10
	SfzParserLINE_COMMENT  = 11
	SfzParserHASH_COMMENT  = 12
)

// SfzParser rules.
const (
	SfzParserRULE_sfz             = 0
	SfzParserRULE_sfzObject       = 1
	SfzParserRULE_headerObject    = 2
	SfzParserRULE_header          = 3
	SfzParserRULE_opcodeStatement = 4
	SfzParserRULE_opcode          = 5
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SfzParserT__0)|(1<<SfzParserNEWLINE)|(1<<SfzParserWHITESPACE))) != 0 {
		{
			p.SetState(12)
			p.SfzObject()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(18)
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
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SfzParserNEWLINE || _la == SfzParserWHITESPACE {
		{
			p.SetState(20)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SfzParserNEWLINE || _la == SfzParserWHITESPACE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	{
		p.SetState(26)
		p.HeaderObject()
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(40)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				p.SetState(30)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == SfzParserNEWLINE || _la == SfzParserWHITESPACE {
					{
						p.SetState(27)
						_la = p.GetTokenStream().LA(1)

						if !(_la == SfzParserNEWLINE || _la == SfzParserWHITESPACE) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

					p.SetState(32)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(33)
					p.HeaderObject()
				}

			case 2:
				p.SetState(35)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for ok := true; ok; ok = _la == SfzParserNEWLINE || _la == SfzParserWHITESPACE {
					{
						p.SetState(34)
						_la = p.GetTokenStream().LA(1)

						if !(_la == SfzParserNEWLINE || _la == SfzParserWHITESPACE) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

					p.SetState(37)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(39)
					p.OpcodeStatement()
				}

			}

		}
		p.SetState(44)
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
		p.SetState(45)
		p.Match(SfzParserT__0)
	}
	{
		p.SetState(46)
		p.Header()
	}
	{
		p.SetState(47)
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
		p.SetState(49)
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

func (s *OpcodeStatementContext) STRING() antlr.TerminalNode {
	return s.GetToken(SfzParserSTRING, 0)
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
		p.Opcode()
	}
	{
		p.SetState(52)
		p.Match(SfzParserT__5)
	}
	{
		p.SetState(53)
		p.Match(SfzParserSTRING)
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

func (s *OpcodeContext) STRING() antlr.TerminalNode {
	return s.GetToken(SfzParserSTRING, 0)
}

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
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SfzParserSTRING {
		{
			p.SetState(55)
			p.Match(SfzParserSTRING)
		}

	}

	return localctx
}
