package controller

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"net/http"
	"testing"
)

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
	service = append(service, model.Service{Id: "urn:infai:ses:service:5555", LocalId: "4711", Name: "get", ProtocolId: "1111", Type: "wrongType"})

	controller := Controller{}
	err, code := controller.ValidateService(service)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}