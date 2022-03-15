package app

import (
	"github.com/go-chi/chi/v5"
	"module31/internal/controller"
	"module31/internal/repository"
	"module31/internal/usecase"
	"net/http"
)

func Run(args []string) error {
	repository, err := repository.NewMemorydb()
	//repository, err := repository.NewMongodb()
	if err != nil {
		return err
	}
	usecase := usecase.NewUsecase(repository)
	router := chi.NewRouter()
	controller.Build(router, usecase)

	http.ListenAndServe(":8080", router)
	return nil
}
