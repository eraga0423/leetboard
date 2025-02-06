package main

import (
	"1337b04rd/internal/adapters"
	"1337b04rd/internal/core"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	// Подключение к базе данных PostgreSQL
	db, err := sql.Open("postgres", "user=youruser dbname=1337b04rd sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Инициализация слоев
	repo := adapters.NewPostgresRepository(db)
	service := core.NewPostService(repo)
	handler := adapters.NewPostHandler(service)

	http.HandleFunc("/post", handler.CreatePostHandler)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
