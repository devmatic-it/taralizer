// Generated from /Users/bator/Documents/go/taralizer/pkg/terraform/terraform.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class TerraformLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, VARIABLE=30, PROVIDER=31, 
		IN=32, STAR=33, DOT=34, LCURL=35, RCURL=36, LPAREN=37, RPAREN=38, EOF_=39, 
		NULL_=40, NATURAL_NUMBER=41, BOOL=42, DESCRIPTION=43, MULTILINESTRING=44, 
		STRING=45, IDENTIFIER=46, COMMENT=47, BLOCKCOMMENT=48, WS=49;
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
			"T__25", "T__26", "T__27", "T__28", "DIGIT", "VARIABLE", "PROVIDER", 
			"IN", "STAR", "DOT", "LCURL", "RCURL", "LPAREN", "RPAREN", "EOF_", "NULL_", 
			"NATURAL_NUMBER", "BOOL", "DESCRIPTION", "MULTILINESTRING", "STRING", 
			"IDENTIFIER", "COMMENT", "BLOCKCOMMENT", "WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'terraform'", "'resource'", "'data'", "'output'", "'locals'", 
			"'module'", "'='", "'local'", "'var'", "'?'", "':'", "'for'", "'jsonencode'", 
			"','", "'['", "']'", "'file'", "'+'", "'-'", "'/'", "'%'", "'>'", "'>='", 
			"'<'", "'<='", "'=='", "'!='", "'&&'", "'||'", "'variable'", "'provider'", 
			"'in'", "'*'", "'.'", "'{'", "'}'", "'('", "')'", null, "'nul'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, "VARIABLE", "PROVIDER", "IN", "STAR", 
			"DOT", "LCURL", "RCURL", "LPAREN", "RPAREN", "EOF_", "NULL_", "NATURAL_NUMBER", 
			"BOOL", "DESCRIPTION", "MULTILINESTRING", "STRING", "IDENTIFIER", "COMMENT", 
			"BLOCKCOMMENT", "WS"
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


	public TerraformLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "terraform.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\63\u0188\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\3\2"+
		"\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\n\3\n\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3"+
		"\16\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3"+
		"\22\3\22\3\22\3\22\3\23\3\23\3\24\3\24\3\25\3\25\3\26\3\26\3\27\3\27\3"+
		"\30\3\30\3\30\3\31\3\31\3\32\3\32\3\32\3\33\3\33\3\33\3\34\3\34\3\34\3"+
		"\35\3\35\3\35\3\36\3\36\3\36\3\37\3\37\3 \3 \3 \3 \3 \3 \3 \3 \3 \3!\3"+
		"!\3!\3!\3!\3!\3!\3!\3!\3\"\3\"\3\"\3#\3#\3$\3$\3%\3%\3&\3&\3\'\3\'\3("+
		"\3(\3)\3)\3)\3)\3)\3)\3)\7)\u0107\n)\f)\16)\u010a\13)\3)\3)\3)\3)\3*\3"+
		"*\3*\3*\3+\6+\u0115\n+\r+\16+\u0116\3,\3,\3,\3,\3,\3,\3,\3,\3,\5,\u0122"+
		"\n,\3-\3-\3-\3-\3-\3-\3-\3-\3-\3-\3-\3-\3-\3-\3-\7-\u0133\n-\f-\16-\u0136"+
		"\13-\3-\3-\3-\3-\3-\3-\3-\3-\3-\3-\3-\3-\3.\3.\3.\3.\3.\3.\3.\3.\7.\u014c"+
		"\n.\f.\16.\u014f\13.\3.\3.\3.\3.\3/\3/\3/\3/\7/\u0159\n/\f/\16/\u015c"+
		"\13/\3/\3/\3\60\3\60\7\60\u0162\n\60\f\60\16\60\u0165\13\60\3\61\3\61"+
		"\3\61\5\61\u016a\n\61\3\61\7\61\u016d\n\61\f\61\16\61\u0170\13\61\3\61"+
		"\3\61\3\62\3\62\3\62\3\62\7\62\u0178\n\62\f\62\16\62\u017b\13\62\3\62"+
		"\3\62\3\62\3\62\3\62\3\63\6\63\u0183\n\63\r\63\16\63\u0184\3\63\3\63\6"+
		"\u0108\u0134\u014d\u0179\2\64\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13"+
		"\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61"+
		"\32\63\33\65\34\67\359\36;\37=\2? A!C\"E#G$I%K&M\'O(Q)S*U+W,Y-[.]/_\60"+
		"a\61c\62e\63\3\2\b\3\2\62;\5\2\f\f\17\17$$\4\2C\\c|\7\2//\62;C\\aac|\4"+
		"\2\f\f\17\17\5\2\13\f\17\17\"\"\2\u0192\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3"+
		"\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2"+
		"\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35"+
		"\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)"+
		"\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2"+
		"\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2"+
		"C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3"+
		"\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2"+
		"\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\3g\3\2\2\2\5"+
		"q\3\2\2\2\7z\3\2\2\2\t\177\3\2\2\2\13\u0086\3\2\2\2\r\u008d\3\2\2\2\17"+
		"\u0094\3\2\2\2\21\u0096\3\2\2\2\23\u009c\3\2\2\2\25\u00a0\3\2\2\2\27\u00a2"+
		"\3\2\2\2\31\u00a4\3\2\2\2\33\u00a8\3\2\2\2\35\u00b3\3\2\2\2\37\u00b5\3"+
		"\2\2\2!\u00b7\3\2\2\2#\u00b9\3\2\2\2%\u00be\3\2\2\2\'\u00c0\3\2\2\2)\u00c2"+
		"\3\2\2\2+\u00c4\3\2\2\2-\u00c6\3\2\2\2/\u00c8\3\2\2\2\61\u00cb\3\2\2\2"+
		"\63\u00cd\3\2\2\2\65\u00d0\3\2\2\2\67\u00d3\3\2\2\29\u00d6\3\2\2\2;\u00d9"+
		"\3\2\2\2=\u00dc\3\2\2\2?\u00de\3\2\2\2A\u00e7\3\2\2\2C\u00f0\3\2\2\2E"+
		"\u00f3\3\2\2\2G\u00f5\3\2\2\2I\u00f7\3\2\2\2K\u00f9\3\2\2\2M\u00fb\3\2"+
		"\2\2O\u00fd\3\2\2\2Q\u00ff\3\2\2\2S\u010f\3\2\2\2U\u0114\3\2\2\2W\u0121"+
		"\3\2\2\2Y\u0123\3\2\2\2[\u0143\3\2\2\2]\u0154\3\2\2\2_\u015f\3\2\2\2a"+
		"\u0169\3\2\2\2c\u0173\3\2\2\2e\u0182\3\2\2\2gh\7v\2\2hi\7g\2\2ij\7t\2"+
		"\2jk\7t\2\2kl\7c\2\2lm\7h\2\2mn\7q\2\2no\7t\2\2op\7o\2\2p\4\3\2\2\2qr"+
		"\7t\2\2rs\7g\2\2st\7u\2\2tu\7q\2\2uv\7w\2\2vw\7t\2\2wx\7e\2\2xy\7g\2\2"+
		"y\6\3\2\2\2z{\7f\2\2{|\7c\2\2|}\7v\2\2}~\7c\2\2~\b\3\2\2\2\177\u0080\7"+
		"q\2\2\u0080\u0081\7w\2\2\u0081\u0082\7v\2\2\u0082\u0083\7r\2\2\u0083\u0084"+
		"\7w\2\2\u0084\u0085\7v\2\2\u0085\n\3\2\2\2\u0086\u0087\7n\2\2\u0087\u0088"+
		"\7q\2\2\u0088\u0089\7e\2\2\u0089\u008a\7c\2\2\u008a\u008b\7n\2\2\u008b"+
		"\u008c\7u\2\2\u008c\f\3\2\2\2\u008d\u008e\7o\2\2\u008e\u008f\7q\2\2\u008f"+
		"\u0090\7f\2\2\u0090\u0091\7w\2\2\u0091\u0092\7n\2\2\u0092\u0093\7g\2\2"+
		"\u0093\16\3\2\2\2\u0094\u0095\7?\2\2\u0095\20\3\2\2\2\u0096\u0097\7n\2"+
		"\2\u0097\u0098\7q\2\2\u0098\u0099\7e\2\2\u0099\u009a\7c\2\2\u009a\u009b"+
		"\7n\2\2\u009b\22\3\2\2\2\u009c\u009d\7x\2\2\u009d\u009e\7c\2\2\u009e\u009f"+
		"\7t\2\2\u009f\24\3\2\2\2\u00a0\u00a1\7A\2\2\u00a1\26\3\2\2\2\u00a2\u00a3"+
		"\7<\2\2\u00a3\30\3\2\2\2\u00a4\u00a5\7h\2\2\u00a5\u00a6\7q\2\2\u00a6\u00a7"+
		"\7t\2\2\u00a7\32\3\2\2\2\u00a8\u00a9\7l\2\2\u00a9\u00aa\7u\2\2\u00aa\u00ab"+
		"\7q\2\2\u00ab\u00ac\7p\2\2\u00ac\u00ad\7g\2\2\u00ad\u00ae\7p\2\2\u00ae"+
		"\u00af\7e\2\2\u00af\u00b0\7q\2\2\u00b0\u00b1\7f\2\2\u00b1\u00b2\7g\2\2"+
		"\u00b2\34\3\2\2\2\u00b3\u00b4\7.\2\2\u00b4\36\3\2\2\2\u00b5\u00b6\7]\2"+
		"\2\u00b6 \3\2\2\2\u00b7\u00b8\7_\2\2\u00b8\"\3\2\2\2\u00b9\u00ba\7h\2"+
		"\2\u00ba\u00bb\7k\2\2\u00bb\u00bc\7n\2\2\u00bc\u00bd\7g\2\2\u00bd$\3\2"+
		"\2\2\u00be\u00bf\7-\2\2\u00bf&\3\2\2\2\u00c0\u00c1\7/\2\2\u00c1(\3\2\2"+
		"\2\u00c2\u00c3\7\61\2\2\u00c3*\3\2\2\2\u00c4\u00c5\7\'\2\2\u00c5,\3\2"+
		"\2\2\u00c6\u00c7\7@\2\2\u00c7.\3\2\2\2\u00c8\u00c9\7@\2\2\u00c9\u00ca"+
		"\7?\2\2\u00ca\60\3\2\2\2\u00cb\u00cc\7>\2\2\u00cc\62\3\2\2\2\u00cd\u00ce"+
		"\7>\2\2\u00ce\u00cf\7?\2\2\u00cf\64\3\2\2\2\u00d0\u00d1\7?\2\2\u00d1\u00d2"+
		"\7?\2\2\u00d2\66\3\2\2\2\u00d3\u00d4\7#\2\2\u00d4\u00d5\7?\2\2\u00d58"+
		"\3\2\2\2\u00d6\u00d7\7(\2\2\u00d7\u00d8\7(\2\2\u00d8:\3\2\2\2\u00d9\u00da"+
		"\7~\2\2\u00da\u00db\7~\2\2\u00db<\3\2\2\2\u00dc\u00dd\t\2\2\2\u00dd>\3"+
		"\2\2\2\u00de\u00df\7x\2\2\u00df\u00e0\7c\2\2\u00e0\u00e1\7t\2\2\u00e1"+
		"\u00e2\7k\2\2\u00e2\u00e3\7c\2\2\u00e3\u00e4\7d\2\2\u00e4\u00e5\7n\2\2"+
		"\u00e5\u00e6\7g\2\2\u00e6@\3\2\2\2\u00e7\u00e8\7r\2\2\u00e8\u00e9\7t\2"+
		"\2\u00e9\u00ea\7q\2\2\u00ea\u00eb\7x\2\2\u00eb\u00ec\7k\2\2\u00ec\u00ed"+
		"\7f\2\2\u00ed\u00ee\7g\2\2\u00ee\u00ef\7t\2\2\u00efB\3\2\2\2\u00f0\u00f1"+
		"\7k\2\2\u00f1\u00f2\7p\2\2\u00f2D\3\2\2\2\u00f3\u00f4\7,\2\2\u00f4F\3"+
		"\2\2\2\u00f5\u00f6\7\60\2\2\u00f6H\3\2\2\2\u00f7\u00f8\7}\2\2\u00f8J\3"+
		"\2\2\2\u00f9\u00fa\7\177\2\2\u00faL\3\2\2\2\u00fb\u00fc\7*\2\2\u00fcN"+
		"\3\2\2\2\u00fd\u00fe\7+\2\2\u00feP\3\2\2\2\u00ff\u0100\7>\2\2\u0100\u0101"+
		"\7>\2\2\u0101\u0102\7G\2\2\u0102\u0103\7Q\2\2\u0103\u0104\7H\2\2\u0104"+
		"\u0108\3\2\2\2\u0105\u0107\13\2\2\2\u0106\u0105\3\2\2\2\u0107\u010a\3"+
		"\2\2\2\u0108\u0109\3\2\2\2\u0108\u0106\3\2\2\2\u0109\u010b\3\2\2\2\u010a"+
		"\u0108\3\2\2\2\u010b\u010c\7G\2\2\u010c\u010d\7Q\2\2\u010d\u010e\7H\2"+
		"\2\u010eR\3\2\2\2\u010f\u0110\7p\2\2\u0110\u0111\7w\2\2\u0111\u0112\7"+
		"n\2\2\u0112T\3\2\2\2\u0113\u0115\5=\37\2\u0114\u0113\3\2\2\2\u0115\u0116"+
		"\3\2\2\2\u0116\u0114\3\2\2\2\u0116\u0117\3\2\2\2\u0117V\3\2\2\2\u0118"+
		"\u0119\7v\2\2\u0119\u011a\7t\2\2\u011a\u011b\7w\2\2\u011b\u0122\7g\2\2"+
		"\u011c\u011d\7h\2\2\u011d\u011e\7c\2\2\u011e\u011f\7n\2\2\u011f\u0120"+
		"\7u\2\2\u0120\u0122\7g\2\2\u0121\u0118\3\2\2\2\u0121\u011c\3\2\2\2\u0122"+
		"X\3\2\2\2\u0123\u0124\7>\2\2\u0124\u0125\7>\2\2\u0125\u0126\7F\2\2\u0126"+
		"\u0127\7G\2\2\u0127\u0128\7U\2\2\u0128\u0129\7E\2\2\u0129\u012a\7T\2\2"+
		"\u012a\u012b\7K\2\2\u012b\u012c\7R\2\2\u012c\u012d\7V\2\2\u012d\u012e"+
		"\7K\2\2\u012e\u012f\7Q\2\2\u012f\u0130\7P\2\2\u0130\u0134\3\2\2\2\u0131"+
		"\u0133\13\2\2\2\u0132\u0131\3\2\2\2\u0133\u0136\3\2\2\2\u0134\u0135\3"+
		"\2\2\2\u0134\u0132\3\2\2\2\u0135\u0137\3\2\2\2\u0136\u0134\3\2\2\2\u0137"+
		"\u0138\7F\2\2\u0138\u0139\7G\2\2\u0139\u013a\7U\2\2\u013a\u013b\7E\2\2"+
		"\u013b\u013c\7T\2\2\u013c\u013d\7K\2\2\u013d\u013e\7R\2\2\u013e\u013f"+
		"\7V\2\2\u013f\u0140\7K\2\2\u0140\u0141\7Q\2\2\u0141\u0142\7P\2\2\u0142"+
		"Z\3\2\2\2\u0143\u0144\7>\2\2\u0144\u0145\7>\2\2\u0145\u0146\7/\2\2\u0146"+
		"\u0147\7G\2\2\u0147\u0148\7Q\2\2\u0148\u0149\7H\2\2\u0149\u014d\3\2\2"+
		"\2\u014a\u014c\13\2\2\2\u014b\u014a\3\2\2\2\u014c\u014f\3\2\2\2\u014d"+
		"\u014e\3\2\2\2\u014d\u014b\3\2\2\2\u014e\u0150\3\2\2\2\u014f\u014d\3\2"+
		"\2\2\u0150\u0151\7G\2\2\u0151\u0152\7Q\2\2\u0152\u0153\7H\2\2\u0153\\"+
		"\3\2\2\2\u0154\u015a\7$\2\2\u0155\u0156\7^\2\2\u0156\u0159\7$\2\2\u0157"+
		"\u0159\n\3\2\2\u0158\u0155\3\2\2\2\u0158\u0157\3\2\2\2\u0159\u015c\3\2"+
		"\2\2\u015a\u0158\3\2\2\2\u015a\u015b\3\2\2\2\u015b\u015d\3\2\2\2\u015c"+
		"\u015a\3\2\2\2\u015d\u015e\7$\2\2\u015e^\3\2\2\2\u015f\u0163\t\4\2\2\u0160"+
		"\u0162\t\5\2\2\u0161\u0160\3\2\2\2\u0162\u0165\3\2\2\2\u0163\u0161\3\2"+
		"\2\2\u0163\u0164\3\2\2\2\u0164`\3\2\2\2\u0165\u0163\3\2\2\2\u0166\u016a"+
		"\7%\2\2\u0167\u0168\7\61\2\2\u0168\u016a\7\61\2\2\u0169\u0166\3\2\2\2"+
		"\u0169\u0167\3\2\2\2\u016a\u016e\3\2\2\2\u016b\u016d\n\6\2\2\u016c\u016b"+
		"\3\2\2\2\u016d\u0170\3\2\2\2\u016e\u016c\3\2\2\2\u016e\u016f\3\2\2\2\u016f"+
		"\u0171\3\2\2\2\u0170\u016e\3\2\2\2\u0171\u0172\b\61\2\2\u0172b\3\2\2\2"+
		"\u0173\u0174\7\61\2\2\u0174\u0175\7,\2\2\u0175\u0179\3\2\2\2\u0176\u0178"+
		"\13\2\2\2\u0177\u0176\3\2\2\2\u0178\u017b\3\2\2\2\u0179\u017a\3\2\2\2"+
		"\u0179\u0177\3\2\2\2\u017a\u017c\3\2\2\2\u017b\u0179\3\2\2\2\u017c\u017d"+
		"\7,\2\2\u017d\u017e\7\61\2\2\u017e\u017f\3\2\2\2\u017f\u0180\b\62\2\2"+
		"\u0180d\3\2\2\2\u0181\u0183\t\7\2\2\u0182\u0181\3\2\2\2\u0183\u0184\3"+
		"\2\2\2\u0184\u0182\3\2\2\2\u0184\u0185\3\2\2\2\u0185\u0186\3\2\2\2\u0186"+
		"\u0187\b\63\3\2\u0187f\3\2\2\2\17\2\u0108\u0116\u0121\u0134\u014d\u0158"+
		"\u015a\u0163\u0169\u016e\u0179\u0184\4\2\3\2\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}