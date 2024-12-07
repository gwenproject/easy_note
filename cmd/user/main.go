package main

import (
	"log"

	_ "github.com/gwen0x4c3/easy_note/cmd/user/dal"
	"github.com/gwen0x4c3/easy_note/kitex_gen/kuser/userservice"
)

func main() {
	svr := userservice.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
