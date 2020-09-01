/*
 *
 * Copyright 2019 InfAI (CC SES)
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

package controller

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"net/http"
	"testing"
)

func TestValidDeviceClass(t *testing.T) {
	deviceclass := model.DeviceClass{}
	deviceclass.Id = "urn:infai:ses:deviceclass:1234"
	deviceclass.Name = "Lamp"
	deviceclass.RdfType = model.SES_ONTOLOGY_DEVICE_CLASS
	deviceclass.Image = "image"

	controller := Controller{}
	err, code := controller.ValidateDeviceClass(deviceclass)
	if err == nil && code == http.StatusOK {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestDeviceClassMissingId(t *testing.T) {
	deviceclass := model.DeviceClass{}
	deviceclass.Id = ""

	controller := Controller{}
	err, code := controller.ValidateDeviceClass(deviceclass)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestDeviceClassMissingName(t *testing.T) {
	deviceclass := model.DeviceClass{}
	deviceclass.Id = "urn:infai:ses:deviceclass:1234"
	deviceclass.Name = ""

	controller := Controller{}
	err, code := controller.ValidateDeviceClass(deviceclass)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestDeviceClassWrongType(t *testing.T) {
	deviceclass := model.DeviceClass{}
	deviceclass.Id = "urn:infai:ses:deviceclass:1234"
	deviceclass.Name = "Lamp"
	deviceclass.RdfType = "wrongType"

	controller := Controller{}
	err, code := controller.ValidateDeviceClass(deviceclass)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}
