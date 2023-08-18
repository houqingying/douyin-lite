package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

const MQCONURL = "amqp://guest:guest@106.14.252.145:5672/"

var BaseRmq *RabbitMQ

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 连接信息
	MqUrl string
}

// 初始化消息队列
func InitRabbitMQ() {
	BaseRmq = &RabbitMQ{
		MqUrl: MQCONURL,
	}
	//连接
	conn, err := amqp.Dial(BaseRmq.MqUrl)
	BaseRmq.failOnError(err, "Failed to connect to RabbitMQ")
	BaseRmq.conn = conn
	BaseRmq.channel, err = conn.Channel()
	BaseRmq.failOnError(err, "Failed to get channel")
}

// 异常处理
func (r *RabbitMQ) failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// 关闭
func (r *RabbitMQ) close() {
	r.conn.Close()
	r.channel.Close()
}
