// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	domain "api/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	r "api/utils/r"
)

// SearchedRepository is an autogenerated mock type for the SearchedRepository type
type SearchedRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, d
func (_m *SearchedRepository) Add(ctx context.Context, d *domain.Searched) (primitive.ObjectID, r.Ex) {
	ret := _m.Called(ctx, d)

	var r0 primitive.ObjectID
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Searched) primitive.ObjectID); ok {
		r0 = rf(ctx, d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(primitive.ObjectID)
		}
	}

	var r1 r.Ex
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Searched) r.Ex); ok {
		r1 = rf(ctx, d)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(r.Ex)
		}
	}

	return r0, r1
}

// DeleteByID provides a mock function with given fields: ctx, ID
func (_m *SearchedRepository) DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex {
	ret := _m.Called(ctx, ID)

	var r0 r.Ex
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) r.Ex); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(r.Ex)
		}
	}

	return r0
}

// FetchAll provides a mock function with given fields: ctx
func (_m *SearchedRepository) FetchAll(ctx context.Context) ([]*domain.Searched, r.Ex) {
	ret := _m.Called(ctx)

	var r0 []*domain.Searched
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Searched); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Searched)
		}
	}

	var r1 r.Ex
	if rf, ok := ret.Get(1).(func(context.Context) r.Ex); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(r.Ex)
		}
	}

	return r0, r1
}

// FetchByID provides a mock function with given fields: ctx, ID
func (_m *SearchedRepository) FetchByID(ctx context.Context, ID primitive.ObjectID) (*domain.Searched, r.Ex) {
	ret := _m.Called(ctx, ID)

	var r0 *domain.Searched
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) *domain.Searched); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Searched)
		}
	}

	var r1 r.Ex
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) r.Ex); ok {
		r1 = rf(ctx, ID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(r.Ex)
		}
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: ctx, ID, d
func (_m *SearchedRepository) UpdateByID(ctx context.Context, ID primitive.ObjectID, d *domain.Searched) r.Ex {
	ret := _m.Called(ctx, ID, d)

	var r0 r.Ex
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, *domain.Searched) r.Ex); ok {
		r0 = rf(ctx, ID, d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(r.Ex)
		}
	}

	return r0
}

type NewSearchedRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewSearchedRepository creates a new instance of SearchedRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSearchedRepository(t NewSearchedRepositoryT) *SearchedRepository {
	mock := &SearchedRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
