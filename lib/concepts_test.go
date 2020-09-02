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
	"context"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/controller"
	"github.com/SENERGY-Platform/semantic-repository/lib/database"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/SENERGY-Platform/semantic-repository/lib/testutil"
	"github.com/SENERGY-Platform/semantic-repository/lib/testutil/producer"
	"sync"
	"testing"
)

func TestConcepts(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	defer wg.Wait()
	defer cancel()
	err = testutil.GetDockerEnv(ctx, &wg, &conf)
	if err != nil {
		t.Error(err)
		return
	}
	db, err := database.New(conf)
	if err != nil {
		t.Error(err)
		return
	}

	ctrl, err := controller.New(conf, db)
	if err != nil {
		t.Error(err)
		return
	}

	prod, err := testutil.StartSourceMock(conf, ctrl)
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("testProduceValidConcept1withCharIdAndBaseCharId", testProduceValidConcept1withCharIdAndBaseCharId(prod))
	t.Run("testProduceValidConcept1withNoCharId", testProduceValidConcept1withNoCharId(prod))
	t.Run("testProduceValidConcept1withCharId", testProduceValidConcept1withCharId(prod))
	t.Run("testProduceValidCharacteristicDependencie", testProduceValidCharacteristicDependencie(prod))
	t.Run("testReadConcept1WithoutSubClass", testReadConcept1WithoutSubClass(ctrl))
	t.Run("testReadConcept1WithSubClass", testReadConcept1WithSubClass(ctrl))
	t.Run("testDeleteConcept1", testDeleteConcept1(prod))
}

func testProduceValidConcept1withCharIdAndBaseCharId(producer *producer.Producer) func(t *testing.T) {
	return func(t *testing.T) {
		concept := model.Concept{}
		concept.Id = "urn:ses:infai:concept:1a1a1a1-28-11-2019"
		concept.Name = "color1"
		concept.RdfType = "xxx"
		concept.BaseCharacteristicId = "urn:ses:infai:characteristic:544433333"
		concept.CharacteristicIds = []string{"urn:ses:infai:characteristic:544433333"}
		err := producer.PublishConcept(concept, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testProduceValidConcept1withNoCharId(producer *producer.Producer) func(t *testing.T) {
	return func(t *testing.T) {
		concept := model.Concept{}
		concept.Id = "urn:ses:infai:concept:1a1a1a"
		concept.Name = "color1"
		concept.RdfType = "xxx"
		concept.CharacteristicIds = nil
		err := producer.PublishConcept(concept, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testProduceValidConcept1withCharId(producer *producer.Producer) func(t *testing.T) {
	return func(t *testing.T) {
		concept := model.Concept{}
		concept.Id = "urn:ses:infai:concept:1a1a1a"
		concept.Name = "color1"
		concept.RdfType = "xxx"
		concept.CharacteristicIds = []string{"urn:ses:infai:characteristic:544433333"}
		concept.BaseCharacteristicId = "urn:ses:infai:characteristic:544433333"
		err := producer.PublishConcept(concept, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testProduceValidCharacteristicDependencie(producer *producer.Producer) func(t *testing.T) {
	return func(t *testing.T) {
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
		err := producer.PublishCharacteristic("urn:ses:infai:concept:1a1a1a", characteristic, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testReadConcept1WithoutSubClass(con *controller.Controller) func(t *testing.T) {
	return func(t *testing.T) {
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
}

func testReadConcept1WithSubClass(con *controller.Controller) func(t *testing.T) {
	return func(t *testing.T) {
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
}

func testDeleteConcept1(producer *producer.Producer) func(t *testing.T) {
	return func(t *testing.T) {
		err := producer.PublishConceptDelete("urn:ses:infai:concept:1a1a1a", "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}
	}
}
