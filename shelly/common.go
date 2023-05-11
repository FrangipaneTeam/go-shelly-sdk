package shelly

import (
	"encoding/json"
)

// Empty is an empty struct.
type Empty struct{}

type responseReader struct {
	Response interface{}
}

// NewResponseReader creates a new response reader.
func NewResponseReader() *responseReader {
	return &responseReader{
		Response: interface{}(nil),
	}
}

// Read read the response into the given interface.
func (r *responseReader) Read(dst interface{}) error {
	return readResponse(r.Response, dst)
}

// read reads the response into the given interface.
func readResponse(src, dst interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, dst)
}
