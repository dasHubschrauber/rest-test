package main

import (
	"github.com/dasHubschrauber/rest-test/pkg/db"
	"github.com/dasHubschrauber/rest-test/pkg/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/addBook", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/book/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/book/{id}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)

}
