package main

import (
	"testing"
)

func TestUnmarshalJSON(t *testing.T) { // WIP
	s := State{}
	err := s.UnmarshalJSON([]byte(`{
		"time": 1541448600,
		"states":
			[[
				"ab1644",
				"UAL1254 ",
				"United States",
				1541448598,
				1541448599,
				-84.8207,
				38.5694,
				11262.36,
				false,
				274.2,
				36.76,
				0,
				null,
				11513.82,
				"5226",
				false,
				0
			]]}`))
	if err != nil {
		// t.Error(err)
	}
}
