package main

import services "github.com/usernamenenad/ftn-acr/services/rabbitmq"

func main() {
	// s := server.NewServer(":8000")
	// if err := s.Run(); err != nil {
	// 	log.Fatal(err)
	// }

	sub := services.RabbitMqSubscriber{}

	sub.Subscribe("hello")

}
