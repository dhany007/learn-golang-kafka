package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

const (
	topic      = "my-kafka-topic"
	brokerAddr = "localhost:9092"
)

func NewWriter() *kafka.Writer {
	l := log.New(os.Stdout, "kafka writer: ", 0)
	return kafka.NewWriter(
		kafka.WriterConfig{
			Brokers: []string{brokerAddr},
			Topic:   topic,
			Logger:  l,
		},
	)
}

func NewReader() *kafka.Reader {
	l := log.New(os.Stdout, "kafka reader: ", 0)
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddr},
		Topic:   topic,
		GroupID: "my-group",
		Logger:  l,
	})
}

func produce(ctx context.Context) {
	w := NewWriter()

	for i := 1; i < 10; i++ {
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("this is message " + strconv.Itoa(i)),
		})

		if err != nil {
			log.Fatalf("could not write message (err: %v)", err)
		}

		fmt.Println("writes: ", i)

		time.Sleep(time.Second)
	}
}

func consume(ctx context.Context) {
	r := NewReader()
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Fatalf("failed when read message (err: %v)", err)
		}

		fmt.Println("reveived= ", string(msg.Value))
	}
}

func main() {
	fmt.Println("learn kafka")
	ctx := context.Background()
	go produce(ctx)
	consume(ctx)
}
