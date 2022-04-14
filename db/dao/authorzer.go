package dao

import (
	"time"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
	"gorm.io/gorm/clause"
)

const authorizerTableName = "authorizers"

// CreateOrUpdateAuthorizerRecord 创建或更新授权账号信息
func CreateOrUpdateAuthorizerRecord(record *model.Authorizer) error {
	var err error
	cli := db.Get()
	if err = cli.Table(authorizerTableName).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(record).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// BatchCreateOrUpdateAuthorizerRecord 批量创建或更新授权账号信息
func BatchCreateOrUpdateAuthorizerRecord(record *[]model.Authorizer) error {
	var err error

	cli := db.Get()
	if err = cli.Table(authorizerTableName).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(record).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// ClearAuthorizerRecordsBefore 删除指定时间前的所有授权账号记录
func ClearAuthorizerRecordsBefore(time time.Time) error {
	var err error

	cli := db.Get()
	if err = cli.Table(authorizerTableName).
		Where("updatetime < ?", time).
		Delete(model.Authorizer{}).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// GetAuthorizerRecords 获取授权账号记录
func GetAuthorizerRecords(appid string, offset int, limit int) ([]*model.Authorizer, int64, error) {
	var records = []*model.Authorizer{}
	cli := db.Get()
	result := cli.Table(authorizerTableName)
	if appid != "" {
		result = result.Where("appid = ?", appid)
	}
	var count int64
	result = result.Count(&count).Offset(offset).Limit(limit).Find(&records)
	return records, count, result.Error
}

// DelAuthorizerRecord 删除授权账号记录
func DelAuthorizerRecord(appid string) error {
	var err error

	cli := db.Get()
	if err = cli.Table(authorizerTableName).
		Where("appid = ?", appid).Delete(model.Authorizer{}).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// GetDevWeAppRecords 获取代开发小程序
func GetDevWeAppRecords(offset int, limit int, appid string) ([]*model.Authorizer, int64, error) {
	var records = []*model.Authorizer{}
	cli := db.Get()
	result := cli.Table(authorizerTableName)
	var count int64
	result = result.Where("apptype = ? AND funcinfo LIKE ?", 0, "%18%")
	if len(appid) != 0 {
		result = result.Where("appid = ?", appid)
	}
	result = result.Count(&count).Offset(offset).Limit(limit).Find(&records)
	return records, count, result.Error
}
