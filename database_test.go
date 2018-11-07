package main

import (
	"github.com/globalsign/mgo"
	"testing"
)

func setupDB(t *testing.T) *Database {
	db := Database{
		"mongodb://localhost",
		"test",
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

func TestAddMany(t *testing.T) {

}

func TestCount(t *testing.T) {

}

func TestGetICAO24(t *testing.T) {

}

func TestGetOriginCountry(t *testing.T) {

}
