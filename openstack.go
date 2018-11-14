package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Run _
func Run() {

	for {
		// Concurrent updating of states and flights
		// This might go horribly wrong at some point
		updateStates()
		fmt.Println("States")
		updateFlights()
		fmt.Println("States")
		updateAirports()
		fmt.Println("Airports")
		fmt.Println("Next update in 15 min")
		time.Sleep(15 * time.Minute)
	}
}

func updateStates() {
	var state States

	if err := json.Unmarshal(Body("https://opensky-network.org/api/states/all"), &state); err != nil {
		fmt.Println(err)
	}

	var documents []interface{}
	for i := range state.States {
		documents = append(documents, state.States[i])
	}
	if err := DBValues.Add(documents, DBValues.CollectionState); err != nil {
		fmt.Println(err)
	}
}
func updateFlights() {
	var flights []Flight

	if err := json.Unmarshal(Body(timeFlights()), &flights); err != nil {
		fmt.Println(err)
	}

	var documents []interface{}
	for i := range flights {
		documents = append(documents, flights[i])
	}
	if err := DBValues.Add(documents, DBValues.CollectionFlight); err != nil {
		fmt.Println(err)
	}
}

func updateAirports() {
	var airports []Airport

	if err := json.Unmarshal(Body("https://raw.githubusercontent.com/Onhil/FlightTracker/master/Airports.json"), &airports); err != nil {
		fmt.Println(err)
	}
	var documents []interface{}
	for i := range airports {
		documents = append(documents, airports[i])
	}
	if err := DBValues.Add(documents, DBValues.CollectionAirport); err != nil {
		fmt.Println(err)
	}
}

// Body gets the body from the url and returns it in []byte format
func Body(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	return body
}
