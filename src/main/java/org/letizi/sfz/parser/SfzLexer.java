// Generated from Sfz.g4 by ANTLR 4.8
package org.letizi.sfz.parser;
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SfzLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, IDENTIFIER=7, GENERATOR=8, 
		INT=9, FLOAT=10, PATH=11, NEWLINE=12, WHITESPACE=13, BLOCK_COMMENT=14, 
		LINE_COMMENT=15, HASH_COMMENT=16;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "IDENTIFIER", "GENERATOR", 
			"INT", "FLOAT", "PATH", "SEPARATOR", "PATH_SEGMENT", "NEWLINE", "WHITESPACE", 
			"BLOCK_COMMENT", "LINE_COMMENT", "HASH_COMMENT"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'<'", "'>'", "'global'", "'group'", "'region'", "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, "IDENTIFIER", "GENERATOR", 
			"INT", "FLOAT", "PATH", "NEWLINE", "WHITESPACE", "BLOCK_COMMENT", "LINE_COMMENT", 
			"HASH_COMMENT"
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


	public SfzLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Sfz.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\22\u009c\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\3\2\3\2\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\b\6\bC\n\b\r\b\16"+
		"\bD\3\t\3\t\6\tI\n\t\r\t\16\tJ\3\n\6\nN\n\n\r\n\16\nO\3\13\3\13\5\13T"+
		"\n\13\3\13\3\13\5\13X\n\13\3\f\7\f[\n\f\f\f\16\f^\13\f\3\f\3\f\6\fb\n"+
		"\f\r\f\16\fc\3\f\3\f\7\fh\n\f\f\f\16\fk\13\f\3\r\3\r\3\16\6\16p\n\16\r"+
		"\16\16\16q\3\17\3\17\3\17\5\17w\n\17\3\20\3\20\3\21\3\21\3\21\3\21\7\21"+
		"\177\n\21\f\21\16\21\u0082\13\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3"+
		"\22\3\22\7\22\u008d\n\22\f\22\16\22\u0090\13\22\3\22\3\22\3\23\3\23\7"+
		"\23\u0096\n\23\f\23\16\23\u0099\13\23\3\23\3\23\3\u0080\2\24\3\3\5\4\7"+
		"\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\2\33\2\35\16\37\17!\20#\21"+
		"%\22\3\2\b\5\2\62;aac|\3\2c|\4\2\61\61^^\7\2/\60\62;C\\aac|\4\2\f\f\17"+
		"\17\4\2\13\13\"\"\2\u00a6\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2"+
		"\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2"+
		"\25\3\2\2\2\2\27\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2"+
		"\2\2%\3\2\2\2\3\'\3\2\2\2\5)\3\2\2\2\7+\3\2\2\2\t\62\3\2\2\2\138\3\2\2"+
		"\2\r?\3\2\2\2\17B\3\2\2\2\21F\3\2\2\2\23M\3\2\2\2\25W\3\2\2\2\27\\\3\2"+
		"\2\2\31l\3\2\2\2\33o\3\2\2\2\35v\3\2\2\2\37x\3\2\2\2!z\3\2\2\2#\u0088"+
		"\3\2\2\2%\u0093\3\2\2\2\'(\7>\2\2(\4\3\2\2\2)*\7@\2\2*\6\3\2\2\2+,\7i"+
		"\2\2,-\7n\2\2-.\7q\2\2./\7d\2\2/\60\7c\2\2\60\61\7n\2\2\61\b\3\2\2\2\62"+
		"\63\7i\2\2\63\64\7t\2\2\64\65\7q\2\2\65\66\7w\2\2\66\67\7r\2\2\67\n\3"+
		"\2\2\289\7t\2\29:\7g\2\2:;\7i\2\2;<\7k\2\2<=\7q\2\2=>\7p\2\2>\f\3\2\2"+
		"\2?@\7?\2\2@\16\3\2\2\2AC\t\2\2\2BA\3\2\2\2CD\3\2\2\2DB\3\2\2\2DE\3\2"+
		"\2\2E\20\3\2\2\2FH\7,\2\2GI\t\3\2\2HG\3\2\2\2IJ\3\2\2\2JH\3\2\2\2JK\3"+
		"\2\2\2K\22\3\2\2\2LN\4\62;\2ML\3\2\2\2NO\3\2\2\2OM\3\2\2\2OP\3\2\2\2P"+
		"\24\3\2\2\2QX\5\23\n\2RT\5\23\n\2SR\3\2\2\2ST\3\2\2\2TU\3\2\2\2UV\7\60"+
		"\2\2VX\5\23\n\2WQ\3\2\2\2WS\3\2\2\2X\26\3\2\2\2Y[\5\31\r\2ZY\3\2\2\2["+
		"^\3\2\2\2\\Z\3\2\2\2\\]\3\2\2\2]_\3\2\2\2^\\\3\2\2\2_i\5\33\16\2`b\5\31"+
		"\r\2a`\3\2\2\2bc\3\2\2\2ca\3\2\2\2cd\3\2\2\2de\3\2\2\2ef\5\33\16\2fh\3"+
		"\2\2\2ga\3\2\2\2hk\3\2\2\2ig\3\2\2\2ij\3\2\2\2j\30\3\2\2\2ki\3\2\2\2l"+
		"m\t\4\2\2m\32\3\2\2\2np\t\5\2\2on\3\2\2\2pq\3\2\2\2qo\3\2\2\2qr\3\2\2"+
		"\2r\34\3\2\2\2st\7\17\2\2tw\7\f\2\2uw\t\6\2\2vs\3\2\2\2vu\3\2\2\2w\36"+
		"\3\2\2\2xy\t\7\2\2y \3\2\2\2z{\7\61\2\2{|\7,\2\2|\u0080\3\2\2\2}\177\13"+
		"\2\2\2~}\3\2\2\2\177\u0082\3\2\2\2\u0080\u0081\3\2\2\2\u0080~\3\2\2\2"+
		"\u0081\u0083\3\2\2\2\u0082\u0080\3\2\2\2\u0083\u0084\7,\2\2\u0084\u0085"+
		"\7\61\2\2\u0085\u0086\3\2\2\2\u0086\u0087\b\21\2\2\u0087\"\3\2\2\2\u0088"+
		"\u0089\7\61\2\2\u0089\u008a\7\61\2\2\u008a\u008e\3\2\2\2\u008b\u008d\n"+
		"\6\2\2\u008c\u008b\3\2\2\2\u008d\u0090\3\2\2\2\u008e\u008c\3\2\2\2\u008e"+
		"\u008f\3\2\2\2\u008f\u0091\3\2\2\2\u0090\u008e\3\2\2\2\u0091\u0092\b\22"+
		"\2\2\u0092$\3\2\2\2\u0093\u0097\7%\2\2\u0094\u0096\n\6\2\2\u0095\u0094"+
		"\3\2\2\2\u0096\u0099\3\2\2\2\u0097\u0095\3\2\2\2\u0097\u0098\3\2\2\2\u0098"+
		"\u009a\3\2\2\2\u0099\u0097\3\2\2\2\u009a\u009b\b\23\2\2\u009b&\3\2\2\2"+
		"\20\2DJOSW\\ciqv\u0080\u008e\u0097\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}