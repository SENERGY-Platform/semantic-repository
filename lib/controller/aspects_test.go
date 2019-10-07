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

func TestValidAspect (t *testing.T) {
	aspects := []model.Aspect{}
	aspects = append(aspects, model.Aspect{Id: "urn:infai:ses:aspect:122", Name: "Air", RdfType: model.SES_ONTOLOGY_ASPECT})

	controller := Controller{}
	err, code := controller.ValidateAspects(aspects)
	if err == nil && code == http.StatusOK {
		t.Log(aspects)
	} else {
		t.Fatal(err, code)
	}
}

func TestAspectNoData (t *testing.T) {
	aspects := []model.Aspect{}

	controller := Controller{}
	err, code := controller.ValidateAspects(aspects)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestAspectMissingId (t *testing.T) {
	aspects := []model.Aspect{}
	aspects = append(aspects, model.Aspect{Id: ""})

	controller := Controller{}
	err, code := controller.ValidateAspects(aspects)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestAspectMissingName (t *testing.T) {
	aspects := []model.Aspect{}
	aspects = append(aspects, model.Aspect{Id: "urn:infai:ses:aspect:122", Name: ""})

	controller := Controller{}
	err, code := controller.ValidateAspects(aspects)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestAspectWrongType (t *testing.T) {
	aspects := []model.Aspect{}
	aspects = append(aspects, model.Aspect{Id: "urn:infai:ses:aspect:122", Name: "Air", RdfType: "wrongType"})

	controller := Controller{}
	err, code := controller.ValidateAspects(aspects)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}
