package system

import (
	"github.com/gin-gonic/gin"
)

type OperationRecordRouter struct{}

func (s *OperationRecordRouter) InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("sysOperationRecord")
	{
		operationRecordRouter.DELETE("deleteSysOperationRecord", operationRecordApi.DeleteSysOperationRecord)           // SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecordByIds", operationRecordApi.DeleteSysOperationRecordByIds) // SysOperationRecord
		operationRecordRouter.GET("findSysOperationRecord", operationRecordApi.FindSysOperationRecord)                  // IDSysOperationRecord
		operationRecordRouter.GET("getSysOperationRecordList", operationRecordApi.GetSysOperationRecordList)            // SysOperationRecord

	}
}
