package handler

import (
	"net/http"
	"travel/common/result"

	"travel/app/intelligence/cmd/api/internal/logic"
	"travel/app/intelligence/cmd/api/internal/svc"
)

func ConversationListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewConversationListLogic(r.Context(), svcCtx)
		resp, err := l.ConversationList()
		result.HttpResult(r, w, resp, err)
	}
}
