package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/michee-04/dbtest/pkg/models"
	"github.com/michee-04/dbtest/pkg/utils"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	BookCreate := &models.Book{}
	utils.ParseBody(r, BookCreate)
	b := BookCreate.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	NewBook := models.GetAllBooks()
	res, _ := json.Marshal(NewBook)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetBookById(Id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookUpdate := &models.Book{}
	utils.ParseBody(r, bookUpdate)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails, db:=models.GetBookById(Id)
	if bookUpdate.Name != ""{
		bookDetails.Name = bookUpdate.Name
	}
	if bookUpdate.Author != ""{
		bookDetails.Author = bookUpdate.Author
	}
	if bookUpdate.Publication != ""{
		bookDetails.Publication = bookUpdate.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil{
		fmt.Println("Error while parsing")
	}

	book := models.DeleteBookId(Id)
	res, _ := json.Marshal(book)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
