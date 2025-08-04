package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("hello world")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
		fmt.Fprintf(w, "Request Method: %s\n", r.Method)
	})
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running on port 8080")
}
