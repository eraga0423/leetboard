package main

import (
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func main() {
	// 	// Подключение к базе данных PostgreSQL
	// 	db, err := sql.Open("postgres", "user=youruser dbname=1337b04rd sslmode=disable")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer db.Close()

	// 	// Инициализация слоев
	// 	repo := adapters.NewPostgresRepository(db)
	// 	service := core.NewPostService(repo)
	// 	handler := adapters.NewPostHandler(service)

	// 	http.HandleFunc("/post", handler.CreatePostHandler)

	// 	log.Println("Server running on port 8080")
	// 	http.ListenAndServe(":8080", nil)
}

var templates = map[string]*template.Template{}

func init() {
	LoadTemplates()
}

const (
	archivePage = "../frontend/archive.html"
	catalogPage = "../frontend/catalog.html"
)

func LoadTemplates() {
	pages := []string{archivePage, catalogPage}
	for _, page := range pages {
		tmpl, err := template.ParseFiles(page)
		if err != nil {
			log.Fatalf("loading error %s : %v", page, err)
		}
		templates[page] = tmpl
	}
}

func CatalogHandler(w http.ResponseWriter, r *http.Request) {
	templates[catalogPage].Execute(w, nil)
}
