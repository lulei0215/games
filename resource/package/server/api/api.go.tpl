package {{.Package}}

import (
	{{if not .OnlyTemplate}}
	"{{.Module}}/global"
    "{{.Module}}/model/common/response"
    "{{.Module}}/model/{{.Package}}"
    {{- if not .IsTree}}
    {{.Package}}Req "{{.Module}}/model/{{.Package}}/request"
    {{- end }}
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    {{- if .AutoCreateResource}}
    "{{.Module}}/utils"
    {{- end }}
    {{- else}}
    "{{.Module}}/model/common/response"
    "github.com/gin-gonic/gin"
    {{- end}}
)

type {{.StructName}}Api struct {}

{{if not .OnlyTemplate}}

// Create{{.StructName}} {{.Description}}
// @Tags {{.StructName}}
// @Summary {{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body {{.Package}}.{{.StructName}} true "{{.Description}}"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /{{.Abbreviation}}/create{{.StructName}} [post]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Create{{.StructName}}(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	var {{.Abbreviation}} {{.Package}}.{{.StructName}}
	err := c.ShouldBindJSON(&{{.Abbreviation}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	{{- if .AutoCreateResource }}
    {{.Abbreviation}}.CreatedBy = utils.GetUserID(c)
	{{- end }}
	err = {{.Abbreviation}}Service.Create{{.StructName}}(ctx,&{{.Abbreviation}})
	if err != nil {
        global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":" + err.Error(), c)
		return
	}
    response.OkWithMessage("", c)
}

// Delete{{.StructName}} {{.Description}}
// @Tags {{.StructName}}
// @Summary {{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body {{.Package}}.{{.StructName}} true "{{.Description}}"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /{{.Abbreviation}}/delete{{.StructName}} [delete]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Delete{{.StructName}}(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	{{.PrimaryField.FieldJson}} := c.Query("{{.PrimaryField.FieldJson}}")
		{{- if .AutoCreateResource }}
    userID := utils.GetUserID(c)
        {{- end }}
	err := {{.Abbreviation}}Service.Delete{{.StructName}}(ctx,{{.PrimaryField.FieldJson}} {{- if .AutoCreateResource -}},userID{{- end -}})
	if err != nil {
        global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":" + err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// Delete{{.StructName}}ByIds {{.Description}}
// @Tags {{.StructName}}
// @Summary {{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} ""
// @Router /{{.Abbreviation}}/delete{{.StructName}}ByIds [delete]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Delete{{.StructName}}ByIds(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	{{.PrimaryField.FieldJson}}s := c.QueryArray("{{.PrimaryField.FieldJson}}s[]")
    	{{- if .AutoCreateResource }}
    userID := utils.GetUserID(c)
        {{- end }}
	err := {{.Abbreviation}}Service.Delete{{.StructName}}ByIds(ctx,{{.PrimaryField.FieldJson}}s{{- if .AutoCreateResource }},userID{{- end }})
	if err != nil {
        global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":" + err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// Update{{.StructName}} {{.Description}}
// @Tags {{.StructName}}
// @Summary {{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body {{.Package}}.{{.StructName}} true "{{.Description}}"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /{{.Abbreviation}}/update{{.StructName}} [put]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Update{{.StructName}}(c *gin.Context) {
    // ctxcontext
    ctx := c.Request.Context()

	var {{.Abbreviation}} {{.Package}}.{{.StructName}}
	err := c.ShouldBindJSON(&{{.Abbreviation}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	    {{- if .AutoCreateResource }}
    {{.Abbreviation}}.UpdatedBy = utils.GetUserID(c)
        {{- end }}
	err = {{.Abbreviation}}Service.Update{{.StructName}}(ctx,{{.Abbreviation}})
	if err != nil {
        global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":" + err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// Find{{.StructName}} id{{.Description}}
// @Tags {{.StructName}}
// @Summary id{{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param {{.PrimaryField.FieldJson}} query {{.PrimaryField.FieldType}} true "id{{.Description}}"
// @Success 200 {object} response.Response{data={{.Package}}.{{.StructName}},msg=string} ""
// @Router /{{.Abbreviation}}/find{{.StructName}} [get]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Find{{.StructName}}(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	{{.PrimaryField.FieldJson}} := c.Query("{{.PrimaryField.FieldJson}}")
	re{{.Abbreviation}}, err := {{.Abbreviation}}Service.Get{{.StructName}}(ctx,{{.PrimaryField.FieldJson}})
	if err != nil {
        global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":" + err.Error(), c)
		return
	}
	response.OkWithData(re{{.Abbreviation}}, c)
}

{{- if .IsTree }}
// Get{{.StructName}}List {{.Description}},Tree
// @Tags {{.StructName}}
// @Summary {{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} ""
// @Router /{{.Abbreviation}}/get{{.StructName}}List [get]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Get{{.StructName}}List(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	list, err := {{.Abbreviation}}Service.Get{{.StructName}}InfoList(ctx)
	if err != nil {
	    global.GVA_LOG.Error("!", zap.Error(err))
        response.FailWithMessage(":" + err.Error(), c)
        return
    }
    response.OkWithDetailed(list, "", c)
}
{{- else }}
// Get{{.StructName}}List {{.Description}}
// @Tags {{.StructName}}
// @Summary {{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query {{.Package}}Req.{{.StructName}}Search true "{{.Description}}"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} ""
// @Router /{{.Abbreviation}}/get{{.StructName}}List [get]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Get{{.StructName}}List(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	var pageInfo {{.Package}}Req.{{.StructName}}Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := {{.Abbreviation}}Service.Get{{.StructName}}InfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("!", zap.Error(err))
        response.FailWithMessage(":" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "", c)
}
{{- end }}

{{- if .HasDataSource }}
// Get{{.StructName}}DataSource {{.StructName}}
// @Tags {{.StructName}}
// @Summary {{.StructName}}
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} ""
// @Router /{{.Abbreviation}}/get{{.StructName}}DataSource [get]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Get{{.StructName}}DataSource(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

    // 
    dataSource, err := {{.Abbreviation}}Service.Get{{.StructName}}DataSource(ctx)
    if err != nil {
        global.GVA_LOG.Error("!", zap.Error(err))
   		response.FailWithMessage(":" + err.Error(), c)
   		return
    }
   response.OkWithData(dataSource, c)
}
{{- end }}

{{- end }}

// Get{{.StructName}}Public {{.Description}}
// @Tags {{.StructName}}
// @Summary {{.Description}}
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} ""
// @Router /{{.Abbreviation}}/get{{.StructName}}Public [get]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Get{{.StructName}}Public(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

    // 
    // ，C，
    {{.Abbreviation}}Service.Get{{.StructName}}Public(ctx)
    response.OkWithDetailed(gin.H{
       "info": "{{.Description}}",
    }, "", c)
}
