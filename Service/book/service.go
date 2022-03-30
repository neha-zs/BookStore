package book

import (
	"Project/BookStore/Model"
	"Project/BookStore/Store"
)

type service struct {
	store Store.StoreBook
}

// NewService is a factory function
func NewService(store Store.StoreBook) service {
	return service{store: store}
}

func (s service) ServiceCreate(b *Model.Book) (Model.Book, error) {
	if resp, err := s.store.Create(b); err != nil {
		return resp, err
	}
	return *b, nil
}

func (s service) ServiceGet(author, publisher string) ([]Model.Book, error) {
	if resp, err := s.store.Get(author, publisher); err != nil {
		return resp, err
	}
	return []Model.Book{}, nil
}
