// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"travel/app/social/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/social/comment/list",
				Handler: CommentListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/community/dynamic/detail",
				Handler: CommunityDynamicDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/community/dynamic/list",
				Handler: CommunityDynamicListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/community/dynamic/specific/list",
				Handler: CommunityDynamicSpecificListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/community/list",
				Handler: CommunityListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/content/detail",
				Handler: ContentDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/content/list",
				Handler: ContentListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/content/similar",
				Handler: ContentSimilarHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/copyright/detail",
				Handler: CopyrightDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/copyright/list",
				Handler: CopyrightListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/favor/list",
				Handler: FavorListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/favorite/detail",
				Handler: FavoriteDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/favorite/list",
				Handler: FavoriteListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/search",
				Handler: SearchHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/userhome/content/list",
				Handler: UserHomeContentListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/userhome/dynamic/list",
				Handler: UserHomeDynamicListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/userhome/list",
				Handler: UserHomeListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/social/comment/create",
				Handler: CommentCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/comment/delete",
				Handler: CommentDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/community/create",
				Handler: CommunityCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/community/delete",
				Handler: CommunityDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/community/detail",
				Handler: CommunityDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/community/dynamic/create",
				Handler: CommunityDynamicCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/community/dynamic/delete",
				Handler: CommunityDynamicDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/community/join",
				Handler: CommunityJoinHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/community/quit",
				Handler: CommunityQuitHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/community/update",
				Handler: CommunityUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/content/create",
				Handler: ContentCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/content/delete",
				Handler: ContentDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/content/update",
				Handler: ContentUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/copyright/create",
				Handler: CopyrightCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/copyright/mint",
				Handler: CopyrightMintHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/favor/cancel",
				Handler: FavorCancelHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/favor/create",
				Handler: FavorHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/favor/delete",
				Handler: FavorDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/favorite/create",
				Handler: FavoriteCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/favorite/delete",
				Handler: FavoriteDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/history/create",
				Handler: HistoryCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/history/delete",
				Handler: HistoryDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/history/list",
				Handler: HistoryListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/like",
				Handler: LikeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/message/create",
				Handler: MessageCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/message/delete",
				Handler: MessageDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/message/list",
				Handler: MessageListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/social/message/update",
				Handler: MessageUpdateHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
