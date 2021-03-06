// Code generated by MockGen. DO NOT EDIT.
// Source: Store/interface.go

// Package Store is a generated GoMock package.
package Store

import (
	Model "Project/BookStore/Model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStoreBook is a mock of StoreBook interface.
type MockStoreBook struct {
	ctrl     *gomock.Controller
	recorder *MockStoreBookMockRecorder
}

// MockStoreBookMockRecorder is the mock recorder for MockStoreBook.
type MockStoreBookMockRecorder struct {
	mock *MockStoreBook
}

// NewMockStoreBook creates a new mock instance.
func NewMockStoreBook(ctrl *gomock.Controller) *MockStoreBook {
	mock := &MockStoreBook{ctrl: ctrl}
	mock.recorder = &MockStoreBookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreBook) EXPECT() *MockStoreBookMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStoreBook) Create(b *Model.Book) (Model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", b)
	ret0, _ := ret[0].(Model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockStoreBookMockRecorder) Create(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStoreBook)(nil).Create), b)
}

// Get mocks base method.
func (m *MockStoreBook) Get(author, publication string) ([]Model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", author, publication)
	ret0, _ := ret[0].([]Model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStoreBookMockRecorder) Get(author, publication interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStoreBook)(nil).Get), author, publication)
}
