package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"travel/app/data/cmd/model"

	"travel/app/data/cmd/rpc/internal/svc"
	"travel/app/data/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserTagLogic {
	return &UpdateUserTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserTagLogic) UpdateUserTag(in *pb.UpdateUserTagReq) (*pb.UpdateUserTagResp, error) {
	userId := in.UserId
	tagJson := in.TagJson
	err := l.svcCtx.DB.Model(model.UserTag{}).Select("id").Where("user_id = ?", userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var userTag model.UserTag
			userTag.UserId = userId
			userTag.Tag = datatypes.JSON(tagJson)
			l.svcCtx.DB.Create(&userTag)
		}
	} else {
		l.svcCtx.DB.Model(model.UserTag{}).Where("user_id = ?", userId).Update("tag", tagJson)
	}

	return &pb.UpdateUserTagResp{}, nil
}
