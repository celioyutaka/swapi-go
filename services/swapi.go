package swapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"swapi-go/config"
	swapi "swapi-go/services/schema"
)

func GetRequest(url_params string) string {
	log.Println("Get Request SWAPI: " + url_params)
	url_swapi := config.GetEnv("SWAPI_URL")
	request, err := http.Get(url_swapi + url_params)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}
	stringBody := string(body)
	return stringBody
}

func GetPlanetsSWAPI(planetName string) swapi.PlanetsSWApi {
	log.Println("GetPlanetsSWAPI SWAPI: " + planetName)
	url_params := "planets/?search=" + planetName
	stringBody := GetRequest(url_params)

	var planetsSWApi swapi.PlanetsSWApi
	json.Unmarshal([]byte(stringBody), &planetsSWApi)
	return planetsSWApi

}

func GetApperancesFilms(planetName string) int {
	log.Println("GetApperancesFilms SWAPI: " + planetName)
	planetsSWApi := GetPlanetsSWAPI(planetName)
	for _, value := range planetsSWApi.Results {
		return len(value.Films)
	}
	return 0
}
