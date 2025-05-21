package initialize

import (
	bbs2 "github.com/flipped-aurora/gin-vue-admin/server/model/entity/bbs"
	example2 "github.com/flipped-aurora/gin-vue-admin/server/model/entity/example"
	system2 "github.com/flipped-aurora/gin-vue-admin/server/model/entity/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/entity/video"
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(

		system2.SysApi{},
		system2.SysUser{},
		system2.SysBaseMenu{},
		system2.JwtBlacklist{},
		system2.SysAuthority{},
		system2.SysDictionary{},
		system2.SysOperationRecord{},
		system2.SysAutoCodeHistory{},
		system2.SysDictionaryDetail{},
		system2.SysBaseMenuParameter{},
		system2.SysBaseMenuBtn{},
		system2.SysAuthorityBtn{},
		system2.SysAutoCode{},
		system2.SysChatGptOption{},

		example2.ExaFile{},
		example2.ExaCustomer{},
		example2.ExaFileChunk{},

		// 注册帖子列表
		bbs2.XkBbs{},
		// 注册帖子分类
		bbs2.BbsCategory{},
		// 注册视频分类表
		video.XkVideoCategory{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
