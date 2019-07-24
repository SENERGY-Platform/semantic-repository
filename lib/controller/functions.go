/*
 * Copyright 2019 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package controller

import (
	"encoding/json"
	"github.com/SENERGY-Platform/semantic-repository/lib/database/sparql/rdf"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/piprate/json-gold/ld"
	"net/http"
	"strings"
)

/////////////////////////
//		api
/////////////////////////

func (this *Controller) GetFunctions(funcType string) (result []model.Function, err error, errCode int) {
	functions, err := this.db.GetConstruct("", model.RDF_TYPE, funcType)
	if err != nil {
		return result, err, http.StatusInternalServerError
	}

	triples, err := rdf.NewTripleDecoder(strings.NewReader(string(functions)), rdf.RDFXML).DecodeAll()
	if err != nil {
		return result, err, http.StatusInternalServerError
	}

	turtle := []string{}
	for _, triple := range triples {
		turtle = append(turtle, triple.Serialize(rdf.Turtle))
	}

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.ProcessingMode = ld.JsonLd_1_1

	doc, _ := proc.FromRDF(strings.Join(turtle, ""), options)

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
		return result, err, http.StatusInternalServerError
	}
	b, err := json.Marshal(framedDoc["@graph"])
	var function []model.Function
	err = json.Unmarshal(b, &function)
	if err != nil {
		return result, err, http.StatusInternalServerError
	}

	return function, nil, http.StatusOK
}


/////////////////////////
//		source
/////////////////////////

