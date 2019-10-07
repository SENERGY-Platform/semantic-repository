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
	"testing"
)

func TestReadAspect(t *testing.T) {
	err, con, db := StartUpScript(t)
	/// Aspect Lightning
	err = db.InsertData(
		`<urn:infai:ses:aspect:4444> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <https://senergy.infai.org/ontology/Aspect> .
<urn:infai:ses:aspect:4444> <http://www.w3.org/2000/01/rdf-schema#label> "Lightning" .`)
	if err != nil {
		t.Fatal(err)
	}
	/// Aspect Air
	err = db.InsertData(
		`<urn:infai:ses:aspect:2222> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <https://senergy.infai.org/ontology/Aspect> .
<urn:infai:ses:aspect:2222> <http://www.w3.org/2000/01/rdf-schema#label> "Air" .`)
	if err != nil {
		t.Fatal(err)
	}
	/// Aspect Connectivity
	err = db.InsertData(
		`<urn:infai:ses:aspect:4545> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <https://senergy.infai.org/ontology/Aspect> .
<urn:infai:ses:aspect:4545> <http://www.w3.org/2000/01/rdf-schema#label> "Connectivity" .`)
	if err != nil {
		t.Fatal(err)
	}
	res, err, code := con.GetAspects()
	if err != nil {
		t.Fatal(res, err, code)
	} else {
		t.Log(res)
	}

	if res[0].Id != "urn:infai:ses:aspect:2222" {
		t.Fatal("error id")
	}
	if res[0].Name != "Air" {
		t.Fatal("error Name")
	}

	if res[1].Id != "urn:infai:ses:aspect:4545" {
		t.Fatal("error id")
	}
	if res[1].Name != "Connectivity" {
		t.Fatal("error Name")
	}

	if res[2].Id != "urn:infai:ses:aspect:4444" {
		t.Fatal("error id")
	}
	if res[2].Name != "Lightning" {
		t.Fatal("error Name")
	}

}
