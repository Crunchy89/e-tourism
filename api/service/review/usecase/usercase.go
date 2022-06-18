package usecase

import (
	"context"

	"api/domain"
	"api/utils/r"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReviewUseCase struct {
	Review domain.ReviewRepository
}

// funciton add review
func (p *ReviewUseCase) AddReview(domain *domain.Review) (primitive.ObjectID, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Review.Add(ctx, domain)
}

// fetch all
func (p *ReviewUseCase) FetchAllReview() ([]*domain.Review, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Review.FetchAll(ctx)
}

// function to get review by id
func (p *ReviewUseCase) GetReviewByID(ID primitive.ObjectID) (*domain.Review, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Review.FetchByID(ctx, ID)
}

// function update by id
func (p *ReviewUseCase) UpdateReviewByID(ID primitive.ObjectID, domain *domain.Review) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Review.UpdateByID(ctx, ID, domain)
}

// fucntion delete
func (p *ReviewUseCase) DeleteReviewByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Review.DeleteByID(ctx, ID)
}
