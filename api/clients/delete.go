package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Delete() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandlerDeleteClient).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func HandlerDeleteClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>DELETED Client</h1>")
}
