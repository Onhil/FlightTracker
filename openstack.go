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

	resp, _ := http.Get("https://opensky-network.org/api/states/all")
	for {
		body, _ := ioutil.ReadAll(resp.Body)
		var state States
		if err := json.Unmarshal(body, &state); err != nil {
			fmt.Println("error")
		}
		var sarray []interface{}
		for i := range state.States {
			sarray = append(sarray, state.States[i])
		}
		DBValues.UpdateState(sarray)
		time.Sleep(15 * time.Minute)
	}
}
