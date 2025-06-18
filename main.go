package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

//  @Tag ,
// swag init  @Tag ,  main.go
//  --generalInfo flag
// @Tag.Name        Base
// @Tag.Name        SysUser
// @Tag.Description

// @title                       Gin-Vue-Admin Swagger API
// @version                     v2.8.2
// @description                 gin+vue
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	//
	initializeSystem()
	//
	core.RunServer()
}

// initializeSystem
func initializeSystem() {
	global.GVA_VP = core.Viper() // Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // zap
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm
	initialize.Timer()
	initialize.DBList()
	initialize.SetupHandlers() //
	if global.GVA_DB != nil {
		initialize.RegisterTables() //
	}
}
