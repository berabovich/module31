package app

import (
	"github.com/go-chi/chi/v5"
	"module31/internal/controller"
	"module31/internal/repository"
	"module31/internal/usecase"
	"net/http"
)

//var repositoryA, err = repository.NewMemorydb()
//var useCase = usecase.NewUsecase(repositoryA)

func Run(args []string) error {
	repositoryA, err := repository.NewMemorydb()
	//repository, err := repository.NewMongodb()
	if err != nil {
		return err
	}
	useCase := usecase.NewUsecase(repositoryA)
	router := chi.NewRouter()
	controller.Build(router, useCase)

	err = http.ListenAndServe("localhost:8080", router)
	if err != nil {
		return err
	}
	return nil
}

func Run2(args []string) error {
	repositoryA, err := repository.NewMemorydb()
	//repository, err := repository.NewMongodb()
	if err != nil {
		return err
	}
	useCase := usecase.NewUsecase(repositoryA)
	router := chi.NewRouter()
	controller.Build(router, useCase)

	err = http.ListenAndServe(":8081", router)
	if err != nil {
		return err
	}
	return nil
}
