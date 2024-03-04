package handler

import (
	"net/http"
	"travel/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"travel/app/social/cmd/api/internal/logic"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
)

func CommunityCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommunityCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewCommunityCreateLogic(r.Context(), svcCtx)
		err := l.CommunityCreate(&req)
		result.HttpResult(r, w, nil, err)
	}
}
