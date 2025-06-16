{{if .IsPlugin}}
// {{.FuncName}} {{.FuncDesc}}
// @Tags {{.StructName}}
// @Summary {{.FuncDesc}}
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} ""
// @Router /{{.Abbreviation}}/{{.Router}} [{{.Method}}]
func (a *{{.Abbreviation}}) {{.FuncName}}(c *gin.Context) {
    // Context
    ctx := c.Request.Context()
    // 
    err := service{{ .StructName }}.{{.FuncName}}(ctx)
       if err != nil {
    		global.GVA_LOG.Error("!", zap.Error(err))
            response.FailWithMessage("", c)
    		return
       }
    response.OkWithData("",c)
}

{{- else -}}

// {{.FuncName}} {{.FuncDesc}}
// @Tags {{.StructName}}
// @Summary {{.FuncDesc}}
// @Accept application/json
// @Produce application/json
// @Param data query {{.Package}}Req.{{.StructName}}Search true ""
// @Success 200 {object} response.Response{data=object,msg=string} ""
// @Router /{{.Abbreviation}}/{{.Router}} [{{.Method}}]
func ({{.Abbreviation}}Api *{{.StructName}}Api){{.FuncName}}(c *gin.Context) {
    // Context
    ctx := c.Request.Context()
    // 
    err := {{.Abbreviation}}Service.{{.FuncName}}(ctx)
    if err != nil {
        global.GVA_LOG.Error("!", zap.Error(err))
   		response.FailWithMessage("", c)
   		return
   	}
   	response.OkWithData("",c)
}
{{end}}
