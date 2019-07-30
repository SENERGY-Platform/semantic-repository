package controller

import (
	"encoding/json"
	"github.com/SENERGY-Platform/semantic-repository/lib/database/sparql/rdf"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/piprate/json-gold/ld"
	"github.com/pkg/errors"
	"log"
	"runtime/debug"
	"strings"
)

func (*Controller) RdfXmlToModel(rdfxml string, result interface{}) (err error) {
	triples, err := rdf.NewTripleDecoder(strings.NewReader(rdfxml), rdf.RDFXML).DecodeAll()
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
		"id":           "@id",
		"type":         "@type",
		"name":         model.RDFS_LABEL,
		"device_class": model.SES_ONTOLOGY_HAS_DEVICE_CLASS,
		"services":     model.SES_ONTOLOGY_HAS_SERVICE,
		"aspects":      model.SES_ONTOLOGY_REFERS_TO,
		"functions":    model.SES_ONTOLOGY_EXPOSES_FUNCTION,
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

func (*Controller) RdfXmlToSingleResult(rdfxml string, result *model.DeviceType) (err error) {
	triples, err := rdf.NewTripleDecoder(strings.NewReader(rdfxml), rdf.RDFXML).DecodeAll()
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
		return err
	}

	flattenDocGraph, ok := flattenDoc["@graph"].([]interface{})
	if !ok {
		debug.PrintStack()
		log.Println("Error: Flatten()", ok)
		return errors.New("Error casting")
	}

	b, err := json.Marshal(flattenDocGraph[0])

	err = json.Unmarshal(b, &result)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Unmarshal()", err)
		return err
	}
	return nil
}