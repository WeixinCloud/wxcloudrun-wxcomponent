package dao

import (
	"fmt"
	"time"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
)

const componentTableName = "wxcallback_component"
const bizTableName = "wxcallback_biz"

// AddComponentCallBackRecord 增加第三方事件记录
func AddComponentCallBackRecord(callbackMessage *model.WxCallbackComponentRecord) error {
	log.Info("AddComponentCallBackRecord: " + callbackMessage.InfoType)
	var err error

	cli := db.Get()
	err = cli.Table(componentTableName).Create(callbackMessage).Error
	if err != nil {
		return err
	}

	return nil
}

// GetComponentCallBackRecordList 获取第三方事件记录
func GetComponentCallBackRecordList(startTime time.Time, endTime time.Time,
	infoType string, offset int, limit int) ([]*model.WxCallbackComponentRecord, int64, error) {
	var records = []*model.WxCallbackComponentRecord{}
	cli := db.Get()
	result := cli.Table(componentTableName).Where("receivetime between ? and ?", startTime, endTime)
	if infoType != "" {
		result = result.Where("infotype = ?", infoType)
	}
	var count int64
	result = result.Count(&count).Order("receivetime desc").Offset(offset).Limit(limit).Find(&records)
	return records, count, result.Error
}

// AddBizCallBackRecord 增加小程序事件记录
func AddBizCallBackRecord(callbackMessage *model.WxCallbackBizRecord) error {
	fmt.Println("[AddBizCallBackRecord]" + callbackMessage.ToUserName)
	var err error

	cli := db.Get()
	err = cli.Table(bizTableName).Create(callbackMessage).Error
	if err != nil {
		return err
	}

	return nil
}

// GetBizCallBackRecordList 获取小程序事件记录
func GetBizCallBackRecordList(startTime time.Time, endTime time.Time, appid string,
	msgType string, event string, offset int, limit int) ([]*model.WxCallbackBizRecord, int64, error) {
	var records = []*model.WxCallbackBizRecord{}
	cli := db.Get()
	result := cli.Table(bizTableName).Where("receivetime between ? and ?", startTime, endTime)
	if appid != "" {
		result = result.Where("appid = ?", appid)
	}
	if msgType != "" {
		result = result.Where("msgtype = ?", msgType)
	}
	if event != "" {
		result = result.Where("event = ?", event)
	}
	var count int64
	result = result.Count(&count).Order("receivetime desc").Offset(offset).Limit(limit).Find(&records)
	return records, count, result.Error
}
