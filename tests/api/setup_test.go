package api

import (
	"testing"
	"os"
	"flag"
)

func TestSetup(t *testing.T) {
	t.Log("All API tests should be in sub folders")
}

func TestMain(m *testing.M) {
	// flags must be duplicated in testdata/setup.go file
	_ = flag.Bool("update", false, "update golden files")
	_ = flag.Int("truncate", 0, "truncate log outputs")
	//flag.Parse()

	//setup()
	code := m.Run()
	//shutdown()
	os.Exit(code)
}
