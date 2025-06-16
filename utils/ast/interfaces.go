package ast

import (
	"go/ast"
	"io"
)

type Ast interface {
	// Parse /
	Parse(filename string, writer io.Writer) (file *ast.File, err error)
	// Rollback
	Rollback(file *ast.File) error
	// Injection
	Injection(file *ast.File) error
	// Format
	Format(filename string, writer io.Writer, file *ast.File) error
}
