package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	log.Print("Starting...")
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/basic", BasicHandler)

	http.ListenAndServe("localhost:9000", nil)
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "POST to either /basic or /showcase. See project readme for more information.")
}

func BasicHandler(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	fmt.Fprintln(res, strings.ToUpper(string(body)))
}
