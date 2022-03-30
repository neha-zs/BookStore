package book

import (
	"Project/BookStore/Model"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

// test for create book in datastore layer
func TestCreateBook(t *testing.T) {

	b1 := Model.Book{
		Name:             "Harry Potter",
		Author:           "JK ROWLING",
		PublishedOn:      "1977",
		Publisher:        "Bloomsbury Publishing (UK)",
		NumberOfBookSold: 20,
	}

	testcases := []struct {
		desc       string
		input      Model.Book
		outputBook Model.Book
		err        error
	}{
		{desc: "Success Case", input: b1, outputBook: b1, err: nil},
		{desc: "Nil Body", input: Model.Book{}, outputBook: Model.Book{}, err: nil},
	}
	db, err, _ := sqlmock.New()
	if err != nil {
		return
	}

	st := NewStore(db)

	for i, tc := range testcases {
		book, err := st.Create(&tc.input)

		if err != tc.err {
			t.Errorf("Test case %d: Expected error %v, but got %v", i, tc.err, err)
		}

		if book != tc.outputBook {
			t.Errorf("Test case %d: Expected book %v, but got %v", i, tc.outputBook, book)
		}
	}
}
