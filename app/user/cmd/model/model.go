package model

import "time"

type BaseModel struct {
	Id         int64     `gorm:"primary_key"`
	CreateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	IsDeleted  bool
}

type User struct {
	BaseModel
	Account   string `json:"account" gorm:"size:50"`
	Password  string `json:"password" gorm:"size:255"`
	Avatar    string `json:"avatar" gorm:"size:255"`
	Signature string `json:"signature" gorm:"size:255"`
	Email     string `json:"email" gorm:"size:50"`
}

func (User) TableName() string {
	return "travel_user_info"
}

type Follow struct {
	BaseModel
	UserId       int64 `json:"userId"`
	FollowUserId int64 `json:"followUserId"`
}

func (Follow) TableName() string {
	return "travel_user_follow"
}
