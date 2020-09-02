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
	"github.com/SENERGY-Platform/semantic-repository/lib/source/consumer/listener"
	"log"
)

type Producer struct {
	config   config.Config
	listener map[string]listener.Listener
}

func New(conf config.Config) *Producer {
	return &Producer{config: conf, listener: map[string]listener.Listener{}}
}

func (this *Producer) Handle(topic string, handler listener.Listener) {
	this.listener[topic] = handler
}

func (this *Producer) callListener(topic string, message []byte) error {
	handler, ok := this.listener[topic]
	if !ok {
		return errors.New("unknown topic:" + topic)
	}
	err := handler(message)
	if err != nil {
		log.Println("TEST-WARNING:", err)
	}
	return err
}
