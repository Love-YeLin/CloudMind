// Code generated by goctl. DO NOT EDIT.
package types

type RegisterReq struct {
	NickName string `json:"nickName"`
	PassWord string `json:"passWord"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type RegisterResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type EmailLoginReq struct {
	Email    string `json:"email"`
	PassWord string `json:"passWord"`
}

type EmailLoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type QqLoginReq struct {
}

type QqLoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type WxLoginReq struct {
}

type WxLoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type SendEmailReq struct {
	Email string `json:"email"`
}

type SendEmailResp struct {
}