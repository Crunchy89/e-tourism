package usecase

import (
	"context"

	"api/domain"
	"api/utils/r"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUseCase struct {
	User domain.UserRepository
}

// funciton add user
func (p *UserUseCase) AddUser(domain *domain.User) (primitive.ObjectID, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.User.Add(ctx, domain)
}

// function to get user by id
func (p *UserUseCase) GetUserByID(ID primitive.ObjectID) (*domain.User, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.User.Fetch(ctx, ID)
}

// function to get user by username
func (p *UserUseCase) GetUserByUsername(username string) (*domain.User, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.User.FetchByUsername(ctx, username)
}

// function get user by token
func (p *UserUseCase) GetUserByToken(token string) (*domain.User, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.User.FetchByToken(ctx, token)
}
