package ast

import (
	"fmt"
	"go/ast"
	"io"
)

type PluginInitializeV2 struct {
	Base
	Type         Type   //
	Path         string //
	PluginPath   string //
	RelativePath string //
	ImportPath   string //
	StructName   string //
	PackageName  string //
}

func (a *PluginInitializeV2) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
	if filename == "" {
		if a.RelativePath == "" {
			filename = a.PluginPath
			a.RelativePath = a.Base.RelativePath(a.PluginPath)
			return a.Base.Parse(filename, writer)
		}
		a.PluginPath = a.Base.AbsolutePath(a.RelativePath)
		filename = a.PluginPath
	}
	return a.Base.Parse(filename, writer)
}

func (a *PluginInitializeV2) Injection(file *ast.File) error {
	if !CheckImport(file, a.ImportPath) {
		NewImport(a.ImportPath).Injection(file)
		funcDecl := FindFunction(file, "bizPluginV2")
		stmt := CreateStmt(fmt.Sprintf("PluginInitV2(engine, %s.Plugin)", a.PackageName))
		funcDecl.Body.List = append(funcDecl.Body.List, stmt)
	}
	return nil
}

func (a *PluginInitializeV2) Rollback(file *ast.File) error {
	return nil
}

func (a *PluginInitializeV2) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.PluginPath
	}
	return a.Base.Format(filename, writer, file)
}
