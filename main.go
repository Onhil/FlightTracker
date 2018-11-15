package main

import (
	"encoding/json"
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
	if data, err := DBValues.GetPlanes(bson.M{"estdepartureairport": depAirport}); err != nil {
		http.Error(w, "Departure not in database", http.StatusBadRequest)
	} else {
		render.JSON(w, r, data)
	}
}

// ArrivalHandler handles arrivals
func ArrivalHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	arrAirport := parts[len(parts)-1]
	if data, err := DBValues.GetPlanes(bson.M{"estarrivalairport": arrAirport}); err != nil {
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
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")

	icao24 := parts[len(parts)-1]

	temp, err := DBValues.GetState(bson.M{"icao24": icao24})
	if err != nil {
		http.Error(w, "Error getting plane info", http.StatusBadRequest)
		return
	}

	PlaneJSON, err := json.Marshal(temp)
	if err != nil {
		http.Error(w, "Error parsing plane info", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(PlaneJSON)
}

// PlaneFieldHandler Returns information about a certain field for the plane
func PlaneFieldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")

	icao24 := parts[len(parts)-2]
	field := parts[len(parts)-1]

	temp, err := DBValues.GetState(bson.M{"icao24": icao24})
	if err != nil {
		http.Error(w, "Error getting plane info", http.StatusBadRequest)
		return
	}

	plane := temp[0]

	var Response interface{}

	switch field {
	case "Callsign":
		Response = plane.Callsign
	case "OriginCountry":
		Response = plane.OriginCountry
	case "Longitude":
		Response = plane.Longitude
	case "Latitude":
		Response = plane.Latitude
	case "BaroAltitude":
		Response = plane.BaroAltitude
	case "OnGround":
		Response = plane.OnGround
	case "Velocity":
		Response = plane.Velocity
	case "TrueTrack":
		Response = plane.TrueTrack
	case "VerticalRate":
		Response = plane.VerticalRate
	case "GeoAltitude":
		Response = plane.GeoAltitude
	case "Squawk":
		Response = plane.Squawk
	case "Spi":
		Response = plane.Spi
	default:
		http.Error(w, "Error no such field", http.StatusBadRequest)
		return
	}
	FieldJSON, err := json.Marshal(Response)
	if err != nil {
		http.Error(w, "Error parsing field", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(FieldJSON)
}

// CountryHandler Returns all planes from a country
func CountryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")

	country := parts[len(parts)-1]

	plane, err := DBValues.GetState(bson.M{"origincountry": country})
	if err != nil {
		http.Error(w, "Unable to find any Planes with given OriginCountry", http.StatusBadRequest)
		return
	}
	PlaneNames := []string{}

	for i := 0; i < len(plane); i++ {
		PlaneNames = append(PlaneNames, plane[i].Icao24)
	}

	portJSON, err := json.Marshal(PlaneNames)
	if err != nil {
		http.Error(w, "Unable to parse the Plane names", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(portJSON)
}

// AirportListHandler Lists all airports by ICAO
func AirportListHandler(w http.ResponseWriter, r *http.Request) {
	airports, err := DBValues.GetAirport(nil)
	if err != nil {
		http.Error(w, "Unable to find any Airports with the given ICAO", http.StatusBadRequest)
		return
	}
	AirportIcao := []string{}

	for i := 0; i < len(airports); i++ {
		AirportIcao = append(AirportIcao, airports[i].ICAO)
	}

	portJSON, err := json.Marshal(AirportIcao)
	if err != nil {
		http.Error(w, "Unable to parse the Airport names", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(portJSON)
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

	portJSON, err := json.Marshal(port)
	if err != nil {
		http.Error(w, "Unable to parse the Airport", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(portJSON)
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

	var Response interface{}

	switch field {
	case "ID":
		Response = port.ID
	case "Name":
		Response = port.Name
	case "City":
		Response = port.City
	case "Country":
		Response = port.Country
	case "IATA":
		Response = port.IATA
	case "Latitude":
		Response = port.Latitude
	case "Longitude":
		Response = port.Longitude
	case "Altitude":
		Response = port.Altitude
	case "Timezone":
		Response = port.Timezone
	case "DST":
		Response = port.DST
	case "Tz_Database_Timezone":
		Response = port.TzDatabaseTimezone
	case "Type":
		Response = port.Type
	case "Source":
		Response = port.Source
	default:
		http.Error(w, "Field is not included in Airport!", http.StatusBadRequest)
		return
	}

	portJSON, err := json.Marshal(Response)
	if err != nil {
		http.Error(w, "Unable to parse the Airport", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(portJSON)
}

// AirportCountryHandler Returns all countries with an airport
func AirportCountryHandler(w http.ResponseWriter, r *http.Request) {
	CountryJSON, err := json.Marshal(Country)
	if err != nil {
		http.Error(w, "Unable to parse the countries", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(CountryJSON)
}

// AirportInCountryHandler Names all the airports in the given country
func AirportInCountryHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	country := parts[len(parts)-1]

	airports, err := DBValues.GetAirport(bson.M{"country": country})
	if err != nil {
		http.Error(w, "Unable to find any Airports in the country", http.StatusBadRequest)
		return
	}
	AirportNames := []string{}

	for i := 0; i < len(airports); i++ {
		AirportNames = append(AirportNames, airports[i].ICAO)
	}

	portJSON, err := json.Marshal(AirportNames)
	if err != nil {
		http.Error(w, "Unable to parse the Airport names", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(portJSON)
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
	router.HandleFunc("/flight-tracker/{departing:[A-Z]{4}}", DepartureHandler)
	router.HandleFunc("/flight-tracker/{arriving:[A-Z]{4}}", ArrivalHandler)
	// Handle functions
	log.Fatal(http.ListenAndServe(":"+port, router))
}
