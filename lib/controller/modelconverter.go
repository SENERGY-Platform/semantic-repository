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
	turtle, err := rdfxmlToTurtle(rdfxml)
	if len(turtle) == 0 {
		return nil
	}
	if err != nil {
		debug.PrintStack()
		log.Println("Error: FromRDF()", err)
		return err
	}

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.ProcessingMode = ld.JsonLd_1_1
	doc, err := proc.FromRDF(strings.Join(turtle, ""), options)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: FromRDF()", err)
		return err
	}
	contextDoc := getContext()
	graph := map[string]interface{}{}
	graph["@context"] = contextDoc
	graph["@graph"] = doc

	cont := map[string]interface{}{}
	cont["@context"] = contextDoc

	flattenDoc, err := proc.Flatten(graph, cont, options)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Flatten()", err)
		return err
	}

	flattenDocCast, ok := flattenDoc.(map[string]interface{})
	if !ok {
		debug.PrintStack()
		log.Println("Error: FlattenDoc casting()", ok)
		return errors.New("Error casting")
	}
	if flattenDocCast["@graph"] == nil {
		return nil
	}

	flattenDocGraph := flattenDocCast["@graph"].([]interface{})
	b, err := json.Marshal(flattenDocGraph)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Marshal()", err)
		return err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Unmarshal()", err)
		return err
	}
	return nil
}

func (*Controller) RdfXmlToSingleResult(rdfxml string, result *model.DeviceType) (err error) {
	turtle, err := rdfxmlToTurtle(rdfxml)
	if len(turtle) == 0 {
		return nil
	}
	if err != nil {
		debug.PrintStack()
		log.Println("Error: FromRDF()", err)
		return err
	}

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.ProcessingMode = ld.JsonLd_1_1
	doc, err := proc.FromRDF(strings.Join(turtle, ""), options)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: FromRDF()", err)
		return err
	}
	contextDoc := getContext()
	graph := map[string]interface{}{}
	graph["@context"] = contextDoc
	graph["@graph"] = doc

	cont := map[string]interface{}{}
	cont["@context"] = contextDoc
	cont["@type"] = model.SES_ONTOLOGY_DEVICE_TYPE

	frameDoc, err := proc.Frame(graph, cont, options)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Frame()", err)
		return err
	}

	frameDocGraph, ok := frameDoc["@graph"].([]interface{})
	if !ok {
		debug.PrintStack()
		log.Println("Error: FrameDoc casting()", ok)
		return errors.New("Error casting")
	}

	b, err := json.Marshal(frameDocGraph[0])
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Marshal()", err)
		return err
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Unmarshal()", err)
		return err
	}
	return nil
}


func rdfxmlToTurtle(rdfxml string) (result []string, err error) {
	triples, err := rdf.NewTripleDecoder(strings.NewReader(rdfxml), rdf.RDFXML).DecodeAll()
	if err != nil {
		debug.PrintStack()
		log.Println("Error: NewTripleDecoder()", err)
		return nil, err
	}
	if len(triples) == 0 {
		log.Println("No triples found")
		return nil, nil
	}
	turtle := []string{}
	for _, triple := range triples {
		turtle = append(turtle, triple.Serialize(rdf.Turtle))
	}
	return turtle, nil
}

func getContext() map[string]interface{} {
	contextDoc := map[string]interface{}{
		"id":           "@id",
		"rdf_type":         "@type",
		"name":         model.RDFS_LABEL,
		"device_class": model.SES_ONTOLOGY_HAS_DEVICE_CLASS,
		"services": map[string]interface{}{
			"@id":        model.SES_ONTOLOGY_HAS_SERVICE,
			"@container": "@set",
		},
		"aspects": map[string]interface{}{
			"@id":        model.SES_ONTOLOGY_REFERS_TO,
			"@container": "@set",
		},
		"functions": map[string]interface{}{
			"@id":        model.SES_ONTOLOGY_EXPOSES_FUNCTION,
			"@container": "@set",
		},
		"concept_ids": map[string]interface{}{
			"@id":        model.SES_ONTOLOGY_HAS_CONCEPT,
			"@type":      "@id",
			"@container": "@set",
		},
	}
	return contextDoc
}
