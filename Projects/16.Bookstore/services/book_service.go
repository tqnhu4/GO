package services

import (
	"bookstore/models"
	"bookstore/repository"
	"errors"
)

// BookService defines the interface for book-related business logic
type BookService interface {
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id uint) (models.Book, error)
	CreateBook(book *models.Book) error
	UpdateBook(id uint, updatedBook *models.Book) error
	DeleteBook(id uint) error
}

// bookServiceImpl implements BookService
type bookServiceImpl struct {
	bookRepo repository.BookRepository
}

// NewBookService creates a new BookService instance
func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookServiceImpl{bookRepo: bookRepo}
}

func (s *bookServiceImpl) GetAllBooks() ([]models.Book, error) {
	return s.bookRepo.GetAllBooks()
}

func (s *bookServiceImpl) GetBookByID(id uint) (models.Book, error) {
	book, err := s.bookRepo.GetBookByID(id)
	if err != nil {
		if errors.Is(err, errors.New("record not found")) { // GORM's ErrRecordNotFound
			return models.Book{}, errors.New("book not found")
		}
		return models.Book{}, err
	}
	return book, nil
}

func (s *bookServiceImpl) CreateBook(book *models.Book) error {
	// Add business logic here, e.g., ISBN validation, price checks
	existingBook, err := s.bookRepo.GetBookByISBN(book.ISBN)
	if err == nil && existingBook.ID != 0 {
		return errors.New("book with this ISBN already exists")
	}
	return s.bookRepo.CreateBook(book)
}

func (s *bookServiceImpl) UpdateBook(id uint, updatedBook *models.Book) error {
	book, err := s.bookRepo.GetBookByID(id)
	if err != nil {
		if errors.Is(err, errors.New("record not found")) {
			return errors.New("book not found for update")
		}
		return err
	}

	// Update fields
	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.PublishedDate = updatedBook.PublishedDate
	book.Price = updatedBook.Price
	book.Stock = updatedBook.Stock
	// ISBN should ideally not be updated

	return s.bookRepo.UpdateBook(&book)
}

func (s *bookServiceImpl) DeleteBook(id uint) error {
	// Check if book exists before deleting
	_, err := s.bookRepo.GetBookByID(id)
	if err != nil {
		if errors.Is(err, errors.New("record not found")) {
			return errors.New("book not found for deletion")
		}
		return err
	}
	return s.bookRepo.DeleteBook(id)
}