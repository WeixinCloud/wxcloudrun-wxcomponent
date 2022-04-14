package model

import "time"

// Counter 计数器
type Counter struct {
	Key        string    `gorm:"column:key;uniqueIndex"`
	Value      uint      `gorm:"column:value"`
	CreateTime time.Time `gorm:"column:createtime;default:null" json:"createTime"`
	UpdateTime time.Time `gorm:"column:updatetime;default:null" json:"updatetime"`
}
