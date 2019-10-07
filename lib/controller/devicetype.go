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
	"sort"
)

/////////////////////////
//		api
/////////////////////////

func (this *Controller) GetDeviceType(deviceTypeId string) (result model.DeviceType, err error, errCode int) {
	deviceType, err := this.db.GetDeviceType(deviceTypeId, "", "", "")
	if err != nil {
		log.Println("GetDeviceType ERROR: GetDeviceType", err)
		return result, err, http.StatusInternalServerError
	}

	res := []model.DeviceType{}
	err = this.RdfXmlFrame(deviceType, &res, deviceTypeId)
	if err != nil {
		log.Println("GetDeviceType ERROR: RdfXmlFrame", err)
		return result, err, http.StatusInternalServerError
	}

	sort.Slice(res[0].Services, func(i, j int) bool {
		return res[0].Services[i].Name < res[0].Services[j].Name
	})

	return res[0], nil, http.StatusOK
}

func (this *Controller) GetDeviceTypesFiltered(deviceClassId string, functionId string, aspectId string) (result []model.DeviceType, err error, errCode int) {
	deviceTypes, err := this.db.GetDeviceType("", deviceClassId, functionId, aspectId)
	if err != nil {
		log.Println("GetDeviceType ERROR: GetDeviceTypesFiltered", err)
		return result, err, http.StatusInternalServerError
	}

	err = this.RdfXmlFrame(deviceTypes, &result, "")
	if err != nil {
		log.Println("GetDeviceType ERROR: RdfXmlFrame", err)
		return result, err, http.StatusInternalServerError
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result, nil, http.StatusOK
}

func (this *Controller) ValidateDeviceType(dt model.DeviceType) (err error, code int) {
	if dt.Id == "" {
		return errors.New("missing device-type id"), http.StatusBadRequest
	}
	if dt.Name == "" {
		return errors.New("missing device-type name"), http.StatusBadRequest
	}

	if dt.RdfType != model.SES_ONTOLOGY_DEVICE_TYPE {
		return errors.New("wrong device type"), http.StatusBadRequest
	}

	err, code = this.ValidateService(dt.Services)
	if err != nil {
		return err, code
	}

	err, code = this.ValidateDeviceClass(dt.DeviceClass)
	if err != nil {
		return err, code
	}

	return nil, http.StatusOK
}

/////////////////////////
//		source
/////////////////////////

func (this *Controller) SetDeviceType(deviceType model.DeviceType, owner string) (err error) {
	SetDevicetypeRdfTypes(&deviceType)

	err, code := this.ValidateDeviceType(deviceType)
	if err != nil {
		debug.PrintStack()
		log.Println("Error Validation:", err, code)
		return
	}
	// delete is required for the update of devicetypes
	err = this.DeleteDeviceType(deviceType.Id)
	if err != nil {
		debug.PrintStack()
		log.Println("Error Delete Device Type:", err, code)
		return
	}

	b, err := json.Marshal(deviceType)
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

func SetDevicetypeRdfTypes(deviceType *model.DeviceType) {
	deviceType.RdfType = model.SES_ONTOLOGY_DEVICE_TYPE
	deviceType.DeviceClass.RdfType = model.SES_ONTOLOGY_DEVICE_CLASS
	for serviceIndex, _ := range deviceType.Services {
		deviceType.Services[serviceIndex].RdfType = model.SES_ONTOLOGY_SERVICE
		for aspectIndex, _ := range deviceType.Services[serviceIndex].Aspects {
			deviceType.Services[serviceIndex].Aspects[aspectIndex].RdfType = model.SES_ONTOLOGY_ASPECT
		}
	}
}

func (this *Controller) DeleteDeviceType(id string) (err error) {
	if id == "" {
		debug.PrintStack()
		return errors.New("missing deviceType id")
	}

	err = this.db.DeleteDeviceType(id)
	if err != nil {
		debug.PrintStack()
		return err
	}
	return nil
}
