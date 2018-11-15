package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	)

func TestPlaneHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(PlaneHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/")
	if err != nil {
	t.Errorf("Error creating the POST request, %s", err)
	}
	if resp.StatusCode != http.StatusOK {
	t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
	}
}

func TestCountryMapHandler(t *testing.T) { // The function to be tested is not yet implemented

}

func TestPlaneMapHandler(t *testing.T) { // The function to be tested is not yet implemented

}