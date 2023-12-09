package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Update() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandlerUpdateClient).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func HandlerUpdateClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Updated Client</h1>")
}
