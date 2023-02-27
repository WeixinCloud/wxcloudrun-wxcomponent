package wxcallback

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx"
	"github.com/subosito/gotenv"

	wxbase "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/base"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/dao"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
	"github.com/gin-gonic/gin"
)

var appid string
var token string
var aesKey string

func init() {
	gotenv.Load()
	aesKey = os.Getenv("WX_AESKEY")
	token = os.Getenv("WX_TOKEN")
	appid = os.Getenv("WX_APPID")
}

type wxCallbackComponentRecord struct {
	CreateTime int64  `json:"CreateTime"`
	InfoType   string `json:"InfoType"`
}

type EncryptRequestBody struct {
	XMLName                      xml.Name `xml:"xml"`
	AppId                        string   `xml:"AppId" json:"app_id"`
	CreateTime                   int64    `xml:"CreateTime" json:"create_time"`
	InfoType                     string   `xml:"InfoType" json:"info_type"`
	ComponentVerifyTicket        string   `xml:"ComponentVerifyTicket" json:"component_verify_ticket"`
	AuthorizerAppid              string   `xml:"AuthorizerAppid" json:"authorizer_appid"`
	AuthorizationCode            string   `xml:"AuthorizationCode" json:"authorization_code"`
	AuthorizationCodeExpiredTime string   `xml:"AuthorizationCodeExpiredTime" json:"authorization_code_expired_time"`
	PreAuthCode                  string   `xml:"PreAuthCode" json:"pre_auth_code"`
}

func ParseEncryptRequestBody(r *http.Request) *EncryptRequestBody {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil
	}
	//  mlog.AppendObj(nil, "Wechat Message Service: RequestBody--", body)
	requestBody := &EncryptRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody

}

func componentHandler(c *gin.Context) {
	// 记录到数据库
	body, _ := ioutil.ReadAll(c.Request.Body)
	wxReq := WxReq{}
	if err := c.ShouldBindQuery(&wxReq); err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
	}
	instance := NewWechatMsgCrypt(token, aesKey, appid)

	bytes := []byte(body)
	s := strings.TrimSpace(string(bytes))
	xmlData := EventEncryptRequest{}
	_ = xml.Unmarshal([]byte(s), &xmlData)
	encrypttBody := instance.WechatEventDecrypt(wxReq, xmlData)
	log.Info(encrypttBody)
	r := model.WxCallbackComponentRecord{
		CreateTime:  time.Unix(int64(encrypttBody.CreateTime), 0),
		ReceiveTime: time.Now(),
		InfoType:    encrypttBody.InfoType,
		PostBody:    string(body),
	}
	if encrypttBody.CreateTime == 0 {
		r.CreateTime = time.Unix(1, 0)
	}
	if err := dao.AddComponentCallBackRecord(&r); err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}

	// 处理授权相关的消息
	var err error
	switch encrypttBody.InfoType {
	case "component_verify_ticket":
		err = ticketHandler(encrypttBody)
	case "authorized":
		fallthrough
	case "updateauthorized":
		err = newAuthHander(encrypttBody)
	case "unauthorized":
		err = unAuthHander(encrypttBody)
	}
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}

	// 转发到用户配置的地址
	var proxyOpen bool
	proxyOpen, err = proxyCallbackMsg(encrypttBody.InfoType, "", "", string(body), c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	if !proxyOpen {
		c.String(http.StatusOK, "success")
	}
}

type ticketRecord struct {
	ComponentVerifyTicket string `json:"ComponentVerifyTicket"`
}

func ticketHandler(encrypttBody EventMessageBody) error {
	record := ticketRecord{ComponentVerifyTicket: encrypttBody.ComponentVerifyTicket}
	log.Info("[new ticket]" + record.ComponentVerifyTicket)
	if err := wxbase.SetTicket(record.ComponentVerifyTicket); err != nil {
		return err
	}
	return nil
}

type newAuthRecord struct {
	CreateTime                   int64  `json:"CreateTime"`
	AuthorizerAppid              string `json:"AuthorizerAppid"`
	AuthorizationCode            string `json:"AuthorizationCode"`
	AuthorizationCodeExpiredTime string `json:"AuthorizationCodeExpiredTime"`
}

func newAuthHander(encrypttBody EventMessageBody) error {
	var err error
	var refreshtoken string
	var appinfo wx.AuthorizerInfoResp
	record := newAuthRecord{
		CreateTime:                   int64(encrypttBody.CreateTime),
		AuthorizerAppid:              encrypttBody.AuthorizerAppid,
		AuthorizationCode:            encrypttBody.AuthorizationCode,
		AuthorizationCodeExpiredTime: encrypttBody.AuthorizationCodeExpiredTime,
	}
	if refreshtoken, err = queryAuth(record.AuthorizationCode); err != nil {
		return err
	}
	if err = wx.GetAuthorizerInfo(record.AuthorizerAppid, &appinfo); err != nil {
		return err
	}
	if err = dao.CreateOrUpdateAuthorizerRecord(&model.Authorizer{
		Appid:         record.AuthorizerAppid,
		AppType:       appinfo.AuthorizerInfo.AppType,
		ServiceType:   appinfo.AuthorizerInfo.ServiceType.Id,
		NickName:      appinfo.AuthorizerInfo.NickName,
		UserName:      appinfo.AuthorizerInfo.UserName,
		HeadImg:       appinfo.AuthorizerInfo.HeadImg,
		QrcodeUrl:     appinfo.AuthorizerInfo.QrcodeUrl,
		PrincipalName: appinfo.AuthorizerInfo.PrincipalName,
		RefreshToken:  refreshtoken,
		FuncInfo:      appinfo.AuthorizationInfo.StrFuncInfo,
		VerifyInfo:    appinfo.AuthorizerInfo.VerifyInfo.Id,
		AuthTime:      time.Unix(record.CreateTime, 0),
	}); err != nil {
		return err
	}
	return nil
}

type queryAuthReq struct {
	ComponentAppid    string `wx:"component_appid"`
	AuthorizationCode string `wx:"authorization_code"`
}

type authorizationInfo struct {
	AuthorizerRefreshToken string `wx:"authorizer_refresh_token"`
}
type queryAuthResp struct {
	AuthorizationInfo authorizationInfo `wx:"authorization_info"`
}

func queryAuth(authCode string) (string, error) {
	req := queryAuthReq{
		ComponentAppid:    wxbase.GetAppid(),
		AuthorizationCode: authCode,
	}
	var resp queryAuthResp
	_, body, err := wx.PostWxJsonWithComponentToken("/cgi-bin/component/api_query_auth", "", req)
	if err != nil {
		return "", err
	}
	if err := wx.WxJson.Unmarshal(body, &resp); err != nil {
		log.Errorf("Unmarshal err, %v", err)
		return "", err
	}
	return resp.AuthorizationInfo.AuthorizerRefreshToken, nil
}

type unAuthRecord struct {
	CreateTime      int64  `json:"CreateTime"`
	AuthorizerAppid string `json:"AuthorizerAppid"`
}

func unAuthHander(encrypttBody EventMessageBody) error {
	record := unAuthRecord{
		CreateTime:      int64(encrypttBody.CreateTime),
		AuthorizerAppid: encrypttBody.AuthorizerAppid,
	}
	if err := dao.DelAuthorizerRecord(record.AuthorizerAppid); err != nil {
		log.Errorf("DelAuthorizerRecord err %v", err)
		return err
	}
	return nil
}
