package admin

import (
	"os"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/encrypt"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/dao"
)

// InitAdmin 初始化管理员
func InitAdmin(username, password string) error {
	if err := dao.AddUserRecordIfNeeded(username, password); err != nil {
		log.Errorf("InitAuth err %v", err)
		return err
	}
	return nil
}

// Init 初始化管理员
func Init() error {
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	log.Debugf("GetUser user[%s] pwd[%s]", username, password)
	// conv password like website
	md5Pwd := encrypt.GenerateMd5(password)
	_ = InitAdmin(username, md5Pwd)
	return nil
}
