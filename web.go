package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	log.Print("Starting Golang Rest Showcase...")
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/basic", BasicHandler)
	http.HandleFunc("/showcase", ShowcaseHandler)

	http.ListenAndServe("localhost:9000", nil)
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "POST to either /basic or /showcase. See project readme for more information.")
}

func BasicHandler(res http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	fmt.Fprintln(res, strings.ToUpper(string(reqBody)))
}

func ShowcaseHandler(res http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	log.Printf("request body: %s", string(reqBody))
	var dat map[string]interface{}
	if err := json.Unmarshal(reqBody, &dat); err != nil {
		panic(err)
	}

	input := dat["input"].(string)

	hash := md5.New()
	io.WriteString(hash, input)
	md5 := strings.ToUpper(fmt.Sprintf("%x", hash.Sum(nil)))

	md5Out := MD5Out{md5, input}

	jsonOut, _ := json.Marshal(md5Out)

	fmt.Fprint(res, string(jsonOut))
}

type MD5Out struct {
	MD5            string `json:"md5"`
	OriginalString string `json:"originalString"`
}
