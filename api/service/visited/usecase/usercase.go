package usecase

import (
	"context"

	"api/domain"
	"api/utils/r"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VisitedUseCase struct {
	Visited domain.VisitedRepository
}

// funciton add visited
func (p *VisitedUseCase) AddVisited(domain *domain.Visited) (primitive.ObjectID, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Visited.Add(ctx, domain)
}

// fetch all
func (p *VisitedUseCase) FetchAllVisited() ([]*domain.Visited, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Visited.FetchAll(ctx)
}

// function to get visited by id
func (p *VisitedUseCase) GetVisitedByID(ID primitive.ObjectID) (*domain.Visited, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Visited.FetchByID(ctx, ID)
}

// function update by id
func (p *VisitedUseCase) UpdateVisitedByID(ID primitive.ObjectID, domain *domain.Visited) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Visited.UpdateByID(ctx, ID, domain)
}

// fucntion delete
func (p *VisitedUseCase) DeleteVisitedByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Visited.DeleteByID(ctx, ID)
}
