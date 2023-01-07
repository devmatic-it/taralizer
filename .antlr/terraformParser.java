// Generated from /Users/bator/Documents/go/taralizer/terraform.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class terraformParser extends Parser {
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
	public static final int
		RULE_file_ = 0, RULE_terraform = 1, RULE_resource = 2, RULE_data = 3, 
		RULE_provider = 4, RULE_output = 5, RULE_local = 6, RULE_module = 7, RULE_variable = 8, 
		RULE_block = 9, RULE_blocktype = 10, RULE_resourcetype = 11, RULE_name = 12, 
		RULE_label = 13, RULE_blockbody = 14, RULE_argument = 15, RULE_identifier = 16, 
		RULE_identifierchain = 17, RULE_inline_index = 18, RULE_expression = 19, 
		RULE_forloop = 20, RULE_section = 21, RULE_val = 22, RULE_functioncall = 23, 
		RULE_functionname = 24, RULE_functionarguments = 25, RULE_index = 26, 
		RULE_filedecl = 27, RULE_list_ = 28, RULE_map_ = 29, RULE_string = 30, 
		RULE_signed_number = 31, RULE_operator_ = 32, RULE_number = 33;
	private static String[] makeRuleNames() {
		return new String[] {
			"file_", "terraform", "resource", "data", "provider", "output", "local", 
			"module", "variable", "block", "blocktype", "resourcetype", "name", "label", 
			"blockbody", "argument", "identifier", "identifierchain", "inline_index", 
			"expression", "forloop", "section", "val", "functioncall", "functionname", 
			"functionarguments", "index", "filedecl", "list_", "map_", "string", 
			"signed_number", "operator_", "number"
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

	@Override
	public String getGrammarFileName() { return "terraform.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public terraformParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class File_Context extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(terraformParser.EOF, 0); }
		public List<LocalContext> local() {
			return getRuleContexts(LocalContext.class);
		}
		public LocalContext local(int i) {
			return getRuleContext(LocalContext.class,i);
		}
		public List<ModuleContext> module() {
			return getRuleContexts(ModuleContext.class);
		}
		public ModuleContext module(int i) {
			return getRuleContext(ModuleContext.class,i);
		}
		public List<OutputContext> output() {
			return getRuleContexts(OutputContext.class);
		}
		public OutputContext output(int i) {
			return getRuleContext(OutputContext.class,i);
		}
		public List<ProviderContext> provider() {
			return getRuleContexts(ProviderContext.class);
		}
		public ProviderContext provider(int i) {
			return getRuleContext(ProviderContext.class,i);
		}
		public List<VariableContext> variable() {
			return getRuleContexts(VariableContext.class);
		}
		public VariableContext variable(int i) {
			return getRuleContext(VariableContext.class,i);
		}
		public List<DataContext> data() {
			return getRuleContexts(DataContext.class);
		}
		public DataContext data(int i) {
			return getRuleContext(DataContext.class,i);
		}
		public List<ResourceContext> resource() {
			return getRuleContexts(ResourceContext.class);
		}
		public ResourceContext resource(int i) {
			return getRuleContext(ResourceContext.class,i);
		}
		public List<TerraformContext> terraform() {
			return getRuleContexts(TerraformContext.class);
		}
		public TerraformContext terraform(int i) {
			return getRuleContext(TerraformContext.class,i);
		}
		public File_Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_file_; }
	}

	public final File_Context file_() throws RecognitionException {
		File_Context _localctx = new File_Context(_ctx, getState());
		enterRule(_localctx, 0, RULE_file_);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(76); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				setState(76);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case T__4:
					{
					setState(68);
					local();
					}
					break;
				case T__5:
					{
					setState(69);
					module();
					}
					break;
				case T__3:
					{
					setState(70);
					output();
					}
					break;
				case PROVIDER:
					{
					setState(71);
					provider();
					}
					break;
				case VARIABLE:
					{
					setState(72);
					variable();
					}
					break;
				case T__2:
					{
					setState(73);
					data();
					}
					break;
				case T__1:
					{
					setState(74);
					resource();
					}
					break;
				case T__0:
					{
					setState(75);
					terraform();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(78); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__3) | (1L << T__4) | (1L << T__5) | (1L << VARIABLE) | (1L << PROVIDER))) != 0) );
			setState(80);
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

	public static class TerraformContext extends ParserRuleContext {
		public BlockbodyContext blockbody() {
			return getRuleContext(BlockbodyContext.class,0);
		}
		public TerraformContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_terraform; }
	}

	public final TerraformContext terraform() throws RecognitionException {
		TerraformContext _localctx = new TerraformContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_terraform);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(82);
			match(T__0);
			setState(83);
			blockbody();
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

	public static class ResourceContext extends ParserRuleContext {
		public ResourcetypeContext resourcetype() {
			return getRuleContext(ResourcetypeContext.class,0);
		}
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public BlockbodyContext blockbody() {
			return getRuleContext(BlockbodyContext.class,0);
		}
		public ResourceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_resource; }
	}

	public final ResourceContext resource() throws RecognitionException {
		ResourceContext _localctx = new ResourceContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_resource);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(85);
			match(T__1);
			setState(86);
			resourcetype();
			setState(87);
			name();
			setState(88);
			blockbody();
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

	public static class DataContext extends ParserRuleContext {
		public ResourcetypeContext resourcetype() {
			return getRuleContext(ResourcetypeContext.class,0);
		}
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public BlockbodyContext blockbody() {
			return getRuleContext(BlockbodyContext.class,0);
		}
		public DataContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_data; }
	}

	public final DataContext data() throws RecognitionException {
		DataContext _localctx = new DataContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_data);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(90);
			match(T__2);
			setState(91);
			resourcetype();
			setState(92);
			name();
			setState(93);
			blockbody();
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

	public static class ProviderContext extends ParserRuleContext {
		public TerminalNode PROVIDER() { return getToken(terraformParser.PROVIDER, 0); }
		public ResourcetypeContext resourcetype() {
			return getRuleContext(ResourcetypeContext.class,0);
		}
		public BlockbodyContext blockbody() {
			return getRuleContext(BlockbodyContext.class,0);
		}
		public ProviderContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_provider; }
	}

	public final ProviderContext provider() throws RecognitionException {
		ProviderContext _localctx = new ProviderContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_provider);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(95);
			match(PROVIDER);
			setState(96);
			resourcetype();
			setState(97);
			blockbody();
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

	public static class OutputContext extends ParserRuleContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public BlockbodyContext blockbody() {
			return getRuleContext(BlockbodyContext.class,0);
		}
		public OutputContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_output; }
	}

	public final OutputContext output() throws RecognitionException {
		OutputContext _localctx = new OutputContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_output);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(99);
			match(T__3);
			setState(100);
			name();
			setState(101);
			blockbody();
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

	public static class LocalContext extends ParserRuleContext {
		public BlockbodyContext blockbody() {
			return getRuleContext(BlockbodyContext.class,0);
		}
		public LocalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_local; }
	}

	public final LocalContext local() throws RecognitionException {
		LocalContext _localctx = new LocalContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_local);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(103);
			match(T__4);
			setState(104);
			blockbody();
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

	public static class ModuleContext extends ParserRuleContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public BlockbodyContext blockbody() {
			return getRuleContext(BlockbodyContext.class,0);
		}
		public ModuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_module; }
	}

	public final ModuleContext module() throws RecognitionException {
		ModuleContext _localctx = new ModuleContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_module);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(106);
			match(T__5);
			setState(107);
			name();
			setState(108);
			blockbody();
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

	public static class VariableContext extends ParserRuleContext {
		public TerminalNode VARIABLE() { return getToken(terraformParser.VARIABLE, 0); }
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public BlockbodyContext blockbody() {
			return getRuleContext(BlockbodyContext.class,0);
		}
		public VariableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variable; }
	}

	public final VariableContext variable() throws RecognitionException {
		VariableContext _localctx = new VariableContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_variable);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(110);
			match(VARIABLE);
			setState(111);
			name();
			setState(112);
			blockbody();
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

	public static class BlockContext extends ParserRuleContext {
		public BlocktypeContext blocktype() {
			return getRuleContext(BlocktypeContext.class,0);
		}
		public BlockbodyContext blockbody() {
			return getRuleContext(BlockbodyContext.class,0);
		}
		public List<LabelContext> label() {
			return getRuleContexts(LabelContext.class);
		}
		public LabelContext label(int i) {
			return getRuleContext(LabelContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
			blocktype();
			setState(118);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==STRING) {
				{
				{
				setState(115);
				label();
				}
				}
				setState(120);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(121);
			blockbody();
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

	public static class BlocktypeContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(terraformParser.IDENTIFIER, 0); }
		public BlocktypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blocktype; }
	}

	public final BlocktypeContext blocktype() throws RecognitionException {
		BlocktypeContext _localctx = new BlocktypeContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_blocktype);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(123);
			match(IDENTIFIER);
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

	public static class ResourcetypeContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(terraformParser.STRING, 0); }
		public ResourcetypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_resourcetype; }
	}

	public final ResourcetypeContext resourcetype() throws RecognitionException {
		ResourcetypeContext _localctx = new ResourcetypeContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_resourcetype);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(125);
			match(STRING);
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

	public static class NameContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(terraformParser.STRING, 0); }
		public NameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_name; }
	}

	public final NameContext name() throws RecognitionException {
		NameContext _localctx = new NameContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(127);
			match(STRING);
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

	public static class LabelContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(terraformParser.STRING, 0); }
		public LabelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_label; }
	}

	public final LabelContext label() throws RecognitionException {
		LabelContext _localctx = new LabelContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_label);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(129);
			match(STRING);
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

	public static class BlockbodyContext extends ParserRuleContext {
		public TerminalNode LCURL() { return getToken(terraformParser.LCURL, 0); }
		public TerminalNode RCURL() { return getToken(terraformParser.RCURL, 0); }
		public List<ArgumentContext> argument() {
			return getRuleContexts(ArgumentContext.class);
		}
		public ArgumentContext argument(int i) {
			return getRuleContext(ArgumentContext.class,i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public BlockbodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blockbody; }
	}

	public final BlockbodyContext blockbody() throws RecognitionException {
		BlockbodyContext _localctx = new BlockbodyContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_blockbody);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(131);
			match(LCURL);
			setState(136);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__2) | (1L << T__5) | (1L << T__7) | (1L << T__8) | (1L << VARIABLE) | (1L << PROVIDER) | (1L << IN) | (1L << STAR) | (1L << NATURAL_NUMBER) | (1L << IDENTIFIER))) != 0)) {
				{
				setState(134);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
				case 1:
					{
					setState(132);
					argument();
					}
					break;
				case 2:
					{
					setState(133);
					block();
					}
					break;
				}
				}
				setState(138);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(139);
			match(RCURL);
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

	public static class ArgumentContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ArgumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argument; }
	}

	public final ArgumentContext argument() throws RecognitionException {
		ArgumentContext _localctx = new ArgumentContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_argument);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(141);
			identifier();
			setState(142);
			match(T__6);
			setState(143);
			expression(0);
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

	public static class IdentifierContext extends ParserRuleContext {
		public IdentifierchainContext identifierchain() {
			return getRuleContext(IdentifierchainContext.class,0);
		}
		public TerminalNode DOT() { return getToken(terraformParser.DOT, 0); }
		public IdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifier; }
	}

	public final IdentifierContext identifier() throws RecognitionException {
		IdentifierContext _localctx = new IdentifierContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_identifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(147);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__2) | (1L << T__5) | (1L << T__7) | (1L << T__8))) != 0)) {
				{
				setState(145);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__2) | (1L << T__5) | (1L << T__7) | (1L << T__8))) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(146);
				match(DOT);
				}
			}

			setState(149);
			identifierchain();
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

	public static class IdentifierchainContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(terraformParser.IDENTIFIER, 0); }
		public TerminalNode IN() { return getToken(terraformParser.IN, 0); }
		public TerminalNode VARIABLE() { return getToken(terraformParser.VARIABLE, 0); }
		public TerminalNode PROVIDER() { return getToken(terraformParser.PROVIDER, 0); }
		public IndexContext index() {
			return getRuleContext(IndexContext.class,0);
		}
		public List<TerminalNode> DOT() { return getTokens(terraformParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(terraformParser.DOT, i);
		}
		public List<IdentifierchainContext> identifierchain() {
			return getRuleContexts(IdentifierchainContext.class);
		}
		public IdentifierchainContext identifierchain(int i) {
			return getRuleContext(IdentifierchainContext.class,i);
		}
		public TerminalNode STAR() { return getToken(terraformParser.STAR, 0); }
		public Inline_indexContext inline_index() {
			return getRuleContext(Inline_indexContext.class,0);
		}
		public IdentifierchainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifierchain; }
	}

	public final IdentifierchainContext identifierchain() throws RecognitionException {
		IdentifierchainContext _localctx = new IdentifierchainContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_identifierchain);
		int _la;
		try {
			int _alt;
			setState(178);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VARIABLE:
			case PROVIDER:
			case IN:
			case IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(151);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << VARIABLE) | (1L << PROVIDER) | (1L << IN) | (1L << IDENTIFIER))) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(153);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
				case 1:
					{
					setState(152);
					index();
					}
					break;
				}
				setState(159);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(155);
						match(DOT);
						setState(156);
						identifierchain();
						}
						} 
					}
					setState(161);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
				}
				}
				break;
			case STAR:
				enterOuterAlt(_localctx, 2);
				{
				setState(162);
				match(STAR);
				setState(167);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(163);
						match(DOT);
						setState(164);
						identifierchain();
						}
						} 
					}
					setState(169);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
				}
				}
				break;
			case NATURAL_NUMBER:
				enterOuterAlt(_localctx, 3);
				{
				setState(170);
				inline_index();
				setState(175);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(171);
						match(DOT);
						setState(172);
						identifierchain();
						}
						} 
					}
					setState(177);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
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

	public static class Inline_indexContext extends ParserRuleContext {
		public TerminalNode NATURAL_NUMBER() { return getToken(terraformParser.NATURAL_NUMBER, 0); }
		public Inline_indexContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_inline_index; }
	}

	public final Inline_indexContext inline_index() throws RecognitionException {
		Inline_indexContext _localctx = new Inline_indexContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_inline_index);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(180);
			match(NATURAL_NUMBER);
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

	public static class ExpressionContext extends ParserRuleContext {
		public SectionContext section() {
			return getRuleContext(SectionContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(terraformParser.LPAREN, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode RPAREN() { return getToken(terraformParser.RPAREN, 0); }
		public ForloopContext forloop() {
			return getRuleContext(ForloopContext.class,0);
		}
		public Operator_Context operator_() {
			return getRuleContext(Operator_Context.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 38;
		enterRecursionRule(_localctx, 38, RULE_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(189);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__2:
			case T__5:
			case T__7:
			case T__8:
			case T__12:
			case T__14:
			case T__16:
			case T__17:
			case T__18:
			case VARIABLE:
			case PROVIDER:
			case IN:
			case STAR:
			case LCURL:
			case EOF_:
			case NULL_:
			case NATURAL_NUMBER:
			case BOOL:
			case DESCRIPTION:
			case MULTILINESTRING:
			case STRING:
			case IDENTIFIER:
				{
				setState(183);
				section();
				}
				break;
			case LPAREN:
				{
				setState(184);
				match(LPAREN);
				setState(185);
				expression(0);
				setState(186);
				match(RPAREN);
				}
				break;
			case T__11:
				{
				setState(188);
				forloop();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(203);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(201);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(191);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(192);
						operator_();
						setState(193);
						expression(5);
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(195);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(196);
						match(T__9);
						setState(197);
						expression(0);
						setState(198);
						match(T__10);
						setState(199);
						expression(3);
						}
						break;
					}
					} 
				}
				setState(205);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ForloopContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode IN() { return getToken(terraformParser.IN, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public ForloopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forloop; }
	}

	public final ForloopContext forloop() throws RecognitionException {
		ForloopContext _localctx = new ForloopContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_forloop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(206);
			match(T__11);
			setState(207);
			identifier();
			setState(208);
			match(IN);
			setState(209);
			expression(0);
			setState(210);
			match(T__10);
			setState(211);
			expression(0);
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

	public static class SectionContext extends ParserRuleContext {
		public List_Context list_() {
			return getRuleContext(List_Context.class,0);
		}
		public Map_Context map_() {
			return getRuleContext(Map_Context.class,0);
		}
		public ValContext val() {
			return getRuleContext(ValContext.class,0);
		}
		public SectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_section; }
	}

	public final SectionContext section() throws RecognitionException {
		SectionContext _localctx = new SectionContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_section);
		try {
			setState(216);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__14:
				enterOuterAlt(_localctx, 1);
				{
				setState(213);
				list_();
				}
				break;
			case LCURL:
				enterOuterAlt(_localctx, 2);
				{
				setState(214);
				map_();
				}
				break;
			case T__2:
			case T__5:
			case T__7:
			case T__8:
			case T__12:
			case T__16:
			case T__17:
			case T__18:
			case VARIABLE:
			case PROVIDER:
			case IN:
			case STAR:
			case EOF_:
			case NULL_:
			case NATURAL_NUMBER:
			case BOOL:
			case DESCRIPTION:
			case MULTILINESTRING:
			case STRING:
			case IDENTIFIER:
				enterOuterAlt(_localctx, 3);
				{
				setState(215);
				val();
				}
				break;
			default:
				throw new NoViableAltException(this);
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

	public static class ValContext extends ParserRuleContext {
		public TerminalNode NULL_() { return getToken(terraformParser.NULL_, 0); }
		public Signed_numberContext signed_number() {
			return getRuleContext(Signed_numberContext.class,0);
		}
		public StringContext string() {
			return getRuleContext(StringContext.class,0);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode BOOL() { return getToken(terraformParser.BOOL, 0); }
		public TerminalNode DESCRIPTION() { return getToken(terraformParser.DESCRIPTION, 0); }
		public FiledeclContext filedecl() {
			return getRuleContext(FiledeclContext.class,0);
		}
		public FunctioncallContext functioncall() {
			return getRuleContext(FunctioncallContext.class,0);
		}
		public TerminalNode EOF_() { return getToken(terraformParser.EOF_, 0); }
		public ValContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_val; }
	}

	public final ValContext val() throws RecognitionException {
		ValContext _localctx = new ValContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_val);
		try {
			setState(227);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(218);
				match(NULL_);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(219);
				signed_number();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(220);
				string();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(221);
				identifier();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(222);
				match(BOOL);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(223);
				match(DESCRIPTION);
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(224);
				filedecl();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(225);
				functioncall();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(226);
				match(EOF_);
				}
				break;
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

	public static class FunctioncallContext extends ParserRuleContext {
		public FunctionnameContext functionname() {
			return getRuleContext(FunctionnameContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(terraformParser.LPAREN, 0); }
		public FunctionargumentsContext functionarguments() {
			return getRuleContext(FunctionargumentsContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(terraformParser.RPAREN, 0); }
		public FunctioncallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functioncall; }
	}

	public final FunctioncallContext functioncall() throws RecognitionException {
		FunctioncallContext _localctx = new FunctioncallContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_functioncall);
		try {
			int _alt;
			setState(243);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(229);
				functionname();
				setState(230);
				match(LPAREN);
				setState(231);
				functionarguments();
				setState(232);
				match(RPAREN);
				}
				break;
			case T__12:
				enterOuterAlt(_localctx, 2);
				{
				setState(234);
				match(T__12);
				setState(235);
				match(LPAREN);
				setState(239);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
				while ( _alt!=1 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1+1 ) {
						{
						{
						setState(236);
						matchWildcard();
						}
						} 
					}
					setState(241);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
				}
				setState(242);
				match(RPAREN);
				}
				break;
			default:
				throw new NoViableAltException(this);
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

	public static class FunctionnameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(terraformParser.IDENTIFIER, 0); }
		public FunctionnameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionname; }
	}

	public final FunctionnameContext functionname() throws RecognitionException {
		FunctionnameContext _localctx = new FunctionnameContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_functionname);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(245);
			match(IDENTIFIER);
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

	public static class FunctionargumentsContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public FunctionargumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionarguments; }
	}

	public final FunctionargumentsContext functionarguments() throws RecognitionException {
		FunctionargumentsContext _localctx = new FunctionargumentsContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_functionarguments);
		int _la;
		try {
			setState(256);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RPAREN:
				enterOuterAlt(_localctx, 1);
				{
				}
				break;
			case T__2:
			case T__5:
			case T__7:
			case T__8:
			case T__11:
			case T__12:
			case T__14:
			case T__16:
			case T__17:
			case T__18:
			case VARIABLE:
			case PROVIDER:
			case IN:
			case STAR:
			case LCURL:
			case LPAREN:
			case EOF_:
			case NULL_:
			case NATURAL_NUMBER:
			case BOOL:
			case DESCRIPTION:
			case MULTILINESTRING:
			case STRING:
			case IDENTIFIER:
				enterOuterAlt(_localctx, 2);
				{
				setState(248);
				expression(0);
				setState(253);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__13) {
					{
					{
					setState(249);
					match(T__13);
					setState(250);
					expression(0);
					}
					}
					setState(255);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
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

	public static class IndexContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public IndexContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_index; }
	}

	public final IndexContext index() throws RecognitionException {
		IndexContext _localctx = new IndexContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_index);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(258);
			match(T__14);
			setState(259);
			expression(0);
			setState(260);
			match(T__15);
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

	public static class FiledeclContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(terraformParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(terraformParser.RPAREN, 0); }
		public FiledeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_filedecl; }
	}

	public final FiledeclContext filedecl() throws RecognitionException {
		FiledeclContext _localctx = new FiledeclContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_filedecl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(262);
			match(T__16);
			setState(263);
			match(LPAREN);
			setState(264);
			expression(0);
			setState(265);
			match(RPAREN);
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

	public static class List_Context extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List_Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_; }
	}

	public final List_Context list_() throws RecognitionException {
		List_Context _localctx = new List_Context(_ctx, getState());
		enterRule(_localctx, 56, RULE_list_);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(267);
			match(T__14);
			setState(279);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__2) | (1L << T__5) | (1L << T__7) | (1L << T__8) | (1L << T__11) | (1L << T__12) | (1L << T__14) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << VARIABLE) | (1L << PROVIDER) | (1L << IN) | (1L << STAR) | (1L << LCURL) | (1L << LPAREN) | (1L << EOF_) | (1L << NULL_) | (1L << NATURAL_NUMBER) | (1L << BOOL) | (1L << DESCRIPTION) | (1L << MULTILINESTRING) | (1L << STRING) | (1L << IDENTIFIER))) != 0)) {
				{
				setState(268);
				expression(0);
				setState(273);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(269);
						match(T__13);
						setState(270);
						expression(0);
						}
						} 
					}
					setState(275);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
				}
				setState(277);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__13) {
					{
					setState(276);
					match(T__13);
					}
				}

				}
			}

			setState(281);
			match(T__15);
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

	public static class Map_Context extends ParserRuleContext {
		public TerminalNode LCURL() { return getToken(terraformParser.LCURL, 0); }
		public TerminalNode RCURL() { return getToken(terraformParser.RCURL, 0); }
		public List<ArgumentContext> argument() {
			return getRuleContexts(ArgumentContext.class);
		}
		public ArgumentContext argument(int i) {
			return getRuleContext(ArgumentContext.class,i);
		}
		public Map_Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_map_; }
	}

	public final Map_Context map_() throws RecognitionException {
		Map_Context _localctx = new Map_Context(_ctx, getState());
		enterRule(_localctx, 58, RULE_map_);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(283);
			match(LCURL);
			setState(290);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__2) | (1L << T__5) | (1L << T__7) | (1L << T__8) | (1L << VARIABLE) | (1L << PROVIDER) | (1L << IN) | (1L << STAR) | (1L << NATURAL_NUMBER) | (1L << IDENTIFIER))) != 0)) {
				{
				{
				setState(284);
				argument();
				setState(286);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__13) {
					{
					setState(285);
					match(T__13);
					}
				}

				}
				}
				setState(292);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(293);
			match(RCURL);
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

	public static class StringContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(terraformParser.STRING, 0); }
		public TerminalNode MULTILINESTRING() { return getToken(terraformParser.MULTILINESTRING, 0); }
		public StringContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_string; }
	}

	public final StringContext string() throws RecognitionException {
		StringContext _localctx = new StringContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_string);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(295);
			_la = _input.LA(1);
			if ( !(_la==MULTILINESTRING || _la==STRING) ) {
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

	public static class Signed_numberContext extends ParserRuleContext {
		public NumberContext number() {
			return getRuleContext(NumberContext.class,0);
		}
		public Signed_numberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_signed_number; }
	}

	public final Signed_numberContext signed_number() throws RecognitionException {
		Signed_numberContext _localctx = new Signed_numberContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_signed_number);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(298);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__17 || _la==T__18) {
				{
				setState(297);
				_la = _input.LA(1);
				if ( !(_la==T__17 || _la==T__18) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
			}

			setState(300);
			number();
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

	public static class Operator_Context extends ParserRuleContext {
		public TerminalNode STAR() { return getToken(terraformParser.STAR, 0); }
		public Operator_Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator_; }
	}

	public final Operator_Context operator_() throws RecognitionException {
		Operator_Context _localctx = new Operator_Context(_ctx, getState());
		enterRule(_localctx, 64, RULE_operator_);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(302);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << STAR))) != 0)) ) {
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

	public static class NumberContext extends ParserRuleContext {
		public List<TerminalNode> NATURAL_NUMBER() { return getTokens(terraformParser.NATURAL_NUMBER); }
		public TerminalNode NATURAL_NUMBER(int i) {
			return getToken(terraformParser.NATURAL_NUMBER, i);
		}
		public TerminalNode DOT() { return getToken(terraformParser.DOT, 0); }
		public NumberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_number; }
	}

	public final NumberContext number() throws RecognitionException {
		NumberContext _localctx = new NumberContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_number);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(304);
			match(NATURAL_NUMBER);
			setState(307);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				{
				setState(305);
				match(DOT);
				setState(306);
				match(NATURAL_NUMBER);
				}
				break;
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 19:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 4);
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\63\u0138\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\6\2O\n\2\r\2\16\2P\3"+
		"\2\3\2\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6"+
		"\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\13"+
		"\3\13\7\13w\n\13\f\13\16\13z\13\13\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16"+
		"\3\17\3\17\3\20\3\20\3\20\7\20\u0089\n\20\f\20\16\20\u008c\13\20\3\20"+
		"\3\20\3\21\3\21\3\21\3\21\3\22\3\22\5\22\u0096\n\22\3\22\3\22\3\23\3\23"+
		"\5\23\u009c\n\23\3\23\3\23\7\23\u00a0\n\23\f\23\16\23\u00a3\13\23\3\23"+
		"\3\23\3\23\7\23\u00a8\n\23\f\23\16\23\u00ab\13\23\3\23\3\23\3\23\7\23"+
		"\u00b0\n\23\f\23\16\23\u00b3\13\23\5\23\u00b5\n\23\3\24\3\24\3\25\3\25"+
		"\3\25\3\25\3\25\3\25\3\25\5\25\u00c0\n\25\3\25\3\25\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\3\25\3\25\7\25\u00cc\n\25\f\25\16\25\u00cf\13\25\3\26\3\26"+
		"\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\5\27\u00db\n\27\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\3\30\3\30\3\30\5\30\u00e6\n\30\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\7\31\u00f0\n\31\f\31\16\31\u00f3\13\31\3\31\5\31\u00f6"+
		"\n\31\3\32\3\32\3\33\3\33\3\33\3\33\7\33\u00fe\n\33\f\33\16\33\u0101\13"+
		"\33\5\33\u0103\n\33\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\36"+
		"\3\36\3\36\3\36\7\36\u0112\n\36\f\36\16\36\u0115\13\36\3\36\5\36\u0118"+
		"\n\36\5\36\u011a\n\36\3\36\3\36\3\37\3\37\3\37\5\37\u0121\n\37\7\37\u0123"+
		"\n\37\f\37\16\37\u0126\13\37\3\37\3\37\3 \3 \3!\5!\u012d\n!\3!\3!\3\""+
		"\3\"\3#\3#\3#\5#\u0136\n#\3#\3\u00f1\3($\2\4\6\b\n\f\16\20\22\24\26\30"+
		"\32\34\36 \"$&(*,.\60\62\64\668:<>@BD\2\7\5\2\5\5\b\b\n\13\4\2 \"\60\60"+
		"\3\2./\3\2\24\25\4\2\24\37##\2\u0140\2N\3\2\2\2\4T\3\2\2\2\6W\3\2\2\2"+
		"\b\\\3\2\2\2\na\3\2\2\2\fe\3\2\2\2\16i\3\2\2\2\20l\3\2\2\2\22p\3\2\2\2"+
		"\24t\3\2\2\2\26}\3\2\2\2\30\177\3\2\2\2\32\u0081\3\2\2\2\34\u0083\3\2"+
		"\2\2\36\u0085\3\2\2\2 \u008f\3\2\2\2\"\u0095\3\2\2\2$\u00b4\3\2\2\2&\u00b6"+
		"\3\2\2\2(\u00bf\3\2\2\2*\u00d0\3\2\2\2,\u00da\3\2\2\2.\u00e5\3\2\2\2\60"+
		"\u00f5\3\2\2\2\62\u00f7\3\2\2\2\64\u0102\3\2\2\2\66\u0104\3\2\2\28\u0108"+
		"\3\2\2\2:\u010d\3\2\2\2<\u011d\3\2\2\2>\u0129\3\2\2\2@\u012c\3\2\2\2B"+
		"\u0130\3\2\2\2D\u0132\3\2\2\2FO\5\16\b\2GO\5\20\t\2HO\5\f\7\2IO\5\n\6"+
		"\2JO\5\22\n\2KO\5\b\5\2LO\5\6\4\2MO\5\4\3\2NF\3\2\2\2NG\3\2\2\2NH\3\2"+
		"\2\2NI\3\2\2\2NJ\3\2\2\2NK\3\2\2\2NL\3\2\2\2NM\3\2\2\2OP\3\2\2\2PN\3\2"+
		"\2\2PQ\3\2\2\2QR\3\2\2\2RS\7\2\2\3S\3\3\2\2\2TU\7\3\2\2UV\5\36\20\2V\5"+
		"\3\2\2\2WX\7\4\2\2XY\5\30\r\2YZ\5\32\16\2Z[\5\36\20\2[\7\3\2\2\2\\]\7"+
		"\5\2\2]^\5\30\r\2^_\5\32\16\2_`\5\36\20\2`\t\3\2\2\2ab\7!\2\2bc\5\30\r"+
		"\2cd\5\36\20\2d\13\3\2\2\2ef\7\6\2\2fg\5\32\16\2gh\5\36\20\2h\r\3\2\2"+
		"\2ij\7\7\2\2jk\5\36\20\2k\17\3\2\2\2lm\7\b\2\2mn\5\32\16\2no\5\36\20\2"+
		"o\21\3\2\2\2pq\7 \2\2qr\5\32\16\2rs\5\36\20\2s\23\3\2\2\2tx\5\26\f\2u"+
		"w\5\34\17\2vu\3\2\2\2wz\3\2\2\2xv\3\2\2\2xy\3\2\2\2y{\3\2\2\2zx\3\2\2"+
		"\2{|\5\36\20\2|\25\3\2\2\2}~\7\60\2\2~\27\3\2\2\2\177\u0080\7/\2\2\u0080"+
		"\31\3\2\2\2\u0081\u0082\7/\2\2\u0082\33\3\2\2\2\u0083\u0084\7/\2\2\u0084"+
		"\35\3\2\2\2\u0085\u008a\7%\2\2\u0086\u0089\5 \21\2\u0087\u0089\5\24\13"+
		"\2\u0088\u0086\3\2\2\2\u0088\u0087\3\2\2\2\u0089\u008c\3\2\2\2\u008a\u0088"+
		"\3\2\2\2\u008a\u008b\3\2\2\2\u008b\u008d\3\2\2\2\u008c\u008a\3\2\2\2\u008d"+
		"\u008e\7&\2\2\u008e\37\3\2\2\2\u008f\u0090\5\"\22\2\u0090\u0091\7\t\2"+
		"\2\u0091\u0092\5(\25\2\u0092!\3\2\2\2\u0093\u0094\t\2\2\2\u0094\u0096"+
		"\7$\2\2\u0095\u0093\3\2\2\2\u0095\u0096\3\2\2\2\u0096\u0097\3\2\2\2\u0097"+
		"\u0098\5$\23\2\u0098#\3\2\2\2\u0099\u009b\t\3\2\2\u009a\u009c\5\66\34"+
		"\2\u009b\u009a\3\2\2\2\u009b\u009c\3\2\2\2\u009c\u00a1\3\2\2\2\u009d\u009e"+
		"\7$\2\2\u009e\u00a0\5$\23\2\u009f\u009d\3\2\2\2\u00a0\u00a3\3\2\2\2\u00a1"+
		"\u009f\3\2\2\2\u00a1\u00a2\3\2\2\2\u00a2\u00b5\3\2\2\2\u00a3\u00a1\3\2"+
		"\2\2\u00a4\u00a9\7#\2\2\u00a5\u00a6\7$\2\2\u00a6\u00a8\5$\23\2\u00a7\u00a5"+
		"\3\2\2\2\u00a8\u00ab\3\2\2\2\u00a9\u00a7\3\2\2\2\u00a9\u00aa\3\2\2\2\u00aa"+
		"\u00b5\3\2\2\2\u00ab\u00a9\3\2\2\2\u00ac\u00b1\5&\24\2\u00ad\u00ae\7$"+
		"\2\2\u00ae\u00b0\5$\23\2\u00af\u00ad\3\2\2\2\u00b0\u00b3\3\2\2\2\u00b1"+
		"\u00af\3\2\2\2\u00b1\u00b2\3\2\2\2\u00b2\u00b5\3\2\2\2\u00b3\u00b1\3\2"+
		"\2\2\u00b4\u0099\3\2\2\2\u00b4\u00a4\3\2\2\2\u00b4\u00ac\3\2\2\2\u00b5"+
		"%\3\2\2\2\u00b6\u00b7\7+\2\2\u00b7\'\3\2\2\2\u00b8\u00b9\b\25\1\2\u00b9"+
		"\u00c0\5,\27\2\u00ba\u00bb\7\'\2\2\u00bb\u00bc\5(\25\2\u00bc\u00bd\7("+
		"\2\2\u00bd\u00c0\3\2\2\2\u00be\u00c0\5*\26\2\u00bf\u00b8\3\2\2\2\u00bf"+
		"\u00ba\3\2\2\2\u00bf\u00be\3\2\2\2\u00c0\u00cd\3\2\2\2\u00c1\u00c2\f\6"+
		"\2\2\u00c2\u00c3\5B\"\2\u00c3\u00c4\5(\25\7\u00c4\u00cc\3\2\2\2\u00c5"+
		"\u00c6\f\4\2\2\u00c6\u00c7\7\f\2\2\u00c7\u00c8\5(\25\2\u00c8\u00c9\7\r"+
		"\2\2\u00c9\u00ca\5(\25\5\u00ca\u00cc\3\2\2\2\u00cb\u00c1\3\2\2\2\u00cb"+
		"\u00c5\3\2\2\2\u00cc\u00cf\3\2\2\2\u00cd\u00cb\3\2\2\2\u00cd\u00ce\3\2"+
		"\2\2\u00ce)\3\2\2\2\u00cf\u00cd\3\2\2\2\u00d0\u00d1\7\16\2\2\u00d1\u00d2"+
		"\5\"\22\2\u00d2\u00d3\7\"\2\2\u00d3\u00d4\5(\25\2\u00d4\u00d5\7\r\2\2"+
		"\u00d5\u00d6\5(\25\2\u00d6+\3\2\2\2\u00d7\u00db\5:\36\2\u00d8\u00db\5"+
		"<\37\2\u00d9\u00db\5.\30\2\u00da\u00d7\3\2\2\2\u00da\u00d8\3\2\2\2\u00da"+
		"\u00d9\3\2\2\2\u00db-\3\2\2\2\u00dc\u00e6\7*\2\2\u00dd\u00e6\5@!\2\u00de"+
		"\u00e6\5> \2\u00df\u00e6\5\"\22\2\u00e0\u00e6\7,\2\2\u00e1\u00e6\7-\2"+
		"\2\u00e2\u00e6\58\35\2\u00e3\u00e6\5\60\31\2\u00e4\u00e6\7)\2\2\u00e5"+
		"\u00dc\3\2\2\2\u00e5\u00dd\3\2\2\2\u00e5\u00de\3\2\2\2\u00e5\u00df\3\2"+
		"\2\2\u00e5\u00e0\3\2\2\2\u00e5\u00e1\3\2\2\2\u00e5\u00e2\3\2\2\2\u00e5"+
		"\u00e3\3\2\2\2\u00e5\u00e4\3\2\2\2\u00e6/\3\2\2\2\u00e7\u00e8\5\62\32"+
		"\2\u00e8\u00e9\7\'\2\2\u00e9\u00ea\5\64\33\2\u00ea\u00eb\7(\2\2\u00eb"+
		"\u00f6\3\2\2\2\u00ec\u00ed\7\17\2\2\u00ed\u00f1\7\'\2\2\u00ee\u00f0\13"+
		"\2\2\2\u00ef\u00ee\3\2\2\2\u00f0\u00f3\3\2\2\2\u00f1\u00f2\3\2\2\2\u00f1"+
		"\u00ef\3\2\2\2\u00f2\u00f4\3\2\2\2\u00f3\u00f1\3\2\2\2\u00f4\u00f6\7("+
		"\2\2\u00f5\u00e7\3\2\2\2\u00f5\u00ec\3\2\2\2\u00f6\61\3\2\2\2\u00f7\u00f8"+
		"\7\60\2\2\u00f8\63\3\2\2\2\u00f9\u0103\3\2\2\2\u00fa\u00ff\5(\25\2\u00fb"+
		"\u00fc\7\20\2\2\u00fc\u00fe\5(\25\2\u00fd\u00fb\3\2\2\2\u00fe\u0101\3"+
		"\2\2\2\u00ff\u00fd\3\2\2\2\u00ff\u0100\3\2\2\2\u0100\u0103\3\2\2\2\u0101"+
		"\u00ff\3\2\2\2\u0102\u00f9\3\2\2\2\u0102\u00fa\3\2\2\2\u0103\65\3\2\2"+
		"\2\u0104\u0105\7\21\2\2\u0105\u0106\5(\25\2\u0106\u0107\7\22\2\2\u0107"+
		"\67\3\2\2\2\u0108\u0109\7\23\2\2\u0109\u010a\7\'\2\2\u010a\u010b\5(\25"+
		"\2\u010b\u010c\7(\2\2\u010c9\3\2\2\2\u010d\u0119\7\21\2\2\u010e\u0113"+
		"\5(\25\2\u010f\u0110\7\20\2\2\u0110\u0112\5(\25\2\u0111\u010f\3\2\2\2"+
		"\u0112\u0115\3\2\2\2\u0113\u0111\3\2\2\2\u0113\u0114\3\2\2\2\u0114\u0117"+
		"\3\2\2\2\u0115\u0113\3\2\2\2\u0116\u0118\7\20\2\2\u0117\u0116\3\2\2\2"+
		"\u0117\u0118\3\2\2\2\u0118\u011a\3\2\2\2\u0119\u010e\3\2\2\2\u0119\u011a"+
		"\3\2\2\2\u011a\u011b\3\2\2\2\u011b\u011c\7\22\2\2\u011c;\3\2\2\2\u011d"+
		"\u0124\7%\2\2\u011e\u0120\5 \21\2\u011f\u0121\7\20\2\2\u0120\u011f\3\2"+
		"\2\2\u0120\u0121\3\2\2\2\u0121\u0123\3\2\2\2\u0122\u011e\3\2\2\2\u0123"+
		"\u0126\3\2\2\2\u0124\u0122\3\2\2\2\u0124\u0125\3\2\2\2\u0125\u0127\3\2"+
		"\2\2\u0126\u0124\3\2\2\2\u0127\u0128\7&\2\2\u0128=\3\2\2\2\u0129\u012a"+
		"\t\4\2\2\u012a?\3\2\2\2\u012b\u012d\t\5\2\2\u012c\u012b\3\2\2\2\u012c"+
		"\u012d\3\2\2\2\u012d\u012e\3\2\2\2\u012e\u012f\5D#\2\u012fA\3\2\2\2\u0130"+
		"\u0131\t\6\2\2\u0131C\3\2\2\2\u0132\u0135\7+\2\2\u0133\u0134\7$\2\2\u0134"+
		"\u0136\7+\2\2\u0135\u0133\3\2\2\2\u0135\u0136\3\2\2\2\u0136E\3\2\2\2\35"+
		"NPx\u0088\u008a\u0095\u009b\u00a1\u00a9\u00b1\u00b4\u00bf\u00cb\u00cd"+
		"\u00da\u00e5\u00f1\u00f5\u00ff\u0102\u0113\u0117\u0119\u0120\u0124\u012c"+
		"\u0135";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}