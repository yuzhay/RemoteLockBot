package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./remotelock"
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

	var response = remotelock.Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		notifyAllUsers("asd" + message)
		return
	}
	notifyAllUsers(parse(&response))
}

func notifyAllUsers(message string) {
	var ids []int64
	db.Table("users").Pluck("id", &ids)
	for _, id := range ids {
		sendMessage(id, message)
	}
}

func parse(response *remotelock.Response) string {
	if response == nil {
		return ""
	}
	data := response.Data

	message := "Unknown type event"
	switch data.Type {
	case "locked_event":
		message = remotelock.LockedEventDecorator(data)
	case "unlocked_event":
		message = remotelock.UnlockedEventDecorator(data)
	}
	return message
}
