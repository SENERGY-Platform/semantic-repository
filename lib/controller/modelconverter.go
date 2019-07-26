package controller

import (
	"encoding/json"
	"github.com/SENERGY-Platform/semantic-repository/lib/database/sparql/rdf"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/piprate/json-gold/ld"
	"log"
	"reflect"
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
	log.Println(reflect.TypeOf(doc))
	contextDoc := map[string]interface{}{

			"rdfs": "http://www.w3.org/2000/01/rdf-schema#",
			"id": "@id",
			"type": "@type",
			"name": "rdfs:label",
			"functions": model.SES_ONTOLOGY_EXPOSES_FUNCTION,
		    "concept_ids": model.SES_ONTOLOGY_HAS_CONCEPT,
	}

	contextFrame := map[string]interface{}{

		"rdfs": "http://www.w3.org/2000/01/rdf-schema#",
		"id": "@id",
		"type": "@type",
		"name": "rdfs:label",
		"functions": model.SES_ONTOLOGY_EXPOSES_FUNCTION,
		"concept_ids": model.SES_ONTOLOGY_HAS_CONCEPT,
	}

	graph := map[string]interface{} {
	}
	graph["@context"] = contextDoc
	graph["@graph"] = doc

	cont := map[string]interface{} {}
	cont["@context"] = contextFrame

	log.Println("graph:", graph)
	framedDoc, err := proc.Flatten(graph, cont, options)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Frame()", err)
		return err
	}
	log.Println("framedDoc:", framedDoc)
	frameDocM, _ := json.Marshal(framedDoc)
	log.Println("graphNew:", string(frameDocM))
	xxx, _ := framedDoc.(map[string]interface{})
	gr := xxx["@graph"].([]interface{})
	log.Println("gr:", gr)
	b, err := json.Marshal(gr)
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
