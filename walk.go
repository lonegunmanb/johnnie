// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package johnnie

import (
	"fmt"
	"go/ast"
	"reflect"
)

// A Walker's BeginWalk method is invoked for each node encountered by Visit.
// If the result visitor w is not nil, Visit visits each of the children
// of node with the visitor w, followed by a call of w.BeginWalk(nil).
//go:generate mockgen -package=johnnie -destination=./mock_walker_test.go github.com/lonegunmanb/johnnie Walker
type Walker interface {
	BeginWalk(node ast.Node) (w Walker)
	EndWalk(node ast.Node)
	WalkComment(comment *ast.Comment) bool
	WalkCommentGroup(group *ast.CommentGroup) bool
	EndWalkComment(comment *ast.Comment)
	EndWalkCommentGroup(group *ast.CommentGroup)
	WalkField(field *ast.Field) bool
	EndWalkField(field *ast.Field)
	WalkFieldList(list *ast.FieldList) bool
	EndWalkFieldList(list *ast.FieldList)
	WalkEllipsis(ellipsis *ast.Ellipsis) bool
	EndWalkEllipsis(ellipsis *ast.Ellipsis)
	WalkFuncLit(lit *ast.FuncLit) bool
	EndWalkFuncLit(lit *ast.FuncLit)
	WalkCompositeLit(lit *ast.CompositeLit) bool
	EndWalkCompositeLit(lit *ast.CompositeLit)
	WalkParenExpr(expr *ast.ParenExpr) bool
	EndWalkParenExpr(expr *ast.ParenExpr)
	WalkSelectorExpr(expr *ast.SelectorExpr) bool
	EndWalkSelectorExpr(expr *ast.SelectorExpr)
	WalkIndexExpr(expr *ast.IndexExpr) bool
	EndWalkIndexExpr(expr *ast.IndexExpr)
	WalkSliceExpr(expr *ast.SliceExpr) bool
	EndWalkSliceExpr(expr *ast.SliceExpr)
	WalkTypeAssertExpr(expr *ast.TypeAssertExpr) bool
	EndWalkTypeAssertExpr(expr *ast.TypeAssertExpr)
	WalkCallExpr(expr *ast.CallExpr) bool
	EndWalkCallExpr(expr *ast.CallExpr)
	WalkStarExpr(expr *ast.StarExpr) bool
	EndWalkStarExpr(expr *ast.StarExpr)
	WalkUnaryExpr(expr *ast.UnaryExpr) bool
	EndWalkUnaryExpr(expr *ast.UnaryExpr)
	WalkBinaryExpr(expr *ast.BinaryExpr) bool
	EndWalkBinaryExpr(expr *ast.BinaryExpr)
	WalkKeyValueExpr(expr *ast.KeyValueExpr) bool
	EndWalkKeyValueExpr(expr *ast.KeyValueExpr)
	WalkArrayType(arrayType *ast.ArrayType) bool
	EndWalkArrayType(arrayType *ast.ArrayType)
	WalkStructType(structType *ast.StructType) bool
	EndWalkStructType(structType *ast.StructType)
	WalkFuncType(funcType *ast.FuncType) bool
	EndWalkFuncType(funcType *ast.FuncType)
	WalkInterfaceType(interfaceType *ast.InterfaceType) bool
	EndWalkInterfaceType(interfaceType *ast.InterfaceType)
	WalkMapType(mapType *ast.MapType) bool
	EndWalkMapType(mapType *ast.MapType)
	WalkChanType(chanType *ast.ChanType) bool
	EndWalkChanType(chanType *ast.ChanType)
	WalkDeclStmt(stmt *ast.DeclStmt) bool
	EndWalkDeclStmt(stmt *ast.DeclStmt)
	WalkLabeledStmt(stmt *ast.LabeledStmt) bool
	EndWalkLabeledStmt(stmt *ast.LabeledStmt)
	WalkExprStmt(stmt *ast.ExprStmt) bool
	EndWalkExprStmt(stmt *ast.ExprStmt)
	WalkSendStmt(stmt *ast.SendStmt) bool
	EndWalkSendStmt(stmt *ast.SendStmt)
	WalkIncDecStmt(stmt *ast.IncDecStmt) bool
	EndWalkIncDecStmt(stmt *ast.IncDecStmt)
	WalkAssignStmt(stmt *ast.AssignStmt) bool
	EndWalkAssignStmt(stmt *ast.AssignStmt)
	WalkGoStmt(stmt *ast.GoStmt) bool
	EndWalkGoStmt(stmt *ast.GoStmt)
	WalkDeferStmt(stmt *ast.DeferStmt) bool
	EndWalkDeferStmt(stmt *ast.DeferStmt)
	WalkReturnStmt(stmt *ast.ReturnStmt) bool
	EndWalkReturnStmt(stmt *ast.ReturnStmt)
	WalkBranchStmt(stmt *ast.BranchStmt) bool
	EndWalkBranchStmt(stmt *ast.BranchStmt)
	WalkBlockStmt(stmt *ast.BlockStmt) bool
	EndWalkBlockStmt(stmt *ast.BlockStmt)
	WalkIfStmt(stmt *ast.IfStmt) bool
	EndWalkIfStmt(stmt *ast.IfStmt)
	WalkCaseClause(clause *ast.CaseClause) bool
	EndWalkCaseClause(clause *ast.CaseClause)
	WalkSwitchStmt(stmt *ast.SwitchStmt) bool
	EndWalkSwitchStmt(stmt *ast.SwitchStmt)
	WalkTypeSwitchStmt(stmt *ast.TypeSwitchStmt) bool
	EndWalkTypeSwitchStmt(stmt *ast.TypeSwitchStmt)
	WalkCommClause(clause *ast.CommClause) bool
	EndWalkCommClause(clause *ast.CommClause)
	WalkSelectStmt(stmt *ast.SelectStmt) bool
	EndWalkSelectStmt(stmt *ast.SelectStmt)
	WalkForStmt(stmt *ast.ForStmt) bool
	EndWalkForStmt(stmt *ast.ForStmt)
	WalkRangeStmt(stmt *ast.RangeStmt) bool
	EndWalkRangeStmt(stmt *ast.RangeStmt)
	WalkImportSpec(spec *ast.ImportSpec) bool
	EndWalkImportSpec(spec *ast.ImportSpec)
	WalkValueSpec(spec *ast.ValueSpec) bool
	EndWalkValueSpec(spec *ast.ValueSpec)
	WalkTypeSpec(spec *ast.TypeSpec) bool
	EndWalkTypeSpec(spec *ast.TypeSpec)
	WalkGenDecl(decl *ast.GenDecl) bool
	EndWalkGenDecl(decl *ast.GenDecl)
	WalkFuncDecl(decl *ast.FuncDecl) bool
	EndWalkFuncDecl(decl *ast.FuncDecl)
	WalkFile(file *ast.File) bool
	EndWalkFile(file *ast.File)
	WalkPackage(n *ast.Package) bool
	EndWalkPackage(n *ast.Package)
	WalkBadExpr(n *ast.BadExpr) bool
	EndWalkBadExpr(n *ast.BadExpr)
	WalkIdent(n *ast.Ident) bool
	EndWalkIdent(n *ast.Ident)
	WalkBasicLit(n *ast.BasicLit) bool
	EndWalkBasicLit(n *ast.BasicLit)
	WalkBadStmt(n *ast.BadStmt) bool
	EndWalkBadStmt(n *ast.BadStmt)
	WalkEmptyStmt(n *ast.EmptyStmt) bool
	EndWalkEmptyStmt(n *ast.EmptyStmt)
	WalkBadDecl(n *ast.BadDecl) bool
	EndWalkBadDecl(n *ast.BadDecl)
}

// Helper functions for common node lists. They may be empty.

func walkIdentList(v Walker, list []*ast.Ident) {
	for _, x := range list {
		Visit(v, x)
	}
}

func walkExprList(v Walker, list []ast.Expr) {
	for _, x := range list {
		Visit(v, x)
	}
}

func walkStmtList(v Walker, list []ast.Stmt) {
	for _, x := range list {
		Visit(v, x)
	}
}

func walkDeclList(v Walker, list []ast.Decl) {
	for _, x := range list {
		Visit(v, x)
	}
}

// TODO(gri): Investigate if providing a closure to Visit leads to
//            simpler use (and may help eliminate Inspect in turn).

// Visit traverses an AST in depth-first order: It starts by calling
// v.BeginWalk(node); node must not be nil. If the visitor w returned by
// v.BeginWalk(node) is not nil, Visit is invoked recursively with visitor
// w for each of the non-nil children of node, followed by a call of
// w.BeginWalk(nil).
//
func Visit(v Walker, node ast.Node) {
	if v.BeginWalk(node) == nil {
		return
	}
	if !walkNode(v, node) {
		return
	}

	// walk children
	// (the order of the cases matches the order
	// of the corresponding node types in ast.go)
	switch n := node.(type) {
	// Comments and fields
	case *ast.Comment:
		// nothing to do

	case *ast.CommentGroup:
		for _, c := range n.List {
			Visit(v, c)
		}

	case *ast.Field:
		if n.Doc != nil {
			Visit(v, n.Doc)
		}
		walkIdentList(v, n.Names)
		Visit(v, n.Type)
		if n.Tag != nil {
			Visit(v, n.Tag)
		}
		if n.Comment != nil {
			Visit(v, n.Comment)
		}

	case *ast.FieldList:
		for _, f := range n.List {
			Visit(v, f)
		}

	// Expressions
	case *ast.BadExpr, *ast.Ident, *ast.BasicLit:
		// nothing to do

	case *ast.Ellipsis:
		if n.Elt != nil {
			Visit(v, n.Elt)
		}

	case *ast.FuncLit:
		Visit(v, n.Type)
		Visit(v, n.Body)

	case *ast.CompositeLit:
		if n.Type != nil {
			Visit(v, n.Type)
		}
		walkExprList(v, n.Elts)

	case *ast.ParenExpr:
		Visit(v, n.X)

	case *ast.SelectorExpr:
		Visit(v, n.X)
		Visit(v, n.Sel)

	case *ast.IndexExpr:
		Visit(v, n.X)
		Visit(v, n.Index)

	case *ast.SliceExpr:
		Visit(v, n.X)
		if n.Low != nil {
			Visit(v, n.Low)
		}
		if n.High != nil {
			Visit(v, n.High)
		}
		if n.Max != nil {
			Visit(v, n.Max)
		}

	case *ast.TypeAssertExpr:
		Visit(v, n.X)
		if n.Type != nil {
			Visit(v, n.Type)
		}

	case *ast.CallExpr:
		Visit(v, n.Fun)
		walkExprList(v, n.Args)

	case *ast.StarExpr:
		Visit(v, n.X)

	case *ast.UnaryExpr:
		Visit(v, n.X)

	case *ast.BinaryExpr:
		Visit(v, n.X)
		Visit(v, n.Y)

	case *ast.KeyValueExpr:
		Visit(v, n.Key)
		Visit(v, n.Value)

	// Types
	case *ast.ArrayType:
		if n.Len != nil {
			Visit(v, n.Len)
		}
		Visit(v, n.Elt)

	case *ast.StructType:
		Visit(v, n.Fields)

	case *ast.FuncType:
		if n.Params != nil {
			Visit(v, n.Params)
		}
		if n.Results != nil {
			Visit(v, n.Results)
		}

	case *ast.InterfaceType:
		Visit(v, n.Methods)

	case *ast.MapType:
		Visit(v, n.Key)
		Visit(v, n.Value)

	case *ast.ChanType:
		Visit(v, n.Value)

	// Statements
	case *ast.BadStmt:
		// nothing to do

	case *ast.DeclStmt:
		Visit(v, n.Decl)

	case *ast.EmptyStmt:
		// nothing to do

	case *ast.LabeledStmt:
		Visit(v, n.Label)
		Visit(v, n.Stmt)

	case *ast.ExprStmt:
		Visit(v, n.X)

	case *ast.SendStmt:
		Visit(v, n.Chan)
		Visit(v, n.Value)

	case *ast.IncDecStmt:
		Visit(v, n.X)

	case *ast.AssignStmt:
		walkExprList(v, n.Lhs)
		walkExprList(v, n.Rhs)

	case *ast.GoStmt:
		Visit(v, n.Call)

	case *ast.DeferStmt:
		Visit(v, n.Call)

	case *ast.ReturnStmt:
		walkExprList(v, n.Results)

	case *ast.BranchStmt:
		if n.Label != nil {
			Visit(v, n.Label)
		}

	case *ast.BlockStmt:
		walkStmtList(v, n.List)

	case *ast.IfStmt:
		if n.Init != nil {
			Visit(v, n.Init)
		}
		Visit(v, n.Cond)
		Visit(v, n.Body)
		if n.Else != nil {
			Visit(v, n.Else)
		}

	case *ast.CaseClause:
		walkExprList(v, n.List)
		walkStmtList(v, n.Body)

	case *ast.SwitchStmt:
		if n.Init != nil {
			Visit(v, n.Init)
		}
		if n.Tag != nil {
			Visit(v, n.Tag)
		}
		Visit(v, n.Body)

	case *ast.TypeSwitchStmt:
		if n.Init != nil {
			Visit(v, n.Init)
		}
		Visit(v, n.Assign)
		Visit(v, n.Body)

	case *ast.CommClause:
		if n.Comm != nil {
			Visit(v, n.Comm)
		}
		walkStmtList(v, n.Body)

	case *ast.SelectStmt:
		Visit(v, n.Body)

	case *ast.ForStmt:
		if n.Init != nil {
			Visit(v, n.Init)
		}
		if n.Cond != nil {
			Visit(v, n.Cond)
		}
		if n.Post != nil {
			Visit(v, n.Post)
		}
		Visit(v, n.Body)

	case *ast.RangeStmt:
		if n.Key != nil {
			Visit(v, n.Key)
		}
		if n.Value != nil {
			Visit(v, n.Value)
		}
		Visit(v, n.X)
		Visit(v, n.Body)

	// Declarations
	case *ast.ImportSpec:
		if n.Doc != nil {
			Visit(v, n.Doc)
		}
		if n.Name != nil {
			Visit(v, n.Name)
		}
		Visit(v, n.Path)
		if n.Comment != nil {
			Visit(v, n.Comment)
		}

	case *ast.ValueSpec:
		if n.Doc != nil {
			Visit(v, n.Doc)
		}
		walkIdentList(v, n.Names)
		if n.Type != nil {
			Visit(v, n.Type)
		}
		walkExprList(v, n.Values)
		if n.Comment != nil {
			Visit(v, n.Comment)
		}

	case *ast.TypeSpec:
		if n.Doc != nil {
			Visit(v, n.Doc)
		}
		Visit(v, n.Name)
		Visit(v, n.Type)
		if n.Comment != nil {
			Visit(v, n.Comment)
		}

	case *ast.BadDecl:
		// nothing to do

	case *ast.GenDecl:
		if n.Doc != nil {
			Visit(v, n.Doc)
		}
		for _, s := range n.Specs {
			Visit(v, s)
		}

	case *ast.FuncDecl:
		if n.Doc != nil {
			Visit(v, n.Doc)
		}
		if n.Recv != nil {
			Visit(v, n.Recv)
		}
		Visit(v, n.Name)
		Visit(v, n.Type)
		if n.Body != nil {
			Visit(v, n.Body)
		}

	// Files and packages
	case *ast.File:
		if n.Doc != nil {
			Visit(v, n.Doc)
		}
		Visit(v, n.Name)
		walkDeclList(v, n.Decls)
		// don't walk n.Comments - they have been
		// visited already through the individual
		// nodes

	case *ast.Package:
		for _, f := range n.Files {
			Visit(v, f)
		}

	default:
		panic(fmt.Sprintf("ast.Visit: unexpected node type %T", n))
	}

	endWalkNode(v, node)

	v.BeginWalk(nil)
}

func walkNode(v Walker, node ast.Node) bool {
	rets := callMethod(node, v, getWalkNodeMethodName(node))
	return rets[0].Bool()
}

func endWalkNode(v Walker, node ast.Node) {
	callMethod(node, v, getEndWalkNodeMethodName(node))
}

func getWalkNodeMethodName(node ast.Node) string {
	nodeName := getNodeName(node)
	return fmt.Sprintf("Walk%s", nodeName)
}

func getEndWalkNodeMethodName(node ast.Node) string {
	nodeName := getNodeName(node)
	return fmt.Sprintf("EndWalk%s", nodeName)
}

func callMethod(node ast.Node, v Walker, walkMethodName string) []reflect.Value {
	walkMethod := getMethod(v, walkMethodName)
	return walkMethod.Call([]reflect.Value{reflect.ValueOf(node)})
}

func getMethod(v Walker, walkMethodName string) reflect.Value {
	value := reflect.ValueOf(v)
	walkMethod := value.MethodByName(walkMethodName)
	return walkMethod
}

func getNodeName(node ast.Node) string {
	nodeType := reflect.TypeOf(node).Elem()
	nodeTypeName := nodeType.Name()
	return nodeTypeName
}
