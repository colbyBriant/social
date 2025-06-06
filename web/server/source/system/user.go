package system

import (
	"context"
	system2 "github.com/flipped-aurora/gin-vue-admin/server/model/entity/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const initOrderUser = initOrderAuthority + 1

type initUser struct{}

// auto run
func init() {
	system.RegisterInit(initOrderUser, &initUser{})
}

func (i *initUser) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&system2.SysUser{}, &system2.SysChatGptOption{})
}

func (i *initUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&system2.SysUser{})
}

func (i initUser) InitializerName() string {
	return system2.SysUser{}.TableName()
}

func (i *initUser) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	password := utils.BcryptHash("6447985")
	adminPassword := utils.BcryptHash("123456")

	entities := []system2.SysUser{
		{
			UUID:        uuid.NewV4(),
			Username:    "admin",
			Password:    adminPassword,
			NickName:    "Mr.奇淼",
			HeaderImg:   "https://qmplusimg.henrongyi.top/gva_header.jpg",
			AuthorityId: 888,
			Phone:       "17611111111",
			Email:       "333333333@qq.com",
		},
		{
			UUID:        uuid.NewV4(),
			Username:    "a303176530",
			Password:    password,
			NickName:    "用户1",
			HeaderImg:   "https:///qmplusimg.henrongyi.top/1572075907logo.png",
			AuthorityId: 9528,
			Phone:       "17611111111",
			Email:       "333333333@qq.com"},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, system2.SysUser{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	authorityEntities, ok := ctx.Value(initAuthority{}.InitializerName()).([]system2.SysAuthority)
	if !ok {
		return next, errors.Wrap(system.ErrMissingDependentContext, "创建 [用户-权限] 关联失败, 未找到权限表初始化数据")
	}
	if err = db.Model(&entities[0]).Association("Authorities").Replace(authorityEntities); err != nil {
		return next, err
	}
	if err = db.Model(&entities[1]).Association("Authorities").Replace(authorityEntities[:1]); err != nil {
		return next, err
	}
	return next, err
}

func (i *initUser) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var record system2.SysUser
	if errors.Is(db.Where("username = ?", "a303176530").
		Preload("Authorities").First(&record).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return len(record.Authorities) > 0 && record.Authorities[0].AuthorityId == 888
}
