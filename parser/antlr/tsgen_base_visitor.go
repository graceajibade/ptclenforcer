// Code generated from antlr/TSgen.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TSgen

import "github.com/antlr4-go/antlr/v4"

type BaseTSgenVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTSgenVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitStateDecl(ctx *StateDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitTransitionDecl(ctx *TransitionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitTArgs(ctx *TArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitTParam(ctx *TParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitDocComment(ctx *DocCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitPkgDecl(ctx *PkgDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitImportDecl(ctx *ImportDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitProtocolDecl(ctx *ProtocolDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitFileDecl(ctx *FileDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitRawDecl(ctx *RawDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitRawInputs(ctx *RawInputsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitRawParam(ctx *RawParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitRawOutputs(ctx *RawOutputsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitTypeExpr(ctx *TypeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitQualID(ctx *QualIDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitS0Decl(ctx *S0DeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSgenVisitor) VisitFDecl(ctx *FDeclContext) interface{} {
	return v.VisitChildren(ctx)
}
