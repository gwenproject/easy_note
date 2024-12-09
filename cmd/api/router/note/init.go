package note

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/gwen0x4c3/easy_note/cmd/api/middleware"
)

func InitNoteV1(v1 *route.RouterGroup) {
	group := v1.Group("/note")
	group.Use(middleware.AuthMiddleware.MiddlewareFunc())
	// group.POST("/query", )
}

// InitNoteV2 to be used in the future
func InitNoteV2() {

}
