package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ExaAttachmentCategory struct {
	global.GVA_MODEL
	Name     string                   `json:"name" form:"name" gorm:"default:null;type:varchar(255);column:name;comment:;"`
	Pid      uint                     `json:"pid" form:"pid" gorm:"default:0;type:int;column:pid;comment:ID;"`
	Children []*ExaAttachmentCategory `json:"children" gorm:"-"`
}

func (ExaAttachmentCategory) TableName() string {
	return "exa_attachment_category"
}
