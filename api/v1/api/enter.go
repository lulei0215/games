package api

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ SysTransactionsApi }

var sysTransactionsService = service.ServiceGroupApp.ApiServiceGroup.SysTransactionsService
