// SysOperationRecord
package system

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// time.Time import time
type SysOperationRecord struct {
	global.GVA_MODEL
	Ip           string        `json:"ip" form:"ip" gorm:"column:ip;comment:ip"`                                   // ip
	Method       string        `json:"method" form:"method" gorm:"column:method;comment:"`                         //
	Path         string        `json:"path" form:"path" gorm:"column:path;comment:"`                               //
	Status       int           `json:"status" form:"status" gorm:"column:status;comment:"`                         //
	Latency      time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:" swaggertype:"string"` //
	Agent        string        `json:"agent" form:"agent" gorm:"type:text;column:agent;comment:"`                  //
	ErrorMessage string        `json:"error_message" form:"error_message" gorm:"column:error_message;comment:"`    //
	Body         string        `json:"body" form:"body" gorm:"type:text;column:body;comment:Body"`                 // Body
	Resp         string        `json:"resp" form:"resp" gorm:"type:text;column:resp;comment:Body"`                 // Body
	UserID       int           `json:"user_id" form:"user_id" gorm:"column:user_id;comment:id"`                    // id
	User         SysUser       `json:"user"`
}
