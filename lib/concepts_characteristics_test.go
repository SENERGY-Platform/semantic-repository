package lib

import (
	"context"
	"encoding/json"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/controller"
	"github.com/SENERGY-Platform/semantic-repository/lib/database"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/SENERGY-Platform/semantic-repository/lib/testutil"
	"reflect"
	"sort"
	"sync"
	"testing"
)

func TestConceptCharacteristic(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	defer wg.Wait()
	defer cancel()
	err = testutil.GetDockerEnv(ctx, &wg, &conf)
	if err != nil {
		t.Error(err)
		return
	}
	db, err := database.New(conf)
	if err != nil {
		t.Error(err)
		return
	}

	ctrl, err := controller.New(conf, db)
	if err != nil {
		t.Error(err)
		return
	}

	concept := model.Concept{
		Id:                "urn:ses:infai:concept:c",
		Name:              "cn",
		CharacteristicIds: nil,
	}

	expectedConcept := model.Concept{
		Id:                "urn:ses:infai:concept:c",
		Name:              "cn",
		CharacteristicIds: []string{"urn:ses:infai:characteristicch1", "urn:ses:infai:characteristicch2"},
		RdfType:           "https://senergy.infai.org/ontology/Concept",
	}

	characterisitc1 := model.Characteristic{
		Id:   "urn:ses:infai:characteristicch1",
		Name: "chn",
		Type: model.Boolean,
	}
	characterisitc2 := model.Characteristic{
		Id:   "urn:ses:infai:characteristicch2",
		Name: "chn",
		Type: model.Boolean,
	}

	err = ctrl.SetConcept(concept, "owner")
	if err != nil {
		t.Error(err)
		return
	}
	err = ctrl.SetCharacteristic("urn:ses:infai:concept:c", characterisitc1, "owner")
	if err != nil {
		t.Error(err)
		return
	}
	err = ctrl.SetCharacteristic("urn:ses:infai:concept:c", characterisitc2, "owner")
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("check after init", checkConcept(ctrl, "urn:ses:infai:concept:c", expectedConcept))

	err = ctrl.SetConcept(concept, "owner")
	if err != nil {
		t.Error(err)
		return
	}
	t.Run("check after reset concept", checkConcept(ctrl, "urn:ses:infai:concept:c", expectedConcept))

	err = ctrl.SetCharacteristic("urn:ses:infai:concept:c", characterisitc1, "owner")
	if err != nil {
		t.Error(err)
		return
	}
	t.Run("check after reset characteristic", checkConcept(ctrl, "urn:ses:infai:concept:c", expectedConcept))

}

func checkConcept(ctrl *controller.Controller, id string, expected model.Concept) func(t *testing.T) {
	return func(t *testing.T) {
		result, err, _ := ctrl.GetConceptWithoutCharacteristics(id)
		if err != nil {
			t.Error(err, result)
			return
		}
		sort.Strings(result.CharacteristicIds)
		sort.Strings(expected.CharacteristicIds)
		if !reflect.DeepEqual(result, expected) {
			resultJson, _ := json.Marshal(result)
			expectedJson, _ := json.Marshal(expected)
			t.Error(string(resultJson), string(expectedJson))
		}
	}
}
