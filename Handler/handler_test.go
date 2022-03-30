package Handler

import (
	"Project/BookStore/Model"
	"Project/BookStore/Service"
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := Service.NewMockServiceBook(ctrl)
	s := NewHandler(mockService)

	id := uuid.New()

	b1 := Model.Book{ID: id, Name: "Harry Potter", Author: "JK ROWLING", PublishedOn: "1977", Publisher: "Bloomsbury Publishing (UK)", NumberOfBookSold: 20}

	testCases := []struct {
		desc       string
		book       Model.Book
		statusCode int
		err        error
	}{

		{desc: "success", book: b1, statusCode: http.StatusCreated, err: nil}}

	gomock.InOrder(
		mockService.EXPECT().ServiceCreate(&b1).Return(b1, nil),
	)

	for _, tc := range testCases {
		body, _ := json.Marshal(tc.book)

		req := httptest.NewRequest("POST", "/books/", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		s.HandlerCreate(res, req)

		if res.Code != tc.statusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", tc.statusCode, res.Code)
		}
	}
}

func TestHandlerGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := Service.NewMockServiceBook(ctrl)
	s := NewHandler(mockService)

	b1 := Model.Book{Name: "Harry Potter", Author: "JK ROWLING", PublishedOn: "1977",
		Publisher: "Bloomsbury Publishing (UK)", NumberOfBookSold: 20}
	//b2 := Model.Book{Name: "ABC", Author: "JK", PublishedOn: "1977", Publisher: "Bloomsbury Publishing (UK)", NumberOfBookSold: 200}

	testCases := []struct {
		desc           string
		inputAuthor    string
		inputPublisher string
		//expectedOutput []Model.Book
		statusCode int
		mock       []*gomock.Call
	}{
		{
			desc:           "success case",
			inputAuthor:    "JK ROWLING",
			statusCode:     http.StatusOK,
			inputPublisher: "Bloomsbury Publishing (UK)",
			mock: []*gomock.Call{mockService.EXPECT().ServiceGet("", "").
				Return([]Model.Book{b1}, nil)}},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest("GET", "/books", nil)

		res := httptest.NewRecorder()

		s.HandlerGet(res, req)

		if res.Code != tc.statusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", tc.statusCode, res.Code)
		}
	}
}
