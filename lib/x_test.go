package lib

import (
	"encoding/json"
	"fmt"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/database"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/piprate/json-gold/ld"
	"log"
	"os"
	"testing"
)

func TestInsertSparql(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.New(conf)
	success, err := db.InsertData("<urn:infai:ses:category:1> <https://senergy.infai.org/ontology/hasCharacteristic> <urn:infai:ses:characteristic:345> .")
	t.Log(success)
}

func TestConstructSparql(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}

	db, err := database.New(conf)
	body, err := db.ReadData()
	t.Log(string(body))
	//proc := ld.NewJsonLdProcessor()
	//options := ld.NewJsonLdOptions("")
	//// add the processing mode explicitly if you need JSON-LD 1.1 features
	//options.ProcessingMode = ld.JsonLd_1_1
	//options.Format = "application/nquads"
	//doc, err := proc.FromRDF(body, options)
	t.Log(ld.ParseNQuads(string(body)))
	//t.Log(doc, err)
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
		"@context": map[string]interface{}{
			"rdfs":   "http://www.w3.org/2000/01/rdf-schema#",
			"xsd":    "http://www.w3.org/2001/XMLSchema#",
			"schema": "http://schema.org/",
			"name":   "rdfs:label",
			"id":     "@id",
			"type":   "@type",
		},
	}

	function := model.Function{Id: "urn:infai:ses:function:1", Name: "colorFunction", ConceptIds: []string{"1", "2"}, Type: "https://senergy.infai.org/ontology/Function"}

	doc["id"] = function.Id
	doc["name"] = function.Name
	doc["type"] = function.Type
	doc["concept_ids"] = function.ConceptIds

	log.Println(doc)
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
