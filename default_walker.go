package johnnie

import "go/ast"

type DefaultWalker struct {
}

func (walker *DefaultWalker) BeginWalk(node ast.Node) (w Walker) {
	if node == nil {
		w = nil
	} else {
		w = walker
	}
	return
}

func (*DefaultWalker) EndWalk(node ast.Node) {

}

func (*DefaultWalker) WalkComment(comment *ast.Comment) bool {
	return true
}

func (*DefaultWalker) WalkCommentGroup(group *ast.CommentGroup) bool {
	return true
}

func (*DefaultWalker) EndWalkComment(comment *ast.Comment) {

}

func (*DefaultWalker) EndWalkCommentGroup(group *ast.CommentGroup) {

}

func (*DefaultWalker) WalkField(field *ast.Field) bool {
	return true
}

func (*DefaultWalker) EndWalkField(field *ast.Field) {

}

func (*DefaultWalker) WalkFieldList(list *ast.FieldList) bool {
	return true
}

func (*DefaultWalker) EndWalkFieldList(list *ast.FieldList) {

}

func (*DefaultWalker) WalkEllipsis(ellipsis *ast.Ellipsis) bool {
	return true
}

func (*DefaultWalker) EndWalkEllipsis(ellipsis *ast.Ellipsis) {

}

func (*DefaultWalker) WalkFuncLit(lit *ast.FuncLit) bool {
	return true
}

func (*DefaultWalker) EndWalkFuncLit(lit *ast.FuncLit) {

}

func (*DefaultWalker) WalkCompositeLit(lit *ast.CompositeLit) bool {
	return true
}

func (*DefaultWalker) EndWalkCompositeLit(lit *ast.CompositeLit) {

}

func (*DefaultWalker) WalkParenExpr(expr *ast.ParenExpr) bool {
	return true
}

func (*DefaultWalker) EndWalkParenExpr(expr *ast.ParenExpr) {

}

func (*DefaultWalker) WalkSelectorExpr(expr *ast.SelectorExpr) bool {
	return true
}

func (*DefaultWalker) EndWalkSelectorExpr(expr *ast.SelectorExpr) {

}

func (*DefaultWalker) WalkIndexExpr(expr *ast.IndexExpr) bool {
	return true
}

func (*DefaultWalker) EndWalkIndexExpr(expr *ast.IndexExpr) {

}

func (*DefaultWalker) WalkSliceExpr(expr *ast.SliceExpr) bool {
	return true
}

func (*DefaultWalker) EndWalkSliceExpr(expr *ast.SliceExpr) {

}

func (*DefaultWalker) WalkTypeAssertExpr(expr *ast.TypeAssertExpr) bool {
	return true
}

func (*DefaultWalker) EndWalkTypeAssertExpr(expr *ast.TypeAssertExpr) {

}

func (*DefaultWalker) WalkCallExpr(expr *ast.CallExpr) bool {
	return true
}

func (*DefaultWalker) EndWalkCallExpr(expr *ast.CallExpr) {

}

func (*DefaultWalker) WalkStarExpr(expr *ast.StarExpr) bool {
	return true
}

func (*DefaultWalker) EndWalkStarExpr(expr *ast.StarExpr) {

}

func (*DefaultWalker) WalkUnaryExpr(expr *ast.UnaryExpr) bool {
	return true
}

func (*DefaultWalker) EndWalkUnaryExpr(expr *ast.UnaryExpr) {

}

func (*DefaultWalker) WalkBinaryExpr(expr *ast.BinaryExpr) bool {
	return true
}

func (*DefaultWalker) EndWalkBinaryExpr(expr *ast.BinaryExpr) {

}

func (*DefaultWalker) WalkKeyValueExpr(expr *ast.KeyValueExpr) bool {
	return true
}

func (*DefaultWalker) EndWalkKeyValueExpr(expr *ast.KeyValueExpr) {

}

func (*DefaultWalker) WalkArrayType(arrayType *ast.ArrayType) bool {
	return true
}

func (*DefaultWalker) EndWalkArrayType(arrayType *ast.ArrayType) {

}

func (*DefaultWalker) WalkStructType(structType *ast.StructType) bool {
	return true
}

func (*DefaultWalker) EndWalkStructType(structType *ast.StructType) {

}

func (*DefaultWalker) WalkFuncType(funcType *ast.FuncType) bool {
	return true
}

func (*DefaultWalker) EndWalkFuncType(funcType *ast.FuncType) {

}

func (*DefaultWalker) WalkInterfaceType(interfaceType *ast.InterfaceType) bool {
	return true
}

func (*DefaultWalker) EndWalkInterfaceType(interfaceType *ast.InterfaceType) {

}

func (*DefaultWalker) WalkMapType(mapType *ast.MapType) bool {
	return true
}

func (*DefaultWalker) EndWalkMapType(mapType *ast.MapType) {

}

func (*DefaultWalker) WalkChanType(chanType *ast.ChanType) bool {
	return true
}

func (*DefaultWalker) EndWalkChanType(chanType *ast.ChanType) {

}

func (*DefaultWalker) WalkDeclStmt(stmt *ast.DeclStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkDeclStmt(stmt *ast.DeclStmt) {

}

func (*DefaultWalker) WalkLabeledStmt(stmt *ast.LabeledStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkLabeledStmt(stmt *ast.LabeledStmt) {

}

func (*DefaultWalker) WalkExprStmt(stmt *ast.ExprStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkExprStmt(stmt *ast.ExprStmt) {

}

func (*DefaultWalker) WalkSendStmt(stmt *ast.SendStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkSendStmt(stmt *ast.SendStmt) {

}

func (*DefaultWalker) WalkIncDecStmt(stmt *ast.IncDecStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkIncDecStmt(stmt *ast.IncDecStmt) {

}

func (*DefaultWalker) WalkAssignStmt(stmt *ast.AssignStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkAssignStmt(stmt *ast.AssignStmt) {

}

func (*DefaultWalker) WalkGoStmt(stmt *ast.GoStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkGoStmt(stmt *ast.GoStmt) {

}

func (*DefaultWalker) WalkDeferStmt(stmt *ast.DeferStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkDeferStmt(stmt *ast.DeferStmt) {

}

func (*DefaultWalker) WalkReturnStmt(stmt *ast.ReturnStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkReturnStmt(stmt *ast.ReturnStmt) {

}

func (*DefaultWalker) WalkBranchStmt(stmt *ast.BranchStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkBranchStmt(stmt *ast.BranchStmt) {

}

func (*DefaultWalker) WalkBlockStmt(stmt *ast.BlockStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkBlockStmt(stmt *ast.BlockStmt) {

}

func (*DefaultWalker) WalkIfStmt(stmt *ast.IfStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkIfStmt(stmt *ast.IfStmt) {

}

func (*DefaultWalker) WalkCaseClause(clause *ast.CaseClause) bool {
	return true
}

func (*DefaultWalker) EndWalkCaseClause(clause *ast.CaseClause) {

}

func (*DefaultWalker) WalkSwitchStmt(stmt *ast.SwitchStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkSwitchStmt(stmt *ast.SwitchStmt) {

}

func (*DefaultWalker) WalkTypeSwitchStmt(stmt *ast.TypeSwitchStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkTypeSwitchStmt(stmt *ast.TypeSwitchStmt) {

}

func (*DefaultWalker) WalkCommClause(clause *ast.CommClause) bool {
	return true
}

func (*DefaultWalker) EndWalkCommClause(clause *ast.CommClause) {

}

func (*DefaultWalker) WalkSelectStmt(stmt *ast.SelectStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkSelectStmt(stmt *ast.SelectStmt) {

}

func (*DefaultWalker) WalkForStmt(stmt *ast.ForStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkForStmt(stmt *ast.ForStmt) {

}

func (*DefaultWalker) WalkRangeStmt(stmt *ast.RangeStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkRangeStmt(stmt *ast.RangeStmt) {

}

func (*DefaultWalker) WalkImportSpec(spec *ast.ImportSpec) bool {
	return true
}

func (*DefaultWalker) EndWalkImportSpec(spec *ast.ImportSpec) {

}

func (*DefaultWalker) WalkValueSpec(spec *ast.ValueSpec) bool {
	return true
}

func (*DefaultWalker) EndWalkValueSpec(spec *ast.ValueSpec) {

}

func (*DefaultWalker) WalkTypeSpec(spec *ast.TypeSpec) bool {
	return true
}

func (*DefaultWalker) EndWalkTypeSpec(spec *ast.TypeSpec) {

}

func (*DefaultWalker) WalkGenDecl(decl *ast.GenDecl) bool {
	return true
}

func (*DefaultWalker) EndWalkGenDecl(decl *ast.GenDecl) {

}

func (*DefaultWalker) WalkFuncDecl(decl *ast.FuncDecl) bool {
	return true
}

func (*DefaultWalker) EndWalkFuncDecl(decl *ast.FuncDecl) {

}

func (*DefaultWalker) WalkFile(file *ast.File) bool {
	return true
}

func (*DefaultWalker) EndWalkFile(file *ast.File) {

}

func (*DefaultWalker) WalkPackage(n *ast.Package) bool {
	return true
}

func (*DefaultWalker) EndWalkPackage(n *ast.Package) {

}

func (*DefaultWalker) WalkBadExpr(n *ast.BadExpr) bool {
	return true
}

func (*DefaultWalker) EndWalkBadExpr(n *ast.BadExpr) {

}

func (*DefaultWalker) WalkIdent(n *ast.Ident) bool {
	return true
}

func (*DefaultWalker) EndWalkIdent(n *ast.Ident) {

}

func (*DefaultWalker) WalkBasicLit(n *ast.BasicLit) bool {
	return true
}

func (*DefaultWalker) EndWalkBasicLit(n *ast.BasicLit) {

}

func (*DefaultWalker) WalkBadStmt(n *ast.BadStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkBadStmt(n *ast.BadStmt) {

}

func (*DefaultWalker) WalkEmptyStmt(n *ast.EmptyStmt) bool {
	return true
}

func (*DefaultWalker) EndWalkEmptyStmt(n *ast.EmptyStmt) {

}

func (*DefaultWalker) WalkBadDecl(n *ast.BadDecl) bool {
	return true
}

func (*DefaultWalker) EndWalkBadDecl(n *ast.BadDecl) {

}
