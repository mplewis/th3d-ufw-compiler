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

	"./internal/render"
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

func requestHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.ServerError(w, err)
		return
	}

	cr, err := validateAndParse(data)
	if err != nil {
		render.ClientError(w, err)
		return
	}

	hex, err := compile(cr)
	if err != nil {
		render.ServerError(w, err)
		return
	}

	render.ValidResult(w, map[string]string{"compiled_hex": hex})
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
