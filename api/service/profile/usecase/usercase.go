package usecase

import (
	"context"

	"api/domain"
	"api/utils/r"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileUseCase struct {
	Profile domain.ProfileRepository
}

// funciton add profile
func (p *ProfileUseCase) AddProfile(domain *domain.Profile) (primitive.ObjectID, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Profile.Add(ctx, domain)
}

// fetch all
func (p *ProfileUseCase) FetchAllProfile() ([]*domain.Profile, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Profile.FetchAll(ctx)
}

// function to get profile by id
func (p *ProfileUseCase) GetProfileByID(ID primitive.ObjectID) (*domain.Profile, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Profile.FetchByID(ctx, ID)
}

// function update by id
func (p *ProfileUseCase) UpdateProfileByID(ID primitive.ObjectID, domain *domain.Profile) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Profile.UpdateByID(ctx, ID, domain)
}

// fucntion delete
func (p *ProfileUseCase) DeleteProfileByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Profile.DeleteByID(ctx, ID)
}
