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
	err := controller.RdfXmlToSingleResult(rdfxml, &deviceType)
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
	err := controller.RdfXmlToSingleResult(rdfxml, &deviceType)
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

<rdf:Description rdf:about="urn:ses:infai:characteristic:2b2c2d">
	<rdf:type rdf:resource="https://senergy.infai.org/ontology/Characteristic"/>
	<rdfs:label>structure</rdfs:label>
	<hasSubCharacteristic xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:ses:infai:characteristic:3e3e3e"/>
	<hasSubCharacteristic xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:ses:infai:characteristic:4r4r4r"/>
	<hasSubCharacteristic xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:ses:infai:characteristic:6z6z6z"/>
	<hasSubCharacteristic xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:ses:infai:characteristic:7u7u7u"/>
	<hasSubCharacteristic xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:ses:infai:characteristic:5t5t5t"/>
	<hasValueType xmlns="https://senergy.infai.org/ontology/" rdf:resource="https://schema.org/StructuredValue"/>
</rdf:Description>

<rdf:Description rdf:about="urn:ses:infai:concept:1a1a1a">
	<rdf:type rdf:resource="https://senergy.infai.org/ontology/Concept"/>
	<rdfs:label>color</rdfs:label>
	<hasCharacteristic xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:ses:infai:characteristic:2b2c2d"/>
</rdf:Description>

<rdf:Description rdf:about="urn:ses:infai:characteristic:3e3e3e">
	<hasValue xmlns="https://senergy.infai.org/ontology/">100</hasValue>
	<rdf:type rdf:resource="https://senergy.infai.org/ontology/Characteristic"/>
	<rdfs:label>nameString</rdfs:label>
	<hasValueType xmlns="https://senergy.infai.org/ontology/" rdf:resource="https://schema.org/Text"/>
</rdf:Description>

<rdf:Description rdf:about="urn:ses:infai:characteristic:4r4r4r">
	<hasMaxValue xmlns="https://senergy.infai.org/ontology/" rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">255</hasMaxValue>
	<hasMinValue xmlns="https://senergy.infai.org/ontology/" rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">0</hasMinValue>
	<hasValue xmlns="https://senergy.infai.org/ontology/" rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">122</hasValue>
	<rdf:type rdf:resource="https://senergy.infai.org/ontology/Characteristic"/>
	<rdfs:label>nameInteger</rdfs:label>
	<hasValueType xmlns="https://senergy.infai.org/ontology/" rdf:resource="https://schema.org/Integer"/>
</rdf:Description>

<rdf:Description rdf:about="urn:ses:infai:characteristic:6z6z6z">
	<hasMaxValue xmlns="https://senergy.infai.org/ontology/" rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">255</hasMaxValue>
	<hasMinValue xmlns="https://senergy.infai.org/ontology/" rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">0</hasMinValue>
	<hasValue xmlns="https://senergy.infai.org/ontology/" rdf:datatype="http://www.w3.org/2001/XMLSchema#double">1.2222E2</hasValue>
	<rdf:type rdf:resource="https://senergy.infai.org/ontology/Characteristic"/>
	<rdfs:label>nameFloat</rdfs:label>
	<hasValueType xmlns="https://senergy.infai.org/ontology/" rdf:resource="https://schema.org/Float"/>
</rdf:Description>

<rdf:Description rdf:about="urn:ses:infai:characteristic:7u7u7u">
	<hasMaxValue xmlns="https://senergy.infai.org/ontology/" rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">10</hasMaxValue>
	<hasMinValue xmlns="https://senergy.infai.org/ontology/" rdf:datatype="http://www.w3.org/2001/XMLSchema#integer">0</hasMinValue>
	<rdf:type rdf:resource="https://senergy.infai.org/ontology/Characteristic"/>
	<rdfs:label>nameNoValue</rdfs:label>
	<hasValueType xmlns="https://senergy.infai.org/ontology/" rdf:resource="https://schema.org/Float"/>
</rdf:Description>

<rdf:Description rdf:about="urn:ses:infai:characteristic:5t5t5t">
	<hasValue xmlns="https://senergy.infai.org/ontology/" rdf:datatype="http://www.w3.org/2001/XMLSchema#boolean">true</hasValue>
	<rdf:type rdf:resource="https://senergy.infai.org/ontology/Characteristic"/>
	<rdfs:label>nameBoolean</rdfs:label>
	<hasValueType xmlns="https://senergy.infai.org/ontology/" rdf:resource="https://schema.org/Boolean"/>
</rdf:Description>

</rdf:RDF>
`

	controller := Controller{}
	err := controller.RdfXmlToSingleResult(rdfxml, &concept)
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
		//if concept.Characteristics[0].Id != "urn:ses:infai:characteristic:2b2c2d" {
		//	t.Fatal("wrong Characteristics id", concept.Characteristics[0].Id)
		//}
		//if concept.Characteristics[0].Name != "structure" {
		//	t.Fatal("wrong Characteristics name")
		//}
		//if concept.Characteristics[0].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
		//	t.Fatal("wrong Characteristics rdf_type")
		//}
		//if concept.Characteristics[0].MinValue != nil {
		//	t.Fatal("wrong Characteristics MinValue", concept.Characteristics[0].MinValue)
		//}
		//if concept.Characteristics[0].MaxValue != nil {
		//	t.Fatal("wrong Characteristics MaxValue", concept.Characteristics[0].MaxValue)
		//}
		//if concept.Characteristics[0].Value != nil {
		//	t.Fatal("wrong Characteristics Value", concept.Characteristics[0].Value)
		//}
		//if concept.Characteristics[0].Type != model.Structure {
		//	t.Fatal("wrong Characteristics Type", concept.Characteristics[0].Type)
		//}
		//// id = urn:ses:infai:characteristic:3e3e3e
		//index := 0
		//if concept.Characteristics[0].SubCharacteristics[index].Id != "urn:ses:infai:characteristic:3e3e3e" {
		//	t.Fatal("wrong SubCharacteristics Id", concept.Characteristics[0].SubCharacteristics[index].Id)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].Name != "nameString" {
		//	t.Fatal("wrong SubCharacteristics Name", concept.Characteristics[0].SubCharacteristics[index].Name)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].Type != model.String {
		//	t.Fatal("wrong SubCharacteristics Type", concept.Characteristics[0].SubCharacteristics[index].Type)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
		//	t.Fatal("wrong SubCharacteristics RdfType", concept.Characteristics[0].SubCharacteristics[index].RdfType)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].Value != "100" {
		//	t.Fatal("wrong SubCharacteristics Value", concept.Characteristics[0].SubCharacteristics[index].Value)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].MaxValue != nil {
		//	t.Fatal("wrong SubCharacteristics MaxValue", concept.Characteristics[0].SubCharacteristics[index].MaxValue)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].MinValue != nil {
		//	t.Fatal("wrong SubCharacteristics MinValue", concept.Characteristics[0].SubCharacteristics[index].MinValue)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics != nil {
		//	t.Fatal("wrong SubCharacteristics SubCharacteristics", concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics)
		//}
		//// id = urn:ses:infai:characteristic:4r4r4r
		//index = 1
		//if concept.Characteristics[0].SubCharacteristics[index].Id != "urn:ses:infai:characteristic:4r4r4r" {
		//	t.Fatal("wrong SubCharacteristics Id", concept.Characteristics[0].SubCharacteristics[index].Id)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].Name != "nameInteger" {
		//	t.Fatal("wrong SubCharacteristics Name", concept.Characteristics[0].SubCharacteristics[index].Name)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].Type != model.Integer {
		//	t.Fatal("wrong SubCharacteristics Type", concept.Characteristics[0].SubCharacteristics[index].Type)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
		//	t.Fatal("wrong SubCharacteristics RdfType", concept.Characteristics[0].SubCharacteristics[index].RdfType)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].Value != float64(122) {
		//	t.Fatal("wrong SubCharacteristics Value", concept.Characteristics[0].SubCharacteristics[index].Value)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].MaxValue != float64(255) {
		//	t.Fatal("wrong SubCharacteristics MaxValue", concept.Characteristics[0].SubCharacteristics[index].MaxValue)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].MinValue != float64(0) {
		//	t.Fatal("wrong SubCharacteristics MinValue", concept.Characteristics[0].SubCharacteristics[index].MinValue)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics != nil {
		//	t.Fatal("wrong SubCharacteristics SubCharacteristics", concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics)
		//}
		//// id = urn:ses:infai:characteristic:6z6z6z
		//index = 2
		//if concept.Characteristics[0].SubCharacteristics[index].Id != "urn:ses:infai:characteristic:6z6z6z" {
		//	t.Fatal("wrong SubCharacteristics Id", concept.Characteristics[0].SubCharacteristics[index].Id)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].Name != "nameFloat" {
		//	t.Fatal("wrong SubCharacteristics Name", concept.Characteristics[0].SubCharacteristics[index].Name)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].Type != model.Float {
		//	t.Fatal("wrong SubCharacteristics Type", concept.Characteristics[0].SubCharacteristics[index].Type)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
		//	t.Fatal("wrong SubCharacteristics RdfType", concept.Characteristics[0].SubCharacteristics[index].RdfType)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].Value != float64(122.22) {
		//	t.Fatal("wrong SubCharacteristics Value", concept.Characteristics[0].SubCharacteristics[index].Value)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].MaxValue != float64(255) {
		//	t.Fatal("wrong SubCharacteristics MaxValue", concept.Characteristics[0].SubCharacteristics[index].MaxValue)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].MinValue != float64(0) {
		//	t.Fatal("wrong SubCharacteristics MinValue", concept.Characteristics[0].SubCharacteristics[index].MinValue)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics != nil {
		//	t.Fatal("wrong SubCharacteristics SubCharacteristics", concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics)
		//}
		//// id = urn:ses:infai:characteristic:7u7u7u
		//index = 3
		//if concept.Characteristics[0].SubCharacteristics[index].Id != "urn:ses:infai:characteristic:7u7u7u" {
		//	t.Fatal("wrong SubCharacteristics Id", concept.Characteristics[0].SubCharacteristics[index].Id)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].Name != "nameNoValue" {
		//	t.Fatal("wrong SubCharacteristics Name", concept.Characteristics[0].SubCharacteristics[index].Name)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].Type != model.Float {
		//	t.Fatal("wrong SubCharacteristics Type", concept.Characteristics[0].SubCharacteristics[index].Type)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
		//	t.Fatal("wrong SubCharacteristics RdfType", concept.Characteristics[0].SubCharacteristics[index].RdfType)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].Value != nil {
		//	t.Fatal("wrong SubCharacteristics Value", concept.Characteristics[0].SubCharacteristics[index].Value)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].MaxValue != float64(10) {
		//	t.Fatal("wrong SubCharacteristics MaxValue", concept.Characteristics[0].SubCharacteristics[index].MaxValue)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].MinValue != float64(0) {
		//	t.Fatal("wrong SubCharacteristics MinValue", concept.Characteristics[0].SubCharacteristics[index].MinValue)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics != nil {
		//	t.Fatal("wrong SubCharacteristics SubCharacteristics", concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics)
		//}
		//// id = urn:ses:infai:characteristic:5t5t5t
		//index = 4
		//if concept.Characteristics[0].SubCharacteristics[index].Id != "urn:ses:infai:characteristic:5t5t5t" {
		//	t.Fatal("wrong SubCharacteristics Id", concept.Characteristics[0].SubCharacteristics[index].Id)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].Name != "nameBoolean" {
		//	t.Fatal("wrong SubCharacteristics Name", concept.Characteristics[0].SubCharacteristics[index].Name)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].Type != model.Boolean {
		//	t.Fatal("wrong SubCharacteristics Type", concept.Characteristics[0].SubCharacteristics[index].Type)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
		//	t.Fatal("wrong SubCharacteristics RdfType", concept.Characteristics[0].SubCharacteristics[index].RdfType)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].Value != true {
		//	t.Fatal("wrong SubCharacteristics Value", concept.Characteristics[0].SubCharacteristics[index].Value)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].MaxValue != nil {
		//	t.Fatal("wrong SubCharacteristics MaxValue", concept.Characteristics[0].SubCharacteristics[index].MaxValue)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].MinValue != nil {
		//	t.Fatal("wrong SubCharacteristics MinValue", concept.Characteristics[0].SubCharacteristics[index].MinValue)
		//}
		//if concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics != nil {
		//	t.Fatal("wrong SubCharacteristics SubCharacteristics", concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics)
		//}
		t.Log(concept)
	} else {
		t.Fatal(err)
	}

}
