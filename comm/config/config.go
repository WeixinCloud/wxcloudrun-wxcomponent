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

var ServerConf = &Server{}

var cfg *ini.File

// Init 初始化
func Init() error {
	var err error
	cfg, err = ini.Load("comm/config/server.conf")
	if err != nil {
		log.Errorf("load server.conf': %v", err)
		return err
	}
	mapTo("server", ServerConf)
	if ServerConf.AesKey == "" {
		ServerConf.AesKey = encrypt.GenerateMd5(os.Getenv("MYSQL_PASSWORD"))
	}
	if ServerConf.JwtSecret == "" {
		ServerConf.JwtSecret = encrypt.GenerateMd5(os.Getenv("MYSQL_PASSWORD"))
	}
	log.Info(ServerConf)
	return nil
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Errorf("%s err: %v", section, err)
	}
}
