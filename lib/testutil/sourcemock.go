package testutil

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/source/consumer/listener"
	"github.com/SENERGY-Platform/semantic-repository/lib/testutil/producer"
	"log"
)

func StartSourceMock(config config.Config, control listener.Controller) (prod *producer.Producer, err error) {
	prod = producer.New(config)
	for _, factory := range listener.Factories {
		topic, handler, err := factory(config, control)
		if err != nil {
			log.Println("ERROR: listener.factory", topic, err)
			return prod, err
		}
		prod.Handle(topic, handler)
	}
	return prod, err
}
