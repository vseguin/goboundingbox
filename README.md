# goboundingbox

Simple API providing the following methods:

- a GET request `/id/<12345>` where <12345> is the cartodb_id
it should return a JSON document including the `cartodb_id` requested, the `name`, the `polulation` and the point `coordinates` :
ex :
```
curl -ks http://localhost:8080/id/744

{"cartodb_id": 744,
"name": "oriel",
"population": 2500,
"coordinates": [-80.643498,43.069946]
}
```

## Installation

Ensure to run `go get github.com/gorilla/mux` and `github.com/zemirco/couchdb` that are dependencies for the API.

The file `utils/locationparser.go` is used to parse cities and save them in a CouchDB database, so you need to install CouchDB first. Also, set an environment variable as `COUCHDB_URL` with value such as `http://127.0.0.1:5984/` to ensure it runs.

Finally, main executable is `main.go`.