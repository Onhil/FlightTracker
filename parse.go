package main

import (
	"encoding/json"
	"fmt"
)

/*
###  EXAMPLE USE:

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
	session.DB(DBValues.DatabaseName).C(DBValues.CollectionName).RemoveAll(nil)
	var sarray []interface{}
	for i := range state.States {
		sarray = append(sarray, state.States[i])
	}
	var s []State
	session.DB(DBValues.DatabaseName).C(DBValues.CollectionName).Insert(sarray...)
	session.DB(DBValues.DatabaseName).C(DBValues.CollectionName).Find(nil).All(&s)
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
	if v[1] == nil {
		s.Callsign = ""
	} else {
		s.Callsign = v[1].(string) // Null
	}
	s.OriginCountry = v[2].(string)
	///	s.TimePosition = v[3].(int)
	///	s.LastContact = v[4].(int)
	if v[5] == nil {
		s.Longitude = 0
	} else {
		s.Longitude = v[5].(float64) // Null
	}

	if v[6] == nil {
		s.Latitude = 0
	} else {
		s.Latitude = v[6].(float64) // Null
	}

	if v[7] == nil {
		s.BaroAltitude = 0
	} else {
		s.BaroAltitude = v[7].(float64) // Null
	}
	s.OnGround = v[8].(bool)
	if v[9] == nil {
		s.Velocity = 0
	} else {
		s.Velocity = v[9].(float64) // Null
	}

	if v[10] == nil {
		s.TrueTrack = 0
	} else {
		s.TrueTrack = v[10].(float64) // Null
	}

	if v[11] == nil {
		s.VerticalRate = 0
	} else {
		s.VerticalRate = v[11].(float64) // Null
	}

	///	s.Sensors = v[12].([]int)
	if v[13] == nil {
		s.GeoAltitude = 0
	} else {
		s.GeoAltitude = v[13].(float64) // Null
	}

	if v[14] == nil {
		s.Squawk = ""
	} else {
		s.Squawk = v[14].(string) // Null
	}

	s.Spi = v[15].(bool)
	/// s.PositionSource = v[16].(int)

	return nil
}
