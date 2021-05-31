# swapi-go
StarWarsAPI - Rest API for Star Wars

## Features
### Planet
- List
- Create/Insert
- Update
- Delete
- Search By ID
- Search By Name
- List By Name
- Check and count in how many films this planet appeared (using swapi.dev)

### TO DO
- Database in MongoDB

## How to use
### Config
To configure server, server port, database connection, etc

Change in file **config/config.json**

Example, by default the server port is **8000**, you can change SERVER_PORT to **8080** or **9000**
```json
{
    "SERVER_PORT": "8000",

    "SWAPI_URL": "https://swapi.dev/api/",

    "DATABASE": "mongodb",

    "MYSQL_HOST": "localhost",
    "MYSQL_USER": "user",
    "MYSQL_PASSWORD": "password",
    "MYSQL_DATABASE": "swapi",
    "MYSQL_PORT": "3306",

    "MONGODB_HOST":"localhost",
    "MONGODB_USER":"",
    "MONGODB_PASSWORD":"",
    "MONGODB_DATABASE":"swapi",
    "MONGODB_PORT":"27017",
    "MONGODB_CONNECTION":"mongodb://127.0.0.1:27017/?readPreference=primary&appname=SWAPI&ssl=false&connect=direct",

    "TEST_YODA": "Tests you should do"
}
```

### Build
Build the project
```GO
go build
```
Run
```cmd
swapi-go.exe
```
**Ready to use!**


### Test
Test
```GO
go test ./...

or 

go test -v ./...
```

## Make a request
### CREATE/INSERT
**Request**
**POST** `localhost:8000/api/planet`
```json
{
    "Name": "Hoth",
    "Climate": "temperature",
    "Terrain": "ocean"
}
```
**Response**
```json
{
    "success": true,
    "data": {
        "Id": 10,
        "Name": "Hoth",
        "Climate": "temperature",
        "Terrain": "ocean",
        "ApperancesFilms": 1
    }
}
```

### READ - Search By Id
**Request**
**GET** `localhost:8000/api/planet/3`
```json
{
    "success": true,
    "data": {
        "Id": 3,
        "Name": "Coruscant",
        "Climate": "temperature",
        "Terrain": "cityscape",
        "ApperancesFilms": 4
    }
}
```

### READ - Search By Name
**Request**
**GET** `localhost:8000/api/planet/search/Coruscant`
```json
{
    "success": true,
    "data": {
        "Id": 3,
        "Name": "Coruscant",
        "Climate": "temperature",
        "Terrain": "cityscape",
        "ApperancesFilms": 4
    }
}
```


### UPDATE
**Request**
**PUT** `localhost:8000/api/planet/10`
```json
{
    "Name": "Hoth",
    "Climate": "temperature, ocean",
    "Terrain": "ocean"
}
```
**Response**
```json
{
    "success": true,
    "data": {
        "Id": 10,
        "Name": "Hoth",
        "Climate": "temperature, ocean",
        "Terrain": "ocean",
        "ApperancesFilms": 1
    }
}
```


### DELETE
**Request**
**DELETE** `localhost:8000/api/planet/10`

**Response**
```json
{
    "success": true,
    "data": {
        "Id": 0,
        "Name": "",
        "Climate": "",
        "Terrain": "",
        "ApperancesFilms": 0
    }
}
```

### LIST ALL
**Request**
**GET** `localhost:8000/api/planets`
```json
{
    "success": true,
    "data": [
        {
            "Id": 10,
            "Name": "Hoth",
            "Climate": "temperature",
            "Terrain": "ocean",
            "ApperancesFilms": 1
        },
        {
            "Id": 11,
            "Name": "Tatooine",
            "Climate": "arid",
            "Terrain": "desert",
            "ApperancesFilms": 5
        }
    ]
}
```



## MongoDB
Default database name: **swapi**

Default collection name: **planets**



## MySQL
If you using MySQL, this is the query to create database, tables, users
```SQL
CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON * . * TO 'user'@'localhost';
FLUSH PRIVILEGES

CREATE DATABASE IF NOT EXISTS `swapi`
USE `swapi`;

CREATE TABLE IF NOT EXISTS `planets` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) DEFAULT NULL,
  `Climate` varchar(255) DEFAULT NULL,
  `Terrain` varchar(255) DEFAULT NULL,
  `ApperancesFilms` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;
```


## Bibliografia - Links
- https://golang.org/doc/faq
- https://www.mongodb.com/languages/golang
- https://swapi.dev/documentation#planets
