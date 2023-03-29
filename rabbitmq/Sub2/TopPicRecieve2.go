package main

import "go-learn/rabbitmq/RabbitMQ"

func main() {
	kutengOne := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "kuteng.*.two")
	kutengOne.RecieveTopic()
}
