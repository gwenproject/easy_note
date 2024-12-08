package user

import (
	"github.com/cloudwego/hertz/pkg/route"
	v1 "github.com/gwen0x4c3/easy_note/cmd/api/handler/user_handler/v1"
)

func InitUserV1(g *route.RouterGroup) {
	group := g.Group("/user")
	// group.POST("/login")
	group.POST("/register", v1.Register)
}
