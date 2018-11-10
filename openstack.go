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
		resp, err := http.Get("https://opensky-network.org/api/states/all")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Enter loop")
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var state States
		if err = json.Unmarshal(body, &state); err != nil {
			fmt.Println(err)
		}
		var sarray []interface{}
		for i := range state.States {
			sarray = append(sarray, state.States[i])
		}
		err = DBValues.Add(sarray, DBValues.CollectionName)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(30 * time.Second)
		fmt.Println("Next update in 15 min")
	}
}
