// point of this file is to have functions that directly communicate with the database
// only running database queries, no logic here -> controllers folder

package models

import (
	"example.com/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm" // ORM - alleviates load of writing raw SQL
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json: "publication"`
}

func init() {
	// initialises the database
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// CreateBook creates a new book entry to send to the database
func (book *Book) CreateBook() *Book {
	// creates a new book record
	db.NewRecord(book)
	db.Create(&book)
	// returns the created book
	return book
}

// GetBooks queries the database for all the books stored therein
func GetBooks() []Book {
	// create a slice of Books to temporarily store the queried data
	var Books []Book
	// finds records matching Book in the db and points them to be stored in the created buffer
	db.Find(&Books)
	// returns the buffer containing the books
	return Books
}

// GetBook queries the database for a specific book, defined by its ID
func GetBook(Id int64) (*Book, *gorm.DB) {
	// create a temporary variable to store the queried data
	var searched_book Book
	// finds record matching the Id argument, and points it to be stored in the buffer
	db := db.Where("ID=?", Id).Find(&searched_book)
	// returns the book (buffer), and the db
	return &searched_book, db
}

// DeleteBook deletes a book entry that matches the given Id
func DeleteBook(Id int64) Book {
	// create a buffer variable
	var deleted_book Book
	// find book by ID, delete it
	db.Where("ID=?", Id).Delete(deleted_book)
	// return the deleted book
	return deleted_book
}

// UpdateBook -> get, delete, create pipeline
