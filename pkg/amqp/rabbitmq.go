package amqp

import (
	"github.com/google/wire"
	"github.com/streadway/amqp"
)

var RabbitMqConn *amqp.Connection

var Provider = wire.NewSet(NewRabbitMq)

func NewRabbitMq() *amqp.Connection {
	//var err error
	// 创建连接
	/*RabbitMqConn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/%s", appConfig.Data.RabbitMq.User, appConfig.Data.RabbitMq.Password, appConfig.Data.RabbitMq.Host, appConfig.Data.RabbitMq.Port, appConfig.Data.RabbitMq.Vhost))
	if err != nil {
		log.Fatalf("failed to connect RabbitMQ: %v", err)
	}*/
	return RabbitMqConn
}
