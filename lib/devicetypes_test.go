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

package lib

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/controller"
	"github.com/SENERGY-Platform/semantic-repository/lib/database"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/SENERGY-Platform/semantic-repository/lib/source/producer"
	"testing"
)

func TestProduceValidDeviceType(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:1e1e"
	devicetype.Name = "Philips Hue Color"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:2e2e",
		Name: "Lamp",
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3e3e",
		"localId",
		"setBrightness2",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4e4e", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5e5e", Name: "brightnessAdjustment", ConceptId: "urn:ses:infai:concept:1a1a1a", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3f3f",
		"localId",
		"setBrightness1",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4e4e", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5e5e", Name: "brightnessAdjustment", ConceptId: "urn:ses:infai:concept:1a1a1a", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")

	devicetype1 := model.DeviceType{}
	devicetype1.Id = "urn:infai:ses:devicetype:1e1e_2"
	devicetype1.Name = "Lifx"
	devicetype1.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:2e2e",
		Name: "Lamp",
	}
	devicetype1.Description = "description"
	devicetype1.Image = "image"
	devicetype1.Services = []model.Service{}
	devicetype1.Services = append(devicetype1.Services, model.Service{
		"urn:infai:ses:service:3e3e_2",
		"localId",
		"setBrightness2",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4e4e", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5e5e", Name: "brightnessAdjustment", ConceptId: "urn:ses:infai:concept:1a1a1a", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	devicetype1.Services = append(devicetype1.Services, model.Service{
		"urn:infai:ses:service:3f3f_2",
		"localId",
		"setBrightness1",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4e4e", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5e5e", Name: "brightnessAdjustment", ConceptId: "urn:ses:infai:concept:1a1a1a", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype1, "sdfdsfsf")
}

func TestReadDeviceTypesWithDeviceClassIdAndFunctionId(t *testing.T) {
	err, con, _ := StartUpScript(t)
	deviceType, err, code := con.GetDeviceTypesFiltered("urn:infai:ses:deviceclass:2e2e", "urn:infai:ses:function:5e5e", "")

	if deviceType[0].Id != "urn:infai:ses:devicetype:1e1e" {
		t.Fatal("error id")
	}

	if deviceType[0].RdfType != model.SES_ONTOLOGY_DEVICE_TYPE {
		t.Fatal("error model")
	}

	if deviceType[0].Name != "Philips Hue Color" {
		t.Fatal("error name")
	}

	if deviceType[0].Description != "" {
		t.Fatal("error description")
	}

	if deviceType[0].Image != "" {
		t.Fatal("error image")
	}
	// DeviceClass
	if deviceType[0].DeviceClass.Id != "urn:infai:ses:deviceclass:2e2e" {
		t.Fatal("error deviceclass id")
	}
	if deviceType[0].DeviceClass.Name != "Lamp" {
		t.Fatal("error deviceclass name")
	}
	if deviceType[0].DeviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("error deviceclass rdf type")
	}
	// Service
	if deviceType[0].Services[0].Id != "urn:infai:ses:service:3f3f" {
		t.Fatal("error service -> 0 -> id")
	}
	if deviceType[0].Services[0].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType[0].Services[0].Name != "setBrightness1" {
		t.Log(deviceType[0].Services[0].Name)
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType[0].Services[0].Description != "" {
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType[0].Services[0].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType[0].Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType[0].Services[0].Aspects[0].Name != "Lighting" {
		t.Log(deviceType[0].Services[0].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType[0].Services[0].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}
	if deviceType[0].Services[0].Functions[0].Id != "urn:infai:ses:function:5e5e" {
		t.Fatal("error function -> 0/0 -> id")
	}
	if deviceType[0].Services[0].Functions[0].Name != "brightnessAdjustment" {
		t.Fatal("error function -> 0/0 -> Name")
	}
	if deviceType[0].Services[0].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 0/0 -> RdfType")
	}
	if deviceType[0].Services[0].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 0/0/0 -> ConceptIds")
	}
	/// service 2
	if deviceType[0].Services[1].Id != "urn:infai:ses:service:3e3e" {
		t.Fatal("error service -> 1 -> id")
	}
	if deviceType[0].Services[1].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 1 -> RdfType")
	}
	if deviceType[0].Services[1].Name != "setBrightness2" {
		t.Log(deviceType[0].Services[1].Name)
		t.Fatal("error service -> 1 -> name")
	}
	if deviceType[0].Services[1].Description != "" {
		t.Fatal("error service -> 1 -> description")
	}
	if deviceType[0].Services[1].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 1 -> LocalId")
	}
	if deviceType[0].Services[1].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 1/0 -> id")
	}
	if deviceType[0].Services[1].Aspects[0].Name != "Lighting" {
		t.Log(deviceType[0].Services[1].Aspects[0].Name)
		t.Fatal("error aspect -> 1/0 -> Name")
	}
	if deviceType[0].Services[1].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 1/0 -> RdfType")
	}
	if deviceType[0].Services[1].Functions[0].Id != "urn:infai:ses:function:5e5e" {
		t.Fatal("error function -> 1/0 -> id")
	}
	if deviceType[0].Services[1].Functions[0].Name != "brightnessAdjustment" {
		t.Fatal("error function -> 1/0 -> Name")
	}
	if deviceType[0].Services[1].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 1/0 -> RdfType")
	}
	if deviceType[0].Services[1].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 1/0/0 -> ConceptIds")
	}

	if deviceType[1].Id != "urn:infai:ses:devicetype:1e1e_2" {
		t.Fatal("error id")
	}

	if deviceType[1].RdfType != model.SES_ONTOLOGY_DEVICE_TYPE {
		t.Fatal("error model")
	}

	if deviceType[1].Name != "Lifx" {
		t.Fatal("error name")
	}

	if deviceType[1].Description != "" {
		t.Fatal("error description")
	}

	if deviceType[1].Image != "" {
		t.Fatal("error image")
	}
	// DeviceClass
	if deviceType[1].DeviceClass.Id != "urn:infai:ses:deviceclass:2e2e" {
		t.Fatal("error deviceclass id")
	}
	if deviceType[1].DeviceClass.Name != "Lamp" {
		t.Fatal("error deviceclass name")
	}
	if deviceType[1].DeviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("error deviceclass rdf type")
	}
	// Service
	if deviceType[1].Services[0].Id != "urn:infai:ses:service:3f3f_2" {
		t.Fatal("error service -> 0 -> id")
	}
	if deviceType[1].Services[0].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType[1].Services[0].Name != "setBrightness1" {
		t.Log(deviceType[0].Services[0].Name)
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType[1].Services[0].Description != "" {
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType[1].Services[0].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType[1].Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType[1].Services[0].Aspects[0].Name != "Lighting" {
		t.Log(deviceType[0].Services[0].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType[1].Services[0].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}
	if deviceType[1].Services[0].Functions[0].Id != "urn:infai:ses:function:5e5e" {
		t.Fatal("error function -> 0/0 -> id")
	}
	if deviceType[1].Services[0].Functions[0].Name != "brightnessAdjustment" {
		t.Fatal("error function -> 0/0 -> Name")
	}
	if deviceType[1].Services[0].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 0/0 -> RdfType")
	}
	if deviceType[1].Services[0].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 0/0/0 -> ConceptIds")
	}
	/// service 2
	if deviceType[1].Services[1].Id != "urn:infai:ses:service:3e3e_2" {
		t.Fatal("error service -> 1 -> id", deviceType[1].Services[1].Id)
	}
	if deviceType[1].Services[1].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 1 -> RdfType")
	}
	if deviceType[1].Services[1].Name != "setBrightness2" {
		t.Log(deviceType[0].Services[1].Name)
		t.Fatal("error service -> 1 -> name")
	}
	if deviceType[1].Services[1].Description != "" {
		t.Fatal("error service -> 1 -> description")
	}
	if deviceType[1].Services[1].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 1 -> LocalId")
	}
	if deviceType[1].Services[1].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 1/0 -> id")
	}
	if deviceType[1].Services[1].Aspects[0].Name != "Lighting" {
		t.Log(deviceType[0].Services[1].Aspects[0].Name)
		t.Fatal("error aspect -> 1/0 -> Name")
	}
	if deviceType[1].Services[1].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 1/0 -> RdfType")
	}
	if deviceType[1].Services[1].Functions[0].Id != "urn:infai:ses:function:5e5e" {
		t.Fatal("error function -> 1/0 -> id")
	}
	if deviceType[1].Services[1].Functions[0].Name != "brightnessAdjustment" {
		t.Fatal("error function -> 1/0 -> Name")
	}
	if deviceType[1].Services[1].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 1/0 -> RdfType")
	}
	if deviceType[1].Services[1].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 1/0/0 -> ConceptIds")
	}

	if err != nil {
		t.Fatal(deviceType, err, code)
	} else {
		t.Log(deviceType)
	}
}

func TestReadDeviceTypeWithId1(t *testing.T) {
	err, con, _ := StartUpScript(t)
	deviceType, err, code := con.GetDeviceType("urn:infai:ses:devicetype:1e1e")

	if deviceType.Id != "urn:infai:ses:devicetype:1e1e" {
		t.Fatal("error id")
	}

	if deviceType.RdfType != model.SES_ONTOLOGY_DEVICE_TYPE {
		t.Fatal("error model")
	}

	if deviceType.Name != "Philips Hue Color" {
		t.Fatal("error name")
	}

	if deviceType.Description != "" {
		t.Fatal("error description")
	}

	if deviceType.Image != "" {
		t.Fatal("error image")
	}
	// DeviceClass
	if deviceType.DeviceClass.Id != "urn:infai:ses:deviceclass:2e2e" {
		t.Fatal("error deviceclass id")
	}
	if deviceType.DeviceClass.Name != "Lamp" {
		t.Fatal("error deviceclass name")
	}
	if deviceType.DeviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("error deviceclass rdf type")
	}
	// Service
	if deviceType.Services[0].Id != "urn:infai:ses:service:3f3f" {
		t.Fatal("error service -> 0 -> id")
	}
	if deviceType.Services[0].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType.Services[0].Name != "setBrightness1" {
		t.Log(deviceType.Services[0].Name)
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType.Services[0].Description != "" {
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType.Services[0].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType.Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType.Services[0].Aspects[0].Name != "Lighting" {
		t.Log(deviceType.Services[0].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType.Services[0].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}
	if deviceType.Services[0].Functions[0].Id != "urn:infai:ses:function:5e5e" {
		t.Fatal("error function -> 0/0 -> id")
	}
	if deviceType.Services[0].Functions[0].Name != "brightnessAdjustment" {
		t.Fatal("error function -> 0/0 -> Name")
	}
	if deviceType.Services[0].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 0/0 -> RdfType")
	}
	if deviceType.Services[0].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 0/0/0 -> ConceptIds")
	}
	/// service 2
	if deviceType.Services[1].Id != "urn:infai:ses:service:3e3e" {
		t.Fatal("error service -> 1 -> id")
	}
	if deviceType.Services[1].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 1 -> RdfType")
	}
	if deviceType.Services[1].Name != "setBrightness2" {
		t.Log(deviceType.Services[1].Name)
		t.Fatal("error service -> 1 -> name")
	}
	if deviceType.Services[1].Description != "" {
		t.Fatal("error service -> 1 -> description")
	}
	if deviceType.Services[1].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 1 -> LocalId")
	}
	if deviceType.Services[1].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 1/0 -> id")
	}
	if deviceType.Services[1].Aspects[0].Name != "Lighting" {
		t.Log(deviceType.Services[1].Aspects[0].Name)
		t.Fatal("error aspect -> 1/0 -> Name")
	}
	if deviceType.Services[1].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 1/0 -> RdfType")
	}
	if deviceType.Services[1].Functions[0].Id != "urn:infai:ses:function:5e5e" {
		t.Fatal("error function -> 1/0 -> id")
	}
	if deviceType.Services[1].Functions[0].Name != "brightnessAdjustment" {
		t.Fatal("error function -> 1/0 -> Name")
	}
	if deviceType.Services[1].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 1/0 -> RdfType")
	}
	if deviceType.Services[1].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 1/0/0 -> ConceptIds")
	}

	if err != nil {
		t.Fatal(deviceType, err, code)
	} else {
		t.Log(deviceType)
	}
}

func TestReadDeviceTypeWithId2(t *testing.T) {
	err, con, _ := StartUpScript(t)
	deviceType, err, code := con.GetDeviceType("urn:infai:ses:devicetype:1e1e_2")

	if deviceType.Id != "urn:infai:ses:devicetype:1e1e_2" {
		t.Fatal("error id")
	}

	if deviceType.RdfType != model.SES_ONTOLOGY_DEVICE_TYPE {
		t.Fatal("error model")
	}

	if deviceType.Name != "Lifx" {
		t.Fatal("error name")
	}

	if deviceType.Description != "" {
		t.Fatal("error description")
	}

	if deviceType.Image != "" {
		t.Fatal("error image")
	}
	// DeviceClass
	if deviceType.DeviceClass.Id != "urn:infai:ses:deviceclass:2e2e" {
		t.Fatal("error deviceclass id")
	}
	if deviceType.DeviceClass.Name != "Lamp" {
		t.Fatal("error deviceclass name")
	}
	if deviceType.DeviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("error deviceclass rdf type")
	}
	// Service
	if deviceType.Services[0].Id != "urn:infai:ses:service:3f3f_2" {
		t.Fatal("error service -> 0 -> id")
	}
	if deviceType.Services[0].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType.Services[0].Name != "setBrightness1" {
		t.Log(deviceType.Services[0].Name)
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType.Services[0].Description != "" {
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType.Services[0].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType.Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType.Services[0].Aspects[0].Name != "Lighting" {
		t.Log(deviceType.Services[0].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType.Services[0].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}
	if deviceType.Services[0].Functions[0].Id != "urn:infai:ses:function:5e5e" {
		t.Fatal("error function -> 0/0 -> id")
	}
	if deviceType.Services[0].Functions[0].Name != "brightnessAdjustment" {
		t.Fatal("error function -> 0/0 -> Name")
	}
	if deviceType.Services[0].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 0/0 -> RdfType")
	}
	if deviceType.Services[0].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 0/0/0 -> ConceptIds")
	}
	/// service 2
	if deviceType.Services[1].Id != "urn:infai:ses:service:3e3e_2" {
		t.Fatal("error service -> 1 -> id")
	}
	if deviceType.Services[1].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 1 -> RdfType")
	}
	if deviceType.Services[1].Name != "setBrightness2" {
		t.Log(deviceType.Services[1].Name)
		t.Fatal("error service -> 1 -> name")
	}
	if deviceType.Services[1].Description != "" {
		t.Fatal("error service -> 1 -> description")
	}
	if deviceType.Services[1].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 1 -> LocalId")
	}
	if deviceType.Services[1].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 1/0 -> id")
	}
	if deviceType.Services[1].Aspects[0].Name != "Lighting" {
		t.Log(deviceType.Services[1].Aspects[0].Name)
		t.Fatal("error aspect -> 1/0 -> Name")
	}
	if deviceType.Services[1].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 1/0 -> RdfType")
	}
	if deviceType.Services[1].Functions[0].Id != "urn:infai:ses:function:5e5e" {
		t.Fatal("error function -> 1/0 -> id")
	}
	if deviceType.Services[1].Functions[0].Name != "brightnessAdjustment" {
		t.Fatal("error function -> 1/0 -> Name")
	}
	if deviceType.Services[1].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 1/0 -> RdfType")
	}
	if deviceType.Services[1].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 1/0/0 -> ConceptIds")
	}

	if err != nil {
		t.Fatal(deviceType, err, code)
	} else {
		t.Log(deviceType)
	}
}

func TestCreateAndDeleteDeviceTypePart1(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:1"
	devicetype.Name = "Philips Hue Color"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:2",
		Name: "Lamp",
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3a",
		"localId",
		"setBrightness",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4a", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5a", Name: "brightnessAdjustment", ConceptId: "urn:infai:ses:concept:6a", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3b",
		"localId",
		"setBrightness",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4b", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5b", Name: "brightnessAdjustment", ConceptId: "urn:infai:ses:concept:6b", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")
}

func TestCreateAndDeleteDeviceTypePart2(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	err = producer.PublishDeviceTypeDelete("urn:infai:ses:devicetype:1", "sdfdsfsf")
	if err != nil {
		t.Fatal(err)
	}
}

func StartUpScript(t *testing.T) (error, *controller.Controller, database.Database) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.New(conf)
	if err != nil {
		t.Fatal(err)
	}
	_, err = producer.New(conf)
	if err != nil {
		t.Fatal(err)
	}
	con, err := controller.New(conf, db)
	if err != nil {
		t.Fatal(err)
	}
	return err, con, db
}
