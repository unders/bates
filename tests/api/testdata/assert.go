package testdata

import (
	"testing"
	"encoding/json"
)

// AssertNil asserts that err is nil
func AssertNil(t *testing.T, err error, str string) {
	t.Helper()
	if err != nil {
		if str != "" {
			t.Log(str)
		}
		t.Fatalf("\nWant: nil\n Got: %s\n", err)
	}
}

// AssertErr asserts that err is not nil
func AssertErr(t *testing.T, err error, str string) {
	t.Helper()
	if err == nil {
		if str != "" {
			t.Log(str)
		}
		t.Fatal("\nWant: err\n Got: nil\n")
	}
}

// AssertValidJSON asserts json string is valid
func AssertValidJSON(t *testing.T, input []byte) {
	t.Helper()

	if json.Valid(input) {
		return
	}

	t.Fatalf("\nWant: Valid JSON\n Got:\n %s\n", JSONPrettify(input))
}
