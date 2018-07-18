package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"net/url"
	"strings"
)


func TestCreateBirdHandler(t *testing.T){

	data := url.Values{}
	data.Set("species", "foo")
	data.Add("description", "bar")

	req, err := http.NewRequest("POST", "bird", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(createBirdHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if len(birds) != 1{
		t.Fatal("Bird length isn't right")
	}
}


func TestGetBirdHandler(t *testing.T){
	req, err := http.NewRequest("GET", "bird", nil)


	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(getBirdHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected, err := json.Marshal(birds)
	actual := recorder.Body.String()
	if string(expected) != actual{
		t.Fatal("Strings dont match")
	}
}
