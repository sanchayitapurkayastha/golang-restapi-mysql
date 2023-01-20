package controllers

import (
    "fmt"

    "github.com/gorilla/mux"
    "github.com/sanchayitapurkayastha/golang-rest-api-mysql/pkg/utils"
    "github.com/sanchayitapurkayastha/golang-rest-api-mysql/pkg/models"

    "encoding/json"
    "net/http"
    "strconv"
)

var NewBook models.Book

func enableCors(w *http.ResponseWriter) {
   (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    w.Header().Set("Content-Type", "application/json")

    b := models.GetBooks()

    res, _ := json.Marshal(b)
    w.WriteHeader(http.StatusOK)    //it will give us 200=> everything is working fine
    w.Write(res)    //the response we'll be sending to the frontend or postman
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    w.Header().Set("Content-Type", "application/json")

    vars := mux.Vars(r)
    //access the bookId in the request
    bookId := vars["bookId"]

    id, err := strconv.ParseInt(bookId, 0, 0)
    if err !=  nil {
        fmt.Println("Parsing Error")
    }

    bookDetails, _ := models.GetBookById(id)

    res, _ := json.Marshal(bookDetails)
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    w.Header().Set("Content-Type", "application/json")

    CreateBook := &models.Book{}
    utils.ParseBody(r, CreateBook)
    b := CreateBook.CreateBook()

    res, _ := json.Marshal(b)
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    w.Header().Set("Content-Type", "application/json")

    vars := mux.Vars(r)
    bookId := vars["bookId"]

    id, err := strconv.ParseInt(bookId, 0, 0)
    if err !=  nil {
        fmt.Println("Parsing Error")
    }

    bookDetails:= models.DeleteBook(id)

    res, _ := json.Marshal(bookDetails)
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    w.Header().Set("Content-Type", "application/json")

    updatedBook := &models.Book{}
    utils.ParseBody(r, updatedBook)

    vars := mux.Vars(r)
    bookId := vars["bookId"]

    id, err := strconv.ParseInt(bookId, 0, 0)
    if err !=  nil {
        fmt.Println("Parsing Error")
    }

    bookDetails, db := models.GetBookById(id)
    if updatedBook.Name != ""  {
        bookDetails.Name = updatedBook.Name
    }
    if updatedBook.Author != ""  {
        bookDetails.Author = updatedBook.Author
    }
    if updatedBook.Publication != ""  {
        bookDetails.Publication = updatedBook.Publication
    }
    db.Save(&bookDetails)

    res, _ := json.Marshal(bookDetails)
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}