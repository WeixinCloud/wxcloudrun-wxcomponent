package admin

import (
	"net/http"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/api/proxy"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"
	"github.com/gin-gonic/gin"
)

func getProxyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, errno.OK.WithData(proxy.GetProxyConfig()))
}

type updateProxyReq struct {
	Open bool `json:"open"`
	Port int  `json:"port"`
}

func updateProxyHandler(c *gin.Context) {
	var req updateProxyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData(err.Error()))
		return
	}
	if err := proxy.SetProxyConfig(req.Open, req.Port, ""); err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK)
}
