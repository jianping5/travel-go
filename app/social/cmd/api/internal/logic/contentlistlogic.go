package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"
	"travel/common/enum"
	"travel/common/xerr"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentListLogic {
	return &ContentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContentListLogic) ContentList(req *types.ContentListReq) (resp *types.ContentListResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	// TODO：根据类型分类查询
	itemType := req.ItemType
	offset := (req.PageNum - 1) * req.PageSize
	var total int64
	var contents []types.ContentView
	switch enum.FileType(itemType) {
	case enum.Text:
		// todo：文章
		tx := l.svcCtx.DB.Model(&model.Article{})
		// 记录总数
		countTx := tx
		countTx.Count(&total)

		tx.Offset(offset).Limit(req.PageSize).Scan(&contents)

		// 用户信息
		for i, a := range contents {
			var userInfoView types.UserInfoView
			userId := a.UserId
			info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: userId})
			_ = copier.Copy(&userInfoView, &info)
			contents[i].UserInfo = userInfoView

			// 是否点赞
			var isLiked bool
			l.svcCtx.DB.Model(&model.Like{}).Select("likedStatus").Where("userId = ? and itemType = ? and itemId = ?", loginUserId, enum.ARTICLE, a.Id).Scan(&isLiked)
			contents[i].IsLiked = isLiked
		}
		break
	case enum.Video:
		// 视频
		tx := l.svcCtx.DB.Model(&model.Video{})
		// 记录总数
		countTx := tx
		countTx.Count(&total)

		tx.Offset(offset).Limit(req.PageSize).Scan(&contents)

		// 用户信息
		for i, v := range contents {
			var userInfoView types.UserInfoView
			userId := v.UserId
			info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: userId})
			_ = copier.Copy(&userInfoView, &info)
			contents[i].UserInfo = userInfoView

			// 是否点赞
			var isLiked bool
			l.svcCtx.DB.Model(&model.Like{}).Select("likedStatus").Where("userId = ? and itemType = ? and itemId = ?", loginUserId, enum.Video, v.Id).Scan(&isLiked)
			contents[i].IsLiked = isLiked
		}
		break
	default:
		return nil, errors.Wrap(xerr.NewErrMsg("参数错误"), "参数错误")
	}
	return &types.ContentListResp{
		List:  contents,
		Total: int(total),
	}, nil
}
