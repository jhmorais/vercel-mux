package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetID() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandlerGetClientByID).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func HandlerGetClientByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>CREATED Client by id</h1>")
}
