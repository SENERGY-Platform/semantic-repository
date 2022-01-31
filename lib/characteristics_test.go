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

func TestCharacteristics(t *testing.T) {
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

	t.Run("testProduceValidCharacteristic1", testProduceValidCharacteristic1(prod))
	t.Run("testProduceValidCharacteristic2", testProduceValidCharacteristic2(prod))
	t.Run("testReadCharacteristic1", testReadCharacteristic1(ctrl))
	t.Run("testReadAllCharacteristic", testReadAllCharacteristic(ctrl))
	t.Run("testDeleteCharacteristic1", testDeleteCharacteristic1(prod))
}

func TestSpecialCaseCharacteristics(t *testing.T) {
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

	t.Run("produce characteristic with special characters", testProduceValidCharacteristic3(prod))
	t.Run("read characteristic with special characters", testReadCharacteristic3(ctrl))
}

func testProduceValidCharacteristic1(producer *producer.Producer) func(t *testing.T) {
	return func(t *testing.T) {
		characteristic := model.Characteristic{}
		characteristic.Id = "urn:ses:infai:characteristic:1d1e1f"
		characteristic.Name = "struct1"
		characteristic.RdfType = "xxx"
		characteristic.DisplayUnit = "°C"
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
		}, {
			Id:       "urn:infai:ses:characteristic:3t3t3t",
			RdfType:  "",
			MinValue: nil,
			MaxValue: nil,
			Value:    nil,
			Type:     model.Structure,
			Name:     "charInnerStructure1",
			SubCharacteristics: []model.Characteristic{
				{
					Id:                 "urn:infai:ses:characteristic:4z4z4z",
					RdfType:            "",
					MinValue:           nil,
					MaxValue:           nil,
					Value:              true,
					Type:               model.Boolean,
					Name:               "charBoolean",
					SubCharacteristics: nil}},
		}}
		err := producer.PublishCharacteristic("urn:ses:infai:concept:1a1a1a", characteristic, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testProduceValidCharacteristic2(producer *producer.Producer) func(t *testing.T) {
	return func(t *testing.T) {
		characteristic := model.Characteristic{}
		characteristic.Id = "urn:ses:infai:characteristic:4711111-20.03.2020"
		characteristic.Name = "bool"
		characteristic.DisplayUnit = "°F"
		characteristic.RdfType = "xxx"
		characteristic.Type = model.Boolean
		err := producer.PublishCharacteristic("urn:ses:infai:concept:1a1a1a", characteristic, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testProduceValidCharacteristic3(producer *producer.Producer) func(t *testing.T) {
	return func(t *testing.T) {
		characteristic := model.Characteristic{}
		characteristic.Id = "urn:ses:infai:characteristic:1111111-30.03.2021"
		characteristic.Name = "µg/m³"
		characteristic.RdfType = "xxx"
		characteristic.Type = model.Boolean
		err := producer.PublishCharacteristic("urn:ses:infai:concept:1a1a1a", characteristic, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testReadCharacteristic1(con *controller.Controller) func(t *testing.T) {
	return func(t *testing.T) {
		characteristic, err, _ := con.GetCharacteristic("urn:ses:infai:characteristic:1d1e1f")
		if err == nil {
			if characteristic.Id != "urn:ses:infai:characteristic:1d1e1f" {
				t.Fatal("wrong id", characteristic.Id)
			}
			if characteristic.Name != "struct1" {
				t.Fatal("wrong name")
			}
			if characteristic.DisplayUnit != "°C" {
				t.Fatal("wrong display unit:", characteristic.DisplayUnit)
			}
			if characteristic.RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
				t.Fatal("wrong rdf_type")
			}
			if characteristic.Type != model.Structure {
				t.Fatal("wrong Type")
			}
			if characteristic.Value != nil {
				t.Fatal("wrong Value")
			}
			if characteristic.MaxValue != nil {
				t.Fatal("wrong MaxValue")
			}
			if characteristic.MinValue != nil {
				t.Fatal("wrong MinValue")
			}
			///////// index -> 0
			if characteristic.SubCharacteristics[0].Id != "urn:infai:ses:characteristic:2r2r2r" {
				t.Fatal("wrong id")
			}
			if characteristic.SubCharacteristics[0].Name != "charFloat" {
				t.Fatal("wrong id")
			}
			if characteristic.SubCharacteristics[0].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
				t.Fatal("wrong id")
			}
			if characteristic.SubCharacteristics[0].Type != model.Float {
				t.Fatal("wrong id")
			}
			if characteristic.SubCharacteristics[0].Value != 2.2 {
				t.Fatal("wrong Value")
			}
			if characteristic.SubCharacteristics[0].MaxValue != 3.0 {
				t.Fatal("wrong MaxValue")
			}
			if characteristic.SubCharacteristics[0].MinValue != -2.0 {
				t.Fatal("wrong MinValue")
			}
			if characteristic.SubCharacteristics[0].SubCharacteristics != nil {
				t.Fatal("wrong SubCharacteristics")
			}
			///////// index -> 1
			if characteristic.SubCharacteristics[1].Id != "urn:infai:ses:characteristic:3t3t3t" {
				t.Fatal("wrong id")
			}
			if characteristic.SubCharacteristics[1].Name != "charInnerStructure1" {
				t.Fatal("wrong id")
			}
			if characteristic.SubCharacteristics[1].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
				t.Fatal("wrong id")
			}
			if characteristic.SubCharacteristics[1].Type != model.Structure {
				t.Fatal("wrong id")
			}
			if characteristic.SubCharacteristics[1].Value != nil {
				t.Fatal("wrong Value")
			}
			if characteristic.SubCharacteristics[1].MaxValue != nil {
				t.Fatal("wrong MaxValue")
			}
			if characteristic.SubCharacteristics[1].MinValue != nil {
				t.Fatal("wrong MinValue")
			}
			///////// index -> 1 -> 0
			if characteristic.SubCharacteristics[1].SubCharacteristics[0].Id != "urn:infai:ses:characteristic:4z4z4z" {
				t.Fatal("wrong id")
			}
			if characteristic.SubCharacteristics[1].SubCharacteristics[0].Name != "charBoolean" {
				t.Fatal("wrong id")
			}
			if characteristic.SubCharacteristics[1].SubCharacteristics[0].RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
				t.Fatal("wrong id")
			}
			if characteristic.SubCharacteristics[1].SubCharacteristics[0].Type != model.Boolean {
				t.Fatal("wrong id")
			}
			if characteristic.SubCharacteristics[1].SubCharacteristics[0].Value != true {
				t.Fatal("wrong Value")
			}
			if characteristic.SubCharacteristics[1].SubCharacteristics[0].MaxValue != nil {
				t.Fatal("wrong MaxValue")
			}
			if characteristic.SubCharacteristics[1].SubCharacteristics[0].MinValue != nil {
				t.Fatal("wrong MinValue")
			}
			t.Log(characteristic)
		} else {
			t.Fatal(err)
		}
	}
}

func testReadCharacteristic3(con *controller.Controller) func(t *testing.T) {
	return func(t *testing.T) {
		characteristic, err, _ := con.GetCharacteristic("urn:ses:infai:characteristic:1111111-30.03.2021")
		if err == nil {
			if characteristic.Id != "urn:ses:infai:characteristic:1111111-30.03.2021" {
				t.Fatal("wrong id", characteristic.Id)
			}
			if characteristic.Name != "µg/m³" {
				t.Fatal("wrong name", characteristic.Name)
			}
		} else {
			t.Fatal(err)
		}
	}
}

func testReadAllCharacteristic(con *controller.Controller) func(t *testing.T) {
	return func(t *testing.T) {
		characteristics, err, _ := con.GetLeafCharacteristics()

		if err == nil {
			if len(characteristics) != 3 {
				t.Fatal("wrong number of elements")
			}
			t.Logf("%+v\n", characteristics)
			t.Log(len(characteristics))
		} else {
			t.Fatal(err)
		}
	}
}

func testDeleteCharacteristic1(producer *producer.Producer) func(t *testing.T) {
	return func(t *testing.T) {
		err := producer.PublishCharacteristicDelete("urn:ses:infai:characteristic:1d1e1f", "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}
	}
}
