package controller

import (
	"encoding/json"
	"github.com/SENERGY-Platform/semantic-repository/lib/database/sparql/rdf"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/piprate/json-gold/ld"
	"log"
	"runtime/debug"
	"strings"
	"testing"
)

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
	triples, err := rdf.NewTripleDecoder(strings.NewReader(rdfxml), rdf.RDFXML).DecodeAll()
	if err != nil {
		debug.PrintStack()
		log.Println("Error: NewTripleDecoder()", err)
		//return err
	}

	if len(triples) == 0 {
		log.Println("No triples found")
		//return nil
	}

	turtle := []string{}
	for _, triple := range triples {
		turtle = append(turtle, triple.Serialize(rdf.Turtle))
	}

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	//options.CompactArrays = false
	options.ProcessingMode = ld.JsonLd_1_1
	doc, err := proc.FromRDF(strings.Join(turtle, ""), options)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: FromRDF()", err)
		//return err
	}
	contextDoc := map[string]interface{}{
		"id":"@id",
		"type":"@type",
		"name": model.RDFS_LABEL,
		"device_class": model.SES_ONTOLOGY_HAS_DEVICE_CLASS,
		"services": map[string]interface{}{
			"@id": model.SES_ONTOLOGY_HAS_SERVICE,
			"@container": "@set",
		},
		"aspects": map[string]interface{}{
			"@id": model.SES_ONTOLOGY_REFERS_TO,
			"@container": "@set",
		},
		"functions": map[string]interface{}{
			"@id": model.SES_ONTOLOGY_EXPOSES_FUNCTION,
			"@container": "@set",
		},
	}

	frameContext := contextDoc

	graph := map[string]interface{}{
	}
	graph["@context"] = contextDoc
	graph["@graph"] = doc


	cont := map[string]interface{}{}
	cont["@context"] = frameContext
	cont["@type"] = model.SES_ONTOLOGY_DEVICE_TYPE




	flattenDoc, err := proc.Frame(graph, cont, options)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Flatten()", err)
	}

	//flattenDocCast := flattenDoc.(map[string]interface{})
	//if flattenDocCast["@graph"] == nil {
	//	//return nil
	//}




	flattenDocGraph, ok := flattenDoc["@graph"].([]interface{})
	if !ok {
		//return nil
	}


	b, err := json.Marshal(flattenDocGraph[0])

	err = json.Unmarshal(b, &deviceType)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Unmarshal()", err)
		//return err
	}
	log.Println(string(b))
	log.Println(deviceType)
	//return nil

}
