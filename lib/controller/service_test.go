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

func TestValidService (t *testing.T) {
	service := []model.Service{}
	service = append(service, model.Service{
		Id: "urn:infai:ses:service:5555",
		LocalId: "4711", Name: "get",
		ProtocolId: "1111",
		RdfType: model.SES_ONTOLOGY_SERVICE,
		Aspects: []model.Aspect{{Id: "urn:infai:ses:aspect:1", Name: "aspect", RdfType: model.SES_ONTOLOGY_ASPECT}},
		Functions: []model.Function{{Id: "urn:infai:ses:function:1", Name: "function", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION, ConceptId: "urn:infai:ses:concept:1"}},
	})

	controller := Controller{}
	err, code := controller.ValidateService(service)
	if err == nil && code == http.StatusOK {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestServiceNoData (t *testing.T) {
	service := []model.Service{}

	controller := Controller{}
	err, code := controller.ValidateService(service)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestServiceMissingId (t *testing.T) {
	service := []model.Service{}
	service = append(service, model.Service{Id: ""})

	controller := Controller{}
	err, code := controller.ValidateService(service)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestServiceMissingLocalId (t *testing.T) {
	service := []model.Service{}
	service = append(service, model.Service{Id: "urn:infai:ses:service:5555", LocalId: ""})

	controller := Controller{}
	err, code := controller.ValidateService(service)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestServiceMissingName (t *testing.T) {
	service := []model.Service{}
	service = append(service, model.Service{Id: "urn:infai:ses:service:5555", LocalId: "4711", Name: ""})

	controller := Controller{}
	err, code := controller.ValidateService(service)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestServiceMissingProtocolId (t *testing.T) {
	service := []model.Service{}
	service = append(service, model.Service{Id: "urn:infai:ses:service:5555", LocalId: "4711", Name: "get", ProtocolId: ""})

	controller := Controller{}
	err, code := controller.ValidateService(service)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestServiceWrongType (t *testing.T) {
	service := []model.Service{}
	service = append(service, model.Service{Id: "urn:infai:ses:service:5555", LocalId: "4711", Name: "get", ProtocolId: "1111", RdfType: "wrongType"})

	controller := Controller{}
	err, code := controller.ValidateService(service)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}