package mcpTool

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&{{.Name | title}}{})
}

type {{.Name | title}} struct {
}

// {{.Description}}
func (t *{{.Name | title}}) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// TODO: 
	// :
	// {{- range .Params}}
	// {{.Name}} := request.Params.Arguments["{{.Name}}"]
	// {{- end}}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			{{- range .Response}}
			mcp.{{.Type | title}}Content{
				Type: "{{.Type}}",
				// TODO: {{.Type}}
			},
			{{- end}}
		},
	}, nil
}

func (t *{{.Name | title}}) New() mcp.Tool {
	return mcp.NewTool("{{.Name}}",
		mcp.WithDescription("{{.Description}}"),
		{{- range .Params}}
		mcp.With{{.Type | title}}("{{.Name}}",
			{{- if .Required}}mcp.Required(),{{end}}
			mcp.Description("{{.Description}}"),
			{{- if .Default}}
              {{- if eq .Type "string"}}
              mcp.DefaultString("{{.Default}}"),
              {{- else if eq .Type "number"}}
              mcp.DefaultNumber({{.Default}}),
              {{- else if eq .Type "boolean"}}
              mcp.DefaultBoolean({{if or (eq .Default "true") (eq .Default "True")}}true{{else}}false{{end}}),
              {{- else if eq .Type "array"}}
              // ：
              // mcp.DefaultArray({{.Default}}),
              {{- end}}
            {{- end}}
		),
		{{- end}}
	)
}
