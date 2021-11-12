package main

import (
	"bytes"
	golang_rss_reader_package "github.com/Welith/golang-rss-reader-package"
	"github.com/streadway/amqp"
	"net/http"
	"os"
)

//ProcessIncomingMessages
func ProcessIncomingMessages() {

	conn, err := amqp.Dial(os.Getenv("AMQP_TRANSPORT"))

	if err != nil {

		LogError(err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {

		LogError(err.Error())
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		os.Getenv("QUEUE_NAME"), // name
		true,                    // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)

	if err != nil {

		LogError(err.Error())
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {

		LogError(err.Error())
	}

	forever := make(chan bool)

	go func() {

		for d := range msgs {

			messageData := Request{}

			if err = Deserialize(d.Body, &messageData.Urls); err != nil {

				LogError(err.Error())
			}

			result := golang_rss_reader_package.Parse(messageData.Urls)
			var response ResponseItem

			response.Items = result
			for _, item := range response.Items {

				byteData, err := Serialize(item)

				if err != nil {

					LogError(err.Error())
				}

				_, err = http.Post(os.Getenv("LARAVEL_APP_URL") + "/api/feeds", "application/json", bytes.NewBuffer(byteData))

				if err != nil {

					LogError(err.Error())
				}
			}

			ch.Ack(d.DeliveryTag, false)
		}
	}()

	<-forever
}

