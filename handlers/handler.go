package handlers

import (
	"fmt"
	"net/http"
)

func GetGreeting(w http.ResponseWriter, r *http.Request) {
	greeting := "Hello, World!"
	fmt.Fprintln(w, greeting)
}
