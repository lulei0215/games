package api

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ SysTransactionsRouter }

var sysTransactionsApi = api.ApiGroupApp.ApiApiGroup.SysTransactionsApi
