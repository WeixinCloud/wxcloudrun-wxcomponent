package model

import "time"

// WxToken 微信相关的token
type WxToken struct {
	Type       int       `gorm:"column:type"`
	Appid      string    `gorm:"column:appid"`
	Token      string    `gorm:"column:token"`
	Expiretime time.Time `gorm:"column:expiretime"`
	CreateTime time.Time `gorm:"column:createtime;default:null"`
	UpdateTime time.Time `gorm:"column:updatetime;default:null"`
}

const WXTOKENTYPE_AUTH = 1
const WXTOKENTYPE_OWN = 2
