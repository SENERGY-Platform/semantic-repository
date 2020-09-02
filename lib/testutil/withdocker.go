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
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"net"
	"sync"
)

func GetDockerEnv(ctx context.Context, wg *sync.WaitGroup, conf *config.Config) (err error) {
	_, ip, err := MongoContainer(ctx, wg)
	if err != nil {
		return err
	}

	port, _, err := RyaContainer(ctx, wg, ip)
	if err != nil {
		return err
	}

	conf.RyaUrl = "http://localhost:" + port
	return
}

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}
