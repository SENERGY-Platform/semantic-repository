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
	"encoding/json"
	"fmt"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/controller"
	"github.com/SENERGY-Platform/semantic-repository/lib/database"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/SENERGY-Platform/semantic-repository/lib/source/producer"
	"testing"
)

func TestProduceValidDeviceTypes(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:device-type:eb4a3337-01a1-4434-9dcc-064b3955eeef"
	devicetype.Name = "Philips-Extended-Color-Light"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:device-class:14e56881-16f9-4120-bb41-270a43070c86",
		Name: "Lamp",
	}
	devicetype.Description = "Philips Hue Extended Color Light"
	devicetype.Image = "https://i.imgur.com/OZOqLcR.png"
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:1b0ef253-16f7-4b65-8a15-fe79fccf7e70",
		"setColor",
		"setColorService",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:a7470d73-dde3-41fc-92bd-f16bb28f2da6", Name: "Lighting", RdfType: "https://senergy.infai.org/ontology/Aspect"}},
		"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:controlling-function:c54e2a89-1fb8-4ecb-8993-a7b40b355599", Name: "setColorFunction", ConceptId: "urn:infai:ses:concept:8b1161d5-7878-4dd2-a36c-6f98f6b94bf8", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")

	////////////////////////////////
	/// DANFOSS THERMOSTAT       ///
	////////////////////////////////

	devicetype = model.DeviceType{}
	devicetype.Id = "urn:infai:ses:device-type:662d9c9f-949d-4577-9485-9cb7255f547f"
	devicetype.Name = "Danfoss Radiator Thermostat"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:device-class:997937d6-c5f3-4486-b67c-114675038393",
		Name: "Thermostat",
	}
	devicetype.Description = ""
	devicetype.Image = ""
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:de9252b9-5492-4fe5-8c9c-b4b8460f65f6",
		"exact:67-1",
		"setTemperatureService",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6", Name: "Air", RdfType: "https://senergy.infai.org/ontology/Aspect"}},
		"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:controlling-function:99240d90-02dd-4d4f-a47c-069cfe77629c", Name: "setTemperatureFunction", ConceptId: "urn:infai:ses:concept:0bc81398-3ed6-4e2b-a6c4-b754583aac37", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:f306de41-a55b-45ed-afc9-039bbe53db1b",
		"get_level:67-1",
		"getTemperatureService",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6", Name: "Air", RdfType: "https://senergy.infai.org/ontology/Aspect"}},
		"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:measuring-function:f2769eb9-b6ad-4f7e-bd28-e4ea043d2f8b", Name: "getTemperatureFunction", ConceptId: "urn:infai:ses:concept:0bc81398-3ed6-4e2b-a6c4-b754583aac37", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")

	////////////////////////////////
	/// BLEBOX                   ///
	////////////////////////////////

	devicetype = model.DeviceType{}
	devicetype.Id = "urn:infai:ses:device-type:a8cbd322-9d8c-4f4c-afec-ae4b7986b6ed"
	devicetype.Name = "Blebox-Air-Sensor"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:device-class:8bd38ea2-1835-4a1e-ac02-6b3169513fd3",
		Name: "AirQualityMeter",
	}
	devicetype.Description = ""
	devicetype.Image = ""
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:422fd899-a2cc-4e43-8d81-4e330a7ca8ab",
		"reading_pm10",
		"getParticleAmountPM10Service",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6", Name: "Air", RdfType: "https://senergy.infai.org/ontology/Aspect"}},
		"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0", Name: "getParticleAmountPM10", ConceptId: "urn:infai:ses:concept:a63a960d-1f36-4d95-ad5b-dfb4a7fe3b5b", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION}},
		"asdasdsadsadasd",
	})

	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:1d20a68b-7136-456c-ace5-c3adb66866bf",
		"reading_pm1",
		"getParticleAmountPM1Service",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6", Name: "Air", RdfType: "https://senergy.infai.org/ontology/Aspect"}},
		"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:measuring-function:0e19d094-70c6-402c-8523-3aaff2ce6dd9", Name: "getParticleAmountPM1", ConceptId: "urn:infai:ses:concept:a63a960d-1f36-4d95-ad5b-dfb4a7fe3b5b", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")

}

func TestReadDeviceType(t *testing.T) {

	err, con, _ := StartUpScript(t)
	deviceType, err, code := con.GetDeviceType("urn:infai:ses:device-type:eb4a3337-01a1-4434-9dcc-064b3955eeef")

	deviceTypeStringified := `{"id":"urn:infai:ses:device-type:eb4a3337-01a1-4434-9dcc-064b3955eeef","name":"Philips-Extended-Color-Light","description":"","image":"","services":[{"id":"urn:infai:ses:service:1b0ef253-16f7-4b65-8a15-fe79fccf7e70","local_id":"","name":"setColorService","description":"","aspects":[{"id":"urn:infai:ses:aspect:a7470d73-dde3-41fc-92bd-f16bb28f2da6","name":"Lighting","rdf_type":"https://senergy.infai.org/ontology/Aspect"}],"protocol_id":"","inputs":null,"outputs":null,"functions":[{"id":"urn:infai:ses:controlling-function:c54e2a89-1fb8-4ecb-8993-a7b40b355599","name":"setColorFunction","concept_id":"urn:infai:ses:concept:8b1161d5-7878-4dd2-a36c-6f98f6b94bf8","rdf_type":"https://senergy.infai.org/ontology/ControllingFunction"}],"rdf_type":"https://senergy.infai.org/ontology/Service"}],"device_class":{"id":"urn:infai:ses:device-class:14e56881-16f9-4120-bb41-270a43070c86","name":"Lamp","rdf_type":"https://senergy.infai.org/ontology/DeviceClass"},"rdf_type":"https://senergy.infai.org/ontology/DeviceType"}`

	b, err := json.Marshal(deviceType)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(string(b))

	if err != nil {
		t.Fatal(deviceType, err, code)
	} else {
		if string(b) != deviceTypeStringified {
			t.Log("expected:", deviceTypeStringified)
			t.Log("was:", string(b))
			t.Fatal("error")
		}
	}
}

func TestReadDeviceTypeCF(t *testing.T) {
	// ControllingFunctionId + DeviceClassId
	err, con, _ := StartUpScript(t)
	deviceType, err, code := con.GetDeviceTypesFiltered([]model.DeviceTypesFilter{{FunctionId: "urn:infai:ses:controlling-function:c54e2a89-1fb8-4ecb-8993-a7b40b355599", DeviceClassId: "urn:infai:ses:device-class:14e56881-16f9-4120-bb41-270a43070c86", AspectId: ""}})

	deviceTypeStringified := `[{"id":"urn:infai:ses:device-type:eb4a3337-01a1-4434-9dcc-064b3955eeef","name":"Philips-Extended-Color-Light","description":"","image":"","services":[{"id":"urn:infai:ses:service:1b0ef253-16f7-4b65-8a15-fe79fccf7e70","local_id":"","name":"setColorService","description":"","aspects":[{"id":"urn:infai:ses:aspect:a7470d73-dde3-41fc-92bd-f16bb28f2da6","name":"Lighting","rdf_type":"https://senergy.infai.org/ontology/Aspect"}],"protocol_id":"","inputs":null,"outputs":null,"functions":[{"id":"urn:infai:ses:controlling-function:c54e2a89-1fb8-4ecb-8993-a7b40b355599","name":"setColorFunction","concept_id":"urn:infai:ses:concept:8b1161d5-7878-4dd2-a36c-6f98f6b94bf8","rdf_type":"https://senergy.infai.org/ontology/ControllingFunction"}],"rdf_type":"https://senergy.infai.org/ontology/Service"}],"device_class":{"id":"urn:infai:ses:device-class:14e56881-16f9-4120-bb41-270a43070c86","name":"Lamp","rdf_type":"https://senergy.infai.org/ontology/DeviceClass"},"rdf_type":"https://senergy.infai.org/ontology/DeviceType"}]`

	b, err := json.Marshal(deviceType)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(string(b))

	if err != nil {
		t.Fatal(deviceType, err, code)
	} else {
		if string(b) != deviceTypeStringified {
			t.Log("expected:", deviceTypeStringified)
			t.Log("was:", string(b))
			t.Fatal("error")
		}
	}
}

func TestReadDeviceType_1MF(t *testing.T) {
	// 1 MeasuringFunctionId + Aspect
	err, con, _ := StartUpScript(t)
	//deviceType, err, code := con.GetDeviceTypesFiltered("", []string{"urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0", "urn:infai:ses:measuring-function:0e19d094-70c6-402c-8523-3aaff2ce6dd9"}, []string{})
	deviceType, err, code := con.GetDeviceTypesFiltered([]model.DeviceTypesFilter{{FunctionId: "urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0", DeviceClassId: "", AspectId: "urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"}})

	deviceTypeStringified := `[{"id":"urn:infai:ses:device-type:a8cbd322-9d8c-4f4c-afec-ae4b7986b6ed","name":"Blebox-Air-Sensor","description":"","image":"","services":[{"id":"urn:infai:ses:service:422fd899-a2cc-4e43-8d81-4e330a7ca8ab","local_id":"","name":"getParticleAmountPM10Service","description":"","aspects":[{"id":"urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6","name":"Air","rdf_type":"https://senergy.infai.org/ontology/Aspect"}],"protocol_id":"","inputs":null,"outputs":null,"functions":[{"id":"urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0","name":"getParticleAmountPM10","concept_id":"urn:infai:ses:concept:a63a960d-1f36-4d95-ad5b-dfb4a7fe3b5b","rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction"}],"rdf_type":"https://senergy.infai.org/ontology/Service"}],"device_class":{"id":"urn:infai:ses:device-class:8bd38ea2-1835-4a1e-ac02-6b3169513fd3","name":"AirQualityMeter","rdf_type":"https://senergy.infai.org/ontology/DeviceClass"},"rdf_type":"https://senergy.infai.org/ontology/DeviceType"}]`

	b, err := json.Marshal(deviceType)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(string(b))

	if err != nil {
		t.Fatal(deviceType, err, code)
	} else {
		if string(b) != deviceTypeStringified {
			t.Log("expected:", deviceTypeStringified)
			t.Log("was:", string(b))
			t.Fatal("error")
		}
	}
}

func TestReadDeviceType_2MF_sameAspect(t *testing.T) {
	// 2 MeasuringFunctionId + same Aspect
	err, con, _ := StartUpScript(t)
	//deviceType, err, code := con.GetDeviceTypesFiltered("", []string{"urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0", "urn:infai:ses:measuring-function:0e19d094-70c6-402c-8523-3aaff2ce6dd9"}, []string{})
	deviceType, err, code := con.GetDeviceTypesFiltered([]model.DeviceTypesFilter{{FunctionId: "urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0", DeviceClassId: "", AspectId: "urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"},
		{FunctionId: "urn:infai:ses:measuring-function:0e19d094-70c6-402c-8523-3aaff2ce6dd9", DeviceClassId: "", AspectId: "urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"}})

	deviceTypeStringified := `[{"id":"urn:infai:ses:device-type:a8cbd322-9d8c-4f4c-afec-ae4b7986b6ed","name":"Blebox-Air-Sensor","description":"","image":"","services":[{"id":"urn:infai:ses:service:422fd899-a2cc-4e43-8d81-4e330a7ca8ab","local_id":"","name":"getParticleAmountPM10Service","description":"","aspects":[{"id":"urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6","name":"Air","rdf_type":"https://senergy.infai.org/ontology/Aspect"}],"protocol_id":"","inputs":null,"outputs":null,"functions":[{"id":"urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0","name":"getParticleAmountPM10","concept_id":"urn:infai:ses:concept:a63a960d-1f36-4d95-ad5b-dfb4a7fe3b5b","rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction"}],"rdf_type":"https://senergy.infai.org/ontology/Service"},{"id":"urn:infai:ses:service:1d20a68b-7136-456c-ace5-c3adb66866bf","local_id":"","name":"getParticleAmountPM1Service","description":"","aspects":[{"id":"urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6","name":"Air","rdf_type":"https://senergy.infai.org/ontology/Aspect"}],"protocol_id":"","inputs":null,"outputs":null,"functions":[{"id":"urn:infai:ses:measuring-function:0e19d094-70c6-402c-8523-3aaff2ce6dd9","name":"getParticleAmountPM1","concept_id":"urn:infai:ses:concept:a63a960d-1f36-4d95-ad5b-dfb4a7fe3b5b","rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction"}],"rdf_type":"https://senergy.infai.org/ontology/Service"}],"device_class":{"id":"urn:infai:ses:device-class:8bd38ea2-1835-4a1e-ac02-6b3169513fd3","name":"AirQualityMeter","rdf_type":"https://senergy.infai.org/ontology/DeviceClass"},"rdf_type":"https://senergy.infai.org/ontology/DeviceType"}]`

	b, err := json.Marshal(deviceType)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(string(b))

	if err != nil {
		t.Fatal(deviceType, err, code)
	} else {
		if string(b) != deviceTypeStringified {
			t.Log("expected:", deviceTypeStringified)
			t.Log("was:", string(b))
			t.Fatal("error")
		}
	}
}

func TestReadDeviceType_2MF_sameAspect_DifferentDeviceClasses(t *testing.T) {
	// 2 MeasuringFunctionId + same Aspect + 2 different DeviceClasses
	err, con, _ := StartUpScript(t)
	//deviceType, err, code := con.GetDeviceTypesFiltered("", []string{"urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0", "urn:infai:ses:measuring-function:0e19d094-70c6-402c-8523-3aaff2ce6dd9"}, []string{})
	deviceType, err, code := con.GetDeviceTypesFiltered([]model.DeviceTypesFilter{{FunctionId: "urn:infai:ses:measuring-function:f2769eb9-b6ad-4f7e-bd28-e4ea043d2f8b", DeviceClassId: "", AspectId: "urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"},
		{FunctionId: "urn:infai:ses:measuring-function:0e19d094-70c6-402c-8523-3aaff2ce6dd9", DeviceClassId: "", AspectId: "urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"}})

	b, err := json.Marshal(deviceType)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(string(b))

	if err != nil {
		t.Fatal(deviceType, err, code)
	} else {
		if string(b) != "null" {
			t.Log("expected: null")
			t.Log("was:", string(b))
			t.Fatal("error")
		}
	}
}

func TestReadDeviceTypeWithId1(t *testing.T) {
	err, con, _ := StartUpScript(t)
	deviceType, err, code := con.GetDeviceType("urn:infai:ses:device-type:eb4a3337-01a1-4434-9dcc-064b3955eeef")

	if deviceType.Id != "urn:infai:ses:device-type:eb4a3337-01a1-4434-9dcc-064b3955eeef" {
		t.Fatal("error id")
	}

	if deviceType.RdfType != model.SES_ONTOLOGY_DEVICE_TYPE {
		t.Fatal("error model")
	}

	if deviceType.Name != "Philips-Extended-Color-Light" {
		t.Fatal("error name")
	}

	if deviceType.Description != "" {
		t.Fatal("error description")
	}

	if deviceType.Image != "" {
		t.Fatal("error image")
	}
	// DeviceClass
	if deviceType.DeviceClass.Id != "urn:infai:ses:device-class:14e56881-16f9-4120-bb41-270a43070c86" {
		t.Fatal("error deviceclass id")
	}
	if deviceType.DeviceClass.Name != "Lamp" {
		t.Fatal("error deviceclass name")
	}
	if deviceType.DeviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("error deviceclass rdf type")
	}
	// Service
	if deviceType.Services[0].Id != "urn:infai:ses:service:1b0ef253-16f7-4b65-8a15-fe79fccf7e70" {
		t.Fatal("error service -> 0 -> id")
	}
	if deviceType.Services[0].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType.Services[0].Name != "setColorService" {
		t.Log(deviceType.Services[0].Name)
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType.Services[0].Description != "" {
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType.Services[0].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType.Services[0].Aspects[0].Id != "urn:infai:ses:aspect:a7470d73-dde3-41fc-92bd-f16bb28f2da6" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType.Services[0].Aspects[0].Name != "Lighting" {
		t.Log(deviceType.Services[0].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType.Services[0].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}
	if deviceType.Services[0].Functions[0].Id != "urn:infai:ses:controlling-function:c54e2a89-1fb8-4ecb-8993-a7b40b355599" {
		t.Fatal("error function -> 0/0 -> id")
	}
	if deviceType.Services[0].Functions[0].Name != "setColorFunction" {
		t.Fatal("error function -> 0/0 -> Name")
	}
	if deviceType.Services[0].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 0/0 -> RdfType")
	}
	if deviceType.Services[0].Functions[0].ConceptId != "urn:infai:ses:concept:8b1161d5-7878-4dd2-a36c-6f98f6b94bf8" {
		t.Fatal("error function -> 0/0/0 -> ConceptIds")
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

func TestProduceValidDeviceTypeWithoutConceptId(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:1_4-12-2019"
	devicetype.Name = "Philips Hue Color"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:2_4-12-2019",
		Name: "Lamp",
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3_4-12-2019",
		"localId",
		"setBrightness2",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4_4-12-2019", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5_4-12-2019", Name: "brightnessAdjustment", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")
}

func TestReadDeviceTypeWithoutConceptId(t *testing.T) {
	err, con, _ := StartUpScript(t)
	deviceType, err, code := con.GetDeviceType("urn:infai:ses:devicetype:1_4-12-2019")

	t.Log(deviceType)

	if deviceType.Id != "urn:infai:ses:devicetype:1_4-12-2019" {
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
	if deviceType.DeviceClass.Id != "urn:infai:ses:deviceclass:2_4-12-2019" {
		t.Fatal("error deviceclass id")
	}
	if deviceType.DeviceClass.Name != "Lamp" {
		t.Fatal("error deviceclass name")
	}
	if deviceType.DeviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("error deviceclass rdf type")
	}
	// Service
	if deviceType.Services[0].Id != "urn:infai:ses:service:3_4-12-2019" {
		t.Fatal("error service -> 0 -> id")
	}
	if deviceType.Services[0].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType.Services[0].Name != "setBrightness2" {
		t.Log(deviceType.Services[0].Name)
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType.Services[0].Description != "" {
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType.Services[0].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType.Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4_4-12-2019" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType.Services[0].Aspects[0].Name != "Lighting" {
		t.Log(deviceType.Services[0].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType.Services[0].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}
	if deviceType.Services[0].Functions[0].Id != "urn:infai:ses:function:5_4-12-2019" {
		t.Fatal("error function -> 0/0 -> id")
	}
	if deviceType.Services[0].Functions[0].Name != "brightnessAdjustment" {
		t.Fatal("error function -> 0/0 -> Name")
	}
	if deviceType.Services[0].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 0/0 -> RdfType")
	}
	if deviceType.Services[0].Functions[0].ConceptId != "" {
		t.Fatal("error function -> 0/0/0 -> ConceptIds", deviceType.Services[0].Functions[0].ConceptId)
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
