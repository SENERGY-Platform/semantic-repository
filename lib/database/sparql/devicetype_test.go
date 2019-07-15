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

package sparql

import (
	"context"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/ory/dockertest"
	"testing"
	"time"
)

func TestMongoDeviceType(t *testing.T) {

	conf, err := config.Load("../../../config.json")
	if err != nil {
		t.Fatal(err)
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatal("Could not connect to docker: ", err)
	}
	closer, port, _, err := SparqlTestServer(pool)
	if err != nil {
		t.Fatal(err, port)
	}
	if true {
		defer closer()
	}
	//TODO: use port to set your config: conf.MongoUrl = "mongodb://localhost:" + port
	m, err := New(conf)
	if err != nil {
		t.Fatal(err)
	}

	//simple test
	//TODO: do more tests
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	_, exists, err := m.GetDeviceType(ctx, "does_not_exist")
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("device type should not exist")
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = m.SetDeviceType(ctx, model.DeviceType{
		Id:       "foobar1",
		Name:     "foo1",
		Services: []model.Service{},
	})
	if err != nil {
		t.Fatal(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	device, exists, err := m.GetDeviceType(ctx, "foobar1")
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("device should exist")
	}
	if device.Id != "foobar1" || device.Name != "foo1" {
		t.Fatal("unexpected result", device)
	}

}
