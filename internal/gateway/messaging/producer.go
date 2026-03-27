package messaging

import (
	"encoding/json"
	"first-go-project/internal/model"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type Producer[T model.Event] struct {
	Producer sarama.SyncProducer
	Topic string
	Log *logrus.Logger
}

func (p *Producer[T]) GetTopic() *string {
	return &p.Topic
}

func (p *Producer[T]) Send(event T) error {
	value, err := json.Marshal(event)
	if err != nil {
		p.Log.Warnf("Failed to marshal event %s to bytes", *p.GetTopic())
		return err
	}

	message := &sarama.ProducerMessage{
		Topic: p.Topic,
		Key: sarama.StringEncoder(event.GetId()),
		Value: sarama.ByteEncoder(value),
	}

	partition, offset, err := p.Producer.SendMessage(message)
	if err != nil {
		p.Log.Warnf("Failed to publish message for event %s", *p.GetTopic())
		return err
	}

	p.Log.Debugf("Published message for event %s, partition %d, offset %d", *p.GetTopic(), partition, offset)
	return nil
}
