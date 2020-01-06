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
	"log"
	"net/http"
	"sort"
)

/////////////////////////
//		api
/////////////////////////

func (this *Controller) GetDeviceClasses() (result []model.DeviceClass, err error, errCode int) {
	deviceClasses, err := this.db.GetListWithoutSubProperties(model.RDF_TYPE, model.SES_ONTOLOGY_DEVICE_CLASS)
	if err != nil {
		log.Println("GetDeviceClasses ERROR: GetListWithoutSubProperties", err)
		return result, err, http.StatusInternalServerError
	}

	err = this.RdfXmlToModel(deviceClasses, &result)
	if err != nil {
		log.Println("GetDeviceClasses ERROR: RdfXmlToModel", err)
		return result, err, http.StatusInternalServerError
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result, nil, http.StatusOK
}

func (this *Controller) GetDeviceClassesWithControllingFunctions() (result []model.DeviceClass, err error, errCode int) {
	deviceClasses, err := this.db.GetDeviceClassesWithControllingFunctions()
	if err != nil {
		log.Println("GetDeviceClassesWithControllingFunctions ERROR: GetDeviceClassesWithControllingFunctions", err)
		return result, err, http.StatusInternalServerError
	}

	err = this.RdfXmlToModel(deviceClasses, &result)
	if err != nil {
		log.Println("GetDeviceClassesWithControllingFunctions ERROR: RdfXmlToModel", err)
		return result, err, http.StatusInternalServerError
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result, nil, http.StatusOK
}

func (this *Controller) GetDeviceClassesFunctions(subject string) (result []model.Function, err error, errCode int) {
	functions, err := this.db.GetDeviceClassesFunctions(subject)
	if err != nil {
		log.Println("GetDeviceClassesFunctions ERROR: GetDeviceClassesFunctions", err)
		return result, err, http.StatusInternalServerError
	}

	err = this.RdfXmlToModel(functions, &result)
	if err != nil {
		log.Println("GetDeviceClassesFunctions ERROR: RdfXmlToModel", err)
		return result, err, http.StatusInternalServerError
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result, nil, http.StatusOK
}

func (this *Controller) GetDeviceClassesControllingFunctions(subject string) (result []model.Function, err error, errCode int) {
	functions, err := this.db.GetDeviceClassesControllingFunctions(subject)
	if err != nil {
		log.Println("GetDeviceClassesControllingFunctions ERROR: GetDeviceClassesControllingFunctions", err)
		return result, err, http.StatusInternalServerError
	}

	err = this.RdfXmlToModel(functions, &result)
	if err != nil {
		log.Println("GetDeviceClassesControllingFunctions ERROR: RdfXmlToModel", err)
		return result, err, http.StatusInternalServerError
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result, nil, http.StatusOK
}

func (this *Controller) ValidateDeviceClass(deviceClass model.DeviceClass) (error, int) {

	if deviceClass.Id == "" {
		return errors.New("missing device class id"), http.StatusBadRequest
	}
	if deviceClass.Name == "" {
		return errors.New("missing device class name"), http.StatusBadRequest
	}
	if deviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		return errors.New("wrong device class type"), http.StatusBadRequest
	}

	return nil, http.StatusOK
}

/////////////////////////
//		source
/////////////////////////
