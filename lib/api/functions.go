package api

import (
	"encoding/json"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/SmartEnergyPlatform/jwt-http-router"
	"log"
	"net/http"
)

func init() {
	endpoints = append(endpoints, FunctionsEndpoints)
}

func FunctionsEndpoints(config config.Config, control Controller, router *jwt_http_router.Router) {
	controllingResource := "/controlling-functions"
	measuringResource := "/measuring-functions"

	router.GET(controllingResource, func(writer http.ResponseWriter, request *http.Request, params jwt_http_router.Params, jwt jwt_http_router.Jwt) {
		result, err, errCode := control.GetFunctions(model.SES_ONTOLOGY_CONTROLLING_FUNCTION)
		if err != nil {
			http.Error(writer, err.Error(), errCode)
			return
		}
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		err = json.NewEncoder(writer).Encode(result)
		if err != nil {
			log.Println("ERROR: unable to encode response", err)
		}
		return
	})

	router.GET(measuringResource, func(writer http.ResponseWriter, request *http.Request, params jwt_http_router.Params, jwt jwt_http_router.Jwt) {
		result, err, errCode := control.GetFunctions(model.SES_ONTOLOGY_MEASURING_FUNCTION)
		if err != nil {
			http.Error(writer, err.Error(), errCode)
			return
		}
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		err = json.NewEncoder(writer).Encode(result)
		if err != nil {
			log.Println("ERROR: unable to encode response", err)
		}
		return
	})

}