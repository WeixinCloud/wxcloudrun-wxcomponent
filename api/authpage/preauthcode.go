package authpage

import (
	"encoding/json"
	"net/http"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/httputils"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	wxbase "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/base"
	"github.com/gin-gonic/gin"
)

type getPreAuthCodeReq struct {
	ComponentAppid string `json:"component_appid"`
}

type getPreAuthCodeResp struct {
	PreAuthCode string `json:"pre_auth_code"`
}

func getPreAuthCodeHandler(c *gin.Context) {
	req := getPreAuthCodeReq{
		ComponentAppid: wxbase.GetAppid(),
	}
	_, respbody, err := httputils.PostWxJson("/cgi-bin/component/api_create_preauthcode", req, true)
	if err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	var resp getPreAuthCodeResp
	if err := json.Unmarshal(respbody, &resp); err != nil {
		log.Errorf("Unmarshal err, %v", err)
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{
		"preAuthCode": resp.PreAuthCode,
	}))
}
