package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// ## EXAMPLE USES
/*
	test := []byte(`{
		"time": 1541448600,
		"states":
			[[
				"ab1644",
				"UAL1254 ",
				"United States",
				1541448598,
				1541448599,
				-84.8207,
				38.5694,
				11262.36,
				false,
				274.2,
				36.76,
				0,
				null,
				11513.82,
				"5226",
				false,
				0
			]]}`)

	session, err := mgo.Dial(DBValues.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()
	var state States
	if err := json.Unmarshal(test, &state); err != nil {
		fmt.Println("error")
	}

	session.DB(DBValues.DatabaseName).C(DBValues.CollectionState).RemoveAll(nil)

	session.DB(DBValues.DatabaseName).C(DBValues.CollectionState).Insert(state.States[0])
	fmt.println(state.States[0])
*/

/*
	session, err := mgo.Dial(DBValues.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()
	resp, _ := http.Get("https://opensky-network.org/api/states/all")

	body, _ := ioutil.ReadAll(resp.Body)
	var state States
	if err := json.Unmarshal(body, &state); err != nil {
		fmt.Println("error")
	}
	session.DB(DBValues.DatabaseName).C(DBValues.CollectionState).RemoveAll(nil)
	var sarray []interface{}
	for i := range state.States {
		sarray = append(sarray, state.States[i])
	}
	var s []State
	session.DB(DBValues.DatabaseName).C(DBValues.CollectionState).Insert(sarray...)
	session.DB(DBValues.DatabaseName).C(DBValues.CollectionState).Find(nil).All(&s)
	fmt.Println(s)
*/

// UnmarshalJSON states from GET /states/all from OpenSky
func (s *State) UnmarshalJSON(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		fmt.Printf("Error whilde decoding %v\n", err)
		return err
	}
	s.Icao24 = v[0].(string)
	if v[1] != nil {
		s.Callsign = v[1].(string) // Null
	}
	s.OriginCountry = v[2].(string)
	///	s.TimePosition = v[3].(int)
	///	s.LastContact = v[4].(int)
	if v[5] != nil {
		s.Longitude = v[5].(float64) // Null
	}

	if v[6] != nil {
		s.Latitude = v[6].(float64) // Null
	}

	if v[7] != nil {
		s.BaroAltitude = v[7].(float64) // Null
	}
	s.OnGround = v[8].(bool)
	if v[9] != nil {
		s.Velocity = v[9].(float64) // Null
	}

	if v[10] != nil {
		s.TrueTrack = v[10].(float64) // Null
	}

	if v[11] != nil {
		s.VerticalRate = v[11].(float64) // Null
	}

	///	s.Sensors = v[12].([]int)
	if v[13] != nil {
		s.GeoAltitude = v[13].(float64) // Null
	}

	if v[14] != nil {
		s.Squawk = v[14].(string) // Null
	}

	s.Spi = v[15].(bool)
	/// s.PositionSource = v[16].(int)

	return nil
}

func timeFlights() string {
	end := time.Now().Unix()
	begin := end - 7200
	url := fmt.Sprintf("https://opensky-network.org/api/flights/all?begin=%d&end=%d", begin, end)
	return url
}

func mergeStatesAndFlights(s []State, f []Flight) []Planes {
	var planes []Planes
	var p Planes
	// Adds j from for loop f that has the same callsing as s
	var fCall []int

	for i := range s {

		if s[i].Callsign != "" {
			for j := range f {
				// If s and f callsign is the same it merges them into planes
				if s[i].Callsign == f[j].Callsign {
					p = Planes{s[i], f[j]}
					planes = append(planes, p)
					fCall = append(fCall, j)
				}
			}
			// If s callsign is empty add it seperatly to planes
		} else {
			p = Planes{s[i], Flight{}}
			planes = append(planes, p)
		}
	}

	for i := range f {
		b := false
		// Checks wether or not index has been used before
		for j := range fCall {
			if i == fCall[j] {
				b = true
			}
		}
		// If index have not been used it adds flight to planes
		if !b {
			p = Planes{State{}, f[i]}
			planes = append(planes, p)
		}
	}
	return planes
}

// ## EXAMPLE USES OF AIPORT MARSHAL
/*
test := []byte(`{
		[[
			"23",
			"Gjovik Airport",
			"Gjovik",
			"Norway",
			"GJO",		//Can be null
			"GJOV"		//Can be null
			-84.8207,
			38.5694,
			350,
			1,
			"E",		//E, A, S, O, Z, N or U
			"Europe/Oslo",
			"airport",
			"Example"
		]]}`)
*/

//ParseAirport parses the data for airports
func (a *Airport) ParseAirport(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		fmt.Printf("Error whilde decoding %v\n", err)
		return err
	}
	a.ID = v[0].(int)
	a.Name = v[1].(string)
	a.City = v[2].(string)
	a.Country = v[3].(string)

	if v[4] != nil {
		a.IATA = v[4].(string) //Null
	}

	if v[5] != nil {
		a.ICAO = v[5].(string) //Null
	}
	fmt.Println(a.Name)
	a.Latitude = v[6].(float64)
	a.Longitude = v[7].(float64)
	a.Altitude = v[8].(float64)
	a.Timezone = v[9].(string)
	a.DST = v[10].(string)
	a.TzDatabaseTimezone = v[11].(string)
	a.Type = v[12].(string)
	a.Source = v[13].(string)

	return nil
}
