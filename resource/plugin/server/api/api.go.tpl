package api

import (
{{if not .OnlyTemplate}}
	"{{.Module}}/global"
    "{{.Module}}/model/common/response"
    "{{.Module}}/plugin/{{.Package}}/model"
    {{- if not .IsTree}}
    "{{.Module}}/plugin/{{.Package}}/model/request"
    {{- end }}
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    {{- if .AutoCreateResource}}
    "{{.Module}}/utils"
    {{- end }}
{{- else }}
    "{{.Module}}/model/common/response"
    "github.com/gin-gonic/gin"
{{- end }}
)

var {{.StructName}} = new({{.Abbreviation}})

type {{.Abbreviation}} struct {}
{{if not .OnlyTemplate}}
// Create{{.StructName}} {{.Description}}
// @Tags {{.StructName}}
// @Summary {{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "{{.Description}}"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /{{.Abbreviation}}/create{{.StructName}} [post]
func (a *{{.Abbreviation}}) Create{{.StructName}}(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	var info model.{{.StructName}}
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	{{- if .AutoCreateResource }}
    info.CreatedBy = utils.GetUserID(c)
	{{- end }}
	err = service{{ .StructName }}.Create{{.StructName}}(ctx,&info)
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
// @Param data body model.{{.StructName}} true "{{.Description}}"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /{{.Abbreviation}}/delete{{.StructName}} [delete]
func (a *{{.Abbreviation}}) Delete{{.StructName}}(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	{{.PrimaryField.FieldJson}} := c.Query("{{.PrimaryField.FieldJson}}")
{{- if .AutoCreateResource }}
    userID := utils.GetUserID(c)
{{- end }}
	err := service{{ .StructName }}.Delete{{.StructName}}(ctx,{{.PrimaryField.FieldJson}} {{- if .AutoCreateResource -}},userID{{- end -}})
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
func (a *{{.Abbreviation}}) Delete{{.StructName}}ByIds(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	{{.PrimaryField.FieldJson}}s := c.QueryArray("{{.PrimaryField.FieldJson}}s[]")
{{- if .AutoCreateResource }}
    userID := utils.GetUserID(c)
{{- end }}
	err := service{{ .StructName }}.Delete{{.StructName}}ByIds(ctx,{{.PrimaryField.FieldJson}}s{{- if .AutoCreateResource }},userID{{- end }})
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
// @Param data body model.{{.StructName}} true "{{.Description}}"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /{{.Abbreviation}}/update{{.StructName}} [put]
func (a *{{.Abbreviation}}) Update{{.StructName}}(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	var info model.{{.StructName}}
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
{{- if .AutoCreateResource }}
    info.UpdatedBy = utils.GetUserID(c)
{{- end }}
	err = service{{ .StructName }}.Update{{.StructName}}(ctx,info)
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
// @Success 200 {object} response.Response{data=model.{{.StructName}},msg=string} ""
// @Router /{{.Abbreviation}}/find{{.StructName}} [get]
func (a *{{.Abbreviation}}) Find{{.StructName}}(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	{{.PrimaryField.FieldJson}} := c.Query("{{.PrimaryField.FieldJson}}")
	re{{.Abbreviation}}, err := service{{ .StructName }}.Get{{.StructName}}(ctx,{{.PrimaryField.FieldJson}})
	if err != nil {
        global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":" + err.Error(), c)
		return
	}
    response.OkWithData(re{{.Abbreviation}}, c)
}

{{- if .IsTree }}
// Get{{.StructName}}List {{.Description}}
// @Tags {{.StructName}}
// @Summary {{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} ""
// @Router /{{.Abbreviation}}/get{{.StructName}}List [get]
func (a *{{.Abbreviation}}) Get{{.StructName}}List(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	list, err := service{{ .StructName }}.Get{{.StructName}}InfoList(ctx)
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
// @Param data query request.{{.StructName}}Search true "{{.Description}}"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} ""
// @Router /{{.Abbreviation}}/get{{.StructName}}List [get]
func (a *{{.Abbreviation}}) Get{{.StructName}}List(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	var pageInfo request.{{.StructName}}Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := service{{ .StructName }}.Get{{.StructName}}InfoList(ctx,pageInfo)
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
func (a *{{.Abbreviation}}) Get{{.StructName}}DataSource(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

    // 
   dataSource, err := service{{ .StructName }}.Get{{.StructName}}DataSource(ctx)
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
func (a *{{.Abbreviation}}) Get{{.StructName}}Public(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

    //  ，C，
    service{{ .StructName }}.Get{{.StructName}}Public(ctx)
    response.OkWithDetailed(gin.H{"info": "{{.Description}}"}, "", c)
}
