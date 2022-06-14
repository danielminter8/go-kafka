package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Go-Consumer API is alive")
	})

	app.Get("api/consumer/:topic", func(c *fiber.Ctx) error {
		topic := c.Params("topic")
		go consumer(topic)

		return c.JSON(map[string]interface{}{
			"topic":   topic,
			"data":    consumedData,
			"message": "returns all consumed data for specified topic or you can check logs",
		})
	})

	app.Listen(":8091")

}

var consumedData []string

// consumer - consumes messages for specifeid topic
func consumer(topic string) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "broker:29092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			consumedData = append(consumedData, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

}
