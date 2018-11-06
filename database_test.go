package main

import (
	"github.com/globalsign/mgo"
	"testing"
)

func setupDB(t *testing.T) *Database {
	db := Database{
		"mongodb://localhost",
		"testTrackDB",
		"Tracks",
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

}

func TestAddMany(t *testing.T) {

}

func TestCount(t *testing.T) {

}

func TestGetICAO24(t *testing.T) {

}

func TestGetOriginCountry(t *testing.T) {

}
