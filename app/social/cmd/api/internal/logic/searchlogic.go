package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"
	"travel/common/enum"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (resp *types.SearchResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	offset := (req.PageNum - 1) * req.PageSize
	var total int64
	var search types.SearchResp
	switch enum.ItemType(req.ItemType) {
	case enum.ARTICLE:
		l.svcCtx.DB.Model(&model.Content{}).Where("item_type = ? and title like ?", enum.ARTICLE, "%"+req.Keyword+"%").Count(&total)
		contents, _ := l.getSortedArticleList(req.SortType, offset, req.PageSize, req.Keyword)
		for i, a := range contents {
			// 用户信息
			var userInfoView types.UserInfoView
			userId := a.UserId
			info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: userId, LoginUserId: loginUserId})
			_ = copier.Copy(&userInfoView, &info)
			contents[i].UserInfo = userInfoView

			// 是否点赞
			var isLiked bool
			l.svcCtx.DB.Model(&model.Like{}).Select("liked_status").Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, enum.ARTICLE, a.Id).Scan(&isLiked)
			contents[i].IsLiked = isLiked

			// 是否收藏
			var favor model.Favor
			if err := l.svcCtx.DB.Model(&model.Favor{}).Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, enum.ARTICLE, a.Id).First(&favor).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					contents[i].IsFavored = false
				}
			} else {
				contents[i].IsFavored = true
			}
		}
		search.Total = int(total)
		search.ContentList = contents
		break
	case enum.VIDEO:
		l.svcCtx.DB.Model(&model.Content{}).Where("item_type = ? and title like ?", enum.VIDEO, "%"+req.Keyword+"%").Count(&total)

		fmt.Println(total)

		contents, _ := l.getSortedVideoList(req.SortType, offset, req.PageSize, req.Keyword)
		for i, a := range contents {
			// 用户信息
			var userInfoView types.UserInfoView
			userId := a.UserId
			info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: userId, LoginUserId: loginUserId})
			_ = copier.Copy(&userInfoView, &info)
			contents[i].UserInfo = userInfoView

			// 是否点赞
			var isLiked bool
			l.svcCtx.DB.Model(&model.Like{}).Select("liked_status").Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, enum.VIDEO, a.Id).Scan(&isLiked)
			contents[i].IsLiked = isLiked

			// 是否收藏
			var favor model.Favor
			if err := l.svcCtx.DB.Model(&model.Favor{}).Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, enum.VIDEO, a.Id).First(&favor).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					contents[i].IsFavored = false
				}
			} else {
				contents[i].IsFavored = true
			}
		}
		search.Total = int(total)
		search.ContentList = contents
		break
	case enum.USER:
		searchUser, _ := l.svcCtx.UserRpc.SearchUser(l.ctx, &user.SearchUserReq{
			SortType: int32(req.SortType),
			Offset:   int32(offset),
			PageSize: int32(req.PageSize),
			Keyword:  req.Keyword,
		})
		users := searchUser.Users
		var userInfos []types.UserInfoView
		_ = copier.Copy(&userInfos, &users)
		search.Total = int(searchUser.Total)
		search.UserList = userInfos
		break
	case enum.COMMUNITY:
		l.svcCtx.DB.Model(&model.Community{}).Where("name like ?", "%"+req.Keyword+"%").Count(&total)
		communityList, _ := l.getSortedCommunityList(req.SortType, offset, req.PageSize, req.Keyword)
		search.Total = int(total)
		search.CommunityList = communityList
		break
	case enum.DYNAMIC:
		l.svcCtx.DB.Model(&model.Dynamic{}).Where("title like ?", "%"+req.Keyword+"%").Count(&total)
		dynamics, _ := l.getSortedDynamicList(req.SortType, offset, req.PageSize, req.Keyword)
		for i, dynamic := range dynamics {
			// 用户信息
			var userInfoView types.UserInfoView
			userInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
				Id:          loginUserId,
				LoginUserId: loginUserId,
			})
			_ = copier.Copy(&userInfoView, &userInfo)
			dynamics[i].UserInfo = userInfoView
			// 社区信息
			var communityView types.CommunityView
			l.svcCtx.DB.Take(&model.Community{}).Where("id = ?", dynamic.CommunityId).Scan(&communityView)
			dynamics[i].Community = communityView

			// 是否点赞
			var isLiked bool
			l.svcCtx.DB.Model(&model.Like{}).Select("liked_status").Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, enum.DYNAMIC, dynamic.Id).Scan(&isLiked)
			dynamics[i].IsLiked = isLiked
		}
		search.Total = int(total)
		search.DynamicList = dynamics
		break
	case enum.COPYRIGHT:
		l.svcCtx.DB.Model(&model.Copyright{}).Where("title like ?", "%"+req.Keyword+"%").Count(&total)
		copyrights, _ := l.getSortedCopyrightList(req.SortType, offset, req.PageSize, req.Keyword)
		search.Total = int(total)
		search.CopyrightList = copyrights
		break
	default:
		break
	}

	return &search, nil
}

func (l *SearchLogic) getSortedArticleList(sortType, offset, pageSize int, keyword string) ([]types.ContentView, error) {
	var contents []types.ContentView
	switch enum.SortType(sortType) {
	case enum.Newest:
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("create_time DESC").
			Where("item_type = ? and title like ?", enum.ARTICLE, "%"+keyword+"%").Scan(&contents)
		break
	case enum.Popular:
		// TODO: 考虑评论的正负性
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("like_count+comment_count+favor_count DESC").
			Where("item_type = ? and title like ?", enum.ARTICLE, "%"+keyword+"%").Scan(&contents)
		break
	case enum.Oldest:
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("create_time ASC").
			Where("item_type = ? and title like ?", enum.ARTICLE, "%"+keyword+"%").Scan(&contents)
		break
	default:
		break
	}
	return contents, nil
}

func (l *SearchLogic) getSortedVideoList(sortType, offset, pageSize int, keyword string) ([]types.ContentView, error) {
	var contents []types.ContentView
	switch enum.SortType(sortType) {
	case enum.Newest:
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("create_time DESC").
			Where("item_type = ? and title like ?", enum.VIDEO, "%"+keyword+"%").Scan(&contents)
		break
	case enum.Popular:
		// TODO: 考虑评论的正负性
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("like_count+comment_count+favor_count DESC").
			Where("item_type = ? and title like ?", enum.VIDEO, "%"+keyword+"%").Scan(&contents)
		break
	case enum.Oldest:
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("create_time ASC").
			Where("item_type = ? and title like ?", enum.VIDEO, "%"+keyword+"%").Scan(&contents)
		break
	default:
		break
	}
	return contents, nil
}

func (l *SearchLogic) getSortedCommunityList(sortType, offset, pageSize int, keyword string) ([]types.CommunityView, error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	var communityList []types.CommunityView
	switch enum.SortType(sortType) {
	case enum.Newest:
		l.svcCtx.DB.Model(&model.Community{}).Offset(offset).Limit(pageSize).Order("create_time DESC").
			Where("name like ?", "%"+keyword+"%").Scan(&communityList)
		break
	case enum.Popular:
		// TODO: 考虑评论的正负性
		l.svcCtx.DB.Model(&model.Community{}).Offset(offset).Limit(pageSize).Order("member_count DESC").
			Where("name like ?", "%"+keyword+"%").Scan(&communityList)
		break
	case enum.Oldest:
		l.svcCtx.DB.Model(&model.Community{}).Offset(offset).Limit(pageSize).Order("create_time ASC").
			Where("name like ?", "%"+keyword+"%").Scan(&communityList)
		break
	default:
		break
	}

	for i, c := range communityList {
		// 获取是否已经加入
		var id int64
		if l.svcCtx.DB.Model(&model.UserCommunity{}).Select("id").Where("user_id = ? and community_id = ?", loginUserId, c.Id).Scan(&id); id == 0 {
			communityList[i].IsJoined = false
		} else {
			communityList[i].IsJoined = true
		}
	}

	return communityList, nil
}

func (l *SearchLogic) getSortedDynamicList(sortType, offset, pageSize int, keyword string) ([]types.CommunityDynamicView, error) {
	var dynamics []types.CommunityDynamicView
	switch enum.SortType(sortType) {
	case enum.Newest:
		l.svcCtx.DB.Model(&model.Dynamic{}).Offset(offset).Limit(pageSize).Order("create_time DESC").
			Where("title like ?", "%"+keyword+"%").Scan(&dynamics)
		break
	case enum.Popular:
		// TODO: 考虑评论的正负性
		l.svcCtx.DB.Model(&model.Dynamic{}).Offset(offset).Limit(pageSize).Order("like_count+comment_count DESC").
			Where("title like ?", "%"+keyword+"%").Scan(&dynamics)
		break
	case enum.Oldest:
		l.svcCtx.DB.Model(&model.Dynamic{}).Offset(offset).Limit(pageSize).Order("create_time ASC").
			Where("title like ?", "%"+keyword+"%").Scan(&dynamics)
		break
	default:
		break
	}
	return dynamics, nil
}

func (l *SearchLogic) getSortedCopyrightList(sortType, offset, pageSize int, keyword string) ([]types.CopyrightView, error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	var copyrights []types.CopyrightView
	switch enum.SortType(sortType) {
	case enum.Newest:
		l.svcCtx.DB.Model(&model.Copyright{}).Offset(offset).Limit(pageSize).Order("create_time DESC").
			Where("title like ?", "%"+keyword+"%").Scan(&copyrights)
		break
	case enum.Oldest:
		l.svcCtx.DB.Model(&model.Community{}).Offset(offset).Limit(pageSize).Order("create_time ASC").
			Where("title like ?", "%"+keyword+"%").Scan(&copyrights)
		break
	default:
		break
	}

	for i, c := range copyrights {
		switch enum.ItemType(c.ItemType) {
		case enum.ARTICLE:
			var article model.Content
			l.svcCtx.DB.Model(&model.Content{}).Select("title", "cover_url").Where("id = ?", c.ItemId).Scan(&article)
			copyrights[i].Title = article.Title
			copyrights[i].CoverUrl = article.CoverUrl

			// 用户信息
			info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: c.UserId, LoginUserId: loginUserId})
			copyrights[i].Account = info.Account
			copyrights[i].CoverUrl = info.Avatar
			break
		case enum.VIDEO:
			var video model.Content
			l.svcCtx.DB.Model(&model.Content{}).Select("title", "cover_url").Where("id = ?", c.ItemId).Scan(&video)
			copyrights[i].Title = video.Title
			copyrights[i].CoverUrl = video.CoverUrl
			// 用户信息
			info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: c.UserId, LoginUserId: loginUserId})
			copyrights[i].Account = info.Account
			copyrights[i].CoverUrl = info.Avatar
			break
		default:
			break
		}
	}
	return copyrights, nil
}
