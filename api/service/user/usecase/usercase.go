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

// function get all user
func (p *UserUseCase) GetAllUser() ([]*domain.User, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.User.FetchAll(ctx)
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

// function pagination
func (p *UserUseCase) Pagination(page int64, limit int64) (*domain.UserPagination, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := p.User.Pagination(ctx, page, limit)
	if err != nil {
		return nil, err
	}
	return p.CreateUserPagination(res, page, limit)
}

// function serach pagination with UserPagination result
func (p *UserUseCase) SearchPagination(page int64, limit int64, search string) (*domain.UserPagination, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := p.User.Search(ctx, page, limit, search)
	if err != nil {
		return nil, err
	}
	return p.CreateUserPagination(res, page, limit)
}

// function create user pagination
func (p *UserUseCase) CreateUserPagination(data []*domain.User, page int64, limit int64) (*domain.UserPagination, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	count, err := p.User.CountAll(ctx)
	if err != nil {
		return nil, err
	}
	result := &domain.UserPagination{
		Data:  data,
		Page:  page,
		Limit: limit,
		Total: count,
	}
	return result, nil
}
