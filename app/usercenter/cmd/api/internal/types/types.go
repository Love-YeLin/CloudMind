// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id          int64   `json:"id"`          // 用户ID
	Email       string  `json:"email"`       // 邮箱号
	NickName    string  `json:"nickName"`    // 用户名
	Sex         int64   `json:"sex"`         // 性别
	Avatar      string  `json:"avatar"`      // 用户头像
	Name        string  `json:"name"`        // 真实姓名
	IdCard      string  `json:"idCard"`      // 身份证号
	Create_time int64   `json:"create_time"` // 创建账号的时间
	Update_time int64   `json:"update_time"` // 上次修改用户名的时间
	Memory      float64 `json:"memory"`      // 空间
	Flow        float64 `json:"flow"`        // 流量
	Money       float64 `json:"money"`       // 余额
}

type Post struct {
	Title  string `json:"title"`  // 帖子标题
	Avatar string `json:"avatar"` // 帖子图片
	Id     int64  `json:"id"`     // 帖子ID
}

type File struct {
	Title  string `json:"title"`  // 文件标题
	Avatar string `json:"avatar"` // 文件图片
	Id     int64  `json:"id"`     // 文件ID
}

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

type GetUserInfoReq struct {
}

type GetUserInfoResp struct {
	UserInfo User `json:"userInfo"`
}

type UpdateNickNameReq struct {
	NickName string `json:"nickName"`
}

type UpdateNickNameResp struct {
}

type UpdatePassWordReq struct {
	PassWord string `json:"passWord"`
}

type UpdatePassWordResp struct {
}

type UpdateSexReq struct {
	Sex int64 `json:"sex"`
}

type UpdateSexResp struct {
}

type UpdateAvatarReq struct {
	File []byte `form:"file"`
}

type UpdateAvatarResp struct {
}

type RealNameAuthenticationReq struct {
	Name   string `json:"name"`
	IdCard string `json:"idCard"`
}

type RealNameAuthenticationResp struct {
}

type UpdateStarReq struct {
	Id int64 `json:"id"`
}

type UpdateStarResp struct {
}

type LogoutReq struct {
}

type LogoutResp struct {
}
