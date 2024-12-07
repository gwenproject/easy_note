package main

import (
	"log"

	"github.com/gwen0x4c3/easy_note/kitex_gen/knote/noteservice"
)

func main() {
	svr := noteservice.NewServer(new(NoteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
