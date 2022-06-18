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
	return p.User.FetchByID(ctx, ID)
}

// function to get user by username
func (p *UserUseCase) GetUserByUsername(username string) (*domain.User, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.User.FetchByUsername(ctx, username)
}

// function get user by email
func (p *UserUseCase) GetUserByEmail(email string) (*domain.User, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.User.FetchByEmail(ctx, email)
}

// function update by id
func (p *UserUseCase) UpdateUserByID(ID primitive.ObjectID, domain *domain.User) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.User.UpdateByID(ctx, ID, domain)
}

// function update by username
func (p *UserUseCase) UpdateUserByUsername(username string, domain *domain.User) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.User.UpdateByUsername(ctx, username, domain)
}

// function update by email
func (p *UserUseCase) UpdateUserByEmail(email string, domain *domain.User) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.User.UpdateByEmail(ctx, email, domain)
}

// function delete
func (p *UserUseCase) DeleteUserByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.User.DeleteByID(ctx, ID)
}

// function active
func (p *UserUseCase) ActiveUserByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.User.ActiveByID(ctx, ID)
}
