package swapi

type PlanetsSWApi struct {
	Count    string        `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []PlanetSWApi `json:"results"`
}

type PlanetSWApi struct {
	Name            string        `json:"name"`
	Rotation_period string        `json:"rotation_period"`
	Orbital_period  string        `json:"orbital_period"`
	Diameter        string        `json:"diameter"`
	Climate         string        `json:"climate"`
	Gravity         string        `json:"gravity"`
	Terrain         string        `json:"terrain"`
	Surface_water   string        `json:"surface_water"`
	Population      string        `json:"population"`
	Residents       []interface{} `json:"residents"`
	Films           []interface{} `json:"films"`
	Created         string        `json:"created"`
	Edited          string        `json:"edited"`
	Url             string        `json:"url"`
}
