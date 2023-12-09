package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerGabi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You SUCCESS DEPLOY GOLANG page Gabi</h1>")
}

func HandlerGabiTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You SUCCESS DEPLOY GOLANG page Gabi Test</h1>")
}

func Gabi() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandlerGabi).Methods("GET")
	r.HandleFunc("/test", HandlerGabiTest).Methods("GET")

	http.ListenAndServe(":8080", r)
}
