package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/gwen0x4c3/easy_note/kitex_gen/kuser"
	"github.com/gwen0x4c3/easy_note/kitex_gen/kuser/userservice"
	"github.com/gwen0x4c3/easy_note/pkg/constants"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func InitUserRpc() {
	var err error
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	userClient, err = userservice.NewClient(
		constants.UserServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(100*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
	)
	if err != nil {
		panic(err)
	}
}

func CreateUser(ctx context.Context, req *kuser.CreateUserRequest) error {
	resp, err := userClient.CreateUser(ctx, req)
	err = CheckError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage, err)
	return err
}

func CheckUser(ctx context.Context, req *kuser.CheckUserRequest) (resp *kuser.CheckUserResponse, err error) {
	resp, err = userClient.CheckUser(ctx, req)
	err = CheckError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage, err)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
