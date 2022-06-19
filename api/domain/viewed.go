package domain

import (
	"api/utils/r"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Viewed struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ForeignID primitive.ObjectID `json:"foreign_id" bson:"foreign_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	IsDelete  bool               `json:"is_delete" bson:"is_delete"`
	Log       *Log               `json:"log" bson:"log,omitempty"`
}

type ViewedRepository interface {
	Add(ctx context.Context, d *Viewed) (primitive.ObjectID, r.Ex)
	FetchAll(ctx context.Context) ([]*Viewed, r.Ex)
	FetchByID(ctx context.Context, ID primitive.ObjectID) (*Viewed, r.Ex)
	UpdateByID(ctx context.Context, ID primitive.ObjectID, d *Viewed) r.Ex
	DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex
}
