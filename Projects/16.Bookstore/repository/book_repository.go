package repository

import (
	"bookstore/models"
	"gorm.io/gorm"
)

// BookRepository defines the interface for book data operations
type BookRepository interface {
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id uint) (models.Book, error)
	GetBookByISBN(isbn string) (models.Book, error)
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(id uint) error
}

// bookRepositoryImpl implements BookRepository using GORM
type bookRepositoryImpl struct {
	db *gorm.DB
}

// NewBookRepository creates a new BookRepository instance
func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepositoryImpl{db: db}
}

func (r *bookRepositoryImpl) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	result := r.db.Find(&books)
	return books, result.Error
}

func (r *bookRepositoryImpl) GetBookByID(id uint) (models.Book, error) {
	var book models.Book
	result := r.db.First(&book, id)
	return book, result.Error
}

func (r *bookRepositoryImpl) GetBookByISBN(isbn string) (models.Book, error) {
	var book models.Book
	result := r.db.Where("isbn = ?", isbn).First(&book)
	return book, result.Error
}

func (r *bookRepositoryImpl) CreateBook(book *models.Book) error {
	result := r.db.Create(book)
	return result.Error
}

func (r *bookRepositoryImpl) UpdateBook(book *models.Book) error {
	result := r.db.Save(book) // Save updates all fields
	return result.Error
}

func (r *bookRepositoryImpl) DeleteBook(id uint) error {
	result := r.db.Delete(&models.Book{}, id)
	return result.Error
}