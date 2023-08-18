package rabbitmq

import "log"

type LikeMQ struct {
	RabbitMQ
	QueueName string
	Exchange  string
	Key       string
}

var LikeAddMQ *LikeMQ
var LikeDelMQ *LikeMQ

func (r *LikeMQ) PublishModel(mes string) error {
	_, err := r.channel.QueueDeclare(r.QueueName, false, false,
		false, false, nil)

	if err != nil {
		log.Println(err)
		return err
	}

}
