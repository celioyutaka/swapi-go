package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

/*
IMPORTANT
For every new configuration you should add in Configuration Struct ALWAYS as string
*/
type Configuration struct {
	SERVER_PORT    string
	DATABASE       string
	MYSQL_HOST     string
	MYSQL_USER     string
	MYSQL_PASSWORD string
	MYSQL_DATABASE string
	MYSQL_PORT     string

	MONGODB_HOST       string
	MONGODB_USER       string
	MONGODB_PASSWORD   string
	MONGODB_DATABASE   string
	MONGODB_PORT       string
	MONGODB_CONNECTION string

	SWAPI_URL string

	TEST_YODA string
}

/*
Use reflection to get property dynamically
*/
func getField(v *Configuration, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}

/*
Get environment config
*/
func GetEnv(key string) string {
	filename := "config/config.json"
	//In tests, the path is considered the subdir of the file (config/config_test.go, to fix, add ../ before path)
	if _, err := os.Stat("config/config.json"); os.IsNotExist(err) {
		filename = "../config/config.json"
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		log.Fatalln("You should create config.json file in /config directory")

	}
	byteValue, _ := ioutil.ReadAll(file)
	var value Configuration
	json.Unmarshal(byteValue, &value)

	file.Close()
	return getField(&value, key)
}
