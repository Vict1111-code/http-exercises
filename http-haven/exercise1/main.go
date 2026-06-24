package main

import (
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong\n"))
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
