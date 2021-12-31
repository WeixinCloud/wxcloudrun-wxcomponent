package token

import (
	"encoding/json"
	"errors"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/httputils"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	wxbase "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/base"
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

// GetNewComponentAccessToken 获取新的ComponentAccessToken
func GetNewComponentAccessToken() (string, error) {
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
