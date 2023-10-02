package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dasHubschrauber/rest-test/pkg/models"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	var book models.Book

	if result := h.DB.Find(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.NrPages = updatedBook.NrPages
	book.Desc = updatedBook.Desc

	h.DB.Save(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated " + book.Title)

}
