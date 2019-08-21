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

	//zero := float64(0)
	//two55 := float64(255)


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
