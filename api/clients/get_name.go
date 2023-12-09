package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetName() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandlerGetClientNyName).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func HandlerGetClientNyName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Get Client by name</h1>")
}
