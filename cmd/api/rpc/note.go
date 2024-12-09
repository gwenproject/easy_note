package rpc

import (
    "context"
    "github.com/gwen0x4c3/easy_note/kitex_gen/knote"
    "github.com/gwen0x4c3/easy_note/pkg/errno"
    "time"

    "github.com/cloudwego/kitex/client"
    "github.com/cloudwego/kitex/pkg/retry"
    "github.com/gwen0x4c3/easy_note/kitex_gen/knote/noteservice"
    "github.com/gwen0x4c3/easy_note/pkg/constants"
    etcd "github.com/kitex-contrib/registry-etcd"
)

var noteClient noteservice.Client

func InitNoteRPC() {
    var err error
    r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
    if err != nil {
        panic(err)
    }
    noteClient, err = noteservice.NewClient(
        constants.NoteServiceName,
        client.WithResolver(r),
        // client.WithMuxConnection(1),                       // mux
        client.WithRPCTimeout(3*time.Second),              // rpc timeout
        client.WithConnectTimeout(100*time.Millisecond),   // conn timeout
        client.WithFailureRetry(retry.NewFailurePolicy()), // retry
    )
    if err != nil {
        panic(err)
    }
}

func QueryNote(ctx context.Context, req *knote.QueryNoteRequest) ([]*knote.Note, int64, error) {
    resp, err := noteClient.QueryNote(ctx, req)
    if err != nil {
        return nil, 0, err
    }
    if resp.BaseResp.StatusCode != errno.SuccessCode {
        return nil, 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
    }
    return resp.Notes, resp.Total, nil
}

func CreateNote(ctx context.Context, req *knote.CreateNoteRequest) error {
    resp, err := noteClient.CreateNote(ctx, req)
    if err != nil {
        return err
    }
    if resp.BaseResp.StatusCode != errno.SuccessCode {
        return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
    }
    return nil
}
