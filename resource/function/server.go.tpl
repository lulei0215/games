{{- $db := "" }}
{{- if eq .BusinessDB "" }}
 {{- $db = "global.GVA_DB" }}
{{- else}}
 {{- $db =  printf "global.MustGetGlobalDBByDBName(\"%s\")" .BusinessDB   }}
{{- end}}
{{if .IsPlugin}}

// {{.FuncName}} {{.FuncDesc}}
// Author [yourname](https://github.com/yourname)
func (s *{{.Abbreviation}}) {{.FuncName}}(ctx context.Context) (err error) {
	db := {{$db}}.Model(&model.{{.StructName}}{})
    return db.Error
}

{{- else -}}

// {{.FuncName}} {{.FuncDesc}}
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service){{.FuncName}}(ctx context.Context) (err error) {
	// 
	db := {{$db}}.Model(&{{.Package}}.{{.StructName}}{})
    return db.Error
}
{{end}}
