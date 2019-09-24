package sparql

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"net/url"
)

func (*Database) getConstructListWithoutSubProperties(p string, o string) (string) {
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

func (*Database) getConstructWithAllSubProperties(subject string) (string) {

	return url.QueryEscape("prefix x: <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> " +
		"construct { ?s ?p ?o } " +
		"where {<" + subject + "> (x:|!x:)* ?s . ?s ?p ?o . }")
}

func (*Database) getConstructWithoutSubProperties(subject string) (string) {
	//construct {<urn:ses:infai:concept:1a1a1a> ?p ?o .}
	//where {<urn:ses:infai:concept:1a1a1a> ?p ?o .}
	return url.QueryEscape(
		"construct { <" + subject + "> ?p ?o .} " +
		"where {<" + subject + "> ?p ?o .}")
}

func (*Database) getDeleteDeviceTypeQuery(subject string) (string) {

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

func (*Database) getDeviceTypeQuery(subject string) (string) {

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
	//
	//	?function rdfs:label ?f_label;
	//	rdf:type ?f_type;
	//	ses:hasConcept ?concept_id.
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
	//
	//	?function rdfs:label ?f_label;
	//	rdf:type ?f_type;
	//	ses:hasConcept ?concept_id.
	//	?aspect rdfs:label ?a_label;
	//	rdf:type ?a_type.
	//	?deviceclass rdfs:label ?dc_label;
	//	rdf:type ?dc_type.
	//}

	return url.QueryEscape(
		model.PREFIX_SES +
			model.PREFIX_RDF +
			"construct {" +
			"<" + subject + ">" +
			"rdf:type ?type;" +
			"rdfs:label ?label;" +
			"ses:hasDeviceClass ?deviceclass;" +
			"ses:hasService ?service ." +
			"?service rdf:type ?s_type;" +
			"rdfs:label ?s_label;" +
			"ses:refersTo ?aspect;" +
			"ses:exposesFunction ?function." +

			"?function rdfs:label ?f_label;" +
			"rdf:type ?f_type;" +
			"ses:hasConcept ?concept_id." +
			"?aspect rdfs:label ?a_label;" +
			"rdf:type ?a_type." +
			"?deviceclass rdfs:label ?dc_label;" +
			"rdf:type ?dc_type." +
			"} where {" +
			"<" + subject + ">" +
			"rdf:type ?type;" +
			"rdfs:label ?label;" +
			"ses:hasDeviceClass ?deviceclass;" +
			"ses:hasService ?service ." +
			"?service rdf:type ?s_type;" +
			"rdfs:label ?s_label;" +
			"ses:refersTo ?aspect;" +
			"ses:exposesFunction ?function." +

			"?function rdfs:label ?f_label;" +
			"rdf:type ?f_type;" +
			"ses:hasConcept ?concept_id." +
			"?aspect rdfs:label ?a_label;" +
			"rdf:type ?a_type." +
			"?deviceclass rdfs:label ?dc_label;" +
			"rdf:type ?dc_type." +
			"}")
}

func (*Database) getDeleteConceptWithNestedQuery(subject string) (string) {

	return url.QueryEscape("prefix x: <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> " +
		"delete { ?s ?p ?o } " +
		"where {<" + subject + "> (x:|!x:)* ?s . ?s ?p ?o . }")
}

func (*Database) getDeleteConceptWithouthNestedQuery(subject string) (string) {

	return url.QueryEscape(
		model.PREFIX_SES +
			model.PREFIX_RDF +
			"delete { " +
			"<" + subject + "> rdf:type ?type;" +
			"rdfs:label ?label." +
			"<" + subject + "> ses:hasCharacteristic ?characteristic ." +
			"} where {" +
			"<" + subject + "> rdf:type ?type;" +
			"	rdfs:label ?label." +
			" OPTIONAL {<" + subject + "> ses:hasCharacteristic ?characteristic .}" +
			"}")
}

func (*Database) getDeleteCharacteristicQuery(subject string) (string) {

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
