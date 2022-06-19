package usecase

import (
	"context"

	"api/domain"
	"api/utils/r"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ViewedUseCase struct {
	Viewed domain.ViewedRepository
}

// funciton add viewed
func (p *ViewedUseCase) AddViewed(domain *domain.Viewed) (primitive.ObjectID, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Viewed.Add(ctx, domain)
}

// fetch all
func (p *ViewedUseCase) FetchAllViewed() ([]*domain.Viewed, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Viewed.FetchAll(ctx)
}

// function to get viewed by id
func (p *ViewedUseCase) GetViewedByID(ID primitive.ObjectID) (*domain.Viewed, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Viewed.FetchByID(ctx, ID)
}

// function update by id
func (p *ViewedUseCase) UpdateViewedByID(ID primitive.ObjectID, domain *domain.Viewed) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Viewed.UpdateByID(ctx, ID, domain)
}

// fucntion delete
func (p *ViewedUseCase) DeleteViewedByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Viewed.DeleteByID(ctx, ID)
}
