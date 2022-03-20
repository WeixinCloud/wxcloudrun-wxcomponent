package dao

import (
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/utils"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
	"strconv"
	"time"
)

const proxyLogTableName = "ldy_proxy_log"

func AddRecord(method string, uri string, appid string, req []byte, resp []byte, user *utils.Claims) error {
	userId, _ := strconv.Atoi(user.ID)
	record := model.ProxyLog{
		Method:     method,
		Uri:        uri,
		Appid:      appid,
		Req:        string(req),
		Resp:       string(resp),
		Userid:     userId,
		Username:   user.UserName,
		CreateTime: time.Now(),
	}

	cli := db.Get()
	if err := cli.Table(proxyLogTableName).
		Create(&record).Error; err != nil {
		log.Errorf("save proxy log error: err=%v, data:%v", err, record)
		return err
	}

	log.Infof("Save Proxy Log: %v", record)
	return nil
}
