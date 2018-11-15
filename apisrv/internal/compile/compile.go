package compile

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"../structs"
)

var configHeaderPath = path.Join("/build", "src", "Configuration.h")
var pioConfigPath = path.Join("/build", "platformio.ini")
var firmwarePath = path.Join("/build", ".pioenvs", "printer", "firmware.hex")

// Compile compiles firmware requested by a CompileRequest.
func Compile(cr structs.CompileRequest) (string, error) {
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
