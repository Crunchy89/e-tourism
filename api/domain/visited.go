package domain

import (
	"api/utils/r"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Visited struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ForeignID primitive.ObjectID `json:"foreign_id" bson:"foreign_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	IsDelete  bool               `json:"is_delete" bson:"is_delete"`
	Log       *Log               `json:"log" bson:"log,omitempty"`
}

type VisitedRepository interface {
	Add(ctx context.Context, d *Visited) (primitive.ObjectID, r.Ex)
	FetchAll(ctx context.Context) ([]*Visited, r.Ex)
	FetchByID(ctx context.Context, ID primitive.ObjectID) (*Visited, r.Ex)
	UpdateByID(ctx context.Context, ID primitive.ObjectID, d *Visited) r.Ex
	DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex
}
