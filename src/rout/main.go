package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/gorilla/handlers"
	"os"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/products", ProductHandler).Methods("GET")
	r.HandleFunc("/products/{key}", ProductHandler).Methods("GET")
	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler).Methods("GET")
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	//http.Handle("/", r)
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["key"])
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Key: %v\n", vars["key"])
}

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The Home Page.")
}

