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
	"context"
	"encoding/json"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/database"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/SENERGY-Platform/semantic-repository/lib/testutil"
	"net/http"
	"reflect"
	"sync"
	"testing"
)

func TestValidDeviceType(t *testing.T) {

	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:1111"
	devicetype.Name = "Philips Hue Color"
	devicetype.DeviceClassId = "urn:infai:ses:deviceclass:2222"
	devicetype.Description = "description"
	devicetype.RdfType = model.SES_ONTOLOGY_DEVICE_TYPE
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3333",
		"localId",
		"setBrigthness",
		"",
		"",
		[]string{"urn:infai:ses:aspect:4444"},
		"urn:infai:ses:protocolId",
		[]model.Content{},
		[]model.Content{},
		[]string{"urn:infai:ses:function:5555"},
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

func TestValidationDeviceTypeInvalidId(t *testing.T) {
	devicetype := model.DeviceType{}
	devicetype.Id = "foo"

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

func TestValidationDeviceTypeNoDeviceClassId(t *testing.T) {
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

func TestValidationDeviceTypeNoServiceData(t *testing.T) {
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:5555"
	devicetype.Name = "philips hue color"
	devicetype.RdfType = model.SES_ONTOLOGY_DEVICE_TYPE
	devicetype.DeviceClassId = "urn:infai:ses:deviceclass:555556498"

	controller := Controller{}
	err, code := controller.ValidateDeviceType(devicetype)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestServiceWithInteraction(t *testing.T) {
	dt := model.DeviceType{
		Id:            model.URN_PREFIX + "dtid",
		Name:          "dtname",
		DeviceClassId: model.URN_PREFIX + "dcid",
		Services: []model.Service{
			{
				Id:          model.URN_PREFIX + "sid",
				LocalId:     "lsid",
				Name:        "sname",
				Interaction: model.EVENT,
				ProtocolId:  model.URN_PREFIX + "pid",
				AspectIds:   []string{model.URN_PREFIX + "aid"},
				FunctionIds: []string{model.URN_PREFIX + "fid"},
			},
		},
	}
	SetDevicetypeRdfTypes(&dt)

	err, _ := (&Controller{}).ValidateDeviceType(dt)
	if err != nil {
		t.Error(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	defer wg.Wait()
	defer cancel()

	conf := config.Config{}
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

	ctrl, err := New(conf, db)
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("create device-type", testCreateDeviceType(ctrl, dt))

	//cleanup unsaved properties
	dt.Services[0].LocalId = ""

	t.Run("check device-type", testCheckDeviceType(ctrl, dt.Id, dt))
}

func TestProtocolIdChange(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	defer wg.Wait()
	defer cancel()
	conf := config.Config{}
	err := testutil.GetDockerEnv(ctx, &wg, &conf)
	if err != nil {
		t.Error(err)
		return
	}
	db, err := database.New(conf)
	if err != nil {
		t.Error(err)
		return
	}

	ctrl, err := New(conf, db)
	if err != nil {
		t.Error(err)
		return
	}

	dt := model.DeviceType{
		Id:            model.URN_PREFIX + "dtid",
		Name:          "dtname",
		DeviceClassId: model.URN_PREFIX + "dcid",
		Services: []model.Service{
			{
				Id:          model.URN_PREFIX + "sid",
				LocalId:     "lsid",
				Name:        "sname",
				Interaction: model.EVENT,
				ProtocolId:  model.URN_PREFIX + "pid",
				AspectIds:   []string{model.URN_PREFIX + "aid"},
				FunctionIds: []string{model.URN_PREFIX + "fid"},
			},
		},
	}
	SetDevicetypeRdfTypes(&dt)

	//tstr := triples.(string)
	//t.Log(tstr)
	tstr := `<urn:infai:ses:dtid> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <https://senergy.infai.org/ontology/DeviceType> .
        <urn:infai:ses:dtid> <http://www.w3.org/2000/01/rdf-schema#label> "dtname" .
        <urn:infai:ses:dtid> <https://senergy.infai.org/ontology/hasDeviceClass> <urn:infai:ses:dcid> .
        <urn:infai:ses:dtid> <https://senergy.infai.org/ontology/hasService> <urn:infai:ses:sid> .
        <urn:infai:ses:sid> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <https://senergy.infai.org/ontology/Service> .
        <urn:infai:ses:sid> <http://www.w3.org/2000/01/rdf-schema#label> "sname" .
        <urn:infai:ses:sid> <https://senergy.infai.org/ontology/exposesFunction> <urn:infai:ses:fid> .
        <urn:infai:ses:sid> <https://senergy.infai.org/ontology/hasProtocol> "urn:infai:ses:pid" .
        <urn:infai:ses:sid> <https://senergy.infai.org/ontology/interaction> "event" .
        <urn:infai:ses:sid> <https://senergy.infai.org/ontology/refersTo> <urn:infai:ses:aid> .`

	err = db.InsertData(tstr)

	//cleanup unsaved properties
	dt.Services[0].LocalId = ""

	//protocol id as id fix removes protocol ids
	protocolId := dt.Services[0].ProtocolId
	dt.Services[0].ProtocolId = ""

	t.Run("check device-type without protocol", testCheckDeviceType(ctrl, dt.Id, dt))

	dt.Services[0].ProtocolId = protocolId
	t.Run("update device-type with new protocol", testCreateDeviceType(ctrl, dt))

	t.Run("check device-type with protocol", testCheckDeviceType(ctrl, dt.Id, dt))
}

func testCheckDeviceType(ctrl *Controller, id string, expected model.DeviceType) func(t *testing.T) {
	return func(t *testing.T) {
		actualDt, err, _ := ctrl.GetDeviceType(id)
		if err != nil {
			t.Error(err)
			return
		}
		if !reflect.DeepEqual(actualDt, expected) {
			actualJson, _ := json.Marshal(actualDt)
			expectedJson, _ := json.Marshal(expected)
			t.Error(string(actualJson), string(expectedJson))
			return
		}
	}
}

func testCreateDeviceType(ctrl *Controller, deviceType model.DeviceType) func(t *testing.T) {
	return func(t *testing.T) {
		err := ctrl.SetDeviceType(deviceType, "owner")
		if err != nil {
			t.Error(err)
			return
		}
	}
}
