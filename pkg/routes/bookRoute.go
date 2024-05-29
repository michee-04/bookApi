package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michee-04/dbtest/pkg/controller"
)

const port = ":5000"

// var RegisterBookStoreRoutes = func (router *mux.Router) {
 func RegisterBookStoreRoutes(){

	r := mux.NewRouter()

	r.HandleFunc("/book/", controller.CreateBook).Methods("POST")

	r.HandleFunc("/book/", controller.GetBook).Methods("GET")

	r.HandleFunc("/book/{bookId}", controller.GetBookById).Methods("GET")

	r.HandleFunc("/book/{bookId}", controller.UpdateBook).Methods("PUT")

	r.HandleFunc("/book/{bookId}", controller.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(port, r))
}