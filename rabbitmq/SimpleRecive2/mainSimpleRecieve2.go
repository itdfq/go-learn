package main

import "go-learn/rabbitmq/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMqSimple("dfq")
	rabbitmq.ConsumeSimple("消费者二")
}
