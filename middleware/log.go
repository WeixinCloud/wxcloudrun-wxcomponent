package middleware

import (
	"bytes"
	"io/ioutil"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"

	"github.com/gin-gonic/gin"
)

// 日志中间件
func LogMiddleWare(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	log.Debugf("%s", string(body))
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	log.Debugf("%s", "---header---")
	for k, v := range c.Request.Header {
		log.Debugf("%s %s", k, v)
	}
	log.Debugf("%s", "---header---")
}
