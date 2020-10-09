// Generated from Sfz.g4 by ANTLR 4.8
package org.letizi.sfz.parser;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SfzParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, STRING=4, NEWLINE=5, WHITESPACE=6, BLOCK_COMMENT=7, 
		LINE_COMMENT=8, HASH_COMMENT=9;
	public static final int
		RULE_sfz = 0, RULE_sfzObject = 1, RULE_headerObject = 2, RULE_header = 3, 
		RULE_opcodeStatement = 4, RULE_opcode = 5, RULE_value = 6;
	private static String[] makeRuleNames() {
		return new String[] {
			"sfz", "sfzObject", "headerObject", "header", "opcodeStatement", "opcode", 
			"value"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'<'", "'>'", "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, "STRING", "NEWLINE", "WHITESPACE", "BLOCK_COMMENT", 
			"LINE_COMMENT", "HASH_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Sfz.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SfzParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class SfzContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(SfzParser.EOF, 0); }
		public List<SfzObjectContext> sfzObject() {
			return getRuleContexts(SfzObjectContext.class);
		}
		public SfzObjectContext sfzObject(int i) {
			return getRuleContext(SfzObjectContext.class,i);
		}
		public SfzContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sfz; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).enterSfz(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).exitSfz(this);
		}
	}

	public final SfzContext sfz() throws RecognitionException {
		SfzContext _localctx = new SfzContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_sfz);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(17);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << NEWLINE) | (1L << WHITESPACE))) != 0)) {
				{
				{
				setState(14);
				sfzObject();
				}
				}
				setState(19);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(20);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SfzObjectContext extends ParserRuleContext {
		public List<HeaderObjectContext> headerObject() {
			return getRuleContexts(HeaderObjectContext.class);
		}
		public HeaderObjectContext headerObject(int i) {
			return getRuleContext(HeaderObjectContext.class,i);
		}
		public List<TerminalNode> WHITESPACE() { return getTokens(SfzParser.WHITESPACE); }
		public TerminalNode WHITESPACE(int i) {
			return getToken(SfzParser.WHITESPACE, i);
		}
		public List<TerminalNode> NEWLINE() { return getTokens(SfzParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(SfzParser.NEWLINE, i);
		}
		public List<OpcodeStatementContext> opcodeStatement() {
			return getRuleContexts(OpcodeStatementContext.class);
		}
		public OpcodeStatementContext opcodeStatement(int i) {
			return getRuleContext(OpcodeStatementContext.class,i);
		}
		public SfzObjectContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sfzObject; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).enterSfzObject(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).exitSfzObject(this);
		}
	}

	public final SfzObjectContext sfzObject() throws RecognitionException {
		SfzObjectContext _localctx = new SfzObjectContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_sfzObject);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(25);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==NEWLINE || _la==WHITESPACE) {
				{
				{
				setState(22);
				_la = _input.LA(1);
				if ( !(_la==NEWLINE || _la==WHITESPACE) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				}
				setState(27);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			{
			setState(28);
			headerObject();
			setState(44);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					setState(42);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
					case 1:
						{
						setState(32);
						_errHandler.sync(this);
						_la = _input.LA(1);
						while (_la==NEWLINE || _la==WHITESPACE) {
							{
							{
							setState(29);
							_la = _input.LA(1);
							if ( !(_la==NEWLINE || _la==WHITESPACE) ) {
							_errHandler.recoverInline(this);
							}
							else {
								if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
								_errHandler.reportMatch(this);
								consume();
							}
							}
							}
							setState(34);
							_errHandler.sync(this);
							_la = _input.LA(1);
						}
						setState(35);
						headerObject();
						}
						break;
					case 2:
						{
						setState(37); 
						_errHandler.sync(this);
						_la = _input.LA(1);
						do {
							{
							{
							setState(36);
							_la = _input.LA(1);
							if ( !(_la==NEWLINE || _la==WHITESPACE) ) {
							_errHandler.recoverInline(this);
							}
							else {
								if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
								_errHandler.reportMatch(this);
								consume();
							}
							}
							}
							setState(39); 
							_errHandler.sync(this);
							_la = _input.LA(1);
						} while ( _la==NEWLINE || _la==WHITESPACE );
						setState(41);
						opcodeStatement();
						}
						break;
					}
					} 
				}
				setState(46);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			}
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class HeaderObjectContext extends ParserRuleContext {
		public HeaderContext header() {
			return getRuleContext(HeaderContext.class,0);
		}
		public HeaderObjectContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_headerObject; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).enterHeaderObject(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).exitHeaderObject(this);
		}
	}

	public final HeaderObjectContext headerObject() throws RecognitionException {
		HeaderObjectContext _localctx = new HeaderObjectContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_headerObject);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(47);
			match(T__0);
			setState(48);
			header();
			setState(49);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class HeaderContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(SfzParser.STRING, 0); }
		public HeaderContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_header; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).enterHeader(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).exitHeader(this);
		}
	}

	public final HeaderContext header() throws RecognitionException {
		HeaderContext _localctx = new HeaderContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_header);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(52);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==STRING) {
				{
				setState(51);
				match(STRING);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OpcodeStatementContext extends ParserRuleContext {
		public OpcodeContext opcode() {
			return getRuleContext(OpcodeContext.class,0);
		}
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public OpcodeStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_opcodeStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).enterOpcodeStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).exitOpcodeStatement(this);
		}
	}

	public final OpcodeStatementContext opcodeStatement() throws RecognitionException {
		OpcodeStatementContext _localctx = new OpcodeStatementContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_opcodeStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(54);
			opcode();
			setState(55);
			match(T__2);
			setState(56);
			value();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OpcodeContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(SfzParser.STRING, 0); }
		public OpcodeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_opcode; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).enterOpcode(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).exitOpcode(this);
		}
	}

	public final OpcodeContext opcode() throws RecognitionException {
		OpcodeContext _localctx = new OpcodeContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_opcode);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(59);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==STRING) {
				{
				setState(58);
				match(STRING);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ValueContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(SfzParser.STRING, 0); }
		public ValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_value; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).enterValue(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).exitValue(this);
		}
	}

	public final ValueContext value() throws RecognitionException {
		ValueContext _localctx = new ValueContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_value);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(62);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==STRING) {
				{
				setState(61);
				match(STRING);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\13C\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\3\2\7\2\22\n\2\f\2\16\2\25"+
		"\13\2\3\2\3\2\3\3\7\3\32\n\3\f\3\16\3\35\13\3\3\3\3\3\7\3!\n\3\f\3\16"+
		"\3$\13\3\3\3\3\3\6\3(\n\3\r\3\16\3)\3\3\7\3-\n\3\f\3\16\3\60\13\3\3\4"+
		"\3\4\3\4\3\4\3\5\5\5\67\n\5\3\6\3\6\3\6\3\6\3\7\5\7>\n\7\3\b\5\bA\n\b"+
		"\3\b\2\2\t\2\4\6\b\n\f\16\2\3\3\2\7\b\2D\2\23\3\2\2\2\4\33\3\2\2\2\6\61"+
		"\3\2\2\2\b\66\3\2\2\2\n8\3\2\2\2\f=\3\2\2\2\16@\3\2\2\2\20\22\5\4\3\2"+
		"\21\20\3\2\2\2\22\25\3\2\2\2\23\21\3\2\2\2\23\24\3\2\2\2\24\26\3\2\2\2"+
		"\25\23\3\2\2\2\26\27\7\2\2\3\27\3\3\2\2\2\30\32\t\2\2\2\31\30\3\2\2\2"+
		"\32\35\3\2\2\2\33\31\3\2\2\2\33\34\3\2\2\2\34\36\3\2\2\2\35\33\3\2\2\2"+
		"\36.\5\6\4\2\37!\t\2\2\2 \37\3\2\2\2!$\3\2\2\2\" \3\2\2\2\"#\3\2\2\2#"+
		"%\3\2\2\2$\"\3\2\2\2%-\5\6\4\2&(\t\2\2\2\'&\3\2\2\2()\3\2\2\2)\'\3\2\2"+
		"\2)*\3\2\2\2*+\3\2\2\2+-\5\n\6\2,\"\3\2\2\2,\'\3\2\2\2-\60\3\2\2\2.,\3"+
		"\2\2\2./\3\2\2\2/\5\3\2\2\2\60.\3\2\2\2\61\62\7\3\2\2\62\63\5\b\5\2\63"+
		"\64\7\4\2\2\64\7\3\2\2\2\65\67\7\6\2\2\66\65\3\2\2\2\66\67\3\2\2\2\67"+
		"\t\3\2\2\289\5\f\7\29:\7\5\2\2:;\5\16\b\2;\13\3\2\2\2<>\7\6\2\2=<\3\2"+
		"\2\2=>\3\2\2\2>\r\3\2\2\2?A\7\6\2\2@?\3\2\2\2@A\3\2\2\2A\17\3\2\2\2\13"+
		"\23\33\"),.\66=@";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}