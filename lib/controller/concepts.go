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
	"errors"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/piprate/json-gold/ld"
	"log"
	"net/http"
	"runtime/debug"
)

/////////////////////////
//		api
/////////////////////////

//func (this *Controller) GetDeviceType(subject string) (result model.DeviceType, err error, errCode int) {
//	deviceType, err := this.db.GetConstructWithProperties(subject)
//	if err != nil {
//		log.Println("GetDeviceClasses ERROR: GetConstructWithoutProperties", err)
//		return result, err, http.StatusInternalServerError
//	}
//
//	err = this.RdfXmlToSingleResult(deviceType, &result)
//	if err != nil {
//		log.Println("GetDeviceClasses ERROR: RdfXmlToModel", err)
//		return result, err, http.StatusInternalServerError
//	}
//
//	sort.Slice(result.Services, func(i, j int) bool {
//		return result.Services[i].Name < result.Services[j].Name
//	})
//
//
//	return result, nil, http.StatusOK
//}

func (this *Controller) ValidateConcept(concept model.Concept) (err error, code int) {
	if concept.Id == "" {
		return errors.New("missing concept id"), http.StatusBadRequest
	}
	if concept.Name == "" {
		return errors.New("missing concept name"), http.StatusBadRequest
	}

	if concept.RdfType != model.SES_ONTOLOGY_CONCEPT {
		return errors.New("wrong concept type"), http.StatusBadRequest
	}

	err, code = this.ValidateCharacteristics(concept.Characteristics)
	if err != nil {
		return err, code
	}

	return nil, http.StatusOK
}

/////////////////////////
//		source
/////////////////////////

func (this *Controller) SetConcept(concept model.Concept, owner string) (err error) {
	SetConceptRdfTypes(&concept)

	err, code := this.ValidateConcept(concept)
	if err != nil {
		debug.PrintStack()
		log.Println("Error Validation:", err, code)
		return
	}
	// delete is required for the update of concepts
	err = this.DeleteConcept(concept.Id)
	if err != nil {
		debug.PrintStack()
		log.Println("Error Delete Device Type:", err, code)
		return
	}

	b, err := json.Marshal(concept)
	var deviceTypeJsonLd map[string]interface{}
	err = json.Unmarshal(b, &deviceTypeJsonLd)

	deviceTypeJsonLd["@context"] = getDeviceTypeContext()

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
		log.Println("Error insert devicetype:", err)
		return err
	}
	return nil
}

func SetConceptRdfTypes(concept *model.Concept) {
	concept.RdfType = model.SES_ONTOLOGY_CONCEPT
	SetCharacteristicRdfTypes(concept.Characteristics)
}

func SetCharacteristicRdfTypes(characteristic []model.Characteristic) {
	for charIndex, _ := range characteristic {
		characteristic[charIndex].RdfType = model.SES_ONTOLOGY_CHARACTERISTIC
		SetCharacteristicRdfTypes(characteristic[charIndex].SubCharacteristics)
	}
}

func (this *Controller) DeleteConcept(id string) (err error) {
	if id == "" {
		debug.PrintStack()
		return errors.New("missing concept id")
	}

	//err = this.db.DeleteDeviceType(id)
	//if err != nil {
	//	debug.PrintStack()
	//	return err
	//}
	return nil
}