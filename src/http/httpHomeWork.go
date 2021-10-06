package main

import (
	"log"
	"net/http"
	"os"
)

// cmd line test as : curl -H 'colin:me' http://localhost/healthz

func main() {
	http.HandleFunc("/healthz", homework)
	err := http.ListenAndServe(":80", nil)
	log.Fatal(err)
}

func homework(w http.ResponseWriter, r *http.Request) {
	rHaeder := r.Header.Clone()
	customHeader := rHaeder.Get("colin")
	w.Header().Set("colin", customHeader)
	statusCode := http.StatusOK
	versionString := "VERSION"
	err := os.Setenv("VERSION", "VERSION-VALUE")
	w.Header().Set(versionString, os.Getenv(versionString))
	w.WriteHeader(statusCode)
	if err != nil {
		statusCode = http.StatusInternalServerError
		w.WriteHeader(statusCode)
	}
	log.Printf("custom header : %s \n", customHeader)
	log.Printf("VERSION : %s \n", os.Getenv(versionString))
	log.Printf("statusCode : %d \n", statusCode)
	log.Printf("request IP : %s \n", r.RemoteAddr)
}
