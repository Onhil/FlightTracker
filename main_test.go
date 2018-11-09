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

func TestOriginCountryHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(PlaneHandler))
	defer ts.Close()

	// resp, err := http.Get(ts.URL + "/country/{country:[A-Za-z_ ]+}")

}

func TestDepartureHandler(t *testing.T) {

}

func TestArrivalHandler(t *testing.T) {

}
