package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
)

var configHeaderPath = path.Join("/build", "src", "Configuration.h")
var pioConfigPath = path.Join("/build", "platformio.ini")
var firmwarePath = path.Join("/build", ".pioenvs", "printer", "firmware.hex")

type compileRequest struct {
	ConfigHeader string `json:"config_header"`
	PioConfig    string `json:"pio_config"`
}

func validateAndParse(raw []byte) (compileRequest, error) {
	var cr compileRequest
	err := json.Unmarshal(raw, &cr)
	if err != nil {
		return cr, err
	}

	if cr.ConfigHeader == "" {
		return cr, errors.New("config_header must not be empty")
	}
	if cr.PioConfig == "" {
		return cr, errors.New("pio_config must not be empty")
	}

	return cr, nil
}

func renderJSON(w http.ResponseWriter, status int, values interface{}) {
	json, err := json.Marshal(values)
	if err != nil {
		renderServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
}

func renderValidResult(w http.ResponseWriter, values interface{}) {
	renderJSON(w, http.StatusOK, values)
}

func renderError(w http.ResponseWriter, status int, e error) {
	renderJSON(w, status, map[string]string{"error": e.Error()})
}

func renderClientError(w http.ResponseWriter, e error) {
	renderError(w, http.StatusBadRequest, e)
}

func renderServerError(w http.ResponseWriter, e error) {
	renderError(w, http.StatusInternalServerError, e)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		renderServerError(w, err)
		return
	}

	cr, err := validateAndParse(data)
	if err != nil {
		renderClientError(w, err)
		return
	}

	hex, err := compile(cr)
	if err != nil {
		renderServerError(w, err)
		return
	}

	renderValidResult(w, map[string]string{"compiled_hex": hex})
}

func compile(cr compileRequest) (string, error) {
	err := ioutil.WriteFile(configHeaderPath, []byte(cr.ConfigHeader), 0644)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(pioConfigPath, []byte(cr.PioConfig), 0644)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("platformio", "run")
	cmd.Dir = "/build"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	hex, err := ioutil.ReadFile(firmwarePath)
	if err != nil {
		return "", err
	}

	return string(hex), nil
}

func main() {
	http.HandleFunc("/", requestHandler)
	log.Print("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
