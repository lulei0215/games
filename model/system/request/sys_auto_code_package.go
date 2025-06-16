package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type SysAutoCodePackageCreate struct {
	Desc        string `json:"desc" example:""`
	Label       string `json:"label" example:""`
	Template    string `json:"template"  example:""`
	PackageName string `json:"packageName" example:""`
	Module      string `json:"-" example:""`
}

func (r *SysAutoCodePackageCreate) AutoCode() AutoCode {
	return AutoCode{
		Package: r.PackageName,
		Module:  global.GVA_CONFIG.AutoCode.Module,
	}
}

func (r *SysAutoCodePackageCreate) Create() model.SysAutoCodePackage {
	return model.SysAutoCodePackage{
		Desc:        r.Desc,
		Label:       r.Label,
		Template:    r.Template,
		PackageName: r.PackageName,
		Module:      global.GVA_CONFIG.AutoCode.Module,
	}
}
