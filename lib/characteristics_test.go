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

package lib

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/SENERGY-Platform/semantic-repository/lib/source/producer"
	"testing"
)

func TestProduceValidCharacteristic1(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}

	producer, _ := producer.New(conf)
	characteristic := model.Characteristic{}
	characteristic.Id = "urn:ses:infai:characteristic:1d1e1f"
	characteristic.Name = "color"
	characteristic.RdfType = "xxx"
	characteristic.Type = model.Structure
	characteristic.SubCharacteristics = []model.Characteristic{{
		Id:                 "urn:infai:ses:characteristic:2r2r2r",
		RdfType:            "",
		MinValue:           -2,
		MaxValue:           3,
		Value:              2.2,
		Type:               model.Float,
		Name:               "charFloat",
		SubCharacteristics: nil,
	}}
	producer.PublishCharacteristic("urn:ses:infai:concept:1a1a1a", characteristic, "sdfdsfsf")
}


//func TestReadConcept1(t *testing.T) {
//	err, con, _ := StartUpScript(t)
//	concept, err, _ := con.GetConcept("urn:ses:infai:concept:1a1a1a")
//
//	if err == nil {
//		if concept.Id != "urn:ses:infai:concept:1a1a1a" {
//			t.Fatal("wrong id")
//		}
//		if concept.Name != "color" {
//			t.Fatal("wrong name")
//		}
//		if concept.RdfType != model.SES_ONTOLOGY_CONCEPT {
//			t.Fatal("wrong rdf_type")
//		}
//		if concept.CharacteristicIds[0] != "urn:ses:infai:characteristic:1a1a1a" {
//			t.Fatal("wrong CharacteristicIds")
//		}
//		t.Log(concept)
//	} else {
//		t.Fatal(err)
//	}
//}



func TestDeleteCharacteristic1(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	err = producer.PublishConceptDelete("urn:ses:infai:characteristic:1d1e1f", "sdfdsfsf")
	if err != nil {
		t.Fatal(err)
	}
}



