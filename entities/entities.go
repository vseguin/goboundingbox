package entities

type CitiesList struct {
	Cities []City `json:"features"`
}

type City struct {
	Geometry   Point          `json:"geometry"`
	ID         string         `json:"_id"`
	Properties CityProperties `json:"properties"`
}

func (city City) GetID() string {
	return city.ID
}

func (city City) GetRev() string {
	return "0"
}

type Point struct {
	Coordinates []float64 `json:"coordinates"`
}

type CityProperties struct {
	Name       string `json:"name"`
	PlaceKey   string `json:"place_key"`
	Capital    string `json:"capital"`
	Population int64  `json:"population"`
	PClass     string `json:"pclass"`
	ID         int    `json:"cartodb_id"`
}
