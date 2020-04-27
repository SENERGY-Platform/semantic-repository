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
	"sort"
)

/////////////////////////
//		api
/////////////////////////

func (this *Controller) GetFunctionsByType(funcType string) (result []model.Function, err error, errCode int) {
	functions, err := this.db.GetListWithoutSubProperties(model.RDF_TYPE, funcType)
	if err != nil {
		log.Println("GetFunctionsByType ERROR: GetListWithoutSubProperties", err)
		return result, err, http.StatusInternalServerError
	}

	err = this.RdfXmlToModel(functions, &result)
	if err != nil {
		log.Println("GetFunctionsByType ERROR: RdfXmlToModel", err)
		return result, err, http.StatusInternalServerError
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result, nil, http.StatusOK

}

func (this *Controller) GetFunctions(limit int, offset int, search string, direction string) (result model.FunctionList, err error, errCode int) {
	functions, totalcount, err := this.db.GetFunctionsWithoutSubPropertiesLimitOffsetSearch(limit, offset, search, direction)
	if err != nil {
		log.Println("GetFunctions ERROR: GetListWithoutSubProperties", err)
		return result, err, http.StatusInternalServerError
	}

	err = this.RdfXmlToModel(functions, &result.Functions)
	if err != nil {
		log.Println("GetFunctions ERROR: RdfXmlToModel", err)
		return result, err, http.StatusInternalServerError
	}

	count := []model.TotalCount{}
	err = this.RdfXmlFrame(totalcount, &count, model.SES_ONTOLOGY_COUNT)
	if err != nil {
		log.Println("GetFunctions ERROR: RdfXmlToModel", err)
		return result, err, http.StatusInternalServerError
	}
	result.TotalCount = count[0].TotalCount

	sort.Slice(result.Functions, func(i, j int) bool {
		if direction == "desc" {
			return result.Functions[i].Name > result.Functions[j].Name
		} else {
			return result.Functions[i].Name < result.Functions[j].Name
		}
	})

	return result, nil, http.StatusOK

}

func (this *Controller) ValidateFunctions(functions []model.Function) (error, int) {
	if (len(functions)) == 0 {
		return errors.New("expect at least one function"), http.StatusBadRequest
	}

	for _, function := range functions {
		if function.Id == "" {
			return errors.New("missing function id"), http.StatusBadRequest
		}
		if function.Name == "" {
			return errors.New("missing function name"), http.StatusBadRequest
		}
		if !(function.RdfType == model.SES_ONTOLOGY_CONTROLLING_FUNCTION || function.RdfType == model.SES_ONTOLOGY_MEASURING_FUNCTION) {
			return errors.New("wrong function type"), http.StatusBadRequest
		}

	}

	return nil, http.StatusOK
}

/////////////////////////
//		source
/////////////////////////
