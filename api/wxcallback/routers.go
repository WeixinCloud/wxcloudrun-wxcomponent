package wxcallback

import (
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/middleware"
	"github.com/gin-gonic/gin"
)

// Routers 路由
func Routers(e *gin.Engine) {
	g := e.Group("/wxcallback", middleware.WXSourceMiddleWare)
	g.POST("/component", componentHandler)
	g.POST("/biz/:appid", bizHandler)
}
