package auth

import (
	"net/http"
	"travel/app/user/cmd/api/internal/logic/auth"
	"travel/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"travel/app/user/cmd/api/internal/svc"
	"travel/app/user/cmd/api/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := auth.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		result.HttpResult(r, w, resp, err)
	}
}
