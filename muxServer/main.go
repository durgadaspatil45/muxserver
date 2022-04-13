package main

import (
	"log"
	h "muxServer/Handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	/*The router is the main router for your web application and will later be passed as
	  parameter to the server. It will receive all HTTP connections and pass it on to the
	  request handlers you will register on it*/
	r := mux.NewRouter()
	r.HandleFunc("/book/add", h.CreateBook).Methods("POST")
	r.HandleFunc("/book/all", h.ReadBook).Methods("GET")
	r.HandleFunc("/book/delete/{Id}", h.DeleteBook).Methods("DELETE")
	r.HandleFunc("/book/update", h.UpdateBook).Methods("PUT")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8085", r))
	//The biggest strength of the gorilla/mux Router is the ability to extract segments from the request URL.

}
