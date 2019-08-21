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

func (this *Controller) ValidateCharacteristics(characteristics []model.Characteristic) (error, int) {
	for _, characteristic := range characteristics {
		if characteristic.Id == "" {
			return errors.New("missing characteristic id"), http.StatusBadRequest
		}
		if characteristic.Name == "" {
			return errors.New("missing characteristic name"), http.StatusBadRequest
		}
		if characteristic.RdfType != model.SES_ONTOLOGY_CHARACTERISTIC {
			return errors.New("wrong characteristic rdf_type"), http.StatusBadRequest
		}
		if characteristic.Type != model.String &&
			characteristic.Type != model.Integer &&
			characteristic.Type != model.Float &&
			characteristic.Type != model.Boolean &&
			characteristic.Type != model.List &&
			characteristic.Type != model.Structure {
			return errors.New("wrong characteristic type"), http.StatusBadRequest
		}

		err, code := this.ValidateCharacteristics(characteristic.SubCharacteristics)
		if err != nil {
			return err, code
		}
	}

	return nil, http.StatusOK
}
