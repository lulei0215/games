package request

import (
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type SysAutoHistoryCreate struct {
	Table            string            //
	Package          string            // /
	Request          string            //
	StructName       string            //
	BusinessDB       string            //
	Description      string            // Struct
	Injections       map[string]string //
	Templates        map[string]string //
	ApiIDs           []uint            // api
	MenuID           uint              // ID
	ExportTemplateID uint              // ID
}

func (r *SysAutoHistoryCreate) Create() model.SysAutoCodeHistory {
	entity := model.SysAutoCodeHistory{
		Package:          r.Package,
		Request:          r.Request,
		Table:            r.Table,
		StructName:       r.StructName,
		Abbreviation:     r.StructName,
		BusinessDB:       r.BusinessDB,
		Description:      r.Description,
		Injections:       r.Injections,
		Templates:        r.Templates,
		ApiIDs:           r.ApiIDs,
		MenuID:           r.MenuID,
		ExportTemplateID: r.ExportTemplateID,
	}
	if entity.Table == "" {
		entity.Table = r.StructName
	}
	return entity
}

type SysAutoHistoryRollBack struct {
	common.GetById
	DeleteApi   bool `json:"deleteApi" form:"deleteApi"`     //
	DeleteMenu  bool `json:"deleteMenu" form:"deleteMenu"`   //
	DeleteTable bool `json:"deleteTable" form:"deleteTable"` //
}

func (r *SysAutoHistoryRollBack) ApiIds(entity model.SysAutoCodeHistory) common.IdsReq {
	length := len(entity.ApiIDs)
	ids := make([]int, 0)
	for i := 0; i < length; i++ {
		ids = append(ids, int(entity.ApiIDs[i]))
	}
	return common.IdsReq{Ids: ids}
}
