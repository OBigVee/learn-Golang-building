package main
import (
	"fmt"
	"net/http"
	"log"
	)

func formHandler(w http.ResponseWriter, req *http.Request){
	if err := req.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name :=  req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, req *http.Request){
	if req.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello O.BigVee again doing the stunt!")
}

func main() {
	var fileserver = http.FileServer(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Sever connected on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err !=nil{
		log.Fatal(err)
	}
}
