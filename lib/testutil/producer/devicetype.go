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
	"encoding/json"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/SENERGY-Platform/semantic-repository/lib/source/consumer/listener"
	"log"
	"runtime/debug"
)

func (this *Producer) PublishDeviceType(deviceType model.DeviceType, userId string) (err error) {
	cmd := listener.DeviceTypeCommand{Command: "PUT", Id: deviceType.Id, DeviceType: deviceType, Owner: userId}
	return this.PublishDeviceTypeCommand(cmd)
}

func (this *Producer) PublishDeviceTypeDelete(id string, userId string) error {
	cmd := listener.DeviceTypeCommand{Command: "DELETE", Id: id, Owner: userId}
	return this.PublishDeviceTypeCommand(cmd)
}

func (this *Producer) PublishDeviceTypeCommand(cmd listener.DeviceTypeCommand) error {
	if this.config.Debug {
		log.Println("DEBUG: produce devicetype", cmd)
	}
	message, err := json.Marshal(cmd)
	if err != nil {
		debug.PrintStack()
		return err
	}
	return this.callListener(this.config.DeviceTypeTopic, message)
}
