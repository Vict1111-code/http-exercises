package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Query().Get("name") == "" {
		fmt.Fprintln(w, "Hello, Geust!")
		return
	}

	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(fmt.Appendf(nil, "Hello, %s!\n", r.URL.Query().Get("name")))
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server is running on http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
