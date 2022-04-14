package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/patrickmn/go-cache"
)

var dbInstance *gorm.DB
var cacheInstance *cache.Cache

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

	checkTables()

	// 初始化cache
	cacheInstance = cache.New(5*time.Minute, 10*time.Minute)

	return nil
}

func checkTables() {
	dbInstance.Exec("CREATE TABLE IF NOT EXISTS `wxcallback_component` (`id` INT UNSIGNED AUTO_INCREMENT, `receivetime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, `createtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, `infotype` VARCHAR(64) NOT NULL DEFAULT '', `postbody` TEXT NOT NULL, PRIMARY KEY (`id`), INDEX(`receivetime`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;")
	dbInstance.Exec("CREATE TABLE IF NOT EXISTS `wxcallback_biz` (`id` INT UNSIGNED AUTO_INCREMENT, `receivetime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, `createtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, `tousername` VARCHAR(64) NOT NULL DEFAULT '', `appid` VARCHAR(64) NOT NULL DEFAULT '', `msgtype` VARCHAR(64) NOT NULL DEFAULT '', `event` VARCHAR(64) NOT NULL DEFAULT '', `postbody` TEXT NOT NULL, PRIMARY KEY (`id`), INDEX(`receivetime`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;")
	dbInstance.Exec("CREATE TABLE IF NOT EXISTS `comm` (`key` VARCHAR(64) NOT NULL, `value` TEXT NOT NULL, `createtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, `updatetime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, PRIMARY KEY (`key`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;")
	dbInstance.Exec("CREATE TABLE IF NOT EXISTS `user` ( `id` INT NOT NULL AUTO_INCREMENT, `username` VARCHAR(32) NOT NULL, `password` VARCHAR(64) NOT NULL, `createtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, `updatetime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, PRIMARY KEY (`ID`), UNIQUE KEY `user_username_uindex` (`username`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8;")
	dbInstance.Exec("CREATE TABLE IF NOT EXISTS `authorizers` ( `id` INT NOT NULL AUTO_INCREMENT, `appid` VARCHAR(32) NOT NULL, `apptype` INT NOT NULL DEFAULT 0, `servicetype` INT NOT NULL DEFAULT 0, `nickname` VARCHAR(32) NOT NULL NOT NULL DEFAULT '', `username` VARCHAR(32) NOT NULL NOT NULL DEFAULT '', `headimg` VARCHAR(256) NOT NULL DEFAULT '', `qrcodeurl` VARCHAR(256) NOT NULL DEFAULT '',`principalname` VARCHAR(64) NOT NULL DEFAULT '', `refreshtoken` VARCHAR(128) NOT NULL DEFAULT '', `funcinfo` VARCHAR(128) NOT NULL DEFAULT '', `verifyinfo` INT NOT NULL DEFAULT -1, `authtime` TIMESTAMP NOT NULL, `updatetime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, PRIMARY KEY (`id`), UNIQUE KEY(`appid`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;")
	dbInstance.Exec("CREATE TABLE IF NOT EXISTS `wxcallback_rules` (`id` INT UNSIGNED AUTO_INCREMENT, `name` VARCHAR(64) NOT NULL DEFAULT '', `infotype` VARCHAR(64) NOT NULL DEFAULT '', `msgtype` VARCHAR(64) NOT NULL DEFAULT '', `event` VARCHAR(64) NOT NULL DEFAULT '', `type` INT NOT NULL DEFAULT 0, `open` INT NOT NULL DEFAULT 0,  `info` TEXT NOT NULL, `createtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, `updatetime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, PRIMARY KEY (`id`), UNIQUE KEY(infotype, msgtype, event)) ENGINE=InnoDB DEFAULT CHARSET=utf8;")
	dbInstance.Exec("CREATE TABLE IF NOT EXISTS `wxtoken` (`id` INT UNSIGNED AUTO_INCREMENT, `type` INT NOT NULL DEFAULT 0, `appid` VARCHAR(128) NOT NULL DEFAULT '', `token` TEXT NOT NULL, `expiretime` TIMESTAMP NOT NULL, `createtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, `updatetime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, PRIMARY KEY (`id`), UNIQUE KEY `appid_uindex` (`appid`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8;")
	dbInstance.Exec("CREATE TABLE IF NOT EXISTS `counter` (`id` INT UNSIGNED AUTO_INCREMENT, `key` VARCHAR(64) NOT NULL, `value` INT UNSIGNED, `createtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, `updatetime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, PRIMARY KEY (`id`), UNIQUE KEY(`key`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;")
	dbInstance.Exec("CREATE TABLE IF NOT EXISTS `counter` (`id` INT UNSIGNED AUTO_INCREMENT, `key` VARCHAR(64) NOT NULL, `value` INT UNSIGNED, `createtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, `updatetime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, PRIMARY KEY (`id`), UNIQUE KEY(`key`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;")
}

// Get
func Get() *gorm.DB {
	return dbInstance
}

// GetCache
func GetCache() *cache.Cache {
	return cacheInstance
}
