package ast

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

// PackageInitializeRouter
// ModuleName := PackageName.AppName.GroupName
// ModuleName.FunctionName(RouterGroupName)
type PackageInitializeRouter struct {
	Base
	Type                 Type   //
	Path                 string //
	ImportPath           string //
	RelativePath         string //
	AppName              string //
	GroupName            string //
	ModuleName           string //
	PackageName          string //
	FunctionName         string //
	RouterGroupName      string //
	LeftRouterGroupName  string //
	RightRouterGroupName string //
}

func (a *PackageInitializeRouter) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
	if filename == "" {
		if a.RelativePath == "" {
			filename = a.Path
			a.RelativePath = a.Base.RelativePath(a.Path)
			return a.Base.Parse(filename, writer)
		}
		a.Path = a.Base.AbsolutePath(a.RelativePath)
		filename = a.Path
	}
	return a.Base.Parse(filename, writer)
}

func (a *PackageInitializeRouter) Rollback(file *ast.File) error {
	funcDecl := FindFunction(file, "initBizRouter")
	exprNum := 0
	for i := range funcDecl.Body.List {
		if IsBlockStmt(funcDecl.Body.List[i]) {
			if VariableExistsInBlock(funcDecl.Body.List[i].(*ast.BlockStmt), a.ModuleName) {
				for ii, stmt := range funcDecl.Body.List[i].(*ast.BlockStmt).List {
					//  *ast.ExprStmt
					exprStmt, ok := stmt.(*ast.ExprStmt)
					if !ok {
						continue
					}
					//  *ast.CallExpr
					callExpr, ok := exprStmt.X.(*ast.CallExpr)
					if !ok {
						continue
					}
					//
					selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
					if !ok {
						continue
					}
					//  systemRouter.InitApiRouter
					ident, ok := selExpr.X.(*ast.Ident)
					//+1
					if ok && ident.Name == a.ModuleName {
						exprNum++
					}
					//
					if !ok || ident.Name != a.ModuleName || selExpr.Sel.Name != a.FunctionName {
						continue
					}
					exprNum--
					// 。
					funcDecl.Body.List[i].(*ast.BlockStmt).List = append(funcDecl.Body.List[i].(*ast.BlockStmt).List[:ii], funcDecl.Body.List[i].(*ast.BlockStmt).List[ii+1:]...)
					// ，。
					if exprNum == 0 {
						funcDecl.Body.List = append(funcDecl.Body.List[:i], funcDecl.Body.List[i+1:]...)
					}
					break
				}
				break
			}
		}
	}

	return nil
}

func (a *PackageInitializeRouter) Injection(file *ast.File) error {
	funcDecl := FindFunction(file, "initBizRouter")
	hasRouter := false
	var varBlock *ast.BlockStmt
	for i := range funcDecl.Body.List {
		if IsBlockStmt(funcDecl.Body.List[i]) {
			if VariableExistsInBlock(funcDecl.Body.List[i].(*ast.BlockStmt), a.ModuleName) {
				hasRouter = true
				varBlock = funcDecl.Body.List[i].(*ast.BlockStmt)
				break
			}
		}
	}
	if !hasRouter {
		stmt := a.CreateAssignStmt()
		varBlock = &ast.BlockStmt{
			List: []ast.Stmt{
				stmt,
			},
		}
	}
	routerStmt := CreateStmt(fmt.Sprintf("%s.%s(%s,%s)", a.ModuleName, a.FunctionName, a.LeftRouterGroupName, a.RightRouterGroupName))
	varBlock.List = append(varBlock.List, routerStmt)
	if !hasRouter {
		funcDecl.Body.List = append(funcDecl.Body.List, varBlock)
	}
	return nil
}

func (a *PackageInitializeRouter) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}

func (a *PackageInitializeRouter) CreateAssignStmt() *ast.AssignStmt {
	//
	ident := &ast.Ident{
		Name: a.ModuleName,
	}

	//
	selector := &ast.SelectorExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: a.PackageName},
			Sel: &ast.Ident{Name: a.AppName},
		},
		Sel: &ast.Ident{Name: a.GroupName},
	}

	//
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{ident},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{selector},
	}

	return stmt
}
