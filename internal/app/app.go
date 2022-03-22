package app

import (
	"github.com/go-chi/chi/v5"
	"module31/internal/controller"
	"module31/internal/repository"
	"module31/internal/usecase"
	"net/http"
)

func Run(port string) error {
	//repositoryA, err := repository.NewMemorydb()
	repositoryA, err := repository.NewMongodb()
	if err != nil {
		return err
	}
	useCase := usecase.NewUsecase(repositoryA)
	router := chi.NewRouter()
	controller.Build(router, useCase)

	err = http.ListenAndServe(port, router)
	if err != nil {
		return err
	}
	return nil
}
