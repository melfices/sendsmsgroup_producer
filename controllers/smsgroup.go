package controllers

import (
	"encoding/json"
	"net/http"
	"sendsmsgroup-producer/configs"
	"sendsmsgroup-producer/models"
	"sendsmsgroup-producer/responses"

	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
)

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, responses.SMSGroupResponseTxt{Status: http.StatusOK, Message: "success", Data: "OK"})
}
func AddSMSQueue(c echo.Context) error {
	amqpServerURL := configs.EnvAMQPURL()

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	// Let's start by opening a channel to our RabbitMQ
	// instance over the connection we have already
	// established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// With the instance and declare Queues that we can
	// publish and subscribe to.
	_, err = channelRabbitMQ.QueueDeclare(
		configs.EnvQueueName(), // queue name
		true,                   // durable
		false,                  // auto delete
		false,                  // exclusive
		false,                  // no wait
		nil,                    // arguments
	)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic("err")
	}
	smsgroup_param := new(models.SMSGroupParam)
	if err := c.Bind(smsgroup_param); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	data, _ := json.Marshal(smsgroup_param)
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        data,
	}
	// Attempt to publish a message to the queue.
	if err := channelRabbitMQ.Publish(
		"",                     // exchange
		configs.EnvQueueName(), // queue name
		false,                  // mandatory
		false,                  // immediate
		message,                // message to publish
	); err != nil {
		return c.JSON(http.StatusOK, responses.SMSGroupResponseTxt{Status: http.StatusOK, Message: "failed", Data: "error publish to queue"})
	}
	return c.JSON(http.StatusOK, responses.SMSGroupResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": smsgroup_param}})
}
