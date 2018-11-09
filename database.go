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

	err = session.DB(db.DatabaseName).C(db.CollectionName).EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

// AddAirport adds all the airport data to database
func (db *Database) AddAirport(a []Airport) error {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	for _, t := range a {
		err = session.DB(db.DatabaseName).C(db.CollectionName).Insert(t)
		if err != nil {
			return err
		}
	}

	return err
}

// Count Counts the elements in the database
func (db *Database) Count() int {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	count, err := session.DB(db.DatabaseName).C(db.CollectionName).Count()
	if err != nil {
		fmt.Printf("error in Count(): %v", err.Error())
		return -1
	}

	return count
}

//CountAirport counts the elements in airport database
func (db *Database) CountAirport() int {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	count, err := session.DB(db.DatabaseName).C(db.CollectionAirport).Count()
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

	State := State{}
	err = session.DB(db.DatabaseName).C(db.CollectionName).Find(bson.M{"icao24": keyID}).One(&State)
	if err != nil {
		return State, false
	}

	return State, true
}

// GetOriginCountry returns the origin country if it is in the database and an empty object and false if not
func (db *Database) GetOriginCountry(keyID string) ([]State, bool) {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	State := []State{}
	err = session.DB(db.DatabaseName).C(db.CollectionName).Find(bson.M{"origincountry": keyID}).All(&State)
	if err != nil {
		return State, false
	}

	return State, true
}

// GetAirport returns airport after ICAO code and true if in database, and an empty object and false if not
func (db *Database) GetAirport(keyID string) (Airport, bool) {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	Port := Airport{}
	err = session.DB(db.DatabaseName).C(db.CollectionAirport).Find(bson.M{"icao": keyID}).One(&Port)
	if err != nil {
		return Port, false
	}

	return Port, true
}

// UpdateState tries to updates a State and returns an error if that is not possible
func (db *Database) UpdateState(sarray []interface{}) error {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.DB(db.DatabaseName).C(db.CollectionName).RemoveAll(nil)

	err = session.DB(db.DatabaseName).C(db.CollectionName).Insert(sarray...)

	return err
}
