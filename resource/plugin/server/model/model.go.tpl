{{- if .IsAdd}}
// 
{{- range .Fields}}
  {{ GenerateField . }}
{{- end }}

{{ else }}
package model

{{- if not .OnlyTemplate}}
import (
	{{- if .GvaModel }}
	"{{.Module}}/global"
	{{- end }}
	{{- if or .HasTimer }}
	"time"
	{{- end }}
	{{- if .NeedJSON }}
	"gorm.io/datatypes"
	{{- end }}
)
{{- end }}

// {{.StructName}} {{.Description}} 
type {{.StructName}} struct {
{{- if not .OnlyTemplate}}
{{- if .GvaModel }}
    global.GVA_MODEL
{{- end }}
{{- range .Fields}}
  {{ GenerateField . }}
{{- end }}
    {{- if .AutoCreateResource }}
    CreatedBy  uint   `gorm:"column:created_by;comment:"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:"`
    {{- end }}
    {{- if .IsTree }}
    Children   []*{{.StructName}} `json:"children" gorm:"-"`     //
    ParentID   int             `json:"parentID" gorm:"column:parent_id;comment:"`
    {{- end }}
    {{- end }}
}

{{ if .TableName }}
// TableName {{.Description}} {{.StructName}} {{.TableName}}
func ({{.StructName}}) TableName() string {
    return "{{.TableName}}"
}
{{ end }}


{{if .IsTree }}
// GetChildren TreeNode
func (s *{{.StructName}}) GetChildren() []*{{.StructName}} {
    return s.Children
}

// SetChildren TreeNode
func (s *{{.StructName}}) SetChildren(children *{{.StructName}}) {
	s.Children = append(s.Children, children)
}

// GetID TreeNode
func (s *{{.StructName}}) GetID() int {
    return int({{if not .GvaModel}}*{{- end }}s.{{.PrimaryField.FieldName}})
}

// GetParentID TreeNode
func (s *{{.StructName}}) GetParentID() int {
    return s.ParentID
}
{{ end }}


{{ end }}
