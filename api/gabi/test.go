package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerGabiTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You SUCCESS DEPLOY GOLANG page Gabi Test</h1>")
}

func Gabi() {
	r := mux.NewRouter()
	r.HandleFunc("/gabi/test", HandlerGabiTest).Methods("GET")

	http.ListenAndServe(":8080", r)
}
