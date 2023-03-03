package routes

import(
	"github.com/gorilla/mux"
	"github.com/obigvee/Go-BookStore/pkg/controllers"
)

var RegisterBookStoreRoute = func(router *mux.Router){
	router.HandlerFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandlerFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandlerFunc("/book/{boomId}", controllers.GetBook).Methods("GET")
	router.HandlerFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandlerFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
