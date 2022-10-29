package domain

import (
	"api/utils/r"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Travel struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NamaTravel  string             `json:"nama_travel" bson:"nama_travel,omitempty"`
	Pemilik     string             `json:"pemilik" bson:"pemilik,omitempty"`
	Slug        string             `json:"slug" bson:"slug,omitempty"`
	Alamat      string             `json:"alamat" bson:"alamat,omitempty"`
	NoHp        string             `json:"no_hp" bson:"no_hp,omitempty"`
	Email       string             `json:"email" bson:"email,omitempty"`
	Koordinat   *Koordinat         `json:"koordinat" bson:"koordinat,omitempty"`
	MediaSosial *MediaSosial       `json:"media_sosial" bson:"media_sosial,omitempty"`
	Rating      float32            `json:"rating" bson:"rating,omitempty"`
	IsDelete    bool               `json:"is_delete" bson:"is_delete"`
	IsActive    bool               `json:"is_active" bson:"is_active"`
	Log         *Log               `json:"log" bson:"log,omitempty"`
}

type TravelRepository interface {
	Add(ctx context.Context, d *Travel) (primitive.ObjectID, r.Ex)
	FetchAll(ctx context.Context) ([]*Travel, r.Ex)
	FetchByID(ctx context.Context, ID primitive.ObjectID) (*Travel, r.Ex)
	FetchBySlug(ctx context.Context, slug string) (*Travel, r.Ex)
	UpdateByID(ctx context.Context, ID primitive.ObjectID, d *Travel) r.Ex
	DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex
	ActiveByID(ctx context.Context, ID primitive.ObjectID) r.Ex
}
