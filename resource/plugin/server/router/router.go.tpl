package router

import (
	{{if .OnlyTemplate }} // {{end}}"{{.Module}}/middleware"
	"github.com/gin-gonic/gin"
)

var {{.StructName}} = new({{.Abbreviation}})

type {{.Abbreviation}} struct {}

// Init  {{.Description}} 
func (r *{{.Abbreviation}}) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
{{- if not .OnlyTemplate }}
	{
	    group := private.Group("{{.Abbreviation}}").Use(middleware.OperationRecord())
		group.POST("create{{.StructName}}", api{{.StructName}}.Create{{.StructName}})   // {{.Description}}
		group.DELETE("delete{{.StructName}}", api{{.StructName}}.Delete{{.StructName}}) // {{.Description}}
		group.DELETE("delete{{.StructName}}ByIds", api{{.StructName}}.Delete{{.StructName}}ByIds) // {{.Description}}
		group.PUT("update{{.StructName}}", api{{.StructName}}.Update{{.StructName}})    // {{.Description}}
	}
	{
	    group := private.Group("{{.Abbreviation}}")
		group.GET("find{{.StructName}}", api{{.StructName}}.Find{{.StructName}})        // ID{{.Description}}
		group.GET("get{{.StructName}}List", api{{.StructName}}.Get{{.StructName}}List)  // {{.Description}}
	}
	{
	    group := public.Group("{{.Abbreviation}}")
    	{{- if .HasDataSource}}
	    group.GET("get{{.StructName}}DataSource", api{{.StructName}}.Get{{.StructName}}DataSource)  // {{.Description}}
	    {{- end}}
	    group.GET("get{{.StructName}}Public", api{{.StructName}}.Get{{.StructName}}Public)  // {{.Description}}
	}
{{- else}}
     // {
	 //   group := private.Group("{{.Abbreviation}}").Use(middleware.OperationRecord())
	 // }
	 // {
     //   group := private.Group("{{.Abbreviation}}")
     // }
    {
	    group := public.Group("{{.Abbreviation}}")
	    group.GET("get{{.StructName}}Public", api{{.StructName}}.Get{{.StructName}}Public)  // {{.Description}}
    }
{{- end}}
}
