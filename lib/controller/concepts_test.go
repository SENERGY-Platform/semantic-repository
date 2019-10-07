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

func TestConceptsSetRdfTypes(t *testing.T) {

	concept := model.Concept{}

	SetConceptRdfTypes(&concept)
	if concept.RdfType != model.SES_ONTOLOGY_CONCEPT {
		t.Fatal()
	}
	t.Log(concept)
}

func TestValidationConceptMissingId(t *testing.T) {
	concept := model.Concept{}
	concept.Id = ""

	controller := Controller{}
	err, code := controller.ValidateConcept(concept)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationConceptMissingName(t *testing.T) {
	concept := model.Concept{}
	concept.Id = "1"
	concept.Name = ""

	controller := Controller{}
	err, code := controller.ValidateConcept(concept)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationWrongConceptRdfType(t *testing.T) {
	concept := model.Concept{}
	concept.Id = "1"
	concept.Name = "name"
	concept.RdfType = "xxxx"

	controller := Controller{}
	err, code := controller.ValidateConcept(concept)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationValidConceptWithoutCharacteristicId(t *testing.T) {
	concept := model.Concept{}
	concept.Id = "1"
	concept.Name = "name"
	concept.RdfType = model.SES_ONTOLOGY_CONCEPT

	controller := Controller{}
	err, code := controller.ValidateConcept(concept)
	if err == nil && code == http.StatusOK {
		t.Log(concept)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationValidConceptWith1CharacteristicId(t *testing.T) {
	concept := model.Concept{}
	concept.Id = "1"
	concept.Name = "name"
	concept.RdfType = model.SES_ONTOLOGY_CONCEPT
	concept.CharacteristicIds = []string{"urn:infai:ses:characteristic:4444"}

	controller := Controller{}
	err, code := controller.ValidateConcept(concept)
	if err == nil && code == http.StatusOK {
		t.Log(concept)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationValidConceptWith2CharacteristicIds(t *testing.T) {
	concept := model.Concept{}
	concept.Id = "1"
	concept.Name = "name"
	concept.RdfType = model.SES_ONTOLOGY_CONCEPT
	concept.CharacteristicIds = []string{"urn:infai:ses:characteristic:4444","urn:infai:ses:characteristic:5555"}

	controller := Controller{}
	err, code := controller.ValidateConcept(concept)
	if err == nil && code == http.StatusOK {
		t.Log(concept)
	} else {
		t.Fatal(err, code)
	}
}
