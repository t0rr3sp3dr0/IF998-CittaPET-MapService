package main

import (
	"fmt"
	"log"

	citta "cin.ufpe.br/~if998/cittapet/mapservice/proto"

	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

type consumer struct {
	server   *mapServer
	messages <-chan amqp.Delivery
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func NewConsumer(server *mapServer) (*consumer, error) {
	c := consumer{}

	c.server = server

	queueConn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", queueUser, queuePass, queueAddr, queuePort))
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to RabbitMQ: %v", err)
	}
	defer queueConn.Close()

	ch, err := queueConn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"stream", // name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to declare a queue: %v", err)
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to set QoS: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,   // queue
		hostname, // consumer
		false,    // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to register a consumer: %v", err)
	}

	c.messages = msgs

	return &c, nil
}

func (c *consumer) Consume() error {
	for d := range c.messages {
		event := &citta.BusEvent{}
		if err := proto.Unmarshal(d.Body, event); err != nil {
			d.Ack(false)
			return fmt.Errorf("Failed to parse bus-event: %v", err)
		}
		log.Printf(" [x] Consuming bus-event %v", event)

		coordinate, err := UTM2LatLng(event.Coordinate)
		if err != nil {
			return fmt.Errorf("Failed to convert UTM to LatLng: %v", err)
		}

		c.server.Buses.Store(event.Unit, &citta.GetBusesLocationResponse_BusLocation{
			Unit:       event.Unit,
			Timestamp:  event.Timestamp,
			Coordinate: coordinate,
		})

		d.Ack(false)
		log.Printf(" [x] Consumed bus-event %v", event)
	}

	return nil
}
