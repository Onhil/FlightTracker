package main

// Structs

// States holds time and a list of states(not sure if this is important or not)
type States struct {
	Time   int     `json:"time"`
	States []State `json:"states"`
}

// State is a struct witch stores states of data
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

// Flight stores flight data and info about departure and arrival airport
type Flight struct {
	Icao24F             string `json:"icao24"`              // Unique ICAO 24-bit address of the transponder in hex string representation. All letters are lower case.
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

// Planes is a struct with data from states and flights
type Planes struct {
	State
	Flight
}

//Airport is a struct that has data regarding the airports in the world
type Airport struct {
	ID                 int     `json:"ID"`                    //Unique OpenFlights identifier for this airport
	Name               string  `json:"Name"`                  //Name of airport. May or may nor contain the City name.
	City               string  `json:"City"`                  //Main city served by airport. May be spelled differently from Name.
	Country            string  `json:"Country"`               //Country or territory where airport is located.
	IATA               string  `json:"IATA"`                  //3-letter IATA code. Null if not assigned/unknown
	ICAO               string  `json:"ICAO"`                  //4-letter ICAO code. Null if not assigned.
	Latitude           float64 `json:"Latitude"`              //Decimal degrees, usally to six significant digits. Negative is West, positive is East.
	Longitude          float64 `json:"Longitude"`             //Decimal degrees, usally to six significant digits. Negative is West, positive is East.
	Altitude           float64 `json:"Altitude"`              //In feet.
	Timezone           float64 `json:"Timezone"`              //Hours offset from UTC. Fractional hours are expressed as decimals, eg. India is 5.5.
	DST                string  `json:"DST"`                   //Daylight savings time. One of E (Europe), A (US/Canada), S (Sout America), O (Australia), Z (New Zealand), N (None) or U (Unknown)
	TzDatabaseTimeZone string  `json:"Tz_database_time_zone"` //Timezone in "tz" (Olson) format, eg. "America/Los_Angeles".
	Type               string  `json:"Type"`                  //Type of airport. only type=airport included.
	Source             string  `json:"Source"`                //Source of this data. "OurAirports" is the only source
}

// AirportJSON holds the json for Airport
type AirportJSON struct {
	Aport     Airport `json:"Airport"`
	Arriving  string  `json:"Arriving"`
	Departing string  `json:"Departing"`
}

// Database holds database basic data
type Database struct {
	HostURL           string
	DatabaseName      string
	CollectionState   string
	CollectionAirport string
	CollectionFlight  string
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
	Planes map[int]Planes
}

// DBValues is a database element which is accessible everywhere, not sure if this is needed to be honest
var DBValues Database
