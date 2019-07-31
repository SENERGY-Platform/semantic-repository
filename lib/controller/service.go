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

package controller

import (
	"errors"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"net/http"
)

func (this *Controller) ValidateService(services []model.Service) (error, int) {
	if (len(services)) == 0 {
		return errors.New("expect at least one service"), http.StatusBadRequest
	}

	for _, service := range services {
		if service.Id == "" {
			return errors.New("missing service id"), http.StatusBadRequest
		}
		if service.LocalId == "" {
			return errors.New("missing service local id"), http.StatusBadRequest
		}
		if service.Name == "" {
			return errors.New("missing service name"), http.StatusBadRequest
		}

		if service.ProtocolId == "" {
			return errors.New("missing service protocol id"), http.StatusBadRequest
		}

		if service.RdfType != model.SES_ONTOLOGY_SERVICE {
			return errors.New("wrong service type"), http.StatusBadRequest
		}

		err, code := this.ValidateAspects(service.Aspects)
		if err != nil {
			return err, code
		}

		err, code = this.ValidateFunctions(service.Functions)
		if err != nil {
			return err, code
		}
	}

	return nil, http.StatusOK
}
