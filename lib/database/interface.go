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

package database

import (
	"context"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
)

type Database interface {
	Disconnect()

	GetDevice(ctx context.Context, id string) (device model.Device, exists bool, err error)
	SetDevice(ctx context.Context, device model.Device) error
	RemoveDevice(ctx context.Context, id string) error

	GetDeviceType(ctx context.Context, id string) (deviceType model.DeviceType, exists bool, err error)
	SetDeviceType(ctx context.Context, deviceType model.DeviceType) error
	RemoveDeviceType(ctx context.Context, id string) error
}
