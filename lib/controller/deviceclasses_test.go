package controller

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"net/http"
	"testing"
)

func TestValidDeviceClass (t *testing.T) {
	deviceclass := model.DeviceClass{}
	deviceclass.Id = "urn:infai:ses:deviceclass:1234"
	deviceclass.Name = "Lamp"
	deviceclass.Type = model.SES_ONTOLOGY_DEVICE_CLASS

	controller := Controller{}
	err, code := controller.ValidateDeviceClass(deviceclass)
	if err == nil && code == http.StatusOK {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestDeviceClassMissingId (t *testing.T) {
	deviceclass := model.DeviceClass{}
	deviceclass.Id = ""

	controller := Controller{}
	err, code := controller.ValidateDeviceClass(deviceclass)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestDeviceClassMissingName (t *testing.T) {
	deviceclass := model.DeviceClass{}
	deviceclass.Id = "urn:infai:ses:deviceclass:1234"
	deviceclass.Name = ""

	controller := Controller{}
	err, code := controller.ValidateDeviceClass(deviceclass)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestDeviceClassWrongType (t *testing.T) {
	deviceclass := model.DeviceClass{}
	deviceclass.Id = "urn:infai:ses:deviceclass:1234"
	deviceclass.Name = "Lamp"
	deviceclass.Type = "wrongType"

	controller := Controller{}
	err, code := controller.ValidateDeviceClass(deviceclass)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

