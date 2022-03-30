package book

import (
	"Project/BookStore/Model"
	"Project/BookStore/Store"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

// test function for Get book
func TestGet(t *testing.T) {
	b1 := Model.Book{Name: "Harry Potter", Author: "JK ROWLING", PublishedOn: "1977", Publisher: "Bloomsbury Publishing (UK)", NumberOfBookSold: 20}
	b2 := Model.Book{Name: "ABC", Author: "JK", PublishedOn: "1977", Publisher: "Bloomsbury Publishing (UK)", NumberOfBookSold: 200}

	testCases := []struct {
		desc           string
		inputAuthor    string
		inputPublisher string
		expectedOutput []Model.Book
		err            error
		mock           gomock.Call
	}{
		{
			desc:           "Sucess Case",
			inputAuthor:    "JK ROWLING",
			inputPublisher: "",
			expectedOutput: []Model.Book{b1},
			err:            nil,
		},
		{
			desc:           "Author and publisher blank",
			inputAuthor:    "",
			inputPublisher: "",
			expectedOutput: []Model.Book{b1, b2},
			err:            nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStore := Store.NewMockStoreBook(ctrl)
	s := NewService(mockStore)

	//mockStore.EXPECT().Get("JK ROWLING", "").Return([]Model.Book{b1}, nil)
	mockStore.EXPECT().Get("", "").Return([]Model.Book{b1}, nil)

	for i, v := range testCases {

		actualOutput, err := s.ServiceGet(v.inputAuthor, v.inputPublisher)
		if err != v.err {
			t.Errorf("Test case %d: %s failed, expected error: %v, actual error: %v", i, v.desc, v.err, err)
		}

		if reflect.DeepEqual(actualOutput, v.expectedOutput) {
			t.Errorf("Test case %d: %s failed, expected output: %v, actual output: %v", i, v.desc, v.expectedOutput, actualOutput)
		}
	}
}

// test function for create book
func TestCreate(t *testing.T) {
	b1 := Model.Book{Name: "Harry Potter", Author: "JK ROWLING", PublishedOn: "1977", Publisher: "Bloomsbury Publishing (UK)", NumberOfBookSold: 20}

	testCases := []struct {
		desc           string
		inputBook      Model.Book
		expectedOutput Model.Book
		err            error
	}{
		{
			desc:           "Success Case",
			inputBook:      b1,
			expectedOutput: b1,
			err:            nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStore := Store.NewMockStoreBook(ctrl)
	s := NewService(mockStore)

	mockStore.EXPECT().Create(&b1).Return(b1, nil)

	for i, tc := range testCases {

		resp, err := s.ServiceCreate(&tc.inputBook)

		if err != tc.err {
			t.Errorf("Test case %d: Expected error %v, but got %v", i, tc.err, err)
		}

		if resp != tc.expectedOutput {
			t.Errorf("Test case %d: Expected book %v, but got %v", i, tc.expectedOutput, resp)
		}
	}

}
