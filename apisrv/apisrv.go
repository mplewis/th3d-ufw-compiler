package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"./internal/compile"
	"./internal/parse"
	"./internal/render"
	"./internal/structs"
)

var requestQueue = make(chan structs.CompileRequest)
var resultQueue = make(chan structs.CompileResult)

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

	requestQueue <- cr
	render.ValidResult(w, map[string]string{"id": cr.ID})
}

func work() {
	for {
		cr := <-requestQueue
		resultQueue <- compile.Compile(cr)
	}
}

func main() {
	go work()
	http.HandleFunc("/", requestHandler)
	log.Print("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
