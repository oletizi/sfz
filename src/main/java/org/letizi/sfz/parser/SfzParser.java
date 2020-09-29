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
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, Digit=29, Text=30, Newline=31, 
		Whitespace=32, BlockComment=33, LineComment=34;
	public static final int
		RULE_sfz = 0, RULE_line = 1, RULE_header_stmt = 2, RULE_header = 3, RULE_opcode_stmt = 4, 
		RULE_opcode = 5, RULE_value = 6;
	private static String[] makeRuleNames() {
		return new String[] {
			"sfz", "line", "header_stmt", "header", "opcode_stmt", "opcode", "value"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'<'", "'>'", "'region'", "'group'", "'control'", "'global'", "'curve'", 
			"'effect'", "'master'", "'midi'", "'sampler'", "'='", "'ampeg_release'", 
			"'bend_down'", "'bend_up'", "'hikey'", "'hivel'", "'key'", "'lokey'", 
			"'lovel'", "'pitch_keycenter'", "'sample'", "'seq_length'", "'seq_position'", 
			"'sw_default'", "'sw_hikey'", "'sw_last'", "'sw_lokey'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, "Digit", "Text", "Newline", "Whitespace", 
			"BlockComment", "LineComment"
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
		public List<LineContext> line() {
			return getRuleContexts(LineContext.class);
		}
		public LineContext line(int i) {
			return getRuleContext(LineContext.class,i);
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
			setState(15); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(14);
				line();
				}
				}
				setState(17); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27))) != 0) );
			setState(19);
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

	public static class LineContext extends ParserRuleContext {
		public Header_stmtContext header_stmt() {
			return getRuleContext(Header_stmtContext.class,0);
		}
		public Opcode_stmtContext opcode_stmt() {
			return getRuleContext(Opcode_stmtContext.class,0);
		}
		public List<TerminalNode> Newline() { return getTokens(SfzParser.Newline); }
		public TerminalNode Newline(int i) {
			return getToken(SfzParser.Newline, i);
		}
		public LineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_line; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).enterLine(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).exitLine(this);
		}
	}

	public final LineContext line() throws RecognitionException {
		LineContext _localctx = new LineContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_line);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(23);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				{
				setState(21);
				header_stmt();
				}
				break;
			case T__12:
			case T__13:
			case T__14:
			case T__15:
			case T__16:
			case T__17:
			case T__18:
			case T__19:
			case T__20:
			case T__21:
			case T__22:
			case T__23:
			case T__24:
			case T__25:
			case T__26:
			case T__27:
				{
				setState(22);
				opcode_stmt();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(26); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(25);
				match(Newline);
				}
				}
				setState(28); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==Newline );
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

	public static class Header_stmtContext extends ParserRuleContext {
		public HeaderContext header() {
			return getRuleContext(HeaderContext.class,0);
		}
		public Header_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_header_stmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).enterHeader_stmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).exitHeader_stmt(this);
		}
	}

	public final Header_stmtContext header_stmt() throws RecognitionException {
		Header_stmtContext _localctx = new Header_stmtContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_header_stmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(30);
			match(T__0);
			setState(31);
			header();
			setState(32);
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
			setState(34);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__2) | (1L << T__3) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__9) | (1L << T__10))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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

	public static class Opcode_stmtContext extends ParserRuleContext {
		public OpcodeContext opcode() {
			return getRuleContext(OpcodeContext.class,0);
		}
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public List<TerminalNode> Whitespace() { return getTokens(SfzParser.Whitespace); }
		public TerminalNode Whitespace(int i) {
			return getToken(SfzParser.Whitespace, i);
		}
		public Opcode_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_opcode_stmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).enterOpcode_stmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SfzListener ) ((SfzListener)listener).exitOpcode_stmt(this);
		}
	}

	public final Opcode_stmtContext opcode_stmt() throws RecognitionException {
		Opcode_stmtContext _localctx = new Opcode_stmtContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_opcode_stmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(36);
			opcode();
			setState(38);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Whitespace) {
				{
				setState(37);
				match(Whitespace);
				}
			}

			setState(40);
			match(T__11);
			setState(42);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Whitespace) {
				{
				setState(41);
				match(Whitespace);
				}
			}

			setState(44);
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
			setState(46);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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
		public List<TerminalNode> Text() { return getTokens(SfzParser.Text); }
		public TerminalNode Text(int i) {
			return getToken(SfzParser.Text, i);
		}
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
			setState(49); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(48);
				match(Text);
				}
				}
				setState(51); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==Text );
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3$8\4\2\t\2\4\3\t\3"+
		"\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\3\2\6\2\22\n\2\r\2\16\2\23\3"+
		"\2\3\2\3\3\3\3\5\3\32\n\3\3\3\6\3\35\n\3\r\3\16\3\36\3\4\3\4\3\4\3\4\3"+
		"\5\3\5\3\6\3\6\5\6)\n\6\3\6\3\6\5\6-\n\6\3\6\3\6\3\7\3\7\3\b\6\b\64\n"+
		"\b\r\b\16\b\65\3\b\2\2\t\2\4\6\b\n\f\16\2\4\3\2\5\r\3\2\17\36\2\66\2\21"+
		"\3\2\2\2\4\31\3\2\2\2\6 \3\2\2\2\b$\3\2\2\2\n&\3\2\2\2\f\60\3\2\2\2\16"+
		"\63\3\2\2\2\20\22\5\4\3\2\21\20\3\2\2\2\22\23\3\2\2\2\23\21\3\2\2\2\23"+
		"\24\3\2\2\2\24\25\3\2\2\2\25\26\7\2\2\3\26\3\3\2\2\2\27\32\5\6\4\2\30"+
		"\32\5\n\6\2\31\27\3\2\2\2\31\30\3\2\2\2\32\34\3\2\2\2\33\35\7!\2\2\34"+
		"\33\3\2\2\2\35\36\3\2\2\2\36\34\3\2\2\2\36\37\3\2\2\2\37\5\3\2\2\2 !\7"+
		"\3\2\2!\"\5\b\5\2\"#\7\4\2\2#\7\3\2\2\2$%\t\2\2\2%\t\3\2\2\2&(\5\f\7\2"+
		"\')\7\"\2\2(\'\3\2\2\2()\3\2\2\2)*\3\2\2\2*,\7\16\2\2+-\7\"\2\2,+\3\2"+
		"\2\2,-\3\2\2\2-.\3\2\2\2./\5\16\b\2/\13\3\2\2\2\60\61\t\3\2\2\61\r\3\2"+
		"\2\2\62\64\7 \2\2\63\62\3\2\2\2\64\65\3\2\2\2\65\63\3\2\2\2\65\66\3\2"+
		"\2\2\66\17\3\2\2\2\b\23\31\36(,\65";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}