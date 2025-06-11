package models

import (
	"time"
)

// Book represents a book in the bookstore
type Book struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Title        string    `gorm:"not null" json:"title"`
	Author       string    `gorm:"not null" json:"author"`
	ISBN         string    `gorm:"unique;not null" json:"isbn"`
	PublishedDate time.Time `json:"published_date"`
	Price        float64   `json:"price"`
	Stock        int       `json:"stock"`
}