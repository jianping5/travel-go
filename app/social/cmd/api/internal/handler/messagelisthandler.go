package handler

import (
	"net/http"
	"travel/common/result"

	"travel/app/social/cmd/api/internal/logic"
	"travel/app/social/cmd/api/internal/svc"
)

func MessageListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewMessageListLogic(r.Context(), svcCtx)
		resp, err := l.MessageList()
		result.HttpResult(r, w, resp, err)
	}
}
