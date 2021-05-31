package mysql

import (
	"database/sql"
	"swapi-go/config"
)

func Connect() *sql.DB {
	host := config.GetEnv("MYSQL_HOST")
	user := config.GetEnv("MYSQL_USER")
	password := config.GetEnv("MYSQL_PASSWORD")
	database := config.GetEnv("MYSQL_DATABASE")
	port := config.GetEnv("MYSQL_PORT")
	connection_string := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database
	//fmt.Println("connection_string: " + connection_string)
	//example connection string user:password@tcp(127.0.0.1:3306)/swapi

	db, err := sql.Open("mysql", connection_string)

	if err != nil {
		panic(err.Error())
	}
	//close connection
	//defer db.Close()

	return db

}
