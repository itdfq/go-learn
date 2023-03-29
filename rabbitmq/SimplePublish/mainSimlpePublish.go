package main

import (
	"fmt"
	"go-learn/rabbitmq/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMqSimple("dfq")
	for i := 0; i < 5; i++ {
		rabbitmq.PublishSimple("测试第二次发送代码")
	}
	fmt.Println("发送成功!!!")

}
