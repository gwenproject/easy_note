package v1

import (
    "context"
    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/common/hlog"
    "github.com/gwen0x4c3/easy_note/cmd/api/handler"
    "github.com/gwen0x4c3/easy_note/cmd/api/middleware"
    "github.com/gwen0x4c3/easy_note/cmd/api/rpc"
    "github.com/gwen0x4c3/easy_note/kitex_gen/knote"
    "github.com/gwen0x4c3/easy_note/pkg/constants"
    "github.com/gwen0x4c3/easy_note/pkg/errno"
)

func QueryNote(ctx context.Context, c *app.RequestContext) {
    userId := middleware.AuthMiddleware.IdentityHandler(ctx, c).(int64)
    hlog.Info("userId: ", userId)
    var queryVar struct {
        Current  int64  `json:"current",form:"current",query:"current"`
        PageSize int64  `json:"pageSize",form:"pageSize",query:"pageSize"`
        Keyword  string `json:"keyword",form:"keyword",query:"keyword"`
    }

    if err := c.Bind(&queryVar); err != nil {
        handler.SendResponse(c, errno.ConvertErr(err), nil)
        return
    }
    hlog.Info("queryVar: ", queryVar)

    if queryVar.PageSize <= 0 || queryVar.PageSize >= 100 {
        handler.SendResponse(c, errno.ParamErr, nil)
        return
    }

    req := &knote.QueryNoteRequest{
        UserId:   userId,
        Keyword:  &queryVar.Keyword,
        Current:  queryVar.Current,
        PageSize: queryVar.PageSize,
    }

    notes, total, err := rpc.QueryNote(context.Background(), req)
    if err != nil {
        handler.SendResponse(c, errno.ConvertErr(err), nil)
        return
    }
    handler.SendResponse(c, nil, map[string]interface{}{
        constants.Total: total,
        constants.List:  notes,
    })
}
