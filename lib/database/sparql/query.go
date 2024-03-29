/*
 *
 * Copyright 2019 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *
 */

package sparql

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"net/url"
	"strconv"
	"strings"
)

func (*Database) getConstructListWithoutSubProperties(p string, o string) string {
	if p == "" {
		p = " ?p "
	} else {
		p = " <" + p + "> "
	}

	if o == "" {
		o = " ?o "
	} else {
		o = " <" + o + "> "
	}
	return url.QueryEscape("CONSTRUCT {?s ?p ?o} " +
		"WHERE {" + "?s ?p ?o ." +
		" ?s " + p + o + "." +
		"?s <http://www.w3.org/2000/01/rdf-schema#label> ?label" +
		"}" +
		"ORDER BY ASC(?label)")
}

func (*Database) getFunctionsWithoutSubPropertiesLimitOffsetSearch(limit int, offset int, search string, direction string) string {
	//PREFIX ses: <https://senergy.infai.org/ontology/>
	//PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
	//
	//CONSTRUCT {
	//	?s rdfs:label ?label;
	//	rdf:type ?type;
	//	ses:hasConcept ?concept .
	//}
	//
	//
	//WHERE {
	//	?s rdfs:label ?label;
	//	rdf:type ?type.
	//	OPTIONAL {?s ses:hasConcept ?concept .}
	//
	//	VALUES ?type { <https://senergy.infai.org/ontology/ControllingFunction> <https://senergy.infai.org/ontology/MeasuringFunction> }
	//	FILTER CONTAINS (?label, 'Func')
	//
	//}
	//ORDER BY desc(?label)
	//LIMIT 6
	//OFFSET 0

	numberOfFields := 5
	query := "PREFIX ses: <https://senergy.infai.org/ontology/> " +
		"PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>" +
		"CONSTRUCT {" +
		"?s rdfs:label ?label;" +
		"rdfs:comment ?comment;" +
		"ses:displayName ?displayName;" +
		"rdf:type ?type;" +
		"ses:hasConcept ?concept ." +
		"}" +
		"WHERE {" +
		"?s rdfs:label ?label;" +
		"rdf:type ?type." +
		"OPTIONAL {?s ses:hasConcept ?concept .}" +
		"OPTIONAL {?s rdfs:comment ?comment .}" +
		"OPTIONAL {?s ses:displayName ?displayName .}" +
		"VALUES ?type { <" + model.SES_ONTOLOGY_CONTROLLING_FUNCTION + "> <" + model.SES_ONTOLOGY_MEASURING_FUNCTION + "> }"

	if search != "" {
		query += "FILTER CONTAINS (?label, '" + search + "')"
	}

	query += "}" +
		"ORDER BY "
	if direction == "asc" || direction == "desc" {
		query += direction
	} else {
		query += "asc"
	}

	query += "(?label)" +
		"LIMIT " + strconv.Itoa(numberOfFields*limit) + " " +
		"OFFSET " + strconv.Itoa(numberOfFields*offset)

	return url.QueryEscape(query)
}

func (*Database) getFunctionsCount(search string) string {
	//PREFIX ses: <https://senergy.infai.org/ontology/>
	//PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
	//
	//CONSTRUCT
	//{
	//	ses:Count ses:totalCount ?cnt .
	//}
	//WHERE
	//{
	//	Select (count(*) as ?cnt)
	//	{
	//	?s rdfs:label ?label;
	//	rdf:type ?type.
	//	OPTIONAL {?s ses:hasConcept ?concept .}
	//
	//	VALUES ?type { <https://senergy.infai.org/ontology/ControllingFunction> <https://senergy.infai.org/ontology/MeasuringFunction>}
	//	FILTER CONTAINS (?label, "")
	//	}
	//	}

	query := "PREFIX ses: <https://senergy.infai.org/ontology/> " +
		"PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>" +
		"CONSTRUCT { " +
		"ses:Count ses:totalCount ?cnt ." +
		"}" +
		"WHERE" +
		"{" +
		"Select (count(*) as ?cnt) " +
		"{" +
		"?s rdfs:label ?label;" +
		"rdf:type ?type." +
		"OPTIONAL {?s ses:hasConcept ?concept .}" +
		"VALUES ?type { <" + model.SES_ONTOLOGY_CONTROLLING_FUNCTION + "> <" + model.SES_ONTOLOGY_MEASURING_FUNCTION + "> }"

	if search != "" {
		query += "FILTER CONTAINS (?label, '" + search + "')"
	}

	query += "}}"
	return url.QueryEscape(query)
}

func (*Database) getConstructWithAllSubProperties(subject string) string {

	return url.QueryEscape("prefix x: <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> " +
		"construct { ?s ?p ?o } " +
		"where {<" + subject + "> (x:|!x:)* ?s . ?s ?p ?o . }")
}

func (*Database) getConstructWithoutSubProperties(subject string) string {
	//construct {<urn:ses:infai:concept:1a1a1a> ?p ?o .}
	//where {<urn:ses:infai:concept:1a1a1a> ?p ?o .}
	return url.QueryEscape(
		"construct { <" + subject + "> ?p ?o .} " +
			"where {<" + subject + "> ?p ?o .}")
}

func (*Database) getDeleteDeviceTypeQuery(subject string) string {

	return url.QueryEscape(
		model.PREFIX_SES +
			model.PREFIX_RDF +
			"delete {" +
			"<" + subject + "> rdf:type ?type;" +
			"rdfs:label ?label;" +
			"ses:hasDeviceClass ?deviceclass;" +
			"ses:hasService ?service ." +
			"?service ?p ?o." +
			"} where {" +
			"<" + subject + "> rdf:type ?type;" +
			"rdfs:label ?label;" +
			"ses:hasDeviceClass ?deviceclass;" +
			"ses:hasService ?service ." +
			"?service ?p ?o." +
			"}")
}

func (*Database) getDeviceTypeQuery(deviceTypeId string, filters []model.DeviceTypesFilter) string {

	// Example Devicetype

	//PREFIX ses: <https://senergy.infai.org/ontology/>
	//PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
	//construct { <urn:infai:ses:device-type:eb4a3337-01a1-4434-9dcc-064b3955eeef>  rdf:type ?type;
	//	rdfs:label ?label;
	//	ses:hasDeviceClass  ?deviceclass ;
	//	ses:hasService ?service0 .
	//
	//	?service0 rdf:type ?s_type0;
	//	rdfs:label ?s_label0;
	//	ses:refersTo ?aspect0;
	//	ses:exposesFunction ?function0 .}
	//
	//where
	//{ <urn:infai:ses:device-type:eb4a3337-01a1-4434-9dcc-064b3955eeef>  rdf:type ?type;
	//rdfs:label ?label;
	//ses:hasDeviceClass  ?deviceclass ;
	//ses:hasService ?service0 .
	//
	//	?service0 rdf:type ?s_type0;
	//rdfs:label ?s_label0;
	//ses:refersTo ?aspect0;
	//ses:exposesFunction ?function0 .}

	// Example Deviceclass & Function

	//	PREFIX ses: <https://senergy.infai.org/ontology/> PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>construct {?devicetype rdf:type ?type;
	//rdfs:label ?label;
	//ses:hasDeviceClass  <urn:infai:ses:device-class:14e56881-16f9-4120-bb41-270a43070c86> ;
	//ses:hasService ?service0 .
	//	?service0 rdf:type ?s_type0;
	//rdfs:label ?s_label0;
	//ses:refersTo ?aspect0;
	//ses:exposesFunction <urn:infai:ses:controlling-function:c54e2a89-1fb8-4ecb-8993-a7b40b355599> .
	//
	//	?service0 ses:hasProtocol ?s_protocol0.
	//	?service0 ses:interaction ?s_interaction0. }
	//where
	//{?devicetype rdf:type ?type;
	//rdfs:label ?label;
	//ses:hasDeviceClass  <urn:infai:ses:device-class:14e56881-16f9-4120-bb41-270a43070c86> ;
	//ses:hasService ?service0 .
	//?service0 rdf:type ?s_type0;
	//rdfs:label ?s_label0;
	//ses:refersTo ?aspect0;
	//ses:exposesFunction <urn:infai:ses:controlling-function:c54e2a89-1fb8-4ecb-8993-a7b40b355599> .
	//OPTIONAL {?service0 ses:hasProtocol ?s_protocol0.
	//?service0 ses:interaction ?s_interaction0. } }

	devicetype := "?devicetype"
	if deviceTypeId != "" {
		devicetype = " <" + deviceTypeId + "> "
	}
	deviceclass := "?deviceclass"
	for _, filter := range filters {
		if filter.DeviceClassId != "" {
			if deviceclass == "?deviceclass" {
				deviceclass = "<" + filter.DeviceClassId + ">"
			} else {
				if deviceclass != "<"+filter.DeviceClassId+">" {
					// error
					return ""
				}
			}
		}

	}

	// linebreak
	lnb := "\n"
	construct := ""
	where := ""

	construct +=
		devicetype + " rdf:type ?type;" + lnb +
			"rdfs:label ?label;" + lnb +
			"ses:hasDeviceClass  " + deviceclass + " ;" + lnb +
			"ses:hasService "

	services := []string{}
	for i := 0; i < len(filters); i++ {
		services = append(services, "?service"+strconv.Itoa(i))
	}
	construct += strings.Join(services, ", ") + " ." + lnb

	where += construct
	for index, filter := range filters {
		convIndex := strconv.Itoa(index)
		aspect := "?aspect" + convIndex
		if filter.AspectId != "" {
			aspect = "<" + filter.AspectId + ">"
		}
		function := "?function" + convIndex
		if filter.FunctionId != "" {
			function = "<" + filter.FunctionId + ">"
		}

		protocolInService := "?service" + convIndex + " ses:hasProtocol ?s_protocol" + convIndex + ". "
		interactionInService := "?service" + convIndex + " ses:interaction ?s_interaction" + convIndex + ". "

		service :=
			"?service" + convIndex + " rdf:type ?s_type" + convIndex + ";" + lnb +
				"rdfs:label ?s_label" + convIndex + ";" + lnb +
				"ses:refersTo " + aspect + ";" + lnb +
				"ses:exposesFunction " + function + " ." + lnb
		construct += service + lnb + protocolInService + lnb + interactionInService
		where += service + "OPTIONAL {" + protocolInService + lnb + interactionInService + "} "
	}

	query := model.PREFIX_SES +
		model.PREFIX_RDF +
		"construct {" +
		construct +
		"} " + lnb +
		"where" + lnb +
		" {" +
		where +
		"}"
	return url.QueryEscape(
		query)
}

func (*Database) getLeafCharacteristics() string {

	//PREFIX ses: <https://senergy.infai.org/ontology/>
	//PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
	//
	//construct {?characteristics ?p ?o} where {
	//	?characteristics ?p ?o .
	//	?characteristics rdf:type ses:Characteristic.
	//	OPTIONAL {?characteristics ses:hasSubCharacteristic ?subChar } .
	//		FILTER ( !bound(?subChar) )
	//}

	return url.QueryEscape(
		model.PREFIX_SES +
			model.PREFIX_RDF +
			"construct {" +
			"?characteristics ?p ?o ." +
			"} where {" +
			"?characteristics ?p ?o ." +
			"?characteristics rdf:type ses:Characteristic ." +
			"OPTIONAL {?characteristics ses:hasSubCharacteristic ?subChar } ." +
			"FILTER ( !bound(?subChar) )" +
			"}")
}

func (*Database) getDeviceClassesFunctions(subject string) string {

	//PREFIX ses: <https://senergy.infai.org/ontology/>
	//PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
	//construct {
	//	?function ?p ?o
	//} where {
	//	?s ses:hasDeviceClass <urn:infai:ses:device-class:d8473e94-624e-4581-aafd-ff2962a4f81b>;
	//	ses:hasService ?service .
	//	?service ses:exposesFunction ?function.
	//	?function ?p ?o
	//}

	return url.QueryEscape(
		model.PREFIX_SES +
			model.PREFIX_RDF +
			"construct {" +
			"?function ?p ?o" +
			"} where {" +
			"?s ses:hasDeviceClass <" + subject + ">;" +
			"ses:hasService ?service ." +
			"?service ses:exposesFunction ?function." +
			"?function ?p ?o" +
			"}")
}

func (*Database) getDeviceClassesWithControllingFunctions() string {

	//PREFIX ses: <https://senergy.infai.org/ontology/>
	//PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
	//construct {
	//	?deviceclass rdf:type ses:DeviceClass;
	//	rdfs:label ?dc_name.
	//} where {
	//	?deviceclass rdf:type ses:DeviceClass;
	//	rdfs:label ?dc_name.
	//
	//	?devicetype ses:hasDeviceClass ?deviceclass;
	//	ses:hasService ?service.
	//
	//	?service ses:exposesFunction ?function.
	//
	//	?function rdf:type ses:ControllingFunction.
	//}

	return url.QueryEscape(
		model.PREFIX_SES +
			model.PREFIX_RDF +
			"construct {" +
			"?deviceclass rdf:type ses:DeviceClass;" +
			"rdfs:label ?dc_name." +
			"} where {" +
			"?deviceclass rdf:type ses:DeviceClass;" +
			"rdfs:label ?dc_name." +
			"?devicetype ses:hasDeviceClass ?deviceclass;" +
			"ses:hasService ?service." +
			"?service ses:exposesFunction ?function." +
			"?function rdf:type ses:ControllingFunction." +
			"}")
}

func (*Database) getDeviceClassesControllingFunctions(subject string) string {

	//PREFIX ses: <https://senergy.infai.org/ontology/>
	//PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
	//construct {
	//	?function ?p ?o .
	//} where {
	//	?s ses:hasDeviceClass <urn:infai:ses:deviceclass:2e2e-DeviceClassTest>;
	//	ses:hasService ?service .
	//	?service ses:exposesFunction ?function.
	//	?function rdf:type <https://senergy.infai.org/ontology/ControllingFunction>;
	//	?p ?o .
	//}

	return url.QueryEscape(
		model.PREFIX_SES +
			model.PREFIX_RDF +
			"construct {" +
			"?function ?p ?o" +
			"} where {" +
			"?s ses:hasDeviceClass <" + subject + ">;" +
			"ses:hasService ?service ." +
			"?service ses:exposesFunction ?function." +
			"?function rdf:type <" + model.SES_ONTOLOGY_CONTROLLING_FUNCTION + "> ;" +
			"?p ?o" +
			"}")
}

func (*Database) getAspectsMeasuringFunctions(subject string) string {

	//PREFIX ses: <https://senergy.infai.org/ontology/>
	//PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
	//construct {
	//	?function ?p ?o .
	//} where {
	//	?service ses:refersTo <urn:infai:ses:aspect:4e4e-DeviceClassTest>;
	//	ses:exposesFunction ?function.
	//	?function rdf:type <https://senergy.infai.org/ontology/MeasuringFunction>;
	//	?p ?o .
	//}

	return url.QueryEscape(
		model.PREFIX_SES +
			model.PREFIX_RDF +
			"construct {" +
			"?function ?p ?o" +
			"} where {" +
			"?service ses:refersTo <" + subject + ">;" +
			"ses:exposesFunction ?function ." +
			"?function rdf:type <" + model.SES_ONTOLOGY_MEASURING_FUNCTION + "> ;" +
			"?p ?o" +
			"}")
}

func (*Database) getAspectsWithMeasuringFunction() string {

	//PREFIX ses: <https://senergy.infai.org/ontology/>
	//PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
	//construct {
	//	?aspect rdf:type ses:Aspect;
	//	rdfs:label ?dc_name.
	//} where {
	//	?aspect rdf:type ses:Aspect;
	//	rdfs:label ?dc_name.
	//
	//	?service ses:refersTo ?aspect;
	//	ses:exposesFunction ?function.
	//
	//	?function rdf:type ses:MeasuringFunction.
	//}

	return url.QueryEscape(
		model.PREFIX_SES +
			model.PREFIX_RDF +
			"construct {" +
			"?aspect rdf:type ses:Aspect;" +
			"rdfs:label ?dc_name." +
			"} where {" +
			"?aspect rdf:type ses:Aspect;" +
			"rdfs:label ?dc_name." +
			"?service ses:refersTo ?aspect;" +
			"ses:exposesFunction ?function." +
			"?function rdf:type ses:MeasuringFunction." +
			"}")
}

func (*Database) getDeleteConceptWithNestedQuery(subject string) string {

	return url.QueryEscape("prefix x: <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> " +
		"delete { ?s ?p ?o } " +
		"where {<" + subject + "> (x:|!x:)* ?s . ?s ?p ?o . }")
}

func (*Database) getDeleteConceptWithouthNestedQuery(subject string) string {

	return url.QueryEscape(
		model.PREFIX_SES +
			model.PREFIX_RDF +
			"delete { " +
			"<" + subject + "> rdf:type ?type;" +
			"ses:hasBaseCharacteristic ?baseCharacteristic;" +
			"rdfs:label ?label." +
			"} where {" +
			"<" + subject + "> rdf:type ?type;" +
			"ses:hasBaseCharacteristic ?baseCharacteristic;" +
			"	rdfs:label ?label." +
			"}")
}

func (*Database) getDeleteCharacteristicQuery(subject string) string {
	return url.QueryEscape(
		"prefix x: <http://www.w3.org/1999/02/22-rdf-syntax-ns#type>" +
			"delete {?s ?p ?o .} where {" +
			"{" +
			"	<" + subject + "> (x:|!x:)* ?s ." +
			"		?s ?p ?o ." +
			"	}" +
			"	UNION" +
			"	{" +
			"		?s ?p ?o ." +
			"		?s ?p <" + subject + "> ." +
			"	}" +
			"}")
}

func (*Database) getDeleteCharacteristicConceptRelationQuery(characteristicId string) string {
	return url.QueryEscape(`delete where { ?s <` + model.SES_ONTOLOGY_HAS_CHARACTERISTIC + "> <" + characteristicId + `>}`)
}

func (*Database) getDeviceClassQuery(subject string, queryForm string) string {
	return url.QueryEscape(
		model.PREFIX_SES +
			model.PREFIX_RDF +
			queryForm + " {" +
			"<" + subject + "> rdf:type ?type;" +
			"rdfs:label ?label;" +
			"ses:hasImage ?image;" +
			"} where {" +
			"<" + subject + "> rdf:type ?type;" +
			"rdfs:label ?label." +
			"OPTIONAL {<" + subject + "> ses:hasImage ?image .}" +
			"}")
}

func (*Database) getLocationQuery(subject string, queryForm string) string {
	return url.QueryEscape(
		model.PREFIX_SES +
			model.PREFIX_RDF +
			queryForm + " {" +
			"<" + subject + "> rdf:type ?type;" +
			"rdfs:label ?label;" +
			"rdfs:comment ?comment;" +
			"ses:hasImage ?image;" +
			"ses:isLocationOfDevice ?device;" +
			"ses:isLocationOfDeviceGroup ?deviceGroup;" +
			"} where {" +
			"<" + subject + "> rdf:type ?type;" +
			"rdfs:label ?label." +
			"OPTIONAL {" +
			"<" + subject + "> ses:hasImage ?image ." +
			"}" +
			"OPTIONAL {" +
			"<" + subject + "> rdfs:comment ?comment ." +
			"}" +
			"OPTIONAL {" +
			"<" + subject + "> ses:isLocationOfDevice ?device ." +
			"}" +
			"OPTIONAL {" +
			"<" + subject + "> ses:isLocationOfDeviceGroup ?deviceGroup ." +
			"}" +
			"}")
}

func (*Database) getAspectQuery(subject string, queryForm string) string {

	return url.QueryEscape(
		model.PREFIX_SES +
			model.PREFIX_RDF +
			queryForm + " {" +
			"<" + subject + "> rdf:type ?type;" +
			"rdfs:label ?label;" +
			"} where {" +
			"<" + subject + "> rdf:type ?type;" +
			"rdfs:label ?label;" +
			"}")
}

func (*Database) getFunctionQuery(subject string, queryForm string) string {

	//PREFIX ses: <https://senergy.infai.org/ontology/>
	//PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
	//delete {
	//	<urn:infai:ses:measuring-function:321> rdf:type ?type;
	//	rdfs:label ?label;
	//	ses:hasConcept ?concept;
	//} where {
	//	<urn:infai:ses:measuring-function:321> rdf:type ?type;
	//	rdfs:label ?label;
	//	OPTIONAL {<urn:infai:ses:measuring-function:321> ses:hasConcept ?concept .}
	//}

	return url.QueryEscape(
		model.PREFIX_SES +
			model.PREFIX_RDF +
			queryForm + " {" +
			"<" + subject + "> rdf:type ?type;" +
			"rdfs:label ?label;" +
			"rdfs:comment ?comment;" +
			"ses:hasConcept ?concept;" +
			"ses:displayName ?displayName;" +
			"} where {" +
			"<" + subject + "> rdf:type ?type;" +
			"rdfs:label ?label;" +
			"OPTIONAL {<" + subject + "> ses:hasConcept ?concept .}" +
			"OPTIONAL {<" + subject + "> rdfs:comment ?comment .}" +
			"OPTIONAL {<" + subject + "> ses:displayName ?displayName .}" +
			"}")
}
