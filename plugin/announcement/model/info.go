package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// Info
type Info struct {
	global.GVA_MODEL
	Title       string         `json:"title" form:"title" gorm:"column:title;comment:;"`                                              //
	Content     string         `json:"content" form:"content" gorm:"column:content;comment:;type:text;"`                              //
	UserID      *int           `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`                                          //
	Attachments datatypes.JSON `json:"attachments" form:"attachments" gorm:"column:attachments;comment:;" swaggertype:"array,object"` //
}

// TableName  Info gva_announcements_info
func (Info) TableName() string {
	return "gva_announcements_info"
}
