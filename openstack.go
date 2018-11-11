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
		go updateStates()
		go updateFlights()
		fmt.Println("Next update in 15 min")
		time.Sleep(15 * time.Minute)
	}
}

func updateStates() {
	var state States

	if err := json.Unmarshal(body("https://opensky-network.org/api/states/all"), &state); err != nil {
		fmt.Println(err)
	}
	var sarray []interface{}
	for i := range state.States {
		sarray = append(sarray, state.States[i])
	}
	err := DBValues.Add(sarray, DBValues.CollectionState)
	if err != nil {
		fmt.Println(err)
	}
}
func updateFlights() {
	var flights []Flight

	if err := json.Unmarshal(body(timeFlights()), &flights); err != nil {
		fmt.Println(err)
	}
	var sarray []interface{}
	for i := range flights {
		sarray = append(sarray, flights[i])
	}
	err := DBValues.Add(sarray, DBValues.CollectionFlight)
	if err != nil {
		fmt.Println(err)
	}
}

func body(url string) []byte {
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
