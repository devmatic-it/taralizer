// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package terraform // Terraform
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseTerraformListener is a complete listener for a parse tree produced by TerraformParser.
type BaseTerraformListener struct{}

var _ TerraformListener = &BaseTerraformListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTerraformListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTerraformListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTerraformListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTerraformListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile_ is called when production file_ is entered.
func (s *BaseTerraformListener) EnterFile_(ctx *File_Context) {}

// ExitFile_ is called when production file_ is exited.
func (s *BaseTerraformListener) ExitFile_(ctx *File_Context) {}

// EnterTerraform is called when production terraform is entered.
func (s *BaseTerraformListener) EnterTerraform(ctx *TerraformContext) {}

// ExitTerraform is called when production terraform is exited.
func (s *BaseTerraformListener) ExitTerraform(ctx *TerraformContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseTerraformListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseTerraformListener) ExitResource(ctx *ResourceContext) {}

// EnterData is called when production data is entered.
func (s *BaseTerraformListener) EnterData(ctx *DataContext) {}

// ExitData is called when production data is exited.
func (s *BaseTerraformListener) ExitData(ctx *DataContext) {}

// EnterProvider is called when production provider is entered.
func (s *BaseTerraformListener) EnterProvider(ctx *ProviderContext) {}

// ExitProvider is called when production provider is exited.
func (s *BaseTerraformListener) ExitProvider(ctx *ProviderContext) {}

// EnterOutput is called when production output is entered.
func (s *BaseTerraformListener) EnterOutput(ctx *OutputContext) {}

// ExitOutput is called when production output is exited.
func (s *BaseTerraformListener) ExitOutput(ctx *OutputContext) {}

// EnterLocal is called when production local is entered.
func (s *BaseTerraformListener) EnterLocal(ctx *LocalContext) {}

// ExitLocal is called when production local is exited.
func (s *BaseTerraformListener) ExitLocal(ctx *LocalContext) {}

// EnterModule is called when production module is entered.
func (s *BaseTerraformListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseTerraformListener) ExitModule(ctx *ModuleContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseTerraformListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseTerraformListener) ExitVariable(ctx *VariableContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseTerraformListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseTerraformListener) ExitBlock(ctx *BlockContext) {}

// EnterBlocktype is called when production blocktype is entered.
func (s *BaseTerraformListener) EnterBlocktype(ctx *BlocktypeContext) {}

// ExitBlocktype is called when production blocktype is exited.
func (s *BaseTerraformListener) ExitBlocktype(ctx *BlocktypeContext) {}

// EnterResourcetype is called when production resourcetype is entered.
func (s *BaseTerraformListener) EnterResourcetype(ctx *ResourcetypeContext) {}

// ExitResourcetype is called when production resourcetype is exited.
func (s *BaseTerraformListener) ExitResourcetype(ctx *ResourcetypeContext) {}

// EnterName is called when production name is entered.
func (s *BaseTerraformListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseTerraformListener) ExitName(ctx *NameContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseTerraformListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseTerraformListener) ExitLabel(ctx *LabelContext) {}

// EnterBlockbody is called when production blockbody is entered.
func (s *BaseTerraformListener) EnterBlockbody(ctx *BlockbodyContext) {}

// ExitBlockbody is called when production blockbody is exited.
func (s *BaseTerraformListener) ExitBlockbody(ctx *BlockbodyContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseTerraformListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseTerraformListener) ExitArgument(ctx *ArgumentContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseTerraformListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseTerraformListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIdentifierchain is called when production identifierchain is entered.
func (s *BaseTerraformListener) EnterIdentifierchain(ctx *IdentifierchainContext) {}

// ExitIdentifierchain is called when production identifierchain is exited.
func (s *BaseTerraformListener) ExitIdentifierchain(ctx *IdentifierchainContext) {}

// EnterInline_index is called when production inline_index is entered.
func (s *BaseTerraformListener) EnterInline_index(ctx *Inline_indexContext) {}

// ExitInline_index is called when production inline_index is exited.
func (s *BaseTerraformListener) ExitInline_index(ctx *Inline_indexContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTerraformListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTerraformListener) ExitExpression(ctx *ExpressionContext) {}

// EnterForloop is called when production forloop is entered.
func (s *BaseTerraformListener) EnterForloop(ctx *ForloopContext) {}

// ExitForloop is called when production forloop is exited.
func (s *BaseTerraformListener) ExitForloop(ctx *ForloopContext) {}

// EnterSection is called when production section is entered.
func (s *BaseTerraformListener) EnterSection(ctx *SectionContext) {}

// ExitSection is called when production section is exited.
func (s *BaseTerraformListener) ExitSection(ctx *SectionContext) {}

// EnterVal is called when production val is entered.
func (s *BaseTerraformListener) EnterVal(ctx *ValContext) {}

// ExitVal is called when production val is exited.
func (s *BaseTerraformListener) ExitVal(ctx *ValContext) {}

// EnterFunctioncall is called when production functioncall is entered.
func (s *BaseTerraformListener) EnterFunctioncall(ctx *FunctioncallContext) {}

// ExitFunctioncall is called when production functioncall is exited.
func (s *BaseTerraformListener) ExitFunctioncall(ctx *FunctioncallContext) {}

// EnterFunctionname is called when production functionname is entered.
func (s *BaseTerraformListener) EnterFunctionname(ctx *FunctionnameContext) {}

// ExitFunctionname is called when production functionname is exited.
func (s *BaseTerraformListener) ExitFunctionname(ctx *FunctionnameContext) {}

// EnterFunctionarguments is called when production functionarguments is entered.
func (s *BaseTerraformListener) EnterFunctionarguments(ctx *FunctionargumentsContext) {}

// ExitFunctionarguments is called when production functionarguments is exited.
func (s *BaseTerraformListener) ExitFunctionarguments(ctx *FunctionargumentsContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseTerraformListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseTerraformListener) ExitIndex(ctx *IndexContext) {}

// EnterFiledecl is called when production filedecl is entered.
func (s *BaseTerraformListener) EnterFiledecl(ctx *FiledeclContext) {}

// ExitFiledecl is called when production filedecl is exited.
func (s *BaseTerraformListener) ExitFiledecl(ctx *FiledeclContext) {}

// EnterList_ is called when production list_ is entered.
func (s *BaseTerraformListener) EnterList_(ctx *List_Context) {}

// ExitList_ is called when production list_ is exited.
func (s *BaseTerraformListener) ExitList_(ctx *List_Context) {}

// EnterMap_ is called when production map_ is entered.
func (s *BaseTerraformListener) EnterMap_(ctx *Map_Context) {}

// ExitMap_ is called when production map_ is exited.
func (s *BaseTerraformListener) ExitMap_(ctx *Map_Context) {}

// EnterString is called when production string is entered.
func (s *BaseTerraformListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseTerraformListener) ExitString(ctx *StringContext) {}

// EnterSigned_number is called when production signed_number is entered.
func (s *BaseTerraformListener) EnterSigned_number(ctx *Signed_numberContext) {}

// ExitSigned_number is called when production signed_number is exited.
func (s *BaseTerraformListener) ExitSigned_number(ctx *Signed_numberContext) {}

// EnterOperator_ is called when production operator_ is entered.
func (s *BaseTerraformListener) EnterOperator_(ctx *Operator_Context) {}

// ExitOperator_ is called when production operator_ is exited.
func (s *BaseTerraformListener) ExitOperator_(ctx *Operator_Context) {}

// EnterNumber is called when production number is entered.
func (s *BaseTerraformListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseTerraformListener) ExitNumber(ctx *NumberContext) {}
