package config

import (
	"os"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/encrypt"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/go-ini/ini"
)

// Server 系统配置结构体
type Server struct {
	JwtSecret     string
	JwtIssue      string
	JwtExpireTime int32
	AesKey        string
}

// WxApi 配置结构体
type WxApi struct {
	UseCloudBaseAccessToken bool
	UseComponentAccessToken bool
	UseHttps                bool
}

// Comm 常规配置结构体
type Comm struct {
	Version string
}

var ServerConf = &Server{}
var CommConf = &Comm{}
var WxApiConf = &WxApi{}

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("comm/config/server.conf")
	if err != nil {
		log.Errorf("load server.conf': %v", err)
		return
	}
	mapTo("server", ServerConf)
	mapTo("comm", CommConf)
	mapTo("wxapi", WxApiConf)
	if ServerConf.AesKey == "" {
		ServerConf.AesKey = encrypt.GenerateMd5(os.Getenv("MYSQL_PASSWORD"))
	}
	if ServerConf.JwtSecret == "" {
		ServerConf.JwtSecret = encrypt.GenerateMd5(os.Getenv("MYSQL_PASSWORD"))
	}
	log.Info(ServerConf)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Errorf("%s err: %v", section, err)
	}
}
