package domain

import (
	"api/utils/r"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ForeignID    primitive.ObjectID `json:"foreign_id" bson:"foreign_id,omitempty"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	Username     string             `json:"username" bson:"username,omitempty"`
	TotalComment int                `json:"total_comment" bson:"total_comment,omitempty"`
	IsDelete     bool               `json:"is_delete" bson:"is_delete"`
	Comment      string             `json:"comment" bson:"comment,omitempty"`
	Log          *Log               `json:"log" bson:"log,omitempty"`
}

type ReviewRepository interface {
	Add(ctx context.Context, d *Review) (primitive.ObjectID, r.Ex)
	FetchAll(ctx context.Context) ([]*Review, r.Ex)
	FetchByID(ctx context.Context, ID primitive.ObjectID) (*Review, r.Ex)
	UpdateByID(ctx context.Context, ID primitive.ObjectID, d *Review) r.Ex
	DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex
}
