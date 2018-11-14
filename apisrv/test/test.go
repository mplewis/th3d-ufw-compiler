package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	ch, err := ioutil.ReadFile("../../cache_helpers/Configuration.h")
	if err != nil {
		panic(err)
	}

	pc, err := ioutil.ReadFile("../../cache_helpers/platformio.ini")
	if err != nil {
		panic(err)
	}

	obj := map[string]string{
		"config_header": string(ch),
		"pio_config":    string(pc),
	}
	jsonb, _ := json.Marshal(obj)
	fmt.Println("Body:", string(jsonb))

	req, err := http.NewRequest("POST", "http://localhost:8080", bytes.NewBuffer(jsonb))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	fmt.Println("Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Body:", string(body))
}
