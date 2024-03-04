package handler

import (
	"net/http"
	"travel/common/result"

	"travel/app/intelligence/cmd/api/internal/logic"
	"travel/app/intelligence/cmd/api/internal/svc"
)

func ConversationDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewConversationDeleteLogic(r.Context(), svcCtx)
		err := l.ConversationDelete()
		result.HttpResult(r, w, nil, err)
	}
}
