package main

import (
	"encoding/json"
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
}

// ArrivalHandler handles arrivals
func ArrivalHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	arrAirport := parts[len(parts)-1]
	if data, err := DBValues.GetPlanes(bson.M{"estArrivalAirport": arrAirport}); err != nil {
		http.Error(w, "Arrival not in database", http.StatusBadRequest)
	} else {
		render.JSON(w, r, data)
	}
}

// PlaneListHandler Lists all planes by ICAO24
func PlaneListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	planes := []Planes{}
	planes, err := DBValues.GetPlanes(nil)

	for i := 0; i < len(planes); i++ {
		planes = append(planes, planes[i].Icao24)
	}

	IcaoJSON, _ := json.Marshal(planes)
	w.WriteHeader(http.StatusOK)
	w.Write(IcaoJSON)
}

/*
func PlaneInfoHandler(w http.ResponseWriter, r *http.Request) {			// Returns information about plane

}

func PlaneFieldHandler(w http.ResponseWriter, r *http.Request) { 		// Returns information about a certain field for the plane

}

func PlaneMapHandler(w http.ResponseWriter, r *http.Request) {	 		// Shows the plane on the map

}

func CountryHandler(w http.ResponseWriter, r *http.Request) {			// Returns all planes from a country

}

func CountryMapHandler(w http.ResponseWriter, r *http.Request) {		// Shows all planes from country on the map

}

func AirportListHandler(w http.ResponseWriter, r *http.Request) {		// Lists all airports by ICAO

}
																		// Returns information about the airport and
func AirportInfoHandler(w http.ResponseWriter, r *http.Request) {		// the ICAO24 of all planes that arrives and depart from it

}

func AirportFieldHandler(w http.ResponseWriter, r *http.Request) {		// Returns the field information for the airport

}

func AirportCountryHandler(w http.ResponseWriter, r *http.Request) {	// Returns all countries with an airport

}

func AirportInCountryHandler(w http.ResponseWriter, r *http.Request) {	// Names all the airports in the given country

}
*/
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
	router.HandleFunc("/flight-tracker/{country:.+}", OriginCountryHandler)
	router.HandleFunc("/flight-tracker/plane", PlaneListHandler)
	/*router.HandleFunc("/flight-tracker/plane/{icao24:[A-Za-z0-9]+}", PlaneInfoHandler)
	router.HandleFunc("/flight-tracker/plane/{icao24:[A-Za-z0-9]}/{field:[A-Za-z0-9]+}", PlaneFieldHandler)
	router.HandleFunc("/flight-tracker/plane/map/{icao24:[A-Za-z0-9]+}", PlaneMapHandler)
	router.HandleFunc("/flight-tracker/plane/country/{country:.+}", CountryHandler)
	router.HandleFunc("/flight-tracker/plane/country/map/{country:.+}", CountryMapHandler)
	router.HandleFunc("/flight-tracker/airport", AirportListHandler)
	router.HandleFunc("/flight-tracker/airport/{icao:[A-Z]{4}}", AirportInfoHandler)
	router.HandleFunc("/flight-tracker/airport/{icao:[A-Z]{4}}/{field:[A-Za-z0-9]+}", AirportFieldHandler)
	router.HandleFunc("/flight-tracker/airport/country", AirportCountryHandler)
	router.HandleFunc("/flight-tracker/airport/country/{country:.+}", AirportInCountryHandler)*/
	router.HandleFunc("/flight-tracker/{departing:[A-Z]{4}}", DepartureHandler)
	router.HandleFunc("/flight-tracker/{arriving:[A-Z]{4}}", ArrivalHandler)
	// Handle functions
	log.Fatal(http.ListenAndServe(":"+port, router))
}
