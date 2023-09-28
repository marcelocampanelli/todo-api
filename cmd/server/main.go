package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/marcelocampanelli/todo-api/internal/infra/database"
	"net/http"
)

func main() {

	db, err := database.ConnectToDatabase()
	if err != nil {
		fmt.Println("Error to access database:", err)
		panic(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":3000", r)
}
