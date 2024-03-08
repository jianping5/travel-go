package logic

import (
	"context"
	"travel/app/user/cmd/model"
	"travel/common/enum"

	"travel/app/user/cmd/rpc/internal/svc"
	"travel/app/user/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLogic {
	return &SearchUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserLogic) SearchUser(in *pb.SearchUserReq) (*pb.SearchUserResp, error) {
	var total int64
	l.svcCtx.DB.Model(&model.User{}).Where("account like ?", "%"+in.Keyword+"%").Count(&total)
	var users []*pb.UserInfoView
	switch enum.SortType(in.SortType) {
	case enum.Newest:
		l.svcCtx.DB.Model(&model.User{}).Offset(int(in.Offset)).Limit(int(in.PageSize)).Order("createTime DESC").
			Where("account like ?", "%"+in.Keyword+"%").Scan(&users)
		break
	case enum.Oldest:
		l.svcCtx.DB.Model(&model.User{}).Offset(int(in.Offset)).Limit(int(in.PageSize)).Order("createTime ASC").
			Where("account like ?", "%"+in.Keyword+"%").Scan(&users)
		break
	default:
		break
	}

	return &pb.SearchUserResp{
		Users: users,
		Total: int32(total),
	}, nil
}
