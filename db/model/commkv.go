package model

import "time"

// CommKv 通用的kv
type CommKv struct {
	Key        string    `gorm:"column:key"`
	Value      string    `gorm:"column:value"`
	CreateTime time.Time `gorm:"column:createtime;default:null" json:"createTime"`
	UpdateTime time.Time `gorm:"column:updatetime;default:null" json:"updatetime"`
}
