package main

import (
	"log"

	"github.com/usernamenenad/ftn-acr/src/api/server"
)

func main() {
	s := server.NewServer(":8000")
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
