package inits

import (
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/api/admin"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/config"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db"
)

type AppOption func() error

var appOpts []AppOption

func include(opts ...AppOption) {
	appOpts = append(appOpts, opts...)
}

// 初始化
func Init() error {

	// db.Init must be the first
	include(config.Init, db.Init, admin.Init)

	for _, opt := range appOpts {
		if err := opt(); err != nil {
			log.Errorf("inits failed, err:%v\n", err)
			return err
		} else {
			log.Infof("%v init succ", opt)
		}
	}
	return nil
}
