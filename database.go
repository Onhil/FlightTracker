package main

import (
	"errors"
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
//
// Add example
// var documents []interface{}
//	 for i := range flights {
//		 documents = append(documents, flights[i])
// 	 }
// err := DBValues.Add(documents, DBValues.CollectionState)
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

// Getflight accepts bson.M{} to find all flights with choosen paramaters
// Example
// findData == bson.M{"origincountry": "Italy"}
func (db *Database) Getflight(findData bson.M) ([]Flight, error) {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	var flights []Flight
	err = session.DB(db.DatabaseName).C(db.CollectionFlight).Find(findData).All(&flights)

	return flights, nil
}

// GetState accepts bson.M{} to find all flights with choosen paramaters
// Example
// findData == bson.M{"Callsign": "<insert callsign here>"}
func (db *Database) GetState(findData bson.M) ([]State, error) {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	var state []State

	err = session.DB(db.DatabaseName).C(db.CollectionState).Find(findData).All(&state)

	return state, err
}

// GetAirport accepts bson.M{} to find all Airports with choosen paramters
// Example
// FindData == bson.M{"country": "Italy"}
func (db *Database) GetAirport(findData bson.M) ([]Airport, error) {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	var port []Airport

	err = session.DB(db.DatabaseName).C(db.CollectionAirport).Find(findData).One(&port)

	return port, err
}

func (db *Database) GetPlanes(find bson.M) ([]Planes, error) {
	var s []State
	var f []Flight
	var err error

	if s, err = db.GetState(find); err != nil {
		fmt.Println(err)
	}
	if f, err = db.Getflight(find); err != nil {
		fmt.Println(err)
	}
	if s != nil && f != nil {
		// Merges states and flight together
		planes := mergeStatesAndFlights(s, f)
		return planes, nil
	}
	// Returns all flights if find parameter resulted in none
	if s != nil && f == nil {
		f = nil
		if f, err = db.Getflight(nil); err != nil {
			return []Planes{}, nil
		}
		// Returns all states if find parameter resulted in none
	} else if s == nil && f != nil {
		s = nil
		if s, err = db.GetState(nil); err != nil {
			return []Planes{}, nil
		}
	} else {
		return []Planes{}, errors.New("Nothing found in States or Flights with those parameters")
	}

	// Merges states and flight together
	planes := mergeStatesAndFlights(s, f)
	return planes, nil
}
