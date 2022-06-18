package usecase

import (
	"context"

	"api/domain"
	"api/utils/r"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FasilitasUseCase struct {
	Fasilitas domain.FasilitasRepository
}

// funciton add fasilitas
func (p *FasilitasUseCase) AddFasilitas(domain *domain.Fasilitas) (primitive.ObjectID, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Fasilitas.Add(ctx, domain)
}

// fetch all
func (p *FasilitasUseCase) FetchAllFasilitas() ([]*domain.Fasilitas, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Fasilitas.FetchAll(ctx)
}

// function to get fasilitas by id
func (p *FasilitasUseCase) GetFasilitasByID(ID primitive.ObjectID) (*domain.Fasilitas, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Fasilitas.FetchByID(ctx, ID)
}

// function update by id
func (p *FasilitasUseCase) UpdateFasilitasByID(ID primitive.ObjectID, domain *domain.Fasilitas) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Fasilitas.UpdateByID(ctx, ID, domain)
}

// fucntion delete
func (p *FasilitasUseCase) DeleteFasilitasByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Fasilitas.DeleteByID(ctx, ID)
}

// function activate
func (p *FasilitasUseCase) ActivateFasilitasByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Fasilitas.Active(ctx, ID)
}
