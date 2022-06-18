package usecase

import (
	"context"

	"api/domain"
	"api/utils/r"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FotoUseCase struct {
	Foto domain.FotoRepository
}

// funciton add foto
func (p *FotoUseCase) AddFoto(domain *domain.Foto) (primitive.ObjectID, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Foto.Add(ctx, domain)
}

// fetch all
func (p *FotoUseCase) FetchAllFoto() ([]*domain.Foto, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Foto.FetchAll(ctx)
}

// function to get foto by id
func (p *FotoUseCase) GetFotoByID(ID primitive.ObjectID) (*domain.Foto, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Foto.FetchByID(ctx, ID)
}

// function update by id
func (p *FotoUseCase) UpdateFotoByID(ID primitive.ObjectID, domain *domain.Foto) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Foto.UpdateByID(ctx, ID, domain)
}

// fucntion delete
func (p *FotoUseCase) DeleteFotoByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Foto.DeleteByID(ctx, ID)
}
