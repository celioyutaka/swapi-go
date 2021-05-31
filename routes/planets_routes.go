package routes

import (
	"swapi-go/controllers"

	"github.com/gorilla/mux"
)

func RequestHandlerPlanet(router *mux.Router) *mux.Router {
	//route of planets
	planetsController := controllers.PlanetsController{}
	//get all
	router.HandleFunc("/api/planets", planetsController.Index).Methods("GET")
	//create planet
	router.HandleFunc("/api/planet", planetsController.Create).Methods("POST")
	//get planet by name
	router.HandleFunc("/api/planet/search/{name}", planetsController.SearchByName).Methods("GET")
	//get planet by id
	router.HandleFunc("/api/planet/{id}", planetsController.Read).Methods("GET")
	//update planet by id
	router.HandleFunc("/api/planet/{id}", planetsController.Update).Methods("PUT")
	//delete planet by id
	router.HandleFunc("/api/planet/{id}", planetsController.Delete).Methods("DELETE")

	//get all planets contains 'name'
	router.HandleFunc("/api/planets/{name}", planetsController.ListByName).Methods("GET")
	//get planet with specific name
	//router.HandleFunc("/api/planets/{name}", planetsController.SearchByName).Methods("GET")

	return router
}
