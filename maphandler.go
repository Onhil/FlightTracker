package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/globalsign/mgo/bson"
)

// PlaneHandler is the function which handles planes and displays a google map, it is currently in an early stage of development.
func PlaneHandler(w http.ResponseWriter, r *http.Request) {

	var pllanes []Planes
	var airrports []Airport

	pllanes, err := DBValues.GetPlanes(nil)
	if err != nil {
		// TODO better error
		http.Error(w, "Error not able to get planes", http.StatusBadRequest)
		return
	}
	airrports, err = DBValues.GetAirport(nil)
	if err != nil {
		// TODO better error
		http.Error(w, "Error not able to get airports", http.StatusBadRequest)
		return
	}

	planes := make(map[int]Planes)
	airports := make(map[int]Airport)

	for i := 0; i < len(pllanes); i++ {
		planes[i] = pllanes[i]
	}

	for i := 0; i < len(airrports); i++ {
		airports[i] = airrports[i]
	}

	p := Markers{Title: "Plz Work", Planes: planes, Airports: airports}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		// TODO better error
		http.Error(w, "Error in parsing index", http.StatusBadRequest)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		// TODO better error
		http.Error(w, "Error in executing", http.StatusBadRequest)
		return
	}
}

// PlaneMapHandler Shows the plane on the map
func PlaneMapHandler(w http.ResponseWriter, r *http.Request) {
	//Show plane
	//Show arrival and departure airport
	parts := strings.Split(r.URL.Path, "/")

	icao24 := parts[len(parts)-1]

	var pllanes []Planes
	var airrports []Airport
	var airport []Airport

	pllanes, err := DBValues.GetPlanes(bson.M{"icao24": icao24})
	if err != nil {
		http.Error(w, "Error getting planes", http.StatusBadRequest)
		return
	}

	if len(pllanes) == 0 {
		http.Error(w, "Error no plane with that icao", http.StatusBadRequest)
		return
	}

	airport, err = DBValues.GetAirport(bson.M{"icao": pllanes[0].EstArrivalAirport})
	if err != nil {
		http.Error(w, "Error getting airport", http.StatusBadRequest)
		return
	}
	airrports = append(airrports, airport...)

	airport, err = DBValues.GetAirport(bson.M{"icao": pllanes[0].EstDepartureAirport})
	if err != nil {
		http.Error(w, "Error getting airport", http.StatusBadRequest)
		return
	}

	airrports = append(airrports, airport...)

	planes := make(map[int]Planes)
	airports := make(map[int]Airport)

	planes[0] = pllanes[0]

	for i := 0; i < len(airrports)-1; i++ {
		airports[i] = airrports[i]
	}

	p := Markers{Title: "Plz Work", Planes: planes, Airports: airports}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		// TODO better error
		http.Error(w, "Error in parsing index", http.StatusBadRequest)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		// TODO better error
		http.Error(w, "Error in executing", http.StatusBadRequest)
		return
	}
}

// CountryMapHandler Shows all planes from country on the map
func CountryMapHandler(w http.ResponseWriter, r *http.Request) {
	//Show all planes from country
	//Show all airports
	parts := strings.Split(r.URL.Path, "/")

	var pllanes []Planes
	var airrports []Airport

	country := parts[len(parts)-1]

	pllanes, _ = DBValues.GetPlanes(bson.M{"origincountry": country})
	airrports, _ = DBValues.GetAirport(nil)

	planes := make(map[int]Planes)
	airports := make(map[int]Airport)

	for i := 0; i < len(pllanes); i++ {
		planes[i] = pllanes[i]
	}

	for i := 0; i < len(airrports); i++ {
		airports[i] = airrports[i]
	}

	p := Markers{Title: "Plz Work", Planes: planes, Airports: airports}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		// TODO better error
		http.Error(w, "Error in parsing index", http.StatusBadRequest)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		// TODO better error
		http.Error(w, "Error in executing", http.StatusBadRequest)
		return
	}

}
