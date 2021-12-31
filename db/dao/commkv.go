package dao

import (
	"encoding/base64"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/encrypt"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/config"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
	"gorm.io/gorm/clause"
)

const commTableName = "comm"

// SetByStr 写
func SetByStr(key string, value string) error {
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
		return err
	}

	return nil
}

// GetByStr 读
func GetByStr(key string, defaultValue string) string {
	var err error
	var kv model.CommKv
	cli := db.Get()
	if err = cli.Table(commTableName).Where("`key` = ?", key).Take(&kv).Error; err != nil {
		log.Error(err.Error())
		return defaultValue
	}

	return kv.Value
}

// SetByStrEncrypt 加密写
func SetByStrEncrypt(key string, value string) error {
	encryptValue, err := encrypt.AesEncrypt([]byte(value), []byte(config.ServerConf.AesKey))
	if err != nil {
		return err
	}
	return SetByStr(key, base64.StdEncoding.EncodeToString(encryptValue))
}

// GetByStrDecrypt 解密读
func GetByStrDecrypt(key string, defaultValue string) string {
	var dbValue string
	var encryptValue, origValue []byte
	var err error
	if dbValue = GetByStr(key, defaultValue); dbValue == defaultValue {
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
