package main

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Init checks if the Database actually works
func (db *Database) Init() {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	index := mgo.Index{
		Key:        []string{"icao24"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = session.DB(db.DatabaseName).C(db.CollectionState).EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

// Add removes and adds documents to passes collection name
// Example:
// CollectionState
// CollectionAirport
// CollectionFlight
func (db *Database) Add(documents []interface{}, collN string) error {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	_, err = session.DB(db.DatabaseName).C(collN).RemoveAll(nil)
	if err != nil {
		return err
	}

	err = session.DB(db.DatabaseName).C(collN).Insert(documents...)

	return err
}

// Count Counts the documents in a collection
// Example:
// CollectionState
// CollectionAirport
// CollectionFlight
func (db *Database) Count(collN string) int {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	count, err := session.DB(db.DatabaseName).C(collN).Count()
	if err != nil {
		fmt.Printf("error in Count(): %v", err.Error())
		return -1
	}

	return count
}

// GetICAO24 gets the ICA024 from the database or returns false and an empty state object
func (db *Database) GetICAO24(keyID string) (State, bool) {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	state := State{}
	err = session.DB(db.DatabaseName).C(db.CollectionState).Find(bson.M{"icao24": keyID}).One(&state)
	if err != nil {
		return state, false
	}

	return state, true
}

// GetOriginCountry returns the origin country if it is in the database and an empty object and false if not
func (db *Database) GetOriginCountry(keyID string) ([]State, bool) {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	state := []State{}
	err = session.DB(db.DatabaseName).C(db.CollectionState).Find(bson.M{"origincountry": keyID}).All(&state)
	if err != nil {
		return state, false
	}

	return state, true
}

// GetAirport returns airport after ICAO code and true if in database, and an empty object and false if not
func (db *Database) GetAirport(keyID string) (Airport, bool) {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	port := Airport{}
	err = session.DB(db.DatabaseName).C(db.CollectionAirport).Find(bson.M{"icao": keyID}).One(&port)
	if err != nil {
		return port, false
	}

	return port, true
}

func (db *Database) GetAllFlights() ([]State, bool) {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	var states []State

	err = session.DB(db.DatabaseName).C(db.CollectionState).Find(nil).All(&states)

	return states, true

}

func (db *Database) GetFlightFieldData(findData map[string]interface{}) ([]Flight, bool) {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	var flights []Flight

	err = session.DB(db.DatabaseName).C(db.CollectionFlight).Find(findData).All(&flights)

	return flights, true

}
