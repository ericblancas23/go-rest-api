package main

import (
	"fmt"
	"net/http"

	"github.com/ericblancas23/go-rest-api/internal/database"
	transportHTTP "github.com/ericblancas23/go-rest-api/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {
	fmt.Println("setting app")

	var err error

	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to load server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go lang rest")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error occured")
		fmt.Println(err)
	}
}
