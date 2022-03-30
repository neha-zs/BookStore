package book

import (
	"Project/BookStore/Model"
	"database/sql"
	"github.com/google/uuid"
)

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) store {
	return store{db}
}

// create book in stroe layer
func (s store) Create(b *Model.Book) (Model.Book, error) {
	b.ID = uuid.New()
	err := s.db.QueryRow("INSERT INTO books (id, name, author , publishedOn , publisher, numberofBookSold) VALUES ($1, $2, $3, $4 , $5 , $6) RETURNING id",
		b.ID, b.Name, b.Author, b.PublishedOn, b.Publisher, b.PublishedOn).Scan(&b.ID)
	if err != nil {
		return *b, err
	}

	return *b, nil
}

// Get book in store layer
func (s store) Get(author, publication string) ([]Model.Book, error) {

	rows, err := s.db.Query("SELECT * FROM books WHERE author = $1 AND publishedOn = $2", author, publication)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []Model.Book{}
	for rows.Next() {
		b := Model.Book{}
		err := rows.Scan(&b.ID, &b.Name, &b.Author, &b.PublishedOn, &b.Publisher, &b.NumberOfBookSold)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}


