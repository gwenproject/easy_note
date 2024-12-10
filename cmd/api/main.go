package main

import (
    "context"
    "fmt"
    "github.com/gwen0x4c3/easy_note/pkg/tracer"

    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
    "github.com/cloudwego/hertz/pkg/app/server"
    "github.com/cloudwego/hertz/pkg/common/hlog"
    "github.com/cloudwego/hertz/pkg/protocol/consts"
    "github.com/gwen0x4c3/easy_note/cmd/api/router"
    "github.com/gwen0x4c3/easy_note/cmd/api/rpc"
    "github.com/gwen0x4c3/easy_note/pkg/constants"
    "github.com/gwen0x4c3/easy_note/pkg/errno"
)

func Init() {
    tracer.InitJaeger(constants.ApiServiceName)
    rpc.InitRPC()
}

func main() {

    Init()

    hz := server.New(
        server.WithHostPorts(fmt.Sprintf("0.0.0.0:%d", constants.ApiServicePort)),
        server.WithHandleMethodNotAllowed(true),
    )
    router.InitRouter(hz)

    // use recovery
    hz.Use(recovery.Recovery(recovery.WithRecoveryHandler(
        func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
            hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
            c.JSON(consts.StatusInternalServerError, map[string]interface{}{
                "code":    errno.ServiceErrCode,
                "message": "Internal Server Error",
            })
        },
    )))

    err := hz.Run()
    if err != nil {
        panic(err)
    }
}
