package main

import (
	"testing"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func setupDB(t *testing.T) *Database {
	db := Database{
		HostURL:           "mongodb://localhost",
		DatabaseName:      "testing",
		CollectionState:   "testState",
		CollectionAirport: "testAirport",
		CollectionFlight:  "testFlight",
	}

	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		t.Error(err)
	}
	defer session.Close()

	return &db
}

func tearDownDB(t *testing.T, db *Database) {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		t.Error(err)
	}
	defer session.Close()

	err = session.DB(db.DatabaseName).DropDatabase()
	if err != nil {
		t.Error(err)
	}
}

func TestAdd(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count(db.CollectionState) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	var sarray []interface{}
	testState := State{"A", "B", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	testState1 := State{"D", "E", "F", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	sarray = append(sarray, testState)
	sarray = append(sarray, testState1)

	err := db.Add(sarray, db.CollectionState)

	if err != nil {
		t.Error(err)
	}
	if db.Count(db.CollectionState) != 2 {
		t.Error("Database not properly initialized, database count should be 1")
	}
}

func TestGetFlight(t *testing.T) {
	// Setup
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count(db.CollectionFlight) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	var testFlightArray []interface{}
	testFlight := Flight{"A", 1, "B", 1, "Reku", "C"}
	testFlight1 := Flight{"D", 1, "E", 1, "Rek", "F"}
	testFlight2 := Flight{"G", 1, "H", 1, "Reku", "I"}
	testFlightArray = append(testFlightArray, testFlight)
	testFlightArray = append(testFlightArray, testFlight1)
	testFlightArray = append(testFlightArray, testFlight2)

	err := db.Add(testFlightArray, db.CollectionFlight)

	if err != nil {
		t.Error(err)
	}

	if DBValues.Count(db.CollectionFlight) >= 1 {
		t.Error("Database not properly initialized, database count should be 1")
	}

	// Actual test

	FindData := bson.M{"estarrivalairport": "Reku"}

	a, err := db.GetFlight(FindData)
	if err != nil {
		t.Errorf("Error in retrival of Country, %s", err)
	}

	if a[0] != testFlight {
		t.Error("Incorrect airport was returned")
	}

	if len(a) != 2 {
		t.Error("Incorrect airport was returned")
	}
}

func TestGetState(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count(db.CollectionState) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	testState := State{"A", "B", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	testState1 := State{"D", "E", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	testState2 := State{"F", "G", "H", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	var testStateArray []interface{}
	testStateArray = append(testStateArray, testState)
	testStateArray = append(testStateArray, testState1)
	testStateArray = append(testStateArray, testState2)

	err := db.Add(testStateArray, db.CollectionState)
	if err != nil {
		t.Error(err)
	}

	if db.Count(db.CollectionState) != 3 {
		t.Error("Database not properly initialized, database count should be 3")
	}

	FindData := bson.M{"origincountry": "C"}

	s, err := db.GetState(FindData)
	if err != nil {
		t.Errorf("Error in retrival of value, %s", err)
	}

	if s[0] != testState {
		t.Error("Incorrect state was returned")
	}

	if len(s) != 2 {
		t.Error("Incorrect number of states were returned!")
	}
}

func TestGetAirport(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count(db.CollectionAirport) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	testAirport1 := Airport{1, "Gjovik Airport", "Gjovik", "Mekka", "GJO", "GJOV", float64(10), float64(24), float64(500), "100", "E", "Norway/Oslo", "airport", "test"}
	testAirport2 := Airport{2, "Bardufoss Airport", "Bardufoss", "Norway", "BAR", "BARD", float64(10), float64(24), float64(500), "100", "E", "Norway/Oslo", "airport", "test"}
	testAirport3 := Airport{3, "Molvik Airport", "Molvik", "Norway", "MOL", "MOLV", float64(10), float64(24), float64(500), "100", "E", "Norway/Oslo", "airport", "test"}

	var d []interface{}
	d = append(d, testAirport1)
	d = append(d, testAirport2)
	d = append(d, testAirport3)

	err := db.Add(d, db.CollectionAirport)

	if err != nil {
		t.Error("There should not have been any errors!")
	}

	FindData := bson.M{"country": "Norway"}

	a, err := db.GetAirport(FindData)
	if err != nil {
		t.Errorf("Error in retrival of Country, %s", err)
	}

	if a[0] != testAirport2 {
		t.Error("Incorrect airport was returned")
	}

	if len(a) != 2 {
		t.Error("Incorrect number of airports was returned")
	}
}

func TestGetPlanes(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count(db.CollectionState) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	testState := State{"A", "B", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	testState1 := State{"D", "E", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	testState2 := State{"G", "H", "I", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}

	var testStateArray []interface{}
	testStateArray = append(testStateArray, testState)
	testStateArray = append(testStateArray, testState1)
	testStateArray = append(testStateArray, testState2)

	p := Planes{testState, Flight{}}
	err := db.Add(testStateArray, db.CollectionState)
	if err != nil {
		t.Error(err)
	}

	if db.Count(db.CollectionState) != 3 {
		t.Error("Database not properly initialized, database count should be 3")
	}
	FindData := bson.M{"origincountry": "C"}

	s, err := db.GetPlanes(FindData)
	if err != nil {
		t.Errorf("Error in retrival of OriginCountry, %s", err)
	}

	if s[0] != p {
		t.Error("Incorrect State were returned")
	}

	if len(s) != 2 {
		t.Error("Incorrect number of State were returned")
	}
}
