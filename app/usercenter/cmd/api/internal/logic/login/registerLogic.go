package login

import (
	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/common/errorx"
	"context"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	Resp, err := l.svcCtx.UsercenterRpc.Register(l.ctx, &pb.RegisterReq{
		NickName: req.NickName,
		PassWord: req.PassWord,
		Email:    req.Email,
		Code:     req.Code,
		UserAuth: &pb.UserAuth{
			AuthType: "email",
			AuthKey:  req.Email,
		},
	})
	if err != nil {
		return nil, err
	}
	if Resp.Error != "" {
		return nil, errorx.NewDefaultError(Resp.Error)
	}
	var resp types.RegisterResp
	_ = copier.Copy(&resp, Resp)
	return &resp, nil
}
