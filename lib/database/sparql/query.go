package sparql

import "net/url"

func (*Database) getConstructQuery(s string, p string, o string) (string) {
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
