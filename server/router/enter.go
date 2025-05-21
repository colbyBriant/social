package router

import (
	bbsrouter "github.com/flipped-aurora/gin-vue-admin/server/router/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/upload"
	"github.com/flipped-aurora/gin-vue-admin/server/router/video"
)

type RouterGroup struct {
	System          system.RouterGroup
	Example         example.RouterGroup
	XkBbs           bbsrouter.RouterGroup
	LocalUpload     upload.RouterGroup
	XKVideoCategory video.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
