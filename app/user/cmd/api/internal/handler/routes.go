// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"travel/app/user/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/user/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/register",
				Handler: RegisterHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/user/follow",
				Handler: FollowHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/follow/list",
				Handler: FollowListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/userinfo",
				Handler: UserInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
