package render

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

func writeJSON(w http.ResponseWriter, status int, values interface{}) {
	json, err := json.Marshal(values)
	if err != nil {
		ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
}

func writeError(w http.ResponseWriter, status int, e error) {
	etype := reflect.TypeOf(e).String()
	msg := fmt.Sprintf("%s: %s", etype, e.Error())
	writeJSON(w, status, map[string]string{"error": msg})
}

// ValidResult renders a valid `values` object as JSON with 200 OK.
func ValidResult(w http.ResponseWriter, values interface{}) {
	writeJSON(w, http.StatusOK, values)
}

// ClientError renders the message of an `error` as JSON with 400 Bad Request.
func ClientError(w http.ResponseWriter, e error) {
	writeError(w, http.StatusBadRequest, e)
}

// ServerError renders the message of an `error` as JSON with 500 Internal Server Error.
func ServerError(w http.ResponseWriter, e error) {
	writeError(w, http.StatusInternalServerError, e)
}

// NotFound renders the message of an `error` as JSON with 404 Not Found.
func NotFound(w http.ResponseWriter, e error) {
	writeError(w, http.StatusNotFound, e)
}
