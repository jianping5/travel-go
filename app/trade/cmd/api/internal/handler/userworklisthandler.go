package handler

import (
	"net/http"
	"travel/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"travel/app/trade/cmd/api/internal/logic"
	"travel/app/trade/cmd/api/internal/svc"
	"travel/app/trade/cmd/api/internal/types"
)

func UserWorkListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserWorkListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewUserWorkListLogic(r.Context(), svcCtx)
		resp, err := l.UserWorkList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
