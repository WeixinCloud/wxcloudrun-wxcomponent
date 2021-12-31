package model

import "time"

// WxCallbackComponentRecord 第三方授权事件的记录
type WxCallbackComponentRecord struct {
	ReceiveTime time.Time `gorm:"column:receivetime" json:"receiveTime"`
	CreateTime  time.Time `gorm:"column:createtime" json:"createTime"`
	InfoType    string    `gorm:"column:infotype" json:"infoType"`
	PostBody    string    `gorm:"column:postbody" json:"postBody"`
}

// WxCallbackBizRecord 小程序授权事件记录
type WxCallbackBizRecord struct {
	ReceiveTime time.Time `gorm:"column:receivetime" json:"receiveTime"`
	CreateTime  time.Time `gorm:"column:createtime" json:"createTime"`
	Appid       string    `gorm:"column:appid" json:"appid"`
	ToUserName  string    `gorm:"column:tousername" json:"toUserName"`
	MsgType     string    `gorm:"column:msgtype" json:"msgType"`
	Event       string    `gorm:"column:event" json:"event"`
	PostBody    string    `gorm:"column:postbody" json:"postBody"`
}
