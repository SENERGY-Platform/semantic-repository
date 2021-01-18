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
	"fmt"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/piprate/json-gold/ld"
	"log"
	"net/http"
	"net/url"
	"runtime/debug"
	"strings"
)

/////////////////////////
//		api
/////////////////////////

func (this *Controller) GetLocation(id string) (result model.Location, err error, errCode int) {
	location, err := this.db.GetSubject(id, model.SES_ONTOLOGY_LOCATION)
	if err != nil {
		log.Println("GetLocation ERROR: GetSubject", err)
		return result, err, http.StatusInternalServerError
	}

	res := []model.Location{}
	err = this.RdfXmlToModelWithContext(location, &res, getLocationContext())
	//err = this.RdfXmlToModel(location, &res)
	if err != nil {
		log.Println("GetLocation ERROR: RdfXmlToModel", err)
		return result, err, http.StatusInternalServerError
	}

	if len(res) == 0 {
		return result, errors.New("not found"), http.StatusNotFound
	}

	return res[0], nil, http.StatusOK
}

func (this *Controller) ValidateLocation(location model.Location) (error, int) {
	if location.Id == "" {
		return errors.New("missing device class id"), http.StatusBadRequest
	}
	if !strings.HasPrefix(location.Id, model.URN_PREFIX) {
		return errors.New("invalid location id"), http.StatusBadRequest
	}
	if location.Name == "" {
		return errors.New("missing device class name"), http.StatusBadRequest
	}
	if location.Image != "" {
		if _, err := url.ParseRequestURI(location.Image); err != nil {
			return fmt.Errorf("image is not valid URL: %w", err), http.StatusBadRequest
		}
	}
	if location.RdfType != model.SES_ONTOLOGY_LOCATION {
		return errors.New("wrong device class type"), http.StatusBadRequest
	}

	return nil, http.StatusOK
}

func (this *Controller) SetLocation(location model.Location, owner string) (err error) {
	SetLocationRdfType(&location)

	err, code := this.ValidateLocation(location)
	if err != nil {
		debug.PrintStack()
		log.Println("Error Validation:", err, code)
		return
	}

	err = this.DeleteLocation(location.Id)
	if err != nil {
		debug.PrintStack()
		log.Println("Error Delete Location:", err, code)
		return
	}

	b, err := json.Marshal(location)
	var elementLd map[string]interface{}
	err = json.Unmarshal(b, &elementLd)

	elementLd["@context"] = getLocationContext()

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.ProcessingMode = ld.JsonLd_1_1
	options.Format = "application/n-quads"

	triples, err := proc.ToRDF(elementLd, options)
	if err != nil {
		debug.PrintStack()
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return err
	}
	err = this.db.InsertData(triples.(string))
	if err != nil {
		debug.PrintStack()
		log.Println("Error insert Location:", err)
		return err
	}

	return nil
}

/////////////////////////
//		source
/////////////////////////
func SetLocationRdfType(location *model.Location) {
	location.RdfType = model.SES_ONTOLOGY_LOCATION
}

func (this *Controller) DeleteLocation(id string) (err error) {
	if id == "" {
		debug.PrintStack()
		return errors.New("missing location id")
	}

	err = this.db.DeleteSubject(id, model.SES_ONTOLOGY_LOCATION)
	if err != nil {
		debug.PrintStack()
		return err
	}
	return nil
}
