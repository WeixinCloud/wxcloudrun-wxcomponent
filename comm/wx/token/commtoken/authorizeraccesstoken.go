package token

import (
	"encoding/json"
	"errors"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/httputils"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	wxbase "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/base"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/dao"
)

type authorizerAccessTokenReq struct {
	ComponentAppid         string `json:"component_appid"`
	AuthorizerAppid        string `json:"authorizer_appid"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

type authorizerAccessTokenResp struct {
	AuthorizerAccessToken string `json:"authorizer_access_token"`
	ExpiresIn             uint64 `json:"expires_in"`
}

// GetNewAuthorizerAccessToken 获取新AuthorizerAccessToken
func GetNewAuthorizerAccessToken(appid string) (string, error) {
	records, _, err := dao.GetAuthorizerRecords(appid, 0, 1)
	if err != nil {
		return "", err
	}
	if len(records) < 1 {
		return "", errors.New("empty records")
	}
	req := authorizerAccessTokenReq{
		ComponentAppid:         wxbase.GetAppid(),
		AuthorizerAppid:        appid,
		AuthorizerRefreshToken: records[0].RefreshToken,
	}
	_, respbody, err := httputils.PostWxJson("/cgi-bin/component/api_authorizer_token", req, true)
	if err != nil {
		return "", err
	}
	var resp authorizerAccessTokenResp
	if err := json.Unmarshal(respbody, &resp); err != nil {
		log.Errorf("Unmarshal err, %v", err)
		return "", err
	}
	return resp.AuthorizerAccessToken, nil
}
