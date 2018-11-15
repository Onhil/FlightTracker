package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	)

func TestPlaneHandler(t *testing.T) {
	DBValues := setupDB(t)
	defer tearDownDB(t, DBValues)

	DBValues.Init()
	if DBValues.Count(DBValues.CollectionState) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	var sarray []interface{}
	testState := State{"A", "B", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	testState1 := State{"D", "E", "F", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	sarray = append(sarray, testState)
	sarray = append(sarray, testState1)

	err := DBValues.Add(sarray, DBValues.CollectionState)

	if err != nil {
		t.Error(err)
	}
	if DBValues.Count(DBValues.CollectionState) != 2 {
		t.Error("Database not properly initialized, database count should be 1")
	}

	var testFlightArray []interface{}
	testFlight := Flight{"A", 1, "B", 1, "Reku", "C"}
	testFlight1 := Flight{"D", 1, "E", 1, "Rek", "F"}
	testFlight2 := Flight{"G", 1, "H", 1, "Reku", "I"}
	testFlightArray = append(testFlightArray, testFlight)
	testFlightArray = append(testFlightArray, testFlight1)
	testFlightArray = append(testFlightArray, testFlight2)

	err = DBValues.Add(testFlightArray, DBValues.CollectionFlight)

	if err != nil {
		t.Error(err)
	}
	if DBValues.Count(DBValues.CollectionFlight) != 3 {
		t.Error("Database not properly initialized, database count should be 1")
	}

	testAirport1 := Airport{1, "Gjovik Airport", "Gjovik", "Mekka", "GJO", "GJOV", float64(10), float64(24), float64(500), "100", "E", "Norway/Oslo", "airport", "test"}
	testAirport2 := Airport{2, "Bardufoss Airport", "Bardufoss", "Norway", "BAR", "BARD", float64(10), float64(24), float64(500), "100", "E", "Norway/Oslo", "airport", "test"}
	testAirport3 := Airport{3, "Molvik Airport", "Molvik", "Norway", "MOL", "MOLV", float64(10), float64(24), float64(500), "100", "E", "Norway/Oslo", "airport", "test"}

	var d []interface{}
	d = append(d, testAirport1)
	d = append(d, testAirport2)
	d = append(d, testAirport3)

	err = DBValues.Add(d, DBValues.CollectionAirport)

	if err != nil {
		t.Error("There should not have been any errors!")
	}

	ts := httptest.NewServer(http.HandlerFunc(PlaneHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "")

	if err != nil {
		t.Errorf("Error creating the POST request, %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
	}
}

func TestCountryMapHandler(t *testing.T) { // The function to be tested is not yet implemented
	DBValues := setupDB(t)
	defer tearDownDB(t, DBValues)

	DBValues.Init()
	if DBValues.Count(DBValues.CollectionState) != 0 {
		t.Error("Database not properly initialized, database count should be 0")
	}

	var sarray []interface{}
	testState := State{"A", "B", "C", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	testState1 := State{"D", "E", "F", float64(18), float64(12), float64(400), false, float64(250), float64(19), float64(18), float64(16), "", true}
	sarray = append(sarray, testState)
	sarray = append(sarray, testState1)

	err := DBValues.Add(sarray, DBValues.CollectionState)

	if err != nil {
		t.Error(err)
	}
	if DBValues.Count(DBValues.CollectionState) != 2 {
		t.Error("Database not properly initialized, database count should be 1")
	}

	var testFlightArray []interface{}
	testFlight := Flight{"A", 1, "B", 1, "Reku", "C"}
	testFlight1 := Flight{"D", 1, "E", 1, "Reku", "F"}
	testFlight2 := Flight{"G", 1, "H", 1, "Reku", "I"}
	testFlightArray = append(testFlightArray, testFlight)
	testFlightArray = append(testFlightArray, testFlight1)
	testFlightArray = append(testFlightArray, testFlight2)

	err = DBValues.Add(testFlightArray, DBValues.CollectionFlight)

	if err != nil {
		t.Error(err)
	}
	if DBValues.Count(DBValues.CollectionFlight) != 3 {
		t.Error("Database not properly initialized, database count should be 1")
	}

	testAirport1 := Airport{1, "Reku", "Gjovik", "Mekka", "GJO", "GJOV", float64(10), float64(24), float64(500), "100", "E", "Norway/Oslo", "airport", "test"}
	testAirport2 := Airport{2, "A", "Bardufoss", "Norway", "BAR", "BARD", float64(10), float64(24), float64(500), "100", "E", "Norway/Oslo", "airport", "test"}
	testAirport3 := Airport{3, "Molvik Airport", "Molvik", "Norway", "MOL", "MOLV", float64(10), float64(24), float64(500), "100", "E", "Norway/Oslo", "airport", "test"}

	var d []interface{}
	d = append(d, testAirport1)
	d = append(d, testAirport2)
	d = append(d, testAirport3)

	err = DBValues.Add(d, DBValues.CollectionAirport)

	if err != nil {
		t.Error("There should not have been any errors!")
	}

	ts := httptest.NewServer(http.HandlerFunc(PlaneMapHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/A")

	if err != nil {
		t.Errorf("Error creating the POST request, %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
	}
	resp, err = http.Get(ts.URL + "/gjkjkdgf")

	if err != nil {
		t.Errorf("Error creating the POST request, %s", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestPlaneMapHandler(t *testing.T) { // The function to be tested is not yet implemented

}