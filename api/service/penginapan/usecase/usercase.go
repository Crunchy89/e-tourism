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

// fetch all
func (p *PenginapanUseCase) FetchAllPenginapan() ([]*domain.Penginapan, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Penginapan.FetchAll(ctx)
}

// function to get user by id
func (p *PenginapanUseCase) GetPenginapanByID(ID primitive.ObjectID) (*domain.Penginapan, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Penginapan.FetchByID(ctx, ID)
}

// function to get data by slug
func (p *PenginapanUseCase) GetPenginapanBySlug(slug string) (*domain.Penginapan, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Penginapan.FetchBySlug(ctx, slug)
}

// function update by id
func (p *PenginapanUseCase) UpdatePenginapanByID(ID primitive.ObjectID, domain *domain.Penginapan) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Penginapan.UpdateByID(ctx, ID, domain)
}

// fucntion delete
func (p *PenginapanUseCase) DeletePenginapanByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Penginapan.DeleteByID(ctx, ID)
}

// function activate
func (p *PenginapanUseCase) ActivatePenginapanByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Penginapan.Active(ctx, ID)
}
