package Store

import "Project/BookStore/Model"

type StoreBook interface {
	Create(b *Model.Book) (Model.Book, error)
	Get(author, publication string) ([]Model.Book, error)
}
