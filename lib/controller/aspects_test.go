package controller

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"net/http"
	"testing"
)

func TestValidAspect (t *testing.T) {
	aspects := []model.Aspect{}
	aspects = append(aspects, model.Aspect{Id: "urn:infai:ses:aspect:122", Name: "Air", RdfType: model.SES_ONTOLOGY_ASPECT})

	controller := Controller{}
	err, code := controller.ValidateAspects(aspects)
	if err == nil && code == http.StatusOK {
		t.Log(aspects)
	} else {
		t.Fatal(err, code)
	}
}

func TestAspectNoData (t *testing.T) {
	aspects := []model.Aspect{}

	controller := Controller{}
	err, code := controller.ValidateAspects(aspects)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestAspectMissingId (t *testing.T) {
	aspects := []model.Aspect{}
	aspects = append(aspects, model.Aspect{Id: ""})

	controller := Controller{}
	err, code := controller.ValidateAspects(aspects)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestAspectMissingName (t *testing.T) {
	aspects := []model.Aspect{}
	aspects = append(aspects, model.Aspect{Id: "urn:infai:ses:aspect:122", Name: ""})

	controller := Controller{}
	err, code := controller.ValidateAspects(aspects)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}

func TestAspectWrongType (t *testing.T) {
	aspects := []model.Aspect{}
	aspects = append(aspects, model.Aspect{Id: "urn:infai:ses:aspect:122", Name: "Air", RdfType: "wrongType"})

	controller := Controller{}
	err, code := controller.ValidateAspects(aspects)
	if err != nil && code == http.StatusBadRequest {
		t.Log(err)
	} else {
		t.Fatal(err, code)
	}
}
