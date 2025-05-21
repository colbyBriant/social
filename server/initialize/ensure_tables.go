package initialize

import (
	"context"
	adapter "github.com/casbin/gorm-adapter/v3"
	example2 "github.com/flipped-aurora/gin-vue-admin/server/model/entity/example"
	system2 "github.com/flipped-aurora/gin-vue-admin/server/model/entity/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"gorm.io/gorm"
)

const initOrderEnsureTables = system.InitOrderExternal - 1

type ensureTables struct{}

// auto run
func init() {
	system.RegisterInit(initOrderEnsureTables, &ensureTables{})
}

func (ensureTables) InitializerName() string {
	return "ensure_tables_created"
}
func (e *ensureTables) InitializeData(ctx context.Context) (next context.Context, err error) {
	return ctx, nil
}

func (e *ensureTables) DataInserted(ctx context.Context) bool {
	return true
}

func (e *ensureTables) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	tables := []interface{}{
		system2.SysApi{},
		system2.SysUser{},
		system2.SysBaseMenu{},
		system2.SysAuthority{},
		system2.JwtBlacklist{},
		system2.SysDictionary{},
		system2.SysAutoCodeHistory{},
		system2.SysOperationRecord{},
		system2.SysDictionaryDetail{},
		system2.SysBaseMenuParameter{},
		system2.SysBaseMenuBtn{},
		system2.SysAuthorityBtn{},
		system2.SysAutoCode{},

		adapter.CasbinRule{},

		example2.ExaFile{},
		example2.ExaCustomer{},
		example2.ExaFileChunk{},
		example2.ExaFileUploadAndDownload{},
	}
	for _, t := range tables {
		_ = db.AutoMigrate(&t)
		// 视图 authority_menu 会被当成表来创建，引发冲突错误（更新版本的gorm似乎不会）
		// 由于 AutoMigrate() 基本无需考虑错误，因此显式忽略
	}
	return ctx, nil
}

func (e *ensureTables) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	tables := []interface{}{
		system2.SysApi{},
		system2.SysUser{},
		system2.SysBaseMenu{},
		system2.SysAuthority{},
		system2.JwtBlacklist{},
		system2.SysDictionary{},
		system2.SysAutoCodeHistory{},
		system2.SysOperationRecord{},
		system2.SysDictionaryDetail{},
		system2.SysBaseMenuParameter{},
		system2.SysBaseMenuBtn{},
		system2.SysAuthorityBtn{},
		system2.SysAutoCode{},

		adapter.CasbinRule{},

		example2.ExaFile{},
		example2.ExaCustomer{},
		example2.ExaFileChunk{},
		example2.ExaFileUploadAndDownload{},
	}
	yes := true
	for _, t := range tables {
		yes = yes && db.Migrator().HasTable(t)
	}
	return yes
}
