package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerGabiTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You SUCCESS DEPLOY GOLANG page Gabi Test</h1>")
}

func HandlerGabiTestPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "{'message': 'OK'}")
}

func Gabi() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandlerGabiTest).Methods("GET")
	r.HandleFunc("/create", HandlerGabiTestPost).Methods("POST")

	http.ListenAndServe(":8080", r)
}
