package messaging

import (
	"first-go-project/internal/model"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type UserProducer struct {
	Producer[*model.UserEvent]
}

func NewUserProducer(producer sarama.SyncProducer, log *logrus.Logger) *UserProducer {
	return &UserProducer{
		Producer: Producer[*model.UserEvent]{
			Producer: producer,
			Topic: "first_go.users",
			Log: log,
		},
	}
}

