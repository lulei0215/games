package {{.Package}}

import (
	{{if .OnlyTemplate}}// {{ end}}"{{.Module}}/middleware"
	"github.com/gin-gonic/gin"
)

type {{.StructName}}Router struct {}

// Init{{.StructName}}Router  {{.Description}} 
func (s *{{.StructName}}Router) Init{{.StructName}}Router(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	{{- if not .OnlyTemplate}}
	{{.Abbreviation}}Router := Router.Group("{{.Abbreviation}}").Use(middleware.OperationRecord())
	{{.Abbreviation}}RouterWithoutRecord := Router.Group("{{.Abbreviation}}")
	{{- else }}
	// {{.Abbreviation}}Router := Router.Group("{{.Abbreviation}}").Use(middleware.OperationRecord())
    // {{.Abbreviation}}RouterWithoutRecord := Router.Group("{{.Abbreviation}}")
	{{- end}}
	{{.Abbreviation}}RouterWithoutAuth := PublicRouter.Group("{{.Abbreviation}}")
	{{- if not .OnlyTemplate}}
	{
		{{.Abbreviation}}Router.POST("create{{.StructName}}", {{.Abbreviation}}Api.Create{{.StructName}})   // {{.Description}}
		{{.Abbreviation}}Router.DELETE("delete{{.StructName}}", {{.Abbreviation}}Api.Delete{{.StructName}}) // {{.Description}}
		{{.Abbreviation}}Router.DELETE("delete{{.StructName}}ByIds", {{.Abbreviation}}Api.Delete{{.StructName}}ByIds) // {{.Description}}
		{{.Abbreviation}}Router.PUT("update{{.StructName}}", {{.Abbreviation}}Api.Update{{.StructName}})    // {{.Description}}
	}
	{
		{{.Abbreviation}}RouterWithoutRecord.GET("find{{.StructName}}", {{.Abbreviation}}Api.Find{{.StructName}})        // ID{{.Description}}
		{{.Abbreviation}}RouterWithoutRecord.GET("get{{.StructName}}List", {{.Abbreviation}}Api.Get{{.StructName}}List)  // {{.Description}}
	}
	{
	{{- if .HasDataSource}}
	    {{.Abbreviation}}RouterWithoutAuth.GET("get{{.StructName}}DataSource", {{.Abbreviation}}Api.Get{{.StructName}}DataSource)  // {{.Description}}
	{{- end}}
	    {{.Abbreviation}}RouterWithoutAuth.GET("get{{.StructName}}Public", {{.Abbreviation}}Api.Get{{.StructName}}Public)  // {{.Description}}
	}
	{{- else}}
	{
	    {{.Abbreviation}}RouterWithoutAuth.GET("get{{.StructName}}Public", {{.Abbreviation}}Api.Get{{.StructName}}Public)  // {{.Description}}
	}
    {{ end }}
}
