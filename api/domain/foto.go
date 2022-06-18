package domain

import (
	"api/utils/r"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Foto struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ForeignID primitive.ObjectID `json:"foreign_id" bson:"foreign_id,omitempty"`
	Path      string             `json:"path" bson:"path,omitempty"`
	Thumbnail string             `json:"thumbnail" bson:"thumbnail,omitempty"`
	IsDelete  bool               `json:"is_delete" bson:"is_delete"`
	Log       *Log               `json:"log" bson:"log,omitempty"`
}

type FotoRepository interface {
	Add(ctx context.Context, d *Foto) (primitive.ObjectID, r.Ex)
	FetchAll(ctx context.Context) ([]*Foto, r.Ex)
	FetchByID(ctx context.Context, ID primitive.ObjectID) (*Foto, r.Ex)
	UpdateByID(ctx context.Context, ID primitive.ObjectID, d *Foto) r.Ex
	DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex
}
