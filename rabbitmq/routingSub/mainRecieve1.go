package main

import "go-learn/rabbitmq/RabbitMQ"

func main() {
	kutengtwo := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_two")
	kutengtwo.RecieveRouting()
}
