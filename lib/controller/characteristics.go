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

func (this *Controller) GetCharacteristic(subject string) (result model.Characteristic, err error, errCode int) {
	characteristic, err := this.db.GetWithAllSubProperties(subject)
	if err != nil {
		log.Println("GetCharacteristic ERROR: GetWithAllSubProperties", err)
		return result, err, http.StatusInternalServerError
	}

	res := []model.Characteristic{}
	err = this.RdfXmlFrame(characteristic, &res, subject)
	if err != nil {
		log.Println("GetCharacteristic ERROR: RdfXmlFrame", err)
		return result, err, http.StatusInternalServerError
	}

	return res[0], nil, http.StatusOK
}

func (this *Controller) GetLeafCharacteristics() (result []model.Characteristic, err error, errCode int) {
	characteristics, err := this.db.GetLeafCharacteristics()
	if err != nil {
		log.Println("GetLeafCharacteristics ERROR: GetLeafCharacteristics", err)
		return result, err, http.StatusInternalServerError
	}

	res := []model.Characteristic{}
	err = this.RdfXmlFrame(characteristics, &res, "")
	if err != nil {
		log.Println("GetLeafCharacteristics ERROR: RdfXmlFrame", err)
		return result, err, http.StatusInternalServerError
	}

	return res, nil, http.StatusOK
}

func (this *Controller) ValidateCharacteristics(characteristic model.Characteristic) (error, int) {
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

	err, code := this.validateSubCharacteristics(characteristic.SubCharacteristics)
	if err != nil {
		return err, code
	}

	return nil, http.StatusOK
}

func (this *Controller) validateSubCharacteristics(characteristics []model.Characteristic) (error, int) {
	for _, characteristic := range characteristics {
		err, code := this.ValidateCharacteristics(characteristic)
		if err != nil {
			return err, code
		}
	}

	return nil, http.StatusOK
}

/////////////////////////
//		source
/////////////////////////

func SetCharacteristicRdfTypes(characteristic *model.Characteristic) {
	characteristic.RdfType = model.SES_ONTOLOGY_CHARACTERISTIC
	SetSubCharacteristicRdfTypes(characteristic.SubCharacteristics)
}

func SetSubCharacteristicRdfTypes(characteristic []model.Characteristic) {
	for charIndex, _ := range characteristic {
		characteristic[charIndex].RdfType = model.SES_ONTOLOGY_CHARACTERISTIC
		SetSubCharacteristicRdfTypes(characteristic[charIndex].SubCharacteristics)
	}
}

func (this *Controller) SetCharacteristic(conceptId string, characteristic model.Characteristic, owner string) (err error) {
	SetCharacteristicRdfTypes(&characteristic)

	err, code := this.ValidateCharacteristics(characteristic)
	if err != nil {
		debug.PrintStack()
		log.Println("Error Validation:", err, code)
		return
	}

	if conceptId == "" {
		debug.PrintStack()
		log.Println("Error missing conceptId:", err, code)
		return
	}
	// delete is required for the update of characteristics
	err = this.DeleteCharacteristic(characteristic.Id)
	if err != nil {
		debug.PrintStack()
		log.Println("Error Delete Characteristic:", err, code)
		return
	}

	b, err := json.Marshal(characteristic)
	var deviceTypeJsonLd map[string]interface{}
	err = json.Unmarshal(b, &deviceTypeJsonLd)

	deviceTypeJsonLd["@context"] = getCharacteristicsContext()

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
		log.Println("Error insert characteristics:", err)
		return err
	}

	err = this.db.InsertData("<" + conceptId + "> <" + model.SES_ONTOLOGY_HAS_CHARACTERISTIC + "> <" + characteristic.Id + "> .")
	if err != nil {
		debug.PrintStack()
		log.Println("Error insert hasCharacteristics:", err)
		return err
	}
	return nil
}

func (this *Controller) DeleteCharacteristic(id string) (err error) {
	if id == "" {
		debug.PrintStack()
		return errors.New("missing characteristic id")
	}

	err = this.db.DeleteCharacteristic(id)
	if err != nil {
		debug.PrintStack()
		return err
	}
	return nil
}
