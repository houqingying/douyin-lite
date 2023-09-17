package rabbitMQ

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const MQURL = "amqp://user@39.105.199.147:5672/"

type RabbitMQ struct {
	conn  *amqp.Connection
	mqurl string
}

var Rmq *RabbitMQ

// InitRabbitMQ 初始化
func InitRabbitMQ() {
	Rmq = &RabbitMQ{
		mqurl: MQURL,
	}
	dial, err := amqp.Dial(Rmq.mqurl)
	Rmq.failOnErr(err, "创建连接失败")
	Rmq.conn = dial
}

// 连接出错调用
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s\n", err, message)
		panic(fmt.Sprintf("%s:%s\n", err, message))
	}
}

// 关闭连接
func (r *RabbitMQ) destroy() {
	r.conn.Close()
}
