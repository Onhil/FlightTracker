package main

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) { // WIP
	var flights []Flight

	if err := json.Unmarshal(body(timeFlights()), &flights); err != nil {
		t.Error(err)
	}
}
