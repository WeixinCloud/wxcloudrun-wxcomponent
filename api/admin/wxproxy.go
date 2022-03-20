package admin

import (
	"fmt"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/token/wxtoken"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getWxProxyHandler(c *gin.Context) {
	appid := c.Query("ldy_appid")
	uri := c.Query("ldy_uri")
	ttype := c.Query("ldy_ttype")
	if len(appid) == 0 || len(uri) == 0 || len(ttype) == 0 {
		c.JSON(http.StatusOK, &errno.JsonResult{Code: -1, ErrorMsg: "params error"})
		return
	}

	var resp string
	resp, err := wxtoken.WxProxyGet(appid, uri, ttype)
	if err != nil {
		c.JSON(http.StatusOK, &errno.JsonResult{Code: -2, ErrorMsg: err.Error()})
		return
	}

	c.String(http.StatusOK, resp)
}

func postWxProxyHandler(c *gin.Context) {
	appid := c.Query("ldy_appid")
	uri := c.Query("ldy_uri")
	ttype := c.Query("ldy_ttype")
	if len(appid) == 0 || len(uri) == 0 || len(ttype) == 0 {
		c.JSON(http.StatusOK, &errno.JsonResult{Code: -1, ErrorMsg: "params error"})
		return
	}

	var req interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("post body error: %s", err.Error())
		c.JSON(http.StatusOK, &errno.JsonResult{Code: -1, ErrorMsg: msg})
		return
	}

	var resp string
	resp, err := wxtoken.WxProxyPost(appid, uri, ttype, req)
	if err != nil {
		c.JSON(http.StatusOK, &errno.JsonResult{Code: -2, ErrorMsg: err.Error()})
		return
	}

	c.String(http.StatusOK, resp)
}
