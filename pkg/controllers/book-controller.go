package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/snake-py/go-bookstore/pkg/models"
	"github.com/snake-py/go-bookstore/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	if id, err := strconv.ParseInt(bookId, 0, 0); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Invalid book id"}`))
	} else {
		bookDetails, _ := models.GetBookById(id)
		res, _ := json.Marshal(bookDetails)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	bookStruct := &models.Book{}
	utils.ParseBody(r, bookStruct) // this will fill the struct with the recieved data
	book := bookStruct.CreateBook()
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	if id, err := strconv.ParseInt(bookId, 0, 0); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotModified)
		w.Write([]byte(`{"message": "Invalid book id"}`))
	} else {
		book := models.DeleteBook(id)
		res, _ := json.Marshal(book)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write(res)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	if id, err := strconv.ParseInt(bookId, 0, 0); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotModified)
		w.Write([]byte(`{"message": "Invalid book id"}`))
	} else {
		bookStruct := &models.Book{}
		utils.ParseBody(r, bookStruct)
		bookDetails, db := models.GetBookById(id)

		// there must be a better way to do this:
		if bookStruct.Name != "" {
			bookDetails.Name = bookStruct.Name
		}

		if bookStruct.Author != "" {
			bookDetails.Author = bookStruct.Author
		}

		if bookStruct.Publication != "" {
			bookDetails.Publication = bookStruct.Publication
		}

		db.Save(&bookDetails)

		res, _ := json.Marshal(bookDetails)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write(res)
	}
}
