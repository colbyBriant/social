package video

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type XkCategoryRouter struct{}

func (e *XkCategoryRouter) InitXkcategoryRouter(Router *gin.RouterGroup) {

	xkVideoCategoryApi := v1.ApiGroupApp.VideoApiGroup.XkVideoCategoryApi
	// 这个路由多了一个对对post，put请求的中间件处理，而这个中间件做了一些对post和put的参数的处理和一些公共信息的处理
	xkBbsCustomerRouterWithoutRecord := Router.Group("videocategory") //.Use(middleware.OperationRecord())
	{
		// 保存
		xkBbsCustomerRouterWithoutRecord.POST("save", xkVideoCategoryApi.CreateXkVideoCategory)
		// 更新
		xkBbsCustomerRouterWithoutRecord.POST("update", xkVideoCategoryApi.UpdateXkVideoCategory)
		// 更新状态
		xkBbsCustomerRouterWithoutRecord.POST("update/status", xkVideoCategoryApi.UpdateXkVideoCategoryStatus)
		// 更新
		xkBbsCustomerRouterWithoutRecord.DELETE("delete/:id", xkVideoCategoryApi.DeleteById)
	}

	// 这个路由是没有中间件的路由
	xkBbsRouterWithoutRecord := Router.Group("videocategory")
	{
		// 分页查询
		xkBbsRouterWithoutRecord.GET("page", xkVideoCategoryApi.LoadXkVideoCategoryPage)
		// 明细查询
		xkBbsRouterWithoutRecord.GET("get", xkVideoCategoryApi.GetXkVideoCategory)
		// 明细查询
		xkBbsRouterWithoutRecord.GET("find", xkVideoCategoryApi.FindCategories)
	}
}
