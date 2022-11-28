package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"example.com/go-bookstore/pkg/models"
	"example.com/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(writer http.ResponseWriter, request *http.Request) {
	// obtains all books in the db
	queried_books := models.GetBooks()

	// converts the obtained data from golang struct to json
	json_res, _ := json.Marshal(queried_books)

	writer.Header().Set("Content-Type", "pkglication/json")

	// sends 200 OK http status code in the http payload header
	writer.WriteHeader(http.StatusOK)

	// 'writes' the response, i.e. sends the response (http body containing book data) to the client
	writer.Write(json_res)
}

func GetBook(writer http.ResponseWriter, request *http.Request) {
	// since book_id will be sent as a variable in the http request, need to obtain it
	args := mux.Vars(request)
	book_id := args["book_id"] // book_id was set in the definition of the API endpoints in routes file

	// since ID comes in as a string, converting it to int
	ID, err := strconv.ParseInt(book_id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing,", err)
	}

	// saving the queried info in a temporary variable
	book_details, _ := models.GetBook(ID)

	// converts the obtained data from golang struct to json for sending
	json_res, _ := json.Marshal(book_details)

	writer.Header().Set("Content-Type", "pkglication/json")

	// sends 200 OK http status code in the http payload header
	writer.WriteHeader(http.StatusOK)

	// 'writes' the response, i.e. sends the response (http body containing book data) to the client
	writer.Write(json_res)
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	created_book := &models.Book{}                // pointer to an identifier of type Book
	utils.ParseRequestBody(request, created_book) // see utils

	new_book := created_book.CreateBook() // see models

	json_res, _ := json.Marshal(new_book)

	writer.Header().Set("Content-Type", "pkglication/json")

	// sends 200 OK http status code in the http payload header
	writer.WriteHeader(http.StatusOK)

	// 'writes' the response, i.e. sends the response (http body containing book data) to the client
	writer.Write(json_res)
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	// since book_id will be sent as a variable in the http request, need to obtain it
	args := mux.Vars(request)
	book_id := args["book_id"] // book_id was set in the definition of the API endpoints in routes file

	// since ID comes in as a string, converting it to int
	ID, err := strconv.ParseInt(book_id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing,", err)
	}

	// saving the queried info in a temporary variable
	book_details := models.DeleteBook(ID)

	json_res, _ := json.Marshal(book_details)

	writer.Header().Set("Content-Type", "pkglication/json")

	// sends 200 OK http status code in the http payload header
	writer.WriteHeader(http.StatusOK)

	// 'writes' the response, i.e. sends the response (http body containing book data) to the client
	writer.Write(json_res)
}

// UpdateBook updates a record in golang
func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	updated_book := &models.Book{}
	utils.ParseRequestBody(request, updated_book)

	// since book_id will be sent as a variable in the http request, need to obtain it
	args := mux.Vars(request)
	book_id := args["book_id"] // book_id was set in the definition of the API endpoints in routes file

	// since ID comes in as a string, converting it to int
	ID, err := strconv.ParseInt(book_id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing,", err)
	}

	// checking the passed updates in the http request, updating them if they have been changed
	book_details, db := models.GetBook(ID)
	if updated_book.Name != "" {
		book_details.Name = updated_book.Name
	}
	if updated_book.Author != "" {
		book_details.Name = updated_book.Author
	}
	if updated_book.Publication != "" {
		book_details.Publication = updated_book.Publication
	}

	// to save the changes to the db
	db.Save(&book_details)

	// to send data back to the client
	json_res, _ := json.Marshal(book_details)

	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(json_res)
}
