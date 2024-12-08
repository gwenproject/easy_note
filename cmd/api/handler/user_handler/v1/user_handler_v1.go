package v1

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gwen0x4c3/easy_note/cmd/api/handler"
	"github.com/gwen0x4c3/easy_note/cmd/api/rpc"
	"github.com/gwen0x4c3/easy_note/kitex_gen/kuser"
	"github.com/gwen0x4c3/easy_note/pkg/errno"
)

func Register(ctx context.Context, c *app.RequestContext) {
	var req kuser.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		handler.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	req.UserName = strings.Trim(req.UserName, " ")
	req.Password = strings.Trim(req.Password, " ")
	if len(req.UserName) == 0 || len(req.Password) == 0 {
		handler.SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := rpc.CreateUser(context.Background(), &req)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
