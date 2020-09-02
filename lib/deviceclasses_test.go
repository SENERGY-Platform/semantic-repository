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

func TestDeviceClass(t *testing.T) {
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

	t.Run("testProduceDeviceType", testProduceDeviceType(prod))
	t.Run("testDeviceClassRead", testDeviceClassRead(ctrl))
	t.Run("testDeviceClassDelete", testDeviceClassDelete(prod))
}

func testProduceDeviceType(producer *producer.Producer) func(t *testing.T) {
	return func(t *testing.T) {
		producer.PublishDeviceClass(model.DeviceClass{Id: "urn:infai:ses:device-class:eb4a3337-01a1-4434-9dcc-064b3955eeef", Name: "Lamp", Image: "https://i.imgur.com/YHc7cbe.png"}, "sdfdsfsf")
		producer.PublishDeviceClass(model.DeviceClass{Id: "urn:infai:ses:device-class:eb4a3337-01a1-4434-9dcc-123456", Name: "Lamp2"}, "sdfdsfsf")
	}
}

func testDeviceClassRead(con *controller.Controller) func(t *testing.T) {
	return func(t *testing.T) {
		res, err, code := con.GetDeviceClasses()
		if err != nil {
			t.Fatal(res, err, code)
		} else {
			//t.Log(res)
		}
		if res[0].Id != "urn:infai:ses:device-class:eb4a3337-01a1-4434-9dcc-064b3955eeef" {
			t.Fatal("error id", res[0].Id)
		}
		if res[0].Name != "Lamp" {
			t.Fatal("error Name")
		}

		if res[0].RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
			t.Fatal("wrong RdfType")
		}

		if res[0].Image != "https://i.imgur.com/YHc7cbe.png" {
			t.Fatal("wrong Image")
		}

		if res[1].Id != "urn:infai:ses:device-class:eb4a3337-01a1-4434-9dcc-123456" {
			t.Fatal("error id", res[0].Id)
		}
		if res[1].Name != "Lamp2" {
			t.Fatal("error Name")
		}

		if res[1].RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
			t.Fatal("wrong RdfType")
		}

		if res[1].Image != "" {
			t.Fatal("wrong Image")
		}
	}
}

func testDeviceClassDelete(producer *producer.Producer) func(t *testing.T) {
	return func(t *testing.T) {
		err := producer.PublishDeviceClassDelete("urn:infai:ses:device-class:eb4a3337-01a1-4434-9dcc-064b3955eeef", "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}

		err = producer.PublishDeviceClassDelete("urn:infai:ses:device-class:eb4a3337-01a1-4434-9dcc-123456", "sdfdsfsf")
		if err != nil {
			t.Fatal(err)
		}
	}
}
