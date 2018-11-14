package main

import (
	"github.com/globalsign/mgo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlaneHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(PlaneHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Errorf("Error creating the POST request, %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
	}
}

func TestOriginCountryHandler(t *testing.T) {
	// Starts the database
	db := setupDB(t)
	defer tearDownDB(t, db)

	DBValues = Database{
		HostURL:           "mongodb://localhost",
		DatabaseName:      "testing",
		CollectionState:   "testState",
		CollectionAirport: "testAirport",
		CollectionFlight:  "testFlight",
	}

	db.Init()
	if db.Count(db.CollectionState) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	// adds state with country as one of its values
	testState := State{"A", "B", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	var testStateArray []interface{}
	testStateArray = append(testStateArray, testState)

	err := db.Add(testStateArray, db.CollectionState)
	if err != nil {
		t.Error(err)
	}

	if db.Count(db.CollectionState) != 1 {
		t.Error("Database not properly initialized, database count should be 1")
	}

	// Actual test of the handler in question
	ts := httptest.NewServer(http.HandlerFunc(OriginCountryHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/C")

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}
	 // TODO: either fix this test if it is wrong or fix the code so it returns an error when it should
		narr, err := http.Get(ts.URL + "/lasdfkjhfkjhb")

		if narr.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected StatusCode %d, received %d", http.StatusBadRequest, narr.StatusCode)
		}

		if err != nil {
			t.Error(err)
		}

}

func TestDepartureHandler(t *testing.T) {

}

func TestArrivalHandler(t *testing.T) {

	DBValues = Database{
		HostURL:           "mongodb://localhost",
		DatabaseName:      "testing",
		CollectionState:   "testState",
		CollectionAirport: "testAirport",
		CollectionFlight:  "testFlight",
	}

	session, err := mgo.Dial(DBValues.HostURL)
	if err != nil {
		t.Error(err)
	}
	defer session.Close()

	defer tearDownDB(t, &DBValues)

	DBValues.Init()
	if DBValues.Count(DBValues.CollectionFlight) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	testFlight := Flight{"A", 1, "B", 1, "Reku", "C"}
	var testStateArray []interface{}
	testStateArray = append(testStateArray, testFlight)

	err = DBValues.Add(testStateArray, DBValues.CollectionFlight)
	if err != nil {
		t.Error(err)
	}

	if DBValues.Count(DBValues.CollectionFlight) != 1 {
		t.Error("Database not properly initialized, database count should be 1")
	}

	ts := httptest.NewServer(http.HandlerFunc(ArrivalHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/Reku")
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}

	resp, err = http.Get(ts.URL + "/djkfjkndfjkfd")
	if resp.StatusCode == http.StatusOK {
		// TODO: add error on incorrect value in database.go/GetPlanes
		t.Errorf("Expected StatusCode %d, received %d", http.StatusBadRequest, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}
}

func TestPlaneListHandler(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count(db.CollectionState) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	testState := State{"A", "B", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	var testStateArray []interface{}
	testStateArray = append(testStateArray, testState)

	err := db.Add(testStateArray, db.CollectionState)
	if err != nil {
		t.Error(err)
	}

	if db.Count(db.CollectionState) != 1 {
		t.Error("Database not properly initialized, database count should be 1")
	}

	ts := httptest.NewServer(http.HandlerFunc(PlaneListHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "")
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}
}
