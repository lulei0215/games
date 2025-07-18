package ast

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/pkg/errors"
)

type Base struct{}

func (a *Base) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
	fileSet := token.NewFileSet()
	if writer != nil {
		file, err = parser.ParseFile(fileSet, filename, nil, parser.ParseComments)
	} else {
		file, err = parser.ParseFile(fileSet, filename, writer, parser.ParseComments)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "[filepath:%s]/!", filename)
	}
	return file, nil
}

func (a *Base) Rollback(file *ast.File) error {
	return nil
}

func (a *Base) Injection(file *ast.File) error {
	return nil
}

func (a *Base) Format(filename string, writer io.Writer, file *ast.File) error {
	fileSet := token.NewFileSet()
	if writer == nil {
		open, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0666)
		defer open.Close()
		if err != nil {
			return errors.Wrapf(err, "[filepath:%s]!", filename)
		}
		writer = open
	}
	err := format.Node(writer, fileSet, file)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]!", filename)
	}
	return nil
}

// RelativePath
func (a *Base) RelativePath(filePath string) string {
	server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
	hasServer := strings.Index(filePath, server)
	if hasServer != -1 {
		filePath = strings.TrimPrefix(filePath, server)
		keys := strings.Split(filePath, string(filepath.Separator))
		filePath = path.Join(keys...)
	}
	return filePath
}

// AbsolutePath
func (a *Base) AbsolutePath(filePath string) string {
	server := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
	keys := strings.Split(filePath, "/")
	filePath = filepath.Join(keys...)
	filePath = filepath.Join(server, filePath)
	return filePath
}
