package categories

import (
    "net/http"
    "strconv"
    "sample-go-app/internal/database"
    "sample-go-app/internal/models"
	"log"
	"github.com/go-chi/chi/v5"
)

// DeleteCategory handles DELETE requests to delete a category
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
    // Extract the category ID from the URL
	categoryIDStr := chi.URLParam(r, "id")

    // Convert categoryID from string to int
    categoryID, err := strconv.Atoi(categoryIDStr)
    log.Printf("Attempting to delete category with ID: %d, original string was '%s'", categoryID, categoryIDStr)
    if err != nil {
        http.Error(w, "Invalid category ID: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Connect to the database
    db := database.DB
    if db == nil {
        http.Error(w, "database connection is not established", http.StatusInternalServerError)
        return
    }

    // Delete the category from the database
    result := db.Delete(&models.Category{}, categoryID)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    // Check if the category was actually deleted (result.RowsAffected should be 1)
    if result.RowsAffected == 0 {
        http.Error(w, "Category not found", http.StatusNotFound)
        return
    }

    // Respond with success
    w.WriteHeader(http.StatusOK) // 200 OK
    w.Write([]byte("Category deleted successfully"))
}