package models

import (
	"gorm.io/gorm"
	"time"
)

//定义公用模型字段进行数据库迁移

type ID struct {
	ID uint `gorm:"primarykey" json:"id"`
}

// Timestamps 定义公共的时间戳字段
// gorm:"column:create_at" 为该字段在数据库中的名称
// json:"create_at" 为该字段在json中的名称
type Timestamps struct {
	CreateAt time.Time `gorm:"column:create_at" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at" json:"update_at"`
}

// SoftDeletes 定义软删除字段
// gorm :"index" 为该字段创建索引
// json:"delete_at" 为该字段在json中的名称
type SoftDeletes struct {
	DeleteAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}
