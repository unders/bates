package testdata

import "flag"

// Flags:
//          go test github.com/unders/bates/tests/api/gist -update
//          go test github.com/unders/bates/tests/api/gist -truncate 500
//

// Config is the global config
var Config = struct{
	Update bool
	Truncate int
}{

}
func init() {
	// flags must be duplicated in setup_test.go file
	var update = flag.Bool("update", false, "update golden files")
	var truncate = flag.Int("truncate", 0, "truncate outputs")
	flag.Parse()

	Config.Update = *update
	Config.Truncate = *truncate
}
