package handler

import (
	"net/http"
	"travel/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"travel/app/trade/cmd/api/internal/logic"
	"travel/app/trade/cmd/api/internal/svc"
	"travel/app/trade/cmd/api/internal/types"
)

func RecordCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecordCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewRecordCreateLogic(r.Context(), svcCtx)
		err := l.RecordCreate(&req)
		result.HttpResult(r, w, nil, err)
	}
}
