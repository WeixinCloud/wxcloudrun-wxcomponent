package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
)

var dbInstance *gorm.DB

// Init 初始化数据库
func Init() error {
	var user, pwd, addr, dataBase string
	user = os.Getenv("MYSQL_USERNAME")
	pwd = os.Getenv("MYSQL_PASSWORD")
	addr = os.Getenv("MYSQL_ADDRESS")
	dataBase = os.Getenv("MYSQL_DATABASE")
	if dataBase == "" {
		dataBase = "wxcomponent"
	}
	tcp := "tcp"
	source := "%s:%s@" + tcp + "(%s)/%s?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8&loc=Local&&parseTime=true"
	source = fmt.Sprintf(source, user, pwd, addr, dataBase)
	log.Debug("start inits mysql with ::::: " + source)

	db, err := gorm.Open(mysql.Open(source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		}})

	if err != nil {
		fmt.Println("DB Open error,err=", err.Error())
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("DB Init error,err=", err.Error())
		return err
	}

	// 用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(100)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(200)
	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	dbInstance = db

	fmt.Println("finish inits mysql with ", source)
	return nil
}

// Get
func Get() *gorm.DB {
	return dbInstance
}
