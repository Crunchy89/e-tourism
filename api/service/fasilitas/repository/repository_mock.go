package repository

import (
	"api/domain"
	"context"

	"api/utils/r"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FasilitasRepoMock struct {
	mock.Mock
}

func (f *FasilitasRepoMock) Add(ctx context.Context, d *domain.Fasilitas) (primitive.ObjectID, r.Ex) {
	arguments := f.Mock.Called(ctx, d)
	if arguments.Get(0) == nil {
		return arguments.Get(0).(primitive.ObjectID), arguments.Get(1).(r.Ex)
	} else {
		return arguments.Get(0).(primitive.ObjectID), nil
	}
}

func (f *FasilitasRepoMock) Delete(ctx context.Context, ID primitive.ObjectID) r.Ex {
	arguments := f.Mock.Called(ctx, ID)
	if arguments.Get(0) == nil {
		return arguments.Get(1).(r.Ex)
	} else {
		return nil
	}
}
func (f *FasilitasRepoMock) FetchAll(ctx context.Context) ([]*domain.Fasilitas, r.Ex) {
	arguments := f.Mock.Called(ctx)
	if arguments.Get(0) == nil {
		return arguments.Get(0).([]*domain.Fasilitas), arguments.Get(1).(r.Ex)
	} else {
		return arguments.Get(0).([]*domain.Fasilitas), nil
	}
}
func (f *FasilitasRepoMock) FetchByID(ctx context.Context, ID primitive.ObjectID) (*domain.Fasilitas, r.Ex) {
	arguments := f.Mock.Called(ctx, ID)
	if arguments.Get(0) == nil {
		return arguments.Get(0).(*domain.Fasilitas), arguments.Get(1).(r.Ex)
	} else {
		return arguments.Get(0).(*domain.Fasilitas), nil
	}
}

func (f *FasilitasRepoMock) UpdateByID(ctx context.Context, ID primitive.ObjectID, d *domain.Fasilitas) r.Ex {
	arguments := f.Mock.Called(ctx, ID)
	if arguments.Get(0) == nil {
		return arguments.Get(1).(r.Ex)
	} else {
		return nil
	}
}
func (f *FasilitasRepoMock) DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex {
	arguments := f.Mock.Called(ctx, ID)
	if arguments.Get(0) == nil {
		return arguments.Get(1).(r.Ex)
	} else {
		return nil
	}
}
func (f *FasilitasRepoMock) Active(ctx context.Context, ID primitive.ObjectID) r.Ex {
	arguments := f.Mock.Called(ctx, ID)
	if arguments.Get(0) == nil {
		return arguments.Get(1).(r.Ex)
	} else {
		return nil
	}
}
func (f *FasilitasRepoMock) FetchByForeignID(ctx context.Context, foreignID primitive.ObjectID) ([]*domain.Fasilitas, r.Ex) {
	arguments := f.Mock.Called(ctx, foreignID)
	if arguments.Get(0) == nil {
		return arguments.Get(0).([]*domain.Fasilitas), arguments.Get(1).(r.Ex)
	} else {
		return arguments.Get(0).([]*domain.Fasilitas), nil
	}
}
func (f *FasilitasRepoMock) PaginationByForeignID(ctx context.Context, foreignID primitive.ObjectID, page int64, limit int64) ([]*domain.Fasilitas, r.Ex) {
	arguments := f.Mock.Called(ctx, foreignID, page, limit)
	if arguments.Get(0) == nil {
		return arguments.Get(0).([]*domain.Fasilitas), arguments.Get(1).(r.Ex)
	} else {
		return arguments.Get(0).([]*domain.Fasilitas), nil
	}
}
func (f *FasilitasRepoMock) SearchByForeignID(ctx context.Context, foreignID primitive.ObjectID, page int64, limit int64, keyword string) ([]*domain.Fasilitas, r.Ex) {
	arguments := f.Mock.Called(ctx, foreignID, page, limit, keyword)
	if arguments.Get(0) == nil {
		return arguments.Get(0).([]*domain.Fasilitas), arguments.Get(1).(r.Ex)
	} else {
		return arguments.Get(0).([]*domain.Fasilitas), nil
	}
}
