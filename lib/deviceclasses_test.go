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

func TestProduceDeviceType(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, err := producer.New(conf)
	if err != nil {
		t.Fatal(err)
	}
	deviceclass := model.DeviceClass{}
	deviceclass.Id = "urn:infai:ses:device-class:eb4a3337-01a1-4434-9dcc-064b3955eeef"
	deviceclass.Name = "Lamp"

	producer.PublishDeviceClass(deviceclass, "sdfdsfsf")

}

func TestDeviceClassRead(t *testing.T) {
	err, con, _ := StartUpScript(t)

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

}

func TestDeviceClassDelete(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, err := producer.New(conf)
	if err != nil {
		t.Fatal(err)
	}
	err = producer.PublishDeviceClassDelete("urn:infai:ses:device-class:eb4a3337-01a1-4434-9dcc-064b3955eeef", "sdfdsfsf")
	if err != nil {
		t.Fatal(err)
	}

}
