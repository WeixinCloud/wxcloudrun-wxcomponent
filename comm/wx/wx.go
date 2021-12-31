package wx

import (
	"encoding/json"
	"fmt"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/httputils"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	wxbase "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/base"
)

type getAuthorizerInfoReq struct {
	ComponentAppid  string `json:"component_appid"`
	AuthorizerAppid string `json:"authorizer_appid"`
}

type idItem struct {
	Id int `json:id`
}

type funcInfo struct {
	FuncscopeCategory idItem `json:"funcscope_category"`
}

type authorizationInfo struct {
	AuthorizationAppid string     `json:"authorization_appid"`
	RawFuncInfo        []funcInfo `json:"func_info"`
	StrFuncInfo        string
}

type networkInfo struct {
	RequestDomain   []string `json:RequestDomain`
	WsRequestDomain []string `json:WsRequestDomain`
	UploadDomain    []string `json:UploadDomain`
	DownloadDomain  []string `json:DownloadDomain`
	BizDomain       []string `json:BizDomain`
	UDPDomain       []string `json:UDPDomain`
}

type categorieInfo struct {
	First  string `json:first`
	Second string `json:second`
}

type miniProgramInfo struct {
	Network    networkInfo     `json:network`
	Categories []categorieInfo `json:categories`
}

type authorizerInfo struct {
	NickName        string           `json:"nick_name"`
	HeadImg         string           `json:"head_img"`
	Appid           string           `json:"appid"`
	ServiceType     idItem           `json:"service_type_info"`
	VerifyInfo      idItem           `json:"verify_type_info"`
	UserName        string           `json:"user_name"`
	PrincipalName   string           `json:"principal_name"`
	QrcodeUrl       string           `json:"qrcode_url"`
	MiniProgramInfo *miniProgramInfo `json:"MiniProgramInfo"`
	AppType         int
}

// AuthorizerInfoResp 授权账号信息结构体
type AuthorizerInfoResp struct {
	AuthorizationInfo authorizationInfo `json:"authorization_info"`
	AuthorizerInfo    authorizerInfo    `json:"authorizer_info"`
}

// GetAuthorizerInfo 获取授权账号信息
func GetAuthorizerInfo(appid string, resp *AuthorizerInfoResp) error {
	req := getAuthorizerInfoReq{
		ComponentAppid:  wxbase.GetAppid(),
		AuthorizerAppid: appid,
	}
	_, respbody, err := httputils.PostWxJson("/cgi-bin/component/api_get_authorizer_info", req, true)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(respbody, &resp); err != nil {
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
