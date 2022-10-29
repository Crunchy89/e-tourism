// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "api/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	r "api/utils/r"
)

// FasilitasRepository is an autogenerated mock type for the FasilitasRepository type
type FasilitasRepository struct {
	mock.Mock
}

// Active provides a mock function with given fields: ctx, ID
func (_m *FasilitasRepository) Active(ctx context.Context, ID primitive.ObjectID) r.Ex {
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
func (_m *FasilitasRepository) Add(ctx context.Context, d *domain.Fasilitas) (primitive.ObjectID, r.Ex) {
	ret := _m.Called(ctx, d)

	var r0 primitive.ObjectID
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Fasilitas) primitive.ObjectID); ok {
		r0 = rf(ctx, d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(primitive.ObjectID)
		}
	}

	var r1 r.Ex
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Fasilitas) r.Ex); ok {
		r1 = rf(ctx, d)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(r.Ex)
		}
	}

	return r0, r1
}

// DeleteByID provides a mock function with given fields: ctx, ID
func (_m *FasilitasRepository) DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex {
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
func (_m *FasilitasRepository) FetchAll(ctx context.Context) ([]*domain.Fasilitas, r.Ex) {
	ret := _m.Called(ctx)

	var r0 []*domain.Fasilitas
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Fasilitas); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Fasilitas)
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

// FetchByForeignID provides a mock function with given fields: ctx, foreignID
func (_m *FasilitasRepository) FetchByForeignID(ctx context.Context, foreignID primitive.ObjectID) ([]*domain.Fasilitas, r.Ex) {
	ret := _m.Called(ctx, foreignID)

	var r0 []*domain.Fasilitas
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) []*domain.Fasilitas); ok {
		r0 = rf(ctx, foreignID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Fasilitas)
		}
	}

	var r1 r.Ex
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) r.Ex); ok {
		r1 = rf(ctx, foreignID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(r.Ex)
		}
	}

	return r0, r1
}

// FetchByID provides a mock function with given fields: ctx, ID
func (_m *FasilitasRepository) FetchByID(ctx context.Context, ID primitive.ObjectID) (*domain.Fasilitas, r.Ex) {
	ret := _m.Called(ctx, ID)

	var r0 *domain.Fasilitas
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) *domain.Fasilitas); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Fasilitas)
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

// PaginationByForeignID provides a mock function with given fields: ctx, foreignID, page, limit
func (_m *FasilitasRepository) PaginationByForeignID(ctx context.Context, foreignID primitive.ObjectID, page int64, limit int64) ([]*domain.Fasilitas, r.Ex) {
	ret := _m.Called(ctx, foreignID, page, limit)

	var r0 []*domain.Fasilitas
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, int64, int64) []*domain.Fasilitas); ok {
		r0 = rf(ctx, foreignID, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Fasilitas)
		}
	}

	var r1 r.Ex
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID, int64, int64) r.Ex); ok {
		r1 = rf(ctx, foreignID, page, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(r.Ex)
		}
	}

	return r0, r1
}

// SearchByForeignID provides a mock function with given fields: ctx, foreignID, page, limit, keyword
func (_m *FasilitasRepository) SearchByForeignID(ctx context.Context, foreignID primitive.ObjectID, page int64, limit int64, keyword string) ([]*domain.Fasilitas, r.Ex) {
	ret := _m.Called(ctx, foreignID, page, limit, keyword)

	var r0 []*domain.Fasilitas
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, int64, int64, string) []*domain.Fasilitas); ok {
		r0 = rf(ctx, foreignID, page, limit, keyword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Fasilitas)
		}
	}

	var r1 r.Ex
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID, int64, int64, string) r.Ex); ok {
		r1 = rf(ctx, foreignID, page, limit, keyword)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(r.Ex)
		}
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: ctx, ID, d
func (_m *FasilitasRepository) UpdateByID(ctx context.Context, ID primitive.ObjectID, d *domain.Fasilitas) r.Ex {
	ret := _m.Called(ctx, ID, d)

	var r0 r.Ex
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, *domain.Fasilitas) r.Ex); ok {
		r0 = rf(ctx, ID, d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(r.Ex)
		}
	}

	return r0
}

type mockConstructorTestingTNewFasilitasRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewFasilitasRepository creates a new instance of FasilitasRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFasilitasRepository(t mockConstructorTestingTNewFasilitasRepository) *FasilitasRepository {
	mock := &FasilitasRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
