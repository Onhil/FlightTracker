package main

import (
	"encoding/json"
	"testing"
)

func TestBody(t *testing.T) {
	var state States

	if err := json.Unmarshal(Body("https://opensky-network.org/api/states/all"), &state); err != nil {
		t.Error(err)
	}
}
