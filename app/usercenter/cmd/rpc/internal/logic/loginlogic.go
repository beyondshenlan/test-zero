package logic

import (
	"context"
	"fmt"
	"time"

	"test-zero/app/usercenter/cmd/rpc/internal/svc"
	"test-zero/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	fmt.Println("hello login me")
	return &pb.LoginResp{
		AccessToken:  "abc123",
		AccessExpire: time.Now().Unix(),
	}, nil
}
