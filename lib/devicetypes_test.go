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
	"github.com/satori/go.uuid"
	"testing"
)

var devicetype1id = "urn:infai:ses:device-type:2cc43032-207e-494e-8de4-94784cd4961d"
var devicetype1name = uuid.NewV4().String()
var devicetype2id = uuid.NewV4().String()
var devicetype2name = uuid.NewV4().String()

func TestProduceValidDeviceType(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:1111"
	devicetype.Name = "Philips Hue Color"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:2222",
		Name: "Lamp",
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3333",
		"localId",
		"setBrightness1",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4444", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5555", Name: "brightnessAdjustment", ConceptIds: []string{"urn:infai:ses:concept:6666", "urn:infai:ses:concept:7777"}, RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3333bbbb",
		"localId",
		"setBrightness2",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4444", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5555", Name: "brightnessAdjustment", ConceptIds: []string{"urn:infai:ses:concept:6666", "urn:infai:ses:concept:7777"}, RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")
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
		[]model.Function{{Id: "urn:infai:ses:function:5a", Name: "brightnessAdjustment", ConceptIds: []string{"urn:infai:ses:concept:6a", "urn:infai:ses:concept:7a"}, RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
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
		[]model.Function{{Id: "urn:infai:ses:function:5b", Name: "brightnessAdjustment", ConceptIds: []string{"urn:infai:ses:concept:6b", "urn:infai:ses:concept:7b"}, RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION}},
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

func TestReadControllingFunction(t *testing.T) {
	err, con, _ := StartUpScript(t)
	res, err, code := con.GetFunctions(model.SES_ONTOLOGY_CONTROLLING_FUNCTION)
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}
}

func TestReadMeasuringFunction(t *testing.T) {
	err, con, _ := StartUpScript(t)
	res, err, code := con.GetFunctions(model.SES_ONTOLOGY_MEASURING_FUNCTION)
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}
}

func TestReadAspect(t *testing.T) {
	err, con, db := StartUpScript(t)
	/// Aspect Lightning
	err = db.InsertData(
		`<urn:infai:ses:aspect:4444> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <https://senergy.infai.org/ontology/Aspect> .
<urn:infai:ses:aspect:4444> <http://www.w3.org/2000/01/rdf-schema#label> "Lightning" .`)
	if err != nil {
		t.Fatal(err)
	}
	/// Aspect Air
	err = db.InsertData(
		`<urn:infai:ses:aspect:2222> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <https://senergy.infai.org/ontology/Aspect> .
<urn:infai:ses:aspect:2222> <http://www.w3.org/2000/01/rdf-schema#label> "Air" .`)
	if err != nil {
		t.Fatal(err)
	}
	/// Aspect Connectivity
	err = db.InsertData(
		`<urn:infai:ses:aspect:4545> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <https://senergy.infai.org/ontology/Aspect> .
<urn:infai:ses:aspect:4545> <http://www.w3.org/2000/01/rdf-schema#label> "Connectivity" .`)
	if err != nil {
		t.Fatal(err)
	}
	res, err, code := con.GetAspects()
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}

	if res[0].Id != "urn:infai:ses:aspect:2222" {
		t.Fatal("error id")
	}
	if res[0].Name != "Air" {
		t.Fatal("error Name")
	}

	if res[1].Id != "urn:infai:ses:aspect:4545" {
		t.Fatal("error id")
	}
	if res[1].Name != "Connectivity" {
		t.Fatal("error Name")
	}

	if res[2].Id != "urn:infai:ses:aspect:4444" {
		t.Fatal("error id")
	}
	if res[2].Name != "Lightning" {
		t.Fatal("error Name")
	}


}

func TestReadDeviceClass(t *testing.T) {
	err, con, _ := StartUpScript(t)
	res, err, code := con.GetDeviceClasses()
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}
}

func TestReadDeviceType(t *testing.T) {
	err, con, _ := StartUpScript(t)
	deviceType, err, code := con.GetDeviceType("urn:infai:ses:devicetype:1111")

	if deviceType.Id != "urn:infai:ses:devicetype:1111" {
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
	if deviceType.DeviceClass.Id != "urn:infai:ses:deviceclass:2222" {
		t.Fatal("error deviceclass id")
	}
	if deviceType.DeviceClass.Name != "Lamp" {
		t.Fatal("error deviceclass name")
	}
	if deviceType.DeviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("error deviceclass rdf type")
	}
	// Service
	if deviceType.Services[0].Id != "urn:infai:ses:service:3333" {
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
	if deviceType.Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4444" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType.Services[0].Aspects[0].Name != "Lighting" {
		t.Log(deviceType.Services[0].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType.Services[0].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}
	if deviceType.Services[0].Functions[0].Id != "urn:infai:ses:function:5555" {
		t.Fatal("error function -> 0/0 -> id")
	}
	if deviceType.Services[0].Functions[0].Name != "brightnessAdjustment" {
		t.Fatal("error function -> 0/0 -> Name")
	}
	if deviceType.Services[0].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 0/0 -> RdfType")
	}
	if deviceType.Services[0].Functions[0].ConceptIds[0] != "urn:infai:ses:concept:7777" {
		t.Fatal("error function -> 0/0/0 -> ConceptIds")
	}
	if deviceType.Services[0].Functions[0].ConceptIds[1] != "urn:infai:ses:concept:6666" {
		t.Fatal("error function -> 0/0/1 -> ConceptIds")
	}
	/// service 2
	if deviceType.Services[1].Id != "urn:infai:ses:service:3333bbbb" {
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
	if deviceType.Services[1].Aspects[0].Id != "urn:infai:ses:aspect:4444" {
		t.Fatal("error aspect -> 1/0 -> id")
	}
	if deviceType.Services[1].Aspects[0].Name != "Lighting" {
		t.Log(deviceType.Services[1].Aspects[0].Name)
		t.Fatal("error aspect -> 1/0 -> Name")
	}
	if deviceType.Services[1].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 1/0 -> RdfType")
	}
	if deviceType.Services[1].Functions[0].Id != "urn:infai:ses:function:5555" {
		t.Fatal("error function -> 1/0 -> id")
	}
	if deviceType.Services[1].Functions[0].Name != "brightnessAdjustment" {
		t.Fatal("error function -> 1/0 -> Name")
	}
	if deviceType.Services[1].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 1/0 -> RdfType")
	}
	if deviceType.Services[1].Functions[0].ConceptIds[0] != "urn:infai:ses:concept:7777" {
		t.Fatal("error function -> 1/0/0 -> ConceptIds")
	}
	if deviceType.Services[1].Functions[0].ConceptIds[1] != "urn:infai:ses:concept:6666" {
		t.Fatal("error function -> 1/0/1 -> ConceptIds")
	}

	if err != nil {
		t.Fatal(deviceType, err, code)
	} else {
		t.Log(deviceType)
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
	producer, err := producer.New(conf)
	if err != nil {
		t.Fatal(err)
	}
	con, err := controller.New(conf, db, producer)
	if err != nil {
		t.Fatal(err)
	}
	return err, con, db
}
