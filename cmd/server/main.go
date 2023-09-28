package main

import (
	"fmt"
	"github.com/marcelocampanelli/todo-api/internal/infra/database"
	"github.com/marcelocampanelli/todo-api/internal/infra/handlers"
	"github.com/marcelocampanelli/todo-api/internal/infra/webserver"
)

func main() {

	db, err := database.ConnectToDatabase()
	if err != nil {
		fmt.Println("Error to access database:", err)
		panic(err)
	}
	defer db.Close()

	server := webserver.NewWebServer(":3000")
	userHandler := handlers.NewUserHadler(database.NewUserRepositoy(db))
	server.AddHandler("/users", userHandler.Create)
	fmt.Println("O SERVER SUBIU CARALHOOO")
	server.Start()
}
