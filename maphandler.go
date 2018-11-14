package main

import (
	"html/template"
	"net/http"
)

// PlaneHandler is the function which handles planes and displays a google map, it is currently in an early stage of development.
func PlaneHandler(w http.ResponseWriter, r *http.Request) {

	var pllanes []Planes
	var airrports []Airport

	pllanes, _ = DBValues.GetPlanes(nil)
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
	}
	err = t.Execute(w, p)
	if err != nil {
		// TODO better error
		http.Error(w, "Error in executing", http.StatusBadRequest)
	}
}

// PlaneMapHandler Shows the plane on the map
func PlaneMapHandler(w http.ResponseWriter, r *http.Request) {
	//Show plane
	//Show arrival and departure airport
}

// CountryMapHandler Shows all planes from country on the map
func CountryMapHandler(w http.ResponseWriter, r *http.Request) {
	//Show all planes from country
	//Show all airports
}
