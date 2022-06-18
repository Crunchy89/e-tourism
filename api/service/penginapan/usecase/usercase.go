package usecase

import (
	"context"

	"api/domain"
	"api/utils/r"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PenginapanUseCase struct {
	Penginapan domain.PenginapanRepository
}

// funciton add user
func (p *PenginapanUseCase) AddPenginapan(domain *domain.Penginapan) (primitive.ObjectID, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Penginapan.Add(ctx, domain)
}

// function to get user by id
func (p *PenginapanUseCase) GetPenginapanByID(ID primitive.ObjectID) (*domain.Penginapan, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Penginapan.Fetch(ctx, ID)
}
