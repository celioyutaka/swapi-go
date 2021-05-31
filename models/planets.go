package models

type Planet struct {
	Id              string `json:"id" bson:"_id,omitempty"`
	Name            string `json:"name" bson:"name,omitempty"`
	Climate         string `json:"climate" bson:"climate,omitempty"`
	Terrain         string `json:"terrain" bson:"terrain,omitempty"`
	ApperancesFilms int    `json:"apperancesfilms" bson:"apperancesfilms,omitempty"`
}
