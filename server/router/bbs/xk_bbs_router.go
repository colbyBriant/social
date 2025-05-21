package bbs

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type XkBbsRouter struct{}

func (e *XkBbsRouter) InitXkBbsRouter(Router *gin.RouterGroup) {

	xkBbsApi := v1.ApiGroupApp.BbsApiGroup.XkBbsApi
	// 这个路由多了一个对对post，put请求的中间件处理，而这个中间件做了一些对post和put的参数的处理和一些公共信息的处理
	xkBbsCustomerRouterWithoutRecord := Router.Group("bbs") //.Use(middleware.OperationRecord())
	{
		// 保存
		xkBbsCustomerRouterWithoutRecord.POST("save", xkBbsApi.CreateXkBbs)
		// 更新
		xkBbsCustomerRouterWithoutRecord.POST("update", xkBbsApi.UpdateXkBbs)
		// 更新
		xkBbsCustomerRouterWithoutRecord.DELETE("delete/:id", xkBbsApi.DeleteById)
		// 分页查询
		xkBbsCustomerRouterWithoutRecord.POST("page", xkBbsApi.LoadXkBbsPage)
	}

	// 这个路由是没有中间件的路由
	xkBbsRouterWithoutRecord := Router.Group("bbs")
	{
		// 明细查询
		xkBbsRouterWithoutRecord.GET("get", xkBbsApi.GetXkBbs)
	}
}
