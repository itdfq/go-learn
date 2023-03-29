package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//定义连接信息
const MQURL = "amqp://godchin:godchin@121.36.77.21:5672/"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换器名称
	Exchange string
	//bind key
	Key string
	//连接信息
	Mqurl string
}

//创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

//断开channel和 connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

//订阅模式创建RabbitMQ实例
func NewRabbitMQPubSub(exchangeMame string) *RabbitMQ {
	//创建RabbitMQ实例
	rabbitmq := NewRabbitMQ("", exchangeMame, "")
	var err error
	//获取连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to Open a channel")
	return rabbitmq
}

//订阅模式生产
func (r *RabbitMQ) PublishPub(message string) {
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(r.Exchange, "fanout", true, false, false, false, nil)
	r.failOnErr(err, "Failed to declare an exchange")

	//发送消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		false, //如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false, //如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	r.failOnErr(err, "failed send message ")

}

//订阅模式消费代码
func (r *RabbitMQ) RecieveSub() {
	//创建交换机
	err := r.channel.ExchangeDeclare(r.Exchange, "fanout", true, false, false, false, nil)
	r.failOnErr(err, "consume failed to declare an exchange")
	//试探性创建队列
	q, err := r.channel.QueueDeclare(
		"",    //随机生产队列名称
		false, //是否持久化
		false, //是否自动删除
		true,  //排他性
		false, //是否阻塞处理
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")

	//绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		"", //在pub/sub模式下，key为空
		r.Exchange,
		false,
		nil,
	)
	r.failOnErr(err, "queue bind to exchange fail")

	//消费消息
	messages, err := r.channel.Consume(q.Name, "", true, false, false, false, nil)
	forever := make(chan bool)
	go func() {
		for d := range messages {
			log.Fatalf("Received a message：%s", d.Body)
		}
	}()
	fmt.Println("退出请按 CTRL+C \n")
	<-forever
}

//简单模式声明rabbitMQ
func NewRabbitMqSimple(queueName string) *RabbitMQ {
	//创建rabbitMq 实例
	rabbitmq := NewRabbitMQ(queueName, "", "")
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}

//直接模式生产
func (r *RabbitMQ) PublishSimple(message string) {
	//申请队列，如果不存在会自动创建，存在则跳过创建
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false, false, false, nil)
	r.failOnErr(err, "申请队列失败")
	r.channel.Publish(r.Exchange, r.QueueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})

}

//简单模式消费
func (r *RabbitMQ) ConsumeSimple() {
	//申请队列
	q, err := r.channel.QueueDeclare(r.QueueName, false, false, false, false, nil)
	r.failOnErr(err, "消费申请队列失败")
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true, //是否自动应答
		false,
		false, //设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, //是否阻塞
		nil)
	forever := make(chan bool)
	//启动协程处理消息
	go func() {
		for d := range msgs {
			log.Printf("simple received a message :%s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
