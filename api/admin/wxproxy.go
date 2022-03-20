package admin

import (
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/utils"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/token/wxtoken"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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

	jwt, _ := c.Get("jwt")
	claims := jwt.(*utils.Claims)

	var resp string
	resp, err := wxtoken.WxProxyGet(appid, uri, ttype, claims)
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

	jwt, _ := c.Get("jwt")
	claims := jwt.(*utils.Claims)
	buffer, _ := ioutil.ReadAll(c.Request.Body)

	var resp string
	resp, err := wxtoken.WxProxyPost(appid, uri, ttype, buffer, claims)
	if err != nil {
		c.JSON(http.StatusOK, &errno.JsonResult{Code: -2, ErrorMsg: err.Error()})
		return
	}

	c.String(http.StatusOK, resp)
}
