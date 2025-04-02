package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/neric1/mail/controller"
	"github.com/neric1/mail/helper"
	"github.com/neric1/mail/router"
	"github.com/rs/zerolog/log"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	log.Info().Msg("Started Server!")
	// Database
	godotenv.Load()
	// User Setup
	userController := controller.NewUsersController()

	// Router
	routes := router.NewRouter(

		userController,
	)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
