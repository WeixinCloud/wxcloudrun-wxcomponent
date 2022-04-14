package dao

import (
	"fmt"
	"time"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
	"gorm.io/gorm"
)

const counterTableName = "counter"

// GetCurrent 获取当前数值
func GetCurrent(key string) (uint, error) {
	var err error
	var counter model.Counter
	cli := db.Get()
	if err = cli.Table(counterTableName).Where("`key` = ?", key).Take(&counter).Error; err != nil {
		log.Error(err)
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}
	return counter.Value, nil
}

// AddOne +1 结果不可超过limit
func AddOne(key string, limit uint) error {
	var err error
	var counter model.Counter
	cli := db.Get()
	if err = cli.Table(counterTableName).Where("`key` = ?", key).Take(&counter).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			counter.Key = key
			counter.Value = 1
			if err := cli.Table(counterTableName).Create(counter).Error; err != nil {
				log.Error(err)
				return err
			}
		} else {
			log.Error(err)
			return err
		}
	}
	result := cli.Table(counterTableName).Where("`key` = ? and value < ? ", key, limit).
		UpdateColumn("value", gorm.Expr("value + ?", 1))
	if err := result.Error; err != nil {
		log.Error(err)
		return err
	}
	if result.RowsAffected == 0 {
		log.Error("affect zero row")
		return fmt.Errorf("affect zero row")
	}
	return nil
}

func clearExpiredRecord() {
	var err error
	cli := db.Get()
	result := cli.Table(counterTableName).
		Where("updatetime < ?", time.Now().Add(-time.Hour)).
		Delete(model.Counter{})
	if err = result.Error; err != nil {
		log.Error(err)
	}
	log.Info("delete expired record: ", result.RowsAffected)
}

func startClearExpiredRecordTask() {
	now := time.Now()
	next := now.Add(time.Hour * 24)
	next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
	t := time.NewTimer(next.Sub(now))
	<-t.C
	clearExpiredRecord()
	timer := time.NewTicker(24 * time.Hour)
	for range timer.C {
		clearExpiredRecord()
	}
}
