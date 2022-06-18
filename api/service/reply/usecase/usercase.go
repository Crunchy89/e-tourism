package usecase

import (
	"context"

	"api/domain"
	"api/utils/r"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReplyUseCase struct {
	Reply domain.ReplyRepository
}

// funciton add reply
func (p *ReplyUseCase) AddReply(domain *domain.Reply) (primitive.ObjectID, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Reply.Add(ctx, domain)
}

// fetch all
func (p *ReplyUseCase) FetchAllReply() ([]*domain.Reply, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Reply.FetchAll(ctx)
}

// function to get reply by id
func (p *ReplyUseCase) GetReplyByID(ID primitive.ObjectID) (*domain.Reply, r.Ex) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Reply.FetchByID(ctx, ID)
}

// function update by id
func (p *ReplyUseCase) UpdateReplyByID(ID primitive.ObjectID, domain *domain.Reply) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Reply.UpdateByID(ctx, ID, domain)
}

// fucntion delete
func (p *ReplyUseCase) DeleteReplyByID(ID primitive.ObjectID) r.Ex {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return p.Reply.DeleteByID(ctx, ID)
}
