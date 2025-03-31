package main

import (
	"context"
	"log"
	"log/slog"

	"1337b0rd/internal/config"
	"1337b0rd/internal/governor"
	"1337b0rd/internal/posgres"
	"1337b0rd/internal/rest"

	_ "github.com/lib/pq"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	gov := governor.New()

	r := rest.New(gov)
	conf := config.NewConfig()
	_, err := posgres.NewPosgres(&conf.Postgres)
	if err != nil {
		slog.Any("failed start database", "postgres")
		panic(err)
	}
	
	go func(ctx context.Context, cancelFunc context.CancelFunc) {
		err := r.Start(ctx)
		if err != nil {
			log.Fatal(err)
		}
		cancelFunc()
	}(ctx, cancel)

	//http.HandleFunc("GET /catalog", CatalogHandler)
	//fmt.Println("start server")
	//err := http.ListenAndServe(":9090", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
}

// var templates = map[string]*template.Template{}

// func init() {
// 	LoadTemplates()
// }

// func LoadTemplates() {
// 	pages := []string{archivePage, catalogPage}
// 	for _, page := range pages {
// 		tmpl, err := template.ParseFiles(page)
// 		if err != nil {
// 			log.Fatalf("loading error %s : %v", page, err)
// 		}
// 		templates[page] = tmpl
// 	}
// }

// func CatalogHandler(w http.ResponseWriter, r *http.Request) {
// 	templates[catalogPage].Execute(w, nil)
// }
