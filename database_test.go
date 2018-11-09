package main

import (
	"fmt"
	"testing"

	"github.com/globalsign/mgo"
)

func setupDB(t *testing.T) *Database {
	db := Database{
		HostURL:           "mongodb://localhost",
		DatabaseName:      "testing",
		CollectionName:    "testdata",
		CollectionAirport: "testAirport",
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

/// Commented out likely not needed as UpdateState is changed to removing all documents and adding new ones therefore no duplicates

/* func TestAddDuplicates(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count() != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	testState := State{"A", "B", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	err := db.Add(testState)
	if err != nil {
		t.Error(err)
	}
	if db.Count() != 1 {
		t.Error("Database not properly initialized, database count should be 1")
	}

	err = db.Add(testState)
	if err == nil {
		t.Error("An error should have been returned when trying to insert a duplicate element!")
	}
	if db.Count() != 1 {
		t.Error("Duplicate got added, database count should be 1")
	}
} */

/* func TestAddManyDuplicates(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count() != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	testState1 := State{"A", "D", "I", float64(18), float64(12), float64(400), false, float64(250), float64(25), float64(28), float64(31), "a", true}
	testState2 := State{"B", "E", "H", float64(19), float64(13), float64(500), false, float64(251), float64(26), float64(29), float64(32), "b", true}
	testState3 := State{"C", "F", "G", float64(20), float64(14), float64(600), false, float64(252), float64(27), float64(30), float64(33), "c", true}

	d := []State{}
	d = append(d, testState1)
	d = append(d, testState2)
	d = append(d, testState3)

	db.AddMany(d)

	if db.Count() != 3 {
		fmt.Print(db.Count()) // DEBUG
		fmt.Print("\n")       // DEBUG
		t.Error("Database not properly initialized, database count should be 3")
	}

	err := db.AddMany(d)

	if err == nil {
		t.Error("An error should have been returned when trying to insert duplicate elements!")
	}

	if db.Count() != 3 {
		fmt.Print(db.Count()) // DEBUG
		fmt.Print("\n")       // DEBUG
		t.Error("Duplicates added, database count should be 3")
	}

} */

func TestAddAirport(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count(db.CollectionAirport) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	testAirport1 := Airport{1, "Gjovik Airport", "Gjovik", "Norway", "GJO", "GJOV", float64(10), float64(24), float64(500), float64(100), "E", "Norway/Oslo", "airport", "test"}
	testAirport2 := Airport{2, "Bardufoss Airport", "Bardufoss", "Norway", "BAR", "BARD", float64(10), float64(24), float64(500), float64(100), "E", "Norway/Oslo", "airport", "test"}
	testAirport3 := Airport{3, "Molvik Airport", "Molvik", "Norway", "MOL", "MOLV", float64(10), float64(24), float64(500), float64(100), "E", "Norway/Oslo", "airport", "test"}

	var d []interface{}
	d = append(d, testAirport1)
	d = append(d, testAirport2)
	d = append(d, testAirport3)

	err := db.Add(d, db.CollectionAirport)

	if err != nil {
		t.Error("There should not have been any errors!")
	}

	if db.Count(db.CollectionAirport) != 3 {
		fmt.Print(db.Count(db.CollectionAirport)) // DEBUG
		fmt.Print("\n")                           // DEBUG
		t.Error("Database not properly initialized, database count should be 3")
	}
}

func TestGetICAO24(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count(db.CollectionName) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	testState := State{"A", "B", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	var testStateArray []interface{}
	testStateArray = append(testStateArray, testState)

	err := db.Add(testStateArray, db.CollectionName)
	if err != nil {
		t.Error(err)
	}
	if db.Count(db.CollectionName) != 1 {
		t.Error("Database not properly initialized, database count should be 1")
	}

	s, ok := db.GetICAO24("A")
	if !ok {
		t.Error("Error in retrival of ICAO24")
	}
	if s != testState {
		t.Error("Incorrect State were returned")
	}
}

func TestGetOriginCountry(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count(db.CollectionName) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	testState := State{"A", "B", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	var testStateArray []interface{}
	testStateArray = append(testStateArray, testState)

	err := db.Add(testStateArray, db.CollectionName)
	if err != nil {
		t.Error(err)
	}
	if db.Count(db.CollectionName) != 1 {
		t.Error("Database not properly initialized, database count should be 1")
	}

	s, ok := db.GetOriginCountry("C")
	if !ok {
		t.Error("Error in retrival of OriginCountry")
	}
	if s[0] != testState {
		t.Error("Incorrect State were returned")
	}
}

func TestGetAirport(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count(db.CollectionAirport) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	testAirport1 := Airport{1, "Gjovik Airport", "Gjovik", "Norway", "GJO", "GJOV", float64(10), float64(24), float64(500), float64(100), "E", "Norway/Oslo", "airport", "test"}
	testAirport2 := Airport{2, "Bardufoss Airport", "Bardufoss", "Norway", "BAR", "BARD", float64(10), float64(24), float64(500), float64(100), "E", "Norway/Oslo", "airport", "test"}
	testAirport3 := Airport{3, "Molvik Airport", "Molvik", "Norway", "MOL", "MOLV", float64(10), float64(24), float64(500), float64(100), "E", "Norway/Oslo", "airport", "test"}

	var d []interface{}
	d = append(d, testAirport1)
	d = append(d, testAirport2)
	d = append(d, testAirport3)

	err := db.Add(d, db.CollectionAirport)

	if err != nil {
		t.Error("There should not have been any errors!")
	}

	a, ok := db.GetAirport("BARD")
	if !ok {
		t.Error("Error in retrival of OriginCountry")
	}

	if a != testAirport2 {
		t.Error("Incorrect airport was returned")
	}
}

/// Commented out likely not needed as UpdateState is changed to removing all documents and adding new ones

/* func TestUpdateState(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count() != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	testState := State{"A", "B", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	updateState := State{"A", "D", "E", float64(13), float64(22), float64(410), false, float64(250), float64(19), float64(18), float64(16), "", true}
	err := db.Add(testState)
	if err != nil {
		t.Error(err)
	}
	if db.Count() != 1 {
		t.Error("Database not properly initialized, database count should be 1")
	}

	err = db.UpdateState(updateState)
	if err != nil {
		t.Error(err)
	}

	s, ok := db.GetICAO24("A")
	if !ok {
		t.Error("Error in retrival of ICAO24")
	}
	if s != updateState {
		t.Error("State not updated")
	}

} */
