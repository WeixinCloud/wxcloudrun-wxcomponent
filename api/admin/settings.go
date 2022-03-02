package admin

import (
	"encoding/json"
	"net/http"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"

	wxbase "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/base"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/dao"
	"github.com/gin-gonic/gin"
)

type secretReq struct {
	Secret string `json:"secret"`
}

func setWxSecretHandler(c *gin.Context) {
	var req secretReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData(err.Error()))
		return
	}
	if err := wxbase.SetSecret(req.Secret); err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK)
}

func getWxSecretHandler(c *gin.Context) {
	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{"secret": wxbase.GetSecret()}))
}

type componentInfo struct {
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	RedirectUrl string `json:"redirectUrl"`
}

func setComponentInfoHandler(c *gin.Context) {
	var req componentInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData(err.Error()))
		return
	}
	value, _ := json.Marshal(req)
	if err := dao.SetCommKv("authinfo", string(value)); err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK)
}
