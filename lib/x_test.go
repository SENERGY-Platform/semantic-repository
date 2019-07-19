package lib

import (
	"fmt"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"io/ioutil"
	"net/http"
	"net/url"
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

