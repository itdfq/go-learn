package main

import (
	"fmt"
	"go-learn/rabbitmq/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMqSimple("dfq")
	rabbitmq.PublishSimple("测试第二次发送代码")
	fmt.Println("发送成功!!!")
}
