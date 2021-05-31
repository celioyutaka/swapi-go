package mongodb

import (
	"context"
	"log"
	"swapi-go/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlanetsMongoDB struct {
}

func (planets_mongodb *PlanetsMongoDB) Insert(p models.Planet) string {
	db, _ := Connect()
	_, err := db.Collection("planets").InsertOne(context.TODO(), p)
	if err != nil {
		panic(err.Error())
	}

	doc := db.Collection("planets").FindOne(context.TODO(), bson.M{}, options.FindOne().SetSort(bson.M{"$natural": -1}))
	var planet models.Planet
	doc.Decode(&planet)

	log.Println(planet)

	//get last insert id
	id := planet.Id
	//id := oid[:]
	if err != nil {
		panic(err.Error())
	}

	return id
}

func (planets_mongodb *PlanetsMongoDB) ListAll() []models.Planet {
	db, ctx := Connect()
	var planets []models.Planet
	filter := bson.D{{}}
	results, err := db.Collection("planets").Find(context.TODO(), filter)
	if err != nil {
		log.Println("Erro planets.Find")
		panic(err.Error())
	}

	for results.Next(ctx) {
		var planet models.Planet
		err := results.Decode(&planet)
		if err != nil {
			panic(err.Error())
		}
		planets = append(planets, planet)
	}
	return planets
}

func (planets_mongodb *PlanetsMongoDB) SearchByFilter(filter interface{}) (models.Planet, bool) {
	db, _ := Connect()

	doc := db.Collection("planets").FindOne(context.TODO(), filter, options.FindOne())
	var planet models.Planet
	doc.Decode(&planet)

	return planet, true
}

func (planets_mongodb *PlanetsMongoDB) UpdateById(id string, p models.Planet) int64 {
	db, _ := Connect()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}
	p.Id = id
	updateQuery := bson.M{
		"$set": bson.M{
			"name":    p.Name,
			"terrain": p.Terrain,
			"climate": p.Climate,
		},
	}
	result, errUpdate := db.Collection("planets").UpdateOne(context.TODO(), bson.M{"_id": objectId}, updateQuery, options.Update())
	if errUpdate != nil {
		log.Println(errUpdate)
	}

	//get number of rows affected
	rowsAffected := result.ModifiedCount
	return rowsAffected
}

func (planets_mongodb *PlanetsMongoDB) DeleteById(id string) int64 {
	db, _ := Connect()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}
	result, errUpdate := db.Collection("planets").DeleteOne(context.TODO(), bson.M{"_id": objectId}, options.Delete())
	if errUpdate != nil {
		log.Println(errUpdate)
	}

	//get number of rows affected
	rowsAffected := result.DeletedCount
	return rowsAffected
}
