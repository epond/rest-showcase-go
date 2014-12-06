package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting...")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)

	http.ListenAndServe(":80", nil)
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "This is home")
}
