package main

import (
	"context"
	"fmt"

	"github.com/usernamenenad/ftn-acr/models/database"
	"github.com/usernamenenad/ftn-acr/models/repository"
)

func main() {
	database.Open("postgres://postgres:password@postgres:5432/postgres?sslmode=disable")
	defer database.Close()

	ch := make(chan bool)

	go func(ch chan bool) {
		fmt.Println(repository.GetProgramByShortName(context.Background(), "RA"))
		ch <- true
	}(ch)

	<-ch
}
