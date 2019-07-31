package controller

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"net/http"
	"testing"
)

func TestValidMeasuringFunction (t *testing.T) {
	functions := []model.Function{}
	functions = append(functions, model.Function{Id: "urn:infai:ses:function:122", Name: "Air", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION, ConceptIds: []string{"urn:infai:ses:concept:1"}})

	controller := Controller{}
	err, code := controller.ValidateFunctions(functions)
	if err == nil && code == http.StatusOK {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidControllingFunction (t *testing.T) {
	functions := []model.Function{}
	functions = append(functions, model.Function{Id: "urn:infai:ses:function:122", Name: "Air", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION, ConceptIds: []string{"urn:infai:ses:concept:1"}})

	controller := Controller{}
	err, code := controller.ValidateFunctions(functions)
	if err == nil && code == http.StatusOK {
		t.Log(err)
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

func TestFunctionMissingConcept (t *testing.T) {
	functions := []model.Function{}
	functions = append(functions, model.Function{Id: "urn:infai:ses:function:122", Name: "Air", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION, ConceptIds: []string{}})

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
	functions = append(functions, model.Function{Id: "urn:infai:ses:function:122", Name: "Air", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION, ConceptIds: []string{""}})

	controller := Controller{}
	err, code := controller.ValidateFunctions(functions)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}