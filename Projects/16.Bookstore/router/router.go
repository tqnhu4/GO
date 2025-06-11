package router

import (
	"bookstore/config"
	"bookstore/handlers"
	"bookstore/repository"
	"bookstore/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRouter initializes the API routes
func InitRouter() *mux.Router {
	r := mux.NewRouter()

	// Initialize dependencies
	bookRepo := repository.NewBookRepository(config.DB)
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	// Book API routes
	r.HandleFunc("/books", bookHandler.GetAllBooks).Methods("GET")
	r.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", bookHandler.GetBookByID).Methods("GET")
	r.HandleFunc("/books/{id}", bookHandler.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", bookHandler.DeleteBook).Methods("DELETE")

	// Middleware (optional, but good practice)
	r.Use(loggingMiddleware)

	return r
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}