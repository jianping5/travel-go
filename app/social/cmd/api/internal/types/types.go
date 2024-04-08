// Code generated by goctl. DO NOT EDIT.
package types

type CommentCreateReq struct {
	CommentItemId   int64  `json:"commentItemId"`
	CommentItemType int    `json:"commentItemType"`
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
	LikeCount       int          `json:"likeCount"`
	ReplyCount      int          `json:"replyCount"`
	CreateTime      string       `json:"createTime"`
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
	Description string `json:"description"`
	Content     string `json:"content"`
	FileType    int    `json:"fileType"`
}

type CommunityDynamicDeleteReq struct {
	Id int64 `json:"id"`
}

type CommunityDynamicDetailReq struct {
	Id int64 `json:"id"`
}

type CommunityDynamicDetailResp struct {
	DynamicDetail CommunityDynamicView `json:"dynamicDetail"`
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

type CommunityDynamicSpecificListReq struct {
	SortType    int   `json:"sortType"`
	CommunityId int64 `json:"communityId"`
	PageNum     int   `json:"pageNum"`
	PageSize    int   `json:"pageSize"`
}

type CommunityDynamicSpecificListResp struct {
	List  []CommunityDynamicView `json:"list"`
	Total int                    `json:"total"`
}

type CommunityDynamicView struct {
	Id           int64         `json:"id"`
	UserInfo     UserInfoView  `json:"userInfo"`
	CommunityId  int64         `json:"communityId"`
	Community    CommunityView `json:"communityView"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Content      string        `json:"content"`
	FileType     int           `json:"fileType"`
	LikeCount    int           `json:"likeCount"`
	CommentCount int           `json:"commentCount"`
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
	Description string `json:"description,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
}

type CommunityView struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
	MemberCount int    `json:"memberCount"`
	CreateTime  string `json:"createTime"`
	IsJoined    bool   `json:"isJoined"`
}

type ContentCreateReq struct {
	ItemType    int      `json:"itemType"`
	Title       string   `json:"title"`
	CoverUrl    string   `json:"coverUrl"`
	Content     string   `json:"content"`
	Description string   `json:"description"`
	Tag         []string `json:"tag"`
}

type ContentDeleteReq struct {
	Id int64 `json:"id"`
}

type ContentDetailReq struct {
	Id int64 `json:"id"`
}

type ContentDetailResp struct {
	ContentDetail ContentView `json:"contentDetail"`
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

type ContentSimilarReq struct {
	ItemType int   `json:"itemType"`
	ItemId   int64 `json:"itemId"`
}

type ContentSimilarResp struct {
	List []ContentView `json:"list"`
}

type ContentSimpleView struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	CoverUrl   string `json:"coverUrl"`
	LikeCount  int    `json:"likeCount"`
	CreateTime string `json:"CreateTime"`
}

type ContentUpdateReq struct {
	Id          int64    `json:"id"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	CoverUrl    string   `json:"coverUrl"`
	Description string   `json:"description"`
	Tag         []string `json:"tag"`
}

type ContentView struct {
	Id           int64        `json:"id"`
	UserId       int64        `json:"userId"`
	UserInfo     UserInfoView `json:"userInfo"`
	ItemType     int          `json:"itemType"`
	Title        string       `json:"title"`
	CoverUrl     string       `json:"coverUrl"`
	Content      string       `json:"content"`
	Description  string       `json:"description"`
	Tag          []string     `json:"tag"`
	LikeCount    int          `json:"likeCount"`
	CommentCount int          `json:"commentCount"`
	FavorCount   int          `json:"favorCount"`
	CreateTime   string       `json:"createTime"`
	IsLiked      bool         `json:"isLiked"`
	IsFavored    bool         `json:"isFavored"`
}

type CopyrightCreateReq struct {
	ItemType     int    `json:"itemType"`
	ItemId       int64  `json:"itemId"`
	ContentUrl   string `json:"contentUrl"`
	UploadSwitch bool   `json:"uploadSwitch"`
}

type CopyrightDetailReq struct {
	Id int64 `json:"id"`
}

type CopyrightDetailResp struct {
	Copyright CopyrightView `json:"copyright"`
	UserInfo  UserInfoView  `json:"userInfo"`
}

type CopyrightListReq struct {
	UserId int64 `json:"userId"`
}

type CopyrightListResp struct {
	List []CopyrightView `json:"list"`
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
	Title      string `json:"title"`
	CoverUrl   string `json:"coverUrl"`
	Account    string `json:"account"`
	Avatar     string `json:"avatar"`
}

type FavorCancelReq struct {
	FavoriteId int64 `json:"favoriteId"`
	ItemId     int64 `json:"itemId"`
}

type FavorDeleteReq struct {
	Id int64 `json:"id"`
}

type FavorListReq struct {
	FavoriteId int64 `json:"favoriteId"`
}

type FavorListResp struct {
	List []FavorView `json:"list"`
}

type FavorReq struct {
	FavoriteId int64 `json:"favoriteId"`
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
	ItemId int64 `json:"itemId"`
}

type FavoriteListResp struct {
	List []FavoriteListView `json:"list"`
}

type FavoriteListView struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"userId"`
	Name      string `json:"name"`
	CoverUrl  string `json:"coverUrl"`
	IsFavored bool   `json:"isFavored"`
}

type HistoryCreateReq struct {
	ItemId int64 `json:"itemId"`
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

type SearchReq struct {
	Keyword  string `json:"keyword"`
	ItemType int    `json:"itemType"`
	SortType int    `json:"sortType"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
}

type SearchResp struct {
	ContentList   []ContentView          `json:"contentList"`
	CommunityList []CommunityView        `json:"communityList"`
	UserList      []UserInfoView         `json:"userList"`
	DynamicList   []CommunityDynamicView `json:"dynamicList"`
	CopyrightList []CopyrightView        `json:"copyrightList"`
	Total         int                    `json:"total"`
}

type UserHomeContentListReq struct {
	UserId   int64 `json:"userId"`
	ItemType int   `json:"itemType"`
	SortType int   `json:"sortType"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

type UserHomeContentListResp struct {
	List  []ContentView `json:"list"`
	Total int           `json:"total"`
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
	RecentArticleList    []ContentView `json:"recentArticleList"`
	RecentVideoList      []ContentView `json:"recentVideoList"`
	RecommendArticleList []ContentView `json:"recommendArticleList"`
	RecommendVideoList   []ContentView `json:"RecommendVideoList"`
}

type UserInfoView struct {
	Id         int64  `json:"id"`
	Account    string `json:"account"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	IsFollowed bool   `json:"isFollowed"`
}
