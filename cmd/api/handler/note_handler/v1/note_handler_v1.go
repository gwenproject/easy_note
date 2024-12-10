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
    "strconv"
)

func QueryNote(ctx context.Context, c *app.RequestContext) {
    userId := middleware.AuthMiddleware.IdentityHandler(ctx, c).(int64)
    hlog.Info("userId: ", userId)
    var queryVar struct {
        Current  int64  `json:"current"`
        PageSize int64  `json:"pageSize"`
        Keyword  string `json:"keyword"`
    }

    if err := c.Bind(&queryVar); err != nil {
        handler.SendResponse(c, errno.ConvertErr(err), nil)
        return
    }

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

func CreateNote(ctx context.Context, c *app.RequestContext) {
    userId := middleware.AuthMiddleware.IdentityHandler(ctx, c).(int64)
    req := new(knote.CreateNoteRequest)
    if err := c.Bind(req); err != nil {
        handler.SendResponse(c, errno.ConvertErr(err), nil)
        return
    }
    req.UserId = userId
    if len(req.Title) == 0 {
        handler.SendResponse(c, errno.ParamErr.WithMessage("标题不能为空"), nil)
        return
    }
    if len(req.Content) == 0 {
        handler.SendResponse(c, errno.ParamErr.WithMessage("内容不能为空"), nil)
        return
    }

    err := rpc.CreateNote(context.Background(), req)
    if err != nil {
        handler.SendResponse(c, errno.ConvertErr(err), nil)
        return
    }
    handler.SendResponse(c, nil, nil)
}

func UpdateNote(ctx context.Context, c *app.RequestContext) {
    noteId := c.Param("note_id")
    req := new(knote.UpdateNoteRequest)
    if err := c.Bind(req); err != nil {
        handler.SendResponse(c, errno.ParamErr.WithMessage("参数错误"), nil)
        return
    }
    if req.Title != nil && len(*req.Title) == 0 {
        handler.SendResponse(c, errno.ParamErr.WithMessage("标题不能为空"), nil)
        return
    }
    if req.Content != nil && len(*req.Content) == 0 {
        handler.SendResponse(c, errno.ParamErr.WithMessage("内容不能为空"), nil)
        return
    }
    noteIdInt, err := strconv.ParseInt(noteId, 10, 64)
    if err != nil {
        handler.SendResponse(c, errno.ParamErr.WithMessage("路由参数错误"), nil)
        return
    }
    req.NoteId = noteIdInt
    if err := rpc.UpdateNote(context.Background(), req); err != nil {
        handler.SendResponse(c, errno.ConvertErr(err), nil)
        return
    }
    handler.SendResponse(c, nil, nil)
}

func DeleteNote(ctx context.Context, c *app.RequestContext) {
    noteId := c.Param("note_id")
    noteIdInt, err := strconv.ParseInt(noteId, 10, 64)
    if err != nil {
        handler.SendResponse(c, errno.ParamErr.WithMessage("路由参数错误"), nil)
        return
    }
    req := new(knote.DeleteNoteRequest)
    req.NoteId = noteIdInt
    if err := rpc.DeleteNote(context.Background(), req); err != nil {
        handler.SendResponse(c, errno.ConvertErr(err), nil)
        return
    }
    handler.SendResponse(c, nil, nil)
}
