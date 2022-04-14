package dao

import (
	"encoding/base64"
	"time"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/encrypt"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/config"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
	"gorm.io/gorm/clause"
)

const commTableName = "comm"

// SetCommKvWithCache 写kv并更新缓存
func SetCommKvWithCache(key string, value string, d time.Duration) error {
	if err := SetCommKv(key, value); err != nil {
		log.Error(err.Error())
		return err
	}
	db.GetCache().Set(key, value, d)
	return nil
}

// GetCommKvWithCache 先读缓存再读数据库，读取数据之后写到缓存
func GetCommKvWithCache(key string, defaultValue string, d time.Duration) string {
	cacheCli := db.GetCache()
	if value, found := cacheCli.Get(key); found {
		return value.(string)
	}
	value := GetCommKv(key, defaultValue)
	if value != defaultValue {
		cacheCli.Set(key, value, d)
	}
	return value
}

// SetCommKv 覆盖写
func SetCommKv(key string, value string) error {
	log.Infof("SetCommKv: %s %s", key, value)
	var err error
	var kv = model.CommKv{
		Key:   key,
		Value: value,
	}

	cli := db.Get()
	err = cli.Table(commTableName).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(kv).Error
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

// AddCommKv 添加一个记录 key重复会报错
func AddCommKv(key string, value string) error {
	log.Infof("AddCommKv: %s %s", key, value)
	var kv = model.CommKv{
		Key:   key,
		Value: value,
	}

	cli := db.Get()
	if err := cli.Table(commTableName).Create(kv).Error; err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

// DelCommKv 删除记录
func DelCommKv(key string) error {
	log.Infof("DelCommKv: %s", key)
	cli := db.Get()
	if err := cli.Table(commTableName).Where("`key` = ?", key).
		Delete(&model.CommKv{}).Error; err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

// DelExpiredCommKv 删除超时的记录
func DelExpiredCommKv(key string, d time.Duration) (int64, error) {
	log.Infof("DelExpiredCommKv: %s", key)
	cli := db.Get()
	result := cli.Table(commTableName).
		Where("`key` = ? and UpdateTime < ?", key,
			time.Now().Add(-d)).Delete(&model.CommKv{})
	if result.Error != nil {
		log.Error(result.Error)
		return result.RowsAffected, result.Error
	}
	return result.RowsAffected, nil
}

// GetCommKv 读
func GetCommKv(key string, defaultValue string) string {
	var err error
	var kv model.CommKv
	cli := db.Get()
	if err = cli.Table(commTableName).Where("`key` = ?", key).Take(&kv).Error; err != nil {
		log.Error(err.Error())
		return defaultValue
	}
	return kv.Value
}

// SetCommKvEncrypt 加密写
func SetCommKvEncrypt(key string, value string) error {
	encryptValue, err := encrypt.AesEncrypt([]byte(value), []byte(config.ServerConf.AesKey))
	if err != nil {
		return err
	}
	return SetCommKv(key, base64.StdEncoding.EncodeToString(encryptValue))
}

// GetCommKvDecrypt 解密读
func GetCommKvDecrypt(key string, defaultValue string) string {
	var dbValue string
	var encryptValue, origValue []byte
	var err error
	if dbValue = GetCommKv(key, defaultValue); dbValue == defaultValue {
		return defaultValue
	}
	if encryptValue, err = base64.StdEncoding.DecodeString(dbValue); err != nil {
		return defaultValue
	}
	if origValue, err = encrypt.AesDecrypt(encryptValue, []byte(config.ServerConf.AesKey)); err != nil {
		log.Error(err.Error())
		return defaultValue
	}

	return string(origValue)
}

// Lock 申请锁
func Lock(key string, value string, expire time.Duration) error {
	if err := AddCommKv(key, value); err != nil {
		affectedRows, delErr := DelExpiredCommKv(key, expire)
		if delErr != nil {
			return delErr
		}
		if affectedRows != 0 {
			log.Debug("Del Expired Lock")
			return AddCommKv(key, value)
		}
		return err
	}
	return nil
}

// Lock 释放锁
func UnLock(key string) error {
	return DelCommKv(key)
}
