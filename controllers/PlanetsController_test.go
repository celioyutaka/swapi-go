package controllers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"swapi-go/database"
	"swapi-go/models"

	"github.com/gorilla/mux"
)

func createTempPlanet() string {

	planets_db := database.PlanetsDB{}
	planet := models.Planet{Name: "Tatooine", Climate: "arid", Terrain: "desert"}
	return planets_db.Insert(planet)
}
func TestSetDefault(t *testing.T) {
	planets_db := database.PlanetsDB{}
	planets_db.ResetDatabase()
}

func TestIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/planets", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	planetsController := PlanetsController{}
	handler := http.HandlerFunc(planetsController.Index)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `"success": `
	//log.Println(rr.Body.String())
	if strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Returned unexpected body: got %v want success:true", expected)
	}

}

func TestCreate(t *testing.T) {
	var jsonStr = []byte(`{"Name": "Tatooine","Climate": "arid","Terrain": "desert"}`)
	req, err := http.NewRequest("POST", "/api/planet", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	planetsController := PlanetsController{}
	handler := http.HandlerFunc(planetsController.Create)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `"success":false`
	//log.Println(rr.Body.String())
	if strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Returned unexpected body: got %v want success:true", expected)
	}

}

func TestRead(t *testing.T) {
	id := createTempPlanet()
	request_url := "/api/planet/" + id
	req, err := http.NewRequest("GET", request_url, nil)

	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	planetsController := PlanetsController{}
	handler := http.HandlerFunc(planetsController.Read)

	//bypass to work with mux
	vars := map[string]string{
		"id": id,
	}
	req = mux.SetURLVars(req, vars)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `"success":false`
	//log.Println(rr.Body.String())
	if strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Returned unexpected body: got %v want success:true", expected)
	}
}

func TestUpdate(t *testing.T) {
	var jsonStr = []byte(`{"Name": "Hoth","Climate": "temperature","Terrain": "ocean"}`)
	id := createTempPlanet()
	request_url := "/api/planet/" + id

	req, err := http.NewRequest("PUT", request_url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	planetsController := PlanetsController{}
	handler := http.HandlerFunc(planetsController.Update)

	//bypass to work with mux
	vars := map[string]string{
		"id": id,
	}
	req = mux.SetURLVars(req, vars)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `"success":false`
	//log.Println(rr.Body.String())
	if strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Returned unexpected body: got %v want success:true", expected)
	}
}

func TestReadName(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/planet/search/Hoth", nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	planetsController := PlanetsController{}
	handler := http.HandlerFunc(planetsController.Read)

	//bypass to work with mux
	vars := map[string]string{
		"name": "Hoth",
	}
	req = mux.SetURLVars(req, vars)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `"success":false`
	//log.Println(rr.Body.String())
	if strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Returned unexpected body: got %v want success:true", expected)
	}
}

func TestDelete(t *testing.T) {
	id := createTempPlanet()
	request_url := "/api/planet/" + id
	req, err := http.NewRequest("DELETE", request_url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	planetsController := PlanetsController{}
	handler := http.HandlerFunc(planetsController.Delete)

	//bypass to work with mux
	vars := map[string]string{
		"id": id,
	}
	req = mux.SetURLVars(req, vars)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `"success":false`
	log.Println(rr.Body.String())
	if strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Returned unexpected body: got %v want success:true", expected)
	}
}
func TestListByName(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/planets/Tatooine", nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	planetsController := PlanetsController{}
	handler := http.HandlerFunc(planetsController.ListByName)

	//bypass to work with mux
	vars := map[string]string{
		"id": "Tatooine",
	}
	req = mux.SetURLVars(req, vars)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `"success":false`
	//log.Println(rr.Body.String())
	if strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Returned unexpected body: got %v want success:true", expected)
	}
}
