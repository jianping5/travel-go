package handler

import (
	"net/http"
	"travel/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"travel/app/social/cmd/api/internal/logic"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
)

func FavoriteDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewFavoriteDetailLogic(r.Context(), svcCtx)
		resp, err := l.FavoriteDetail(&req)
		result.HttpResult(r, w, resp, err)
	}
}
