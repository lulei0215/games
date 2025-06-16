// SysExportTemplate
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SysExportTemplate
type SysExportTemplate struct {
	global.GVA_MODEL
	DBName       string         `json:"dbName" form:"dbName" gorm:"column:db_name;comment:;"`                    //
	Name         string         `json:"name" form:"name" gorm:"column:name;comment:;"`                           //
	TableName    string         `json:"tableName" form:"tableName" gorm:"column:table_name;comment:;"`           //
	TemplateID   string         `json:"templateID" form:"templateID" gorm:"column:template_id;comment:;"`        //
	TemplateInfo string         `json:"templateInfo" form:"templateInfo" gorm:"column:template_info;type:text;"` //
	Limit        *int           `json:"limit" form:"limit" gorm:"column:limit;comment:"`
	Order        string         `json:"order" form:"order" gorm:"column:order;comment:"`
	Conditions   []Condition    `json:"conditions" form:"conditions" gorm:"foreignKey:TemplateID;references:TemplateID;comment:"`
	JoinTemplate []JoinTemplate `json:"joinTemplate" form:"joinTemplate" gorm:"foreignKey:TemplateID;references:TemplateID;comment:"`
}

type JoinTemplate struct {
	global.GVA_MODEL
	TemplateID string `json:"templateID" form:"templateID" gorm:"column:template_id;comment:"`
	JOINS      string `json:"joins" form:"joins" gorm:"column:joins;comment:"`
	Table      string `json:"table" form:"table" gorm:"column:table;comment:"`
	ON         string `json:"on" form:"on" gorm:"column:on;comment:"`
}

func (JoinTemplate) TableName() string {
	return "sys_export_template_join"
}

type Condition struct {
	global.GVA_MODEL
	TemplateID string `json:"templateID" form:"templateID" gorm:"column:template_id;comment:"`
	From       string `json:"from" form:"from" gorm:"column:from;comment:key"`
	Column     string `json:"column" form:"column" gorm:"column:column;comment:"`
	Operator   string `json:"operator" form:"operator" gorm:"column:operator;comment:"`
}

func (Condition) TableName() string {
	return "sys_export_template_condition"
}
