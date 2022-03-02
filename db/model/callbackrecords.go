package model

import (
	"encoding/json"
	"time"
)

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

func (r WxCallbackComponentRecord) MarshalJSON() ([]byte, error) {
	type Alias WxCallbackComponentRecord
	return json.Marshal(&struct {
		Alias
		ReceiveTime int64 `json:"receiveTime"`
		CreateTime  int64 `json:"createTime"`
	}{
		Alias:       (Alias)(r),
		ReceiveTime: r.ReceiveTime.UnixNano() / 1e6,
		CreateTime:  r.CreateTime.UnixNano() / 1e6,
	})
}

func (r WxCallbackBizRecord) MarshalJSON() ([]byte, error) {
	type Alias WxCallbackBizRecord
	return json.Marshal(&struct {
		Alias
		ReceiveTime int64 `json:"receiveTime"`
		CreateTime  int64 `json:"createTime"`
	}{
		Alias:       (Alias)(r),
		ReceiveTime: r.ReceiveTime.UnixNano() / 1e6,
		CreateTime:  r.CreateTime.UnixNano() / 1e6,
	})
}
