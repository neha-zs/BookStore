package main

import (
	"Project/BookStore/Handler"
	book2 "Project/BookStore/Service/book"
	"Project/BookStore/Store/book"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "neha:password@/BookStore")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected!")

	st := book.NewStore(db)
	svc := book2.NewService(st)
	h := Handler.NewHandler(svc)

	r := mux.NewRouter()

	r.HandleFunc("/books", h.HandlerGet).Methods(http.MethodGet)
	r.HandleFunc("/books", h.HandlerCreate).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8000", r))

}
