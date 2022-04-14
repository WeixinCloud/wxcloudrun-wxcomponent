package admin

import (
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/api/innerservice"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/middleware"
	"github.com/gin-gonic/gin"
)

// Routers 路由
func Routers(e *gin.RouterGroup) {
	// auth
	e.PUT("/auth", authHandler)

	g := e.Group("/admin", middleware.JWTMiddleWare)

	// 第三方token
	g.GET("/cloudbase-access-token", getCloudbaseAccessTokenHandler)
	g.GET("/component-access-token", innerservice.GetComponentAccessTokenHandler)
	g.GET("/authorizer-access-token", innerservice.GetAuthorizerAccessTokenHandler)
	g.GET("/ticket", innerservice.GetTicketHandler)

	// 消息与事件
	g.GET("/wx-component-records", getWxComponentRecordsHandler)
	g.GET("/wx-biz-records", getWxBizRecordsHandler)
	g.GET("/callback-config", getWxCallBackConfigHandler)
	g.GET("/callback-proxy-rule-list", getCallBackProxyRuleListHandler)
	g.POST("/callback-proxy-rule", updateCallBackProxyRuleHandler)
	g.PUT("/callback-proxy-rule", addCallBackProxyRuleHandler)
	g.DELETE("/callback-proxy-rule", delCallBackProxyRuleHandler)
	g.POST("/callback-test", testCallbackRuleHandler)

	// 授权小程序管理
	g.POST("/pull-authorizer-list", pullAuthorizerListHandler)
	g.GET("/authorizer-list", getAuthorizerListHandler)

	// 代开发小程序管理
	g.GET("/dev-weapp-list", getDevWeAppListHandler)
	g.POST("/submit-audit", submitAuditHandler)
	g.GET("/dev-versions", devVersionsHandler)
	g.GET("/template-list", templateListHandler)
	g.POST("/revoke-audit", revokeAuditHandler)
	g.POST("/speed-up-audit", speedUpAuditHandler)
	g.POST("/commit-code", commitCodeHandler)
	g.POST("/release-code", releaseCodeHandler)
	g.POST("/upload-media", uploadMediaHandler)
	g.POST("/change-visit-status", changeVisitStatusHandler)
	g.POST("/rollback-release-version", rollbackReleaseVersionHandler)
	g.GET("/page-list", getPageListHandler)
	g.GET("/category", getCategoryHandler)
	g.GET("/qrcode", getQRCodeHandler)

	// 设置
	g.POST("/secret", setWxSecretHandler)
	g.GET("/secret", getWxSecretHandler)
	g.POST("/componentinfo", setComponentInfoHandler)

	// 用户管理
	g.POST("/username", updateUserNameHandler)
	g.POST("/userpwd", updateUserPwdHandler)

	// 刷新token
	g.GET("/refresh-auth", refreshAuthHandler)

	// 转发设置
	g.GET("/proxy", getProxyHandler)
	g.POST("/proxy", updateProxyHandler)
}
