package main

import (
	"fmt"
	"log"
	"swapi-go/database/mongodb"
	routes "swapi-go/routes"
)

func main() {

	planets_mongodb := mongodb.PlanetsMongoDB{}
	planets := planets_mongodb.ListAll()
	log.Println(planets)

	//os.Exit(1)

	fmt.Println("")
	fmt.Println("EPISODE IV")
	fmt.Println("")
	fmt.Println("A NEW HOPE")
	fmt.Println("")
	fmt.Println("A long time ago in a galaxy far, far away...")
	fmt.Println("")
	fmt.Println("It is a COVID-19 pandemic period.")
	fmt.Println("Rebel spaceships bring the first vaccine \nagainst the evil Coronavirus.")
	fmt.Println("")
	fmt.Println("")

	routes.RequestHandler()

}
