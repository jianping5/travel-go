package logic

import (
	"context"
	"travel/app/social/cmd/model"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CopyrightMintLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopyrightMintLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyrightMintLogic {
	return &CopyrightMintLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopyrightMintLogic) CopyrightMint(req *types.CopyrightMintReq) error {
	// 更新对应版权的 token_id 和 account_address
	l.svcCtx.DB.Model(model.Copyright{}).Where("item_id = ?", req.ItemId).Updates(map[string]interface{}{
		"token_id":        req.TokenId,
		"account_address": req.AccountAddress,
	})

	return nil
}
