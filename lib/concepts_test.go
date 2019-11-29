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

func TestProduceValidConcept1withCharIdAndBaseCharId(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}

	producer, _ := producer.New(conf)
	concept := model.Concept{}
	concept.Id = "urn:ses:infai:concept:1a1a1a1-28-11-2019"
	concept.Name = "color1"
	concept.RdfType = "xxx"
	concept.BaseCharacteristicId = "urn:ses:infai:characteristic:544433333"
	concept.CharacteristicIds = []string{"urn:ses:infai:characteristic:544433333"}
	producer.PublishConcept(concept, "sdfdsfsf")
}

func TestProduceValidConcept1withNoCharId(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}

	producer, _ := producer.New(conf)
	concept := model.Concept{}
	concept.Id = "urn:ses:infai:concept:1a1a1a"
	concept.Name = "color1"
	concept.RdfType = "xxx"
	concept.CharacteristicIds = nil
	producer.PublishConcept(concept, "sdfdsfsf")
}

func TestProduceValidConcept1withCharId(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}

	producer, _ := producer.New(conf)
	concept := model.Concept{}
	concept.Id = "urn:ses:infai:concept:1a1a1a"
	concept.Name = "color1"
	concept.RdfType = "xxx"
	concept.CharacteristicIds = []string{"urn:ses:infai:characteristic:544433333"}
	concept.BaseCharacteristicId = "urn:ses:infai:characteristic:544433333"
	producer.PublishConcept(concept, "sdfdsfsf")
}

func TestProduceValidCharacteristicDependencie(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}

	producer, _ := producer.New(conf)
	characteristic := model.Characteristic{}
	characteristic.Id = "urn:ses:infai:characteristic:544433333"
	characteristic.Name = "struct1"
	characteristic.RdfType = "xxx"
	characteristic.Type = model.Structure
	characteristic.SubCharacteristics = []model.Characteristic{{
		Id:                 "urn:infai:ses:characteristic:123456789999",
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

func TestReadConcept1WithoutSubClass(t *testing.T) {
	err, con, _ := StartUpScript(t)
	concept, err, _ := con.GetConceptWithoutCharacteristics("urn:ses:infai:concept:1a1a1a")

	if err == nil {
		if concept.Id != "urn:ses:infai:concept:1a1a1a" {
			t.Fatal("wrong id")
		}
		if concept.Name != "color1" {
			t.Fatal("wrong name")
		}
		if concept.RdfType != model.SES_ONTOLOGY_CONCEPT {
			t.Fatal("wrong rdf_type")
		}
		if concept.CharacteristicIds[0] != "urn:ses:infai:characteristic:544433333" {
			t.Fatal("wrong CharacteristicIds")
		}
		t.Log(concept)
	} else {
		t.Fatal(err)
	}
}

func TestReadConcept1WithSubClass(t *testing.T) {
	err, con, _ := StartUpScript(t)
	concept, err, _ := con.GetConceptWithCharacteristics("urn:ses:infai:concept:1a1a1a")

	if err == nil {
		if concept.Id != "urn:ses:infai:concept:1a1a1a" {
			t.Fatal("wrong id")
		}
		if concept.Name != "color1" {
			t.Fatal("wrong name")
		}
		if concept.RdfType != model.SES_ONTOLOGY_CONCEPT {
			t.Fatal("wrong rdf_type")
		}
		if concept.BaseCharacteristicId != "urn:ses:infai:characteristic:544433333" {
			t.Fatal("wrong BaseCharacteristicId")
		}
		if concept.Characteristics[0].Id != "urn:ses:infai:characteristic:544433333" {
			t.Fatal("wrong Characteristics")
		}
		if concept.Characteristics[0].SubCharacteristics[0].Id != "urn:infai:ses:characteristic:123456789999" {
			t.Fatal("wrong SubCharacteristics")
		}
		if concept.Characteristics[0].SubCharacteristics[0].Name != "charFloat" {
			t.Fatal("wrong SubCharacteristics")
		}
		t.Log(concept)
	} else {
		t.Fatal(err)
	}
}

func TestDeleteConcept1(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	err = producer.PublishConceptDelete("urn:ses:infai:concept:1a1a1a", "sdfdsfsf")
	if err != nil {
		t.Fatal(err)
	}
}
