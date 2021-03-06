package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/render"
	"github.com/gorilla/mux"
)

// Functions

// OriginCountryHandler handles origin country
func OriginCountryHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	country := parts[len(parts)-1]
	country = strings.Replace(country, "_", " ", -1)

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
	if data, err := DBValues.GetFlight(bson.M{"estdepartureairport": depAirport}); err != nil {
		http.Error(w, "Departure not in database", http.StatusBadRequest)
	} else {
		render.JSON(w, r, data)
	}
}

// ArrivalHandler handles arrivals
func ArrivalHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	arrAirport := parts[len(parts)-1]
	if data, err := DBValues.GetFlight(bson.M{"estarrivalairport": arrAirport}); err != nil {
		http.Error(w, "Arrival not in database", http.StatusBadRequest)
	} else {
		render.JSON(w, r, data)
	}
}

// PlaneListHandler Lists all planes by ICAO24
func PlaneListHandler(w http.ResponseWriter, r *http.Request) {

	planes, err := DBValues.GetState(nil)
	if err != nil {
		http.Error(w, "Error getting planes", http.StatusBadRequest)
		return
	}

	render.JSON(w, r, planes)
}

// PlaneInfoHandler Returns information about plane
func PlaneInfoHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	icao24 := parts[len(parts)-1]

	temp, err := DBValues.GetState(bson.M{"icao24": icao24})
	if err != nil {
		http.Error(w, "Error getting plane info", http.StatusBadRequest)
		return
	}

	render.JSON(w, r, temp)
}

// PlaneFieldHandler Returns information about a certain field for the plane
func PlaneFieldHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	icao24 := parts[len(parts)-2]
	field := parts[len(parts)-1]

	temp, err := DBValues.GetState(bson.M{"icao24": icao24})
	if err != nil {
		http.Error(w, "Error getting plane info", http.StatusBadRequest)
		return
	}

	plane := temp[0]

	response, err := plane.getField(field)
	if err != nil {
		http.Error(w, "Unable to find Field", http.StatusBadRequest)
		return
	}
	render.JSON(w, r, response)
}

// CountryHandler Returns all planes from a country
func CountryHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	country := parts[len(parts)-1]

	country = strings.Replace(country, "_", " ", -1)

	plane, err := DBValues.GetState(bson.M{"origincountry": country})
	if err != nil {
		http.Error(w, "Unable to find any Planes with given OriginCountry", http.StatusBadRequest)
		return
	}

	render.JSON(w, r, plane)
}

// AirportListHandler Lists all airports by ICAO
func AirportListHandler(w http.ResponseWriter, r *http.Request) {
	airports, err := DBValues.GetAirport(nil)
	if err != nil {
		http.Error(w, "Unable to find any Airports with the given ICAO", http.StatusBadRequest)
		return
	}

	render.JSON(w, r, airports)
}

// AirportInfoHandler Returns information about the airport and the ICAO24 of all planes that arrives and depart from it
func AirportInfoHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	icao := parts[len(parts)-1]

	airport, err := DBValues.GetAirport(bson.M{"icao": icao})
	if err != nil {
		http.Error(w, "Unable to find any Airports with given ICAO", http.StatusBadRequest)
		return
	}

	port := airport[0] //Convert array to single airport, in case of more than one airport with the ICAO which should not happen

	render.JSON(w, r, port)
}

// AirportFieldHandler Returns the field information for the airport
func AirportFieldHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	icao := parts[len(parts)-2]
	field := parts[len(parts)-1]

	airport, err := DBValues.GetAirport(bson.M{"icao": icao})
	if err != nil {
		http.Error(w, "Unable to find any Airports with given ICAO", http.StatusBadRequest)
		return
	}

	port := airport[0] //Convert array to single airport, in case of more than one airport with the ICAO which should not happen
	response, err := port.getField(field)
	if err != nil {
		http.Error(w, "Unable to find Field", http.StatusBadRequest)
		return
	}
	render.JSON(w, r, response)
}

// AirportCountryHandler Returns all countries with an airport
func AirportCountryHandler(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Country)
}

// AirportInCountryHandler Names all the airports in the given country
func AirportInCountryHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	country := parts[len(parts)-1]

	country = strings.Replace(country, "_", " ", -1)

	airports, err := DBValues.GetAirport(bson.M{"country": country})
	if err != nil {
		http.Error(w, "Unable to find any Airports in the country", http.StatusBadRequest)
		return
	}
	render.JSON(w, r, airports)
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
	router.HandleFunc("/flight-tracker/country/{country:.+}", OriginCountryHandler)
	router.HandleFunc("/flight-tracker/plane", PlaneListHandler)
	router.HandleFunc("/flight-tracker/plane/{icao24:[A-Za-z0-9]+}", PlaneInfoHandler)
	router.HandleFunc("/flight-tracker/plane/{icao24:[A-Za-z0-9]+}/{field:[A-Za-z0-9]+}", PlaneFieldHandler)
	router.HandleFunc("/flight-tracker/map/plane/{icao24:[A-Za-z0-9]+}", PlaneMapHandler)
	router.HandleFunc("/flight-tracker/plane/country/{country:.+}", CountryHandler)
	router.HandleFunc("/flight-tracker/map/country/{country:.+}", CountryMapHandler)
	router.HandleFunc("/flight-tracker/airport", AirportListHandler)
	router.HandleFunc("/flight-tracker/airport/{icao:[A-Z]{4}}", AirportInfoHandler)
	router.HandleFunc("/flight-tracker/airport/{icao:[A-Z]{4}}/{field:[A-Za-z0-9]+}", AirportFieldHandler)
	router.HandleFunc("/flight-tracker/airport/country", AirportCountryHandler)
	router.HandleFunc("/flight-tracker/airport/country/{country:.+}", AirportInCountryHandler)
	router.HandleFunc("/flight-tracker/departure/{departing:[A-Z]{4}}", DepartureHandler)
	router.HandleFunc("/flight-tracker/arrival/{arriving:[A-Z]{4}}", ArrivalHandler)
	// Handle functions
	log.Fatal(http.ListenAndServe(":"+port, router))
}
