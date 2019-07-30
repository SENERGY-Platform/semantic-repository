package sparql

import "net/url"

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
	return url.QueryEscape("construct {?s ?p ?o} where {" + "?s ?p ?o ." +s +p +o +".}")
}


func (*Database) getSubjectWithAllPropertiesQuery(subject string) (string) {

	return url.QueryEscape("prefix x: <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> " +
		"construct { ?s ?p ?o } " +
		"where {<" + subject + "> (x:|!x:)* ?s . ?s ?p ?o . }")
}

