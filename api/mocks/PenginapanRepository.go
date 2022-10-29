// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	domain "api/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	r "api/utils/r"
)

// PenginapanRepository is an autogenerated mock type for the PenginapanRepository type
type PenginapanRepository struct {
	mock.Mock
}

// Active provides a mock function with given fields: ctx, ID
func (_m *PenginapanRepository) Active(ctx context.Context, ID primitive.ObjectID) r.Ex {
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
func (_m *PenginapanRepository) Add(ctx context.Context, d *domain.Penginapan) (primitive.ObjectID, r.Ex) {
	ret := _m.Called(ctx, d)

	var r0 primitive.ObjectID
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Penginapan) primitive.ObjectID); ok {
		r0 = rf(ctx, d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(primitive.ObjectID)
		}
	}

	var r1 r.Ex
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Penginapan) r.Ex); ok {
		r1 = rf(ctx, d)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(r.Ex)
		}
	}

	return r0, r1
}

// CountAll provides a mock function with given fields: ctx
func (_m *PenginapanRepository) CountAll(ctx context.Context) (int64, r.Ex) {
	ret := _m.Called(ctx)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
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

// DeleteByID provides a mock function with given fields: ctx, ID
func (_m *PenginapanRepository) DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex {
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
func (_m *PenginapanRepository) FetchAll(ctx context.Context) ([]*domain.Penginapan, r.Ex) {
	ret := _m.Called(ctx)

	var r0 []*domain.Penginapan
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Penginapan); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Penginapan)
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
func (_m *PenginapanRepository) FetchByID(ctx context.Context, ID primitive.ObjectID) (*domain.Penginapan, r.Ex) {
	ret := _m.Called(ctx, ID)

	var r0 *domain.Penginapan
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) *domain.Penginapan); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Penginapan)
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
func (_m *PenginapanRepository) FetchBySlug(ctx context.Context, slug string) (*domain.Penginapan, r.Ex) {
	ret := _m.Called(ctx, slug)

	var r0 *domain.Penginapan
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Penginapan); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Penginapan)
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

// Pagination provides a mock function with given fields: ctx, page, limit
func (_m *PenginapanRepository) Pagination(ctx context.Context, page int64, limit int64) ([]*domain.Penginapan, r.Ex) {
	ret := _m.Called(ctx, page, limit)

	var r0 []*domain.Penginapan
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) []*domain.Penginapan); ok {
		r0 = rf(ctx, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Penginapan)
		}
	}

	var r1 r.Ex
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) r.Ex); ok {
		r1 = rf(ctx, page, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(r.Ex)
		}
	}

	return r0, r1
}

// Search provides a mock function with given fields: ctx, page, limit, search
func (_m *PenginapanRepository) Search(ctx context.Context, page int64, limit int64, search string) ([]*domain.Penginapan, r.Ex) {
	ret := _m.Called(ctx, page, limit, search)

	var r0 []*domain.Penginapan
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, string) []*domain.Penginapan); ok {
		r0 = rf(ctx, page, limit, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Penginapan)
		}
	}

	var r1 r.Ex
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, string) r.Ex); ok {
		r1 = rf(ctx, page, limit, search)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(r.Ex)
		}
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: ctx, ID, d
func (_m *PenginapanRepository) UpdateByID(ctx context.Context, ID primitive.ObjectID, d *domain.Penginapan) r.Ex {
	ret := _m.Called(ctx, ID, d)

	var r0 r.Ex
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, *domain.Penginapan) r.Ex); ok {
		r0 = rf(ctx, ID, d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(r.Ex)
		}
	}

	return r0
}

type NewPenginapanRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewPenginapanRepository creates a new instance of PenginapanRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPenginapanRepository(t NewPenginapanRepositoryT) *PenginapanRepository {
	mock := &PenginapanRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
