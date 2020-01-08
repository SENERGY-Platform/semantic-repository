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

package lib

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/SENERGY-Platform/semantic-repository/lib/source/producer"
	"testing"
)

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

func TestProduceDeviceTypeforAspectTest(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:1e1e-AspectTest"
	devicetype.Name = "Philips Hue Color"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:2e2e-AspectTest",
		Name: "Lamp",
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3e3e-AspectTest",
		"localId",
		"setBrightness",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4e4e-AspectTest", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5e5e1-AspectTest", Name: "brightnessAdjustment1", ConceptId: "urn:ses:infai:concept:1a1a1a", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION},
			{Id: "urn:infai:ses:function:5e5e2-AspectTest", Name: "brightnessAdjustment2", ConceptId: "urn:ses:infai:concept:1a1a1a", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3f3f-AspectTest",
		"localId",
		"getBrightness",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4e4e-AspectTest", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5e5e3-AspectTest", Name: "brightnessFunction4", ConceptId: "urn:ses:infai:concept:1a1a1a", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION},
			{Id: "urn:infai:ses:function:5e5e4-AspectTest", Name: "brightnessFunction2", ConceptId: "urn:ses:infai:concept:1a1a1a", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")

}

func TestReadAspectMeasuringFunctions(t *testing.T) {
	err, con, _ := StartUpScript(t)

	res, err, code := con.GetAspectsMeasuringFunctions("urn:infai:ses:aspect:4e4e-AspectTest")
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}
	if res[0].Id != "urn:infai:ses:function:5e5e4-AspectTest" {
		t.Fatal("error id", res[0].Id)
	}
	if res[0].Name != "brightnessFunction2" {
		t.Fatal("error Name")
	}
	if res[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error ConceptId")
	}
	if res[0].RdfType != model.SES_ONTOLOGY_MEASURING_FUNCTION {
		t.Fatal("wrong RdfType")
	}

	if res[1].Id != "urn:infai:ses:function:5e5e3-AspectTest" {
		t.Fatal("error id", res[1].Id)
	}
	if res[1].Name != "brightnessFunction4" {
		t.Fatal("error Name")
	}
	if res[1].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error ConceptId")
	}
	if res[1].RdfType != model.SES_ONTOLOGY_MEASURING_FUNCTION {
		t.Fatal("wrong RdfType")
	}

}

func Test_2_ProduceDeviceTypeforAspectTest(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:08-01-20"
	devicetype.Name = "Philips Hue Color"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:08-01-20",
		Name: "Lamp",
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:08-01-20_1",
		"localId",
		"setBrightness",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:08-01-20_1", Name: "aspect1", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:08-01-20_1", Name: "func1", ConceptId: "urn:ses:infai:concept:08-01-20", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION},
			{Id: "urn:infai:ses:function:08-01-20_2", Name: "func1", ConceptId: "urn:ses:infai:concept:08-01-20", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:08-01-20_2",
		"localId",
		"getBrightness",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:08-01-20_2", Name: "aspect2", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:08-01-20_3", Name: "func3", ConceptId: "urn:ses:infai:concept:08-01-20", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION},
			{Id: "urn:infai:ses:function:08-01-20_4", Name: "func4", ConceptId: "urn:ses:infai:concept:08-01-20", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")

}

func Test_2_ReadAspectsWithMeasuringFunctions(t *testing.T) {
	err, con, _ := StartUpScript(t)

	res, err, code := con.GetAspectsWithMeasuringFunction()
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}
	if res[0].Id != "urn:infai:ses:aspect:08-01-20_2" {
		t.Fatal("error id", res[0].Id)
	}
	if res[0].Name != "aspect2" {
		t.Fatal("error Name")
	}
	if res[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("wrong RdfType")
	}

}
