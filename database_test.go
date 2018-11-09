package main

import (
	"fmt"
	"testing"

	"github.com/globalsign/mgo"
)

func setupDB(t *testing.T) *Database {
	db := Database{
		"mongodb://user1:pas123@ds155823.mlab.com:55823/testing",
		"testing",
		"testdata",
	}

	session, err := mgo.Dial(db.HostURL)
	defer session.Close()

	if err != nil {
		t.Error(err)
	}

	return &db
}

func tearDownDB(t *testing.T, db *Database) {
	session, err := mgo.Dial(db.HostURL)
	defer session.Close()

	if err != nil {
		t.Error(err)
	}

	err = session.DB(db.DatabaseName).DropDatabase()
	if err != nil {
		t.Error(err)
	}
}

func TestAdd(t *testing.T) {
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
}

func TestAddDuplicates(t *testing.T) {
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
}

func TestAddMany(t *testing.T) {
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

	err := db.AddMany(d)

	if err != nil {
		t.Error("There should not have been any errors!")
	}

	if db.Count() != 3 {
		fmt.Print(db.Count()) // DEBUG
		fmt.Print("\n")       // DEBUG
		t.Error("Database not properly initialized, database count should be 3")
	}

}

func TestAddManyDuplicates(t *testing.T) {
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

}

func TestGetICAO24(t *testing.T) {
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

	s, ok := db.GetOriginCountry("C")
	if !ok {
		t.Error("Error in retrival of OriginCountry")
	}
	if s[0] != testState {
		t.Error("Incorrect State were returned")
	}
}

func TestUpdateState(t *testing.T) {
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

}
