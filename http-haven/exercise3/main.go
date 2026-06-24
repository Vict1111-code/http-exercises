package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/count" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "Send a POST request with text (1 to 20) to count words")
		return

	case http.MethodPost:
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusBadRequest)
			return
		}

		count := 0
		for range string(b) {
			count++ 
		}
		w.Write([]byte(strconv.Itoa((count)) + "\n"))
		return

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	http.HandleFunc("/count", countHandler)
	fmt.Println("Server is running on http://localhost:8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
