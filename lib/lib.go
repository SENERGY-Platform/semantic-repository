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
	"github.com/SENERGY-Platform/semantic-repository/lib/api"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/controller"
	"github.com/SENERGY-Platform/semantic-repository/lib/database"
	"github.com/SENERGY-Platform/semantic-repository/lib/source/consumer"
	"log"
)

func Start(conf config.Config) (stop func(), err error) {
	db, err := database.New(conf)
	if err != nil {
		log.Println("ERROR: unable to connect to database", err)
		return stop, err
	}

	ctrl, err := controller.New(conf, db)
	if err != nil {
		log.Println("ERROR: unable to start control", err)
		return stop, err
	}

	sourceStop, err := consumer.Start(conf, ctrl)
	if err != nil {
		log.Println("ERROR: unable to start source", err)
		return stop, err
	}

	err = api.Start(conf, ctrl)
	if err != nil {
		sourceStop()
		log.Println("ERROR: unable to start api", err)
		return stop, err
	}

	return func() {
		sourceStop()
	}, err
}
