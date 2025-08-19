// Code generated from antlr/TSgen.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TSgen

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by TSgenParser.
type TSgenVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TSgenParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by TSgenParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by TSgenParser#stateDecl.
	VisitStateDecl(ctx *StateDeclContext) interface{}

	// Visit a parse tree produced by TSgenParser#transitionDecl.
	VisitTransitionDecl(ctx *TransitionDeclContext) interface{}

	// Visit a parse tree produced by TSgenParser#tArgs.
	VisitTArgs(ctx *TArgsContext) interface{}

	// Visit a parse tree produced by TSgenParser#tParam.
	VisitTParam(ctx *TParamContext) interface{}

	// Visit a parse tree produced by TSgenParser#docComment.
	VisitDocComment(ctx *DocCommentContext) interface{}

	// Visit a parse tree produced by TSgenParser#pkgDecl.
	VisitPkgDecl(ctx *PkgDeclContext) interface{}

	// Visit a parse tree produced by TSgenParser#importDecl.
	VisitImportDecl(ctx *ImportDeclContext) interface{}

	// Visit a parse tree produced by TSgenParser#protocolDecl.
	VisitProtocolDecl(ctx *ProtocolDeclContext) interface{}

	// Visit a parse tree produced by TSgenParser#fileDecl.
	VisitFileDecl(ctx *FileDeclContext) interface{}

	// Visit a parse tree produced by TSgenParser#rawDecl.
	VisitRawDecl(ctx *RawDeclContext) interface{}

	// Visit a parse tree produced by TSgenParser#rawInputs.
	VisitRawInputs(ctx *RawInputsContext) interface{}

	// Visit a parse tree produced by TSgenParser#rawParam.
	VisitRawParam(ctx *RawParamContext) interface{}

	// Visit a parse tree produced by TSgenParser#rawOutputs.
	VisitRawOutputs(ctx *RawOutputsContext) interface{}

	// Visit a parse tree produced by TSgenParser#typeExpr.
	VisitTypeExpr(ctx *TypeExprContext) interface{}

	// Visit a parse tree produced by TSgenParser#qualID.
	VisitQualID(ctx *QualIDContext) interface{}

	// Visit a parse tree produced by TSgenParser#s0Decl.
	VisitS0Decl(ctx *S0DeclContext) interface{}

	// Visit a parse tree produced by TSgenParser#fDecl.
	VisitFDecl(ctx *FDeclContext) interface{}
}
