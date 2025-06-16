package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ExaFileUploadAndDownload struct {
	global.GVA_MODEL
	Name    string `json:"name" form:"name" gorm:"column:name;comment:"`                                 //
	ClassId int    `json:"classId" form:"classId" gorm:"default:0;type:int;column:class_id;comment:id;"` // id
	Url     string `json:"url" form:"url" gorm:"column:url;comment:"`                                    //
	Tag     string `json:"tag" form:"tag" gorm:"column:tag;comment:"`                                    //
	Key     string `json:"key" form:"key" gorm:"column:key;comment:"`                                    //
}

func (ExaFileUploadAndDownload) TableName() string {
	return "exa_file_upload_and_downloads"
}
