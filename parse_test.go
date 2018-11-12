package main

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) { // WIP
	/*test := []byte(`{
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
		]]}`)*/
	// {ab1644 UAL1254  United States -84.8207 38.5694 11262.36 false 274.2 36.76 0 11513.82 5226 false}

	// s := State{"ab1644", "UAL1254",  "United States", -84.8207, 38.5694, 11262.36, false, 274.2, 36.76, 0, 11513.82, "5226", false}
	var flights []Flight

	if err := json.Unmarshal(body(timeFlights()), &flights); err != nil {
		t.Error(err)
	}
}

func TestParseAirport(t *testing.T) {
	/*

		"encoding/json"
		"fmt"
		"github.com/globalsign/mgo"

		db := setupDB(t)
		defer tearDownDB(t, db)

		db.Init()
		if db.Count(db.CollectionState) != 0 {
			t.Error("Database not properly initialized, database count should be 0")
		}

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
		session, err := mgo.Dial(db.HostURL)
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
		fmt.Println(state.States[0])


	*/
}
