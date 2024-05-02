package handler

import (
	"net/http"
	"travel/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"travel/app/social/cmd/api/internal/logic"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
)

func CopyrightMintHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CopyrightMintReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewCopyrightMintLogic(r.Context(), svcCtx)
		err := l.CopyrightMint(&req)
		result.HttpResult(r, w, nil, err)
	}
}
