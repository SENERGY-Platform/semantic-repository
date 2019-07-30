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

	if len(triples) == 0 {
		log.Println("No triples found")
		return nil
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
		return err
	}
	contextDoc := map[string]interface{}{
		"rdfs":        "http://www.w3.org/2000/01/rdf-schema#",
		"id":          "@id",
		"type":        "@type",
		"name":        "rdfs:label",
		"functions":   model.SES_ONTOLOGY_EXPOSES_FUNCTION,
		"concept_ids": map[string]interface{}{
			"@id":        model.SES_ONTOLOGY_HAS_CONCEPT,
			"@type":      "@id",
			"@container": "@set",
		},
	}

	contextFrame := contextDoc

	graph := map[string]interface{}{
	}
	graph["@context"] = contextDoc
	graph["@graph"] = doc

	cont := map[string]interface{}{}
	cont["@context"] = contextFrame

	flattenDoc, err := proc.Flatten(graph, cont, options)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Flatten()", err)
		return err
	}

	flattenDocCast := flattenDoc.(map[string]interface{})
	if flattenDocCast["@graph"] == nil {
		return nil
	}

	flattenDocGraph := flattenDocCast["@graph"].([]interface{})
	b, err := json.Marshal(flattenDocGraph)
	err = json.Unmarshal(b, &result)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Unmarshal()", err)
		return err
	}
	return nil
}
