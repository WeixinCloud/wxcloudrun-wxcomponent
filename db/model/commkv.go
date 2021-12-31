package model

// CommKv 通用的kv
type CommKv struct {
	Key   string `gorm:"column:key"`
	Value string `gorm:"column:value"`
}
