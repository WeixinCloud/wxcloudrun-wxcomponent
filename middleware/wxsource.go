package middleware

import (
	"github.com/gin-gonic/gin"
)

// WXSourceMiddleWare 中间件 判断是否来源于微信
func WXSourceMiddleWare(c *gin.Context) {
	c.Next()
	/* if _, ok := c.Request.Header[http.CanonicalHeaderKey("x-wx-source")]; ok {
		fmt.Println("[WXSourceMiddleWare]from wx")
		c.Next()
	} else {
		c.Abort()
		c.JSON(http.StatusUnauthorized, errno.ErrNotAuthorized)
	} */
}
