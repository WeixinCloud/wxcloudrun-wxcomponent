package model

import (
	"time"
)

// UserRecord 用户信息
type UserRecord struct {
	ID         int32     `gorm:"column:id;primaryKey" json:"id"`  // 唯一ID
	Username   string    `gorm:"column:username" json:"username"` // 用户名
	Password   string    `gorm:"column:password" json:"password"` // 密码md5
	CreateTime time.Time `gorm:"column:createtime;default:null"`  // 创建时间
	UpdateTime time.Time `gorm:"column:updatetime;default:null"`  // 更新时间
}
