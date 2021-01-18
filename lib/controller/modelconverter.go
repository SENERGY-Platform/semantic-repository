/*
 *
 * Copyright 2019 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *
 */

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

//deprecated
func (this *Controller) RdfXmlToModel(rdfxml string, result interface{}) (err error) {
	return this.RdfXmlToModelWithContext(rdfxml, result, getDeviceTypeContext())
}

func (*Controller) RdfXmlToModelWithContext(rdfxml string, result interface{}, contextDoc map[string]interface{}) (err error) {
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

func (*Controller) RdfXmlFrame(rdfxml string, result interface{}, rootElement string) (err error) {
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
	case *[]model.DeviceType:
		contextDoc = getDeviceTypeContext()
	case *[]model.Concept:
		contextDoc = getConceptContext()
	case *[]model.Characteristic:
		contextDoc = getCharacteristicsContext()
	case *[]model.ConceptWithCharacteristics:
		contextDoc = getConceptCharacteristicContext()
	case *[]model.TotalCount:
		contextDoc = getTotalCountContext()
	case *[]model.Aspect:
		contextDoc = getAspectContext()
	case *[]model.DeviceClass:
		contextDoc = getDeviceClassContext()
	case *[]model.Function:
		contextDoc = getFunctionContext()
	default:
		debug.PrintStack()
		log.Println("Unknown model type:", reflect.TypeOf(result))
		return errors.New("Unknown model type:" + reflect.TypeOf(result).String())
	}

	graph := map[string]interface{}{}
	graph["@context"] = contextDoc
	graph["@graph"] = doc

	cont := map[string]interface{}{}
	cont["@context"] = contextDoc
	switch result.(type) {
	case *[]model.DeviceType:
		cont["@type"] = model.SES_ONTOLOGY_DEVICE_TYPE
	case *[]model.Concept:
		cont["@type"] = model.SES_ONTOLOGY_CONCEPT
	case *[]model.Characteristic:
		cont["@type"] = model.SES_ONTOLOGY_CHARACTERISTIC
		if rootElement != "" {
			cont["@id"] = rootElement
		}
	case *[]model.ConceptWithCharacteristics:
		cont["@type"] = model.SES_ONTOLOGY_CONCEPT
		cont["@id"] = rootElement
	case *[]model.TotalCount:
		cont["@type"] = model.SES_ONTOLOGY_COUNT
		cont["@id"] = rootElement
	case *[]model.Aspect:
		cont["@type"] = model.SES_ONTOLOGY_ASPECT
	case *[]model.DeviceClass:
		cont["@type"] = model.SES_ONTOLOGY_DEVICE_CLASS
	case *[]model.Function:
		if strings.HasPrefix(rootElement, model.URN_PREFIX+"controlling-function:") {
			cont["@type"] = model.SES_ONTOLOGY_CONTROLLING_FUNCTION
		}

		if strings.HasPrefix(rootElement, model.URN_PREFIX+"measuring-function:") {
			cont["@type"] = model.SES_ONTOLOGY_MEASURING_FUNCTION
		}
	default:
		debug.PrintStack()
		log.Println("Unknown model type:", reflect.TypeOf(result))
		return errors.New("Unknown model type:" + reflect.TypeOf(result).String())
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

	b, err := json.Marshal(frameDocGraph)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Marshal()", err)
		return err
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		debug.PrintStack()
		log.Println("Error: Unmarshal()", err, string(b))
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
		"id":          "@id",
		"rdf_type":    "@type",
		"name":        model.RDFS_LABEL,
		"description": model.RDFS_COMMENT,
		"protocol_id": map[string]interface{}{
			"@id":   model.SES_ONTOLOGY_HAS_PROTOCOL,
			"@type": "@id",
		},
		"interaction": model.SES_ONTOLOGY_HAS_INTERACTION,
		"device_class_id": map[string]interface{}{
			"@id":   model.SES_ONTOLOGY_HAS_DEVICE_CLASS,
			"@type": "@id",
		},
		"concept_id": map[string]interface{}{
			"@id":   model.SES_ONTOLOGY_HAS_CONCEPT,
			"@type": "@id",
		},
		"services": map[string]interface{}{
			"@id":        model.SES_ONTOLOGY_HAS_SERVICE,
			"@container": "@set",
		},
		"aspect_ids": map[string]interface{}{
			"@id":        model.SES_ONTOLOGY_REFERS_TO,
			"@type":      "@id",
			"@container": "@set",
		},
		"function_ids": map[string]interface{}{
			"@id":        model.SES_ONTOLOGY_EXPOSES_FUNCTION,
			"@type":      "@id",
			"@container": "@set",
		},
		"image": map[string]interface{}{
			"@id":   model.SES_ONTOLOGY_HAS_IMAGE,
			"@type": "@id",
		},
	}
	return contextDoc
}

func getConceptContext() map[string]interface{} {
	contextDoc := map[string]interface{}{
		"id":                     "@id",
		"rdf_type":               "@type",
		"name":                   model.RDFS_LABEL,
		"base_characteristic_id": model.SES_ONTOLOGY_HAS_BASE_CHARACTERISTIC,
		"characteristic_ids": map[string]interface{}{
			"@id":        model.SES_ONTOLOGY_HAS_CHARACTERISTIC,
			"@type":      "@id",
			"@container": "@set",
		},
	}
	return contextDoc
}

func getDeviceClassContext() map[string]interface{} {
	contextDoc := map[string]interface{}{
		"id":       "@id",
		"rdf_type": "@type",
		"name":     model.RDFS_LABEL,
		"image": map[string]interface{}{
			"@id":   model.SES_ONTOLOGY_HAS_IMAGE,
			"@type": "@id",
		},
	}
	return contextDoc
}

func getLocationContext() map[string]interface{} {
	contextDoc := map[string]interface{}{
		"id":          "@id",
		"rdf_type":    "@type",
		"name":        model.RDFS_LABEL,
		"description": model.RDFS_COMMENT,
		"image": map[string]interface{}{
			"@id":   model.SES_ONTOLOGY_HAS_IMAGE,
			"@type": "@id",
		},
	}
	return contextDoc
}

func getAspectContext() map[string]interface{} {
	contextDoc := map[string]interface{}{
		"id":       "@id",
		"rdf_type": "@type",
		"name":     model.RDFS_LABEL,
	}
	return contextDoc
}

func getFunctionContext() map[string]interface{} {
	contextDoc := map[string]interface{}{
		"id":          "@id",
		"rdf_type":    "@type",
		"name":        model.RDFS_LABEL,
		"description": model.RDFS_COMMENT,
		"concept_id": map[string]interface{}{
			"@id":   model.SES_ONTOLOGY_HAS_CONCEPT,
			"@type": "@id",
		},
	}
	return contextDoc
}

func getConceptCharacteristicContext() map[string]interface{} {
	contextDoc := map[string]interface{}{
		"id":                     "@id",
		"rdf_type":               "@type",
		"name":                   model.RDFS_LABEL,
		"base_characteristic_id": model.SES_ONTOLOGY_HAS_BASE_CHARACTERISTIC,
		"characteristics": map[string]interface{}{
			"@id":        model.SES_ONTOLOGY_HAS_CHARACTERISTIC,
			"@container": "@set",
		},
		"type": map[string]interface{}{
			"@id":   model.SES_ONTOLOGY_HAS_VALUE_TYPE,
			"@type": "@id",
		},
		"sub_characteristics": map[string]interface{}{
			"@id":        model.SES_ONTOLOGY_HAS_SUB_CHARACTERISTIC,
			"@container": "@set",
		},
		"value":     model.SES_ONTOLOGY_HAS_VALUE,
		"min_value": model.SES_ONTOLOGY_HAS_MIN_VALUE,
		"max_value": model.SES_ONTOLOGY_HAS_MAX_VALUE,
	}
	return contextDoc
}

func getTotalCountContext() map[string]interface{} {
	contextDoc := map[string]interface{}{
		"id":          "@id",
		"rdf_type":    "@type",
		"total_count": model.SES_ONTOLOGY_TOTAL_COUNT,
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
		"value":     model.SES_ONTOLOGY_HAS_VALUE,
		"min_value": model.SES_ONTOLOGY_HAS_MIN_VALUE,
		"max_value": model.SES_ONTOLOGY_HAS_MAX_VALUE,
	}
	return contextDoc
}
