package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Go-Producer API is alive")
	})

	app.Get("/api/producer/:topic/:data", func(c *fiber.Ctx) error {
		topic := c.Params("topic")
		data := c.Params("data")
		err := producer(topic, data)
		return c.JSON(map[string]interface{}{
			"topic": topic,
			"data":  data,
			"error": err,
		})
	})

	app.Listen(":8090")

}

// producer - produces message for specifeid topic
func producer(topic, data string) error {

	fmt.Println("kafka producer: ", topic, " ", data)
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "broker:29092"})
	if err != nil {
		return err
	}

	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic},
		Value:          []byte(data),
	}, nil)

	p.Flush(15 * 1000)
	return nil
}
