package ast

import (
	"fmt"
	"go/ast"
	"io"
)

// PluginInitializeRouter
// PackageName.AppName.GroupName.FunctionName()
type PluginInitializeRouter struct {
	Base
	Type                 Type   //
	Path                 string //
	ImportPath           string //
	ImportGlobalPath     string //
	ImportMiddlewarePath string //
	RelativePath         string //
	AppName              string //
	GroupName            string //
	PackageName          string //
	FunctionName         string //
	LeftRouterGroupName  string //
	RightRouterGroupName string //
}

func (a *PluginInitializeRouter) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
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

func (a *PluginInitializeRouter) Rollback(file *ast.File) error {
	funcDecl := FindFunction(file, "Router")
	delI := 0
	routerNum := 0
	for i := len(funcDecl.Body.List) - 1; i >= 0; i-- {
		stmt, ok := funcDecl.Body.List[i].(*ast.ExprStmt)
		if !ok {
			continue
		}

		callExpr, ok := stmt.X.(*ast.CallExpr)
		if !ok {
			continue
		}

		selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			continue
		}

		ident, ok := selExpr.X.(*ast.SelectorExpr)

		if ok {
			if iExpr, ieok := ident.X.(*ast.SelectorExpr); ieok {
				if iden, idok := iExpr.X.(*ast.Ident); idok {
					if iden.Name == "router" {
						routerNum++
					}
				}
			}
			if ident.Sel.Name == a.GroupName && selExpr.Sel.Name == a.FunctionName {
				//
				delI = i
			}
		}
	}

	funcDecl.Body.List = append(funcDecl.Body.List[:delI], funcDecl.Body.List[delI+1:]...)

	if routerNum <= 1 {
		_ = NewImport(a.ImportPath).Rollback(file)
	}

	return nil
}

func (a *PluginInitializeRouter) Injection(file *ast.File) error {
	_ = NewImport(a.ImportPath).Injection(file)
	funcDecl := FindFunction(file, "Router")

	var exists bool

	ast.Inspect(funcDecl, func(n ast.Node) bool {
		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		ident, ok := selExpr.X.(*ast.SelectorExpr)
		if ok && ident.Sel.Name == a.GroupName && selExpr.Sel.Name == a.FunctionName {
			exists = true
			return false
		}
		return true
	})

	if !exists {
		stmtStr := fmt.Sprintf("%s.%s.%s.%s(%s, %s)", a.PackageName, a.AppName, a.GroupName, a.FunctionName, a.LeftRouterGroupName, a.RightRouterGroupName)
		stmt := CreateStmt(stmtStr)
		funcDecl.Body.List = append(funcDecl.Body.List, stmt)
	}
	return nil
}

func (a *PluginInitializeRouter) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}
