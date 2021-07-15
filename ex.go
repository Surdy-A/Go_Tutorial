package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"github.com/gorilla/mux"

	
)

func main() {
	newMux := http.NewServeMux()

	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Intn(100))
	})

	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Intn(100))
	})
	http.ListenAndServe(":8080", newMux)
}