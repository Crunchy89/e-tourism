// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "api/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	r "api/utils/r"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// ActiveByID provides a mock function with given fields: ctx, ID
func (_m *UserRepository) ActiveByID(ctx context.Context, ID primitive.ObjectID) r.Ex {
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
func (_m *UserRepository) Add(ctx context.Context, d *domain.User) (primitive.ObjectID, r.Ex) {
	ret := _m.Called(ctx, d)

	var r0 primitive.ObjectID
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) primitive.ObjectID); ok {
		r0 = rf(ctx, d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(primitive.ObjectID)
		}
	}

	var r1 r.Ex
	if rf, ok := ret.Get(1).(func(context.Context, *domain.User) r.Ex); ok {
		r1 = rf(ctx, d)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(r.Ex)
		}
	}

	return r0, r1
}

// CountAll provides a mock function with given fields: ctx
func (_m *UserRepository) CountAll(ctx context.Context) (int64, r.Ex) {
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
func (_m *UserRepository) DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex {
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
func (_m *UserRepository) FetchAll(ctx context.Context) ([]*domain.User, r.Ex) {
	ret := _m.Called(ctx)

	var r0 []*domain.User
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.User)
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

// FetchByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepository) FetchByEmail(ctx context.Context, email string) (*domain.User, r.Ex) {
	ret := _m.Called(ctx, email)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 r.Ex
	if rf, ok := ret.Get(1).(func(context.Context, string) r.Ex); ok {
		r1 = rf(ctx, email)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(r.Ex)
		}
	}

	return r0, r1
}

// FetchByID provides a mock function with given fields: ctx, ID
func (_m *UserRepository) FetchByID(ctx context.Context, ID primitive.ObjectID) (*domain.User, r.Ex) {
	ret := _m.Called(ctx, ID)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) *domain.User); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
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

// FetchByUsername provides a mock function with given fields: ctx, username
func (_m *UserRepository) FetchByUsername(ctx context.Context, username string) (*domain.User, r.Ex) {
	ret := _m.Called(ctx, username)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.User); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 r.Ex
	if rf, ok := ret.Get(1).(func(context.Context, string) r.Ex); ok {
		r1 = rf(ctx, username)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(r.Ex)
		}
	}

	return r0, r1
}

// Pagination provides a mock function with given fields: ctx, page, limit
func (_m *UserRepository) Pagination(ctx context.Context, page int64, limit int64) ([]*domain.User, r.Ex) {
	ret := _m.Called(ctx, page, limit)

	var r0 []*domain.User
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) []*domain.User); ok {
		r0 = rf(ctx, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.User)
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
func (_m *UserRepository) Search(ctx context.Context, page int64, limit int64, search string) ([]*domain.User, r.Ex) {
	ret := _m.Called(ctx, page, limit, search)

	var r0 []*domain.User
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, string) []*domain.User); ok {
		r0 = rf(ctx, page, limit, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.User)
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

// UpdateByEmail provides a mock function with given fields: ctx, email, d
func (_m *UserRepository) UpdateByEmail(ctx context.Context, email string, d *domain.User) r.Ex {
	ret := _m.Called(ctx, email, d)

	var r0 r.Ex
	if rf, ok := ret.Get(0).(func(context.Context, string, *domain.User) r.Ex); ok {
		r0 = rf(ctx, email, d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(r.Ex)
		}
	}

	return r0
}

// UpdateByID provides a mock function with given fields: ctx, ID, d
func (_m *UserRepository) UpdateByID(ctx context.Context, ID primitive.ObjectID, d *domain.User) r.Ex {
	ret := _m.Called(ctx, ID, d)

	var r0 r.Ex
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, *domain.User) r.Ex); ok {
		r0 = rf(ctx, ID, d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(r.Ex)
		}
	}

	return r0
}

// UpdateByUsername provides a mock function with given fields: ctx, username, d
func (_m *UserRepository) UpdateByUsername(ctx context.Context, username string, d *domain.User) r.Ex {
	ret := _m.Called(ctx, username, d)

	var r0 r.Ex
	if rf, ok := ret.Get(0).(func(context.Context, string, *domain.User) r.Ex); ok {
		r0 = rf(ctx, username, d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(r.Ex)
		}
	}

	return r0
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
