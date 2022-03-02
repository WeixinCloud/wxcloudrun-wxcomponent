package routers

import (
	"net/http"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/api/admin"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/api/authpage"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/api/innerservice"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/api/proxy"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/api/wxcallback"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/middleware"
	"github.com/gin-gonic/gin"
)

type Option func(*gin.RouterGroup)

var options []Option

// Include 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// Init 初始化
func Init() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LogMiddleWare)

	// 微信消息推送
	wxcallback.Routers(r)

	// 微管家
	Include(admin.Routers, authpage.Routers)
	g := r.Group("/wxcomponent")
	for _, opt := range options {
		opt(g)
	}

	// 静态文件
	g.Static("/assets", "client/dist/wxcomponent/assets")
	r.LoadHTMLGlob("client/dist/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.NoRoute(proxy.ProxyHandler)
	return r
}

// InnerServiceInit 内部服务初始化
func InnerServiceInit() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LogMiddleWare)
	innerservice.Routers(r)
	return r
}
