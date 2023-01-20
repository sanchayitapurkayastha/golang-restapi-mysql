package routes

import (
   "log"
   "net/http"
   "github.com/gorilla/mux"
   "github.com/sanchayitapurkayastha/golang-rest-api-mysql/pkg/controllers"
)

func RegisterBookStoreRoutes() {
    r := mux.NewRouter()

    r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
    r.HandleFunc("/book/", controllers.GetBook).Methods("GET")
    r.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
    r.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
    r.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

    http.Handle("/", r)
    log.Fatal(http.ListenAndServe("localhost:9010", r))
}