package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	var resp *http.Response
	var err error
	if resp, err = http.Get("http://token.fastly.com/token"); err != nil {
		t.Error("Error reported when retrieving token from Fastly service", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	token, err := GenerateToken("RmFzdGx5IFRva2VuIFRlc3Q=", 60)
	if err != nil {
		t.Error("Error while generating token", err)
	}

	if token != string(body) {
		t.Errorf("Expected token: %s, Actual Token: %s", body, token)
	}
}

func BenchmarkGenerateToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateToken("RmFzdGx5IFRva2VuIFRlc3Q=", 60)
	}
}
