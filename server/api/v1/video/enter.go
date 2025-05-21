package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	XkVideoCategoryApi
}

var (
	xkcategoryService = service.ServiceGroupApp.XkVideoServiceGroup.XkVideoCategoryService
)
