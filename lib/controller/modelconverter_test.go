package controller

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"testing"
)

func TestControllingFunction(t *testing.T) {
	contFunc := []model.Function{}
	rdfxml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:sesame="http://www.openrdf.org/schema/sesame#" xmlns:owl="http://www.w3.org/2002/07/owl#" xmlns:xsd="http://www.w3.org/2001/XMLSchema#" xmlns:fn="http://www.w3.org/2005/xpath-functions#">
<rdf:Description rdf:about="urn:infai:ses:function:5555">
<rdf:type rdf:resource="https://senergy.infai.org/ontology/ControllingFunction"/>
<rdfs:label>brightnessAdjustment</rdfs:label>
<hasConcept xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:infai:ses:concept:6666"/>
<hasConcept xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:infai:ses:concept:7777"/>
</rdf:Description>
</rdf:RDF>`

	controller := Controller{}
	err := controller.RdfXmlToModel(rdfxml, &contFunc)
	if err == nil {
		// check content
		if contFunc[0].Id != "urn:infai:ses:function:5555" {
			t.Fatal("wrong id")
		}
		if contFunc[0].Name != "brightnessAdjustment" {
			t.Fatal("wrong name")
		}
		if contFunc[0].Type != "https://senergy.infai.org/ontology/ControllingFunction" {
			t.Fatal("wrong type")
		}
		if contFunc[0].ConceptIds[0] != "urn:infai:ses:concept:6666" {
			t.Fatal("wrong concept id -> index 0")
		}
		if contFunc[0].ConceptIds[1] != "urn:infai:ses:concept:7777" {
			t.Fatal("wrong concept id -> index 1")
		}
		t.Log(contFunc)

	} else {
		t.Fatal(err)
	}
}

func TestAspects(t *testing.T) {
	contFunc := []model.Function{}
	rdfxml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:sesame="http://www.openrdf.org/schema/sesame#" xmlns:owl="http://www.w3.org/2002/07/owl#" xmlns:xsd="http://www.w3.org/2001/XMLSchema#" xmlns:fn="http://www.w3.org/2005/xpath-functions#">
<rdf:Description rdf:about="urn:infai:ses:function:5555">
<rdf:type rdf:resource="https://senergy.infai.org/ontology/ControllingFunction"/>
<rdfs:label>brightnessAdjustment</rdfs:label>
<hasConcept xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:infai:ses:concept:6666"/>
<hasConcept xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:infai:ses:concept:7777"/>
</rdf:Description>
</rdf:RDF>`

	controller := Controller{}
	err := controller.RdfXmlToModel(rdfxml, &contFunc)
	if err == nil {
		t.Log(contFunc)
	} else {
		t.Fatal(err)
	}
}

func TestDeviceType(t *testing.T) {

	deviceType := model.DeviceType{}

	rdfxml := `<?xml version="1.0" encoding="UTF-8"?>
<rdf:RDF
	xmlns:x="http://www.w3.org/1999/02/22-rdf-syntax-ns#type"
	xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
	xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#"
	xmlns:sesame="http://www.openrdf.org/schema/sesame#"
	xmlns:owl="http://www.w3.org/2002/07/owl#"
	xmlns:xsd="http://www.w3.org/2001/XMLSchema#"
	xmlns:fn="http://www.w3.org/2005/xpath-functions#">

<rdf:Description rdf:about="urn:infai:ses:function:5555">
	<hasConcept xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:infai:ses:concept:7777"/>
	<hasConcept xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:infai:ses:concept:6666"/>
	<rdf:type rdf:resource="https://senergy.infai.org/ontology/ControllingFunction"/>
	<rdfs:label>brightnessAdjustment</rdfs:label>
</rdf:Description>

<rdf:Description rdf:about="urn:infai:ses:service:3333">
	<exposesFunction xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:infai:ses:function:5555"/>
	<refersTo xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:infai:ses:aspect:4444"/>
	<rdf:type rdf:resource="https://senergy.infai.org/ontology/Service"/>
	<rdfs:label>setBrigthness</rdfs:label>
</rdf:Description>

<rdf:Description rdf:about="urn:infai:ses:deviceclass:2222">
	<rdf:type rdf:resource="https://senergy.infai.org/ontology/DeviceClass"/>
	<rdfs:label>Lamp</rdfs:label>
</rdf:Description>

<rdf:Description rdf:about="urn:infai:ses:devicetype:1111">
	<hasService xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:infai:ses:service:3333"/>
	<hasDeviceClass xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:infai:ses:deviceclass:2222"/>
	<rdf:type rdf:resource="https://senergy.infai.org/ontology/DeviceType"/>
	<rdfs:label>Philips Hue Color</rdfs:label>
</rdf:Description>

<rdf:Description rdf:about="urn:infai:ses:aspect:4444">
	<rdf:type rdf:resource="https://senergy.infai.org/ontology/Aspect"/>
	<rdfs:label>Lighting</rdfs:label>
</rdf:Description>

</rdf:RDF>
`

	controller := Controller{}
	err := controller.RdfXmlToSingleResult(rdfxml, &deviceType)
	if err == nil {
		t.Log(deviceType)
	} else {
		t.Fatal(err)
	}

}
