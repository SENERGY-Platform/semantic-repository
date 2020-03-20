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

func (*Database) getDeviceTypeQuery(deviceTypeId string, deviceClassId string, functionIds []string, aspectIds []string) string {

	if deviceTypeId == "" {
		deviceTypeId = " ?devicetype "
	} else {
		deviceTypeId = " <" + deviceTypeId + "> "
	}

	if deviceClassId == "" {
		deviceClassId = " ?deviceclass "
	} else {
		deviceClassId = " <" + deviceClassId + "> "
	}

	functionFilter := ""
	if len(functionIds) > 0 {
		functionFilter = " FILTER (?function IN ("
		for index, functionId := range functionIds {
			functionFilter = functionFilter + "<" + functionId + ">"
			if index < len(functionIds)-1 {
				functionFilter = functionFilter + ","
			}
		}
		functionFilter = functionFilter + ")) "
	}

	aspectFilter := ""
	if len(aspectIds) > 0 {
		aspectFilter = " FILTER (?aspect IN ("
		for index, aspectId := range aspectIds {
			aspectFilter = aspectFilter + "<" + aspectId + ">"
			if index < len(aspectIds)-1 {
				aspectFilter = aspectFilter + ","
			}
		}
		aspectFilter = aspectFilter + ")) "
	}

	// Example Devicetype

	//PREFIX ses: <https://senergy.infai.org/ontology/>
	//PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
	//construct {
	//	<urn:infai:ses:device-type:8cb7bb2c-e661-42a3-b7d1-5e9fc6d60152>
	//	rdf:type ?type;
	//	rdfs:label ?label;
	//	ses:hasDeviceClass ?deviceclass;
	//	ses:hasService ?service .
	//	?service rdf:type ?s_type;
	//	rdfs:label ?s_label;
	//	ses:refersTo ?aspect;
	//	ses:exposesFunction ?function.
	//	?function rdfs:label ?f_label;
	//	rdf:type ?f_type.
	//	?function ses:hasConcept ?concept_id.
	//	?aspect rdfs:label ?a_label;
	//	rdf:type ?a_type.
	//	?deviceclass rdfs:label ?dc_label;
	//	rdf:type ?dc_type.
	//} where {
	//	<urn:infai:ses:device-type:8cb7bb2c-e661-42a3-b7d1-5e9fc6d60152>
	//	rdf:type ?type;
	//	rdfs:label ?label;
	//	ses:hasDeviceClass ?deviceclass;
	//	ses:hasService ?service .
	//	?service rdf:type ?s_type;
	//	rdfs:label ?s_label;
	//	ses:refersTo ?aspect;
	//	ses:exposesFunction ?function.
	//	?function rdfs:label ?f_label;
	//	rdf:type ?f_type.
	//	OPTIONAL {?function ses:hasConcept ?concept_id.}
	//	?aspect rdfs:label ?a_label;
	//	rdf:type ?a_type.
	//	?deviceclass rdfs:label ?dc_label;
	//	rdf:type ?dc_type.
	//}

	// Example Deviceclass & Function

	//PREFIX ses: <https://senergy.infai.org/ontology/>
	//PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
	//construct {
	//	?devicetype rdf:type ?type;
	//	rdfs:label ?label;
	//	ses:hasDeviceClass  <urn:infai:ses:deviceclass:2e2e> ;
	//	ses:hasService ?service .
	//
	//	?service rdf:type ?s_type;
	//	rdfs:label ?s_label;
	//	ses:refersTo ?aspect;
	//	ses:exposesFunction ?function.
	//
	//	?function rdfs:label ?f_label;
	//	rdf:type ?f_type.
	//
	//	?function ses:hasConcept ?concept_id.
	//
	//	?aspect rdfs:label ?a_label;
	//	rdf:type ?a_type.
	//
	//	<urn:infai:ses:deviceclass:2e2e>  rdfs:label ?dc_label;
	//	rdf:type ?dc_type.
	//}
	//where {
	//
	//	?devicetype rdf:type ?type;
	//	rdfs:label ?label;
	//	ses:hasDeviceClass  <urn:infai:ses:deviceclass:2e2e> ;
	//	ses:hasService ?service .
	//
	//	?service rdf:type ?s_type;
	//	rdfs:label ?s_label;
	//	ses:refersTo ?aspect;
	//	ses:exposesFunction ?function.
	//
	//	?function rdfs:label ?f_label;
	//	rdf:type ?f_type.
	//	FILTER (?function IN (<urn:infai:ses:function:5e5e-1>) )
	//	OPTIONAL {?function ses:hasConcept ?concept_id.}
	//
	//	?aspect rdfs:label ?a_label;
	//	rdf:type ?a_type.
	//
	//	<urn:infai:ses:deviceclass:2e2e>  rdfs:label ?dc_label;
	//	rdf:type ?dc_type.}

	return url.QueryEscape(
		model.PREFIX_SES +
			model.PREFIX_RDF +
			"construct {" +
			deviceTypeId +
			"rdf:type ?type;" +
			"rdfs:label ?label;" +
			"ses:hasDeviceClass " + deviceClassId + ";" +
			"ses:hasService ?service ." +
			"?service rdf:type ?s_type;" +
			"rdfs:label ?s_label;" +
			"ses:refersTo ?aspect;" +
			"ses:exposesFunction ?function." +

			"?function rdfs:label ?f_label;" +
			"rdf:type ?f_type." +
			"?function ses:hasConcept ?concept_id." +
			"?aspect rdfs:label ?a_label;" +
			"rdf:type ?a_type." +
			deviceClassId + " rdfs:label ?dc_label;" +
			"rdf:type ?dc_type." +
			"} where {" +
			deviceTypeId +
			"rdf:type ?type;" +
			"rdfs:label ?label;" +
			"ses:hasDeviceClass " + deviceClassId + ";" +
			"ses:hasService ?service ." +
			"?service rdf:type ?s_type;" +
			"rdfs:label ?s_label;" +
			"ses:refersTo ?aspect;" +
			"ses:exposesFunction ?function." +

			"?function rdfs:label ?f_label;" +
			"rdf:type ?f_type." +
			functionFilter +
			"OPTIONAL {?function " +
			"ses:hasConcept ?concept_id." +
			"}" +
			"?aspect rdfs:label ?a_label;" +
			"rdf:type ?a_type." +
			aspectFilter +
			deviceClassId + " rdfs:label ?dc_label;" +
			"rdf:type ?dc_type." +
			"}")
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
