package controller

import (
	"encoding/json"
	"github.com/SENERGY-Platform/semantic-repository/lib/database/sparql/rdf"
	"github.com/piprate/json-gold/ld"
	"log"
	"strings"
)

func (*Controller) ByteToModel(deviceClasses []byte, result interface{}) (err error)  {
	triples, err := rdf.NewTripleDecoder(strings.NewReader(string(deviceClasses)), rdf.RDFXML).DecodeAll()
	if err != nil {
		log.Println("GetDeviceClasses ERROR: NewTripleDecoder", err)
		return err
	}
	turtle := []string{}
	for _, triple := range triples {
		turtle = append(turtle, triple.Serialize(rdf.Turtle))
	}
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.ProcessingMode = ld.JsonLd_1_1
	doc, _ := proc.FromRDF(strings.Join(turtle, ""), options)
	context := map[string]interface{}{
		"@context": map[string]interface{}{
			"rdfs":   "http://www.w3.org/2000/01/rdf-schema#",
			"xsd":    "http://www.w3.org/2001/XMLSchema#",
			"schema": "http://schema.org/",
			"name":   "rdfs:label",
			"id":     "@id",
			"type":   "@type",
		},
	}
	framedDoc, err := proc.Frame(doc, context, options)
	if err != nil {
		log.Println("GetDeviceClasses ERROR: Frame", err)
		return err
	}
	b, err := json.Marshal(framedDoc["@graph"])
	err = json.Unmarshal(b, &result)
	if err != nil {
		log.Println("GetDeviceClasses ERROR: Unmarshal", err)
		return err
	}
	return nil
}