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
	"github.com/SmartEnergyPlatform/jwt-http-router"
	"github.com/piprate/json-gold/ld"
	"log"
	"net/http"
)

/////////////////////////
//		api
/////////////////////////

func (this *Controller) ReadDeviceType(id string, jwt jwt_http_router.Jwt) (result model.DeviceType, err error, errCode int) {
	panic("not implemented")
	/*
		ctx, _ := getTimeoutContext()
		deviceType, exists, err := this.db.GetDeviceType(ctx, id)
		if err != nil {
			return result, err, http.StatusInternalServerError
		}
		if !exists {
			return result, errors.New("not found"), http.StatusNotFound
		}
		return deviceType, nil, http.StatusOK

	*/
}

func (this *Controller) ValidateDeviceType(dt model.DeviceType) (err error, code int) {
	if dt.Id == "" {
		return errors.New("missing device-type id"), http.StatusBadRequest
	}
	if dt.Name == "" {
		return errors.New("missing device-type name"), http.StatusBadRequest
	}

	if dt.Type != model.SES_ONTOLOGY_DEVICE_TYPE {
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
	SetTypes(&deviceType)

	err, _ = this.ValidateDeviceType(deviceType)
	if err != nil {
		log.Println("Error Validation:", err)
		return
	}
	log.Println(deviceType)

	b, err := json.Marshal(deviceType)
	var result map[string]interface{}
	err = json.Unmarshal(b, &result)

	context := map[string]interface{}{
		"id":           "@id",
		"type":         "@type",
		"name":         model.RDFS_LABEL,
		"device_class": model.SES_ONTOLOGY_HAS_DEVICE_CLASS,
		"services":     model.SES_ONTOLOGY_HAS_SERVICE,
		"aspects":      model.SES_ONTOLOGY_REFERS_TO,
		"functions":    model.SES_ONTOLOGY_EXPOSES_FUNCTION,
		"concept_ids": map[string]interface{}{
			"@id":        model.SES_ONTOLOGY_HAS_CONCEPT,
			"@type":      "@id",
			"@container": "@set",
		},
	}
	result["@context"] = context

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	//options.CompactArrays = false
	options.ProcessingMode = ld.JsonLd_1_1
	options.Format = "application/n-quads"

	triples, err := proc.ToRDF(result, options)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return
	}

	log.Println("---->")
	log.Println(triples)

	//os.Stdout.WriteString(triples.(string))
	success, err := this.db.InsertData(triples.(string))
	log.Println(success)
	/*
		ctx, _ := getTimeoutContext()
		return this.db.SetDeviceType(ctx, deviceType)
	*/
	return
}

func SetTypes(deviceType *model.DeviceType) {
	deviceType.Type = model.SES_ONTOLOGY_DEVICE_TYPE
	deviceType.DeviceClass.Type = model.SES_ONTOLOGY_DEVICE_CLASS
	for serviceIndex, _ := range deviceType.Services {
		deviceType.Services[serviceIndex].Type = model.SES_ONTOLOGY_SERVICE
		for aspectIndex, _ := range deviceType.Services[serviceIndex].Aspects {
			deviceType.Services[serviceIndex].Aspects[aspectIndex].Type = model.SES_ONTOLOGY_ASPECT
		}
	}
}

func (this *Controller) DeleteDeviceType(id string) error {
	return nil
	// todo
	/*
		ctx, _ := getTimeoutContext()
		return this.db.RemoveDeviceType(ctx, id)
	*/
}
