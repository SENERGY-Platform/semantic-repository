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
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

/////////////////////////
//		api
/////////////////////////

func (this *Controller) GetAspects() (result []model.Aspect, err error, errCode int) {
	aspects, err := this.db.GetConstructWithoutProperties("", model.RDF_TYPE, model.SES_ONTOLOGY_ASPECT)
	if err != nil {
		log.Println("GetAspects ERROR: GetConstructWithoutProperties", err)
		return result, err, http.StatusInternalServerError
	}

	err = this.RdfXmlToModel(aspects, &result)
	if err != nil {
		log.Println("GetAspects ERROR: RdfXmlToModel", err)
		return result, err, http.StatusInternalServerError
	}

	return result, nil, http.StatusOK
}

func (this *Controller) ValidateAspects(aspects []model.Aspect) (error, int) {
	if (len(aspects)) == 0 {
		return errors.New("expect at least one aspect"), http.StatusBadRequest
	}

	for _, aspect := range aspects {
		if aspect.Id == "" {
			return errors.New("missing aspect id"), http.StatusBadRequest
		}
		if aspect.Name == "" {
			return errors.New("missing aspect name"), http.StatusBadRequest
		}
		if aspect.RdfType != model.SES_ONTOLOGY_ASPECT {
			return errors.New("wrong aspect type"), http.StatusBadRequest
		}
	}

	return nil, http.StatusOK
}

/////////////////////////
//		source
/////////////////////////
