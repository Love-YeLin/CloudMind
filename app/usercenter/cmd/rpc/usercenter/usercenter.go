// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package usercenter

import (
	"context"

	"CloudMind/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	File                       = pb.File
	GenerateTokenReq           = pb.GenerateTokenReq
	GenerateTokenResp          = pb.GenerateTokenResp
	GetUserAuthByAuthKeyReq    = pb.GetUserAuthByAuthKeyReq
	GetUserAuthByAuthKeyResp   = pb.GetUserAuthByAuthKeyResp
	GetUserAuthByUserIdReq     = pb.GetUserAuthByUserIdReq
	GetUserAuthyUserIdResp     = pb.GetUserAuthyUserIdResp
	GetUserInfoReq             = pb.GetUserInfoReq
	GetUserInfoResp            = pb.GetUserInfoResp
	LoginReq                   = pb.LoginReq
	LoginResp                  = pb.LoginResp
	LogoutReq                  = pb.LogoutReq
	LogoutResp                 = pb.LogoutResp
	Post                       = pb.Post
	QueryFilesHistoryReq       = pb.QueryFilesHistoryReq
	QueryFilesHistoryResp      = pb.QueryFilesHistoryResp
	QueryPostsHistoryReq       = pb.QueryPostsHistoryReq
	QueryPostsHistoryResp      = pb.QueryPostsHistoryResp
	RealNameAuthenticationReq  = pb.RealNameAuthenticationReq
	RealNameAuthenticationResp = pb.RealNameAuthenticationResp
	RegisterReq                = pb.RegisterReq
	RegisterResp               = pb.RegisterResp
	SendEmailReq               = pb.SendEmailReq
	SendEmailResp              = pb.SendEmailResp
	UpdateUserInfoReq          = pb.UpdateUserInfoReq
	UpdateUserInfoResp         = pb.UpdateUserInfoResp
	User                       = pb.User
	UserAuth                   = pb.UserAuth

	Usercenter interface {
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		SendEmail(ctx context.Context, in *SendEmailReq, opts ...grpc.CallOption) (*SendEmailResp, error)
		GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyReq, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResp, error)
		GetUserAuthByUserId(ctx context.Context, in *GetUserAuthByUserIdReq, opts ...grpc.CallOption) (*GetUserAuthyUserIdResp, error)
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
		UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UpdateUserInfoResp, error)
		RealNameAuthentication(ctx context.Context, in *RealNameAuthenticationReq, opts ...grpc.CallOption) (*RealNameAuthenticationResp, error)
		Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error)
	}

	defaultUsercenter struct {
		cli zrpc.Client
	}
)

func NewUsercenter(cli zrpc.Client) Usercenter {
	return &defaultUsercenter{
		cli: cli,
	}
}

func (m *defaultUsercenter) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUsercenter) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUsercenter) SendEmail(ctx context.Context, in *SendEmailReq, opts ...grpc.CallOption) (*SendEmailResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.SendEmail(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyReq, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserAuthByAuthKey(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserAuthByUserId(ctx context.Context, in *GetUserAuthByUserIdReq, opts ...grpc.CallOption) (*GetUserAuthyUserIdResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserAuthByUserId(ctx, in, opts...)
}

func (m *defaultUsercenter) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUsercenter) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UpdateUserInfoResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.UpdateUserInfo(ctx, in, opts...)
}

func (m *defaultUsercenter) RealNameAuthentication(ctx context.Context, in *RealNameAuthenticationReq, opts ...grpc.CallOption) (*RealNameAuthenticationResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.RealNameAuthentication(ctx, in, opts...)
}

func (m *defaultUsercenter) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.Logout(ctx, in, opts...)
}
