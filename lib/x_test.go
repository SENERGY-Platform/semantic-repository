package lib

import (
	"fmt"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/piprate/json-gold/ld"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestInsertSparql(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	query := url.QueryEscape("Insert data {<urn:infai:ses:category:1> <https://senergy.infai.org/ontology/hasCharacteristic> <urn:infai:ses:characteristic:1> .}")
	resp, err := http.Get(conf.RyaUrl + "/web.rya/queryrdf?query.infer=&query.auth=&conf.cv=&query.resultformat=json&query=" + query)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	//var result interface{}
	body, err := ioutil.ReadAll(resp.Body)
	//err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}

func TestConstructSparql(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	query := url.QueryEscape("construct { ?s ?p ?o.} where { ?s ?p ?o. }")
	resp, err := http.Get(conf.RyaUrl + "/web.rya/queryrdf?query.infer=&query.auth=&conf.cv=&query.resultformat=json&query=" + query)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	//var result interface{}
	body, err := ioutil.ReadAll(resp.Body)
	//err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}

func TestJsonLd(t *testing.T) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	// add the processing mode explicitly if you need JSON-LD 1.1 features
	options.ProcessingMode = ld.JsonLd_1_1
	options.Format = "application/n-quads"

	// this JSON-LD document was taken from http://json-ld.org/test-suite/tests/toRdf-0028-in.jsonld
	doc := map[string]interface{}{
		"@context": map[string]interface{}{
			"rdfs":"http://www.w3.org/2000/01/rdf-schema#",
			"name":"rdfs:label",
			"id": "@id",
			"type": "@type",
		},
		"id":"urn:infai:ses:category:1",
		"name":"color",
		"type":"https://senergy.infai.org/ontology/Category",
	}
	triples, err := proc.ToRDF(doc, options)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return
	}

	os.Stdout.WriteString(triples.(string))
}

