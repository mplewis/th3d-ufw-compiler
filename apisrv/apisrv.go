package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"./internal/compile"
	"./internal/parse"
	"./internal/render"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.ServerError(w, err)
		return
	}

	cr, err := parse.Validate(data)
	if err != nil {
		render.ClientError(w, err)
		return
	}

	hex, err := compile.Compile(cr)
	if err != nil {
		render.ServerError(w, err)
		return
	}

	render.ValidResult(w, map[string]string{"compiled_hex": hex})
}

func main() {
	http.HandleFunc("/", requestHandler)
	log.Print("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
