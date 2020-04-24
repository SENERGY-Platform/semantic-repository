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

package lib

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"testing"
)

func TestReadControllingFunction(t *testing.T) {
	err, con, db := StartUpScript(t)
	err = db.InsertData(
		`<urn:infai:ses:controlling-function:333> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <` + model.SES_ONTOLOGY_CONTROLLING_FUNCTION + `> .
<urn:infai:ses:controlling-function:333> <http://www.w3.org/2000/01/rdf-schema#label> "onFunction" .`)
	if err != nil {
		t.Fatal(err)
	}
	err = db.InsertData(
		`<urn:infai:ses:controlling-function:2222> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <` + model.SES_ONTOLOGY_CONTROLLING_FUNCTION + `> .
<urn:infai:ses:controlling-function:2222> <http://www.w3.org/2000/01/rdf-schema#label> "offFunction" .`)
	if err != nil {
		t.Fatal(err)
	}
	err = db.InsertData(
		`<urn:infai:ses:controlling-function:5467567> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <` + model.SES_ONTOLOGY_CONTROLLING_FUNCTION + `> .
<urn:infai:ses:controlling-function:5467567> <http://www.w3.org/2000/01/rdf-schema#label> "colorFunction" .`)
	if err != nil {
		t.Fatal(err)
	}
	err = db.InsertData(
		`<urn:infai:ses:controlling-function:5467567> <https://senergy.infai.org/ontology/hasConcept> <urn:infai:ses:concept:12213s> .`)
	if err != nil {
		t.Fatal(err)
	}
	res, err, code := con.GetFunctionsByType(model.SES_ONTOLOGY_CONTROLLING_FUNCTION)
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}

	if res[0].Id != "urn:infai:ses:controlling-function:5467567" {
		t.Fatal("error id")
	}
	if res[0].Name != "colorFunction" {
		t.Fatal("error Name")
	}
	if res[0].ConceptId != "urn:infai:ses:concept:12213s" {
		t.Fatal("error ConceptId")
	}

	if res[1].Id != "urn:infai:ses:controlling-function:2222" {
		t.Fatal("error id")
	}
	if res[1].Name != "offFunction" {
		t.Fatal("error Name")
	}

	if res[2].Id != "urn:infai:ses:controlling-function:333" {
		t.Fatal("error id")
	}
	if res[2].Name != "onFunction" {
		t.Fatal("error Name")
	}
}

func TestReadMeasuringFunction(t *testing.T) {
	err, con, db := StartUpScript(t)
	err = db.InsertData(
		`<urn:infai:ses:measuring-function:23> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <` + model.SES_ONTOLOGY_MEASURING_FUNCTION + `> .
<urn:infai:ses:measuring-function:23> <http://www.w3.org/2000/01/rdf-schema#label> "aaa" .`)
	if err != nil {
		t.Fatal(err)
	}
	err = db.InsertData(
		`<urn:infai:ses:measuring-function:321> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <` + model.SES_ONTOLOGY_MEASURING_FUNCTION + `> .
<urn:infai:ses:measuring-function:321> <http://www.w3.org/2000/01/rdf-schema#label> "zzz" .`)
	if err != nil {
		t.Fatal(err)
	}
	err = db.InsertData(
		`<urn:infai:ses:measuring-function:467> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <` + model.SES_ONTOLOGY_MEASURING_FUNCTION + `> .
<urn:infai:ses:measuring-function:467> <http://www.w3.org/2000/01/rdf-schema#label> "bbb" .`)
	if err != nil {
		t.Fatal(err)
	}
	res, err, code := con.GetFunctionsByType(model.SES_ONTOLOGY_MEASURING_FUNCTION)
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}

	if res[0].Id != "urn:infai:ses:measuring-function:23" {
		t.Fatal("error id")
	}
	if res[0].Name != "aaa" {
		t.Fatal("error Name")
	}

	if res[1].Id != "urn:infai:ses:measuring-function:467" {
		t.Fatal("error id")
	}
	if res[1].Name != "bbb" {
		t.Fatal("error Name")
	}

	if res[2].Id != "urn:infai:ses:measuring-function:321" {
		t.Fatal("error id")
	}
	if res[2].Name != "zzz" {
		t.Fatal("error Name")
	}
}

func TestReadFunctions(t *testing.T) {
	err, con, _ := StartUpScript(t)

	if err != nil {
		t.Fatal(err)
	}
	// test limit offset
	res, err, code := con.GetFunctions(1, 0, "", "")
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}

	if res[0].Id != "urn:infai:ses:measuring-function:23" ||
		res[0].Name != "aaa" ||
		res[0].ConceptId != "" ||
		res[0].RdfType != model.SES_ONTOLOGY_MEASURING_FUNCTION {
		t.Fatal("error ")
	}

	res, err, code = con.GetFunctions(1, 1, "", "")
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}

	if res[0].Id != "urn:infai:ses:measuring-function:467" ||
		res[0].Name != "bbb" ||
		res[0].ConceptId != "" ||
		res[0].RdfType != model.SES_ONTOLOGY_MEASURING_FUNCTION {
		t.Fatal("error ")
	}

	res, err, code = con.GetFunctions(1, 2, "", "")
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}

	if res[0].Id != "urn:infai:ses:controlling-function:5467567" ||
		res[0].Name != "colorFunction" ||
		res[0].ConceptId != "urn:infai:ses:concept:12213s" ||
		res[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error ")
	}

	// test direction
	res, err, code = con.GetFunctions(1, 0, "", "asc")
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}

	if res[0].Id != "urn:infai:ses:measuring-function:23" ||
		res[0].Name != "aaa" ||
		res[0].ConceptId != "" ||
		res[0].RdfType != model.SES_ONTOLOGY_MEASURING_FUNCTION {
		t.Fatal("error ")
	}

	res, err, code = con.GetFunctions(1, 0, "", "desc")
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}

	if res[0].Id != "urn:infai:ses:measuring-function:321" ||
		res[0].Name != "zzz" ||
		res[0].ConceptId != "" ||
		res[0].RdfType != model.SES_ONTOLOGY_MEASURING_FUNCTION {
		t.Fatal("error ")
	}

	// test search
	res, err, code = con.GetFunctions(1, 0, "aaa", "asc")
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}

	if res[0].Id != "urn:infai:ses:measuring-function:23" ||
		res[0].Name != "aaa" ||
		res[0].ConceptId != "" ||
		res[0].RdfType != model.SES_ONTOLOGY_MEASURING_FUNCTION {
		t.Fatal("error ")
	}

	res, err, code = con.GetFunctions(10, 0, "unc", "desc")
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}

	if len(res) != 3 {
		t.Fatal("error ")
	}

}
