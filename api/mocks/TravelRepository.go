// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "api/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	r "api/utils/r"
)

// TravelRepository is an autogenerated mock type for the TravelRepository type
type TravelRepository struct {
	mock.Mock
}

// ActiveByID provides a mock function with given fields: ctx, ID
func (_m *TravelRepository) ActiveByID(ctx context.Context, ID primitive.ObjectID) r.Ex {
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

// Add provides a mock function with given fields: ctx, d
func (_m *TravelRepository) Add(ctx context.Context, d *domain.Travel) (primitive.ObjectID, r.Ex) {
	ret := _m.Called(ctx, d)

	var r0 primitive.ObjectID
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Travel) primitive.ObjectID); ok {
		r0 = rf(ctx, d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(primitive.ObjectID)
		}
	}

	var r1 r.Ex
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Travel) r.Ex); ok {
		r1 = rf(ctx, d)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(r.Ex)
		}
	}

	return r0, r1
}

// DeleteByID provides a mock function with given fields: ctx, ID
func (_m *TravelRepository) DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex {
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
func (_m *TravelRepository) FetchAll(ctx context.Context) ([]*domain.Travel, r.Ex) {
	ret := _m.Called(ctx)

	var r0 []*domain.Travel
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Travel); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Travel)
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
func (_m *TravelRepository) FetchByID(ctx context.Context, ID primitive.ObjectID) (*domain.Travel, r.Ex) {
	ret := _m.Called(ctx, ID)

	var r0 *domain.Travel
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) *domain.Travel); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Travel)
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

// FetchBySlug provides a mock function with given fields: ctx, slug
func (_m *TravelRepository) FetchBySlug(ctx context.Context, slug string) (*domain.Travel, r.Ex) {
	ret := _m.Called(ctx, slug)

	var r0 *domain.Travel
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Travel); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Travel)
		}
	}

	var r1 r.Ex
	if rf, ok := ret.Get(1).(func(context.Context, string) r.Ex); ok {
		r1 = rf(ctx, slug)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(r.Ex)
		}
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: ctx, ID, d
func (_m *TravelRepository) UpdateByID(ctx context.Context, ID primitive.ObjectID, d *domain.Travel) r.Ex {
	ret := _m.Called(ctx, ID, d)

	var r0 r.Ex
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, *domain.Travel) r.Ex); ok {
		r0 = rf(ctx, ID, d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(r.Ex)
		}
	}

	return r0
}

type mockConstructorTestingTNewTravelRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTravelRepository creates a new instance of TravelRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTravelRepository(t mockConstructorTestingTNewTravelRepository) *TravelRepository {
	mock := &TravelRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
