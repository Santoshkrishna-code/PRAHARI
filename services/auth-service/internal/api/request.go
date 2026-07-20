package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// DecodeJSON parses and decodes the HTTP request body payload into the target interface.
func DecodeJSON(r *http.Request, target interface{}) error {
	defer r.Body.Close()
	
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Security standard to block unmapped parameters injection
	
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("malformed JSON body request: %w", err)
	}
	
	return nil
}
