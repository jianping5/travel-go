package handler

import (
	"net/http"
	"travel/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"travel/app/data/cmd/api/internal/logic"
	"travel/app/data/cmd/api/internal/svc"
	"travel/app/data/cmd/api/internal/types"
)

func ContentTagCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ContentTagCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewContentTagCreateLogic(r.Context(), svcCtx)
		err := l.ContentTagCreate(&req)
		result.HttpResult(r, w, nil, err)
	}
}
