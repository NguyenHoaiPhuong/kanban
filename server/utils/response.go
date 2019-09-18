package utils

import (
	"encoding/json"
	"net/http"
)

// RespondError writes the error messages
func RespondError(w http.ResponseWriter, status int, messages ...string) {
	w.WriteHeader(status)
	w.Write([]byte(messages))
}

// RespondJSON writes the object
func RespondJSON(w http.ResponseWriter, status int, object interface{}) error {
	bs, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		return err
	}

	w.WriteHeader(status)
	w.Write(bs)
	return nil
}
