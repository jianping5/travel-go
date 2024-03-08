package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"travel/app/data/cmd/model"

	"travel/app/data/cmd/rpc/internal/svc"
	"travel/app/data/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLikeContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLikeContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLikeContentLogic {
	return &UserLikeContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLikeContentLogic) UserLikeContent(in *pb.UserLikeContentReq) (*pb.UserLikeContentResp, error) {
	offset := (in.PageNum - 1) * in.PageSize
	var tags []string
	if err := l.svcCtx.DB.Model(&model.UserTag{}).Select("tag").Where("userId = ?", in.UserId).Scan(&tags).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &pb.UserLikeContentResp{}, nil
		}
	}
	var total int64
	var itemIds []int64
	tx := l.svcCtx.DB.Model(&model.ContentTag{}).Select("id").Where("name IN (?)", tags)

	countTx := tx
	countTx.Count(&total)

	tx.Offset(int(offset)).Limit(int(in.PageSize)).Scan(&itemIds)

	return &pb.UserLikeContentResp{
		ItemIds: itemIds,
		Total:   total,
	}, nil
}
