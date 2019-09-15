// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/query/storage/m3 (interfaces: Storage)

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package m3 is a generated GoMock package.
package m3

import (
	"context"
	"reflect"

	"github.com/m3db/m3/src/query/block"
	"github.com/m3db/m3/src/query/storage"

	"github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockStorage) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStorageMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorage)(nil).Close))
}

// CompleteTags mocks base method
func (m *MockStorage) CompleteTags(arg0 context.Context, arg1 *storage.CompleteTagsQuery, arg2 *storage.FetchOptions) (*storage.CompleteTagsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteTags", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.CompleteTagsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompleteTags indicates an expected call of CompleteTags
func (mr *MockStorageMockRecorder) CompleteTags(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteTags", reflect.TypeOf((*MockStorage)(nil).CompleteTags), arg0, arg1, arg2)
}

// ErrorBehavior mocks base method
func (m *MockStorage) ErrorBehavior() storage.ErrorBehavior {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ErrorBehavior")
	ret0, _ := ret[0].(storage.ErrorBehavior)
	return ret0
}

// ErrorBehavior indicates an expected call of ErrorBehavior
func (mr *MockStorageMockRecorder) ErrorBehavior() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrorBehavior", reflect.TypeOf((*MockStorage)(nil).ErrorBehavior))
}

// Fetch mocks base method
func (m *MockStorage) Fetch(arg0 context.Context, arg1 *storage.FetchQuery, arg2 *storage.FetchOptions) (*storage.FetchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.FetchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockStorageMockRecorder) Fetch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockStorage)(nil).Fetch), arg0, arg1, arg2)
}

// FetchBlocks mocks base method
func (m *MockStorage) FetchBlocks(arg0 context.Context, arg1 *storage.FetchQuery, arg2 *storage.FetchOptions) (block.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBlocks", arg0, arg1, arg2)
	ret0, _ := ret[0].(block.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBlocks indicates an expected call of FetchBlocks
func (mr *MockStorageMockRecorder) FetchBlocks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBlocks", reflect.TypeOf((*MockStorage)(nil).FetchBlocks), arg0, arg1, arg2)
}

// FetchCompressed mocks base method
func (m *MockStorage) FetchCompressed(arg0 context.Context, arg1 *storage.FetchQuery, arg2 *storage.FetchOptions) (SeriesFetchResult, Cleanup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCompressed", arg0, arg1, arg2)
	ret0, _ := ret[0].(SeriesFetchResult)
	ret1, _ := ret[1].(Cleanup)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FetchCompressed indicates an expected call of FetchCompressed
func (mr *MockStorageMockRecorder) FetchCompressed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCompressed", reflect.TypeOf((*MockStorage)(nil).FetchCompressed), arg0, arg1, arg2)
}

// Name mocks base method
func (m *MockStorage) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockStorageMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockStorage)(nil).Name))
}

// SearchCompressed mocks base method
func (m *MockStorage) SearchCompressed(arg0 context.Context, arg1 *storage.FetchQuery, arg2 *storage.FetchOptions) (TagResult, Cleanup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchCompressed", arg0, arg1, arg2)
	ret0, _ := ret[0].(TagResult)
	ret1, _ := ret[1].(Cleanup)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SearchCompressed indicates an expected call of SearchCompressed
func (mr *MockStorageMockRecorder) SearchCompressed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCompressed", reflect.TypeOf((*MockStorage)(nil).SearchCompressed), arg0, arg1, arg2)
}

// SearchSeries mocks base method
func (m *MockStorage) SearchSeries(arg0 context.Context, arg1 *storage.FetchQuery, arg2 *storage.FetchOptions) (*storage.SearchResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchSeries", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.SearchResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchSeries indicates an expected call of SearchSeries
func (mr *MockStorageMockRecorder) SearchSeries(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchSeries", reflect.TypeOf((*MockStorage)(nil).SearchSeries), arg0, arg1, arg2)
}

// Type mocks base method
func (m *MockStorage) Type() storage.Type {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(storage.Type)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockStorageMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockStorage)(nil).Type))
}

// Write mocks base method
func (m *MockStorage) Write(arg0 context.Context, arg1 *storage.WriteQuery) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockStorageMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockStorage)(nil).Write), arg0, arg1)
}
