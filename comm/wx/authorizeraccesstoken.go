package wx

import (
	"errors"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	wxbase "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/base"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/dao"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
)

type authorizerAccessTokenReq struct {
	ComponentAppid         string `wx:"component_appid"`
	AuthorizerAppid        string `wx:"authorizer_appid"`
	AuthorizerRefreshToken string `wx:"authorizer_refresh_token"`
}

type authorizerAccessTokenResp struct {
	AuthorizerAccessToken string `wx:"authorizer_access_token"`
	ExpiresIn             uint64 `wx:"expires_in"`
}

// GetAuthorizerAccessToken 获取AuthorizerAccessToken
func GetAuthorizerAccessToken(appid string) (string, error) {
	return getAccessTokenWithRetry(appid, model.WXTOKENTYPE_AUTH)
}

func getNewAuthorizerAccessToken(appid string) (string, error) {
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
	var resp authorizerAccessTokenResp
	_, body, err := PostWxJsonWithComponentToken("/cgi-bin/component/api_authorizer_token", "", req)
	if err != nil {
		return "", err
	}
	if err := WxJson.Unmarshal(body, &resp); err != nil {
		log.Errorf("Unmarshal err, %v", err)
		return "", err
	}
	return resp.AuthorizerAccessToken, nil
}
