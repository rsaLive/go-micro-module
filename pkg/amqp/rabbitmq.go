package amqp

import (
	"github.com/google/wire"
	"github.com/streadway/amqp"
)

var RabbitMqConn *amqp.Connection

var Provider = wire.NewSet(NewRabbitMq)

func NewRabbitMq() *amqp.Connection {
	return RabbitMqConn
}
