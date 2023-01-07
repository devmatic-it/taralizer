// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package terraform // Terraform
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// TerraformListener is a complete listener for a parse tree produced by TerraformParser.
type TerraformListener interface {
	antlr.ParseTreeListener

	// EnterFile_ is called when entering the file_ production.
	EnterFile_(c *File_Context)

	// EnterTerraform is called when entering the terraform production.
	EnterTerraform(c *TerraformContext)

	// EnterResource is called when entering the resource production.
	EnterResource(c *ResourceContext)

	// EnterData is called when entering the data production.
	EnterData(c *DataContext)

	// EnterProvider is called when entering the provider production.
	EnterProvider(c *ProviderContext)

	// EnterOutput is called when entering the output production.
	EnterOutput(c *OutputContext)

	// EnterLocal is called when entering the local production.
	EnterLocal(c *LocalContext)

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlocktype is called when entering the blocktype production.
	EnterBlocktype(c *BlocktypeContext)

	// EnterResourcetype is called when entering the resourcetype production.
	EnterResourcetype(c *ResourcetypeContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterBlockbody is called when entering the blockbody production.
	EnterBlockbody(c *BlockbodyContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterIdentifierchain is called when entering the identifierchain production.
	EnterIdentifierchain(c *IdentifierchainContext)

	// EnterInline_index is called when entering the inline_index production.
	EnterInline_index(c *Inline_indexContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterForloop is called when entering the forloop production.
	EnterForloop(c *ForloopContext)

	// EnterSection is called when entering the section production.
	EnterSection(c *SectionContext)

	// EnterVal is called when entering the val production.
	EnterVal(c *ValContext)

	// EnterFunctioncall is called when entering the functioncall production.
	EnterFunctioncall(c *FunctioncallContext)

	// EnterFunctionname is called when entering the functionname production.
	EnterFunctionname(c *FunctionnameContext)

	// EnterFunctionarguments is called when entering the functionarguments production.
	EnterFunctionarguments(c *FunctionargumentsContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterFiledecl is called when entering the filedecl production.
	EnterFiledecl(c *FiledeclContext)

	// EnterList_ is called when entering the list_ production.
	EnterList_(c *List_Context)

	// EnterMap_ is called when entering the map_ production.
	EnterMap_(c *Map_Context)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterSigned_number is called when entering the signed_number production.
	EnterSigned_number(c *Signed_numberContext)

	// EnterOperator_ is called when entering the operator_ production.
	EnterOperator_(c *Operator_Context)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitFile_ is called when exiting the file_ production.
	ExitFile_(c *File_Context)

	// ExitTerraform is called when exiting the terraform production.
	ExitTerraform(c *TerraformContext)

	// ExitResource is called when exiting the resource production.
	ExitResource(c *ResourceContext)

	// ExitData is called when exiting the data production.
	ExitData(c *DataContext)

	// ExitProvider is called when exiting the provider production.
	ExitProvider(c *ProviderContext)

	// ExitOutput is called when exiting the output production.
	ExitOutput(c *OutputContext)

	// ExitLocal is called when exiting the local production.
	ExitLocal(c *LocalContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlocktype is called when exiting the blocktype production.
	ExitBlocktype(c *BlocktypeContext)

	// ExitResourcetype is called when exiting the resourcetype production.
	ExitResourcetype(c *ResourcetypeContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitBlockbody is called when exiting the blockbody production.
	ExitBlockbody(c *BlockbodyContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitIdentifierchain is called when exiting the identifierchain production.
	ExitIdentifierchain(c *IdentifierchainContext)

	// ExitInline_index is called when exiting the inline_index production.
	ExitInline_index(c *Inline_indexContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitForloop is called when exiting the forloop production.
	ExitForloop(c *ForloopContext)

	// ExitSection is called when exiting the section production.
	ExitSection(c *SectionContext)

	// ExitVal is called when exiting the val production.
	ExitVal(c *ValContext)

	// ExitFunctioncall is called when exiting the functioncall production.
	ExitFunctioncall(c *FunctioncallContext)

	// ExitFunctionname is called when exiting the functionname production.
	ExitFunctionname(c *FunctionnameContext)

	// ExitFunctionarguments is called when exiting the functionarguments production.
	ExitFunctionarguments(c *FunctionargumentsContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitFiledecl is called when exiting the filedecl production.
	ExitFiledecl(c *FiledeclContext)

	// ExitList_ is called when exiting the list_ production.
	ExitList_(c *List_Context)

	// ExitMap_ is called when exiting the map_ production.
	ExitMap_(c *Map_Context)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitSigned_number is called when exiting the signed_number production.
	ExitSigned_number(c *Signed_numberContext)

	// ExitOperator_ is called when exiting the operator_ production.
	ExitOperator_(c *Operator_Context)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
