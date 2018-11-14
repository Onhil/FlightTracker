package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	var state States

	if err := json.Unmarshal(Body("https://opensky-network.org/api/states/all"), &state); err != nil {
		fmt.Println(err)
	}
}
