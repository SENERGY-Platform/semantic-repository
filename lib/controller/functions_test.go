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

func TestValidMeasuringFunction (t *testing.T) {
	functions := []model.Function{}
	functions = append(functions, model.Function{Id: "urn:infai:ses:function:122", Name: "Air", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION, ConceptId: "urn:infai:ses:concept:1"})

	controller := Controller{}
	err, code := controller.ValidateFunctions(functions)
	if err == nil && code == http.StatusOK {
		t.Log(functions)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidControllingFunction (t *testing.T) {
	functions := []model.Function{}
	functions = append(functions, model.Function{Id: "urn:infai:ses:function:122", Name: "Air", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION, ConceptId: "urn:infai:ses:concept:1"})

	controller := Controller{}
	err, code := controller.ValidateFunctions(functions)
	if err == nil && code == http.StatusOK {
		t.Log(functions)
	} else {
		t.Fatal(err, code)
	}
}

func TestFunctionNoData (t *testing.T) {
	functions := []model.Function{}

	controller := Controller{}
	err, code := controller.ValidateFunctions(functions)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestFunctionMissingId (t *testing.T) {
	functions := []model.Function{}
	functions = append(functions, model.Function{Id: ""})

	controller := Controller{}
	err, code := controller.ValidateFunctions(functions)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestFunctionMissingName (t *testing.T) {
	functions := []model.Function{}
	functions = append(functions, model.Function{Id: "urn:infai:ses:function:122", Name: ""})

	controller := Controller{}
	err, code := controller.ValidateFunctions(functions)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestFunctionWrongType (t *testing.T) {
	functions := []model.Function{}
	functions = append(functions, model.Function{Id: "urn:infai:ses:function:122", Name: "Air", RdfType: "wrongType"})

	controller := Controller{}
	err, code := controller.ValidateFunctions(functions)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestFunctionMissingConceptId (t *testing.T) {
	functions := []model.Function{}
	functions = append(functions, model.Function{Id: "urn:infai:ses:function:122", Name: "Air", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION, ConceptId: ""})

	controller := Controller{}
	err, code := controller.ValidateFunctions(functions)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}