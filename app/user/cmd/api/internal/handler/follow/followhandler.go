package follow

import (
	"net/http"
	"travel/app/user/cmd/api/internal/logic/follow"
	"travel/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"travel/app/user/cmd/api/internal/svc"
	"travel/app/user/cmd/api/internal/types"
)

func FollowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := follow.NewFollowLogic(r.Context(), svcCtx)
		err := l.Follow(&req)
		result.HttpResult(r, w, nil, err)
	}
}
