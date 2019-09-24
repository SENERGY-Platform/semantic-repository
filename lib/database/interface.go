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

package database

type Database interface {
	Disconnect()

	InsertData(triples string) (err error)
	ReadData() (body []byte, err error)
	DeleteDeviceType(s string) (err error)
	GetDeviceType(s string) (rdfxml string, err error)
	DeleteConcept(s string, withNested bool) (err error)
	DeleteCharacteristic(s string) (err error)
	GetConcept(s string) (rdfxml string, err error)
	GetConstructWithoutProperties(s string, p string, o string) (rdfxml string, err error)
	GetConstructWithProperties(subject string) (rdfxml string, err error)
}
