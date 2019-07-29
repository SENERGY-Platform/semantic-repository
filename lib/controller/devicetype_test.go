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
		Id:   "urn:infai:ses:deviceclass:2222",
		Name: "Lamp",
		Type: model.SES_ONTOLOGY_DEVICE_CLASS,
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.Type = model.SES_ONTOLOGY_DEVICE_TYPE
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3333",
		"localId",
		"setBrigthness",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4444", Name: "Lighting", Type: model.SES_ONTOLOGY_ASPECT}},
		"protocolId",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5555", Name: "brightnessAdjustment", ConceptIds: []model.ConceptId{{"urn:infai:ses:concept:6666"}, {"urn:infai:ses:concept:7777"}}, Type: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		model.SES_ONTOLOGY_SERVICE,
	})

	controller := Controller{}
	err, code := controller.ValidateDeviceType(devicetype)
	if err == nil && code == http.StatusOK {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestMissingId(t *testing.T) {
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

func TestMissingName(t *testing.T) {
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

func TestWrongType(t *testing.T) {
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:5555"
	devicetype.Name = "philips hue color"
	devicetype.Type = "type"

	controller := Controller{}
	err, code := controller.ValidateDeviceType(devicetype)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
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
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestNoServiceNoDeviceClass(t *testing.T) {
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:5555"
	devicetype.Name = "philips hue color"
	devicetype.Type = model.SES_ONTOLOGY_DEVICE_TYPE
	devicetype.Services = []model.Service{{Id: "urn:infai:ses:service:1",
		Type: model.SES_ONTOLOGY_SERVICE,
		Name: "test", LocalId: "2",
		ProtocolId: "3",
		Aspects:    []model.Aspect{{Id: "urn:infai:ses:aspect:1", Name: "aspect", Type: model.SES_ONTOLOGY_ASPECT}},
		Functions:  []model.Function{{Id: "urn:infai:ses:function:1", Name: "function", Type: model.SES_ONTOLOGY_MEASURING_FUNCTION, ConceptIds: []model.ConceptId{{Id: "urn:infai:ses:concept:1"}}}}}}
	devicetype.DeviceClass = model.DeviceClass{}

	controller := Controller{}
	err, code := controller.ValidateDeviceType(devicetype)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}
