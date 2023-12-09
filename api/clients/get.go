package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Get() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandlerGetClient).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func HandlerGetClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Get Client</h1>")
}
