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




<rdf:Description rdf:about="urn:infai:ses:devicetype:1111">
	<hasService xmlns="https://senergy.infai.org/ontology/" rdf:resource="urn:infai:ses:service:3333"/>
	<rdf:type rdf:resource="https://senergy.infai.org/ontology/DeviceType"/>
	<rdfs:label>Philips Hue Color</rdfs:label>
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
			"@id": "https://senergy.infai.org/ontology/hasService",
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

	test1, _ := json.Marshal(flattenDocGraph)
	log.Println(string(test1))

	b, err := json.Marshal(flattenDocGraph[0])

	err = json.Unmarshal(b, &deviceType)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Unmarshal()", err)
		//return err
	}
	log.Println(deviceType)
	//return nil

}
