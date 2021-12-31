package routers

import (
	"net/http"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/api/admin"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/api/authpage"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/api/wxcallback"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/middleware"
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options []Option

// Include 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// Init 初始化
func Init() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LogMiddleWare)

	// 加载多个APP的路由配置
	Include(admin.Routers, wxcallback.Routers, authpage.Routers)

	for _, opt := range options {
		opt(r)
	}
	r.Static("/assets", "client/dist/assets")
	r.LoadHTMLGlob("client/dist/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	return r
}
