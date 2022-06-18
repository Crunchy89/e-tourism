package usecase

import (
	"context"

	"api/domain"
	"api/utils/r"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SearchedUseCase struct {
	Searched domain.SearchedRepository
}

// funciton add searched
func (p *SearchedUseCase) AddSearched(domain *domain.Searched) (primitive.ObjectID, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Searched.Add(ctx, domain)
}

// fetch all
func (p *SearchedUseCase) FetchAllSearched() ([]*domain.Searched, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Searched.FetchAll(ctx)
}

// function to get searched by id
func (p *SearchedUseCase) GetSearchedByID(ID primitive.ObjectID) (*domain.Searched, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Searched.FetchByID(ctx, ID)
}

// function update by id
func (p *SearchedUseCase) UpdateSearchedByID(ID primitive.ObjectID, domain *domain.Searched) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Searched.UpdateByID(ctx, ID, domain)
}

// fucntion delete
func (p *SearchedUseCase) DeleteSearchedByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Searched.DeleteByID(ctx, ID)
}
