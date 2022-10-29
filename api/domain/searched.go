package domain

import (
	"api/utils/r"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Searched struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	ForeignID primitive.ObjectID `json:"foreign_id" bson:"foreign_id,omitempty"`
	Search    string             `json:"search" bson:"search"`
	IsDelete  bool               `json:"is_delete" bson:"is_delete"`
	Log       *Log               `json:"log" bson:"log,omitempty"`
}

type SearchedRepository interface {
	Add(ctx context.Context, d *Searched) (primitive.ObjectID, r.Ex)
	FetchAll(ctx context.Context) ([]*Searched, r.Ex)
	FetchByID(ctx context.Context, ID primitive.ObjectID) (*Searched, r.Ex)
	UpdateByID(ctx context.Context, ID primitive.ObjectID, d *Searched) r.Ex
	DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex
}
