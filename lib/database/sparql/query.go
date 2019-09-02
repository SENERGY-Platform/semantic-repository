package sparql

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"net/url"
)

func (*Database) getConstructQueryWithoutProperties(s string, p string, o string) (string) {
	if s == "" {
		s = " ?s "
	} else {
		s = " <" + s + "> "
	}

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
		s + p + o + "." +
		"?s <http://www.w3.org/2000/01/rdf-schema#label> ?label" +
		"}" +
		"ORDER BY ASC(?label)")
}

func (*Database) getSubjectWithAllPropertiesQuery(subject string) (string) {

	return url.QueryEscape("prefix x: <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> " +
		"construct { ?s ?p ?o } " +
		"where {<" + subject + "> (x:|!x:)* ?s . ?s ?p ?o . }")
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
