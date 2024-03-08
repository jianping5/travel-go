package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"travel/app/data/cmd/rpc/data"
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
	contentType := req.ContentType
	offset := (req.PageNum - 1) * req.PageSize
	var total int64
	var contents []types.ContentView
	switch enum.ContentType(contentType) {
	case enum.AllContent:
		tx := l.svcCtx.DB.Model(&model.Content{})
		// 记录总数
		countTx := tx
		countTx.Count(&total)

		tx.Offset(offset).Limit(req.PageSize).Scan(&contents)
		l.SetBasicInfo(&contents, loginUserId)

		break
	case enum.ForYouContent:
		userLikeContent, _ := l.svcCtx.DataRpc.UserLikeContent(l.ctx, &data.UserLikeContentReq{
			UserId:   loginUserId,
			PageNum:  int32(req.PageNum),
			PageSize: int32(req.PageSize),
		})
		itemIds := userLikeContent.ItemIds
		total = userLikeContent.Total
		l.svcCtx.DB.Model(&model.Content{}).Where("id IN (?)", itemIds).Scan(&contents)
		l.SetBasicInfo(&contents, loginUserId)

		break
	case enum.RecentContent:
		tx := l.svcCtx.DB.Model(&model.Content{}).Order("createTime DESC")
		// 记录总数
		countTx := tx
		countTx.Count(&total)

		tx.Offset(offset).Limit(req.PageSize).Scan(&contents)
		l.SetBasicInfo(&contents, loginUserId)

		break
	case enum.ArticleContent:
		// todo：文章
		tx := l.svcCtx.DB.Model(&model.Content{}).Where("itemType = ?", enum.ARTICLE)
		// 记录总数
		countTx := tx
		countTx.Count(&total)

		tx.Offset(offset).Limit(req.PageSize).Scan(&contents)
		l.SetBasicInfo(&contents, loginUserId)

		break
	case enum.VideoContent:
		// 视频
		tx := l.svcCtx.DB.Model(&model.Content{}).Where("itemType = ?", enum.VIDEO)
		// 记录总数
		countTx := tx
		countTx.Count(&total)

		tx.Offset(offset).Limit(req.PageSize).Scan(&contents)
		l.SetBasicInfo(&contents, loginUserId)

		break
	default:
		return nil, errors.Wrap(xerr.NewErrMsg("参数错误"), "参数错误")
	}
	return &types.ContentListResp{
		List:  contents,
		Total: int(total),
	}, nil
}

func (l *ContentListLogic) SetBasicInfo(contents *[]types.ContentView, loginUserId int64) {
	for i, a := range *contents {
		// 用户信息
		var userInfoView types.UserInfoView
		userId := a.UserId
		info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: userId})
		_ = copier.Copy(&userInfoView, &info)
		(*contents)[i].UserInfo = userInfoView

		// 未登录
		if loginUserId == 0 {
			(*contents)[i].IsFavored = false
			(*contents)[i].IsLiked = false
		}

		// 是否点赞
		var isLiked bool
		l.svcCtx.DB.Model(&model.Like{}).Select("likedStatus").Where("userId = ? and itemId = ?", loginUserId, a.Id).Scan(&isLiked)
		(*contents)[i].IsLiked = isLiked

		// 是否收藏
		var favor model.Favor
		if err := l.svcCtx.DB.Model(&model.Favor{}).Where("userId = ? and itemType = ? and itemId = ?", loginUserId, a.Id).First(&favor).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				(*contents)[i].IsFavored = false
			}
		} else {
			(*contents)[i].IsFavored = true
		}
	}
}
