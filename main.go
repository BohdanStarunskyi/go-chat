package main

import (
	"chat/app"
	"chat/database"
	"chat/routes"
	"chat/sockets"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = database.InitDb()
	if err != nil {
		panic("Couldn't init database " + err.Error())
	}
	go sockets.HubInstance.Run()
	e := echo.New()
	application := app.NewApp()
	routes.InitRoutes(e, application)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	if err := e.Start(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic("error starting server " + err.Error())
	}
}
