/*
 *
 * Copyright 2020 InfAI (CC SES)
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

package testutil

import (
	"context"
	"errors"
	"github.com/ory/dockertest/v3"
	"log"
	"net/http"
	"sync"
)

func RyaContainer(ctx context.Context, wg *sync.WaitGroup, mongoinstance string) (hostPort string, ipAddress string, err error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return "", "", err
	}
	container, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "fgseitsrancher.wifa.intern.uni-leipzig.de:5000/rya",
		Tag:        "dev",
		Env:        []string{"MONGO_INSTANCE=" + mongoinstance},
	})
	if err != nil {
		return "", "", err
	}
	wg.Add(1)
	go func() {
		<-ctx.Done()
		log.Println("DEBUG: remove container " + container.Container.Name)
		container.Close()
		wg.Done()
	}()
	hostPort = container.GetPort("8080/tcp")
	err = pool.Retry(func() error {
		log.Println("try " + container.Container.Name + " connection...")
		resp, err := http.Get("http://localhost:" + hostPort + "/web.rya/sparqlQuery.jsp")
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			return errors.New(resp.Status)
		}
		return nil
	})
	return hostPort, container.Container.NetworkSettings.IPAddress, err
}
