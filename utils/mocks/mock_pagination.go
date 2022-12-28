/*
 * Copyright (C) 2020-2022 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ARM-software/golang-utils/utils/collection/pagination (interfaces: IPage,IStream,IIterator,IPaginator,IStreamPaginator)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	pagination "github.com/ARM-software/golang-utils/utils/collection/pagination"
)

// MockIStaticPage is a mock of IStaticPage interface.
type MockIStaticPage struct {
	ctrl     *gomock.Controller
	recorder *MockIStaticPageMockRecorder
}

// MockIStaticPageMockRecorder is the mock recorder for MockIStaticPage.
type MockIStaticPageMockRecorder struct {
	mock *MockIStaticPage
}

// NewMockIStaticPage creates a new mock instance.
func NewMockIStaticPage(ctrl *gomock.Controller) *MockIStaticPage {
	mock := &MockIStaticPage{ctrl: ctrl}
	mock.recorder = &MockIStaticPageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStaticPage) EXPECT() *MockIStaticPageMockRecorder {
	return m.recorder
}

// GetItemCount mocks base method.
func (m *MockIStaticPage) GetItemCount() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemCount indicates an expected call of GetItemCount.
func (mr *MockIStaticPageMockRecorder) GetItemCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemCount", reflect.TypeOf((*MockIStaticPage)(nil).GetItemCount))
}

// GetItemIterator mocks base method.
func (m *MockIStaticPage) GetItemIterator() (pagination.IIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemIterator")
	ret0, _ := ret[0].(pagination.IIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemIterator indicates an expected call of GetItemIterator.
func (mr *MockIStaticPageMockRecorder) GetItemIterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemIterator", reflect.TypeOf((*MockIStaticPage)(nil).GetItemIterator))
}

// HasNext mocks base method.
func (m *MockIStaticPage) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockIStaticPageMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockIStaticPage)(nil).HasNext))
}

// MockIPage is a mock of IPage interface.
type MockIPage struct {
	ctrl     *gomock.Controller
	recorder *MockIPageMockRecorder
}

// MockIPageMockRecorder is the mock recorder for MockIPage.
type MockIPageMockRecorder struct {
	mock *MockIPage
}

// NewMockIPage creates a new mock instance.
func NewMockIPage(ctrl *gomock.Controller) *MockIPage {
	mock := &MockIPage{ctrl: ctrl}
	mock.recorder = &MockIPageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPage) EXPECT() *MockIPageMockRecorder {
	return m.recorder
}

// GetItemCount mocks base method.
func (m *MockIPage) GetItemCount() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemCount indicates an expected call of GetItemCount.
func (mr *MockIPageMockRecorder) GetItemCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemCount", reflect.TypeOf((*MockIPage)(nil).GetItemCount))
}

// GetItemIterator mocks base method.
func (m *MockIPage) GetItemIterator() (pagination.IIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemIterator")
	ret0, _ := ret[0].(pagination.IIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemIterator indicates an expected call of GetItemIterator.
func (mr *MockIPageMockRecorder) GetItemIterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemIterator", reflect.TypeOf((*MockIPage)(nil).GetItemIterator))
}

// GetNext mocks base method.
func (m *MockIPage) GetNext(arg0 context.Context) (pagination.IPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNext", arg0)
	ret0, _ := ret[0].(pagination.IPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNext indicates an expected call of GetNext.
func (mr *MockIPageMockRecorder) GetNext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNext", reflect.TypeOf((*MockIPage)(nil).GetNext), arg0)
}

// HasNext mocks base method.
func (m *MockIPage) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockIPageMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockIPage)(nil).HasNext))
}

// MockIStaticPageStream is a mock of IStaticPageStream interface.
type MockIStaticPageStream struct {
	ctrl     *gomock.Controller
	recorder *MockIStaticPageStreamMockRecorder
}

// MockIStaticPageStreamMockRecorder is the mock recorder for MockIStaticPageStream.
type MockIStaticPageStreamMockRecorder struct {
	mock *MockIStaticPageStream
}

// NewMockIStaticPageStream creates a new mock instance.
func NewMockIStaticPageStream(ctrl *gomock.Controller) *MockIStaticPageStream {
	mock := &MockIStaticPageStream{ctrl: ctrl}
	mock.recorder = &MockIStaticPageStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStaticPageStream) EXPECT() *MockIStaticPageStreamMockRecorder {
	return m.recorder
}

// GetItemCount mocks base method.
func (m *MockIStaticPageStream) GetItemCount() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemCount indicates an expected call of GetItemCount.
func (mr *MockIStaticPageStreamMockRecorder) GetItemCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemCount", reflect.TypeOf((*MockIStaticPageStream)(nil).GetItemCount))
}

// GetItemIterator mocks base method.
func (m *MockIStaticPageStream) GetItemIterator() (pagination.IIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemIterator")
	ret0, _ := ret[0].(pagination.IIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemIterator indicates an expected call of GetItemIterator.
func (mr *MockIStaticPageStreamMockRecorder) GetItemIterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemIterator", reflect.TypeOf((*MockIStaticPageStream)(nil).GetItemIterator))
}

// HasFuture mocks base method.
func (m *MockIStaticPageStream) HasFuture() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasFuture")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasFuture indicates an expected call of HasFuture.
func (mr *MockIStaticPageStreamMockRecorder) HasFuture() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasFuture", reflect.TypeOf((*MockIStaticPageStream)(nil).HasFuture))
}

// HasNext mocks base method.
func (m *MockIStaticPageStream) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockIStaticPageStreamMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockIStaticPageStream)(nil).HasNext))
}

// MockIStream is a mock of IStream interface.
type MockIStream struct {
	ctrl     *gomock.Controller
	recorder *MockIStreamMockRecorder
}

// MockIStreamMockRecorder is the mock recorder for MockIStream.
type MockIStreamMockRecorder struct {
	mock *MockIStream
}

// NewMockIStream creates a new mock instance.
func NewMockIStream(ctrl *gomock.Controller) *MockIStream {
	mock := &MockIStream{ctrl: ctrl}
	mock.recorder = &MockIStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStream) EXPECT() *MockIStreamMockRecorder {
	return m.recorder
}

// GetFuture mocks base method.
func (m *MockIStream) GetFuture(arg0 context.Context) (pagination.IStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFuture", arg0)
	ret0, _ := ret[0].(pagination.IStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFuture indicates an expected call of GetFuture.
func (mr *MockIStreamMockRecorder) GetFuture(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFuture", reflect.TypeOf((*MockIStream)(nil).GetFuture), arg0)
}

// GetItemCount mocks base method.
func (m *MockIStream) GetItemCount() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemCount indicates an expected call of GetItemCount.
func (mr *MockIStreamMockRecorder) GetItemCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemCount", reflect.TypeOf((*MockIStream)(nil).GetItemCount))
}

// GetItemIterator mocks base method.
func (m *MockIStream) GetItemIterator() (pagination.IIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemIterator")
	ret0, _ := ret[0].(pagination.IIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemIterator indicates an expected call of GetItemIterator.
func (mr *MockIStreamMockRecorder) GetItemIterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemIterator", reflect.TypeOf((*MockIStream)(nil).GetItemIterator))
}

// GetNext mocks base method.
func (m *MockIStream) GetNext(arg0 context.Context) (pagination.IPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNext", arg0)
	ret0, _ := ret[0].(pagination.IPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNext indicates an expected call of GetNext.
func (mr *MockIStreamMockRecorder) GetNext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNext", reflect.TypeOf((*MockIStream)(nil).GetNext), arg0)
}

// HasFuture mocks base method.
func (m *MockIStream) HasFuture() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasFuture")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasFuture indicates an expected call of HasFuture.
func (mr *MockIStreamMockRecorder) HasFuture() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasFuture", reflect.TypeOf((*MockIStream)(nil).HasFuture))
}

// HasNext mocks base method.
func (m *MockIStream) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockIStreamMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockIStream)(nil).HasNext))
}

// MockIIterator is a mock of IIterator interface.
type MockIIterator struct {
	ctrl     *gomock.Controller
	recorder *MockIIteratorMockRecorder
}

// MockIIteratorMockRecorder is the mock recorder for MockIIterator.
type MockIIteratorMockRecorder struct {
	mock *MockIIterator
}

// NewMockIIterator creates a new mock instance.
func NewMockIIterator(ctrl *gomock.Controller) *MockIIterator {
	mock := &MockIIterator{ctrl: ctrl}
	mock.recorder = &MockIIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIIterator) EXPECT() *MockIIteratorMockRecorder {
	return m.recorder
}

// GetNext mocks base method.
func (m *MockIIterator) GetNext() (*interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNext")
	ret0, _ := ret[0].(*interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNext indicates an expected call of GetNext.
func (mr *MockIIteratorMockRecorder) GetNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNext", reflect.TypeOf((*MockIIterator)(nil).GetNext))
}

// HasNext mocks base method.
func (m *MockIIterator) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockIIteratorMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockIIterator)(nil).HasNext))
}

// MockIPaginator is a mock of IPaginator interface.
type MockIPaginator struct {
	ctrl     *gomock.Controller
	recorder *MockIPaginatorMockRecorder
}

// MockIPaginatorMockRecorder is the mock recorder for MockIPaginator.
type MockIPaginatorMockRecorder struct {
	mock *MockIPaginator
}

// NewMockIPaginator creates a new mock instance.
func NewMockIPaginator(ctrl *gomock.Controller) *MockIPaginator {
	mock := &MockIPaginator{ctrl: ctrl}
	mock.recorder = &MockIPaginatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPaginator) EXPECT() *MockIPaginatorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIPaginator) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIPaginatorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIPaginator)(nil).Close))
}

// GetCurrentPage mocks base method.
func (m *MockIPaginator) GetCurrentPage() (pagination.IPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentPage")
	ret0, _ := ret[0].(pagination.IPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentPage indicates an expected call of GetCurrentPage.
func (mr *MockIPaginatorMockRecorder) GetCurrentPage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentPage", reflect.TypeOf((*MockIPaginator)(nil).GetCurrentPage))
}

// GetNext mocks base method.
func (m *MockIPaginator) GetNext() (*interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNext")
	ret0, _ := ret[0].(*interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNext indicates an expected call of GetNext.
func (mr *MockIPaginatorMockRecorder) GetNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNext", reflect.TypeOf((*MockIPaginator)(nil).GetNext))
}

// HasNext mocks base method.
func (m *MockIPaginator) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockIPaginatorMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockIPaginator)(nil).HasNext))
}

// Stop mocks base method.
func (m *MockIPaginator) Stop() context.CancelFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(context.CancelFunc)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockIPaginatorMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockIPaginator)(nil).Stop))
}

// MockIPaginatorAndPageFetcher is a mock of IPaginatorAndPageFetcher interface.
type MockIPaginatorAndPageFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockIPaginatorAndPageFetcherMockRecorder
}

// MockIPaginatorAndPageFetcherMockRecorder is the mock recorder for MockIPaginatorAndPageFetcher.
type MockIPaginatorAndPageFetcherMockRecorder struct {
	mock *MockIPaginatorAndPageFetcher
}

// NewMockIPaginatorAndPageFetcher creates a new mock instance.
func NewMockIPaginatorAndPageFetcher(ctrl *gomock.Controller) *MockIPaginatorAndPageFetcher {
	mock := &MockIPaginatorAndPageFetcher{ctrl: ctrl}
	mock.recorder = &MockIPaginatorAndPageFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPaginatorAndPageFetcher) EXPECT() *MockIPaginatorAndPageFetcherMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIPaginatorAndPageFetcher) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIPaginatorAndPageFetcherMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIPaginatorAndPageFetcher)(nil).Close))
}

// FetchNextPage mocks base method.
func (m *MockIPaginatorAndPageFetcher) FetchNextPage(arg0 context.Context, arg1 pagination.IStaticPage) (pagination.IStaticPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchNextPage", arg0, arg1)
	ret0, _ := ret[0].(pagination.IStaticPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchNextPage indicates an expected call of FetchNextPage.
func (mr *MockIPaginatorAndPageFetcherMockRecorder) FetchNextPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchNextPage", reflect.TypeOf((*MockIPaginatorAndPageFetcher)(nil).FetchNextPage), arg0, arg1)
}

// GetCurrentPage mocks base method.
func (m *MockIPaginatorAndPageFetcher) GetCurrentPage() (pagination.IStaticPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentPage")
	ret0, _ := ret[0].(pagination.IStaticPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentPage indicates an expected call of GetCurrentPage.
func (mr *MockIPaginatorAndPageFetcherMockRecorder) GetCurrentPage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentPage", reflect.TypeOf((*MockIPaginatorAndPageFetcher)(nil).GetCurrentPage))
}

// GetNext mocks base method.
func (m *MockIPaginatorAndPageFetcher) GetNext() (*interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNext")
	ret0, _ := ret[0].(*interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNext indicates an expected call of GetNext.
func (mr *MockIPaginatorAndPageFetcherMockRecorder) GetNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNext", reflect.TypeOf((*MockIPaginatorAndPageFetcher)(nil).GetNext))
}

// HasNext mocks base method.
func (m *MockIPaginatorAndPageFetcher) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockIPaginatorAndPageFetcherMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockIPaginatorAndPageFetcher)(nil).HasNext))
}

// Stop mocks base method.
func (m *MockIPaginatorAndPageFetcher) Stop() context.CancelFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(context.CancelFunc)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockIPaginatorAndPageFetcherMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockIPaginatorAndPageFetcher)(nil).Stop))
}

// MockIStreamPaginator is a mock of IStreamPaginator interface.
type MockIStreamPaginator struct {
	ctrl     *gomock.Controller
	recorder *MockIStreamPaginatorMockRecorder
}

// MockIStreamPaginatorMockRecorder is the mock recorder for MockIStreamPaginator.
type MockIStreamPaginatorMockRecorder struct {
	mock *MockIStreamPaginator
}

// NewMockIStreamPaginator creates a new mock instance.
func NewMockIStreamPaginator(ctrl *gomock.Controller) *MockIStreamPaginator {
	mock := &MockIStreamPaginator{ctrl: ctrl}
	mock.recorder = &MockIStreamPaginatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStreamPaginator) EXPECT() *MockIStreamPaginatorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIStreamPaginator) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIStreamPaginatorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIStreamPaginator)(nil).Close))
}

// DryUp mocks base method.
func (m *MockIStreamPaginator) DryUp() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DryUp")
	ret0, _ := ret[0].(error)
	return ret0
}

// DryUp indicates an expected call of DryUp.
func (mr *MockIStreamPaginatorMockRecorder) DryUp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DryUp", reflect.TypeOf((*MockIStreamPaginator)(nil).DryUp))
}

// GetCurrentPage mocks base method.
func (m *MockIStreamPaginator) GetCurrentPage() (pagination.IPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentPage")
	ret0, _ := ret[0].(pagination.IPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentPage indicates an expected call of GetCurrentPage.
func (mr *MockIStreamPaginatorMockRecorder) GetCurrentPage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentPage", reflect.TypeOf((*MockIStreamPaginator)(nil).GetCurrentPage))
}

// GetNext mocks base method.
func (m *MockIStreamPaginator) GetNext() (*interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNext")
	ret0, _ := ret[0].(*interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNext indicates an expected call of GetNext.
func (mr *MockIStreamPaginatorMockRecorder) GetNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNext", reflect.TypeOf((*MockIStreamPaginator)(nil).GetNext))
}

// HasNext mocks base method.
func (m *MockIStreamPaginator) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockIStreamPaginatorMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockIStreamPaginator)(nil).HasNext))
}

// IsRunningDry mocks base method.
func (m *MockIStreamPaginator) IsRunningDry() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunningDry")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRunningDry indicates an expected call of IsRunningDry.
func (mr *MockIStreamPaginatorMockRecorder) IsRunningDry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunningDry", reflect.TypeOf((*MockIStreamPaginator)(nil).IsRunningDry))
}

// Stop mocks base method.
func (m *MockIStreamPaginator) Stop() context.CancelFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(context.CancelFunc)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockIStreamPaginatorMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockIStreamPaginator)(nil).Stop))
}

// MockIStreamPaginatorAndPageFetcher is a mock of IStreamPaginatorAndPageFetcher interface.
type MockIStreamPaginatorAndPageFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockIStreamPaginatorAndPageFetcherMockRecorder
}

// MockIStreamPaginatorAndPageFetcherMockRecorder is the mock recorder for MockIStreamPaginatorAndPageFetcher.
type MockIStreamPaginatorAndPageFetcherMockRecorder struct {
	mock *MockIStreamPaginatorAndPageFetcher
}

// NewMockIStreamPaginatorAndPageFetcher creates a new mock instance.
func NewMockIStreamPaginatorAndPageFetcher(ctrl *gomock.Controller) *MockIStreamPaginatorAndPageFetcher {
	mock := &MockIStreamPaginatorAndPageFetcher{ctrl: ctrl}
	mock.recorder = &MockIStreamPaginatorAndPageFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStreamPaginatorAndPageFetcher) EXPECT() *MockIStreamPaginatorAndPageFetcherMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIStreamPaginatorAndPageFetcher) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIStreamPaginatorAndPageFetcherMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIStreamPaginatorAndPageFetcher)(nil).Close))
}

// DryUp mocks base method.
func (m *MockIStreamPaginatorAndPageFetcher) DryUp() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DryUp")
	ret0, _ := ret[0].(error)
	return ret0
}

// DryUp indicates an expected call of DryUp.
func (mr *MockIStreamPaginatorAndPageFetcherMockRecorder) DryUp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DryUp", reflect.TypeOf((*MockIStreamPaginatorAndPageFetcher)(nil).DryUp))
}

// FetchFuturePage mocks base method.
func (m *MockIStreamPaginatorAndPageFetcher) FetchFuturePage(arg0 context.Context, arg1 pagination.IStaticPageStream) (pagination.IStaticPageStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchFuturePage", arg0, arg1)
	ret0, _ := ret[0].(pagination.IStaticPageStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchFuturePage indicates an expected call of FetchFuturePage.
func (mr *MockIStreamPaginatorAndPageFetcherMockRecorder) FetchFuturePage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchFuturePage", reflect.TypeOf((*MockIStreamPaginatorAndPageFetcher)(nil).FetchFuturePage), arg0, arg1)
}

// FetchNextPage mocks base method.
func (m *MockIStreamPaginatorAndPageFetcher) FetchNextPage(arg0 context.Context, arg1 pagination.IStaticPage) (pagination.IStaticPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchNextPage", arg0, arg1)
	ret0, _ := ret[0].(pagination.IStaticPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchNextPage indicates an expected call of FetchNextPage.
func (mr *MockIStreamPaginatorAndPageFetcherMockRecorder) FetchNextPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchNextPage", reflect.TypeOf((*MockIStreamPaginatorAndPageFetcher)(nil).FetchNextPage), arg0, arg1)
}

// GetCurrentPage mocks base method.
func (m *MockIStreamPaginatorAndPageFetcher) GetCurrentPage() (pagination.IStaticPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentPage")
	ret0, _ := ret[0].(pagination.IStaticPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentPage indicates an expected call of GetCurrentPage.
func (mr *MockIStreamPaginatorAndPageFetcherMockRecorder) GetCurrentPage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentPage", reflect.TypeOf((*MockIStreamPaginatorAndPageFetcher)(nil).GetCurrentPage))
}

// GetNext mocks base method.
func (m *MockIStreamPaginatorAndPageFetcher) GetNext() (*interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNext")
	ret0, _ := ret[0].(*interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNext indicates an expected call of GetNext.
func (mr *MockIStreamPaginatorAndPageFetcherMockRecorder) GetNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNext", reflect.TypeOf((*MockIStreamPaginatorAndPageFetcher)(nil).GetNext))
}

// HasNext mocks base method.
func (m *MockIStreamPaginatorAndPageFetcher) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockIStreamPaginatorAndPageFetcherMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockIStreamPaginatorAndPageFetcher)(nil).HasNext))
}

// IsRunningDry mocks base method.
func (m *MockIStreamPaginatorAndPageFetcher) IsRunningDry() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunningDry")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRunningDry indicates an expected call of IsRunningDry.
func (mr *MockIStreamPaginatorAndPageFetcherMockRecorder) IsRunningDry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunningDry", reflect.TypeOf((*MockIStreamPaginatorAndPageFetcher)(nil).IsRunningDry))
}

// Stop mocks base method.
func (m *MockIStreamPaginatorAndPageFetcher) Stop() context.CancelFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(context.CancelFunc)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockIStreamPaginatorAndPageFetcherMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockIStreamPaginatorAndPageFetcher)(nil).Stop))
}
