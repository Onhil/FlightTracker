package main

import (
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
	DBValues = *setupDB(t)
	defer tearDownDB(t, &DBValues)

	DBValues.Init()
	if DBValues.Count(DBValues.CollectionState) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	// adds state with country as one of its values
	testState := State{"A", "B", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	var testStateArray []interface{}
	testStateArray = append(testStateArray, testState)

	err := DBValues.Add(testStateArray, DBValues.CollectionState)
	if err != nil {
		t.Error(err)
	}

	if DBValues.Count(DBValues.CollectionState) != 1 {
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
	narr, err := http.Get(ts.URL + "/lasdfkjhfkjhb")

	if narr.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusBadRequest, narr.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}

}

func TestDepartureHandler(t *testing.T) {
	DBValues = *setupDB(t)
	defer tearDownDB(t, &DBValues)

	DBValues.Init()
	if DBValues.Count(DBValues.CollectionFlight) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	// adds state with country as one of its values
	testState := Flight{"A", 1, "B", 1, "Reku", "C"}
	var testStateArray []interface{}
	testStateArray = append(testStateArray, testState)

	err := DBValues.Add(testStateArray, DBValues.CollectionFlight)
	if err != nil {
		t.Error(err)
	}

	if DBValues.Count(DBValues.CollectionFlight) != 1 {
		t.Error("Database not properly initialized, database count should be 1")
	}

	// Actual test of the handler in question
	ts := httptest.NewServer(http.HandlerFunc(DepartureHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/B")

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}
	narr, err := http.Get(ts.URL + "/lasdfkjhfkjhb")

	if narr.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusBadRequest, narr.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}

}

func TestArrivalHandler(t *testing.T) {
	DBValues = *setupDB(t)
	defer tearDownDB(t, &DBValues)

	DBValues.Init()
	if DBValues.Count(DBValues.CollectionFlight) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	testFlight := Flight{"A", 1, "B", 1, "Reku", "C"}
	var testStateArray []interface{}
	testStateArray = append(testStateArray, testFlight)

	err := DBValues.Add(testStateArray, DBValues.CollectionFlight)
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
		t.Errorf("Expected StatusCode %d, received %d", http.StatusBadRequest, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}
}

func TestPlaneListHandler(t *testing.T) {
	DBValues = *setupDB(t)
	defer tearDownDB(t, &DBValues)

	DBValues.Init()
	if DBValues.Count(DBValues.CollectionState) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	testState := State{"A", "B", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	var testStateArray []interface{}
	testStateArray = append(testStateArray, testState)

	err := DBValues.Add(testStateArray, DBValues.CollectionState)
	if err != nil {
		t.Error(err)
	}

	if DBValues.Count(DBValues.CollectionState) != 1 {
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

func TestPlaneInfoHandler(t *testing.T) {
	DBValues = *setupDB(t)
	defer tearDownDB(t, &DBValues)

	DBValues.Init()
	if DBValues.Count(DBValues.CollectionState) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}
	testState := State{"A", "B", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	var testStateArray []interface{}
	testStateArray = append(testStateArray, testState)

	err := DBValues.Add(testStateArray, DBValues.CollectionState)

	if err != nil {
		t.Error(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(PlaneInfoHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/A")
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}
	/* // TODO: Make the funciton being tested return an error when it should
	resp, err = http.Get(ts.URL + "/skjahfhjksdfukhj")
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusBadRequest, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}*/
}

func TestPlaneFieldHandler(t *testing.T) {
	DBValues = *setupDB(t)
	defer tearDownDB(t, &DBValues)

	DBValues.Init()
	if DBValues.Count(DBValues.CollectionState) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	// Adds element
	testState := State{"A", "B", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	var testStateArray []interface{}
	testStateArray = append(testStateArray, testState)

	err := DBValues.Add(testStateArray, DBValues.CollectionState)
	if err != nil {
		t.Error(err)
	}

	if DBValues.Count(DBValues.CollectionState) != 1 {
		t.Error("Database not properly initialized, database count should be 1")
	}

	// Actual test
	ts := httptest.NewServer(http.HandlerFunc(PlaneFieldHandler))
	defer ts.Close()
	resp, err := http.Get(ts.URL + "/A/Callsign")
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}

	resp, err = http.Get(ts.URL + "/A/rattattat")
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusBadRequest, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}
}

func TestPlaneMapHandler(t *testing.T) { // The function to be tested is not yet implemented

}

func TestCountryHandler(t *testing.T) {
	DBValues = *setupDB(t)
	defer tearDownDB(t, &DBValues)

	DBValues.Init()
	if DBValues.Count(DBValues.CollectionState) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	// Adds element
	testState := State{"A", "B", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	testState1 := State{"D", "E", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	testState2 := State{"G", "H", "I", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	var testStateArray []interface{}
	testStateArray = append(testStateArray, testState)
	testStateArray = append(testStateArray, testState1)
	testStateArray = append(testStateArray, testState2)

	err := DBValues.Add(testStateArray, DBValues.CollectionState)
	if err != nil {
		t.Error(err)
	}

	if DBValues.Count(DBValues.CollectionState) != 3 {
		t.Error("Database not properly initialized, database count should be 3")
	}

	// Actual test
	ts := httptest.NewServer(http.HandlerFunc(CountryHandler))
	defer ts.Close()
	resp, err := http.Get(ts.URL + "/C")

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}
}

func TestCountryMapHandler(t *testing.T) { // The function to be tested is not yet implemented

}

func TestAirportListHandler(t *testing.T) {
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
	if db.Count(db.CollectionAirport) != 3 {
		t.Error("Database not properly initialized, database count should be 3")
	}

	// Actual test
	ts := httptest.NewServer(http.HandlerFunc(AirportListHandler))
	defer ts.Close()
	resp, err := http.Get(ts.URL + "")

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}
}

func TestAirportInfoHandler(t *testing.T) {
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
	if db.Count(db.CollectionAirport) != 3 {
		t.Error("Database not properly initialized, database count should be 3")
	}

	// Actual test
	ts := httptest.NewServer(http.HandlerFunc(AirportInfoHandler))
	defer ts.Close()
	resp, err := http.Get(ts.URL + "/GJOV")

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}
	/*	// TODO: make AirportInfoHandler handle errors correctly
	resp, err = http.Get(ts.URL + "/jklfdkjd") // error originates here!

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusBadRequest, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}*/
}

func TestAirportFieldHandler(t *testing.T) {
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
	if db.Count(db.CollectionAirport) != 3 {
		t.Error("Database not properly initialized, database count should be 3")
	}

	// Actual test
	ts := httptest.NewServer(http.HandlerFunc(AirportFieldHandler))
	defer ts.Close()
	resp, err := http.Get(ts.URL + "/GJOV/ID")

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}/* // TODO: Make tAirportFieldHandler handle errors correctly
	resp, err = http.Get(ts.URL + "/GJOV/Ijsdkjjk")

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusBadRequest, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}
	resp, err = http.Get(ts.URL + "/fdsfjg/ID")

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusBadRequest, resp.StatusCode)
	}

	if err != nil {
		t.Error(err)
	}*/
}

func TestAirportCountryHandler(t *testing.T) {

}

func TestAirportInCountryHandler(t *testing.T) {

}
