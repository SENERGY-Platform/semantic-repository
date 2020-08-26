/*
 *
 * Copyright 2020 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *
 */

package producer

import (
	"context"
	"encoding/json"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/SENERGY-Platform/semantic-repository/lib/source/consumer/listener"
	"github.com/segmentio/kafka-go"
	"log"
	"runtime/debug"
	"time"
)

func (this *Producer) PublishDeviceClass(deviceclass model.DeviceClass, userId string) (err error) {
	cmd := listener.DeviceClassCommand{Command: "PUT", DeviceClass: deviceclass, Owner: userId}
	return this.PublishDeviceClassCommand(cmd)
}

func (this *Producer) PublishDeviceClassDelete(id string, userId string) error {
	cmd := listener.DeviceClassCommand{Command: "DELETE", Id: id, Owner: userId}
	return this.PublishDeviceClassCommand(cmd)
}

func (this *Producer) PublishDeviceClassCommand(cmd listener.DeviceClassCommand) error {
	if this.config.LogLevel == "DEBUG" {
		log.Println("DEBUG: produce deviceclass", cmd)
	}
	message, err := json.Marshal(cmd)
	if err != nil {
		debug.PrintStack()
		return err
	}
	err = this.deviceclass.WriteMessages(
		context.Background(),
		kafka.Message{
			Key:   []byte(cmd.Id),
			Value: message,
			Time:  time.Now(),
		},
	)
	if err != nil {
		debug.PrintStack()
	}
	return err
}
