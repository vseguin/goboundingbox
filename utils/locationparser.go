package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"

	"../entities"

	"github.com/zemirco/couchdb"
)

func main() {
	raw, err := ioutil.ReadFile("data/canada_cities.geojson")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var citiesList entities.CitiesList
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
