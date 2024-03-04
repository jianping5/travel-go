package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"
	"travel/common/enum"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListReq) (resp *types.CommentListResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	offset := (req.PageNum - 1) * req.PageSize
	var total int64
	var topComments []types.CommentView
	tx := l.svcCtx.DB.Model(&model.Comment{}).
		Where("commentItemType = ? and commentItemId = ? and topId = 0", req.CommentItemType, req.CommentItemId)
	// 记录总数
	countTx := tx
	countTx.Count(&total)

	// 顶级评论列表
	tx.Offset(offset).Limit(req.PageSize).Scan(&topComments)
	l.SetUserInfo(loginUserId, &topComments)

	// 注入回复
	var commentListViews []types.CommentListView
	for _, c := range topComments {
		var commentListView types.CommentListView
		var comments []types.CommentView
		l.svcCtx.DB.Model(&model.Comment{}).Where("topId = ?", c.TopId).Scan(&comments)
		l.SetUserInfo(loginUserId, &comments)
		commentListView.TopComment = c
		commentListView.CommentList = comments
		commentListViews = append(commentListViews, commentListView)
	}

	return &types.CommentListResp{
		List:  commentListViews,
		Total: int(total),
	}, nil
}

func (l *CommentListLogic) SetUserInfo(loginUserId int64, comments *[]types.CommentView) {
	for i, c := range *comments {
		var userInfoView types.UserInfoView
		var parentUserInfoView types.UserInfoView
		userInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: c.UserId})
		parentUserInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: c.ParentUserId})
		_ = copier.Copy(&userInfoView, userInfo)
		_ = copier.Copy(&parentUserInfoView, parentUserInfo)
		(*comments)[i].UserInfo = userInfoView
		(*comments)[i].ParentUserInfo = parentUserInfoView

		// 是否点赞
		var isLiked bool
		l.svcCtx.DB.Model(&model.Like{}).Select("likedStatus").Where("userId = ? and itemType = ? and itemId = ?", loginUserId, enum.COMMENT, c.Id).Scan(&isLiked)
		(*comments)[i].IsLiked = isLiked
	}
}
