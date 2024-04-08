package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"travel/app/trade/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"

	"travel/app/trade/cmd/api/internal/svc"
	"travel/app/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecordListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecordListLogic {
	return &RecordListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecordListLogic) RecordList(req *types.RecordListReq) (resp *types.RecordListResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	var records []types.RecordView
	l.svcCtx.DB.Model(&model.Record{}).Where("work_id = ?", req.WorkId).Scan(records)

	for i, r := range records {
		var oldUserInfo types.UserInfoView
		oldInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: r.OldUserId, LoginUserId: loginUserId})
		_ = copier.Copy(&oldUserInfo, &oldInfo)
		records[i].OldUserInfo = oldUserInfo

		var newUserInfo types.UserInfoView
		newInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: r.NewUserId, LoginUserId: loginUserId})
		_ = copier.Copy(&newUserInfo, &newInfo)
		records[i].NewUserInfo = newUserInfo
	}

	return &types.RecordListResp{
		List: records,
	}, nil
}
