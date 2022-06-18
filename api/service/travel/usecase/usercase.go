package usecase

import (
	"context"

	"api/domain"
	"api/utils/r"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TravelUseCase struct {
	Travel domain.TravelRepository
}

// funciton add Travel
func (p *TravelUseCase) AddTravel(domain *domain.Travel) (primitive.ObjectID, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Travel.Add(ctx, domain)
}

// fetch all
func (p *TravelUseCase) FetchAllTravel() ([]*domain.Travel, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Travel.FetchAll(ctx)
}

// function to get Travel by id
func (p *TravelUseCase) GetTravelByID(ID primitive.ObjectID) (*domain.Travel, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Travel.FetchByID(ctx, ID)
}

// function to get Travel by slug
func (p *TravelUseCase) GetTravelBySlug(slug string) (*domain.Travel, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Travel.FetchBySlug(ctx, slug)
}

// function update by id
func (p *TravelUseCase) UpdateTravelByID(ID primitive.ObjectID, domain *domain.Travel) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Travel.UpdateByID(ctx, ID, domain)
}

// function delete
func (p *TravelUseCase) DeleteTravelByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Travel.DeleteByID(ctx, ID)
}

// function active
func (p *TravelUseCase) ActiveTravelByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Travel.ActiveByID(ctx, ID)
}
