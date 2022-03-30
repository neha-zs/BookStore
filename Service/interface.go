package Service

import "Project/BookStore/Model"

type ServiceBook interface {
	ServiceCreate(b *Model.Book) (Model.Book, error)
	ServiceGet(author, publisher string) ([]Model.Book, error)
}

