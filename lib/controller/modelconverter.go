package controller

import (
	"encoding/json"
	"github.com/SENERGY-Platform/semantic-repository/lib/database/sparql/rdf"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/piprate/json-gold/ld"
	"log"
	"runtime/debug"
	"strings"
)

func (*Controller) RdfXmlToModel(deviceClasses string, result interface{}) (err error) {
	triples, err := rdf.NewTripleDecoder(strings.NewReader(deviceClasses), rdf.RDFXML).DecodeAll()
	if err != nil {
		debug.PrintStack()
		log.Println("Error: NewTripleDecoder()", err)
		return err
	}

	log.Println("triple:", triples)

	turtle := []string{}
	for _, triple := range triples {
		turtle = append(turtle, triple.Serialize(rdf.Turtle))
	}
	log.Println("turtle:", turtle)
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	//options.CompactArrays = false
	options.ProcessingMode = ld.JsonLd_1_1
	doc, err := proc.FromRDF(strings.Join(turtle, ""), options)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: FromRDF()", err)
		return err
	}
	log.Println("doc:", doc)
	context := map[string]interface{}{
		"@context": map[string]interface{}{
			"rdfs": "http://www.w3.org/2000/01/rdf-schema#",
			"id": "@id",
			"type": "@type",
			"name": "rdfs:label",
			//"device_class": model.SES_ONTOLOGY_HAS_DEVICE_CLASS,
			//"services": model.SES_ONTOLOGY_HAS_SERVICE,
			//"aspects": model.SES_ONTOLOGY_REFERS_TO,
			"functions": model.SES_ONTOLOGY_EXPOSES_FUNCTION,
			"concept_ids": map[string]interface{}{
				"@id": model.SES_ONTOLOGY_HAS_CONCEPT,
				"@type": "@id",
				"@container": "@set",
			},
		},
	}
	framedDoc, err := proc.Frame(doc, context, options)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Frame()", err)
		return err
	}
	log.Println("framedDoc:", framedDoc)
	b, err := json.Marshal(framedDoc["@graph"])
	log.Println("b:", string(b))
	err = json.Unmarshal(b, &result)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Unmarshal()", err)
		return err
	}
	log.Println("b:", b)
	return nil
}
