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

type Favorite struct {
	BaseModel
	UserId int64  `json:"userId"`
	Name   string `json:"name"`
}

func (Favorite) TableName() string {
	return "travel_social_favorite"
}

type Favor struct {
	BaseModel
	UserId     int64 `json:"userId"`
	FavoriteId int64 `json:"favoriteId"`
	ItemType   int   `json:"itemType"`
	ItemId     int64 `json:"itemId"`
}

func (Favor) TableName() string {
	return "travel_social_favor"
}

type History struct {
	BaseModel
	UserId   int64 `json:"userId"`
	ItemType int   `json:"itemType"`
	ItemId   int64 `json:"itemId"`
}

func (History) TableName() string {
	return "travel_social_history"
}

type Community struct {
	BaseModel
	UserId      int64  `json:"userId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
	MemberCount int    `json:"memberCount"`
}

func (Community) TableName() string {
	return "travel_social_community"
}

type UserCommunity struct {
	BaseModel
	UserId      int64 `json:"userId"`
	Role        int   `json:"role"`
	CommunityId int64 `json:"communityId"`
}

func (UserCommunity) TableName() string {
	return "travel_social_user_community"
}

type Dynamic struct {
	BaseModel
	UserId       int64  `json:"userId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	CommunityId  int64  `json:"communityId"`
	FileType     int    `json:"fileType"`
	Content      string `json:"content"`
	LikeCount    int    `json:"likeCount"`
	CommentCount int    `json:"commentCount"`
}

func (Dynamic) TableName() string {
	return "travel_social_dynamic"
}

type Content struct {
	BaseModel
	UserId       int64          `json:"userId"`
	Title        string         `json:"title"`
	ItemType     int            `json:"itemType"`
	CoverUrl     string         `json:"coverUrl"`
	Content      string         `json:"contentUrl"`
	Description  string         `json:"description"`
	Tag          datatypes.JSON `gorm:"type:json" json:"tag"`
	LikeCount    int            `json:"likeCount"`
	CommentCount int            `json:"commentCount"`
	FavorCount   int            `json:"favorCount"`
}

func (Content) TableName() string {
	return "travel_social_content"
}

type Comment struct {
	BaseModel
	UserId          int64  `json:"userId"`
	CommentItemType int    `json:"commentItemType"`
	CommentItemId   int64  `json:"commentItemId"`
	ParentUserId    int64  `json:"parentUserId"`
	TopId           int64  `json:"topId"`
	Content         string `json:"content"`
	LikeCount       int    `json:"likeCount"`
	ReplyCount      int    `json:"replyCount"`
}

func (Comment) TableName() string {
	return "travel_social_comment"
}

type Message struct {
	BaseModel
	UserId        int64  `json:"userId"`
	ItemType      int    `json:"itemType"`
	ItemId        int64  `json:"itemId"`
	MessageUserId int64  `json:"messageUserId"`
	MessageType   int    `json:"messageType"`
	MessageStatus int    `json:"messageStatus"`
	Content       string `json:"content"`
}

func (Message) TableName() string {
	return "travel_social_message"
}

type Like struct {
	BaseModel
	UserId      int64 `json:"userId"`
	ItemType    int   `json:"itemType"`
	ItemId      int64 `json:"itemId"`
	LikedStatus bool  `json:"likedStatus"`
}

func (Like) TableName() string {
	return "travel_social_like"
}

type Copyright struct {
	BaseModel
	UserId    int64  `json:"userId"`
	ItemType  int    `json:"itemType"`
	ItemId    int64  `json:"itemId"`
	Metadata  string `json:"metadata"`
	TradeHash string `json:"tradeHash"`
	Address   string `json:"address"`
	Status    int    `json:"status"`
}

func (Copyright) TableName() string {
	return "travel_social_copyright"
}

type ContentUpdateReq struct {
	Id          int64          `json:"id"`
	Title       string         `json:"title"`
	Content     string         `json:"content"`
	CoverUrl    string         `json:"coverUrl"`
	Description string         `json:"description"`
	Tag         datatypes.JSON `gorm:"type:json" json:"tag"`
}
