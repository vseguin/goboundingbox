package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"./utils"

	"github.com/gorilla/mux"
	"github.com/zemirco/couchdb"
)

type CityModel struct {
	ID          string    `json:"cartodb_id"`
	Name        string    `json:"name"`
	Population  int64     `json:"population"`
	Coordinates []float64 `json:"coordinates"`
}

func main() {
	u, err := url.Parse(os.Getenv("COUCHDB_URL"))
	if err != nil {
		panic(err)
	}

	dbClient, err := couchdb.NewClient(u)
	if err != nil {
		panic(err)
	}

	db := dbClient.Use("cities")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Home)
	router.HandleFunc("/id/{locationId}", GetLocation(db))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Bounding Box API.")
}

func GetLocation(db couchdb.DatabaseService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		locationId := vars["locationId"]

		doc := &utils.City{}

		if err := db.Get(doc, locationId); err != nil {
			w.WriteHeader(400)
			fmt.Fprintln(w, "location does not exist.")
			return
		}

		model := &CityModel{
			ID:          doc.ID,
			Name:        doc.Properties.Name,
			Population:  doc.Properties.Population,
			Coordinates: doc.Geometry.Coordinates,
		}

		json.NewEncoder(w).Encode(model)
	}
}
