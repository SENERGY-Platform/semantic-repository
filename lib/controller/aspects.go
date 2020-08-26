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

func (this *Controller) GetAspects() (result []model.Aspect, err error, errCode int) {
	aspects, err := this.db.GetListWithoutSubProperties(model.RDF_TYPE, model.SES_ONTOLOGY_ASPECT)
	if err != nil {
		log.Println("GetAspects ERROR: GetListWithoutSubProperties", err)
		return result, err, http.StatusInternalServerError
	}

	err = this.RdfXmlToModel(aspects, &result)
	if err != nil {
		log.Println("GetAspects ERROR: RdfXmlToModel", err)
		return result, err, http.StatusInternalServerError
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result, nil, http.StatusOK
}

func (this *Controller) GetAspectsWithMeasuringFunction() (result []model.Aspect, err error, errCode int) {
	aspects, err := this.db.GetAspectsWithMeasuringFunction()
	if err != nil {
		log.Println("GetAspects ERROR: GetListWithoutSubProperties", err)
		return result, err, http.StatusInternalServerError
	}

	err = this.RdfXmlToModel(aspects, &result)
	if err != nil {
		log.Println("GetAspects ERROR: RdfXmlToModel", err)
		return result, err, http.StatusInternalServerError
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result, nil, http.StatusOK
}

func (this *Controller) GetAspectsMeasuringFunctions(subject string) (result []model.Function, err error, errCode int) {
	functions, err := this.db.GetAspectsMeasuringFunctions(subject)
	if err != nil {
		log.Println("GetAspectMeasuringFunctions ERROR: GetAspectsMeasuringFunctions", err)
		return result, err, http.StatusInternalServerError
	}

	err = this.RdfXmlToModel(functions, &result)
	if err != nil {
		log.Println("GetAspectMeasuringFunctions ERROR: RdfXmlToModel", err)
		return result, err, http.StatusInternalServerError
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

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
		if !strings.HasPrefix(aspect.Id, model.URN_PREFIX) {
			return errors.New("invalid aspect id"), http.StatusBadRequest
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

func (this *Controller) ValidateAspect(aspect model.Aspect) (error, int) {

	if aspect.Id == "" {
		return errors.New("missing aspect id"), http.StatusBadRequest
	}
	if !strings.HasPrefix(aspect.Id, model.URN_PREFIX) {
		return errors.New("invalid aspect id"), http.StatusBadRequest
	}
	if aspect.Name == "" {
		return errors.New("missing aspect name"), http.StatusBadRequest
	}
	if aspect.RdfType != model.SES_ONTOLOGY_ASPECT {
		return errors.New("wrong aspect type"), http.StatusBadRequest
	}

	return nil, http.StatusOK
}

func (this *Controller) SetAspect(aspect model.Aspect, owner string) (err error) {
	SetAspectRdfType(&aspect)

	err, code := this.ValidateAspect(aspect)
	if err != nil {
		debug.PrintStack()
		log.Println("Error Validation:", err, code)
		return
	}

	err = this.DeleteAspect(aspect.Id)
	if err != nil {
		debug.PrintStack()
		log.Println("Error Delete Aspect:", err, code)
		return
	}

	b, err := json.Marshal(aspect)
	var deviceTypeJsonLd map[string]interface{}
	err = json.Unmarshal(b, &deviceTypeJsonLd)

	deviceTypeJsonLd["@context"] = getAspectContext()

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
		log.Println("Error insert Aspect:", err)
		return err
	}

	return nil
}

/////////////////////////
//		source
/////////////////////////
func SetAspectRdfType(aspect *model.Aspect) {
	aspect.RdfType = model.SES_ONTOLOGY_ASPECT
}

func (this *Controller) DeleteAspect(id string) (err error) {
	if id == "" {
		debug.PrintStack()
		return errors.New("missing aspect id")
	}

	err = this.db.DeleteSubject(id, model.SES_ONTOLOGY_ASPECT)
	if err != nil {
		debug.PrintStack()
		return err
	}
	return nil
}
