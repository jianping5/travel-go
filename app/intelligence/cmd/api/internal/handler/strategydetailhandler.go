package handler

import (
	"net/http"
	"travel/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"travel/app/intelligence/cmd/api/internal/logic"
	"travel/app/intelligence/cmd/api/internal/svc"
	"travel/app/intelligence/cmd/api/internal/types"
)

func StrategyDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StrategyDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewStrategyDetailLogic(r.Context(), svcCtx)
		resp, err := l.StrategyDetail(&req)
		result.HttpResult(r, w, resp, err)
	}
}
