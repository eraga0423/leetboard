package router

import (
	"fmt"
	"net/http"
)

func (r *Router) post() {
	r.router.HandleFunc("GET /posts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

}
