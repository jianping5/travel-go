// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"travel/app/trade/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/trade/record/create",
				Handler: RecordCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/trade/record/list",
				Handler: RecordListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/trade/work/create",
				Handler: WorkCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/trade/work/detail",
				Handler: WorkDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/trade/work/list",
				Handler: WorkListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/trade/work/update",
				Handler: WorkUpdateHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
