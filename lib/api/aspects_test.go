package api

import (
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/database"
	"testing"
)

func TestInsertAspects(t *testing.T) {
	conf, err := config.Load("../../config.json")
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.New(conf)
	success, err := db.InsertData(
		`<urn:infai:ses:aspects:444> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <https://senergy.infai.org/ontology/Aspect> .
<urn:infai:ses:aspects:444> <http://www.w3.org/2000/01/rdf-schema#label> "Air" .`)
	t.Log(success)

	success, err = db.InsertData(
		`<urn:infai:ses:aspects:222> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <https://senergy.infai.org/ontology/Aspect> .
<urn:infai:ses:aspects:222> <http://www.w3.org/2000/01/rdf-schema#label> "Lightning" .`)
	t.Log(success)
}
