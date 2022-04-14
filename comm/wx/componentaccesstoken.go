package wx

import (
	"errors"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	wxbase "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/base"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
)

type componentAccessTokenReq struct {
	ComponentAppid        string `wx:"component_appid"`
	ComponentAppSecret    string `wx:"component_appsecret"`
	ComponentVerifyTicket string `wx:"component_verify_ticket"`
}

type componentAccessTokenResp struct {
	ComponentAccessToken string `wx:"component_access_token"`
	ExpiresIn            uint64 `wx:"expires_in"`
}

// GetComponentAccessToken 获取ComponentAccessToken
func GetComponentAccessToken() (string, error) {
	return getAccessTokenWithRetry(wxbase.GetAppid(), model.WXTOKENTYPE_OWN)
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
	var resp componentAccessTokenResp
	_, body, err := PostWxJsonWithoutToken("/cgi-bin/component/api_component_token", "", req)
	if err != nil {
		return "", err
	}
	if err := WxJson.Unmarshal(body, &resp); err != nil {
		log.Errorf("Unmarshal err, %v", err)
		return "", err
	}
	return resp.ComponentAccessToken, nil
}
