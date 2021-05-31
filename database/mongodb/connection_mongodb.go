package mongodb

import (
	"context"
	"log"
	"swapi-go/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Database, context.Context) {
	host := config.GetEnv("MONGODB_HOST")
	//user := config.GetEnv("MONGODB_USER")
	//password := config.GetEnv("MONGODB_PASSWORD")
	database := config.GetEnv("MONGODB_DATABASE")
	port := config.GetEnv("MONGODB_PORT")
	connection := config.GetEnv("MONGODB_CONNECTION")

	/*
	   Connect to my cluster
	*/
	connection_string := "mongodb://" + host + ":" + port
	if len(connection) > 1 {
		connection_string = connection
	}

	//log.Println(connection_string)
	client, err := mongo.NewClient(options.Client().ApplyURI(connection_string))
	if err != nil {
		log.Println("Fail - mongo.NewClient")
		log.Fatal(err)
	}

	/* if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Println("Fail - mongo.Ping")
		log.Fatal(err)
	} */

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Println("Fail - client.Connect")
		log.Fatal(err)
	}

	return client.Database(database), ctx

	/*
	   List databases
	*/
	/* databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases) */

}
