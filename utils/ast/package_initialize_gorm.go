package ast

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

// PackageInitializeGorm gorm
type PackageInitializeGorm struct {
	Base
	Type         Type   //
	Path         string //
	ImportPath   string //
	Business     string //  gva => gva, "gva"
	StructName   string //
	PackageName  string //
	RelativePath string //
	IsNew        bool   // new true: new(PackageName.StructName) false: &PackageName.StructName{}
}

func (a *PackageInitializeGorm) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
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

func (a *PackageInitializeGorm) Rollback(file *ast.File) error {
	packageNameNum := 0
	//
	ast.Inspect(file, func(n ast.Node) bool {
		// dbbusiness
		varDB := a.Business + "Db"

		if a.Business == "" {
			varDB = "db"
		}

		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		//  db.AutoMigrate()
		selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok || selExpr.Sel.Name != "AutoMigrate" {
			return true
		}

		//  db
		ident, ok := selExpr.X.(*ast.Ident)
		if !ok || ident.Name != varDB {
			return true
		}

		//
		for i := 0; i < len(callExpr.Args); i++ {
			if com, comok := callExpr.Args[i].(*ast.CompositeLit); comok {
				if selector, exprok := com.Type.(*ast.SelectorExpr); exprok {
					if x, identok := selector.X.(*ast.Ident); identok {
						if x.Name == a.PackageName {
							packageNameNum++
							if selector.Sel.Name == a.StructName {
								callExpr.Args = append(callExpr.Args[:i], callExpr.Args[i+1:]...)
								i--
							}
						}
					}
				}
			}
		}
		return true
	})

	if packageNameNum == 1 {
		_ = NewImport(a.ImportPath).Rollback(file)
	}
	return nil
}

func (a *PackageInitializeGorm) Injection(file *ast.File) error {
	_ = NewImport(a.ImportPath).Injection(file)
	bizModelDecl := FindFunction(file, "bizModel")
	if bizModelDecl != nil {
		a.addDbVar(bizModelDecl.Body)
	}
	//
	ast.Inspect(file, func(n ast.Node) bool {
		// dbbusiness
		varDB := a.Business + "Db"

		if a.Business == "" {
			varDB = "db"
		}

		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		//  db.AutoMigrate()
		selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok || selExpr.Sel.Name != "AutoMigrate" {
			return true
		}

		//  db
		ident, ok := selExpr.X.(*ast.Ident)
		if !ok || ident.Name != varDB {
			return true
		}

		//
		callExpr.Args = append(callExpr.Args, &ast.CompositeLit{
			Type: &ast.SelectorExpr{
				X:   ast.NewIdent(a.PackageName),
				Sel: ast.NewIdent(a.StructName),
			},
		})
		return true
	})
	return nil
}

func (a *PackageInitializeGorm) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}

// businessDB
func (a *PackageInitializeGorm) addDbVar(astBody *ast.BlockStmt) {
	for i := range astBody.List {
		if assignStmt, ok := astBody.List[i].(*ast.AssignStmt); ok {
			if ident, ok := assignStmt.Lhs[0].(*ast.Ident); ok {
				if (a.Business == "" && ident.Name == "db") || ident.Name == a.Business+"Db" {
					return
				}
			}
		}
	}

	//  businessDb := global.GetGlobalDBByDBName("business")
	assignNode := &ast.AssignStmt{
		Lhs: []ast.Expr{
			&ast.Ident{
				Name: a.Business + "Db",
			},
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "global",
					},
					Sel: &ast.Ident{
						Name: "GetGlobalDBByDBName",
					},
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: fmt.Sprintf("\"%s\"", a.Business),
					},
				},
			},
		},
	}

	//  businessDb.AutoMigrate()
	autoMigrateCall := &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X: &ast.Ident{
					Name: a.Business + "Db",
				},
				Sel: &ast.Ident{
					Name: "AutoMigrate",
				},
			},
		},
	}

	returnNode := astBody.List[len(astBody.List)-1]
	astBody.List = append(astBody.List[:len(astBody.List)-1], assignNode, autoMigrateCall, returnNode)
}
