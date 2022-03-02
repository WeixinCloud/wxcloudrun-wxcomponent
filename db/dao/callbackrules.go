package dao

import (
	"fmt"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

const callbackRuleTableName = "wxcallback_rules"

// GetWxCallBackRules 获取所有转发规则
func GetWxCallBackRuleList(offset int, limit int, callbackType int) ([]*model.WxCallbackRule, int64, error) {
	var records = []*model.WxCallbackRule{}
	cli := db.Get()
	result := cli.Table(callbackRuleTableName)
	if callbackType == model.CALLBACKTYPE_COM {
		result = result.Where("infotype != \"\"")
	} else if callbackType == model.CALLBACKTYPE_BIZ {
		result = result.Where("infotype = \"\"")
	}
	var count int64
	result = result.Count(&count).Offset(offset).Limit(limit).Find(&records)
	return records, count, result.Error
}

// UpdateWxCallBackRule 更新转发规则
func UpdateWxCallBackRule(record *model.WxCallbackRule) error {
	cli := db.Get()
	if result := cli.Table(callbackRuleTableName).
		Where("id = ?", record.ID).
		Select("name", "infotype", "msgtype", "event", "type", "open", "Info").
		Updates(record); result.Error != nil {
		log.Error(result.Error)
		return result.Error
	}
	cacheCli := db.GetCache()
	key := genCallBackRuleKey(record.InfoType, record.MsgType, record.Event)
	cacheCli.Set(key, record, cache.DefaultExpiration)
	return nil
}

// AddWxCallBackRule 添加转发规则
func AddWxCallBackRule(record *model.WxCallbackRule) error {
	cli := db.Get()
	if result := cli.Table(callbackRuleTableName).Create(record); result.Error != nil {
		log.Error(result.Error)
		return result.Error
	}
	cacheCli := db.GetCache()
	key := genCallBackRuleKey(record.InfoType, record.MsgType, record.Event)
	cacheCli.Set(key, record, cache.DefaultExpiration)
	return nil
}

// DelWxCallBackRule 删除转发规则
func DelWxCallBackRule(id int32) error {
	cli := db.Get()
	var record *model.WxCallbackRule
	if result := cli.Table(callbackRuleTableName).
		Where("id = ?", id).Delete(&record); result.Error != nil {
		log.Error(result.Error)
		return result.Error
	}
	log.Debug(record)
	cacheCli := db.GetCache()
	key := genCallBackRuleKey(record.InfoType, record.MsgType, record.Event)
	cacheCli.Delete(key)
	return nil
}

// GetWxCallBackRuleWithCache 通过消息类型获取转发规则 有缓存
func GetWxCallBackRuleWithCache(infoType string, msgType string, event string) (*model.WxCallbackRule, error) {
	cacheCli := db.GetCache()
	key := genCallBackRuleKey(infoType, msgType, event)
	value, found := cacheCli.Get(key)
	if found {
		log.Infof("hit cache key:", key)
		if value == nil {
			return nil, nil
		}
		return value.(*model.WxCallbackRule), nil
	} else {
		result, err := getWxCallBackRule(infoType, msgType, event)
		if err == gorm.ErrRecordNotFound {
			log.Infof("empty record")
			cacheCli.Set(key, nil, cache.DefaultExpiration)
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		cacheCli.Set(key, result, cache.DefaultExpiration)
		return result, nil
	}
}

// GetWxCallBackRuleById 通过id获取转发规则
func GetWxCallBackRuleById(id int32) (*model.WxCallbackRule, error) {
	var record *model.WxCallbackRule
	cli := db.Get()
	result := cli.Table(callbackRuleTableName).
		Where("id = ?", id).
		Take(&record)
	return record, result.Error
}

func getWxCallBackRule(infoType string, msgType string, event string) (*model.WxCallbackRule, error) {
	var record *model.WxCallbackRule
	cli := db.Get()
	result := cli.Table(callbackRuleTableName).
		Where("infotype = ? and msgtype = ? and event = ?", infoType, msgType, event).
		Take(&record)
	return record, result.Error
}

func genCallBackRuleKey(infoType string, msgType string, event string) string {
	return fmt.Sprintf("cb_%s_%s_%s", infoType, msgType, event)
}
