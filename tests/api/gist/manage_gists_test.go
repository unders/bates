package api

import (
	"testing"
	"net/http"
	"time"
	"github.com/unders/bates/tests/api/testdata"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

// Gist Paths
const (
	endpoint = "https://api.github.com"
	gists = endpoint + "/gists"
)

var client = &http.Client{Timeout: time.Second *4}

func TestGistAPI(t *testing.T) {
	tc := newTestCase(t)

	t.Log("Feature: Manage Gists")
	t.Log("    In order to revision small files in the cloud")
	t.Log("    As a user")
	t.Log("    I want to manage a gists store")

	t.Log("Scenario 1: Fetch a list of gists")
	t.Run(fmt.Sprintf("GET  %s", gists), tc.getGists)
}

type testCase struct {
	client *http.Client
}

func newTestCase(t *testing.T) *testCase {
	return &testCase{ }
}

type gistResponse struct {
	ID string `json:"id"`
	Url string `json:"url"`
}

func (g gistResponse) String() string {
	return fmt.Sprintf("{\n    ID: %s\n    Url: %s\n}", g.ID, g.Url)
}

func (tc *testCase) getGists(t *testing. T) {
	resp, err := client.Get(gists)
	testdata.AssertNil(t, err, "")
	defer resp.Body.Close()

	if 200 != resp.StatusCode {
		t.Errorf("\nWant: 200\nGot: %d\n", resp.StatusCode)
	}

	jsonBody, err := ioutil.ReadAll(resp.Body)
	testdata.AssertValidJSON(t, jsonBody)

	gists := make([]gistResponse, 0, 0)
	err = json.Unmarshal(jsonBody, &gists)
	testdata.AssertNil(t, err, testdata.JSONPrettify(jsonBody))

	if len(gists) < 1 {
		t.Fatal("\nWant: a list of items\n Got: 0 items\n")
	}

	for i, g := range gists {
		if "" == g.ID {
			t.Errorf("\nindex:%d\n Want: <ID>\n  Got: '%s'\n", i, g.ID)
		}
		if "" == g.Url {
			t.Errorf("\nindex:%d\n Want: <URL>\n  Got: '%s'\n", i, g.Url)
		}
	}
}