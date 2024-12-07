package main

import (
	"log"

	"github.com/gwen0x4c3/easy_note/kitex_gen/kuser/userservice"
)

func main() {
	svr := userservice.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
