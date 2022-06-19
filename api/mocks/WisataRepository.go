// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	domain "api/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	r "api/utils/r"
)

// WisataRepository is an autogenerated mock type for the WisataRepository type
type WisataRepository struct {
	mock.Mock
}

// Active provides a mock function with given fields: ctx, ID
func (_m *WisataRepository) Active(ctx context.Context, ID primitive.ObjectID) r.Ex {
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
func (_m *WisataRepository) Add(ctx context.Context, d *domain.Wisata) (primitive.ObjectID, r.Ex) {
	ret := _m.Called(ctx, d)

	var r0 primitive.ObjectID
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Wisata) primitive.ObjectID); ok {
		r0 = rf(ctx, d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(primitive.ObjectID)
		}
	}

	var r1 r.Ex
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Wisata) r.Ex); ok {
		r1 = rf(ctx, d)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(r.Ex)
		}
	}

	return r0, r1
}

// DeleteByID provides a mock function with given fields: ctx, ID
func (_m *WisataRepository) DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex {
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
func (_m *WisataRepository) FetchAll(ctx context.Context) ([]*domain.Wisata, r.Ex) {
	ret := _m.Called(ctx)

	var r0 []*domain.Wisata
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Wisata); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Wisata)
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
func (_m *WisataRepository) FetchByID(ctx context.Context, ID primitive.ObjectID) (*domain.Wisata, r.Ex) {
	ret := _m.Called(ctx, ID)

	var r0 *domain.Wisata
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) *domain.Wisata); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Wisata)
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
func (_m *WisataRepository) FetchBySlug(ctx context.Context, slug string) (*domain.Wisata, r.Ex) {
	ret := _m.Called(ctx, slug)

	var r0 *domain.Wisata
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Wisata); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Wisata)
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
func (_m *WisataRepository) UpdateByID(ctx context.Context, ID primitive.ObjectID, d *domain.Wisata) r.Ex {
	ret := _m.Called(ctx, ID, d)

	var r0 r.Ex
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, *domain.Wisata) r.Ex); ok {
		r0 = rf(ctx, ID, d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(r.Ex)
		}
	}

	return r0
}

type NewWisataRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewWisataRepository creates a new instance of WisataRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWisataRepository(t NewWisataRepositoryT) *WisataRepository {
	mock := &WisataRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
