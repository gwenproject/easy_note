package rpc

import (
    "context"
    "github.com/cloudwego/kitex/client"
    "github.com/cloudwego/kitex/pkg/retry"
    "github.com/gwen0x4c3/easy_note/kitex_gen/kuser"
    "github.com/gwen0x4c3/easy_note/kitex_gen/kuser/userservice"
    "github.com/gwen0x4c3/easy_note/pkg/constants"
    "github.com/gwen0x4c3/easy_note/pkg/errno"
    etcd "github.com/kitex-contrib/registry-etcd"
    "time"
)

var userClient userservice.Client

func InitUserRPC() {
    r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
    if err != nil {
        panic(err)
    }
    userClient, err = userservice.NewClient(
        constants.UserServiceName,
        client.WithResolver(r),
        //client.WithMuxConnection(1),
        client.WithRPCTimeout(3*time.Second),
        client.WithConnectTimeout(100*time.Millisecond),
        client.WithFailureRetry(retry.NewFailurePolicy()),
    )
    if err != nil {
        panic(err)
    }
}

func MGetUser(ctx context.Context, req *kuser.MGetUserRequest) (map[int64]*kuser.User, error) {
    resp, err := userClient.MGetUser(ctx, req)
    if err != nil {
        return nil, err
    }
    if resp.BaseResp.StatusCode != errno.SuccessCode {
        return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
    }
    res := make(map[int64]*kuser.User)
    for _, u := range resp.Users {
        res[u.UserId] = u
    }
    return res, nil
}
