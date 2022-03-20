package model

import (
	"time"
)

// UserRecord 用户信息
type ProxyLog struct {
	ID         int       `gorm:"column:id;primary_key" json:"id"` // 唯一ID
	Method     string    `gorm:"column:method" json:"method"`     // http method
	Uri        string    `gorm:"column:uri" json:"uri"`           // 请示路径
	Appid      string    `gorm:"column:appid" json:"appid"`       // appid
	Req        string    `gorm:"column:req" json:"req"`           // 请求wx内容
	Resp       string    `gorm:"column:resp" json:"resp"`         // wx的返回内容
	Userid     int       `gorm:"column:userid" json:"userid"`     // 操作者id
	Username   string    `gorm:"column:username" json:"username"` // 操作者名称
	CreateTime time.Time `gorm:"column:createtime;default:null"`  // 创建时间
}
