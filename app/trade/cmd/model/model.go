package model

import "time"

type BaseModel struct {
	Id         int64     `gorm:"primary_key"`
	CreateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	IsDeleted  bool
}

type Work struct {
	BaseModel
	UserId      int64  `json:"userId"`
	CopyrightId int64  `json:"copyrightId"`
	Price       string `json:"price"`
	Status      int    `json:"status"`
}

func (Work) TableName() string {
	return "travel_trade_work"
}

type UserWork struct {
	BaseModel
	UserId int64 `json:"userId"`
	WorkId int64 `json:"workId"`
}

func (UserWork) TableName() string {
	return "travel_trade_user_work"
}

type Record struct {
	BaseModel
	WorkId    int64 `json:"workId"`
	OldUserId int64 `json:"oldUserId"`
	NewUserId int64 `json:"newUserId"`
}

func (Record) TableName() string {
	return "travel_trade_record"
}
