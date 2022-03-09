package innerservice

import (
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/middleware"
	"github.com/gin-gonic/gin"
)

// Routers 路由
func Routers(e *gin.Engine) {
	g := e.Group("/inner", middleware.InnerServiceMiddleWare)
	g.GET("/component-access-token", GetComponentAccessTokenHandler)
	g.GET("/authorizer-access-token", GetAuthorizerAccessTokenHandler)
	g.GET("/ticket", GetTicketHandler)
}
