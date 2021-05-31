package database

import (
	"log"
	"strconv"
	"swapi-go/config"
	"swapi-go/database/mongodb"
	"swapi-go/database/mysql"
	"swapi-go/models"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
"Interface" for database,
I started using MySQL then implemented in MongoDB

Choose between MySQL and MongoDB in config/config.json file in "DATABASE"
*/

type PlanetsDB struct {
}

func (planets_db *PlanetsDB) ResetDatabase() {
	switch config.GetEnv("DATABASE") {
	case "mysql":
		planets_mysql := mysql.PlanetsMySQL{}
		planets_mysql.ResetDatabase()
	}
	planets_mysql := mysql.PlanetsMySQL{}
	planets_mysql.ResetDatabase()
}

func (planets_db *PlanetsDB) SearchById(id string) (models.Planet, bool) {
	switch config.GetEnv("DATABASE") {
	case "mysql":
		planets_mysql := mysql.PlanetsMySQL{}
		return planets_mysql.SearchById(id)
	case "mongodb":
		planets_mongodb := mongodb.PlanetsMongoDB{}
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Println(err)
		}
		filter := bson.M{"_id": objectId}
		return planets_mongodb.SearchByFilter(filter)

	}
	planets_mysql := mysql.PlanetsMySQL{}
	return planets_mysql.SearchById(id)

}

func (planets_db *PlanetsDB) ListAll() []models.Planet {
	switch config.GetEnv("DATABASE") {
	case "mysql":
		planets_mysql := mysql.PlanetsMySQL{}
		return planets_mysql.ListAll()
	case "mongodb":
		planets_mongodb := mongodb.PlanetsMongoDB{}
		return planets_mongodb.ListAll()
	}
	planets_mysql := mysql.PlanetsMySQL{}
	return planets_mysql.ListAll()
}

func (planets_db *PlanetsDB) Insert(p models.Planet) string {
	switch config.GetEnv("DATABASE") {
	case "mysql":
		planets_mysql := mysql.PlanetsMySQL{}
		id := planets_mysql.Insert(p)
		strId := strconv.FormatInt(id, 10)
		return strId
	case "mongodb":
		planets_mongodb := mongodb.PlanetsMongoDB{}
		id := planets_mongodb.Insert(p)
		return id

	}
	planets_mysql := mysql.PlanetsMySQL{}
	id := planets_mysql.Insert(p)
	strId := strconv.FormatInt(id, 10)
	return strId
}

func (planets_db *PlanetsDB) SearchByName(name string) (models.Planet, bool) {
	log.Println("SearchByName: " + name)
	switch config.GetEnv("DATABASE") {
	case "mysql":
		planets_mysql := mysql.PlanetsMySQL{}
		return planets_mysql.SearchByName(name)
	case "mongodb":
		planets_mongodb := mongodb.PlanetsMongoDB{}
		filter := bson.M{"name": name}
		return planets_mongodb.SearchByFilter(filter)

	}
	planets_mysql := mysql.PlanetsMySQL{}
	return planets_mysql.SearchByName(name)

}
func (planets_db *PlanetsDB) UpdateById(id string, p models.Planet) int64 {
	switch config.GetEnv("DATABASE") {
	case "mysql":
		planets_mysql := mysql.PlanetsMySQL{}
		return planets_mysql.UpdateById(id, p)
	case "mongodb":
		planets_mongodb := mongodb.PlanetsMongoDB{}
		return planets_mongodb.UpdateById(id, p)
	}
	planets_mysql := mysql.PlanetsMySQL{}
	return planets_mysql.UpdateById(id, p)
}

func (planets_db *PlanetsDB) DeleteById(id string) int64 {
	switch config.GetEnv("DATABASE") {
	case "mysql":
		planets_mysql := mysql.PlanetsMySQL{}
		return planets_mysql.DeleteById(id)
	case "mongodb":
		planets_mongodb := mongodb.PlanetsMongoDB{}
		return planets_mongodb.DeleteById(id)
	}
	planets_mysql := mysql.PlanetsMySQL{}
	return planets_mysql.DeleteById(id)
}
func (planets_db *PlanetsDB) ListByName(name string) ([]models.Planet, bool) {
	switch config.GetEnv("DATABASE") {
	case "mysql":
		planets_mysql := mysql.PlanetsMySQL{}
		return planets_mysql.ListByName(name)
	}
	planets_mysql := mysql.PlanetsMySQL{}
	return planets_mysql.ListByName(name)
}
