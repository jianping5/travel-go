package logic

import (
	"context"
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
	// todo: add your logic here and delete this line
	userId := in.UserId
	tagJson := in.TagJson
	l.svcCtx.DB.Model(model.UserTag{}).Update("tag", tagJson).Where("user_id = ?", userId)

	return &pb.UpdateUserTagResp{}, nil
}
