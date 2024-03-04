// Code generated by goctl. DO NOT EDIT.
package types

type CommentCreateReq struct {
	CommentItemType int    `json:"commentItemType"`
	CommentItemId   int64  `json:"commentItemId"`
	ParentUserId    int64  `json:"parentUserId"`
	TopId           int64  `json:"topId"`
	Content         string `json:"content"`
}

type CommentDeleteReq struct {
	Id int64 `json:"id"`
}

type CommentListReq struct {
	CommentItemType int   `json:"commentItemType"`
	CommentItemId   int64 `json:"commentItemId"`
	PageNum         int   `json:"pageNum"`
	PageSize        int   `json:"pageSize"`
}

type CommentListResp struct {
	List  []CommentListView `json:"list"`
	Total int               `json:"total"`
}

type CommentListView struct {
	TopComment  CommentView   `json:"topComment"`
	CommentList []CommentView `json:"commentList"`
}

type CommentView struct {
	Id              int64        `json:"id"`
	UserId          int64        `json:"userId"`
	UserInfo        UserInfoView `json:"userInfo"`
	ParentUserInfo  UserInfoView `json:"parentUserInfo"`
	CommentItemType int          `json:"commentItemType"`
	CommentItemId   int64        `json:"commentItemId"`
	ParentUserId    int64        `json:"parentUserId"`
	TopId           int64        `json:"topId"`
	Content         string       `json:"content"`
	IsLiked         bool         `json:"isLiked"`
}

type CommunityCreateReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CommunityDeleteReq struct {
	Id int64 `json:"id"`
}

type CommunityDetailReq struct {
	Id int64 `json:"id"`
}

type CommunityDetailResp struct {
	Community CommunityView `json:"community"`
	UserId    int64         `json:"userId"`
	Account   string        `json:"account"`
	Avatar    string        `json:"avatar"`
}

type CommunityDynamicCreateReq struct {
	CommunityId int64  `json:"communityId"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	FileType    int    `json:"fileType"`
}

type CommunityDynamicDeleteReq struct {
	Id int64 `json:"id"`
}

type CommunityDynamicListReq struct {
	Type         int  `json:"type"`
	JoinedSwitch bool `json:"joinedSwitch"`
	PageNum      int  `json:"pageNum"`
	PageSize     int  `json:"pageSize"`
}

type CommunityDynamicListResp struct {
	List  []CommunityDynamicView `json:"list"`
	Total int                    `json:"total"`
}

type CommunityDynamicView struct {
	Id           int64         `json:"id"`
	UserInfo     UserInfoView  `json:"userInfo"`
	CommunityId  int64         `json:"communityId"`
	Community    CommunityView `json:"communityView"`
	Title        string        `json:"title"`
	Content      string        `json:"content"`
	FileType     int           `json:"fileType"`
	LikeCount    int           `json:"likeCount"`
	CommentCount int           `json:"CommentCount"`
	CreateTime   string        `json:"createTime"`
	IsLiked      bool          `json:"isLiked"`
}

type CommunityJoinReq struct {
	CommunityId int64 `json:"communityId"`
	Role        int   `json:"role"`
}

type CommunityListReq struct {
	UserId int64 `json:"userId"`
}

type CommunityListResp struct {
	List []CommunityView `json:"list"`
}

type CommunityQuitReq struct {
	CommunityId int64 `json:"communityId"`
}

type CommunityUpdateReq struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
}

type CommunityView struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
	MemberCount int    `json:"memberCount"`
	CreateTime  string `json:"createTime"`
}

type ContentCreateReq struct {
	ItemType    int    `json:"itemType"`
	Title       string `json:"title"`
	CoverUrl    string `json:"coverUrl"`
	Content     string `json:"content"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
}

type ContentDeleteReq struct {
	ItemType int   `json:"itemType"`
	ItemId   int64 `json:"itemId"`
}

type ContentListReq struct {
	ContentType int `json:"contentType"`
	ItemType    int `json:"itemType"`
	PageNum     int `json:"pageNum"`
	PageSize    int `json:"pageSize"`
}

type ContentListResp struct {
	List  []ContentView `json:"list"`
	Total int           `json:"total"`
}

type ContentSimpleView struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	CoverUrl   string `json:"coverUrl"`
	LikeCount  int    `json:"likeCount"`
	CreateTime string `json:"CreateTime"`
}

type ContentUpdateReq struct {
	ItemType    int    `json:"itemType"`
	ItemId      int64  `json:"itemId"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CoverUrl    string `json:"coverUrl"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
}

type ContentView struct {
	Id           int64        `json:"id"`
	UserId       int64        `json:"userId"`
	UserInfo     UserInfoView `json:"userInfo"`
	Title        string       `json:"title"`
	CoverUrl     string       `json:"coverUrl"`
	Content      string       `json:"content"`
	Description  string       `json:"description"`
	Tag          string       `json:"tag"`
	LikeCount    int          `json:"likeCount"`
	CommentCount int          `json:"commentCount"`
	FavorCount   int          `json:"favorCount"`
	CreateTime   string       `json:"createTime"`
	IsLiked      bool         `json:"isLiked"`
	IsFavored    bool         `json:"isFavored"`
}

type CopyrightCreateReq struct {
	ItemType int   `json:"itemType"`
	ItemId   int64 `json:"itemId"`
}

type CopyrightDetailReq struct {
	UserId   int64 `json:"userId"`
	ItemType int   `json:"itemType"`
	ItemId   int64 `json:"itemId"`
}

type CopyrightDetailResp struct {
	Copyright CopyrightView `json:"copyright"`
	UserInfo  UserInfoView  `json:"userInfo"`
}

type CopyrightView struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"userId"`
	ItemType   int    `json:"itemType"`
	ItemId     int64  `json:"itemId"`
	Metadata   string `json:"metadata"`
	TradeHash  string `json:"tradeHash"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	CreateTime string `json:"createTime"`
}

type FavorDeleteReq struct {
	Id int64 `json:"id"`
}

type FavorListReq struct {
	FavoriteId int64 `json:"favoriteId"`
	ItemType   int   `json:"itemType"`
}

type FavorListResp struct {
	List []FavorView `json:"list"`
}

type FavorReq struct {
	FavoriteId int64 `json:"favoriteId"`
	ItemType   int   `json:"itemType"`
	ItemId     int64 `json:"itemId"`
}

type FavorView struct {
	ItemType   int    `json:"itemType"`
	ItemId     int64  `json:"itemId"`
	CoverUrl   string `json:"coverUrl"`
	UserId     int64  `json:"userId"`
	Account    string `json:"account"`
	Title      string `json:"title"`
	LikeCount  int    `json:"likeCount"`
	CreateTime string `json:"createTime"`
}

type FavoriteCreateReq struct {
	Name string `json:"name"`
}

type FavoriteDeleteReq struct {
	Id int64 `json:"id"`
}

type FavoriteListReq struct {
	UserId int64 `json:"userId"`
}

type FavoriteListResp struct {
	List []FavoriteListView `json:"list"`
}

type FavoriteListView struct {
	Id       int64  `json:"id"`
	UserId   int64  `json:"userId"`
	Name     string `json:"name"`
	CoverUrl string `json:"coverUrl"`
}

type HistoryCreateReq struct {
	ItemType int   `json:"itemType"`
	ItemId   int64 `json:"itemId"`
}

type HistoryDeleteReq struct {
	Id int64 `json:"id"`
}

type HistoryListReq struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

type HistoryListResp struct {
	List  []HistoryView `json:"list"`
	Total int           `json:"total"`
}

type HistoryView struct {
	Id         int64  `json:"id"`
	ItemType   int    `json:"itemType"`
	ItemId     int64  `json:"itemId"`
	CoverUrl   string `json:"coverUrl"`
	UserId     int64  `json:"userId"`
	Account    string `json:"account"`
	Title      string `json:"title"`
	LikeCount  int    `json:"likeCount"`
	CreateTime string `json:"createTime"`
}

type LikeReq struct {
	ItemType    int   `json:"itemType"`
	ItemId      int64 `json:"itemId"`
	LikedStatus bool  `json:"likedStatus"`
}

type MessageCreateReq struct {
	UserIds       []int64 `json:"userId"`
	ItemType      int     `json:"itemType"`
	ItemId        int64   `json:"itemId"`
	MessageType   int     `json:"messageType"`
	MessageUserId int64   `json:"messageUserId"`
	Content       string  `json:"content"`
}

type MessageDeleteReq struct {
	Id int64 `json:"id"`
}

type MessageListResp struct {
	List []MessageView `json:"list"`
}

type MessageUpdateReq struct {
	Id            int64 `json:"id"`
	MessageStatus bool  `json:"messageStatus"`
}

type MessageView struct {
	Id            int64  `json:"id"`
	ItemType      int    `json:"itemType"`
	ItemId        int    `json:"itemId"`
	CoverUrl      string `json:"coverUrl"`
	Title         string `json:"title"`
	MessageType   int    `json:"messageType"`
	MessageStatus bool   `json:"messageStatus"`
	Content       string `json:"content"`
	MessageUserId int64  `json:"messageUserId"`
	Account       string `json:"account"`
}

type UserHomeContentListReq struct {
	UserId   int64 `json:"userId"`
	ItemType int   `json:"itemType"`
	SortType int   `json:"sortType"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

type UserHomeContentListResp struct {
	List  []ContentSimpleView `json:"list"`
	Total int                 `json:"total"`
}

type UserHomeDynamicListReq struct {
	UserId   int64 `json:"userId"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

type UserHomeDynamicListResp struct {
	List  []CommunityDynamicView `json:"list"`
	Total int                    `json:"total"'`
}

type UserHomeListReq struct {
	UserId int64 `json:"userId"`
}

type UserHomeListResp struct {
	RecentArticleList    []ContentSimpleView `json:"recentArticleList"`
	RecentVideoList      []ContentSimpleView `json:"recentVideoList"`
	RecommendArticleList []ContentSimpleView `json:"recommendArticleList"`
	RecommendVideoList   []ContentSimpleView `json:"RecommendVideoList"`
}

type UserInfoView struct {
	Id      int64  `json:"id"`
	Account string `json:"account"`
	Avatar  string `json:"avatar"`
	Email   string `json:"email"`
}
