// Code generated by MockGen. DO NOT EDIT.
// Source: Service/interface.go

// Package Service is a generated GoMock package.
package Service

import (
	Model "Project/BookStore/Model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockServiceBook is a mock of ServiceBook interface.
type MockServiceBook struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBookMockRecorder
}

// MockServiceBookMockRecorder is the mock recorder for MockServiceBook.
type MockServiceBookMockRecorder struct {
	mock *MockServiceBook
}

// NewMockServiceBook creates a new mock instance.
func NewMockServiceBook(ctrl *gomock.Controller) *MockServiceBook {
	mock := &MockServiceBook{ctrl: ctrl}
	mock.recorder = &MockServiceBookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBook) EXPECT() *MockServiceBookMockRecorder {
	return m.recorder
}

// ServiceCreate mocks base method.
func (m *MockServiceBook) ServiceCreate(b *Model.Book) (Model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceCreate", b)
	ret0, _ := ret[0].(Model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServiceCreate indicates an expected call of ServiceCreate.
func (mr *MockServiceBookMockRecorder) ServiceCreate(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceCreate", reflect.TypeOf((*MockServiceBook)(nil).ServiceCreate), b)
}

// ServiceGet mocks base method.
func (m *MockServiceBook) ServiceGet(author, publisher string) ([]Model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceGet", author, publisher)
	ret0, _ := ret[0].([]Model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServiceGet indicates an expected call of ServiceGet.
func (mr *MockServiceBookMockRecorder) ServiceGet(author, publisher interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceGet", reflect.TypeOf((*MockServiceBook)(nil).ServiceGet), author, publisher)
}
