package main

import (
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/inits"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/routers"
	"golang.org/x/sync/errgroup"
)

func main() {
	log.Infof("system begin")
	if err := inits.Init(); err != nil {
		log.Errorf("inits failed, err:%v", err)
		return
	}
	log.Infof("inits.Init Succ")

	var g errgroup.Group

	// 内部服务
	g.Go(func() error {
		r := routers.InnerServiceInit()
		if err := r.Run("127.0.0.1:8081"); err != nil {
			log.Error("startup inner service failed, err:%v", err)
			return err
		}
		return nil
	})

	// 外部服务
	g.Go(func() error {
		r := routers.Init()
		if err := r.Run(":80"); err != nil {
			log.Error("startup service failed, err:%v", err)
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Error(err)
	}
}
