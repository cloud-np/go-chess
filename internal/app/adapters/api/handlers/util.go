package handlers

import (
	"encoding/json"
	"io"
)


func readJSON(body io.Reader, input interface{}) error {
	return json.NewDecoder(body).Decode(input)
}
