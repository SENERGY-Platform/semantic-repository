package lib

import (
	"encoding/json"
	"fmt"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/database"
	"github.com/SENERGY-Platform/semantic-repository/lib/database/sparql/rdf"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/piprate/json-gold/ld"
	"log"
	"os"
	"strings"
	"testing"
)

func TestInsertSparql(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.New(conf)
	success, err := db.InsertData(`<urn:infai:ses:concept:6666> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <https://senergy.infai.org/ontology/Concept> .
<urn:infai:ses:concept:6666> <http://www.w3.org/2000/01/rdf-schema#label> "color" .`)
	t.Log(success)
}

func TestConstructSparql(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}

	db, err := database.New(conf)
	body, err := db.GetConstruct("","", "")
	t.Log(string(body))
}

func TestJsonLd(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	// add the processing mode explicitly if you need JSON-LD 1.1 features
	options.ProcessingMode = ld.JsonLd_1_1
	options.Format = "application/n-quads"

	// this JSON-LD document was taken from http://json-ld.org/test-suite/tests/toRdf-0028-in.jsonld
	doc := map[string]interface{}{
		"id": "@id",
		"type": "@type",
		"name": model.RDFS_LABEL,
		"device_class": model.SES_ONTOLOGY_HAS_DEVICE_CLASS,
		"services": model.SES_ONTOLOGY_HAS_SERVICE,
		"aspects": model.SES_ONTOLOGY_REFERS_TO,
		"functions": model.SES_ONTOLOGY_EXPOSES_FUNCTION,
		"concept_ids": map[string]interface{}{
			"@id": model.SES_ONTOLOGY_HAS_CONCEPT,
			"@type": "@id",
			"@container": "@set",
		},
	}

	function := model.Function{Id: "urn:infai:ses:function:1", Name: "colorFunction", ConceptIds: []string{"1", "2"}, Type: "https://senergy.infai.org/ontology/Function"}

	doc["id"] = function.Id
	doc["name"] = function.Name
	doc["type"] = function.Type
	doc["concept_ids"] = function.ConceptIds

	triples, err := proc.ToRDF(doc, options)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return
	}

	os.Stdout.WriteString(triples.(string))
	db, err := database.New(conf)
	success, err := db.InsertData(triples.(string))
	t.Log(success)
}

func TestFromRDF(t *testing.T) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	// add the processing mode explicitly if you need JSON-LD 1.1 features
	options.ProcessingMode = ld.JsonLd_1_1

	triples := `
<urn:infai:ses:function:1> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <https://senergy.infai.org/ontology/Function> .
<urn:infai:ses:function:1> <http://www.w3.org/2000/01/rdf-schema#label> "colorFunction" .
`
	doc, _ := proc.FromRDF(triples, options)
	t.Log(doc)

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
		log.Println("Errorrrrr", err)
	}
	b, err := json.Marshal(framedDoc["@graph"])
	var function []model.Function

	err = json.Unmarshal(b, &function)
	if err != nil {
		fmt.Println("error:", err)
	}
	t.Log(function)

}

func TestFromXmlToStruct(t *testing.T) {
	xml := `<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:sesame="http://www.openrdf.org/schema/sesame#" xmlns:owl="http://www.w3.org/2002/07/owl#" xmlns:xsd="http://www.w3.org/2001/XMLSchema#" xmlns:fn="http://www.w3.org/2005/xpath-functions#">
<rdf:Description rdf:about="urn:infai:ses:function:1">
<rdf:type rdf:resource="https://senergy.infai.org/ontology/Function"/>
<rdfs:label>colorFunction</rdfs:label>
</rdf:Description>
<rdf:Description rdf:about="urn:infai:ses:function:2">
<rdf:type rdf:resource="https://senergy.infai.org/ontology/Function"/>
<rdfs:label>tempFunction</rdfs:label>
</rdf:Description>
<rdf:Description rdf:about="urn:org.apache.rya/2012/05#rts">
<version xmlns="urn:org.apache.rya/2012/05#">3.0.0</version>
</rdf:Description>
</rdf:RDF>`

	triples, err := rdf.NewTripleDecoder(strings.NewReader(xml), rdf.RDFXML).DecodeAll()
	if err != nil {
		t.Fatal(err)
	}

	turtle := []string{}
	for _, trible := range triples {
		turtle = append(turtle, trible.Serialize(rdf.Turtle))
	}

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	// add the processing mode explicitly if you need JSON-LD 1.1 features
	options.ProcessingMode = ld.JsonLd_1_1

	doc, _ := proc.FromRDF(strings.Join(turtle, ""), options)
	t.Log(doc)

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
		log.Println("Errorrrrr", err)
	}
	b, err := json.Marshal(framedDoc["@graph"])
	var function []model.Function

	err = json.Unmarshal(b, &function)
	if err != nil {
		fmt.Println("error:", err)
	}
	t.Log(function)

}

func TestCompact(t *testing.T) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	// add the processing mode explicitly if you need JSON-LD 1.1 features
	options.ProcessingMode = ld.JsonLd_1_1

	doc := map[string]interface{}{
		"@id": "http://example.org/test#book",
		"http://example.org/vocab#contains": map[string]interface{}{
			"@id": "http://example.org/test#chapter",
		},
		"http://purl.org/dc/elements/1.1/title": "Title",
	}

	context := map[string]interface{}{
		"@context": map[string]interface{}{
			"dc": "http://purl.org/dc/elements/1.1/",
			"ex": "http://example.org/vocab#",
			"ex:contains": map[string]interface{}{
				"@type": "@id",
			},
		},
	}

	compactedDoc, _ := proc.Compact(doc, context, options)
	log.Println(compactedDoc)
}