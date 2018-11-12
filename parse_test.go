package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) { // WIP
	var state States

	if err := json.Unmarshal(body("https://opensky-network.org/api/states/all"), &state); err != nil {
		fmt.Println(err)
	}
}
