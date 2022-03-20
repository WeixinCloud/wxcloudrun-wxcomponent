package httputils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/token/cloudbasetoken"
)

type wxCommError struct {
	ErrCode int    `json:errcode default:0`
	ErrMsg  string `json:errmsg`
}

// GetWxApiUrl 拼接微信开放平台的url
func GetWxApiUrl(path string) string {
	return fmt.Sprintf("https://api.weixin.qq.com%s?cloudbase_access_token=%s", path, cloudbasetoken.GetCloudBaseAccessToken())
}

// GetRawWxApiUrl 拼接微信开放平台的url，不带微信令牌
func GetRawWxApiUrl(path string) string {
	return fmt.Sprintf("https://api.weixin.qq.com%s", path)
}

// PostWxJson 向微信开放平台发起post请求
func PostWxJson(path string, data interface{}, withCloudToken bool) (wxCommError, []byte, error) {
	var url string
	var wxerror wxCommError
	var body []byte
	var err error
	if withCloudToken {
		url = GetWxApiUrl(path)
	} else {
		url = GetRawWxApiUrl(path)
	}
	if body, err = PostJson(url, data); err != nil {
		return wxerror, body, err
	}
	if err = json.Unmarshal(body, &wxerror); err != nil {
		log.Errorf("Unmarshal err, %v", err)
	}
	if wxerror.ErrCode != 0 {
		return wxerror, body, fmt.Errorf("WxErrCode != 0, resp: %v", wxerror)
	}
	return wxerror, body, nil
}

// PostWxJson 发起post请求，格式为JSON
func PostJson(url string, data interface{}) ([]byte, error) {
	return Post(url, data, "application/json")
}

// Post 发起post请求
func Post(url string, data interface{}, contentType string) ([]byte, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	log.Debugf("http url: %s", url)
	log.Debugf("http req: %s", jsonStr)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
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

func Get(url string) ([]byte, error) {
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
