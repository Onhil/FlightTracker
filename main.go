package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/render"
	"github.com/gorilla/mux"
)

// Functions

// PlaneHandler is the function which handles planes and displays a google map, it is currently in an early stage of development.
func PlaneHandler(w http.ResponseWriter, r *http.Request) {

	var pllanes []Planes
	pllanes, _ = DBValues.GetPlanes(nil)

	planes := make(map[int]Planes)

	for i := 0; i < len(pllanes); i++ {
		planes[i] = pllanes[i]
	}

	p := Markers{Title: "Plz Work", Planes: planes}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		// TODO better error
		http.Error(w, "Error in parsing index", http.StatusBadRequest)
	}
	err = t.Execute(w, p)
	if err != nil {
		// TODO better error
		http.Error(w, "Error in executing", http.StatusBadRequest)
	}
}

// OriginCountryHandler handles origin country
func OriginCountryHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	country := parts[len(parts)-1]
	if data, err := DBValues.GetPlanes(bson.M{"origincountry": country}); err != nil {
		http.Error(w, "Country not in database", http.StatusBadRequest)
	} else {
		render.JSON(w, r, data)
	}

	// print out data about country?
	// get planes from country?
	// get airports in country?
}

// DepartureHandler handles departures
func DepartureHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	depAirport := parts[len(parts)-1]
	if data, err := DBValues.GetPlanes(bson.M{"estDepartureAirport": depAirport}); err != nil {
		http.Error(w, "Departure not in database", http.StatusBadRequest)
	} else {
		render.JSON(w, r, data)
	}

	// get flights with estimated dep. airport
	// more?
}

// ArrivalHandler handles arrivals
func ArrivalHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	arrAirport := parts[len(parts)-1]
	if data, err := DBValues.GetPlanes(bson.M{"estarrivalaiport": arrAirport}); err != nil {
		http.Error(w, "Arrival not in database", http.StatusBadRequest)
	} else {
		render.JSON(w, r, data)
	}

	// get flights with estimated arr. airport
	// more?
}

// main
func main() {

	// Database values
	DBValues = Database{
		HostURL:           "mongodb://dataAccess:gettingData123@ds253203.mlab.com:53203/opensky",
		DatabaseName:      "opensky",
		CollectionState:   "States",
		CollectionAirport: "Airports",
		CollectionFlight:  "Flights",
	}

	// Sets the port as what it is assigned to be or 8080 if none is found
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()
	router.HandleFunc("/flight-tracker", PlaneHandler)
	router.HandleFunc("/flight-tracker/{country:[A-Za-z_ ]+}", OriginCountryHandler) // fix path might crash with other funcs
	// merge the two funcs below?
	router.HandleFunc("/flight-tracker/{departing:[A-Z]{4}}", DepartureHandler)
	router.HandleFunc("/flight-tracker/{arriving:[A-Z]{4}}", ArrivalHandler)
	// Handle functions
	log.Fatal(http.ListenAndServe(":"+port, router))
}
