package categories

import (
	"encoding/json"
	"net/http"

	"sample-go-app/internal/models"
	"sample-go-app/internal/database"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	db := database.DB
	if db == nil {
		http.Error(w, "database connection is not established", http.StatusInternalServerError)
		return
	}

	if result := db.Create(&newCategory); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategory)
}