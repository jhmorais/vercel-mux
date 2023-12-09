package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Create() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandlerCreateClient).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func HandlerCreateClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>CREATED Client</h1>")
}
