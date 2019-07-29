package controller

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"net/http"
	"testing"
)

func TestValidDeviceType(t *testing.T) {

	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:1111"
	devicetype.Name = "Philips Hue Color"
	devicetype.DeviceClass = model.DeviceClass{
		Id: "urn:infai:ses:deviceclass:2222",
		Name: "Lamp",
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3333",
		"localId",
		"setBrigthness",
		"",
		[]model.Aspect{{Id:"urn:infai:ses:aspect:4444", Name: "Lighting", Type: "asasasdsadas"}},
		"",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id:"urn:infai:ses:function:5555", Name: "brightnessAdjustment", ConceptIds: []model.ConceptId{{"urn:infai:ses:concept:6666"},{"urn:infai:ses:concept:7777"}}, Type: model.SES_ONTOLOGY_CONTROLLING_FUNCTION }},
		"asdasdsadsadasd",
	})

	controller := Controller{}
	err, code := controller.ValidateDeviceType(devicetype)
	if err == nil && code == http.StatusOK {
		t.Log("ok")
	} else {
		t.Fatal("error")
	}
}

func TestMissingId(t *testing.T) {
	devicetype := model.DeviceType{}
	devicetype.Id = ""

	controller := Controller{}
	err, code := controller.ValidateDeviceType(devicetype)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestMissingName(t *testing.T) {
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:5555"
	devicetype.Name = ""

	controller := Controller{}
	err, code := controller.ValidateDeviceType(devicetype)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestWrongType(t *testing.T) {
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:5555"
	devicetype.Name = "philips hue color"
	devicetype.Type = "type"

	controller := Controller{}
	err, code := controller.ValidateDeviceType(devicetype)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestNoServiceData(t *testing.T) {
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:5555"
	devicetype.Name = "philips hue color"
	devicetype.Type = model.SES_ONTOLOGY_DEVICE_TYPE

	controller := Controller{}
	err, code := controller.ValidateDeviceType(devicetype)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}