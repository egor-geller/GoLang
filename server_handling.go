package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Profile struct {
	Name    string   `json:"host"`
	Headers []string `json:"headers"`
	Agent   string   `json:"user-agent"`
	Request string   `json:"request-uri"`
	Accept  string   `json:"accept"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	profile := Profile{r.Host, []string{r.Header.Get("Accept"), r.UserAgent()}, r.UserAgent(), r.RequestURI, r.Header.Get("Accept")}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
