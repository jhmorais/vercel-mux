package handler

import (
	"fmt"
	"myapp/config"
	"myapp/internal/contracts"
	"myapp/internal/infra/di"
	"myapp/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type HandlerCreateClient struct {
	WorkerPort          string
	CreateClientUseCase contracts.CreateClientUseCase
}

func NewHTTPRouterClient(
	createClientUseCase contracts.CreateClientUseCase,
) *mux.Router {
	router := mux.NewRouter()
	handler := HandlerCreateClient{
		CreateClientUseCase: createClientUseCase,
	}
	router.UseEncodedPath()
	router.Use(utils.CommonMiddleware)

	router.HandleFunc("/", handler.Post).Methods("POST")

	return router
}

func Create() {
	config.LoadServerEnvironmentVars()

	dependencies := di.NewBuild()
	router := NewHTTPRouterClient(dependencies.Usecases.CreateClientUseCase)

	http.ListenAndServe(":8080", router)
}

func (h *HandlerCreateClient) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>CREATED Client</h1>")
}
