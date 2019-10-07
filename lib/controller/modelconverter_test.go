/*
 *
 * Copyright 2019 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *
 */

package controller

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"testing"
)

func TestNoData_RdfXmlToModel(t *testing.T) {
	contFunc := []model.Function{}
	rdfxml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:sesame="http://www.openrdf.org/schema/sesame#" xmlns:owl="http://www.w3.org/2002/07/owl#" xmlns:xsd="http://www.w3.org/2001/XMLSchema#" xmlns:fn="http://www.w3.org/2005/xpath-functions#"> </rdf:RDF>`

	controller := Controller{}
	err := controller.RdfXmlToModel(rdfxml, &contFunc)
	if err == nil {
		if len(contFunc) != 0 {
			t.Fatal("wrong response")
		}
		t.Log(contFunc)

	} else {
		t.Fatal(err)
	}
}

func TestNoData_RdfXmlToSingleResult(t *testing.T) {
	deviceType := model.DeviceType{}
	rdfxml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:sesame="http://www.openrdf.org/schema/sesame#" xmlns:owl="http://www.w3.org/2002/07/owl#" xmlns:xsd="http://www.w3.org/2001/XMLSchema#" xmlns:fn="http://www.w3.org/2005/xpath-functions#"> </rdf:RDF>`

	controller := Controller{}
	err := controller.RdfXmlFrame(rdfxml, &deviceType, "")
	if err == nil {
		if deviceType.RdfType != "" {
			t.Fatal("wrong response")
		}
		t.Log(deviceType)

	} else {
		t.Fatal(err)
	}
}

func TestControllingFunction(t *testing.T) {
	contFunc := []model.Function{}
	rdfxml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:sesame="http://www.openrdf.org/schema/sesame#" xmlns:owl="http://www.w3.org/2002/07/owl#" xmlns:xsd="http://www.w3.org/2001/XMLSchema#" xmlns:fn="http://www.w3.org/2005/xpath-functions#">
<rdf:Description rdf:about="urn:infai:ses:function:5555">
<rdf:type rdf:resource="https://senergy.infai.org/ontology/ControllingFunction"/>
<rdfs:label>brightnessAdjustment</rdfs:label>
<hasConcept xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:infai:ses:concept:6666"/>
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
		if contFunc[0].RdfType != "https://senergy.infai.org/ontology/ControllingFunction" {
			t.Fatal("wrong type")
		}
		if contFunc[0].ConceptId != "urn:infai:ses:concept:6666" {
			t.Fatal("wrong concept id")
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
		if aspects[0].RdfType != "https://senergy.infai.org/ontology/Aspect" {
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
		if deviceClass[0].RdfType != "https://senergy.infai.org/ontology/DeviceClass" {
			t.Fatal("wrong type")
		}
		t.Log(deviceClass)
	} else {
		t.Fatal(err)
	}
}

func TestDeviceType(t *testing.T) {

	deviceType := model.DeviceType{}

	rdfxml :=
		`<?xml version="1.0" encoding="UTF-8"?>
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
	err := controller.RdfXmlFrame(rdfxml, &deviceType, "urn:infai:ses:devicetype:1111")
	if err == nil {
		if deviceType.Id != "urn:infai:ses:devicetype:1111" {
			t.Fatal("wrong id")
		}
		if deviceType.RdfType != "https://senergy.infai.org/ontology/DeviceType" {
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
		if deviceType.DeviceClass.RdfType != "https://senergy.infai.org/ontology/DeviceClass" {
			t.Fatal("wrong device class type")
		}
		// Service Index 0
		if deviceType.Services[0].Id != "urn:infai:ses:service:3333" {
			t.Fatal("wrong service id -> index 0")
		}
		if deviceType.Services[0].RdfType != "https://senergy.infai.org/ontology/Service" {
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
		if deviceType.Services[0].Functions[0].RdfType != "https://senergy.infai.org/ontology/ControllingFunction" {
			t.Fatal("wrong functions type -> index 0,0")
		}
		if deviceType.Services[0].Functions[0].Name != "brightnessAdjustment" {
			t.Fatal("wrong functions name -> index 0,0")
		}
		if deviceType.Services[0].Functions[0].ConceptId != "urn:infai:ses:concept:7777" {
			t.Fatal("wrong functions concept id -> index 0,0,0")
		}
		if deviceType.Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4444" {
			t.Fatal("wrong aspects id -> index 0,0")
		}
		if deviceType.Services[0].Aspects[0].RdfType != "https://senergy.infai.org/ontology/Aspect" {
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

func TestConcept(t *testing.T) {

	concept := model.Concept{}

	rdfxml :=
		`<?xml version="1.0" encoding="UTF-8"?>
<rdf:RDF
	xmlns:x="http://www.w3.org/1999/02/22-rdf-syntax-ns#type"
	xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
	xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#"
	xmlns:sesame="http://www.openrdf.org/schema/sesame#"
	xmlns:owl="http://www.w3.org/2002/07/owl#"
	xmlns:xsd="http://www.w3.org/2001/XMLSchema#"
	xmlns:fn="http://www.w3.org/2005/xpath-functions#">

<rdf:Description rdf:about="urn:ses:infai:concept:1a1a1a">
	<rdf:type rdf:resource="https://senergy.infai.org/ontology/Concept"/>
	<rdfs:label>color</rdfs:label>
	<hasCharacteristic xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:ses:infai:characteristic:2b2c2d"/>
</rdf:Description>


</rdf:RDF>
`

	controller := Controller{}
	err := controller.RdfXmlFrame(rdfxml, &concept, "urn:ses:infai:concept:1a1a1a")
	if err == nil {
		if concept.Id != "urn:ses:infai:concept:1a1a1a" {
			t.Fatal("wrong id")
		}
		if concept.Name != "color" {
			t.Fatal("wrong name")
		}
		if concept.RdfType != model.SES_ONTOLOGY_CONCEPT {
			t.Fatal("wrong rdf_type")
		}
		if concept.CharacteristicIds[0] != "urn:ses:infai:characteristic:2b2c2d" {
			t.Fatal("wrong Characteristics id", concept.CharacteristicIds[0])
		}
		t.Log(concept)
	} else {
		t.Fatal(err)
	}

}
