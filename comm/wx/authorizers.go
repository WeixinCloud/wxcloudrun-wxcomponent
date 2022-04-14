package wx

import (
	"fmt"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	wxbase "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/base"
)

type getAuthorizerInfoReq struct {
	ComponentAppid  string `wx:"component_appid"`
	AuthorizerAppid string `wx:"authorizer_appid"`
}

type idItem struct {
	Id int `wx:"id"`
}

type funcInfo struct {
	FuncscopeCategory idItem `wx:"funcscope_category"`
}

type authorizationInfo struct {
	AuthorizationAppid string     `wx:"authorization_appid"`
	RawFuncInfo        []funcInfo `wx:"func_info"`
	StrFuncInfo        string
}

type networkInfo struct {
	RequestDomain   []string `wx:"RequestDomain"`
	WsRequestDomain []string `wx:"WsRequestDomain"`
	UploadDomain    []string `wx:"UploadDomain"`
	DownloadDomain  []string `wx:"DownloadDomain"`
	BizDomain       []string `wx:"BizDomain"`
	UDPDomain       []string `wx:"UDPDomain"`
}

type categorieInfo struct {
	First  string `wx:"first"`
	Second string `wx:"second"`
}

type miniProgramInfo struct {
	Network    networkInfo     `wx:"network"`
	Categories []categorieInfo `wx:"categories"`
}

// AuthorizerBasicConfig 授权账号的基础配置结构体
type AuthorizerBasicConfig struct {
	IsPhoneConfigured bool `json:"isPhoneConfigured" wx:"is_phone_configured"`
	IsEmailConfigured bool `json:"isEmailConfigured" wx:"is_email_configured"`
}
type authorizerInfo struct {
	NickName        string                 `wx:"nick_name"`
	HeadImg         string                 `wx:"head_img"`
	Appid           string                 `wx:"appid"`
	ServiceType     idItem                 `wx:"service_type_info"`
	VerifyInfo      idItem                 `wx:"verify_type_info"`
	UserName        string                 `wx:"user_name"`
	PrincipalName   string                 `wx:"principal_name"`
	QrcodeUrl       string                 `wx:"qrcode_url"`
	MiniProgramInfo *miniProgramInfo       `wx:"MiniProgramInfo"`
	RegisterType    int                    `wx:"register_type"`
	AccountStatus   int                    `wx:"account_status"`
	BasicConfig     *AuthorizerBasicConfig `wx:"basic_config"`
	AppType         int
}

// AuthorizerInfoResp 授权账号信息结构体
type AuthorizerInfoResp struct {
	AuthorizationInfo authorizationInfo `wx:"authorization_info"`
	AuthorizerInfo    authorizerInfo    `wx:"authorizer_info"`
}

// GetAuthorizerInfo 获取授权账号信息
func GetAuthorizerInfo(appid string, resp *AuthorizerInfoResp) error {
	req := getAuthorizerInfoReq{
		ComponentAppid:  wxbase.GetAppid(),
		AuthorizerAppid: appid,
	}
	_, body, err := PostWxJsonWithComponentToken("/cgi-bin/component/api_get_authorizer_info", "", req)
	if err != nil {
		log.Error(err)
		return err
	}
	if err := WxJson.Unmarshal(body, &resp); err != nil {
		log.Errorf("Unmarshal err, %v", err)
		return err
	}
	funcListStr := ""
	for _, v := range resp.AuthorizationInfo.RawFuncInfo {
		funcListStr = fmt.Sprintf("%s|%d", funcListStr, v.FuncscopeCategory.Id)
	}
	resp.AuthorizationInfo.StrFuncInfo = funcListStr
	resp.AuthorizerInfo.AppType = 0 // 小程序
	if resp.AuthorizerInfo.MiniProgramInfo == nil {
		resp.AuthorizerInfo.AppType = 1 // 公众号
	}
	return nil
}
