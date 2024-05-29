package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"

	"file-system/pkg/connection"
	"file-system/pkg/files"
	"file-system/pkg/helpers"
	"file-system/pkg/routes"
	"file-system/pkg/users"
)

func main() {
	log.Info().Msg("Started Server!")

	db := connection.DatabaseConnection()
	validate := validator.New()

	userRepository := users.NewUserRepositoryImpl(db.User)

	userService := users.NewUserServiceImpl(userRepository, validate)

	userController := users.NewUserController(userService)

	fileRepository := files.NewFileRepositoryImpl(db.File)

	fileService := files.NewFileServiceImpl(fileRepository, validate)

	fileController := files.NewFileController(fileService)

	routes := routes.NewRouter(userController, fileController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helpers.ErrorHelper(err)
}
