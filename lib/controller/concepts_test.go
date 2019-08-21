package controller

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"net/http"
	"testing"
)

func TestConceptsSetRdfTypes(t *testing.T) {

	concept := model.Concept{}
	concept.Characteristics = append(concept.Characteristics, model.Characteristic{
		Id:       "",
		Name:     "",
		RdfType:  "",
		Type:     "",
		MaxValue: 0,
		MinValue: 0,
		Value:    0,
		SubCharacteristics: append(concept.Characteristics, model.Characteristic{
			Id:       "",
			Name:     "",
			RdfType:  "",
			Type:     "",
			MaxValue: 0,
			MinValue: 0,
			Value:    0,
			SubCharacteristics: append(concept.Characteristics, model.Characteristic{
				Id:                 "",
				Name:               "",
				RdfType:            "",
				Type:               "",
				MaxValue:           0,
				MinValue:           0,
				Value:              0,
				SubCharacteristics: nil,
			}),
		}),
	})

	SetConceptRdfTypes(&concept)
	if concept.RdfType != model.SES_ONTOLOGY_CONCEPT {
		t.Fatal()
	}
	if concept.Characteristics[0].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
		t.Fatal()
	}
	if concept.Characteristics[0].SubCharacteristics[0].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
		t.Fatal()
	}
	if concept.Characteristics[0].SubCharacteristics[0].SubCharacteristics[0].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
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

func TestValidationValidConceptWithoutCharacteristic(t *testing.T) {
	concept := model.Concept{}
	concept.Id = "1"
	concept.Name = "name"
	concept.RdfType = model.SES_ONTOLOGY_CONCEPT
	concept.Characteristics = []model.Characteristic{}

	controller := Controller{}
	err, code := controller.ValidateConcept(concept)
	if err == nil && code == http.StatusOK {
		t.Log(concept)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationMissingCharacteristicId(t *testing.T) {
	concept := model.Concept{}
	concept.Id = "1"
	concept.Name = "name"
	concept.RdfType = model.SES_ONTOLOGY_CONCEPT
	concept.Characteristics = []model.Characteristic{{
		Id: "",
	}}

	controller := Controller{}
	err, code := controller.ValidateConcept(concept)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationMissingCharacteristicName(t *testing.T) {
	concept := model.Concept{}
	concept.Id = "1"
	concept.Name = "name"
	concept.RdfType = model.SES_ONTOLOGY_CONCEPT
	concept.Characteristics = []model.Characteristic{{
		Id: "2",
		Name: "",
	}}

	controller := Controller{}
	err, code := controller.ValidateConcept(concept)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationMissingCharacteristicRdfType(t *testing.T) {
	concept := model.Concept{}
	concept.Id = "1"
	concept.Name = "name"
	concept.RdfType = model.SES_ONTOLOGY_CONCEPT
	concept.Characteristics = []model.Characteristic{{
		Id: "2",
		Name: "nameChar",
		RdfType: "xxxxas",
	}}

	controller := Controller{}
	err, code := controller.ValidateConcept(concept)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationMissingCharacteristicType(t *testing.T) {
	concept := model.Concept{}
	concept.Id = "1"
	concept.Name = "name"
	concept.RdfType = model.SES_ONTOLOGY_CONCEPT
	concept.Characteristics = []model.Characteristic{{
		Id: "2",
		Name: "nameChar",
		RdfType: model.SES_ONTOLOGY_CONCEPT,
		Type: "xxxxas",
	}}

	controller := Controller{}
	err, code := controller.ValidateConcept(concept)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err, code)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationValidWith1Characteristic(t *testing.T) {
	concept := model.Concept{}
	concept.Id = "1"
	concept.Name = "name"
	concept.RdfType = model.SES_ONTOLOGY_CONCEPT
	concept.Characteristics = []model.Characteristic{{
		Id: "2",
		Name: "nameChar",
		RdfType: model.SES_ONTOLOGY_CHARACTERISTIC,
		Type: model.Integer,
	}}

	controller := Controller{}
	err, code := controller.ValidateConcept(concept)
	if err == nil && code == http.StatusOK {
		t.Log(concept)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationValidWith2Characteristic(t *testing.T) {
	concept := model.Concept{}
	concept.Id = "1"
	concept.Name = "name"
	concept.RdfType = model.SES_ONTOLOGY_CONCEPT
	concept.Characteristics = []model.Characteristic{{
		Id: "2",
		Name: "Char1",
		RdfType: model.SES_ONTOLOGY_CHARACTERISTIC,
		Type: model.Structure,
	}}
	concept.Characteristics = append(concept.Characteristics, model.Characteristic{
		Id: "3",
		Name: "char2",
		RdfType: model.SES_ONTOLOGY_CHARACTERISTIC,
		Type: model.Structure,
	})

	controller := Controller{}
	err, code := controller.ValidateConcept(concept)
	if err == nil && code == http.StatusOK {
		t.Log(concept)
	} else {
		t.Fatal(err, code)
	}
}

func TestValidationValidWith2CharacteristicAndSubSubs(t *testing.T) {
	concept := model.Concept{}
	concept.Id = "1"
	concept.Name = "name"
	concept.RdfType = model.SES_ONTOLOGY_CONCEPT
	concept.Characteristics = []model.Characteristic{{
		Id: "2",
		Name: "Char2",
		RdfType: model.SES_ONTOLOGY_CHARACTERISTIC,
		Type: model.Structure,
		SubCharacteristics: []model.Characteristic{{
			Id: "2a",
			Name: "char2a",
			RdfType: model.SES_ONTOLOGY_CHARACTERISTIC,
			Type: model.Structure,
			SubCharacteristics: []model.Characteristic{{
				Id: "2a.a",
				Name: "char2a",
				RdfType: model.SES_ONTOLOGY_CHARACTERISTIC,
				Type: model.Boolean,
			}},
		}},
	}}
	concept.Characteristics = append(concept.Characteristics, model.Characteristic{
		Id: "3",
		Name: "char3",
		RdfType: model.SES_ONTOLOGY_CHARACTERISTIC,
		Type: model.Boolean,
	})

	controller := Controller{}
	err, code := controller.ValidateConcept(concept)
	if err == nil && code == http.StatusOK {
		t.Log(concept)
	} else {
		t.Fatal(err, code)
	}
}