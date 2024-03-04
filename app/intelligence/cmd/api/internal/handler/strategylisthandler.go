package handler

import (
	"net/http"
	"travel/common/result"

	"travel/app/intelligence/cmd/api/internal/logic"
	"travel/app/intelligence/cmd/api/internal/svc"
)

func StrategyListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewStrategyListLogic(r.Context(), svcCtx)
		resp, err := l.StrategyList()
		result.HttpResult(r, w, resp, err)
	}
}
