package controller

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"net/http"
	"testing"
)

func TestCharacteristicsSetRdfTypes(t *testing.T) {
	characteristic := model.Characteristic{}
	characteristic.RdfType = "xxxx"
	characteristic.SubCharacteristics = []model.Characteristic{{
		Id:       "",
		RdfType:  "",
		MinValue: 0,
		MaxValue: 0,
		Value:    "",
		Type:     "",
		Name:     "",
		SubCharacteristics:
		[]model.Characteristic{{
			Id:       "",
			RdfType:  "",
			MinValue: 0,
			MaxValue: 0,
			Value:    "",
			Type:     "",
			Name:     "",
			SubCharacteristics: []model.Characteristic{{
				Id:                 "",
				RdfType:            "",
				MinValue:           0,
				MaxValue:           0,
				Value:              "",
				Type:               "",
				Name:               "",
				SubCharacteristics: nil,
			}},
		}},
	}}

	SetCharacteristicRdfTypes(&characteristic)
	if characteristic.RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
		t.Fatal("error rdf type")
	}
	if characteristic.SubCharacteristics[0].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
		t.Fatal("error rdf type")
	}
	if characteristic.SubCharacteristics[0].SubCharacteristics[0].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
		t.Fatal("error rdf type")
	}
	if characteristic.SubCharacteristics[0].SubCharacteristics[0].SubCharacteristics[0].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
		t.Fatal("error rdf type")
	}

	t.Log(characteristic)
}

func TestValidationCharacteristicsMissingId(t *testing.T) {
	characteristic := model.Characteristic{}
	characteristic.Id = ""

	controller := Controller{}
	err, code := controller.ValidateCharacteristics(characteristic)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationCharacteristicsMissingName(t *testing.T) {
	characteristic := model.Characteristic{}
	characteristic.Id = "urn:infai:ses:characteristics:1"
	characteristic.Name = ""

	controller := Controller{}
	err, code := controller.ValidateCharacteristics(characteristic)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationCharacteristicsMissingRdfType(t *testing.T) {
	characteristic := model.Characteristic{}
	characteristic.Id = "urn:infai:ses:characteristics:1"
	characteristic.Name = "char"
	characteristic.RdfType = ""

	controller := Controller{}
	err, code := controller.ValidateCharacteristics(characteristic)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationCharacteristicsMissingType(t *testing.T) {
	characteristic := model.Characteristic{}
	characteristic.Id = "urn:infai:ses:characteristics:1"
	characteristic.Name = "char"
	characteristic.RdfType = model.SES_ONTOLOGY_CHARACTERISTIC
	characteristic.Type = ""

	controller := Controller{}
	err, code := controller.ValidateCharacteristics(characteristic)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationCharacteristicsMissingSubId(t *testing.T) {
	characteristic := model.Characteristic{}
	characteristic.Id = "urn:infai:ses:characteristics:1"
	characteristic.Name = "char"
	characteristic.RdfType = model.SES_ONTOLOGY_CHARACTERISTIC
	characteristic.Type = model.Structure
	characteristic.SubCharacteristics = []model.Characteristic{{
		Id:                 "",
		RdfType:            "",
		MinValue:           0,
		MaxValue:           0,
		Value:              "",
		Type:               "",
		Name:               "",
		SubCharacteristics: nil,
	}}

	controller := Controller{}
	err, code := controller.ValidateCharacteristics(characteristic)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationCharacteristicsMissingSubName(t *testing.T) {
	characteristic := model.Characteristic{}
	characteristic.Id = "urn:infai:ses:characteristics:1"
	characteristic.Name = "char"
	characteristic.RdfType = model.SES_ONTOLOGY_CHARACTERISTIC
	characteristic.Type = model.Structure
	characteristic.SubCharacteristics = []model.Characteristic{{
		Id:                 "urn:infai:ses:characteristics:2",
		RdfType:            "",
		MinValue:           0,
		MaxValue:           0,
		Value:              "",
		Type:               "",
		Name:               "",
		SubCharacteristics: nil,
	}}

	controller := Controller{}
	err, code := controller.ValidateCharacteristics(characteristic)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationCharacteristicsMissingSubRdfType(t *testing.T) {
	characteristic := model.Characteristic{}
	characteristic.Id = "urn:infai:ses:characteristics:1"
	characteristic.Name = "char1"
	characteristic.RdfType = model.SES_ONTOLOGY_CHARACTERISTIC
	characteristic.Type = model.Structure
	characteristic.SubCharacteristics = []model.Characteristic{{
		Id:                 "urn:infai:ses:characteristics:2",
		RdfType:            "",
		MinValue:           0,
		MaxValue:           0,
		Value:              "",
		Type:               "",
		Name:               "char2",
		SubCharacteristics: nil,
	}}

	controller := Controller{}
	err, code := controller.ValidateCharacteristics(characteristic)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationCharacteristicsMissingSubType(t *testing.T) {
	characteristic := model.Characteristic{}
	characteristic.Id = "urn:infai:ses:characteristics:1"
	characteristic.Name = "char1"
	characteristic.RdfType = model.SES_ONTOLOGY_CHARACTERISTIC
	characteristic.Type = model.Structure
	characteristic.SubCharacteristics = []model.Characteristic{{
		Id:                 "urn:infai:ses:characteristics:2",
		RdfType:            model.SES_ONTOLOGY_CHARACTERISTIC,
		MinValue:           0,
		MaxValue:           0,
		Value:              "",
		Type:               "",
		Name:               "char2",
		SubCharacteristics: nil,
	}}

	controller := Controller{}
	err, code := controller.ValidateCharacteristics(characteristic)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationCharacteristicsMissingSubSubId(t *testing.T) {
	characteristic := model.Characteristic{}
	characteristic.Id = "urn:infai:ses:characteristics:1"
	characteristic.Name = "char1"
	characteristic.RdfType = model.SES_ONTOLOGY_CHARACTERISTIC
	characteristic.Type = model.Structure
	characteristic.SubCharacteristics = []model.Characteristic{{
		Id:       "urn:infai:ses:characteristics:2",
		RdfType:  model.SES_ONTOLOGY_CHARACTERISTIC,
		MinValue: 0,
		MaxValue: 0,
		Value:    "",
		Type:     model.Structure,
		Name:     "char2",
		SubCharacteristics:
		[]model.Characteristic{{
			Id:                 "",
			RdfType:            "",
			MinValue:           0,
			MaxValue:           0,
			Value:              "",
			Type:               "",
			Name:               "",
			SubCharacteristics: nil,
		}},
	}}

	controller := Controller{}
	err, code := controller.ValidateCharacteristics(characteristic)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}
