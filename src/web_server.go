package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kabukky/httpscerts"
)

const certFile = "certs/cert.pem"
const keyFile = "certs/key.pem"

func runHttpsServer(host string, port int, bindingPath string) {
	err := httpscerts.Check(certFile, keyFile)

	if err != nil {
		err = httpscerts.Generate(certFile, keyFile, fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			log.Fatal("error: couldn't generate certificate")
		}
	}

	http.HandleFunc(bindingPath, handler)
	log.Printf("Server started: %s:%d", host, port)
	http.ListenAndServeTLS(fmt.Sprintf(":%d", port), certFile, keyFile, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error: reading request body", http.StatusInternalServerError)
		}
		log.Printf("%s", body)
	}
	fmt.Fprintf(w, "OK")
}
