package mysql

import (
	"swapi-go/models"

	_ "github.com/go-sql-driver/mysql"
)

type PlanetsMySQL struct {
}

func (planets_mysql *PlanetsMySQL) ResetDatabase() {
	db := Connect()
	db.Query("TRUNCATE planets")
	db.Query("INSERT INTO `planets` (`Id`, `Name`, `Climate`, `Terrain`, `ApperancesFilms`) VALUES (1, 'Hoth', 'temperature', 'ocean', 1);")

}
func (planets_mysql *PlanetsMySQL) SearchByName(name string) (models.Planet, bool) {
	db := Connect()
	var planet models.Planet

	stringSQL := "SELECT Id, Name, Climate, Terrain, ApperancesFilms FROM planets WHERE Name = ?"
	err := db.QueryRow(stringSQL, name).Scan(&planet.Id, &planet.Name, &planet.Climate, &planet.Terrain, &planet.ApperancesFilms)

	if err != nil {
		return models.Planet{}, false
		//panic(err.Error())
	}
	return planet, true
}
func (planets_mysql *PlanetsMySQL) SearchById(id string) (models.Planet, bool) {
	db := Connect()
	var planet models.Planet

	stringSQL := "SELECT Id, Name, Climate, Terrain, ApperancesFilms FROM planets WHERE Id = ?"
	err := db.QueryRow(stringSQL, id).Scan(&planet.Id, &planet.Name, &planet.Climate, &planet.Terrain, &planet.ApperancesFilms)

	if err != nil {
		return models.Planet{}, false
		//panic(err.Error())
	}
	return planet, true
}

func (planets_mysql *PlanetsMySQL) ListAll() []models.Planet {
	db := Connect()
	results, err := db.Query("SELECT Id, Name, Climate, Terrain, ApperancesFilms FROM planets")
	if err != nil {
		panic(err.Error())
	}
	var planets []models.Planet
	for results.Next() {
		var planet models.Planet
		err = results.Scan(&planet.Id, &planet.Name, &planet.Climate, &planet.Terrain, &planet.ApperancesFilms)
		if err != nil {
			panic(err.Error())
		}
		planets = append(planets, planet)
	}
	return planets
}

func (planets_mysql *PlanetsMySQL) Insert(p models.Planet) int64 {
	db := Connect()
	statement, err := db.Prepare("INSERT INTO planets  (Name, Climate, Terrain, ApperancesFilms) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	result, err := statement.Exec(p.Name, p.Climate, p.Terrain, p.ApperancesFilms)

	if err != nil {
		panic(err.Error())
	}

	//get last insert id
	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	return id
}
func (planets_mysql *PlanetsMySQL) UpdateById(id string, p models.Planet) int64 {
	db := Connect()
	statement, err := db.Prepare("UPDATE planets SET Name = ?, Climate = ?, Terrain = ?, ApperancesFilms = ? WHERE Id = ?")
	if err != nil {
		panic(err.Error())
	}
	result, err := statement.Exec(p.Name, p.Climate, p.Terrain, p.ApperancesFilms, p.Id)

	if err != nil {
		panic(err.Error())
	}

	//get number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	return rowsAffected
}

func (planets_mysql *PlanetsMySQL) DeleteById(id string) int64 {
	db := Connect()
	statement, err := db.Prepare("DELETE FROM planets WHERE Id = ?")
	if err != nil {
		panic(err.Error())
	}
	result, err := statement.Exec(id)

	if err != nil {
		panic(err.Error())
	}

	//get number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	return rowsAffected
}

func (planets_mysql *PlanetsMySQL) ListByName(name string) ([]models.Planet, bool) {
	db := Connect()
	var planets []models.Planet

	stringSQL := "SELECT Id, Name, Climate, Terrain, ApperancesFilms FROM planets WHERE Name like ?"
	results, err := db.Query(stringSQL, "%"+name+"%")
	if err != nil {
		return planets, false
	}

	for results.Next() {
		var planet models.Planet
		err = results.Scan(&planet.Id, &planet.Name, &planet.Climate, &planet.Terrain, &planet.ApperancesFilms)
		if err != nil {
			panic(err.Error())
		}
		planets = append(planets, planet)
		//log.Println(planet.Name)
	}
	return planets, true
}
