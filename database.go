package main

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type database struct {
	HostURL string
	DatabaseName string
	CollectionName string
}

func (db *database) Init() {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()
}

func (db *database) Add(s state) {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.CollectionName).Insert(s)
	if err != nil {
		fmt.Printf("error in Insert(): %v", err.Error())
	}
}

func (db *database) AddMany(s []state) {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.DB(db.DatabaseName).C(db.CollectionName).Bulk().Insert(s)
}

func (db *database) Count() int {
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

func (db *database) GetICAO24(keyID string) (state, bool) {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	State := state{}
	err = session.DB(db.DatabaseName).C(db.CollectionName).Find(bson.M{"icao24": keyID}).One(&State)
	if err != nil {
		return State, false
	}

	return State, true
}

func (db *database) GetOriginCountry(keyID string) ([]state, bool) {
	session, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	State := []state{}
	err = session.DB(db.DatabaseName).C(db.CollectionName).Find(bson.M{"origincountry": keyID}).All(&State)
	if err != nil {
		return State, false
	}

	return State, true
}