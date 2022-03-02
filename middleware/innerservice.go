package middleware

import (
	"net/http"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/gin-gonic/gin"
)

var IpWhitelist []string = []string{"127.0.0.1"}

// 内部服务中间件
func InnerServiceMiddleWare(c *gin.Context) {
	whitelisted := false
	clientIp := getRequestIP(c)
	log.Info("clientIp: ", clientIp)
	for _, v := range IpWhitelist {
		if v == clientIp {
			whitelisted = true
		}
	}
	if whitelisted {
		c.Next()
	} else {
		c.Abort()
		c.JSON(http.StatusUnauthorized, errno.ErrNotAuthorized)
	}
}

// 获取ip
func getRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}
