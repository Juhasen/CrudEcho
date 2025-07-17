package kafka

import (
	"RestCrud/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"
	"time"
)

type EventType string

const (
	CREATE EventType = "CREATE"
	EDIT   EventType = "EDIT"
)

var kafkaClient *kgo.Client

func InitClient() error {
	var err error
	kafkaClient, err = kgo.NewClient(
		kgo.SeedBrokers("localhost:29092"),
	)
	return err
}

func ProduceTodoEvent(data interface{}, eventType EventType, id string) error {
	if kafkaClient == nil {
		return fmt.Errorf("kafka client is not initialized")
	}

	var topic string
	var value []byte
	var err error

	switch data.(type) {
	case *model.Task:
		value, err = json.Marshal(map[string]string{
			"type":      string(eventType),
			"taskId":    id,
			"timestamp": time.Now().Format(time.RFC3339),
		})
		topic = "todo-task"
	case *model.User:
		value, err = json.Marshal(map[string]string{
			"type":      string(eventType),
			"userId":    id,
			"timestamp": time.Now().Format(time.RFC3339),
		})
		topic = "todo-user"
	default:
		return fmt.Errorf("unsupported data type")
	}

	if err != nil {
		return err
	}

	record := &kgo.Record{
		Topic: topic,
		Value: value,
	}

	ctx := context.Background()
	return kafkaClient.ProduceSync(ctx, record).FirstErr()
}
