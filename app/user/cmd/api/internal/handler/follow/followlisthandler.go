package follow

import (
	"net/http"
	"travel/app/user/cmd/api/internal/logic/follow"
	"travel/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"travel/app/user/cmd/api/internal/svc"
	"travel/app/user/cmd/api/internal/types"
)

func FollowListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := follow.NewFollowListLogic(r.Context(), svcCtx)
		resp, err := l.FollowList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
