package dao

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/wonderivan/logger"
	"k8s-demo/db"
	"k8s-demo/model"
	"time"
)

var Event event

type event struct{}

type Events struct {
	Items []*model.Event     `json:"items"`
	Total int                `json:"total"`
}

func(*event) GetList(name, cluster string, page, limit int) (*Events, error) {
	//定义分页数据的起始位置
	startSet := (page-1) * limit

	//定义数据库查询返回内容
	var (
		eventList []*model.Event
		total        int
	)

	//数据库查询，Limit方法用于限制条数，Offset方法设置起始位置
	tx := db.GORM.
		Model(&model.Event{}).
		Where("name like ? and cluster = ?", "%" + name + "%", cluster).
		Count(&total).
		Limit(limit).
		Offset(startSet).
		Order("id desc").
		Find(&eventList)

	if tx.Error != nil {
		logger.Error(fmt.Sprintf("获取Event列表失败,%v\n", tx.Error))
		return nil, errors.New(fmt.Sprintf("获取Event列表失败,%v\n", tx.Error))
	}

	return &Events{
		Items: eventList,
		Total: total,
	}, nil
}

//新增event
func(*event) Add(event *model.Event) error {
	tx := db.GORM.Create(&event)
	if tx.Error != nil {
		logger.Error(fmt.Sprintf("添加Event失败, %v\n", tx.Error))
		return errors.New(fmt.Sprintf("添加Event失败, %v\n", tx.Error))
	}
	return nil
}

//查询单个event
func(*event) HasEvent(name, kind, namespace, reason string, eventTime time.Time, cluster string) (*model.Event, bool, error) {
	data := &model.Event{}
	tx := db.GORM.Where("name = ? and kind = ? and namespace = ? and reason = ? and event_time = ? and cluster = ?",
		name, kind, namespace, reason, eventTime, cluster).First(&data)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, false, nil
	}
	if tx.Error != nil {
		logger.Error(fmt.Sprintf("查询Event失败, %v\n", tx.Error))
		return nil, false, errors.New(fmt.Sprintf("查询Event失败, %v\n", tx.Error))
	}

	return data, true, nil
}