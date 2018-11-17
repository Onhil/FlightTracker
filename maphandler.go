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
	//Gets all the planes from database
	pllanes, err := DBValues.GetPlanes(nil)
	if err != nil {
		// TODO better error
		http.Error(w, "Error not able to get planes", http.StatusBadRequest)
		return
	}
	//Gets all the airports from database
	airrports, err = DBValues.GetAirport(nil)
	if err != nil {
		// TODO better error
		http.Error(w, "Error not able to get airports", http.StatusBadRequest)
		return
	}
	//maps so that the html template will loop through all airports and plans correctly
	planes := make(map[int]Planes)
	airports := make(map[int]Airport)

	//Puts all the planes in the map
	for i := 0; i < len(pllanes); i++ {
		planes[i+1] = pllanes[i]
	}
	//Puts all the airports in the map
	for i := 0; i < len(airrports); i++ {
		airports[i+1] = airrports[i]
	}
	//Creates the struct that will be put in html template
	p := Markers{Title: "Tracking all planes", Planes: planes, Airports: airports}
	//Parses html file
	t, err := template.ParseFiles("index.html")
	if err != nil {
		// TODO better error
		http.Error(w, "Error in parsing index", http.StatusBadRequest)
		return
	}
	//Displays all planes and airports
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
	//Gets icao from url
	icao24 := parts[len(parts)-1]

	var pllanes []Planes
	var airrports []Airport
	var airport []Airport
	//Gets specific plane from database
	pllanes, err := DBValues.GetPlanes(bson.M{"icao24": icao24})
	if err != nil {
		http.Error(w, "Error getting planes", http.StatusBadRequest)
		return
	}
	//Check if plane was returned
	if len(pllanes) == 0 {
		http.Error(w, "Error no plane with that icao", http.StatusBadRequest)
		return
	}

	//Gets arrival aiport
	airport, err = DBValues.GetAirport(bson.M{"icao": pllanes[0].EstArrivalAirport})
	if err != nil {

	}
	airrports = append(airrports, airport...)
	//Gets departure airport
	airport, err = DBValues.GetAirport(bson.M{"icao": pllanes[0].EstDepartureAirport})
	if err != nil {

	}

	airrports = append(airrports, airport...)

	planes := make(map[int]Planes)
	airports := make(map[int]Airport)
	//Put plane in map
	planes[0] = pllanes[0]

	//Puts all the airports in map
	//Reason for 1 = 0 instead of 0 = 0 is javascript starts counting at 1 not 0

	if len(airrports) > 0 {
		airports[1] = airrports[0]
		if len(airrports) == 2 {
			airports[2] = airrports[1]
		}

	}

	p := Markers{Title: "Tracking " + planes[0].Icao24, Planes: planes, Airports: airports}

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
	//Gets country from url
	country := parts[len(parts)-1]
	country = strings.Replace(country, "_", " ", -1)
	//Gets country from DB
	pllanes, err := DBValues.GetPlanes(bson.M{"origincountry": country})
	if err != nil {
		http.Error(w, "Country not in database", http.StatusBadRequest)
		return
	}
	airrports, err = DBValues.GetAirport(nil)
	if err != nil {
		http.Error(w, "Airports missing", http.StatusFailedDependency)
		return
	}

	planes := make(map[int]Planes)
	airports := make(map[int]Airport)
	//Put planes in map
	for i := 0; i < len(pllanes); i++ {
		planes[i+1] = pllanes[i]
	}
	//Put airport in map
	for i := 0; i < len(airrports); i++ {
		airports[i+1] = airrports[i]
	}

	p := Markers{Title: "Tracking all planes from " + country, Planes: planes, Airports: airports}

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
