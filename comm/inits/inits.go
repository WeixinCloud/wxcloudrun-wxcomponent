package inits

import (
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/api/admin"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/api/proxy"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/dao"
)

type AppOption func() error

var appOpts []AppOption

func include(opts ...AppOption) {
	appOpts = append(appOpts, opts...)
}

// 初始化
func Init() error {

	// db.Init must be the first
	include(db.Init, dao.Init, admin.Init, proxy.Init)

	for i, opt := range appOpts {
		log.Infof("[%d]--begin init--", i)
		if err := opt(); err != nil {
			log.Errorf("inits failed, err:%v\n", err)
			return err
		} else {
			log.Infof("[%d]--init succ--", i)
		}
	}
	return nil
}
