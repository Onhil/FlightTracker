package main

import (
	"testing"
	"fmt"
)

//TestCountries checks if there are duplicate countries in the array
func TestCountries(t *testing.T) {
	allok := true
	//Country := CountryArray()
	for i := 0; i < len(Country) - 1; i++ {
		for j:= i+1; j < len(Country); j++ {
			if Country[i] == Country[j] {
				fmt.Println(Country[i])
				allok = false
			}
		}
	}

	if !allok {
		t.Errorf("There were duplicate countries!")
	}
}