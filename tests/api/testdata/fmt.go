package testdata

import (
	"bytes"
	"encoding/json"
)

// JSONPrettify makes json more readable
func JSONPrettify(in []byte) string {
	max := Config.Truncate
	b := &bytes.Buffer{}
	json.Indent(b, in, "", "    ")

	if max == 0 {
		return b.String()
	}

	bb := b.Bytes()
	if len(bb) <  max{
		return string(bb[0:len(bb)])
	}

	return string(bb[0:max]) + " ...\n"
}
