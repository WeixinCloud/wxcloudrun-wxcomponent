package middleware

import (
	"bytes"
	"io/ioutil"
	"strings"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"

	"github.com/gin-gonic/gin"
)

// 日志中间件
func LogMiddleWare(c *gin.Context) {
	if len(c.Request.Header["Content-Type"]) == 0 ||
		(len(c.Request.Header["Content-Type"]) > 0 &&
			strings.Contains(strings.ToLower(c.Request.Header["Content-Type"][0]), "application/json")) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		log.Debugf("body: %s", string(body))
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}

	log.Debugf("%s", "---header---")
	for k, v := range c.Request.Header {
		log.Debugf("%s %s", k, v)
	}
	log.Debugf("%s", "---header---")
}
