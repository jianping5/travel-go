package cmd

import (
	"context"
	"encoding/json"
	"travel/app/data/cmd/rpc/data"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/enum"
)

func UpdateUserTag() {
	var userIds []int64
	// 0. 获取所有用户 id
	userIdsResp, _ := svcCtx.UserRpc.GetUserIds(context.Background(), &user.GetUserIdsReq{})
	userIds = userIdsResp.UserIds

	// 遍历每个用户
	for _, userId := range userIds {
		// 1. 分析用户行为，获取 itemIds
		var itemIds []int64
		behavior, _ := svcCtx.DataRpc.AnalyzeUserBehavior(context.Background(), &data.AnalyzeUserBehaviorReq{
			UserId:   userId,
			ItemType: int64(enum.VIDEO),
		})
		itemIds = behavior.ItemIds

		// 2. 根据 itemIds 获取对应的标签数组
		var tagJsons []string
		svcCtx.DB.Model(model.Content{}).Select("tag").Where("id IN (?)", itemIds).Scan(&tagJsons)
		mp := map[string]bool{}
		for _, tagJson := range tagJsons {
			var tmpTags []string
			_ = json.Unmarshal([]byte(tagJson), &tmpTags)
			for _, tmpTag := range tmpTags {
				mp[tmpTag] = true
			}
		}

		// 3. 对标签名进行去重，并更新到 user_tag 表中
		var tags []string
		for tag := range mp {
			tags = append(tags, tag)
		}
		// 转换成 json
		tagJson, _ := json.Marshal(tags)
		// 更新
		svcCtx.DataRpc.UpdateUserTag(context.Background(), &data.UpdateUserTagReq{
			UserId:  userId,
			TagJson: string(tagJson),
		})
	}
}
