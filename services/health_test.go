package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthResponseIsOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != BinanceHealthPath {
			t.Fatalf("Expected to request '/%s', got: %s", BinanceHealthPath, r.URL.Path)
		}
		response, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", "../test", "systemStatus.json"))
		if err != nil {
			t.Fatal(err)
		}

		w.Write(response)
	}))
	defer server.Close()
	BinanceEndpoint = server.URL

	err := Health()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, nil, err)
}
