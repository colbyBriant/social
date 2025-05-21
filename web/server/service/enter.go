package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/video"
)

// 实例创建聚合
type ServicesGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	XkBbsServiceGroup   bbs.ServiceGroup
	XkVideoServiceGroup video.ServiceGroup
}

// 单例设计模式
var ServiceGroupApp = new(ServicesGroup)
