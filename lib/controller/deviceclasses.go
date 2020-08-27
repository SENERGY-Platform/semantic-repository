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
	"strings"
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

func (this *Controller) GetDeviceClass(id string) (result model.DeviceClass, err error, errCode int) {
	deviceclass, err := this.db.GetSubject(id, model.SES_ONTOLOGY_DEVICE_CLASS)
	if err != nil {
		log.Println("GetDeviceClass ERROR: GetSubject", err)
		return result, err, http.StatusInternalServerError
	}

	res := []model.DeviceClass{}
	err = this.RdfXmlToModel(deviceclass, &res)
	if err != nil {
		log.Println("GetDeviceClass ERROR: RdfXmlToModel", err)
		return result, err, http.StatusInternalServerError
	}

	if len(res) == 0 {
		return result, errors.New("not found"), http.StatusNotFound
	}

	return res[0], nil, http.StatusOK
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
	if !strings.HasPrefix(deviceClass.Id, model.URN_PREFIX) {
		return errors.New("invalid deviceClass id"), http.StatusBadRequest
	}
	if deviceClass.Name == "" {
		return errors.New("missing device class name"), http.StatusBadRequest
	}
	if deviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		return errors.New("wrong device class type"), http.StatusBadRequest
	}

	return nil, http.StatusOK
}

func (this *Controller) SetDeviceClass(deviceclass model.DeviceClass, owner string) (err error) {
	SetDeviceClassRdfType(&deviceclass)

	err, code := this.ValidateDeviceClass(deviceclass)
	if err != nil {
		debug.PrintStack()
		log.Println("Error Validation:", err, code)
		return
	}

	err = this.DeleteDeviceClass(deviceclass.Id)
	if err != nil {
		debug.PrintStack()
		log.Println("Error Delete DeviceClass:", err, code)
		return
	}

	b, err := json.Marshal(deviceclass)
	var deviceTypeJsonLd map[string]interface{}
	err = json.Unmarshal(b, &deviceTypeJsonLd)

	deviceTypeJsonLd["@context"] = getDeviceClassContext()

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
		log.Println("Error insert Deviceclass:", err)
		return err
	}

	return nil
}

/////////////////////////
//		source
/////////////////////////
func SetDeviceClassRdfType(deviceclass *model.DeviceClass) {
	deviceclass.RdfType = model.SES_ONTOLOGY_DEVICE_CLASS
}

func (this *Controller) DeleteDeviceClass(id string) (err error) {
	if id == "" {
		debug.PrintStack()
		return errors.New("missing deviceclass id")
	}

	err = this.db.DeleteSubject(id, model.SES_ONTOLOGY_DEVICE_CLASS)
	if err != nil {
		debug.PrintStack()
		return err
	}
	return nil
}
