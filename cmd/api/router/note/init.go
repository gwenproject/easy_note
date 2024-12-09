package note

import (
    "github.com/cloudwego/hertz/pkg/route"
    v1 "github.com/gwen0x4c3/easy_note/cmd/api/handler/note_handler/v1"
    "github.com/gwen0x4c3/easy_note/cmd/api/middleware"
)

func InitNoteV1(g *route.RouterGroup) {
    group := g.Group("/note")
    group.Use(middleware.AuthMiddleware.MiddlewareFunc())
    group.POST("/query", v1.QueryNote)
}

// InitNoteV2 to be used in the future
func InitNoteV2() {

}
