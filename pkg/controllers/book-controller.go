package controllers

import (
	"Book-Management/pkg/models"
	"Book-Management/pkg/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	// w.Header().Set("Content type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func DeleteBook(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	ID, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	book := models.DeleteBook(ID)
// 	res, _ := json.Marshal(book)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, exists := vars["id"]

	log.Println("Extracted bookId:", bookId)
	log.Println("Vars from request:", vars)

	//  Ensure bookId exists before parsing
	if !exists || bookId == "" {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	//  Convert bookId safely
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "Error parsing book ID", http.StatusBadRequest)
		return
	}

	//  Delete the book and check for errors
	err = models.DeleteBook(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//  Send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Book deleted successfully"}`))
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var UpdateBook = &models.Book{}
	utils.ParseBody(r, UpdateBook)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		fmt.Println("Error while pasrsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if UpdateBook.Name != "" {
		bookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		bookDetails.Publication = UpdateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content tyepe", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// to be removed
// func GetCities(w http.ResponseWriter, r *http.Request) {
// 	ListOFCities := models.GetCities()
// 	res, _ := json.Marshal(ListOFCities)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }
