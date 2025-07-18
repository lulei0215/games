{{- if .IsAdd}}
// 
{{- range .Fields}}
    {{- if ne .FieldSearchType ""}}
      {{ GenerateSearchField . }}
    {{- end}}
{{- end }}
{{- if .NeedSort}}
Sort  string `json:"sort" form:"sort"`
Order string `json:"order" form:"order"`
{{- end}}
{{- else }}
package request

import (
{{- if not .OnlyTemplate }}
	"{{.Module}}/model/common/request"
	{{ if or .HasSearchTimer .GvaModel }}"time"{{ end }}
{{- end }}
)

type {{.StructName}}Search struct{
{{- if not .OnlyTemplate}}
{{- if .GvaModel }}
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
{{- end }}
{{- range .Fields}}
    {{- if ne .FieldSearchType ""}}
      {{ GenerateSearchField . }}
    {{- end}}
{{- end }}
    request.PageInfo
    {{- if .NeedSort}}
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
    {{- end}}
{{- end}}
}
{{- end }}
