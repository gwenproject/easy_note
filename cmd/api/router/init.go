package router

import (
	"log"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/gwen0x4c3/easy_note/cmd/api/router/note"
	"github.com/gwen0x4c3/easy_note/cmd/api/router/user"
)

func InitRouter(hz *server.Hertz) {
	log.Println("Start to initialize hertz routes")

	v1 := hz.Group("/v1")
	// v2 := hz.Group("/v2")

	// user
	user.InitUserV1(v1)

	// note
	note.InitNoteV1(v1)
}
