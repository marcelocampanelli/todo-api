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
	userHandler := handlers.NewUserHandler(database.NewUserRepositoy(db))
	taskHandler := handlers.NewTaskHandler(database.NewTaskRepository(db))
	server.AddHandler("/users", userHandler.Create)
	server.AddHandler("/users/{userID}", userHandler.Update)
	server.AddHandler("/tasks", taskHandler.Create)
	server.AddHandler("/tasks/{taskID}", taskHandler.Update)
	server.AddHandler("/tasks/{taskID}", taskHandler.Delete)
	fmt.Println("O SERVER SUBIU CARALHOOO")
	server.Start()
}
