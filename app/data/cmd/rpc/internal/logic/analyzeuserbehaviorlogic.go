package logic

import (
	"context"
	"travel/app/data/cmd/model"

	"travel/app/data/cmd/rpc/internal/svc"
	"travel/app/data/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnalyzeUserBehaviorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnalyzeUserBehaviorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalyzeUserBehaviorLogic {
	return &AnalyzeUserBehaviorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AnalyzeUserBehaviorLogic) AnalyzeUserBehavior(in *pb.AnalyzeUserBehaviorReq) (*pb.AnalyzeUserBehaviorResp, error) {
	userId := in.UserId
	itemType := in.ItemType
	var itemIds []int64
	// 最近 n 个，n 可调整，目前为 10
	l.svcCtx.DB.Model(model.Behavior{}).Select("behavior_item_id").Order("create_time DESC").Limit(10).Where("user_id = ? and behavior_item_type = ?", userId, itemType).Scan(&itemIds)

	return &pb.AnalyzeUserBehaviorResp{
		ItemIds: itemIds,
	}, nil
}
