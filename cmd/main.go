package main

import (
	"MyToDoMadeOnGo"
	"log"
)

func main() {
	srv := new(MyToDoMadeOnGo.MyServer)
	if err := srv.Run("8080"); err != nil {
		log.Fatal("err %s", err.Error())
	}
}
