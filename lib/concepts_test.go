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

func TestProduceValidConcept1(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}

	producer, _ := producer.New(conf)
	concept := model.Concept{}
	concept.Id = "urn:ses:infai:concept:1a1a1a"
	concept.Name = "color"
	concept.RdfType = "xxx"
	concept.Characteristics = []model.Characteristic{{
		Id:      "urn:ses:infai:characteristic:2b2c2d",
		Name:    "structure",
		Type:    model.Structure,
		RdfType: "xxxx",
		SubCharacteristics: []model.Characteristic{
			{
				Id:       "urn:ses:infai:characteristic:3e3e3e",
				Name:     "nameString",
				Value:    "100",
				Type:     model.String,
				RdfType:  "xxxx",
			},
			{
				Id:       "urn:ses:infai:characteristic:4r4r4r",
				Name:     "nameInteger",
				MinValue: 0,
				MaxValue: 255,
				Value:    122,
				Type:     model.Integer,
				RdfType:  "xxxx",
			},
			{
				Id:       "urn:ses:infai:characteristic:5t5t5t",
				Name:     "nameBoolean",
				Value:    true,
				Type:     model.Boolean,
				RdfType:  "xxxx",
			},
			{
				Id:       "urn:ses:infai:characteristic:6z6z6z",
				Name:     "nameFloat",
				MinValue: 0,
				MaxValue: 255,
				Value:    122.22,
				Type:     model.Float,
				RdfType:  "xxxx",
			},
			{
				Id:       "urn:ses:infai:characteristic:7u7u7u",
				Name:     "nameNoValue",
				MinValue: 0,
				MaxValue: 10,
				Type:     model.Float,
				RdfType:  "xxxx",
			},
		},
	},
	}
	producer.PublishConcept(concept, "sdfdsfsf")
}

func TestProduceValidConcept2(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}

	producer, _ := producer.New(conf)
	concept := model.Concept{}
	concept.Id = "urn:ses:infai:concept:2_1a1a1a"
	concept.Name = "color"
	concept.RdfType = "xxx"
	concept.Characteristics = []model.Characteristic{{
		Id:      "urn:ses:infai:characteristic:2_2b2c2d",
		Name:    "structure",
		Type:    model.Structure,
		RdfType: "xxxx",
		SubCharacteristics: []model.Characteristic{
			{
				Id:       "urn:ses:infai:characteristic:2_3e3e3e",
				Name:     "nameString",
				Value:    "100",
				Type:     model.String,
				RdfType:  "xxxx",
			},
			{
				Id:       "urn:ses:infai:characteristic:2_4r4r4r",
				Name:     "nameInteger",
				MinValue: 0,
				MaxValue: 255,
				Value:    122,
				Type:     model.Integer,
				RdfType:  "xxxx",
			},
			{
				Id:       "urn:ses:infai:characteristic:2_5t5t5t",
				Name:     "nameBoolean",
				Value:    true,
				Type:     model.Boolean,
				RdfType:  "xxxx",
			},
			{
				Id:       "urn:ses:infai:characteristic:2_6z6z6z",
				Name:     "nameFloat",
				MinValue: 0,
				MaxValue: 255,
				Value:    122.22,
				Type:     model.Float,
				RdfType:  "xxxx",
			},
			{
				Id:       "urn:ses:infai:characteristic:2_7u7u7u",
				Name:     "nameNoValue",
				MinValue: 0,
				MaxValue: 10,
				Type:     model.Float,
				RdfType:  "xxxx",
			},
		},
	},
	}
	producer.PublishConcept(concept, "sdfdsfsf")
}

func TestProduceValidConcept3SubSubs(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}

	producer, _ := producer.New(conf)
	concept := model.Concept{}
	concept.Id = "urn:ses:infai:concept:3_1"
	concept.Name = "concept3"
	concept.RdfType = "xxx"
	concept.Characteristics = []model.Characteristic{{
		Id:      "urn:ses:infai:characteristic:3_char1",
		Name:    "structure1",
		Type:    model.Structure,
		RdfType: "xxxx",
		SubCharacteristics: []model.Characteristic{
			{
				Id:       "urn:ses:infai:characteristic:3_char2",
				Name:     "structure2",
				Type:     model.Structure,
				RdfType:  "xxxx",
				SubCharacteristics: []model.Characteristic{
					{
						Id:       "urn:ses:infai:characteristic:3_char3",
						Name:     "nameString",
						Type:     model.String,
						RdfType:  "xxxx",
					},
				},
			},
		},
	},
	}
	producer.PublishConcept(concept, "sdfdsfsf")
}


func TestReadConcept1(t *testing.T) {
	err, con, _ := StartUpScript(t)
 	concept, err, _ := con.GetConcept("urn:ses:infai:concept:1a1a1a")

	if err == nil {
		if concept.Id != "urn:ses:infai:concept:1a1a1a" {
			t.Fatal("wrong id")
		}
		if concept.Name != "color" {
			t.Fatal("wrong name")
		}
		if concept.RdfType != model.SES_ONTOLOGY_CONCEPT {
			t.Fatal("wrong rdf_type")
		}
		if concept.Characteristics[0].Id != "urn:ses:infai:characteristic:2b2c2d" {
			t.Fatal("wrong Characteristics id", concept.Characteristics[0].Id)
		}
		if concept.Characteristics[0].Name != "structure" {
			t.Fatal("wrong Characteristics name")
		}
		if concept.Characteristics[0].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
			t.Fatal("wrong Characteristics rdf_type")
		}
		if concept.Characteristics[0].MinValue != nil {
			t.Fatal("wrong Characteristics MinValue", concept.Characteristics[0].MinValue)
		}
		if concept.Characteristics[0].MaxValue != nil {
			t.Fatal("wrong Characteristics MaxValue", concept.Characteristics[0].MaxValue)
		}
		if concept.Characteristics[0].Value != nil {
			t.Fatal("wrong Characteristics Value", concept.Characteristics[0].Value)
		}
		if concept.Characteristics[0].Type != model.Structure {
			t.Fatal("wrong Characteristics Type", concept.Characteristics[0].Type)
		}
		// id = urn:ses:infai:characteristic:3e3e3e
		index := 0
		if concept.Characteristics[0].SubCharacteristics[index].Id != "urn:ses:infai:characteristic:3e3e3e" {
			t.Fatal("wrong SubCharacteristics Id", concept.Characteristics[0].SubCharacteristics[index].Id)
		}
		if concept.Characteristics[0].SubCharacteristics[index].Name != "nameString" {
			t.Fatal("wrong SubCharacteristics Name", concept.Characteristics[0].SubCharacteristics[index].Name)
		}
		if concept.Characteristics[0].SubCharacteristics[index].Type != model.String {
			t.Fatal("wrong SubCharacteristics Type", concept.Characteristics[0].SubCharacteristics[index].Type)
		}
		if concept.Characteristics[0].SubCharacteristics[index].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
			t.Fatal("wrong SubCharacteristics RdfType", concept.Characteristics[0].SubCharacteristics[index].RdfType)
		}
		if concept.Characteristics[0].SubCharacteristics[index].Value != "100" {
			t.Fatal("wrong SubCharacteristics Value", concept.Characteristics[0].SubCharacteristics[index].Value)
		}
		if concept.Characteristics[0].SubCharacteristics[index].MaxValue != nil {
			t.Fatal("wrong SubCharacteristics MaxValue", concept.Characteristics[0].SubCharacteristics[index].MaxValue)
		}
		if concept.Characteristics[0].SubCharacteristics[index].MinValue != nil {
			t.Fatal("wrong SubCharacteristics MinValue", concept.Characteristics[0].SubCharacteristics[index].MinValue)
		}
		if concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics != nil {
			t.Fatal("wrong SubCharacteristics SubCharacteristics", concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics)
		}
		// id = urn:ses:infai:characteristic:4r4r4r
		index = 1
		if concept.Characteristics[0].SubCharacteristics[index].Id != "urn:ses:infai:characteristic:4r4r4r" {
			t.Fatal("wrong SubCharacteristics Id", concept.Characteristics[0].SubCharacteristics[index].Id)
		}
		if concept.Characteristics[0].SubCharacteristics[index].Name != "nameInteger" {
			t.Fatal("wrong SubCharacteristics Name", concept.Characteristics[0].SubCharacteristics[index].Name)
		}
		if concept.Characteristics[0].SubCharacteristics[index].Type != model.Integer {
			t.Fatal("wrong SubCharacteristics Type", concept.Characteristics[0].SubCharacteristics[index].Type)
		}
		if concept.Characteristics[0].SubCharacteristics[index].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
			t.Fatal("wrong SubCharacteristics RdfType", concept.Characteristics[0].SubCharacteristics[index].RdfType)
		}
		if concept.Characteristics[0].SubCharacteristics[index].Value != float64(122) {
			t.Fatal("wrong SubCharacteristics Value", concept.Characteristics[0].SubCharacteristics[index].Value)
		}
		if concept.Characteristics[0].SubCharacteristics[index].MaxValue != float64(255) {
			t.Fatal("wrong SubCharacteristics MaxValue", concept.Characteristics[0].SubCharacteristics[index].MaxValue)
		}
		if concept.Characteristics[0].SubCharacteristics[index].MinValue != float64(0) {
			t.Fatal("wrong SubCharacteristics MinValue", concept.Characteristics[0].SubCharacteristics[index].MinValue)
		}
		if concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics != nil {
			t.Fatal("wrong SubCharacteristics SubCharacteristics", concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics)
		}
		// id = urn:ses:infai:characteristic:6z6z6z
		index = 2
		if concept.Characteristics[0].SubCharacteristics[index].Id != "urn:ses:infai:characteristic:6z6z6z" {
			t.Fatal("wrong SubCharacteristics Id", concept.Characteristics[0].SubCharacteristics[index].Id)
		}
		if concept.Characteristics[0].SubCharacteristics[index].Name != "nameFloat" {
			t.Fatal("wrong SubCharacteristics Name", concept.Characteristics[0].SubCharacteristics[index].Name)
		}
		if concept.Characteristics[0].SubCharacteristics[index].Type != model.Float {
			t.Fatal("wrong SubCharacteristics Type", concept.Characteristics[0].SubCharacteristics[index].Type)
		}
		if concept.Characteristics[0].SubCharacteristics[index].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
			t.Fatal("wrong SubCharacteristics RdfType", concept.Characteristics[0].SubCharacteristics[index].RdfType)
		}
		if concept.Characteristics[0].SubCharacteristics[index].Value != float64(122.22) {
			t.Fatal("wrong SubCharacteristics Value", concept.Characteristics[0].SubCharacteristics[index].Value)
		}
		if concept.Characteristics[0].SubCharacteristics[index].MaxValue != float64(255) {
			t.Fatal("wrong SubCharacteristics MaxValue", concept.Characteristics[0].SubCharacteristics[index].MaxValue)
		}
		if concept.Characteristics[0].SubCharacteristics[index].MinValue != float64(0) {
			t.Fatal("wrong SubCharacteristics MinValue", concept.Characteristics[0].SubCharacteristics[index].MinValue)
		}
		if concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics != nil {
			t.Fatal("wrong SubCharacteristics SubCharacteristics", concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics)
		}
		// id = urn:ses:infai:characteristic:7u7u7u
		index = 3
		if concept.Characteristics[0].SubCharacteristics[index].Id != "urn:ses:infai:characteristic:7u7u7u" {
			t.Fatal("wrong SubCharacteristics Id", concept.Characteristics[0].SubCharacteristics[index].Id)
		}
		if concept.Characteristics[0].SubCharacteristics[index].Name != "nameNoValue" {
			t.Fatal("wrong SubCharacteristics Name", concept.Characteristics[0].SubCharacteristics[index].Name)
		}
		if concept.Characteristics[0].SubCharacteristics[index].Type != model.Float {
			t.Fatal("wrong SubCharacteristics Type", concept.Characteristics[0].SubCharacteristics[index].Type)
		}
		if concept.Characteristics[0].SubCharacteristics[index].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
			t.Fatal("wrong SubCharacteristics RdfType", concept.Characteristics[0].SubCharacteristics[index].RdfType)
		}
		if concept.Characteristics[0].SubCharacteristics[index].Value != nil {
			t.Fatal("wrong SubCharacteristics Value", concept.Characteristics[0].SubCharacteristics[index].Value)
		}
		if concept.Characteristics[0].SubCharacteristics[index].MaxValue != float64(10) {
			t.Fatal("wrong SubCharacteristics MaxValue", concept.Characteristics[0].SubCharacteristics[index].MaxValue)
		}
		if concept.Characteristics[0].SubCharacteristics[index].MinValue != float64(0) {
			t.Fatal("wrong SubCharacteristics MinValue", concept.Characteristics[0].SubCharacteristics[index].MinValue)
		}
		if concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics != nil {
			t.Fatal("wrong SubCharacteristics SubCharacteristics", concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics)
		}
		// id = urn:ses:infai:characteristic:5t5t5t
		index = 4
		if concept.Characteristics[0].SubCharacteristics[index].Id != "urn:ses:infai:characteristic:5t5t5t" {
			t.Fatal("wrong SubCharacteristics Id", concept.Characteristics[0].SubCharacteristics[index].Id)
		}
		if concept.Characteristics[0].SubCharacteristics[index].Name != "nameBoolean" {
			t.Fatal("wrong SubCharacteristics Name", concept.Characteristics[0].SubCharacteristics[index].Name)
		}
		if concept.Characteristics[0].SubCharacteristics[index].Type != model.Boolean {
			t.Fatal("wrong SubCharacteristics Type", concept.Characteristics[0].SubCharacteristics[index].Type)
		}
		if concept.Characteristics[0].SubCharacteristics[index].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
			t.Fatal("wrong SubCharacteristics RdfType", concept.Characteristics[0].SubCharacteristics[index].RdfType)
		}
		if concept.Characteristics[0].SubCharacteristics[index].Value != true {
			t.Fatal("wrong SubCharacteristics Value", concept.Characteristics[0].SubCharacteristics[index].Value)
		}
		if concept.Characteristics[0].SubCharacteristics[index].MaxValue != nil {
			t.Fatal("wrong SubCharacteristics MaxValue", concept.Characteristics[0].SubCharacteristics[index].MaxValue)
		}
		if concept.Characteristics[0].SubCharacteristics[index].MinValue != nil {
			t.Fatal("wrong SubCharacteristics MinValue", concept.Characteristics[0].SubCharacteristics[index].MinValue)
		}
		if concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics != nil {
			t.Fatal("wrong SubCharacteristics SubCharacteristics", concept.Characteristics[0].SubCharacteristics[index].SubCharacteristics)
		}
		t.Log(concept)
	} else {
		t.Fatal(err)
	}
}

func TestReadConcept3(t *testing.T) {
	err, con, _ := StartUpScript(t)
	concept, err, _ := con.GetConcept("urn:ses:infai:concept:3_1")

	if err == nil {
		if concept.Id != "urn:ses:infai:concept:3_1" {
			t.Fatal("wrong id")
		}
		if concept.Name != "concept3" {
			t.Fatal("wrong name")
		}
		if concept.RdfType != model.SES_ONTOLOGY_CONCEPT {
			t.Fatal("wrong rdf_type")
		}
		if concept.Characteristics[0].Id != "urn:ses:infai:characteristic:3_char1" {
			t.Fatal("wrong Characteristics id", concept.Characteristics[0].Id)
		}
		if concept.Characteristics[0].Name != "structure1" {
			t.Fatal("wrong Characteristics name")
		}
		if concept.Characteristics[0].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
			t.Fatal("wrong Characteristics rdf_type")
		}
		if concept.Characteristics[0].MinValue != nil {
			t.Fatal("wrong Characteristics MinValue", concept.Characteristics[0].MinValue)
		}
		if concept.Characteristics[0].MaxValue != nil {
			t.Fatal("wrong Characteristics MaxValue", concept.Characteristics[0].MaxValue)
		}
		if concept.Characteristics[0].Value != nil {
			t.Fatal("wrong Characteristics Value", concept.Characteristics[0].Value)
		}
		if concept.Characteristics[0].Type != model.Structure {
			t.Fatal("wrong Characteristics Type", concept.Characteristics[0].Type)
		}

		if concept.Characteristics[0].SubCharacteristics[0].Id != "urn:ses:infai:characteristic:3_char2" {
			t.Fatal("wrong Characteristics id", concept.Characteristics[0].Id)
		}
		if concept.Characteristics[0].SubCharacteristics[0].Name != "structure2" {
			t.Fatal("wrong Characteristics name")
		}
		if concept.Characteristics[0].SubCharacteristics[0].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
			t.Fatal("wrong Characteristics rdf_type")
		}
		if concept.Characteristics[0].SubCharacteristics[0].MinValue != nil {
			t.Fatal("wrong Characteristics MinValue", concept.Characteristics[0].MinValue)
		}
		if concept.Characteristics[0].SubCharacteristics[0].MaxValue != nil {
			t.Fatal("wrong Characteristics MaxValue", concept.Characteristics[0].MaxValue)
		}
		if concept.Characteristics[0].SubCharacteristics[0].Value != nil {
			t.Fatal("wrong Characteristics Value", concept.Characteristics[0].Value)
		}
		if concept.Characteristics[0].SubCharacteristics[0].Type != model.Structure {
			t.Fatal("wrong Characteristics Type", concept.Characteristics[0].Type)
		}

		if concept.Characteristics[0].SubCharacteristics[0].SubCharacteristics[0].Id != "urn:ses:infai:characteristic:3_char3" {
			t.Fatal("wrong Characteristics id", concept.Characteristics[0].Id)
		}
		if concept.Characteristics[0].SubCharacteristics[0].SubCharacteristics[0].Name != "nameString" {
			t.Fatal("wrong Characteristics name")
		}
		if concept.Characteristics[0].SubCharacteristics[0].SubCharacteristics[0].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
			t.Fatal("wrong Characteristics rdf_type")
		}
		if concept.Characteristics[0].SubCharacteristics[0].SubCharacteristics[0].MinValue != nil {
			t.Fatal("wrong Characteristics MinValue", concept.Characteristics[0].MinValue)
		}
		if concept.Characteristics[0].SubCharacteristics[0].SubCharacteristics[0].MaxValue != nil {
			t.Fatal("wrong Characteristics MaxValue", concept.Characteristics[0].MaxValue)
		}
		if concept.Characteristics[0].SubCharacteristics[0].SubCharacteristics[0].Value != nil {
			t.Fatal("wrong Characteristics Value", concept.Characteristics[0].Value)
		}
		if concept.Characteristics[0].SubCharacteristics[0].SubCharacteristics[0].Type != model.String {
			t.Fatal("wrong Characteristics Type", concept.Characteristics[0].Type)
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

func TestDeleteConcept2(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	err = producer.PublishConceptDelete("urn:ses:infai:concept:2_1a1a1a", "sdfdsfsf")
	if err != nil {
		t.Fatal(err)
	}
}

