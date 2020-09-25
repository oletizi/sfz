// Generated from src/main/java/org/letizi/sfz/parser/Sfz.g4 by ANTLR 4.8
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
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, Text=17, 
		Newline=18, Whitespace=19, BlockComment=20, LineComment=21;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "Text", 
			"Newline", "Whitespace", "BlockComment", "LineComment"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'<'", "'>'", "'region'", "'group'", "'control'", "'global'", "'curve'", 
			"'effect'", "'master'", "'midi'", "'sampler'", "'='", "'lokey'", "'sample'", 
			"'sw_default'", "'sw_lokey'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, "Text", "Newline", "Whitespace", "BlockComment", 
			"LineComment"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\27\u00bb\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\3\2\3\2\3\3\3\3\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\22"+
		"\6\22\u0093\n\22\r\22\16\22\u0094\3\23\3\23\3\23\5\23\u009a\n\23\3\24"+
		"\6\24\u009d\n\24\r\24\16\24\u009e\3\24\3\24\3\25\3\25\3\25\3\25\7\25\u00a7"+
		"\n\25\f\25\16\25\u00aa\13\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3"+
		"\26\7\26\u00b5\n\26\f\26\16\26\u00b8\13\26\3\26\3\26\3\u00a8\2\27\3\3"+
		"\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21"+
		"!\22#\23%\24\'\25)\26+\27\3\2\5\b\2,,\60;C\\^^aac|\4\2\13\13\"\"\4\2\f"+
		"\f\17\17\2\u00bf\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13"+
		"\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2"+
		"\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2"+
		"!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\3-\3"+
		"\2\2\2\5/\3\2\2\2\7\61\3\2\2\2\t8\3\2\2\2\13>\3\2\2\2\rF\3\2\2\2\17M\3"+
		"\2\2\2\21S\3\2\2\2\23Z\3\2\2\2\25a\3\2\2\2\27f\3\2\2\2\31n\3\2\2\2\33"+
		"p\3\2\2\2\35v\3\2\2\2\37}\3\2\2\2!\u0088\3\2\2\2#\u0092\3\2\2\2%\u0099"+
		"\3\2\2\2\'\u009c\3\2\2\2)\u00a2\3\2\2\2+\u00b0\3\2\2\2-.\7>\2\2.\4\3\2"+
		"\2\2/\60\7@\2\2\60\6\3\2\2\2\61\62\7t\2\2\62\63\7g\2\2\63\64\7i\2\2\64"+
		"\65\7k\2\2\65\66\7q\2\2\66\67\7p\2\2\67\b\3\2\2\289\7i\2\29:\7t\2\2:;"+
		"\7q\2\2;<\7w\2\2<=\7r\2\2=\n\3\2\2\2>?\7e\2\2?@\7q\2\2@A\7p\2\2AB\7v\2"+
		"\2BC\7t\2\2CD\7q\2\2DE\7n\2\2E\f\3\2\2\2FG\7i\2\2GH\7n\2\2HI\7q\2\2IJ"+
		"\7d\2\2JK\7c\2\2KL\7n\2\2L\16\3\2\2\2MN\7e\2\2NO\7w\2\2OP\7t\2\2PQ\7x"+
		"\2\2QR\7g\2\2R\20\3\2\2\2ST\7g\2\2TU\7h\2\2UV\7h\2\2VW\7g\2\2WX\7e\2\2"+
		"XY\7v\2\2Y\22\3\2\2\2Z[\7o\2\2[\\\7c\2\2\\]\7u\2\2]^\7v\2\2^_\7g\2\2_"+
		"`\7t\2\2`\24\3\2\2\2ab\7o\2\2bc\7k\2\2cd\7f\2\2de\7k\2\2e\26\3\2\2\2f"+
		"g\7u\2\2gh\7c\2\2hi\7o\2\2ij\7r\2\2jk\7n\2\2kl\7g\2\2lm\7t\2\2m\30\3\2"+
		"\2\2no\7?\2\2o\32\3\2\2\2pq\7n\2\2qr\7q\2\2rs\7m\2\2st\7g\2\2tu\7{\2\2"+
		"u\34\3\2\2\2vw\7u\2\2wx\7c\2\2xy\7o\2\2yz\7r\2\2z{\7n\2\2{|\7g\2\2|\36"+
		"\3\2\2\2}~\7u\2\2~\177\7y\2\2\177\u0080\7a\2\2\u0080\u0081\7f\2\2\u0081"+
		"\u0082\7g\2\2\u0082\u0083\7h\2\2\u0083\u0084\7c\2\2\u0084\u0085\7w\2\2"+
		"\u0085\u0086\7n\2\2\u0086\u0087\7v\2\2\u0087 \3\2\2\2\u0088\u0089\7u\2"+
		"\2\u0089\u008a\7y\2\2\u008a\u008b\7a\2\2\u008b\u008c\7n\2\2\u008c\u008d"+
		"\7q\2\2\u008d\u008e\7m\2\2\u008e\u008f\7g\2\2\u008f\u0090\7{\2\2\u0090"+
		"\"\3\2\2\2\u0091\u0093\t\2\2\2\u0092\u0091\3\2\2\2\u0093\u0094\3\2\2\2"+
		"\u0094\u0092\3\2\2\2\u0094\u0095\3\2\2\2\u0095$\3\2\2\2\u0096\u0097\7"+
		"\17\2\2\u0097\u009a\7\f\2\2\u0098\u009a\7\f\2\2\u0099\u0096\3\2\2\2\u0099"+
		"\u0098\3\2\2\2\u009a&\3\2\2\2\u009b\u009d\t\3\2\2\u009c\u009b\3\2\2\2"+
		"\u009d\u009e\3\2\2\2\u009e\u009c\3\2\2\2\u009e\u009f\3\2\2\2\u009f\u00a0"+
		"\3\2\2\2\u00a0\u00a1\b\24\2\2\u00a1(\3\2\2\2\u00a2\u00a3\7\61\2\2\u00a3"+
		"\u00a4\7,\2\2\u00a4\u00a8\3\2\2\2\u00a5\u00a7\13\2\2\2\u00a6\u00a5\3\2"+
		"\2\2\u00a7\u00aa\3\2\2\2\u00a8\u00a9\3\2\2\2\u00a8\u00a6\3\2\2\2\u00a9"+
		"\u00ab\3\2\2\2\u00aa\u00a8\3\2\2\2\u00ab\u00ac\7,\2\2\u00ac\u00ad\7\61"+
		"\2\2\u00ad\u00ae\3\2\2\2\u00ae\u00af\b\25\2\2\u00af*\3\2\2\2\u00b0\u00b1"+
		"\7\61\2\2\u00b1\u00b2\7\61\2\2\u00b2\u00b6\3\2\2\2\u00b3\u00b5\n\4\2\2"+
		"\u00b4\u00b3\3\2\2\2\u00b5\u00b8\3\2\2\2\u00b6\u00b4\3\2\2\2\u00b6\u00b7"+
		"\3\2\2\2\u00b7\u00b9\3\2\2\2\u00b8\u00b6\3\2\2\2\u00b9\u00ba\b\26\2\2"+
		"\u00ba,\3\2\2\2\b\2\u0094\u0099\u009e\u00a8\u00b6\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}