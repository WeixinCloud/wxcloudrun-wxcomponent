package wxbase

import (
	"os"
	"strings"
	"time"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/dao"
)

var appid string
var envid string
var service string

func init() {
	envid = os.Getenv("CBR_ENV_ID")
	host := os.Getenv("HOSTNAME")
	appid = os.Getenv("WX_APPID")
	if len(appid) == 0 {
		if i := strings.Index(envid, "-"); i != -1 {
			appid = envid[:i]
		}
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
	return dao.GetCommKvDecrypt("secret", "")
}

// SetSecret 更新Secret
func SetSecret(s string) error {
	return dao.SetCommKvEncrypt("secret", s)
}

// GetTicket 获取最新ticket
func GetTicket() string {
	return dao.GetCommKvWithCache("ticket", "", 15*time.Minute)
}

// GetTicket 更新ticket
func SetTicket(s string) error {
	return dao.SetCommKvWithCache("ticket", s, 15*time.Minute)
}
