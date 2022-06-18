package domain

import (
	"api/utils/r"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reply struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ReviewID primitive.ObjectID `json:"review_id" bson:"review_id,omitempty"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	Username string             `json:"username" bson:"username,omitempty"`
	Comment  string             `json:"comment" bson:"comment,omitempty"`
	IsDelete bool               `json:"is_delete" bson:"is_delete"`
	Log      *Log               `json:"log" bson:"log,omitempty"`
}

type ReplyRepository interface {
	Add(ctx context.Context, d *Reply) (primitive.ObjectID, r.Ex)
	FetchAll(ctx context.Context) ([]*Reply, r.Ex)
	FetchByID(ctx context.Context, ID primitive.ObjectID) (*Reply, r.Ex)
	UpdateByID(ctx context.Context, ID primitive.ObjectID, d *Reply) r.Ex
	DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex
}
