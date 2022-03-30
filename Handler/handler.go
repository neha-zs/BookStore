package Handler

import (
	"Project/BookStore/Model"
	"Project/BookStore/Service"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Struct
type handler struct {
	service Service.ServiceBook
}

func NewHandler(service Service.ServiceBook) handler {
	return handler{service}
}

func (h handler) HandlerGet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	author := r.URL.Query().Get("author")
	publisher := r.URL.Query().Get("publication")

	resp, err := h.service.ServiceGet(author, publisher)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	body, _ := json.Marshal(resp)

	_, _ = w.Write(body)
}

func (h handler) HandlerCreate(w http.ResponseWriter, r *http.Request) {
	var b *Model.Book

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := h.service.ServiceCreate(b)
	body, err = json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_, _ = w.Write(body)

}
