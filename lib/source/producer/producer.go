/*
 * Copyright 2019 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package producer

import (
	"errors"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/source/util"
	"github.com/segmentio/kafka-go"
	"io/ioutil"
	"log"
	"os"
)

type Producer struct {
	config          config.Config
	devices         *kafka.Writer
	devicetypes     *kafka.Writer
	concepts        *kafka.Writer
	deviceclass     *kafka.Writer
	aspect          *kafka.Writer
	characteristics *kafka.Writer
}

func New(conf config.Config) (*Producer, error) {
	broker, err := util.GetBroker(conf.ZookeeperUrl)
	if err != nil {
		return nil, err
	}
	if len(broker) == 0 {
		return nil, errors.New("missing kafka broker")
	}
	devices, err := getProducer(broker, conf.DeviceTopic, conf.LogLevel == "DEBUG")
	if err != nil {
		return nil, err
	}

	devicetypes, err := getProducer(broker, conf.DeviceTypeTopic, conf.LogLevel == "DEBUG")
	if err != nil {
		return nil, err
	}

	concepts, err := getProducer(broker, conf.ConceptTopic, conf.LogLevel == "DEBUG")
	if err != nil {
		return nil, err
	}

	characteristics, err := getProducer(broker, conf.CharacteristicTopic, conf.LogLevel == "DEBUG")
	if err != nil {
		return nil, err
	}

	deviceclass, err := getProducer(broker, conf.DeviceClassTopic, conf.LogLevel == "DEBUG")
	if err != nil {
		return nil, err
	}

	aspect, err := getProducer(broker, conf.AspectTopic, conf.LogLevel == "DEBUG")
	if err != nil {
		return nil, err
	}
	return &Producer{config: conf, devices: devices, devicetypes: devicetypes, concepts: concepts, characteristics: characteristics, deviceclass: deviceclass, aspect: aspect}, nil
}

func getProducer(broker []string, topic string, debug bool) (writer *kafka.Writer, err error) {
	var logger *log.Logger
	if debug {
		logger = log.New(os.Stdout, "[KAFKA-PRODUCER] ", 0)
	} else {
		logger = log.New(ioutil.Discard, "", 0)
	}
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:     broker,
		Topic:       topic,
		MaxAttempts: 10,
		Logger:      logger,
	})
	return writer, err
}
