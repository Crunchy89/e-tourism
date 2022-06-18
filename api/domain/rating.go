package domain

import (
	"api/utils/r"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rating struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ForeignID primitive.ObjectID `json:"foreign_id" bson:"foreign_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	Rating    uint8              `json:"rating" bson:"rating,omitempty"`
	Log       *Log               `json:"log" bson:"log,omitempty"`
}
type RatingRepository interface {
	Add(ctx context.Context, d *Rating) (primitive.ObjectID, r.Ex)
	FetchAll(ctx context.Context) ([]*Rating, r.Ex)
	FetchByID(ctx context.Context, ID primitive.ObjectID) (*Rating, r.Ex)
}
