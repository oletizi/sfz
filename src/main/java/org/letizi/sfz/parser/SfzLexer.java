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
		T__0=1, LT=2, GT=3, GLOBAL=4, GROUP=5, REGION=6, BLOCK_COMMENT=7, LINE_COMMENT=8, 
		HASH_COMMENT=9;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "LT", "GT", "GLOBAL", "GROUP", "REGION", "INT", "FLOAT", "TEXT", 
			"NEWLINE", "WHITESPACE", "BLOCK_COMMENT", "LINE_COMMENT", "HASH_COMMENT"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'='", "'<'", "'>'", "'global'", "'group'", "'region'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, "LT", "GT", "GLOBAL", "GROUP", "REGION", "BLOCK_COMMENT", 
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\13t\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7"+
		"\3\7\3\b\6\b;\n\b\r\b\16\b<\3\t\3\t\5\tA\n\t\3\t\3\t\5\tE\n\t\3\n\6\n"+
		"H\n\n\r\n\16\nI\3\13\3\13\3\13\5\13O\n\13\3\f\3\f\3\r\3\r\3\r\3\r\7\r"+
		"W\n\r\f\r\16\rZ\13\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\7\16e\n\16"+
		"\f\16\16\16h\13\16\3\16\3\16\3\17\3\17\7\17n\n\17\f\17\16\17q\13\17\3"+
		"\17\3\17\3X\2\20\3\3\5\4\7\5\t\6\13\7\r\b\17\2\21\2\23\2\25\2\27\2\31"+
		"\t\33\n\35\13\3\2\4\4\2\f\f\17\17\4\2\13\13\"\"\2v\2\3\3\2\2\2\2\5\3\2"+
		"\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\31\3\2\2\2\2\33"+
		"\3\2\2\2\2\35\3\2\2\2\3\37\3\2\2\2\5!\3\2\2\2\7#\3\2\2\2\t%\3\2\2\2\13"+
		",\3\2\2\2\r\62\3\2\2\2\17:\3\2\2\2\21D\3\2\2\2\23G\3\2\2\2\25N\3\2\2\2"+
		"\27P\3\2\2\2\31R\3\2\2\2\33`\3\2\2\2\35k\3\2\2\2\37 \7?\2\2 \4\3\2\2\2"+
		"!\"\7>\2\2\"\6\3\2\2\2#$\7@\2\2$\b\3\2\2\2%&\7i\2\2&\'\7n\2\2\'(\7q\2"+
		"\2()\7d\2\2)*\7c\2\2*+\7n\2\2+\n\3\2\2\2,-\7i\2\2-.\7t\2\2./\7q\2\2/\60"+
		"\7w\2\2\60\61\7r\2\2\61\f\3\2\2\2\62\63\7t\2\2\63\64\7g\2\2\64\65\7i\2"+
		"\2\65\66\7k\2\2\66\67\7q\2\2\678\7p\2\28\16\3\2\2\29;\4\62;\2:9\3\2\2"+
		"\2;<\3\2\2\2<:\3\2\2\2<=\3\2\2\2=\20\3\2\2\2>E\5\17\b\2?A\5\17\b\2@?\3"+
		"\2\2\2@A\3\2\2\2AB\3\2\2\2BC\7\60\2\2CE\5\17\b\2D>\3\2\2\2D@\3\2\2\2E"+
		"\22\3\2\2\2FH\n\2\2\2GF\3\2\2\2HI\3\2\2\2IG\3\2\2\2IJ\3\2\2\2J\24\3\2"+
		"\2\2KL\7\17\2\2LO\7\f\2\2MO\t\2\2\2NK\3\2\2\2NM\3\2\2\2O\26\3\2\2\2PQ"+
		"\t\3\2\2Q\30\3\2\2\2RS\7\61\2\2ST\7,\2\2TX\3\2\2\2UW\13\2\2\2VU\3\2\2"+
		"\2WZ\3\2\2\2XY\3\2\2\2XV\3\2\2\2Y[\3\2\2\2ZX\3\2\2\2[\\\7,\2\2\\]\7\61"+
		"\2\2]^\3\2\2\2^_\b\r\2\2_\32\3\2\2\2`a\7\61\2\2ab\7\61\2\2bf\3\2\2\2c"+
		"e\n\2\2\2dc\3\2\2\2eh\3\2\2\2fd\3\2\2\2fg\3\2\2\2gi\3\2\2\2hf\3\2\2\2"+
		"ij\b\16\2\2j\34\3\2\2\2ko\7%\2\2ln\n\2\2\2ml\3\2\2\2nq\3\2\2\2om\3\2\2"+
		"\2op\3\2\2\2pr\3\2\2\2qo\3\2\2\2rs\b\17\2\2s\36\3\2\2\2\13\2<@DINXfo\3"+
		"\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}