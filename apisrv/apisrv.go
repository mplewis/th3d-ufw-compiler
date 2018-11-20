package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"./internal/compile"
	"./internal/ejectdb"
	"./internal/parse"
	"./internal/render"
	"./internal/structs"
)

const maxResults = 10
const requestBuffer = 200

var requestQueue = make(chan structs.CompileRequest, requestBuffer)
var results = ejectdb.NewEjectDB(maxResults)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getResults(w, r)
	} else if r.Method == http.MethodPost {
		newRequest(w, r)
	} else {
		render.ClientError(w, fmt.Errorf("Unsupported method %s", r.Method))
	}
}

func getResults(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	result, ok := results.Get(id)
	if !ok {
		render.NotFound(w, fmt.Errorf("No result found with ID %s", id))
		return
	}

	var errMsg interface{}
	if result.Error != nil {
		errMsg = result.Error.Error()
	}

	render.ValidResult(w, map[string]interface{}{
		"compiled_hex": result.IntelHex,
		"error":        errMsg,
	})
}

func newRequest(w http.ResponseWriter, r *http.Request) {
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
		creq := <-requestQueue
		cres := compile.Compile(creq)
		results.Put(cres)
		fmt.Println(creq.ID)
	}
}

func fakework() {
	for {
		creq := <-requestQueue
		ns, _ := time.ParseDuration("2s")
		time.Sleep(ns)
		cres := structs.NewCompileSuccess(creq, "fake_hex_string_compiled_result")
		results.Put(cres)
		fmt.Println(creq.ID)
	}
}

func main() {
	go fakework()
	http.HandleFunc("/", requestHandler)
	log.Print("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
