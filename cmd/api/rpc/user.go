package rpc

import (
    "context"
    "github.com/cloudwego/hertz/pkg/common/hlog"
    "github.com/gwen0x4c3/easy_note/pkg/errno"
    "time"

    "github.com/cloudwego/kitex/client"
    "github.com/cloudwego/kitex/pkg/retry"
    "github.com/gwen0x4c3/easy_note/kitex_gen/kuser"
    "github.com/gwen0x4c3/easy_note/kitex_gen/kuser/userservice"
    "github.com/gwen0x4c3/easy_note/pkg/constants"
    etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func InitUserRPC() {
    var err error
    r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
    if err != nil {
        panic(err)
    }
    userClient, err = userservice.NewClient(
        constants.UserServiceName,
        client.WithResolver(r),
        // client.WithMuxConnection(1),
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
    if err != nil {
        return err
    }
    if resp.BaseResp.StatusCode != errno.SuccessCode {
        return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
    }
    return nil
}

func CheckUser(ctx context.Context, req *kuser.CheckUserRequest) (int64, error) {
    resp, err := userClient.CheckUser(ctx, req)
    hlog.Infof("CheckUser resp: %v", resp)
    if err != nil {
        return 0, err
    }
    if resp.BaseResp.StatusCode != errno.SuccessCode {
        return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
    }
    return resp.UserId, nil
}
