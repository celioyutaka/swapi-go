package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"swapi-go/database"
	"swapi-go/models"
	swapi "swapi-go/services"

	"github.com/gorilla/mux"
)

type PlanetsController struct {
}

func (p *PlanetsController) Index(w http.ResponseWriter, r *http.Request) {
	planets_db := database.PlanetsDB{}
	log.Println("ListAll Planets")
	SetHeaderAPI(w)
	planets := planets_db.ListAll()
	result := len(planets) > 0
	setResponse(w, planets, result)
}
func (p *PlanetsController) Create(w http.ResponseWriter, r *http.Request) {
	planets_db := database.PlanetsDB{}
	log.Println("Create Planet")
	SetHeaderAPI(w)
	var planet models.Planet
	_ = json.NewDecoder(r.Body).Decode(&planet)

	//request swapi.dev to get count of apperances in films
	apperancesFilms := swapi.GetApperancesFilms(planet.Name)
	planet.ApperancesFilms = apperancesFilms

	id := planets_db.Insert(planet)
	planet.Id = id
	result := len(id) > 0
	setResponse(w, planet, result)
}
func (p *PlanetsController) Read(w http.ResponseWriter, r *http.Request) {
	planets_db := database.PlanetsDB{}
	log.Println("Get Planet")
	SetHeaderAPI(w)
	params := mux.Vars(r)

	id := params["id"]
	planet, result := planets_db.SearchById(id)
	setResponse(w, planet, result)

	/*
		name := params["id"]
		planet, result := planets_db.SearchByName(name)
		setResponse(w, planet, result)
	*/

}
func (p *PlanetsController) SearchByName(w http.ResponseWriter, r *http.Request) {
	planets_db := database.PlanetsDB{}
	log.Println("Search planet by Name")
	SetHeaderAPI(w)
	params := mux.Vars(r)
	name := params["name"]
	planet, result := planets_db.SearchByName(name)
	setResponse(w, planet, result)

}
func (p *PlanetsController) Update(w http.ResponseWriter, r *http.Request) {
	planets_db := database.PlanetsDB{}
	log.Println("Update planet")
	SetHeaderAPI(w)
	params := mux.Vars(r)
	log.Println("Updating ID: ", params["id"])
	id := params["id"]

	var planet models.Planet
	_ = json.NewDecoder(r.Body).Decode(&planet)
	planet.Id = id
	rowsAffected := planets_db.UpdateById(id, planet)
	result := rowsAffected > 0
	planet, _ = planets_db.SearchById(id)
	setResponse(w, planet, result)

}
func (p *PlanetsController) Delete(w http.ResponseWriter, r *http.Request) {
	planets_db := database.PlanetsDB{}
	log.Println("Using Death Star to delete this planet")
	SetHeaderAPI(w)
	params := mux.Vars(r)
	id := params["id"]
	rowsAffected := planets_db.DeleteById(id)
	result := rowsAffected > 0
	setResponse(w, models.Planet{}, result)

}
func (p *PlanetsController) ListByName(w http.ResponseWriter, r *http.Request) {
	planets_db := database.PlanetsDB{}
	log.Println("List planets by search by name")
	SetHeaderAPI(w)
	params := mux.Vars(r)
	name := params["name"]
	planets, result := planets_db.ListByName(name)
	setResponse(w, planets, result)
}
