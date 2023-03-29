package main

import "go-learn/rabbitmq/RabbitMQ"

func main() {
	kutengone := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_one")
	kutengone.RecieveRouting()
}
