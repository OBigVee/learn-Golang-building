package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	//"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/obigvee/Go-BookStore/pkg/routes"
)

func main(){
	route := mux.NewRouter()
	routes.RegisterBookStoreRoute(route)
	http.Handle("/", route)
	log.Fatal(http.ListenAndServe("localhost:9010", route))
	fmt.Printf("server running on PORT:9010")
}
