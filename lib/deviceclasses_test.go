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
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/SENERGY-Platform/semantic-repository/lib/source/producer"
	"testing"
)


func TestReadDeviceClass(t *testing.T) {
	err, con, db := StartUpScript(t)
	/// DeviceClass Lightning
	err = db.InsertData(
		`<urn:infai:ses:device-class:123> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <` +  model.SES_ONTOLOGY_DEVICE_CLASS + `> .
<urn:infai:ses:device-class:123> <http://www.w3.org/2000/01/rdf-schema#label> "Ventilator" .`)
	if err != nil {
		t.Fatal(err)
	}
	err = db.InsertData(
		`<urn:infai:ses:device-class:234> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <` +  model.SES_ONTOLOGY_DEVICE_CLASS + `> .
<urn:infai:ses:device-class:234> <http://www.w3.org/2000/01/rdf-schema#label> "ElectricityMeter" .`)
	if err != nil {
		t.Fatal(err)
	}
	err = db.InsertData(
		`<urn:infai:ses:device-class:111> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <` +  model.SES_ONTOLOGY_DEVICE_CLASS + `> .
<urn:infai:ses:device-class:111> <http://www.w3.org/2000/01/rdf-schema#label> "CarbonDioxideMeter" .`)
	if err != nil {
		t.Fatal(err)
	}
	res, err, code := con.GetDeviceClasses()
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}
	if res[0].Id != "urn:infai:ses:device-class:111" {
		t.Fatal("error id")
	}
	if res[0].Name != "CarbonDioxideMeter" {
		t.Fatal("error Name")
	}
	if res[1].Id != "urn:infai:ses:device-class:234" {
		t.Fatal("error id")
	}
	if res[1].Name != "ElectricityMeter" {
		t.Fatal("error Name")
	}
	if res[2].Id != "urn:infai:ses:device-class:123" {
		t.Fatal("error id")
	}
	if res[2].Name != "Ventilator" {
		t.Fatal("error Name")
	}

}

func TestProduceDeviceType(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:04-11-2019"
	devicetype.Name = "Philips Hue Color"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:04-11-2019",
		Name: "Lamp",
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:04-11-2019_1",
		"localId",
		"setBrightness_1",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:04-11-2019_1", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:04-11-2019_1", Name: "brightnessAdjustment_1", ConceptId: "urn:ses:infai:concept:04-11-2019_1", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:04-11-2019_2",
		"localId",
		"setBrightness_2",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:04-11-2019_2", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:04-11-2019_2", Name: "brightnessAdjustment_2", ConceptId: "urn:ses:infai:concept:04-11-2019_2", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")
}

func TestReadDeviceType1(t *testing.T) {
	err, con, _ := StartUpScript(t)
	functions, err, code := con.GetDeviceClassesFunctions("urn:infai:ses:deviceclass:04-11-2019")

	if functions[0].Id != "urn:infai:ses:function:04-11-2019_1" {
		t.Fatal("error id")
	}

	if functions[0].Name != "brightnessAdjustment_1" {
		t.Fatal("error Name")
	}

	if functions[0].ConceptId != "urn:ses:infai:concept:04-11-2019_1" {
		t.Fatal("error ConceptId")
	}

	if functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error RdfType")
	}


	if functions[1].Id != "urn:infai:ses:function:04-11-2019_2" {
		t.Fatal("error id")
	}

	if functions[1].Name != "brightnessAdjustment_2" {
		t.Fatal("error Name")
	}

	if functions[1].ConceptId != "urn:ses:infai:concept:04-11-2019_2" {
		t.Fatal("error ConceptId")
	}

	if functions[1].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error RdfType")
	}

	if err != nil {
		t.Fatal(functions, err, code)
	} else {
		t.Log(functions)
	}
}
