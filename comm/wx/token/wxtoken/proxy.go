package wxtoken

import (
	"fmt"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/httputils"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
	"strconv"
)

func WxProxyGet(appid string, uri string, ttype string) (result string, err error) {
	url, err := getUrl(appid, uri, ttype)
	if err != nil {
		return
	}

	temp, err := httputils.Get(url)
	if err != nil {
		return
	}

	result = string(temp)
	return
}

func WxProxyPost(appid string, uri string, ttype string, data interface{}) (result string, err error) {
	url, err := getUrl(appid, uri, ttype)
	if err != nil {
		return
	}

	temp, err := httputils.PostJson(url, data)
	if err != nil {
		return
	}

	result = string(temp)
	return
}

func getUrl(appid string, uri string, ttype string) (url string, err error) {
	tTypeInt, err := strconv.Atoi(ttype)
	if err != nil {
		err = fmt.Errorf("ttype(%s) error", ttype)
		return
	}

	var token string
	var tokenKey string
	if tTypeInt == model.WXTOKENTYPE_AUTH {
		tokenKey = "access_token"
	} else if tTypeInt == model.WXTOKENTYPE_OWN {
		tokenKey = "component_access_token"
	} else {
		err = fmt.Errorf("invalid type(%s)", ttype)
		return
	}

	token, err = getAccessToken(appid, tTypeInt)
	if err != nil || len(token) == 0 {
		err = fmt.Errorf("get access_token error: appid=%s, token_key=%s", appid, tokenKey)
		return
	}

	baseUrl := httputils.GetRawWxApiUrl(uri)
	url = fmt.Sprintf("%s?%s=%s", baseUrl, tokenKey, token)
	return
}