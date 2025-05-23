// Code generated by mockery v2.53.3. DO NOT EDIT.

package repository

import (
	entity "bot/domain/article/scraper/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockArticle is an autogenerated mock type for the Article type
type MockArticle struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *MockArticle) Create(_a0 entity.Article) (int, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Article) (int, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(entity.Article) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(entity.Article) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with no fields
func (_m *MockArticle) Find() []entity.Article {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 []entity.Article
	if rf, ok := ret.Get(0).(func() []entity.Article); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Article)
		}
	}

	return r0
}

// FindByOption provides a mock function with given fields: _a0
func (_m *MockArticle) FindByOption(_a0 map[string]interface{}) []entity.Article {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for FindByOption")
	}

	var r0 []entity.Article
	if rf, ok := ret.Get(0).(func(map[string]interface{}) []entity.Article); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Article)
		}
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *MockArticle) Update(_a0 entity.Article, _a1 []string) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Article, []string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockArticle creates a new instance of MockArticle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockArticle(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockArticle {
	mock := &MockArticle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
