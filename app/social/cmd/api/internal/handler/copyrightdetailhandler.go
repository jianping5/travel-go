package handler

import (
	"net/http"
	"travel/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"travel/app/social/cmd/api/internal/logic"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
)

func CopyrightDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CopyrightDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewCopyrightDetailReqLogic(r.Context(), svcCtx)
		resp, err := l.CopyrightDetailReq(&req)
		result.HttpResult(r, w, resp, err)
	}
}
