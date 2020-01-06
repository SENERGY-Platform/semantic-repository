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

func TestReadDeviceClasses(t *testing.T) {
	err, con, db := StartUpScript(t)
	/// DeviceClass Lightning
	err = db.InsertData(
		`<urn:infai:ses:device-class:123> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <` + model.SES_ONTOLOGY_DEVICE_CLASS + `> .
<urn:infai:ses:device-class:123> <http://www.w3.org/2000/01/rdf-schema#label> "Ventilator" .`)
	if err != nil {
		t.Fatal(err)
	}
	err = db.InsertData(
		`<urn:infai:ses:device-class:234> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <` + model.SES_ONTOLOGY_DEVICE_CLASS + `> .
<urn:infai:ses:device-class:234> <http://www.w3.org/2000/01/rdf-schema#label> "ElectricityMeter" .`)
	if err != nil {
		t.Fatal(err)
	}
	err = db.InsertData(
		`<urn:infai:ses:device-class:111> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <` + model.SES_ONTOLOGY_DEVICE_CLASS + `> .
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

func TestProduceDeviceTypeforDeviceClassTest(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:1e1e-DeviceClassTest"
	devicetype.Name = "Philips Hue Color"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:2e2e-DeviceClassTest",
		Name: "Lamp",
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3e3e-DeviceClassTest",
		"localId",
		"setBrightness",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4e4e-DeviceClassTest", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5e5e1-DeviceClassTest", Name: "brightnessAdjustment1", ConceptId: "urn:ses:infai:concept:1a1a1a", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION},
			{Id: "urn:infai:ses:function:5e5e2-DeviceClassTest", Name: "brightnessAdjustment2", ConceptId: "urn:ses:infai:concept:1a1a1a", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3f3f-DeviceClassTest",
		"localId",
		"getBrightness",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4e4e-DeviceClassTest", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5e5e3-DeviceClassTest", Name: "brightnessFunction4", ConceptId: "urn:ses:infai:concept:1a1a1a", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION},
			{Id: "urn:infai:ses:function:5e5e4-DeviceClassTest", Name: "brightnessFunction2", ConceptId: "urn:ses:infai:concept:1a1a1a", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")

}

func TestReadDeviceClassFunctions(t *testing.T) {
	err, con, _ := StartUpScript(t)

	res, err, code := con.GetDeviceClassesFunctions("urn:infai:ses:deviceclass:2e2e-DeviceClassTest")
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		//t.Log(res)
	}
	if res[0].Id != "urn:infai:ses:function:5e5e1-DeviceClassTest" {
		t.Fatal("error id", res[0].Id)
	}
	if res[0].Name != "brightnessAdjustment1" {
		t.Fatal("error Name")
	}
	if res[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error ConceptId")
	}
	if res[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("wrong RdfType")
	}

	if res[1].Id != "urn:infai:ses:function:5e5e2-DeviceClassTest" {
		t.Fatal("error id", res[1].Id)
	}
	if res[1].Name != "brightnessAdjustment2" {
		t.Fatal("error Name")
	}
	if res[1].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error ConceptId")
	}
	if res[1].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("wrong RdfType")
	}

	if res[2].Id != "urn:infai:ses:function:5e5e4-DeviceClassTest" {
		t.Fatal("error id", res[2].Id)
	}
	if res[2].Name != "brightnessFunction2" {
		t.Fatal("error Name")
	}
	if res[2].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error ConceptId")
	}
	if res[2].RdfType != model.SES_ONTOLOGY_MEASURING_FUNCTION {
		t.Fatal("wrong RdfType")
	}

	if res[3].Id != "urn:infai:ses:function:5e5e3-DeviceClassTest" {
		t.Fatal("error id", res[3].Id)
	}
	if res[3].Name != "brightnessFunction4" {
		t.Fatal("error Name")
	}
	if res[3].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error ConceptId")
	}
	if res[3].RdfType != model.SES_ONTOLOGY_MEASURING_FUNCTION {
		t.Fatal("wrong RdfType")
	}

}

func TestReadDeviceClassControllingFunctions(t *testing.T) {
	err, con, _ := StartUpScript(t)

	res, err, code := con.GetDeviceClassesControllingFunctions("urn:infai:ses:deviceclass:2e2e-DeviceClassTest")
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		//t.Log(res)
	}
	if res[0].Id != "urn:infai:ses:function:5e5e1-DeviceClassTest" {
		t.Fatal("error id", res[0].Id)
	}
	if res[0].Name != "brightnessAdjustment1" {
		t.Fatal("error Name")
	}
	if res[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error ConceptId")
	}
	if res[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("wrong RdfType")
	}

	if res[1].Id != "urn:infai:ses:function:5e5e2-DeviceClassTest" {
		t.Fatal("error id", res[1].Id)
	}
	if res[1].Name != "brightnessAdjustment2" {
		t.Fatal("error Name")
	}
	if res[1].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error ConceptId")
	}
	if res[1].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("wrong RdfType")
	}

}

func Test_2_ProduceDeviceTypeforDeviceClassTest(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:test_2_06-01-2020"
	devicetype.Name = "Philips Hue Color"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:test_2_06-01-2020",
		Name: "Lamp",
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:test_2_06-01-2020",
		"localId",
		"setBrightness",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:test_2_06-01-2020", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:test_2_06-01-2020-f1", Name: "func1", ConceptId: "urn:ses:infai:concept:test_2_06-01-2020", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION},
			{Id: "urn:infai:ses:function:test_2_06-01-2020-f2", Name: "func2", ConceptId: "urn:ses:infai:concept:test_2_06-01-2020", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")

	devicetype.Id = "urn:infai:ses:devicetype:test_2_06-01-2020-2"
	devicetype.Name = "Philips Hue Color"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:test_2_06-01-2020-2",
		Name: "Lamp2",
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:test_2_06-01-2020-2",
		"localId",
		"setBrightness",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:test_2_06-01-2020-2", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:test_2_06-01-2020-2-f1", Name: "func1", ConceptId: "urn:ses:infai:concept:test_2_06-01-2020-2", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION},
			{Id: "urn:infai:ses:function:test_2_06-01-2020-2-f2", Name: "func2", ConceptId: "urn:ses:infai:concept:test_2_06-01-2020-2", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")

	devicetype.Id = "urn:infai:ses:devicetype:test_2_06-01-2020-3"
	devicetype.Name = "Philips Hue Color"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:test_2_06-01-2020-3",
		Name: "Lamp3",
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:test_2_06-01-2020-3",
		"localId",
		"setBrightness",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:test_2_06-01-2020-3", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:test_2_06-01-2020-3-f1", Name: "func1", ConceptId: "urn:ses:infai:concept:test_2_06-01-2020-3", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION},
			{Id: "urn:infai:ses:function:test_2_06-01-2020-3-f2", Name: "func2", ConceptId: "urn:ses:infai:concept:test_2_06-01-2020-3", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")

}

func Test_2_ReadDeviceClassesWithControllingFunctions(t *testing.T) {
	err, con, _ := StartUpScript(t)

	res, err, code := con.GetDeviceClassesWithControllingFunctions()
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		//t.Log(res)
	}
	if res[0].Id != "urn:infai:ses:deviceclass:test_2_06-01-2020-2" {
		t.Fatal("error id", res[0].Id)
	}
	if res[0].Name != "Lamp2" {
		t.Fatal("error Name")
	}
	if res[0].RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("wrong RdfType")
	}

	if res[1].Id != "urn:infai:ses:deviceclass:test_2_06-01-2020-3" {
		t.Fatal("error id", res[0].Id)
	}
	if res[1].Name != "Lamp3" {
		t.Fatal("error Name")
	}
	if res[1].RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("wrong RdfType")
	}

}
