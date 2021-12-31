package admin

import (
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/middleware"
	"github.com/gin-gonic/gin"
)

// Routers 路由
func Routers(e *gin.Engine) {
	// auth
	e.PUT("/auth", authHandler)

	g := e.Group("/admin", middleware.JWTMiddleWare)

	// 第三方token
	g.GET("/component-access-token", componentAccessTokenHandler)
	g.GET("/cloudbase-access-token", cloudbaseAccessTokenHandler)
	g.GET("/authorizer-access-token", authorizerAccessTokenHandler)
	g.GET("/ticket", ticketHandler)

	// 消息与事件
	g.GET("/wx-component-records", wxComponentRecordsHandler)
	g.GET("/wx-biz-records", wxBizRecordsHandler)
	g.GET("/callback-config", wxCallBackConfigHandler)

	// 小程序管理
	g.POST("/pull-authorizer-list", pullAuthorizerListHandler)
	g.GET("/authorizer-list", getAuthorizerListHandler)

	// 设置
	g.POST("/secret", setWxSecretHandler)
	g.GET("/secret", getWxSecretHandler)
	g.POST("/componentinfo", setComponentInfoHandler)

	// 用户管理
	g.POST("/username", updateUserNameHandler)
	g.POST("/userpwd", updateUserPwdHandler)

	// 刷新token
	g.GET("/refresh-auth", refreshAuthHandler)
}
