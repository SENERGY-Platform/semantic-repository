package controller

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"net/http"
	"testing"
)

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
	functions = append(functions, model.Function{Id: "urn:infai:ses:function:122", Name: "Air", Type: "wrongType"})

	controller := Controller{}
	err, code := controller.ValidateFunctions(functions)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}
