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
	log.Printf("Server started: https://%s:%d", host, port)
	http.ListenAndServeTLS(fmt.Sprintf(":%d", port), certFile, keyFile, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error: reading request body", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "OK")

	message := fmt.Sprintf(
		"Host: %s,\nCookies: %s,\nMethod: %s,\nBody: %s,\nReferer: %s\nRemoteAddr: %s,\nRequestURI: %s",
		r.Host,
		r.Cookies,
		r.Method,
		string(body),
		r.Referer,
		r.RemoteAddr,
		r.RequestURI)

	notifyAllUsers(message)
}

func notifyAllUsers(message string) {
	var ids []int64
	db.Table("users").Pluck("id", &ids)
	for _, id := range ids {
		sendMessage(id, message)
	}
}
