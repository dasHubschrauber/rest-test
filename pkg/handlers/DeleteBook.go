package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dasHubschrauber/rest-test/pkg/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
