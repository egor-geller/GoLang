package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "User-Agent= %q\n", r.UserAgent())
	fmt.Fprintf(w, "Request-Uri = %q\n", r.RequestURI)
	fmt.Fprintln(w, "Headers:")
	fmt.Fprintf(w, "Accept = %q\n", r.Header.Get("Accept"))
	fmt.Fprintf(w, "User-Agent = %q\n", r.UserAgent())
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
