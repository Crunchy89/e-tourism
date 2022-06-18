package usecase

import (
	"context"

	"api/domain"
	"api/utils/r"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WisataUseCase struct {
	Wisata domain.WisataRepository
}

// funciton add user
func (p *WisataUseCase) AddWisata(domain *domain.Wisata) (primitive.ObjectID, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Wisata.Add(ctx, domain)
}

// fetch all
func (p *WisataUseCase) FetchAllWisata() ([]*domain.Wisata, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Wisata.FetchAll(ctx)
}

// function to get user by id
func (p *WisataUseCase) GetWisataByID(ID primitive.ObjectID) (*domain.Wisata, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Wisata.FetchByID(ctx, ID)
}

// function update by id
func (p *WisataUseCase) UpdateWisataByID(ID primitive.ObjectID, domain *domain.Wisata) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Wisata.UpdateByID(ctx, ID, domain)
}

// fucntion delete
func (p *WisataUseCase) DeleteWisataByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Wisata.DeleteByID(ctx, ID)
}

// function activate
func (p *WisataUseCase) ActivateWisataByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Wisata.Active(ctx, ID)
}
