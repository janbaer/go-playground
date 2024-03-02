package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Person struct {
	userId       int    `json:"userId"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Phone        string `json:"phoneNumber"`
	EmailAddress string `json:"emailAddress"`
}

var people []Person

func init() {
	jsonBytes, err := ioutil.ReadFile("./data/people.json")
	if err != nil {
		log.Fatalf("Unexpected error reading JSON file: %v", err)
	}

	// jsonBytes := []byte(`[{"firstName":"John", "lastName":"Doe"}]`)

	if err := json.Unmarshal(jsonBytes, &people); err != nil {
		log.Fatalf("Unexpected error while unmarshalling JSON: %v", err)
	}

}

func main() {
	port := 8000

	router := mux.NewRouter()

	router.HandleFunc("/", handleRootGet).Methods(http.MethodGet)

	fmt.Printf("Running rest-api on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func handleRootGet(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request %v\n", r)

	if q := r.URL.Query().Get("q"); len(q) > 0 {
		log.Printf("User is searching for %s", q)
		if person, found := searchPerson(q); found {
			w.WriteHeader(http.StatusOK)
			w.Header().Add("content-type", "application/json")
			_ = json.NewEncoder(w).Encode(person)
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("No person starting with %s found", q)))
		}

	}

}

func searchPerson(name string) (Person, bool) {
	for _, p := range people {
		log.Printf("Compare person %s", p.LastName)
		if strings.HasPrefix(p.LastName, name) {
			return p, true
		}
	}

	return Person{}, false
}
