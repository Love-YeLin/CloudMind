// Code generated by goctl. DO NOT EDIT.
package types

type FeedBack struct {
	FeedbackType string `json:"feedbackType"` // 评分类型: 下载，点赞，收藏，浏览
	UserId       string `json:"userId"`       // 用户Id
	ItemId       string `json:"itemId"`       // 物品Id
	Timestamp    string `json:"timestamp"`    // 评分时间
}

type GetRecommendByUserIdReq struct {
	Number int64 `path:"number"`
}

type GetRecommendByUserIdResp struct {
	ItemIds []string `json:"itemIds"`
}

type GetRecommendByItemIdReq struct {
	ItemId string `path:"ItemId"`
	Number int64  `path:"number"`
}

type GetRecommendByItemIdResp struct {
	ItemIds []string `json:"itemIds"`
}

type InsertFeedBackReq struct {
	FeedBacks []FeedBack `json:"feedBacks"`
}

type InsertFeedBackResp struct {
}
