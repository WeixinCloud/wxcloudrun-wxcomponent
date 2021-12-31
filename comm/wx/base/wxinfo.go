package wxbase

import (
	"os"
	"strings"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/dao"
)

var appid string
var envid string
var service string

func init() {
	envid = os.Getenv("CBR_ENV_ID")
	host := os.Getenv("HOSTNAME")
	if i := strings.Index(envid, "-"); i != -1 {
		appid = envid[:i]
	}
	log.Info("appid: " + appid)
	if i := strings.Index(host, "-"); i != -1 {
		service = host[:i]
	}
}

// GetService 获取当前服务
func GetService() string {
	return service
}

// GetEnvId 获取当前环境id
func GetEnvId() string {
	return envid
}

// GetAppid 获取Appid
func GetAppid() string {
	return appid
}

// GetSecret 获取Secret
func GetSecret() string {
	return dao.GetByStrDecrypt("secret", "")
}

// SetSecret 更新Secret
func SetSecret(s string) error {
	return dao.SetByStrEncrypt("secret", s)
}

// GetTicket 获取最新ticket
func GetTicket() string {
	return dao.GetByStr("ticket", "")
}

// GetTicket 更新ticket
func SetTicket(s string) error {
	return dao.SetByStr("ticket", s)
}
