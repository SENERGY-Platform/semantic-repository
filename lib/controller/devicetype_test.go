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

func TestValidDeviceType(t *testing.T) {

	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:1111"
	devicetype.Name = "Philips Hue Color"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:2222",
		Name: "Lamp",
		RdfType: model.SES_ONTOLOGY_DEVICE_CLASS,
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.RdfType = model.SES_ONTOLOGY_DEVICE_TYPE
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3333",
		"localId",
		"setBrigthness",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4444", Name: "Lighting", RdfType: model.SES_ONTOLOGY_ASPECT}},
		"protocolId",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5555", Name: "brightnessAdjustment", ConceptId: "urn:infai:ses:concept:6666", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		model.SES_ONTOLOGY_SERVICE,
	})

	controller := Controller{}
	err, code := controller.ValidateDeviceType(devicetype)
	if err == nil && code == http.StatusOK {
		t.Log(devicetype)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationDeviceTypeMissingId(t *testing.T) {
	devicetype := model.DeviceType{}
	devicetype.Id = ""

	controller := Controller{}
	err, code := controller.ValidateDeviceType(devicetype)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationDeviceTypeMissingName(t *testing.T) {
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:5555"
	devicetype.Name = ""

	controller := Controller{}
	err, code := controller.ValidateDeviceType(devicetype)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationDeviceTypeWrongType(t *testing.T) {
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:5555"
	devicetype.Name = "philips hue color"
	devicetype.RdfType = "type"

	controller := Controller{}
	err, code := controller.ValidateDeviceType(devicetype)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationDeviceTypeNoServiceData(t *testing.T) {
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:5555"
	devicetype.Name = "philips hue color"
	devicetype.RdfType = model.SES_ONTOLOGY_DEVICE_TYPE

	controller := Controller{}
	err, code := controller.ValidateDeviceType(devicetype)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationDeviceTypeNoDeviceClass(t *testing.T) {
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:5555"
	devicetype.Name = "philips hue color"
	devicetype.RdfType = model.SES_ONTOLOGY_DEVICE_TYPE
	devicetype.Services = []model.Service{{Id: "urn:infai:ses:service:1",
		RdfType: model.SES_ONTOLOGY_SERVICE,
		Name: "test", LocalId: "2",
		ProtocolId: "3",
		Aspects:    []model.Aspect{{Id: "urn:infai:ses:aspect:1", Name: "aspect", RdfType: model.SES_ONTOLOGY_ASPECT}},
		Functions:  []model.Function{{Id: "urn:infai:ses:function:1", Name: "function", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION, ConceptId: "urn:infai:ses:concept:1"}}}}
	devicetype.DeviceClass = model.DeviceClass{}

	controller := Controller{}
	err, code := controller.ValidateDeviceType(devicetype)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}
