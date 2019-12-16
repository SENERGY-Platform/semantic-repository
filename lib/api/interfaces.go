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

package api

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
)

type Controller interface {
	GetDeviceType(deviceTypeId string) (result model.DeviceType, err error, errCode int)
	GetDeviceTypesFiltered(deviceClassId string, functionIds []string, aspectIds []string) (result []model.DeviceType, err error, errCode int)
	ValidateDeviceType(deviceType model.DeviceType) (err error, code int)

	GetFunctions(funcType string) (result []model.Function, err error, errCode int)
	GetDeviceClasses() (result []model.DeviceClass, err error, errCode int)
	GetDeviceClassesFunctions(subject string) (result []model.Function, err error, errCode int)
	GetDeviceClassesControllingFunctions(subject string) (result []model.Function, err error, errCode int)
	GetAspects() (result []model.Aspect, err error, errCode int)
	GetAspectsMeasuringFunctions(subject string) (result []model.Function, err error, errCode int)

	GetConceptWithoutCharacteristics(subject string) (result model.Concept, err error, errCode int)
	GetConceptWithCharacteristics(subject string) (result model.ConceptWithCharacteristics, err error, errCode int)
	ValidateConcept(concept model.Concept) (err error, code int)

	GetCharacteristic(subject string) (result model.Characteristic, err error, errCode int)
	ValidateCharacteristics(concept model.Characteristic) (err error, code int)
}
