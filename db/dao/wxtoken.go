package dao

import (
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const wxTokenTableName = "wxtoken"

// GetAccessToken 获取AccessToken
func GetAccessToken(appid string, tokenType int) (*model.WxToken, bool, error) {
	cli := db.Get()
	var record *model.WxToken
	if result := cli.Table(wxTokenTableName).
		Where("appid = ? and type = ?", appid, tokenType).
		Take(&record); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, false, nil
		}
		return nil, false, result.Error
	} else {
		return record, true, nil
	}
}

// SetAccessToken 创建或更新wxtoken
func SetAccessToken(record *model.WxToken) error {
	cli := db.Get()
	if err := cli.Table(wxTokenTableName).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(record).Error; err != nil {
		return err
	}
	return nil
}
