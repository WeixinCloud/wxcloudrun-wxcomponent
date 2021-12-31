package main

import (
	"fmt"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/inits"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/routers"
)

func main() {
	log.Infof("system begin")
	if err := inits.Init(); err != nil {
		log.Errorf("inits failed, err:%v\n", err)
		return
	}

	log.Infof("inits.Init Succ")

	// 初始化路由
	r := routers.Init()
	if err := r.Run(":80"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
		return
	}
	log.Infof("system ok")
}
