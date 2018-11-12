package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/render"

	"github.com/go-chi/chi"
)

// Functions

// PlaneHandler is the function which handles planes and displays a google map, it is currently in an early stage of development.
func PlaneHandler(w http.ResponseWriter, r *http.Request) {

	var pllanes []State
	pllanes, _ = DBValues.GetAllFlights()

	planes := make(map[int]State)

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
	country := chi.URLParam(r, "country")
	if data, err := DBValues.GetOriginCountry(country); err != nil {
		http.Error(w, "Country not in database", http.StatusBadRequest)
	} else {
		render.JSON(w, r, data)
	}
}

// DepartureHandler handles departures
func DepartureHandler(w http.ResponseWriter, r *http.Request) {

}

// ArrivalHandler handles arrivals
func ArrivalHandler(w http.ResponseWriter, r *http.Request) {

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

	router := chi.NewRouter()
	router.Route("/flight-tracker", func(r chi.Router) {
		r.Get("/", PlaneHandler)
		r.Route("/country", func(r chi.Router) {
			r.Get("/{country:[A-Za-z_ ]+}", OriginCountryHandler)
		})
		r.Route("/airport", func(r chi.Router) {
			r.Route("/departing", func(r chi.Router) {
				r.Get("/{departing:[A-Z]+}", DepartureHandler)
			})
			r.Route("/arriving", func(r chi.Router) {
				r.Get("/{arriving:[A-Z]+}", ArrivalHandler)
			})
		})
	})
	// Handle functions
	log.Fatal(http.ListenAndServe(":"+port, router))
}
