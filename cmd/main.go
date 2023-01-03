package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KeruboDee/Go-Rest-Api/pkg/handlers"
	"github.com/gorilla/mux"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my First Go API!")
	fmt.Println("Article REST API")
}
func handleRequests() {
	//create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/articles", handlers.GetAllArticles).Methods(http.MethodGet)
	myRouter.HandleFunc("/articles/{id}", handlers.GetArticle).Methods(http.MethodGet)
	myRouter.HandleFunc("/articles", handlers.AddArticle).Methods(http.MethodPost)
	myRouter.HandleFunc("/articles/{id}", handlers.UpdateArticle).Methods(http.MethodPut)
	myRouter.HandleFunc("/articles/{id}", handlers.DeleteArticle).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
