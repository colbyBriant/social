package main

import (
	"fmt"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title						Swagger Example API
// @version					0.0.1
// @description				This is a sample Server pets
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						x-token
// @BasePath					/
func main() {
	// 初始化Viper ---解析config.yaml
	global.GVA_VP = core.Viper()

	fmt.Println(global.GVA_CONFIG.XkBbs.Appid)
	fmt.Println(global.GVA_CONFIG.XkBbs.Level)
	initialize.OtherInit()
	// 初始化zap日志库
	global.GVA_LOG = core.Zap()
	zap.ReplaceGlobals(global.GVA_LOG)

	// gorm连接数据库
	global.GVA_DB = initialize.Gorm()

	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

	core.RunWindowsServer()

}
