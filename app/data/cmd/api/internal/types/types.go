// Code generated by goctl. DO NOT EDIT.
package types

type BehaviorCreateReq struct {
	BehaviorItemType int   `json:"behavoirItemType"`
	BehaviorItemId   int64 `json:"behaviorItemId"`
	BehaviorType     int   `json:"behaviorType"`
}

type UploadResp struct {
	FileUrl string `json:"fileUrl"`
}
