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
	"context"
	"encoding/json"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/controller"
	"github.com/SENERGY-Platform/semantic-repository/lib/database"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/SENERGY-Platform/semantic-repository/lib/testutil"
	"github.com/SENERGY-Platform/semantic-repository/lib/testutil/producer"
	"sync"
	"testing"
)

func TestDeviceType(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	defer wg.Wait()
	defer cancel()
	err = testutil.GetDockerEnv(ctx, &wg, &conf)
	if err != nil {
		t.Error(err)
		return
	}

	db, err := database.New(conf)
	if err != nil {
		t.Error(err)
		return
	}

	ctrl, err := controller.New(conf, db)
	if err != nil {
		t.Error(err)
		return
	}

	prod, err := testutil.StartSourceMock(conf, ctrl)
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("testProduceValidDeviceTypes", testProduceValidDeviceTypes(prod))
	t.Run("testReadDeviceType", testReadDeviceType(ctrl))
	t.Run("testReadDeviceTypeCF", testReadDeviceTypeCF(ctrl))
	t.Run("testReadDeviceType_1MF", testReadDeviceType_1MF(ctrl))
	t.Run("testReadDeviceType_1MF_variante2", testReadDeviceType_1MF_variante2(ctrl))
	t.Run("testReadDeviceType_2MF_sameAspect", testReadDeviceType_2MF_sameAspect(ctrl))
	t.Run("testReadDeviceType_2MF_sameAspect_DifferentDeviceClasses", testReadDeviceType_2MF_sameAspect_DifferentDeviceClasses(ctrl))
	t.Run("testReadDeviceTypeWithId1", testReadDeviceTypeWithId1(ctrl))
	t.Run("testCreateAndDeleteDeviceTypePart1", testCreateAndDeleteDeviceTypePart1(prod))
	t.Run("testCreateAndDeleteDeviceTypePart2", testCreateAndDeleteDeviceTypePart2(prod))
}

func testProduceValidDeviceTypes(producer *producer.Producer) func(t *testing.T) {
	return func(t *testing.T) {
		devicetype := model.DeviceType{}
		devicetype.Id = "urn:infai:ses:device-type:eb4a3337-01a1-4434-9dcc-064b3955eeef"
		devicetype.Name = "Philips-Extended-Color-Light"
		devicetype.DeviceClassId = "urn:infai:ses:device-class:14e56881-16f9-4120-bb41-270a43070c86"
		devicetype.Description = "Philips Hue Extended Color Light"
		devicetype.Services = []model.Service{}
		devicetype.Services = append(devicetype.Services, model.Service{
			"urn:infai:ses:service:1b0ef253-16f7-4b65-8a15-fe79fccf7e70",
			"setColor",
			"setColorService",
			"",
			"",
			[]string{"urn:infai:ses:aspect:a7470d73-dde3-41fc-92bd-f16bb28f2da6"},
			"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b",
			[]model.Content{},
			[]model.Content{},
			[]string{"urn:infai:ses:controlling-function:c54e2a89-1fb8-4ecb-8993-a7b40b355599"},
			"asdasdsadsadasd",
		})

		err := producer.PublishDeviceType(devicetype, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}

		////////////////////////////////
		/// DANFOSS THERMOSTAT       ///
		////////////////////////////////

		devicetype = model.DeviceType{}
		devicetype.Id = "urn:infai:ses:device-type:662d9c9f-949d-4577-9485-9cb7255f547f"
		devicetype.Name = "Danfoss Radiator Thermostat"
		devicetype.DeviceClassId = "urn:infai:ses:device-class:997937d6-c5f3-4486-b67c-114675038393"
		devicetype.Description = ""
		devicetype.Services = []model.Service{}
		devicetype.Services = append(devicetype.Services, model.Service{
			"urn:infai:ses:service:de9252b9-5492-4fe5-8c9c-b4b8460f65f6",
			"exact:67-1",
			"setTemperatureService",
			"",
			"",
			[]string{"urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"},
			"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b",
			[]model.Content{},
			[]model.Content{},
			[]string{"urn:infai:ses:controlling-function:99240d90-02dd-4d4f-a47c-069cfe77629c"},
			"asdasdsadsadasd",
		})

		devicetype.Services = append(devicetype.Services, model.Service{
			"urn:infai:ses:service:f306de41-a55b-45ed-afc9-039bbe53db1b",
			"get_level:67-1",
			"getTemperatureService",
			"",
			"",
			[]string{"urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"},
			"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b",
			[]model.Content{},
			[]model.Content{},
			[]string{"urn:infai:ses:measuring-function:f2769eb9-b6ad-4f7e-bd28-e4ea043d2f8b"},
			"asdasdsadsadasd",
		})

		err = producer.PublishDeviceType(devicetype, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}
		////////////////////////////////
		/// CYRUS MULTISENSOR        ///
		////////////////////////////////

		devicetype = model.DeviceType{}
		devicetype.Id = "urn:infai:ses:device-type:3cc09a10-1feb-4f8b-9390-8d08bf3ba22d"
		devicetype.Name = "Cyrus 4-in-1 Multisensor"
		devicetype.DeviceClassId = "urn:infai:ses:device-class:ff64280a-58e6-4cf9-9a44-e70d3831a79d"
		devicetype.Description = ""
		devicetype.Services = []model.Service{}
		devicetype.Services = append(devicetype.Services, model.Service{
			"urn:infai:ses:service:d3dba284-ef6d-4f12-81df-ed11506702b2",
			"get_level:49-1",
			"getTemperatureService",
			"",
			"",
			[]string{"urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"},
			"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b",
			[]model.Content{},
			[]model.Content{},
			[]string{"urn:infai:ses:measuring-function:f2769eb9-b6ad-4f7e-bd28-e4ea043d2f8b"},
			"asdasdsadsadasd",
		})

		err = producer.PublishDeviceType(devicetype, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}
		////////////////////////////////
		/// BLEBOX                   ///
		////////////////////////////////

		devicetype = model.DeviceType{}
		devicetype.Id = "urn:infai:ses:device-type:a8cbd322-9d8c-4f4c-afec-ae4b7986b6ed"
		devicetype.Name = "Blebox-Air-Sensor"
		devicetype.DeviceClassId = "urn:infai:ses:device-class:8bd38ea2-1835-4a1e-ac02-6b3169513fd3"
		devicetype.Description = ""
		devicetype.Services = []model.Service{}
		devicetype.Services = append(devicetype.Services, model.Service{
			"urn:infai:ses:service:422fd899-a2cc-4e43-8d81-4e330a7ca8ab",
			"reading_pm10",
			"getParticleAmountPM10Service",
			"",
			"",
			[]string{"urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"},
			"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b",
			[]model.Content{},
			[]model.Content{},
			[]string{"urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0"},
			"asdasdsadsadasd",
		})

		devicetype.Services = append(devicetype.Services, model.Service{
			"urn:infai:ses:service:1d20a68b-7136-456c-ace5-c3adb66866bf",
			"reading_pm1",
			"getParticleAmountPM1Service",
			"",
			"",
			[]string{"urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"},
			"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b",
			[]model.Content{},
			[]model.Content{},
			[]string{"urn:infai:ses:measuring-function:0e19d094-70c6-402c-8523-3aaff2ce6dd9"},
			"asdasdsadsadasd",
		})

		err = producer.PublishDeviceType(devicetype, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testReadDeviceType(con *controller.Controller) func(t *testing.T) {
	return func(t *testing.T) {
		deviceType, err, code := con.GetDeviceType("urn:infai:ses:device-type:eb4a3337-01a1-4434-9dcc-064b3955eeef")

		deviceTypeStringified := `{"id":"urn:infai:ses:device-type:eb4a3337-01a1-4434-9dcc-064b3955eeef","name":"Philips-Extended-Color-Light","description":"","services":[{"id":"urn:infai:ses:service:1b0ef253-16f7-4b65-8a15-fe79fccf7e70","local_id":"","name":"setColorService","description":"","interaction":"","aspect_ids":["urn:infai:ses:aspect:a7470d73-dde3-41fc-92bd-f16bb28f2da6"],"protocol_id":"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b","inputs":null,"outputs":null,"function_ids":["urn:infai:ses:controlling-function:c54e2a89-1fb8-4ecb-8993-a7b40b355599"],"rdf_type":"https://senergy.infai.org/ontology/Service"}],"device_class_id":"urn:infai:ses:device-class:14e56881-16f9-4120-bb41-270a43070c86","rdf_type":"https://senergy.infai.org/ontology/DeviceType"}`

		if err != nil {
			t.Fatal(deviceType, err, code)
		} else {
			b, err := json.Marshal(deviceType)
			if err != nil {
				t.Fatal(deviceType, err, code)
			}
			if string(b) != deviceTypeStringified {
				t.Log("expected:", deviceTypeStringified)
				t.Log("was:", string(b))
				t.Fatal("error")
			}
		}
	}
}

func testReadDeviceTypeCF(con *controller.Controller) func(t *testing.T) {
	return func(t *testing.T) {
		deviceType, err, code := con.GetDeviceTypesFiltered([]model.DeviceTypesFilter{{FunctionId: "urn:infai:ses:controlling-function:c54e2a89-1fb8-4ecb-8993-a7b40b355599", DeviceClassId: "urn:infai:ses:device-class:14e56881-16f9-4120-bb41-270a43070c86", AspectId: ""}})

		deviceTypeStringified := `[{"id":"urn:infai:ses:device-type:eb4a3337-01a1-4434-9dcc-064b3955eeef","name":"Philips-Extended-Color-Light","description":"","services":[{"id":"urn:infai:ses:service:1b0ef253-16f7-4b65-8a15-fe79fccf7e70","local_id":"","name":"setColorService","description":"","interaction":"","aspect_ids":["urn:infai:ses:aspect:a7470d73-dde3-41fc-92bd-f16bb28f2da6"],"protocol_id":"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b","inputs":null,"outputs":null,"function_ids":["urn:infai:ses:controlling-function:c54e2a89-1fb8-4ecb-8993-a7b40b355599"],"rdf_type":"https://senergy.infai.org/ontology/Service"}],"device_class_id":"urn:infai:ses:device-class:14e56881-16f9-4120-bb41-270a43070c86","rdf_type":"https://senergy.infai.org/ontology/DeviceType"}]`

		if err != nil {
			t.Fatal(deviceType, err, code)
		} else {
			b, err := json.Marshal(deviceType)
			if err != nil {
				t.Fatal(deviceType, err, code)
			}
			if string(b) != deviceTypeStringified {
				t.Log("expected:", deviceTypeStringified)
				t.Log("was:", string(b))
				t.Fatal("error")
			}
		}
	}
}

func testReadDeviceType_1MF(con *controller.Controller) func(t *testing.T) {
	return func(t *testing.T) {
		// 1 MeasuringFunctionId + Aspect
		//deviceType, err, code := con.GetDeviceTypesFiltered("", []string{"urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0", "urn:infai:ses:measuring-function:0e19d094-70c6-402c-8523-3aaff2ce6dd9"}, []string{})
		deviceType, err, code := con.GetDeviceTypesFiltered([]model.DeviceTypesFilter{{FunctionId: "urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0", DeviceClassId: "", AspectId: "urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"}})

		deviceTypeStringified := `[{"id":"urn:infai:ses:device-type:a8cbd322-9d8c-4f4c-afec-ae4b7986b6ed","name":"Blebox-Air-Sensor","description":"","services":[{"id":"urn:infai:ses:service:422fd899-a2cc-4e43-8d81-4e330a7ca8ab","local_id":"","name":"getParticleAmountPM10Service","description":"","interaction":"","aspect_ids":["urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"],"protocol_id":"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b","inputs":null,"outputs":null,"function_ids":["urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0"],"rdf_type":"https://senergy.infai.org/ontology/Service"}],"device_class_id":"urn:infai:ses:device-class:8bd38ea2-1835-4a1e-ac02-6b3169513fd3","rdf_type":"https://senergy.infai.org/ontology/DeviceType"}]`

		if err != nil {
			t.Fatal(deviceType, err, code)
		} else {
			b, err := json.Marshal(deviceType)
			if err != nil {
				t.Fatal(deviceType, err, code)
			}
			if string(b) != deviceTypeStringified {
				t.Log("expected:", deviceTypeStringified)
				t.Log("was:", string(b))
				t.Fatal("error")
			}
		}
	}
}

func testReadDeviceType_1MF_variante2(con *controller.Controller) func(t *testing.T) {
	return func(t *testing.T) {
		// 1 MeasuringFunctionId + Aspect
		//deviceType, err, code := con.GetDeviceTypesFiltered("", []string{"urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0", "urn:infai:ses:measuring-function:0e19d094-70c6-402c-8523-3aaff2ce6dd9"}, []string{})
		deviceType, err, code := con.GetDeviceTypesFiltered([]model.DeviceTypesFilter{{FunctionId: "urn:infai:ses:measuring-function:f2769eb9-b6ad-4f7e-bd28-e4ea043d2f8b", DeviceClassId: "", AspectId: "urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"}})

		deviceTypeStringified := `[{"id":"urn:infai:ses:device-type:3cc09a10-1feb-4f8b-9390-8d08bf3ba22d","name":"Cyrus 4-in-1 Multisensor","description":"","services":[{"id":"urn:infai:ses:service:d3dba284-ef6d-4f12-81df-ed11506702b2","local_id":"","name":"getTemperatureService","description":"","interaction":"","aspect_ids":["urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"],"protocol_id":"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b","inputs":null,"outputs":null,"function_ids":["urn:infai:ses:measuring-function:f2769eb9-b6ad-4f7e-bd28-e4ea043d2f8b"],"rdf_type":"https://senergy.infai.org/ontology/Service"}],"device_class_id":"urn:infai:ses:device-class:ff64280a-58e6-4cf9-9a44-e70d3831a79d","rdf_type":"https://senergy.infai.org/ontology/DeviceType"},{"id":"urn:infai:ses:device-type:662d9c9f-949d-4577-9485-9cb7255f547f","name":"Danfoss Radiator Thermostat","description":"","services":[{"id":"urn:infai:ses:service:f306de41-a55b-45ed-afc9-039bbe53db1b","local_id":"","name":"getTemperatureService","description":"","interaction":"","aspect_ids":["urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"],"protocol_id":"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b","inputs":null,"outputs":null,"function_ids":["urn:infai:ses:measuring-function:f2769eb9-b6ad-4f7e-bd28-e4ea043d2f8b"],"rdf_type":"https://senergy.infai.org/ontology/Service"}],"device_class_id":"urn:infai:ses:device-class:997937d6-c5f3-4486-b67c-114675038393","rdf_type":"https://senergy.infai.org/ontology/DeviceType"}]`

		if err != nil {
			t.Fatal(deviceType, err, code)
		} else {
			b, err := json.Marshal(deviceType)
			if err != nil {
				t.Fatal(deviceType, err, code)
			}
			if string(b) != deviceTypeStringified {
				t.Log("expected:", deviceTypeStringified)
				t.Log("was:", string(b))
				t.Fatal("error")
			}
		}
	}
}

func testReadDeviceType_2MF_sameAspect(con *controller.Controller) func(t *testing.T) {
	return func(t *testing.T) {
		// 2 MeasuringFunctionId + same Aspect
		//deviceType, err, code := con.GetDeviceTypesFiltered("", []string{"urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0", "urn:infai:ses:measuring-function:0e19d094-70c6-402c-8523-3aaff2ce6dd9"}, []string{})
		deviceType, err, code := con.GetDeviceTypesFiltered([]model.DeviceTypesFilter{{FunctionId: "urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0", DeviceClassId: "", AspectId: "urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"},
			{FunctionId: "urn:infai:ses:measuring-function:0e19d094-70c6-402c-8523-3aaff2ce6dd9", DeviceClassId: "", AspectId: "urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"}})

		deviceTypeStringified := `[{"id":"urn:infai:ses:device-type:a8cbd322-9d8c-4f4c-afec-ae4b7986b6ed","name":"Blebox-Air-Sensor","description":"","services":[{"id":"urn:infai:ses:service:422fd899-a2cc-4e43-8d81-4e330a7ca8ab","local_id":"","name":"getParticleAmountPM10Service","description":"","interaction":"","aspect_ids":["urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"],"protocol_id":"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b","inputs":null,"outputs":null,"function_ids":["urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0"],"rdf_type":"https://senergy.infai.org/ontology/Service"},{"id":"urn:infai:ses:service:1d20a68b-7136-456c-ace5-c3adb66866bf","local_id":"","name":"getParticleAmountPM1Service","description":"","interaction":"","aspect_ids":["urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"],"protocol_id":"urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b","inputs":null,"outputs":null,"function_ids":["urn:infai:ses:measuring-function:0e19d094-70c6-402c-8523-3aaff2ce6dd9"],"rdf_type":"https://senergy.infai.org/ontology/Service"}],"device_class_id":"urn:infai:ses:device-class:8bd38ea2-1835-4a1e-ac02-6b3169513fd3","rdf_type":"https://senergy.infai.org/ontology/DeviceType"}]`

		if err != nil {
			t.Fatal(deviceType, err, code)
		} else {
			b, err := json.Marshal(deviceType)
			if err != nil {
				t.Fatal(deviceType, err, code)
				return
			}
			if string(b) != deviceTypeStringified {
				t.Log("expected:", deviceTypeStringified)
				t.Log("was:", string(b))
				t.Fatal("error")
			}
		}
	}
}

func testReadDeviceType_2MF_sameAspect_DifferentDeviceClasses(con *controller.Controller) func(t *testing.T) {
	return func(t *testing.T) {
		// 2 MeasuringFunctionId + same Aspect + 2 different DeviceClasses
		//deviceType, err, code := con.GetDeviceTypesFiltered("", []string{"urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0", "urn:infai:ses:measuring-function:0e19d094-70c6-402c-8523-3aaff2ce6dd9"}, []string{})
		deviceType, err, code := con.GetDeviceTypesFiltered([]model.DeviceTypesFilter{{FunctionId: "urn:infai:ses:measuring-function:f2769eb9-b6ad-4f7e-bd28-e4ea043d2f8b", DeviceClassId: "", AspectId: "urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"},
			{FunctionId: "urn:infai:ses:measuring-function:0e19d094-70c6-402c-8523-3aaff2ce6dd9", DeviceClassId: "", AspectId: "urn:infai:ses:aspect:a14c5efb-b0b6-46c3-982e-9fded75b5ab6"}})

		if err != nil {
			t.Fatal(deviceType, err, code)
		} else {
			b, err := json.Marshal(deviceType)
			if err != nil {
				t.Fatal(deviceType, err, code)
			}
			if string(b) != "null" {
				t.Log("expected: null")
				t.Log("was:", string(b))
				t.Fatal("error")
			}
		}
	}
}

func testReadDeviceTypeWithId1(con *controller.Controller) func(t *testing.T) {
	return func(t *testing.T) {
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

		// DeviceClass
		if deviceType.DeviceClassId != "urn:infai:ses:device-class:14e56881-16f9-4120-bb41-270a43070c86" {
			t.Fatal("error deviceclass id")
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
		if deviceType.Services[0].Interaction != "" {
			t.Fatal("error service -> 0 -> Interaction")
		}
		if deviceType.Services[0].ProtocolId != "urn:infai:ses:protocol:f3a63aeb-187e-4dd9-9ef5-d97a6eb6292b" {
			t.Fatal("error service -> 0 -> ProtocolId")
		}
		if deviceType.Services[0].LocalId != "" { // not stored as TRIPLE
			t.Fatal("error service -> 0 -> LocalId")
		}
		if deviceType.Services[0].AspectIds[0] != "urn:infai:ses:aspect:a7470d73-dde3-41fc-92bd-f16bb28f2da6" {
			t.Fatal("error aspect -> 0 -> AspectIds -> 0")
		}
		if deviceType.Services[0].FunctionIds[0] != "urn:infai:ses:controlling-function:c54e2a89-1fb8-4ecb-8993-a7b40b355599" {
			t.Fatal("error aspect -> 0 -> FunctionIds -> 0")
		}
		if err != nil {
			t.Fatal(deviceType, err, code)
		} else {
			t.Log(deviceType)
		}
	}
}

func testCreateAndDeleteDeviceTypePart1(producer *producer.Producer) func(t *testing.T) {
	return func(t *testing.T) {
		devicetype := model.DeviceType{}
		devicetype.Id = "urn:infai:ses:devicetype:1"
		devicetype.Name = "Philips Hue Color"
		devicetype.DeviceClassId = "urn:infai:ses:deviceclass:2"
		devicetype.Description = "description"
		devicetype.Services = []model.Service{}
		devicetype.Services = append(devicetype.Services, model.Service{
			"urn:infai:ses:service:3a",
			"localId",
			"setBrightness",
			"",
			"",
			[]string{"urn:infai:ses:aspect:4a"},
			"urn:infai:ses:protocol:1425",
			[]model.Content{},
			[]model.Content{},
			[]string{"urn:infai:ses:function:5a"},
			"asdasdsadsadasd",
		})

		devicetype.Services = append(devicetype.Services, model.Service{
			"urn:infai:ses:service:3b",
			"localId",
			"setBrightness",
			"",
			"",
			[]string{"urn:infai:ses:aspect:4b"},
			"urn:infai:ses:protocol:1425",
			[]model.Content{},
			[]model.Content{},
			[]string{"urn:infai:ses:function:5b"},
			"asdasdsadsadasd",
		})

		producer.PublishDeviceType(devicetype, "sdfdsfsf")
	}
}

func testCreateAndDeleteDeviceTypePart2(producer *producer.Producer) func(t *testing.T) {
	return func(t *testing.T) {
		err := producer.PublishDeviceTypeDelete("urn:infai:ses:devicetype:1", "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}
	}
}
