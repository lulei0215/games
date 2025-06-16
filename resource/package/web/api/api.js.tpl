import service from '@/utils/request'

{{- if not .OnlyTemplate}}
// @Tags {{.StructName}}
// @Summary {{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":""}"
// @Router /{{.Abbreviation}}/create{{.StructName}} [post]
export const create{{.StructName}} = (data) => {
  return service({
    url: '/{{.Abbreviation}}/create{{.StructName}}',
    method: 'post',
    data
  })
}

// @Tags {{.StructName}}
// @Summary {{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":""}"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [delete]
export const delete{{.StructName}} = (params) => {
  return service({
    url: '/{{.Abbreviation}}/delete{{.StructName}}',
    method: 'delete',
    params
  })
}

// @Tags {{.StructName}}
// @Summary {{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":""}"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [delete]
export const delete{{.StructName}}ByIds = (params) => {
  return service({
    url: '/{{.Abbreviation}}/delete{{.StructName}}ByIds',
    method: 'delete',
    params
  })
}

// @Tags {{.StructName}}
// @Summary {{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":""}"
// @Router /{{.Abbreviation}}/update{{.StructName}} [put]
export const update{{.StructName}} = (data) => {
  return service({
    url: '/{{.Abbreviation}}/update{{.StructName}}',
    method: 'put',
    data
  })
}

// @Tags {{.StructName}}
// @Summary id{{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.{{.StructName}} true "id{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":""}"
// @Router /{{.Abbreviation}}/find{{.StructName}} [get]
export const find{{.StructName}} = (params) => {
  return service({
    url: '/{{.Abbreviation}}/find{{.StructName}}',
    method: 'get',
    params
  })
}

// @Tags {{.StructName}}
// @Summary {{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":""}"
// @Router /{{.Abbreviation}}/get{{.StructName}}List [get]
export const get{{.StructName}}List = (params) => {
  return service({
    url: '/{{.Abbreviation}}/get{{.StructName}}List',
    method: 'get',
    params
  })
}

{{- if .HasDataSource}}
// @Tags {{.StructName}}
// @Summary 
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":""}"
// @Router /{{.Abbreviation}}/find{{.StructName}}DataSource [get]
export const get{{.StructName}}DataSource = () => {
  return service({
    url: '/{{.Abbreviation}}/get{{.StructName}}DataSource',
    method: 'get',
  })
}
{{- end}}

{{- end}}

// @Tags {{.StructName}}
// @Summary {{.Description}}
// @Accept application/json
// @Produce application/json
// @Param data query {{.Package}}Req.{{.StructName}}Search true "{{.Description}}"
// @Success 200 {object} response.Response{data=object,msg=string} ""
// @Router /{{.Abbreviation}}/get{{.StructName}}Public [get]
export const get{{.StructName}}Public = () => {
  return service({
    url: '/{{.Abbreviation}}/get{{.StructName}}Public',
    method: 'get',
  })
}
