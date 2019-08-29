package controller

import (
	"encoding/json"
	"github.com/SENERGY-Platform/semantic-repository/lib/database/sparql/rdf"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/piprate/json-gold/ld"
	"github.com/pkg/errors"
	"log"
	"reflect"
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
	contextDoc := getDeviceTypeContext()
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

func (*Controller) RdfXmlToSingleResult(rdfxml string, result interface{}, rootElement string) (err error) {
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
	options.UseNativeTypes = true
	doc, err := proc.FromRDF(strings.Join(turtle, ""), options)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: FromRDF()", err)
		return err
	}
	contextDoc := map[string]interface{}{}
	switch result.(type) {
	case *model.DeviceType:
		contextDoc = getDeviceTypeContext()
	case *model.Concept:
		contextDoc = getConceptContext()
	case *model.Characteristic:
		contextDoc = getCharacteristicsContext()
	default:
		debug.PrintStack()
		log.Println("Unknown model type:", reflect.TypeOf(result))
		return err
	}

	graph := map[string]interface{}{}
	graph["@context"] = contextDoc
	graph["@graph"] = doc

	cont := map[string]interface{}{}
	cont["@context"] = contextDoc
	switch result.(type) {
	case *model.DeviceType:
		cont["@type"] = model.SES_ONTOLOGY_DEVICE_TYPE
	case *model.Concept:
		cont["@type"] = model.SES_ONTOLOGY_CONCEPT
	case *model.Characteristic:
		cont["@type"] = model.SES_ONTOLOGY_CHARACTERISTIC
		cont["@id"] = rootElement
	default:
		debug.PrintStack()
		log.Println("Unknown model type:", reflect.TypeOf(result))
		return err
	}

	cont["@embed"] = "@always"

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
		turtle = append(turtle, triple.Serialize(rdf.NQuads))
	}
	return turtle, nil
}

func getDeviceTypeContext() map[string]interface{} {
	contextDoc := map[string]interface{}{
		"id":           "@id",
		"rdf_type":     "@type",
		"name":         model.RDFS_LABEL,
		"device_class": model.SES_ONTOLOGY_HAS_DEVICE_CLASS,
		"concept_id": map[string]interface{}{
			"@id":   model.SES_ONTOLOGY_HAS_CONCEPT,
			"@type": "@id",
		},
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
	}
	return contextDoc
}

func getConceptContext() map[string]interface{} {
	contextDoc := map[string]interface{}{
		"id":       "@id",
		"rdf_type": "@type",
		"name":     model.RDFS_LABEL,
		"characteristic_ids": map[string]interface{}{
			"@id":        model.SES_ONTOLOGY_HAS_CHARACTERISTIC,
			"@type":      "@id",
			"@container": "@set",
		},
	}
	return contextDoc
}

func getCharacteristicsContext() map[string]interface{} {
	contextDoc := map[string]interface{}{
		"id":       "@id",
		"rdf_type": "@type",
		"name":     model.RDFS_LABEL,
		"type": map[string]interface{}{
			"@id":   model.SES_ONTOLOGY_HAS_VALUE_TYPE,
			"@type": "@id",
		},
		"sub_characteristics": map[string]interface{}{
			"@id":        model.SES_ONTOLOGY_HAS_SUB_CHARACTERISTIC,
			"@container": "@set",
		},
		"value": model.SES_ONTOLOGY_HAS_VALUE,
		"min_value": model.SES_ONTOLOGY_HAS_MIN_VALUE,
		"max_value": model.SES_ONTOLOGY_HAS_MAX_VALUE,
	}
	return contextDoc
}
