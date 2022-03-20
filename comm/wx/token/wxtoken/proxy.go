package wxtoken

import (
	"bytes"
	"fmt"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/httputils"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/utils"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/dao"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func WxProxyGet(appid string, uri string, ttype string, user *utils.Claims) (result string, err error) {
	// 获取token及url
	url, err := getUrl(appid, uri, ttype)
	if err != nil {
		return
	}

	// 请求微信接口
	resp, err := proxyGet(url)
	if err != nil {
		return
	}

	// 记录日志
	err = dao.AddRecord("get", uri, appid, []byte(""), resp, user)
	if err != nil {
		return
	}

	result = string(resp)
	return
}

func WxProxyPost(appid string, uri string, ttype string, data []byte, user *utils.Claims) (result string, err error) {
	// 获取token及url
	url, err := getUrl(appid, uri, ttype)
	if err != nil {
		return
	}

	// 请求微信接口
	resp, err := proxyPost(url, data)
	if err != nil {
		return
	}

	// 记录日志
	err = dao.AddRecord("get", uri, appid, []byte(""), resp, user)
	if err != nil {
		return
	}

	result = string(resp)
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

func proxyGet(url string) ([]byte, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	log.Debugf("http get url: %s", url)
	resp, err := client.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("http code: %d", resp.StatusCode)
	}

	result, _ := ioutil.ReadAll(resp.Body)
	log.Debugf("http get resp: %s", result)
	return result, nil
}

func proxyPost(url string, data []byte) ([]byte, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	contentType := "application/json"
	log.Debugf("http url: %s", url)
	log.Debugf("http req: %s", data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("http code: %d", resp.StatusCode)
	}

	result, _ := ioutil.ReadAll(resp.Body)
	log.Debugf("http resp: %s", result)
	return result, nil
}