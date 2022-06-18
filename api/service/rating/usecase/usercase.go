package usecase

import (
	"context"

	"api/domain"
	"api/utils/r"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingUseCase struct {
	Rating domain.RatingRepository
}

// funciton add rating
func (p *RatingUseCase) AddRating(domain *domain.Rating) (primitive.ObjectID, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Rating.Add(ctx, domain)
}

// fetch all
func (p *RatingUseCase) FetchAllRating() ([]*domain.Rating, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Rating.FetchAll(ctx)
}

// function to get rating by id
func (p *RatingUseCase) GetRatingByID(ID primitive.ObjectID) (*domain.Rating, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Rating.FetchByID(ctx, ID)
}
