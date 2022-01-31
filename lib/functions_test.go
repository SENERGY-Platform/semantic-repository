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

func TestFunction(t *testing.T) {
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

	t.Run("testProduceFunctions", testProduceFunctions(prod))
	t.Run("testReadControllingFunction", testReadControllingFunction(ctrl))
	t.Run("testReadMeasuringFunction", testReadMeasuringFunction(ctrl))
	t.Run("testReadFunctions", testReadFunctions(ctrl))
	t.Run("testFunctionDelete", testFunctionDelete(prod))
}

func testProduceFunctions(producer *producer.Producer) func(t *testing.T) {
	return func(t *testing.T) {
		confunction1 := model.Function{}
		confunction1.Id = "urn:infai:ses:controlling-function:333"
		confunction1.Name = "setOnFunction"
		confunction1.Description = "Turn the device on"

		err := producer.PublishFunction(confunction1, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}

		confunction2 := model.Function{}
		confunction2.Id = "urn:infai:ses:controlling-function:2222"
		confunction2.Name = "setOffFunction"
		confunction2.DisplayName = "off-function"
		confunction2.ConceptId = ""

		err = producer.PublishFunction(confunction2, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}

		confunction3 := model.Function{}
		confunction3.Id = "urn:infai:ses:controlling-function:5467567"
		confunction3.Name = "setColorFunction"
		confunction3.ConceptId = "urn:infai:ses:concept:efffsdfd-01a1-4434-9dcc-064b3955000f"

		err = producer.PublishFunction(confunction3, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}

		measfunction1 := model.Function{}
		measfunction1.Id = "urn:infai:ses:measuring-function:23"
		measfunction1.Name = "getOnOffFunction"

		err = producer.PublishFunction(measfunction1, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}

		measfunction2 := model.Function{}
		measfunction2.Id = "urn:infai:ses:measuring-function:321"
		measfunction2.Name = "getTemperatureFunction"
		measfunction2.ConceptId = "urn:infai:ses:concept:efffsdfd-aaaa-bbbb-ccc-0000"

		err = producer.PublishFunction(measfunction2, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}

		measfunction3 := model.Function{}
		measfunction3.Id = "urn:infai:ses:measuring-function:467"
		measfunction3.Name = "getHumidityFunction"
		measfunction3.DisplayName = "hum_display"

		err = producer.PublishFunction(measfunction3, "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testReadControllingFunction(con *controller.Controller) func(t *testing.T) {
	return func(t *testing.T) {
		res, err, code := con.GetFunctionsByType(model.SES_ONTOLOGY_CONTROLLING_FUNCTION)
		if err != nil {
			t.Fatal(res, err, code)
		} else {
			t.Log(res)
		}

		if res[0].Id != "urn:infai:ses:controlling-function:5467567" {
			t.Fatal("error id")
		}
		if res[0].Name != "setColorFunction" {
			t.Fatal("error Name")
		}
		if res[0].ConceptId != "urn:infai:ses:concept:efffsdfd-01a1-4434-9dcc-064b3955000f" {
			t.Fatal("error ConceptId")
		}

		if res[1].Id != "urn:infai:ses:controlling-function:2222" {
			t.Fatal("error id")
		}
		if res[1].Name != "setOffFunction" {
			t.Fatal("error Name")
		}
		if res[1].DisplayName != "off-function" {
			t.Fatal("error Display Name", res[1])
		}
		if res[1].ConceptId != "" {
			t.Fatal("error ConceptId")
		}

		if res[2].Id != "urn:infai:ses:controlling-function:333" {
			t.Fatal("error id")
		}
		if res[2].Name != "setOnFunction" {
			t.Fatal("error Name")
		}
		if res[2].ConceptId != "" {
			t.Fatal("error ConceptId")
		}
		if res[2].Description != "Turn the device on" {
			t.Fatal("error Description")
		}
	}
}

func testReadMeasuringFunction(con *controller.Controller) func(t *testing.T) {
	return func(t *testing.T) {
		res, err, code := con.GetFunctionsByType(model.SES_ONTOLOGY_MEASURING_FUNCTION)
		if err != nil {
			t.Fatal(res, err, code)
		} else {
			t.Log(res)
		}

		if res[0].Id != "urn:infai:ses:measuring-function:467" {
			t.Fatal("error id")
		}
		if res[0].Name != "getHumidityFunction" {
			t.Fatal("error Name")
		}
		if res[0].ConceptId != "" {
			t.Fatal("error ConceptId")
		}

		if res[1].Id != "urn:infai:ses:measuring-function:23" {
			t.Fatal("error id")
		}
		if res[1].Name != "getOnOffFunction" {
			t.Fatal("error Name")
		}
		if res[1].ConceptId != "" {
			t.Fatal("error ConceptId")
		}

		if res[2].Id != "urn:infai:ses:measuring-function:321" {
			t.Fatal("error id")
		}
		if res[2].Name != "getTemperatureFunction" {
			t.Fatal("error Name")
		}
		if res[2].ConceptId != "urn:infai:ses:concept:efffsdfd-aaaa-bbbb-ccc-0000" {
			t.Fatal("error ConceptId")
		}
	}
}

func testReadFunctions(con *controller.Controller) func(t *testing.T) {
	return func(t *testing.T) {
		// test limit offset
		res, err, code := con.GetFunctions(1, 0, "", "")
		if err != nil {
			t.Fatal(res, err, code)
		} else {
			t.Log(res)
		}

		if res.Functions[0].Id != "urn:infai:ses:measuring-function:467" ||
			res.Functions[0].Name != "getHumidityFunction" ||
			res.Functions[0].ConceptId != "" ||
			res.Functions[0].RdfType != model.SES_ONTOLOGY_MEASURING_FUNCTION ||
			res.TotalCount != 6 {
			t.Fatal("error ")
		}

		if res.Functions[0].DisplayName != "hum_display" {
			t.Fatal("unexpected display name:", res.Functions[0].DisplayName)
		}

		res, err, code = con.GetFunctions(1, 1, "", "")
		if err != nil {
			t.Fatal(res, err, code)
		} else {
			t.Log(res)
		}

		if res.Functions[0].Id != "urn:infai:ses:measuring-function:23" ||
			res.Functions[0].Name != "getOnOffFunction" ||
			res.Functions[0].ConceptId != "" ||
			//res.Functions[0].RdfType != model.SES_ONTOLOGY_MEASURING_FUNCTION ||
			res.TotalCount != 6 {
			t.Fatal("error ")
		}

		res, err, code = con.GetFunctions(1, 2, "", "")
		if err != nil {
			t.Fatal(res, err, code)
		} else {
			t.Log(res)
		}

		if res.Functions[0].Id != "urn:infai:ses:measuring-function:321" ||
			res.Functions[0].Name != "getTemperatureFunction" ||
			res.Functions[0].ConceptId != "urn:infai:ses:concept:efffsdfd-aaaa-bbbb-ccc-0000" ||
			res.Functions[0].RdfType != model.SES_ONTOLOGY_MEASURING_FUNCTION ||
			res.TotalCount != 6 {
			t.Fatal("error ", res)
		}

		// test direction
		res, err, code = con.GetFunctions(1, 0, "", "asc")
		if err != nil {
			t.Fatal(res, err, code)
		} else {
			t.Log(res)
		}

		if res.Functions[0].Id != "urn:infai:ses:measuring-function:467" ||
			res.Functions[0].Name != "getHumidityFunction" ||
			res.Functions[0].ConceptId != "" ||
			res.Functions[0].RdfType != model.SES_ONTOLOGY_MEASURING_FUNCTION ||
			res.TotalCount != 6 {
			t.Fatal("error ")
		}

		res, err, code = con.GetFunctions(1, 0, "", "desc")
		if err != nil {
			t.Fatal(res, err, code)
		} else {
			t.Log(res)
		}

		if res.Functions[0].Id != "urn:infai:ses:controlling-function:333" ||
			res.Functions[0].Name != "setOnFunction" ||
			res.Functions[0].Description != "Turn the device on" ||
			res.Functions[0].ConceptId != "" ||
			res.Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION ||
			res.TotalCount != 6 {
			t.Fatal("error ")
		}

		// test search
		res, err, code = con.GetFunctions(1, 0, "Off", "asc")
		if err != nil {
			t.Fatal(res, err, code)
		} else {
			t.Log(res)
		}

		if res.Functions[0].Id != "urn:infai:ses:measuring-function:23" ||
			res.Functions[0].Name != "getOnOffFunction" ||
			res.Functions[0].ConceptId != "" ||
			res.Functions[0].RdfType != model.SES_ONTOLOGY_MEASURING_FUNCTION ||
			res.TotalCount != 2 {
			t.Fatal("error ")
		}

		res, err, code = con.GetFunctions(10, 0, "unc", "desc")
		if err != nil {
			t.Fatal(res, err, code)
		} else {
			t.Log(res)
		}

		if len(res.Functions) != 6 ||
			res.TotalCount != 6 {
			t.Fatal("error ")
		}
	}
}

func testFunctionDelete(producer *producer.Producer) func(t *testing.T) {
	return func(t *testing.T) {
		funcids := [6]string{
			"urn:infai:ses:controlling-function:333",
			"urn:infai:ses:controlling-function:2222",
			"urn:infai:ses:controlling-function:5467567",
			"urn:infai:ses:measuring-function:23",
			"urn:infai:ses:measuring-function:321",
			"urn:infai:ses:measuring-function:467"}

		for _, funcid := range funcids {
			err := producer.PublishFunctionDelete(funcid, "sdfdsfsf")
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
