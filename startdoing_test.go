package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	ts := httptest.NewServer(Router())
	defer ts.Close()
	req, _ := http.NewRequest("GET", ts.URL+"/", nil)
	client := http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "Hello, world!" {
		t.Error("Wrong content:", string(body))
	}
}
