package wxtoken

import (
	"encoding/json"
	"errors"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/httputils"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	wxbase "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/base"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
)

type componentAccessTokenReq struct {
	ComponentAppid        string `json:"component_appid"`
	ComponentAppSecret    string `json:"component_appsecret"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
}

type componentAccessTokenResp struct {
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            uint64 `json:"expires_in"`
}

// GetComponentAccessToken 获取ComponentAccessToken
func GetComponentAccessToken() (string, error) {
	return getAccessToken(wxbase.GetAppid(), model.WXTOKENTYPE_OWN)
}

func getNewComponentAccessToken() (string, error) {
	ticket := wxbase.GetTicket()
	if len(ticket) == 0 {
		return "", errors.New("empty ticket")
	}
	req := componentAccessTokenReq{
		ComponentAppid:        wxbase.GetAppid(),
		ComponentAppSecret:    wxbase.GetSecret(),
		ComponentVerifyTicket: ticket,
	}
	_, respbody, err := httputils.PostWxJson("/cgi-bin/component/api_component_token", req, false)
	if err != nil {
		return "", err
	}
	var resp componentAccessTokenResp
	if err := json.Unmarshal(respbody, &resp); err != nil {
		log.Errorf("Unmarshal err, %v", err)
		return "", err
	}
	return resp.ComponentAccessToken, nil
}
