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
	"log"
	"net/http"
)

/////////////////////////
//		api
/////////////////////////

func (this *Controller) GetAspects() (result []model.Aspect, err error, errCode int) {
	deviceClasses, err := this.db.GetConstruct("", model.RDF_TYPE, model.SES_ONTOLOGY_ASPECT)
	if err != nil {
		log.Println("GetAspects ERROR: GetConstruct", err)
		return result, err, http.StatusInternalServerError
	}

	err = this.ByteToModel(deviceClasses, &result)
	if err != nil {
		log.Println("GetAspects ERROR: ByteToModel", err)
		return result, err, http.StatusInternalServerError
	}

	return result, nil, http.StatusOK
}

/////////////////////////
//		source
/////////////////////////

