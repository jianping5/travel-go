package model

import "time"

type BaseModel struct {
	Id         int64     `gorm:"primary_key"`
	CreateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	IsDeleted  bool
}

type Strategy struct {
	BaseModel
	UserId      int64  `json:"userId"`
	Destination string `json:"destination"`
	Duration    string `json:"duration"`
	Budget      string `json:"budget"`
	TripGroup   string `json:"tripGroup"`
	TripMood    string `json:"tripMood"`
	Strategy    string `json:"strategy"`
}

func (Strategy) TableName() string {
	return "travel_intelligence_strategy"
}

type Conversation struct {
	BaseModel
	UserId      int64  `json:"userId"`
	Content     string `json:"content"`
	IsGenerated bool   `json:"isGenerated"`
}

func (Conversation) TableName() string {
	return "travel_intelligence_conversation"
}
