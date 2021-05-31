package routes

import (
	"log"
	"net/http"

	"swapi-go/config"

	"github.com/gorilla/mux"
)

func RequestHandler() {
	router := mux.NewRouter()

	/*if this file gets bigger, create files for each area
	ex:
		planets_routes.go
		people_routes.go
		species_routes.go
		vehicles_routes.go
	*/

	//planets routes handler
	router = RequestHandlerPlanet(router)

	//start server, specific port
	serverPort := config.GetEnv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(":"+serverPort, router))
	log.Println("Starting server in port: " + serverPort)
}
