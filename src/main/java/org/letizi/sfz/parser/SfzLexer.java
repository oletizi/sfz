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
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, Digit=28, Text=29, Newline=30, Whitespace=31, 
		BlockComment=32, LineComment=33;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
			"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24", 
			"T__25", "T__26", "Digit", "Text", "Newline", "Whitespace", "BlockComment", 
			"LineComment"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'<'", "'>'", "'region'", "'group'", "'control'", "'global'", "'curve'", 
			"'effect'", "'master'", "'midi'", "'sampler'", "'='", "'ampeg_release'", 
			"'bend_down'", "'bend_up'", "'hikey'", "'hivel'", "'lokey'", "'lovel'", 
			"'pitch_keycenter'", "'sample'", "'seq_length'", "'seq_position'", "'sw_default'", 
			"'sw_hikey'", "'sw_last'", "'sw_lokey'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, "Digit", "Text", "Newline", "Whitespace", "BlockComment", 
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2#\u0140\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\3\2\3\2\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\35\3\35\3\36\6\36\u0118\n\36\r\36\16\36\u0119\3\37\3\37\3\37\5"+
		"\37\u011f\n\37\3 \6 \u0122\n \r \16 \u0123\3 \3 \3!\3!\3!\3!\7!\u012c"+
		"\n!\f!\16!\u012f\13!\3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\7\"\u013a\n\"\f\""+
		"\16\"\u013d\13\"\3\"\3\"\3\u012d\2#\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n"+
		"\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30"+
		"/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#\3\2\6\3\2\62;\b\2,,/;C\\"+
		"^^aac|\4\2\13\13\"\"\4\2\f\f\17\17\2\u0144\2\3\3\2\2\2\2\5\3\2\2\2\2\7"+
		"\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2"+
		"\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2"+
		"\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2"+
		"\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2"+
		"\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2"+
		"\2A\3\2\2\2\2C\3\2\2\2\3E\3\2\2\2\5G\3\2\2\2\7I\3\2\2\2\tP\3\2\2\2\13"+
		"V\3\2\2\2\r^\3\2\2\2\17e\3\2\2\2\21k\3\2\2\2\23r\3\2\2\2\25y\3\2\2\2\27"+
		"~\3\2\2\2\31\u0086\3\2\2\2\33\u0088\3\2\2\2\35\u0096\3\2\2\2\37\u00a0"+
		"\3\2\2\2!\u00a8\3\2\2\2#\u00ae\3\2\2\2%\u00b4\3\2\2\2\'\u00ba\3\2\2\2"+
		")\u00c0\3\2\2\2+\u00d0\3\2\2\2-\u00d7\3\2\2\2/\u00e2\3\2\2\2\61\u00ef"+
		"\3\2\2\2\63\u00fa\3\2\2\2\65\u0103\3\2\2\2\67\u010b\3\2\2\29\u0114\3\2"+
		"\2\2;\u0117\3\2\2\2=\u011e\3\2\2\2?\u0121\3\2\2\2A\u0127\3\2\2\2C\u0135"+
		"\3\2\2\2EF\7>\2\2F\4\3\2\2\2GH\7@\2\2H\6\3\2\2\2IJ\7t\2\2JK\7g\2\2KL\7"+
		"i\2\2LM\7k\2\2MN\7q\2\2NO\7p\2\2O\b\3\2\2\2PQ\7i\2\2QR\7t\2\2RS\7q\2\2"+
		"ST\7w\2\2TU\7r\2\2U\n\3\2\2\2VW\7e\2\2WX\7q\2\2XY\7p\2\2YZ\7v\2\2Z[\7"+
		"t\2\2[\\\7q\2\2\\]\7n\2\2]\f\3\2\2\2^_\7i\2\2_`\7n\2\2`a\7q\2\2ab\7d\2"+
		"\2bc\7c\2\2cd\7n\2\2d\16\3\2\2\2ef\7e\2\2fg\7w\2\2gh\7t\2\2hi\7x\2\2i"+
		"j\7g\2\2j\20\3\2\2\2kl\7g\2\2lm\7h\2\2mn\7h\2\2no\7g\2\2op\7e\2\2pq\7"+
		"v\2\2q\22\3\2\2\2rs\7o\2\2st\7c\2\2tu\7u\2\2uv\7v\2\2vw\7g\2\2wx\7t\2"+
		"\2x\24\3\2\2\2yz\7o\2\2z{\7k\2\2{|\7f\2\2|}\7k\2\2}\26\3\2\2\2~\177\7"+
		"u\2\2\177\u0080\7c\2\2\u0080\u0081\7o\2\2\u0081\u0082\7r\2\2\u0082\u0083"+
		"\7n\2\2\u0083\u0084\7g\2\2\u0084\u0085\7t\2\2\u0085\30\3\2\2\2\u0086\u0087"+
		"\7?\2\2\u0087\32\3\2\2\2\u0088\u0089\7c\2\2\u0089\u008a\7o\2\2\u008a\u008b"+
		"\7r\2\2\u008b\u008c\7g\2\2\u008c\u008d\7i\2\2\u008d\u008e\7a\2\2\u008e"+
		"\u008f\7t\2\2\u008f\u0090\7g\2\2\u0090\u0091\7n\2\2\u0091\u0092\7g\2\2"+
		"\u0092\u0093\7c\2\2\u0093\u0094\7u\2\2\u0094\u0095\7g\2\2\u0095\34\3\2"+
		"\2\2\u0096\u0097\7d\2\2\u0097\u0098\7g\2\2\u0098\u0099\7p\2\2\u0099\u009a"+
		"\7f\2\2\u009a\u009b\7a\2\2\u009b\u009c\7f\2\2\u009c\u009d\7q\2\2\u009d"+
		"\u009e\7y\2\2\u009e\u009f\7p\2\2\u009f\36\3\2\2\2\u00a0\u00a1\7d\2\2\u00a1"+
		"\u00a2\7g\2\2\u00a2\u00a3\7p\2\2\u00a3\u00a4\7f\2\2\u00a4\u00a5\7a\2\2"+
		"\u00a5\u00a6\7w\2\2\u00a6\u00a7\7r\2\2\u00a7 \3\2\2\2\u00a8\u00a9\7j\2"+
		"\2\u00a9\u00aa\7k\2\2\u00aa\u00ab\7m\2\2\u00ab\u00ac\7g\2\2\u00ac\u00ad"+
		"\7{\2\2\u00ad\"\3\2\2\2\u00ae\u00af\7j\2\2\u00af\u00b0\7k\2\2\u00b0\u00b1"+
		"\7x\2\2\u00b1\u00b2\7g\2\2\u00b2\u00b3\7n\2\2\u00b3$\3\2\2\2\u00b4\u00b5"+
		"\7n\2\2\u00b5\u00b6\7q\2\2\u00b6\u00b7\7m\2\2\u00b7\u00b8\7g\2\2\u00b8"+
		"\u00b9\7{\2\2\u00b9&\3\2\2\2\u00ba\u00bb\7n\2\2\u00bb\u00bc\7q\2\2\u00bc"+
		"\u00bd\7x\2\2\u00bd\u00be\7g\2\2\u00be\u00bf\7n\2\2\u00bf(\3\2\2\2\u00c0"+
		"\u00c1\7r\2\2\u00c1\u00c2\7k\2\2\u00c2\u00c3\7v\2\2\u00c3\u00c4\7e\2\2"+
		"\u00c4\u00c5\7j\2\2\u00c5\u00c6\7a\2\2\u00c6\u00c7\7m\2\2\u00c7\u00c8"+
		"\7g\2\2\u00c8\u00c9\7{\2\2\u00c9\u00ca\7e\2\2\u00ca\u00cb\7g\2\2\u00cb"+
		"\u00cc\7p\2\2\u00cc\u00cd\7v\2\2\u00cd\u00ce\7g\2\2\u00ce\u00cf\7t\2\2"+
		"\u00cf*\3\2\2\2\u00d0\u00d1\7u\2\2\u00d1\u00d2\7c\2\2\u00d2\u00d3\7o\2"+
		"\2\u00d3\u00d4\7r\2\2\u00d4\u00d5\7n\2\2\u00d5\u00d6\7g\2\2\u00d6,\3\2"+
		"\2\2\u00d7\u00d8\7u\2\2\u00d8\u00d9\7g\2\2\u00d9\u00da\7s\2\2\u00da\u00db"+
		"\7a\2\2\u00db\u00dc\7n\2\2\u00dc\u00dd\7g\2\2\u00dd\u00de\7p\2\2\u00de"+
		"\u00df\7i\2\2\u00df\u00e0\7v\2\2\u00e0\u00e1\7j\2\2\u00e1.\3\2\2\2\u00e2"+
		"\u00e3\7u\2\2\u00e3\u00e4\7g\2\2\u00e4\u00e5\7s\2\2\u00e5\u00e6\7a\2\2"+
		"\u00e6\u00e7\7r\2\2\u00e7\u00e8\7q\2\2\u00e8\u00e9\7u\2\2\u00e9\u00ea"+
		"\7k\2\2\u00ea\u00eb\7v\2\2\u00eb\u00ec\7k\2\2\u00ec\u00ed\7q\2\2\u00ed"+
		"\u00ee\7p\2\2\u00ee\60\3\2\2\2\u00ef\u00f0\7u\2\2\u00f0\u00f1\7y\2\2\u00f1"+
		"\u00f2\7a\2\2\u00f2\u00f3\7f\2\2\u00f3\u00f4\7g\2\2\u00f4\u00f5\7h\2\2"+
		"\u00f5\u00f6\7c\2\2\u00f6\u00f7\7w\2\2\u00f7\u00f8\7n\2\2\u00f8\u00f9"+
		"\7v\2\2\u00f9\62\3\2\2\2\u00fa\u00fb\7u\2\2\u00fb\u00fc\7y\2\2\u00fc\u00fd"+
		"\7a\2\2\u00fd\u00fe\7j\2\2\u00fe\u00ff\7k\2\2\u00ff\u0100\7m\2\2\u0100"+
		"\u0101\7g\2\2\u0101\u0102\7{\2\2\u0102\64\3\2\2\2\u0103\u0104\7u\2\2\u0104"+
		"\u0105\7y\2\2\u0105\u0106\7a\2\2\u0106\u0107\7n\2\2\u0107\u0108\7c\2\2"+
		"\u0108\u0109\7u\2\2\u0109\u010a\7v\2\2\u010a\66\3\2\2\2\u010b\u010c\7"+
		"u\2\2\u010c\u010d\7y\2\2\u010d\u010e\7a\2\2\u010e\u010f\7n\2\2\u010f\u0110"+
		"\7q\2\2\u0110\u0111\7m\2\2\u0111\u0112\7g\2\2\u0112\u0113\7{\2\2\u0113"+
		"8\3\2\2\2\u0114\u0115\t\2\2\2\u0115:\3\2\2\2\u0116\u0118\t\3\2\2\u0117"+
		"\u0116\3\2\2\2\u0118\u0119\3\2\2\2\u0119\u0117\3\2\2\2\u0119\u011a\3\2"+
		"\2\2\u011a<\3\2\2\2\u011b\u011c\7\17\2\2\u011c\u011f\7\f\2\2\u011d\u011f"+
		"\7\f\2\2\u011e\u011b\3\2\2\2\u011e\u011d\3\2\2\2\u011f>\3\2\2\2\u0120"+
		"\u0122\t\4\2\2\u0121\u0120\3\2\2\2\u0122\u0123\3\2\2\2\u0123\u0121\3\2"+
		"\2\2\u0123\u0124\3\2\2\2\u0124\u0125\3\2\2\2\u0125\u0126\b \2\2\u0126"+
		"@\3\2\2\2\u0127\u0128\7\61\2\2\u0128\u0129\7,\2\2\u0129\u012d\3\2\2\2"+
		"\u012a\u012c\13\2\2\2\u012b\u012a\3\2\2\2\u012c\u012f\3\2\2\2\u012d\u012e"+
		"\3\2\2\2\u012d\u012b\3\2\2\2\u012e\u0130\3\2\2\2\u012f\u012d\3\2\2\2\u0130"+
		"\u0131\7,\2\2\u0131\u0132\7\61\2\2\u0132\u0133\3\2\2\2\u0133\u0134\b!"+
		"\2\2\u0134B\3\2\2\2\u0135\u0136\7\61\2\2\u0136\u0137\7\61\2\2\u0137\u013b"+
		"\3\2\2\2\u0138\u013a\n\5\2\2\u0139\u0138\3\2\2\2\u013a\u013d\3\2\2\2\u013b"+
		"\u0139\3\2\2\2\u013b\u013c\3\2\2\2\u013c\u013e\3\2\2\2\u013d\u013b\3\2"+
		"\2\2\u013e\u013f\b\"\2\2\u013fD\3\2\2\2\b\2\u0119\u011e\u0123\u012d\u013b"+
		"\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}