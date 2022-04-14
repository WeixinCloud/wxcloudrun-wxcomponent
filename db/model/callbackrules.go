package model

import "time"

// WxCallbackRule 回调消息转发规则
type WxCallbackRule struct {
	ID         int32     `gorm:"column:id;primaryKey" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	InfoType   string    `gorm:"column:infotype" json:"infoType"`
	MsgType    string    `gorm:"column:msgtype" json:"msgType"`
	Event      string    `gorm:"column:event" json:"event"`
	Type       int       `gorm:"column:type" json:"type"`
	Open       int       `gorm:"column:open" json:"open"`
	Info       string    `gorm:"column:info" json:"info"`
	CreateTime time.Time `gorm:"column:createtime;default:null" json:"createTime"`
	UpdateTime time.Time `gorm:"column:updatetime;default:null" json:"updatetime"`
}

// HttpProxyConfig http转发配置
type HttpProxyConfig struct {
	Port int    `json:"port"`
	Path string `json:"path"`
}

const PROXYTYPE_HTTP = 1
const CALLBACKTYPE_COM = 1
const CALLBACKTYPE_BIZ = 2
