package utils

import (
	"encoding/json"
	"net/http"
)

// RespondError writes the error messages
func RespondError(w http.ResponseWriter, status int, messages ...string) {
	errMsg := ""
	for _, msg := range messages {
		errMsg += msg + "."
	}
	w.WriteHeader(status)
	w.Write([]byte(errMsg))
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
