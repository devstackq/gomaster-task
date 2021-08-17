package main

import (
	"log"

	"github.com/devstackq/server"
)

//run server
//config db
//endpoint: handler
//uuid- randomizer
//get/put - by id, post - create, validation date
//crud - query db

func main() {

	s := server.NewServer()
	err := s.Run()
	if err != nil {
		log.Println(err)
		return
	}

}
