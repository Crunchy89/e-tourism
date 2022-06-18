package usecase

import (
	"context"

	"api/domain"
	"api/utils/r"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KamarUseCase struct {
	Kamar domain.KamarRepository
}

// funciton add kamar
func (p *KamarUseCase) AddKamar(domain *domain.Kamar) (primitive.ObjectID, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Kamar.Add(ctx, domain)
}

// fetch all
func (p *KamarUseCase) FetchAllKamar() ([]*domain.Kamar, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Kamar.FetchAll(ctx)
}

// function to get kamar by id
func (p *KamarUseCase) GetKamarByID(ID primitive.ObjectID) (*domain.Kamar, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Kamar.FetchByID(ctx, ID)
}

// function update by id
func (p *KamarUseCase) UpdateKamarByID(ID primitive.ObjectID, domain *domain.Kamar) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Kamar.UpdateByID(ctx, ID, domain)
}

// fucntion delete
func (p *KamarUseCase) DeleteKamarByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Kamar.DeleteByID(ctx, ID)
}

// function activate
func (p *KamarUseCase) ActivateKamarByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Kamar.Active(ctx, ID)
}
