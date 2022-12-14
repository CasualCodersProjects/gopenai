package types

import (
	"bytes"
	"encoding/json"
	"io"
)

func Encode(v interface{}) (io.Reader, error) {
	// encode v into a reader using json
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return io.NopCloser(bytes.NewReader(jsonBytes)), nil
}
