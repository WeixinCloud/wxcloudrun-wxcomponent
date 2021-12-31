package dao

import (
	"errors"
	"time"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/encrypt"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
)

const userTableName = "user"

// AddUserRecord 增加用户
func AddUserRecord(username string, password string) error {
	md5Pwd := encrypt.GenerateMd5(password)
	nowTime := time.Now()
	newUser := model.UserRecord{Username: username, Password: md5Pwd, CreateTime: nowTime, UpdateTime: nowTime}
	log.Debug(newUser)
	cli := db.Get()
	if err := cli.Table(userTableName).
		Create(&newUser).Error; err != nil {
		return err
	}
	return nil
}

// GetUserRecord 获取用户记录
func GetUserRecord(username string, password string) ([]*model.UserRecord, error) {
	md5Pwd := encrypt.GenerateMd5(password)
	log.Debugf("user[%s] pwd[%s]", username, md5Pwd)
	var records []*model.UserRecord
	cli := db.Get()
	result := cli.Table(userTableName).
		Where("username = ? and password = ?", username, md5Pwd).
		First(&records)
	return records, result.Error
}

// UpdateUserRecord 更新用户
func UpdateUserRecord(id int32, username string, password string, oldPassword string) error {
	log.Debugf("user[%s] pwd[%s] oldpwd[%s]", username, password, oldPassword)
	cli := db.Get()
	result := cli.Table(userTableName).Where("id = ?", id)
	if oldPassword != "" {
		md5OldPwd := encrypt.GenerateMd5(oldPassword)
		var records []*model.UserRecord
		result = result.Where("password = ?", md5OldPwd).First(&records)
		if len(records) == 0 {
			return errors.New("password error")
		}
	}
	if username != "" {
		if err := result.Update("username", username).Error; err != nil {
			log.Error("update username err ", err.Error())
			return err
		}
	}
	if password != "" {
		if oldPassword == "" {
			return errors.New("empty old password")
		}
		md5Pwd := encrypt.GenerateMd5(password)
		if err := result.Update("password", md5Pwd).Error; err != nil {
			log.Error("update password err ", err.Error())
			return err
		}
	}
	if err := result.Update("updatetime", time.Now()).Error; err != nil {
		log.Error("update updatetime err ", err.Error())
		return err
	}
	return nil
}
