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
	aspects := []model.Aspect{}
	rdfxml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:sesame="http://www.openrdf.org/schema/sesame#" xmlns:owl="http://www.w3.org/2002/07/owl#" xmlns:xsd="http://www.w3.org/2001/XMLSchema#" xmlns:fn="http://www.w3.org/2005/xpath-functions#">
<rdf:Description rdf:about="urn:infai:ses:aspect:4444">
<rdfs:label>Lighting</rdfs:label>
<rdf:type rdf:resource="https://senergy.infai.org/ontology/Aspect"/>
</rdf:Description>
</rdf:RDF>`

	controller := Controller{}
	err := controller.RdfXmlToModel(rdfxml, &aspects)
	if err == nil {
		if aspects[0].Id != "urn:infai:ses:aspect:4444" {
			t.Fatal("wrong id")
		}
		if aspects[0].Name != "Lighting" {
			t.Fatal("wrong name")
		}
		if aspects[0].Type != "https://senergy.infai.org/ontology/Aspect" {
			t.Fatal("wrong type")
		}
		t.Log(aspects)
	} else {
		t.Fatal(err)
	}
}

func TestDeviceClass(t *testing.T) {
	deviceClass := []model.DeviceClass{}
	rdfxml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:sesame="http://www.openrdf.org/schema/sesame#" xmlns:owl="http://www.w3.org/2002/07/owl#" xmlns:xsd="http://www.w3.org/2001/XMLSchema#" xmlns:fn="http://www.w3.org/2005/xpath-functions#">
<rdf:Description rdf:about="urn:infai:ses:deviceclass:2222">
<rdf:type rdf:resource="https://senergy.infai.org/ontology/DeviceClass"/>
<rdfs:label>Lamp</rdfs:label>
</rdf:Description>
</rdf:RDF>`

	controller := Controller{}
	err := controller.RdfXmlToModel(rdfxml, &deviceClass)
	if err == nil {
		if deviceClass[0].Id != "urn:infai:ses:deviceclass:2222" {
			t.Fatal("wrong id")
		}
		if deviceClass[0].Name != "Lamp" {
			t.Fatal("wrong name")
		}
		if deviceClass[0].Type != "https://senergy.infai.org/ontology/DeviceClass" {
			t.Fatal("wrong type")
		}
		t.Log(deviceClass)
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
		if deviceType.Id != "urn:infai:ses:devicetype:1111" {
			t.Fatal("wrong id")
		}
		if deviceType.Type != "https://senergy.infai.org/ontology/DeviceType" {
			t.Fatal("wrong type")
		}
		if deviceType.Name != "Philips Hue Color" {
			t.Fatal("wrong name")
		}
		if deviceType.Image != "" {
			t.Fatal("wrong image")
		}
		if deviceType.Description != "" {
			t.Fatal("wrong description")
		}
		// DeviceClass
		if deviceType.DeviceClass.Id != "urn:infai:ses:deviceclass:2222" {
			t.Fatal("wrong device class id")
		}
		if deviceType.DeviceClass.Name != "Lamp" {
			t.Fatal("wrong device class name")
		}
		if deviceType.DeviceClass.Type != "https://senergy.infai.org/ontology/DeviceClass" {
			t.Fatal("wrong device class type")
		}
		// Service Index 0
		if deviceType.Services[0].Id != "urn:infai:ses:service:3333" {
			t.Fatal("wrong service id -> index 0")
		}
		if deviceType.Services[0].Type != "https://senergy.infai.org/ontology/Service" {
			t.Fatal("wrong service type -> index 0")
		}
		if deviceType.Services[0].Name != "setBrigthness" {
			t.Fatal("wrong service name -> index 0")
		}
		if deviceType.Services[0].Description != "" {
			t.Fatal("wrong service description -> index 0")
		}
		if deviceType.Services[0].ProtocolId != "" {
			t.Fatal("wrong service ProtocolId -> index 0")
		}
		if deviceType.Services[0].LocalId != "" {
			t.Fatal("wrong service LocalId -> index 0")
		}
		if deviceType.Services[0].Functions[0].Id != "urn:infai:ses:function:5555" {
			t.Fatal("wrong functions id -> index 0,0")
		}
		if deviceType.Services[0].Functions[0].Type != "https://senergy.infai.org/ontology/ControllingFunction" {
			t.Fatal("wrong functions type -> index 0,0")
		}
		if deviceType.Services[0].Functions[0].Name != "brightnessAdjustment" {
			t.Fatal("wrong functions name -> index 0,0")
		}
		if deviceType.Services[0].Functions[0].ConceptIds[0] != "urn:infai:ses:concept:7777" {
			t.Fatal("wrong functions concept id -> index 0,0,0")
		}
		if deviceType.Services[0].Functions[0].ConceptIds[1] != "urn:infai:ses:concept:6666" {
			t.Fatal("wrong functions concept id -> index 0,0,1")
		}
		if deviceType.Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4444" {
			t.Fatal("wrong aspects id -> index 0,0")
		}
		if deviceType.Services[0].Aspects[0].Type != "https://senergy.infai.org/ontology/Aspect" {
			t.Fatal("wrong aspects type -> index 0,0")
		}
		if deviceType.Services[0].Aspects[0].Name != "Lighting" {
			t.Fatal("wrong aspects name -> index 0,0")
		}
		if deviceType.Services[0].Aspects[0].Name != "Lighting" {
			t.Fatal("wrong aspects name -> index 0,0")
		}
		t.Log(deviceType)
	} else {
		t.Fatal(err)
	}

}
