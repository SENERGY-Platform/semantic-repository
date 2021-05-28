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
	"github.com/SENERGY-Platform/semantic-repository/lib/api"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/controller"
	"github.com/SENERGY-Platform/semantic-repository/lib/database"
	"github.com/SENERGY-Platform/semantic-repository/lib/source/consumer"
	"log"
)

func Start(baseCtx context.Context, conf config.Config) (err error) {
	ctx, cancel := context.WithCancel(baseCtx)
	defer func() {
		if err != nil {
			cancel()
		}
	}()
	db, err := database.New(conf)
	if err != nil {
		log.Println("ERROR: unable to connect to database", err)
		return err
	}

	ctrl, err := controller.New(conf, db)
	if err != nil {
		log.Println("ERROR: unable to start control", err)
		return err
	}

	if !conf.DisableKafkaConsumer {
		err = consumer.Start(ctx, conf, ctrl)
		if err != nil {
			log.Println("ERROR: unable to start source", err)
			return err
		}
	}

	if !conf.DisableHttpApi {
		err = api.Start(conf, ctrl)
		if err != nil {
			log.Println("ERROR: unable to start api", err)
			return err
		}
	}

	return err
}
