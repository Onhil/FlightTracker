package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/go-chi/render"

	"github.com/go-chi/chi"
)

// Structs

// States holds time and a list of states(not sure if this is important or not)
type States struct {
	Time   int     `json:"time"`
	States []State `json:"states"`
}

// State os a struct witch
type State struct {
	Icao24        string `json:"Icao24"`        // Unique ICAO 24-bit address of the transponder in hex string representation.
	Callsign      string `json:"Callsign"`      // Callsign of the vehicle (8 chars). Can be null if no callsign has been received.
	OriginCountry string `json:"OriginCountry"` // Country name inferred from the ICAO 24-bit address.
	//TimePosition  int  `json:"TimePosition"`  // Unix timestamp (seconds) for the last position update. Can be null if no position report was received by OpenSky within the past 15s.
	//LastContact   int  `json:"LastContact"`  // Unix timestamp (seconds) for the last update in general. This field is updated for any new, valid message received from the transponder.
	Longitude    float64 `json:"Longitude"`    // WGS-84 longitude in decimal degrees. Can be null.
	Latitude     float64 `json:"Latitude"`     // WGS-84 latitude in decimal degrees. Can be null.
	BaroAltitude float64 `json:"BaroAltitude"` // Barometric altitude in meters. Can be null.
	OnGround     bool    `json:"OnGround"`     // Boolean value which indicates if the position was retrieved from a surface position report.
	Velocity     float64 `json:"Velocity"`     // Velocity over ground in m/s. Can be null.
	TrueTrack    float64 `json:"TrueTrack"`    // True track in decimal degrees clockwise from north (north=0°). Can be null.
	VerticalRate float64 `json:"VerticalRate"` // Vertical rate in m/s. A positive value indicates that the airplane is climbing, a negative value indicates that it descends. Can be null.
	// Sensors   []int   `json:"Sensors"` 	// IDs of the receivers which contributed to this state vector. Is null if no filtering for sensor was used in the request.
	GeoAltitude float64 `json:"GeoAltitude"` // Geometric altitude in meters. Can be null.
	Squawk      string  `json:"Squawk"`      // The transponder code aka Squawk. Can be null.
	Spi         bool    `json:"spi"`         // Whether flight status indicates special purpose indicator.
	/// PositionSource int     `json:"positionSource"` //Origin of this state’s position: 0 = ADS-B, 1 = ASTERIX, 2 = MLAT
}

// Database holds database basic data
type Database struct {
	HostURL        string
	DatabaseName   string
	CollectionName string
}

type flights struct {
	Icao24              string `json:"icao24"`              // Unique ICAO 24-bit address of the transponder in hex string representation. All letters are lower case.
	FirstSeen           int    `json:"firstSeen"`           // Estimated time of departure for the flight as Unix time (seconds since epoch).
	EstDepartureAirport string `json:"estDepartureAirport"` // ICAO code of the estimated departure airport. Can be null if the airport could not be identified.
	LastSeen            int    `json:"lastSeen"`            // Estimated time of arrival for the flight as Unix time (seconds since epoch)
	EstArrivalAirport   string `json:"estArrivalAiport"`    // ICAO code of the estimated arrival airport. Can be null if the airport could not be identified.
	Callsign            string `json:"callsign"`            // Callsign of the vehicle (8 chars). Can be null if no callsign has been received. If the vehicle transmits multiple callsigns during the flight, we take the one seen most frequently
	// EstDepartureAirportHorizDistance int    // Horizontal distance of the last received airborne position to the estimated departure airport in meters
	// EstDepartureAirportVertDistance  int    // Vertical distance of the last received airborne position to the estimated departure airport in meters
	// EstArrivalAirportHorizDistance   int    // Horizontal distance of the last received airborne position to the estimated arrival airport in meters
	// EstArrivalAirportVertDistance    int    // Vertical distance of the last received airborne position to the estimated arrival airport in meters
	// DepartureAirportCandidatesCount  int    // Number of other possible departure airports. These are airports in short distance to estDepartureAirport.
	// ArrivalAirportCandidatesCount    int    // Number of other possible departure airports. These are airports in short distance to estArrivalAirport.
}

// PlaneMarker is a marker for a plane
type PlaneMarker struct {
	Lat              float64
	Long             float64
	Icao24           string
	Callsign         string
	DepartureAirport string
	DepartureTime    int
	TrueTrack        float64
}

// Markers holds markers values
type Markers struct {
	Title  string
	Planes map[int]PlaneMarker
}

// DBValues is a database element which is accessible everywhere, not sure if this is needed to be honest
var DBValues Database

// Functions

// PlaneHandler is the function which handles planes and displays a google map, it is currently in an early stage of development.
func PlaneHandler(w http.ResponseWriter, r *http.Request) {

	pllanes := make(map[int]PlaneMarker)

	pllanes[0] = PlaneMarker{Lat: 55.508742, Long: -0.120850, Icao24: "IK2314", Callsign: "DEC342", DepartureAirport: "ENBR", DepartureTime: 12, TrueTrack: 0}
	pllanes[1] = PlaneMarker{Lat: 58.508742, Long: -2.120850, Icao24: "rsg34", Callsign: "234fsd", DepartureAirport: "ENBFL", DepartureTime: 12, TrueTrack: 67}
	pllanes[2] = PlaneMarker{Lat: 67.508742, Long: -8.120850, Icao24: "Ywer3", Callsign: "324sdf", DepartureAirport: "ENGR", DepartureTime: 12, TrueTrack: 90}

	p := Markers{Title: "Plz Work", Planes: pllanes}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

// OriginCountryHandler handles origin country
func OriginCountryHandler(w http.ResponseWriter, r *http.Request) {
	country := chi.URLParam(r, "country")
	if data, ok := DBValues.GetOriginCountry(country); !ok {
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

func main() {

	// Database values
	DBValues = Database{
		HostURL:        "mongodb://dataAccess:gettingData123@ds253203.mlab.com:53203/opensky",
		CollectionName: "States",
		DatabaseName:   "opensky",
	}

	// Sets the port as what it is assigned to be or 8080 if none is found
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := chi.NewRouter()
	router.Route("/flight-tracker", func(r chi.Router) {
		//r.Get("", )
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
	http.HandleFunc("/", PlaneHandler)
	http.ListenAndServe(":"+port, router)
}
