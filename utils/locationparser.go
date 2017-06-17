package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"

	"github.com/zemirco/couchdb"
)

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

func main() {
	raw, err := ioutil.ReadFile("data/canada_cities.geojson")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var citiesList CitiesList
	json.Unmarshal(raw, &citiesList)

	u, err := url.Parse(os.Getenv("COUCHDB_URL"))
	if err != nil {
		panic(err)
	}

	client, err := couchdb.NewClient(u)
	if err != nil {
		panic(err)
	}

	if _, err = client.Create("cities"); err != nil {
		panic(err)
	}

	db := client.Use("cities")

	for _, city := range citiesList.Cities {
		city.ID = strconv.Itoa(city.Properties.ID)

		result, err := db.Post(city)

		if err != nil {
			panic(err)
		}

		fmt.Println(result)
	}
}
