package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kabukky/httpscerts"
)

func main() {
	err := httpscerts.Check("cert.pem", "key.pem")

	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", "0.0.0.0:65000")
		if err != nil {
			log.Fatal("error: couldn't generate certificate")
		}
	}

	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":65000", "cert.pem", "key.pem", nil)
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
