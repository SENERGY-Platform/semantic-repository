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
	"encoding/json"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/piprate/json-gold/ld"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"runtime/debug"
	"sort"
	"strings"
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

func (this *Controller) GetFunction(id string) (result model.Function, err error, errCode int) {
	rdftype := ""
	if strings.HasPrefix(id, model.URN_PREFIX+"controlling-function:") {
		rdftype = model.SES_ONTOLOGY_CONTROLLING_FUNCTION
	}

	if strings.HasPrefix(id, model.URN_PREFIX+"measuring-function:") {
		rdftype = model.SES_ONTOLOGY_MEASURING_FUNCTION
	}
	function, err := this.db.GetSubject(id, rdftype)
	if err != nil {
		log.Println("GetFunction ERROR: GetSubject", err)
		return result, err, http.StatusInternalServerError
	}

	res := []model.Function{}
	err = this.RdfXmlToModel(function, &res)
	if err != nil {
		log.Println("GetFunction ERROR: RdfXmlToModel", err)
		return result, err, http.StatusInternalServerError
	}

	if len(res) == 0 {
		return result, errors.New("not found"), http.StatusNotFound
	}

	return res[0], nil, http.StatusOK
}

func (this *Controller) ValidateFunctions(functions []model.Function) (error, int) {
	if (len(functions)) == 0 {
		return errors.New("expect at least one function"), http.StatusBadRequest
	}

	for _, function := range functions {
		if function.Id == "" {
			return errors.New("missing function id"), http.StatusBadRequest
		}
		if !strings.HasPrefix(function.Id, model.URN_PREFIX) {
			return errors.New("invalid function id"), http.StatusBadRequest
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

func (this *Controller) ValidateFunction(function model.Function) (error, int) {

	if function.Id == "" {
		return errors.New("missing function id"), http.StatusBadRequest
	}
	if !strings.HasPrefix(function.Id, model.URN_PREFIX) {
		return errors.New("invalid function id"), http.StatusBadRequest
	}
	if function.Name == "" {
		return errors.New("missing function name"), http.StatusBadRequest
	}
	if !(function.RdfType == model.SES_ONTOLOGY_CONTROLLING_FUNCTION || function.RdfType == model.SES_ONTOLOGY_MEASURING_FUNCTION) {
		return errors.New("wrong function type"), http.StatusBadRequest
	}

	return nil, http.StatusOK
}

func (this *Controller) SetFunction(function model.Function, owner string) (err error) {
	SetFunctionRdfType(&function)

	err, code := this.ValidateFunction(function)
	if err != nil {
		debug.PrintStack()
		log.Println("Error Validation:", err, code)
		return
	}

	err = this.DeleteFunction(function.Id)
	if err != nil {
		debug.PrintStack()
		log.Println("Error Delete Functions:", err, code)
		return
	}

	b, err := json.Marshal(function)
	var deviceTypeJsonLd map[string]interface{}
	err = json.Unmarshal(b, &deviceTypeJsonLd)

	deviceTypeJsonLd["@context"] = getFunctionContext()

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.ProcessingMode = ld.JsonLd_1_1
	options.Format = "application/n-quads"

	triples, err := proc.ToRDF(deviceTypeJsonLd, options)
	if err != nil {
		debug.PrintStack()
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return err
	}

	err = this.db.InsertData(triples.(string))
	if err != nil {
		debug.PrintStack()
		log.Println("Error insert Functions:", err)
		return err
	}

	return nil
}

/////////////////////////
//		source
/////////////////////////
func SetFunctionRdfType(function *model.Function) {
	if strings.HasPrefix(function.Id, model.URN_PREFIX+"controlling-function:") {
		function.RdfType = model.SES_ONTOLOGY_CONTROLLING_FUNCTION
	}

	if strings.HasPrefix(function.Id, model.URN_PREFIX+"measuring-function:") {
		function.RdfType = model.SES_ONTOLOGY_MEASURING_FUNCTION
	}

}

func (this *Controller) DeleteFunction(id string) (err error) {
	if id == "" {
		debug.PrintStack()
		return errors.New("missing functions id")
	}

	rdftype := ""
	if strings.HasPrefix(id, model.URN_PREFIX+"controlling-function:") {
		rdftype = model.SES_ONTOLOGY_CONTROLLING_FUNCTION
	}

	if strings.HasPrefix(id, model.URN_PREFIX+"measuring-function:") {
		rdftype = model.SES_ONTOLOGY_MEASURING_FUNCTION
	}

	err = this.db.DeleteSubject(id, rdftype)
	if err != nil {
		debug.PrintStack()
		return err
	}
	return nil
}
