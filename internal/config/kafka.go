package config

import (
	"strings"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewKafkaProducer(config *viper.Viper, log *logrus.Logger) sarama.SyncProducer {
	if !config.GetBool("KAFKA_PRODUCER_ENABLED") {
		log.Info("Kafka producer is disabled")
		return nil
	}

	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Retry.Max = 3

	brokers := strings.Split(config.GetString("KAFKA_BOOTSTRAP_SERVERS"), ",")

	producer, err := sarama.NewSyncProducer(brokers, saramaConfig)
	if err != nil {
		log.Fatalf("Failed to to create kafka producer: %+v", err)
	}

	return producer
}