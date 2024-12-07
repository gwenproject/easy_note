package main

import (
	"log"

	_ "github.com/gwen0x4c3/easy_note/cmd/note/dal"
	"github.com/gwen0x4c3/easy_note/kitex_gen/knote/noteservice"
)

func main() {
	svr := noteservice.NewServer(new(NoteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
