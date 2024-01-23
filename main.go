package main

import (
	"net/http"

	"github.com/andreaswiidi/restful-api/app"
	"github.com/andreaswiidi/restful-api/controller"
	"github.com/andreaswiidi/restful-api/helper"
	"github.com/andreaswiidi/restful-api/middleware"
	"github.com/andreaswiidi/restful-api/repository"
	"github.com/andreaswiidi/restful-api/service"
	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryControler(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
