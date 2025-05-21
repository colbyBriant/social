package upload

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type OssUploadRouter struct{}

func (e *OssUploadRouter) InitOssUploadRouter(Router *gin.RouterGroup) {

	ossUploadApi := v1.ApiGroupApp.LocalUploadGroup.OSSUploadApi
	// 这个路由多了一个对对post，put请求的中间件处理，而这个中间件做了一些对post和put的参数的处理和一些公共信息的处理
	xkBbsCustomerRouterWithoutRecord := Router.Group("oss") //.Use(middleware.OperationRecord())
	{
		// 保存
		xkBbsCustomerRouterWithoutRecord.POST("upload/file", ossUploadApi.UploadFile)
	}

}
