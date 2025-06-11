package handlers

import (
	"bookstore/models"
	"bookstore/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// BookHandler handles HTTP requests related to books
type BookHandler struct {
	bookService services.BookService
}

// NewBookHandler creates a new BookHandler instance
func NewBookHandler(s services.BookService) *BookHandler {
	return &BookHandler{bookService: s}
}

// GetAllBooks handles GET /books
func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.bookService.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBookByID handles GET /books/{id}
func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := h.bookService.GetBookByID(uint(id))
	if err != nil {
		if err.Error() == "book not found" { // Specific error from service layer
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// CreateBook handles POST /books
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if book.Title == "" || book.Author == "" || book.ISBN == "" {
		http.Error(w, "Title, Author, and ISBN are required", http.StatusBadRequest)
		return
	}

	err = h.bookService.CreateBook(&book)
	if err != nil {
		if err.Error() == "book with this ISBN already exists" {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// UpdateBook handles PUT /books/{id}
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var updatedBook models.Book
	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.bookService.UpdateBook(uint(id), &updatedBook)
	if err != nil {
		if err.Error() == "book not found for update" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook) // Return the updated book
}

// DeleteBook handles DELETE /books/{id}
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	err = h.bookService.DeleteBook(uint(id))
	if err != nil {
		if err.Error() == "book not found for deletion" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content for successful deletion
}