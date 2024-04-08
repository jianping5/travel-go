package model

import (
	"gorm.io/datatypes"
	"time"
)

type BaseModel struct {
	Id         int64     `gorm:"primary_key"`
	CreateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	IsDeleted  bool
}

type Behavior struct {
	BaseModel
	UserId           int64 `json:"userId"`
	BehaviorItemType int   `json:"behavoirItemType"`
	BehaviorItemId   int64 `json:"behaviorItemId"`
}

func (Behavior) TableName() string {
	return "travel_data_behavior"
}

type ContentTag struct {
	BaseModel
	Name     string `json:"name"`
	ItemType int    `json:"itemType"`
	ItemId   int64  `json:"itemId"`
}

func (ContentTag) TableName() string {
	return "travel_data_content_tag"
}

type UserTag struct {
	BaseModel
	UserId int64          `json:"userId"`
	Tag    datatypes.JSON `gorm:"type:json" json:"tag"`
}

func (UserTag) TableName() string {
	return "travel_data_user_tag"
}
