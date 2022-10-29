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

// function to get fasilitas by foreign id
func (p *FasilitasUseCase) GetFasilitasByForeignID(foreignID primitive.ObjectID) ([]*domain.Fasilitas, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Fasilitas.FetchByForeignID(ctx, foreignID)
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

// function to get fasilitas by foreign id pagination
func (p *FasilitasUseCase) GetFasilitasByForeignIDPagination(foreignID primitive.ObjectID, page int64, limit int64) ([]*domain.Fasilitas, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Fasilitas.PaginationByForeignID(ctx, foreignID, page, limit)
}

// function to get fasilitas search pagination by foreign id
func (p *FasilitasUseCase) GetFasilitasSearchByForeignIDPagination(foreignID primitive.ObjectID, page int64, limit int64, keyword string) ([]*domain.Fasilitas, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Fasilitas.SearchByForeignID(ctx, foreignID, page, limit, keyword)
}
