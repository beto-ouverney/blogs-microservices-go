// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/beto-ouverney/blogs-microservices/categories/server/entity"
	errors "github.com/beto-ouverney/blogs-microservices/categories/server/errors"

	mock "github.com/stretchr/testify/mock"
)

// ICategoryModel is an autogenerated mock type for the ICategoryModel type
type ICategoryModel struct {
	mock.Mock
}

// Add provides a mock function with given fields: name
func (_m *ICategoryModel) Add(name string) (*entity.Category, *errors.CustomError) {
	ret := _m.Called(name)

	var r0 *entity.Category
	if rf, ok := ret.Get(0).(func(string) *entity.Category); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Category)
		}
	}

	var r1 *errors.CustomError
	if rf, ok := ret.Get(1).(func(string) *errors.CustomError); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *ICategoryModel) GetAll() (*[]entity.Category, *errors.CustomError) {
	ret := _m.Called()

	var r0 *[]entity.Category
	if rf, ok := ret.Get(0).(func() *[]entity.Category); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.Category)
		}
	}

	var r1 *errors.CustomError
	if rf, ok := ret.Get(1).(func() *errors.CustomError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: name
func (_m *ICategoryModel) GetByName(name string) (*entity.Category, *errors.CustomError) {
	ret := _m.Called(name)

	var r0 *entity.Category
	if rf, ok := ret.Get(0).(func(string) *entity.Category); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Category)
		}
	}

	var r1 *errors.CustomError
	if rf, ok := ret.Get(1).(func(string) *errors.CustomError); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.CustomError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewICategoryModel interface {
	mock.TestingT
	Cleanup(func())
}

// NewICategoryModel creates a new instance of ICategoryModel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICategoryModel(t mockConstructorTestingTNewICategoryModel) *ICategoryModel {
	mock := &ICategoryModel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}