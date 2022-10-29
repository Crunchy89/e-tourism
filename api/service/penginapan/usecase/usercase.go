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

// funciton add penginapan
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

// function to get penginapan by id
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

// function pagination
func (p *PenginapanUseCase) Pagination(page int64, limit int64) (*domain.PenginapanPagination, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := p.Penginapan.Pagination(ctx, page, limit)
	if err != nil {
		return nil, err
	}
	return p.CreatePenginapanPagination(res, page, limit)
}

// function serach pagination with PenginapanPagination result
func (p *PenginapanUseCase) SearchPagination(page int64, limit int64, search string) (*domain.PenginapanPagination, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := p.Penginapan.Search(ctx, page, limit, search)
	if err != nil {
		return nil, err
	}
	return p.CreatePenginapanPagination(res, page, limit)
}

// function create penginapan pagination
func (p *PenginapanUseCase) CreatePenginapanPagination(data []*domain.Penginapan, page int64, limit int64) (*domain.PenginapanPagination, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	count, err := p.Penginapan.CountAll(ctx)
	if err != nil {
		return nil, err
	}
	result := &domain.PenginapanPagination{
		Data:  data,
		Page:  page,
		Limit: limit,
		Total: count,
	}
	return result, nil
}
